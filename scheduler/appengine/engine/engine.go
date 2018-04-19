// Copyright 2015 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package engine implements the core logic of the scheduler service.
package engine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"google.golang.org/api/pubsub/v1"

	ds "go.chromium.org/gae/service/datastore"
	"go.chromium.org/gae/service/info"
	"go.chromium.org/gae/service/memcache"
	"go.chromium.org/gae/service/taskqueue"

	"go.chromium.org/luci/appengine/tq"
	"go.chromium.org/luci/auth/identity"
	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/data/rand/mathrand"
	"go.chromium.org/luci/common/data/stringset"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/proto/google"
	"go.chromium.org/luci/common/retry/transient"
	"go.chromium.org/luci/common/sync/parallel"
	"go.chromium.org/luci/server/auth"
	"go.chromium.org/luci/server/auth/signing"
	"go.chromium.org/luci/server/tokens"

	"go.chromium.org/luci/scheduler/appengine/acl"
	"go.chromium.org/luci/scheduler/appengine/catalog"
	"go.chromium.org/luci/scheduler/appengine/internal"
	"go.chromium.org/luci/scheduler/appengine/task"
)

// TODO(vadimsh): Use annotated errors instead of constants, so they can have
// more information.

var (
	// ErrNoPermission indicates the caller doesn't not have permission to perform
	// desired action, depending on which either either OWNERS or TRIGGERER
	// permission is required.
	ErrNoPermission = errors.New("insufficient rights on a job")
	// ErrNoSuchJob indicates the job doesn't exist or not visible.
	ErrNoSuchJob = errors.New("no such job")
	// ErrNoSuchInvocation indicates the invocation doesn't exist or not visible.
	ErrNoSuchInvocation = errors.New("the invocation doesn't exist")
)

// Engine manages all scheduler jobs: keeps track of their state, runs state
// machine transactions, starts new invocations, etc.
//
// A method returns errors.Transient if the error is non-fatal and the call
// should be retried later. Any other error means that retry won't help.
//
// The general pattern for doing something to a job is to get a reference to
// it via GetVisibleJob() (this call checks READER access), and then pass *Job
// to desired methods (which may additionally check for more permissions).
//
// ACLs are enforced with the following implication:
//  * if caller lacks READER access to Jobs, methods behave as if Jobs do not
//    exist.
//  * if caller lacks TRIGGERER or OWNER access to Jobs, but has READER access,
//    ErrNoPermission will be returned.
//
// Use EngineInternal if you need to skip ACL checks.
type Engine interface {
	// GetVisibleJobs returns all enabled visible jobs.
	//
	// Returns them in no particular order.
	GetVisibleJobs(c context.Context) ([]*Job, error)

	// GetVisibleProjectJobs returns enabled visible jobs belonging to a project.
	//
	// Returns them in no particular order.
	GetVisibleProjectJobs(c context.Context, projectID string) ([]*Job, error)

	// GetVisibleJob returns a single visible job given its full ID.
	//
	// ErrNoSuchJob error is returned if either:
	//   * job doesn't exist,
	//   * job is disabled (i.e. was removed from its project config),
	//   * job isn't visible due to lack of READER access.
	GetVisibleJob(c context.Context, jobID string) (*Job, error)

	// GetVisibleJobBatch is like GetVisibleJob, except it operates on a batch of
	// jobs at once.
	//
	// Returns a mapping (jobID => *Job) with only visible jobs. If the check
	// fails returns a transient error.
	GetVisibleJobBatch(c context.Context, jobIDs []string) (map[string]*Job, error)

	// ListInvocations returns invocations of a given job, sorted by their
	// creation time (most recent first).
	//
	// Can optionally return only active invocations (i.e. ones that are pending,
	// starting or running) or only finished ones. See ListInvocationsOpts.
	//
	// Returns invocations and a cursor string if there's more. Returns only
	// transient errors.
	ListInvocations(c context.Context, job *Job, opts ListInvocationsOpts) ([]*Invocation, string, error)

	// GetInvocation returns an invocation of a given job.
	//
	// ErrNoSuchInvocation is returned if the invocation doesn't exist.
	GetInvocation(c context.Context, job *Job, invID int64) (*Invocation, error)

	// PauseJob prevents new automatic invocations of a job.
	//
	// It clears the pending triggers queue, and makes the job ignore all incoming
	// triggers until it is resumed.
	//
	// For cron jobs it also replaces job's schedule with "triggered", effectively
	// preventing it from running automatically (until it is resumed).
	//
	// Manual invocations (via ForceInvocation) are still allowed. Does nothing if
	// the job is already paused. Any pending or running invocations are still
	// executed.
	PauseJob(c context.Context, job *Job) error

	// ResumeJob resumes paused job. Does nothing if the job is not paused.
	ResumeJob(c context.Context, job *Job) error

	// AbortJob resets the job to scheduled state, aborting all currently pending
	// or running invocations (if any).
	AbortJob(c context.Context, job *Job) error

	// AbortInvocation forcefully moves the invocation to failed state.
	//
	// It opportunistically tries to send "abort" signal to a job runner if it
	// supports cancellation, but it doesn't wait for reply. It proceeds to
	// modifying local state in the scheduler service datastore immediately.
	//
	// AbortInvocation can be used to manually "unstuck" jobs that got stuck due
	// to missing PubSub notifications or other kinds of unexpected conditions.
	//
	// Does nothing if invocation is already in some final state.
	AbortInvocation(c context.Context, job *Job, invID int64) error

	// ForceInvocation launches job invocation right now if job isn't running now.
	//
	// Used by "Run now" UI button.
	//
	// Returns an object that can be waited on to grab a corresponding Invocation
	// when it appears (if ever).
	ForceInvocation(c context.Context, job *Job) (FutureInvocation, error)

	// EmitTriggers puts one or more triggers into pending trigger queues of the
	// specified jobs.
	//
	// If the caller has no permission to trigger at least one job, the entire
	// call is aborted. Otherwise, the call is NOT transactional.
	EmitTriggers(c context.Context, perJob map[*Job][]*internal.Trigger) error

	// ListTriggers returns list of job's pending triggers sorted by time, most
	// recent last.
	ListTriggers(c context.Context, job *Job) ([]*internal.Trigger, error)
}

// EngineInternal is a variant of engine API that skips ACL checks.
//
// Used by the scheduler service guts that executed outside of a context of some
// end user.
type EngineInternal interface {
	// PublicAPI returns ACL-enforcing API.
	PublicAPI() Engine

	// GetAllProjects returns projects that have at least one enabled job.
	GetAllProjects(c context.Context) ([]string, error)

	// UpdateProjectJobs adds new, removes old and updates existing jobs.
	UpdateProjectJobs(c context.Context, projectID string, defs []catalog.Definition) error

	// ResetAllJobsOnDevServer forcefully resets state of all enabled jobs.
	//
	// Supposed to be used only on devserver, where task queue stub state is not
	// preserved between appserver restarts and it messes everything.
	ResetAllJobsOnDevServer(c context.Context) error

	// ExecuteSerializedAction is called via a task queue to execute an action
	// produced by job state machine transition.
	//
	// These actions are POSTed to TimersQueue and InvocationsQueue defined in
	// Config by Engine.
	//
	// 'retryCount' is 0 on first attempt, 1 if task queue service retries
	// request once, 2 - if twice, and so on.
	//
	// Returning transient errors here causes the task queue to retry the task.
	ExecuteSerializedAction(c context.Context, body []byte, retryCount int) error

	// ProcessPubSubPush is called whenever incoming PubSub message is received.
	ProcessPubSubPush(c context.Context, body []byte) error

	// PullPubSubOnDevServer is called on dev server to pull messages from PubSub
	// subscription associated with given publisher.
	//
	// It is needed to be able to manually tests PubSub related workflows on dev
	// server, since dev server can't accept PubSub push messages.
	PullPubSubOnDevServer(c context.Context, taskManagerName, publisher string) error
}

// ListInvocationsOpts are passed to ListInvocations method.
type ListInvocationsOpts struct {
	PageSize     int
	Cursor       string
	FinishedOnly bool
	ActiveOnly   bool
}

// FutureInvocation is returned by ForceInvocation.
//
// It can be used to wait for a triggered invocation to appear.
type FutureInvocation interface {
	// InvocationID returns an ID of the invocation or 0 if not started yet.
	//
	// Returns only transient errors.
	InvocationID(context.Context) (int64, error)
}

// Config contains parameters for the engine.
type Config struct {
	Catalog              catalog.Catalog // provides task.Manager's to run tasks
	Dispatcher           *tq.Dispatcher  // dispatcher for task queue tasks
	TimersQueuePath      string          // URL of a task queue handler for timer ticks
	TimersQueueName      string          // queue name for timer ticks
	InvocationsQueuePath string          // URL of a task queue handler that starts jobs
	InvocationsQueueName string          // queue name for job starts
	PubSubPushPath       string          // URL to use in PubSub push config
}

// NewEngine returns default implementation of EngineInternal.
func NewEngine(cfg Config) EngineInternal {
	eng := &engineImpl{cfg: cfg}
	eng.init()
	return eng
}

type engineImpl struct {
	cfg      Config
	opsCache opsCache

	// configureTopic is used by prepareTopic, mocked in tests.
	configureTopic func(c context.Context, topic, sub, pushURL, publisher string) error
}

// init registers task queue handlers.
func (e *engineImpl) init() {
	// TODO(vadimsh): We probably need some non-default retry policies for all
	// tasks, not just launchInvocationTask.
	e.cfg.Dispatcher.RegisterTask(&internal.LaunchInvocationsBatchTask{}, e.execLaunchInvocationsBatchTask, "batches", nil)
	e.cfg.Dispatcher.RegisterTask(&internal.LaunchInvocationTask{}, e.execLaunchInvocationTask, "launches", &taskqueue.RetryOptions{
		// Give 5 attempts to mark the job as failed. See 'launchInvocationTask'.
		RetryLimit: invocationRetryLimit + 5,
		MinBackoff: time.Second,
		MaxBackoff: maxInvocationRetryBackoff,
		AgeLimit:   time.Duration(invocationRetryLimit+5) * maxInvocationRetryBackoff,
	})
	e.cfg.Dispatcher.RegisterTask(&internal.TriageJobStateTask{}, e.execTriageJobStateTask, "triages", nil)
	e.cfg.Dispatcher.RegisterTask(&internal.KickTriageTask{}, e.execKickTriageTask, "triages", nil)
	e.cfg.Dispatcher.RegisterTask(&internal.InvocationFinishedTask{}, e.execInvocationFinishedTask, "completions", nil)
	e.cfg.Dispatcher.RegisterTask(&internal.FanOutTriggersTask{}, e.execFanOutTriggersTask, "triggers", nil)
	e.cfg.Dispatcher.RegisterTask(&internal.EnqueueTriggersTask{}, e.execEnqueueTriggersTask, "triggers", nil)
	e.cfg.Dispatcher.RegisterTask(&internal.ScheduleTimersTask{}, e.execScheduleTimersTask, "timers", nil)
	e.cfg.Dispatcher.RegisterTask(&internal.TimerTask{}, e.execTimerTask, "timers", nil)
	e.cfg.Dispatcher.RegisterTask(&internal.CronTickTask{}, e.execCronTickTask, "crons", nil)
}

// jobController is a part of engine that directly deals with state transitions
// of a single job and its invocations.
//
// It is short-lived. It is instantiated, used and discarded. Never retained.
//
// All onJob* methods are called from within a Job transaction.
// All onInv* methods are called from within an Invocation transaction.
//
// There are two implementations of the controller (jobControllerV1 and
// jobControllerV2).
type jobController interface {
	onJobScheduleChange(c context.Context, job *Job) error
	onJobEnabled(c context.Context, job *Job) error
	onJobDisabled(c context.Context, job *Job) error
	onJobCronTick(c context.Context, job *Job, tick *internal.CronTickTask) error
	onJobAbort(c context.Context, job *Job) (invs []int64, err error)
	onJobForceInvocation(c context.Context, job *Job) (FutureInvocation, error)

	onInvUpdating(c context.Context, old, fresh *Invocation, timers []*internal.Timer, triggers []*internal.Trigger) error
}

////////////////////////////////////////////////////////////////////////////////
// Engine interface implementation.

// GetVisibleJobs returns all enabled visible jobs.
//
// Part of the public interface, checks ACLs.
func (e *engineImpl) GetVisibleJobs(c context.Context) ([]*Job, error) {
	q := ds.NewQuery("Job").Eq("Enabled", true)
	return e.queryEnabledVisibleJobs(c, q)
}

// GetVisibleProjectJobs enabled visible jobs belonging to a project.
//
// Part of the public interface, checks ACLs.
func (e *engineImpl) GetVisibleProjectJobs(c context.Context, projectID string) ([]*Job, error) {
	q := ds.NewQuery("Job").Eq("Enabled", true).Eq("ProjectID", projectID)
	return e.queryEnabledVisibleJobs(c, q)
}

// GetVisibleJob returns a single visible job given its full ID.
//
// Part of the public interface, checks ACLs.
func (e *engineImpl) GetVisibleJob(c context.Context, jobID string) (*Job, error) {
	job, err := e.getJob(c, jobID)
	switch {
	case err != nil:
		return nil, err
	case job == nil || !job.Enabled:
		return nil, ErrNoSuchJob
	}
	if err := job.CheckRole(c, acl.Reader); err != nil {
		if err == ErrNoPermission {
			err = ErrNoSuchJob // pretend protected jobs don't exist
		}
		return nil, err
	}
	return job, nil
}

// GetVisibleJobBatch is like GetVisibleJob, except it operates on a batch of
// jobs at once.
//
// Part of the public interface.
//
// Part of the public interface, checks ACLs.
func (e *engineImpl) GetVisibleJobBatch(c context.Context, jobIDs []string) (map[string]*Job, error) {
	// TODO(vadimsh): This can be parallelized to be single GetMulti RPC to fetch
	// jobs and single filterForRole to check ACLs. In practice O(len(jobIDs)) is
	// small, so there's no pressing need to do this.
	visible := make(map[string]*Job, len(jobIDs))
	for _, id := range jobIDs {
		switch job, err := e.GetVisibleJob(c, id); {
		case err == nil:
			visible[id] = job
		case err != ErrNoSuchJob:
			return nil, err
		}
	}
	return visible, nil
}

// ListInvocations returns invocations of a given job, most recent first.
//
// Part of the public interface.
//
// Supports both v1 and v2 invocations.
func (e *engineImpl) ListInvocations(c context.Context, job *Job, opts ListInvocationsOpts) ([]*Invocation, string, error) {
	if opts.ActiveOnly && opts.FinishedOnly {
		return nil, "", fmt.Errorf("using both ActiveOnly and FinishedOnly is not allowed")
	}

	if opts.PageSize <= 0 || opts.PageSize > 500 {
		opts.PageSize = 500
	}

	var cursor internal.InvocationsCursor
	if err := decodeInvCursor(opts.Cursor, &cursor); err != nil {
		return nil, "", err
	}

	// We are going to merge results of multiple queries:
	//   1) Over historical finished invocations in the datastore.
	//   2) Over recently finished invocations, stored inline in the Job entity.
	//   3) Over active invocations, also stored inline in the Job entity.
	var qs []invQuery
	if !opts.ActiveOnly {
		// Most of the historical invocations came from the datastore query. But it
		// may not have recently finished invocations yet (due to Datastore eventual
		// consistently).
		q := finishedInvQuery(c, job, cursor.LastScanned)
		defer q.close()
		qs = append(qs, q)
		// Use recently finished invocations from the Job, since they may be more
		// up-to-date and do not depend on Datastore index consistency lag.
		qs = append(qs, recentInvQuery(c, job, cursor.LastScanned))
	}
	if !opts.FinishedOnly {
		qs = append(qs, activeInvQuery(c, job, cursor.LastScanned))
	}

	out := make([]*Invocation, 0, opts.PageSize)

	// Build the full page out of potentially incomplete (due to post-filtering)
	// smaller pages. Note that most of the time 'fetchInvsPage' will return the
	// full page right away.
	var page invsPage
	var err error
	for opts.PageSize > 0 {
		out, page, err = fetchInvsPage(c, qs, opts, out)
		switch {
		case err != nil:
			return nil, "", err
		case page.final:
			return out, "", nil // return empty cursor to indicate we are done
		}
		opts.PageSize -= page.count
	}

	// We end up here if the last fetched mini-page wasn't final, need new cursor.
	cursorStr, err := encodeInvCursor(&internal.InvocationsCursor{
		LastScanned: page.lastScanned,
	})
	if err != nil {
		return nil, "", errors.Annotate(err, "failed to serialize the cursor").Err()
	}
	return out, cursorStr, nil
}

// GetInvocation returns some invocation of a given job.
//
// Part of the public interface.
//
// Supports both v1 and v2 invocations.
func (e *engineImpl) GetInvocation(c context.Context, job *Job, invID int64) (*Invocation, error) {
	// Note: we want public API users to go through GetVisibleJob to check ACLs,
	// thus usage of *Job, even though JobID string is sufficient in this case.
	return e.getInvocation(c, job.JobID, invID)
}

// PauseJob prevents new automatic invocations of a job.
//
// Part of the public interface, checks ACLs.
func (e *engineImpl) PauseJob(c context.Context, job *Job) error {
	if err := job.CheckRole(c, acl.Owner); err != nil {
		return err
	}
	return e.setJobPausedFlag(c, job, true, auth.CurrentIdentity(c))
}

// ResumeJob resumes paused job. Does nothing if the job is not paused.
//
// Part of the public interface, checks ACLs.
func (e *engineImpl) ResumeJob(c context.Context, job *Job) error {
	if err := job.CheckRole(c, acl.Owner); err != nil {
		return err
	}
	return e.setJobPausedFlag(c, job, false, auth.CurrentIdentity(c))
}

// AbortJob resets the job to scheduled state, aborting all currently pending
// or running invocations (if any).
//
// Part of the public interface, checks ACLs.
func (e *engineImpl) AbortJob(c context.Context, job *Job) error {
	// First, we check ACLs.
	if err := job.CheckRole(c, acl.Owner); err != nil {
		return err
	}
	jobID := job.JobID

	// Second, we switch the job to the default state and disassociate the running
	// invocations (if any) from the job entity.
	var invs []int64
	err := e.jobTxn(c, jobID, func(c context.Context, job *Job, isNew bool) (err error) {
		if isNew {
			return errSkipPut // the job was removed, nothing to abort
		}
		invs, err = e.jobController(jobID).onJobAbort(c, job)
		return err
	})
	if err != nil {
		return err
	}

	// Now we kill the invocations. We do it separately because it may involve
	// an RPC to remote service (e.g. to cancel a task) that can't be done from
	// the transaction.
	wg := sync.WaitGroup{}
	wg.Add(len(invs))
	errs := errors.NewLazyMultiError(len(invs))
	for i, invID := range invs {
		go func(i int, invID int64) {
			defer wg.Done()
			errs.Assign(i, e.abortInvocation(c, jobID, invID))
		}(i, invID)
	}
	wg.Wait()
	if err := errs.Get(); err != nil {
		return transient.Tag.Apply(err)
	}
	return nil
}

// AbortInvocation forcefully moves the invocation to failed state.
//
// Part of the public interface, checks ACLs.
func (e *engineImpl) AbortInvocation(c context.Context, job *Job, invID int64) error {
	if err := job.CheckRole(c, acl.Owner); err != nil {
		return err
	}
	return e.abortInvocation(c, job.JobID, invID)
}

// ForceInvocation launches job invocation right now if job isn't running now.
//
// Part of the public interface, checks ACLs.
func (e *engineImpl) ForceInvocation(c context.Context, job *Job) (FutureInvocation, error) {
	if err := job.CheckRole(c, acl.Triggerer); err != nil {
		return nil, err
	}

	var noSuchJob bool
	var future FutureInvocation
	err := e.jobTxn(c, job.JobID, func(c context.Context, job *Job, isNew bool) (err error) {
		if isNew || !job.Enabled {
			noSuchJob = true
			return errSkipPut
		}
		future, err = e.jobController(job.JobID).onJobForceInvocation(c, job)
		return err
	})

	switch {
	case noSuchJob:
		return nil, ErrNoSuchJob
	case err != nil:
		return nil, err
	}

	return future, nil
}

// EmitTriggers puts one or more triggers into pending trigger queues of the
// specified jobs.
func (e *engineImpl) EmitTriggers(c context.Context, perJob map[*Job][]*internal.Trigger) error {
	// Make sure the caller has permissions to add triggers to all jobs.
	jobs := make([]*Job, 0, len(perJob))
	for j := range perJob {
		jobs = append(jobs, j)
	}
	switch filtered, err := e.filterForRole(c, jobs, acl.Triggerer); {
	case err != nil:
		return errors.Annotate(err, "transient error when checking Triggerer role").Err()
	case len(filtered) != len(jobs):
		return ErrNoPermission // some jobs are not triggerable
	}

	// Actually trigger.
	return parallel.FanOutIn(func(tasks chan<- func() error) {
		for job, triggers := range perJob {
			jobID := job.JobID
			triggers := triggers
			if e.isV2Job(jobID) {
				tasks <- func() error {
					return e.execEnqueueTriggersTask(c, &internal.EnqueueTriggersTask{
						JobId:    jobID,
						Triggers: triggers,
					})
				}
			} else {
				tasks <- func() error { return e.newTriggers(c, jobID, triggers) }
			}
		}
	})
}

// ListTriggers returns sorted list of job's pending triggers.
//
// Supports both v1 and v2.
func (e *engineImpl) ListTriggers(c context.Context, job *Job) ([]*internal.Trigger, error) {
	var triggers []*internal.Trigger
	var err error

	if e.isV2Job(job.JobID) {
		// v2 implementation stores triggers in a dsset.
		_, triggers, err = pendingTriggersSet(c, job.JobID).Triggers(c)
	} else {
		// v1 implementation stores triggers inside Job entity itself.
		triggers, err = unmarshalTriggersList(job.State.PendingTriggersRaw)
	}

	if err != nil {
		return nil, transient.Tag.Apply(err)
	}
	sortTriggers(triggers)
	return triggers, nil
}

////////////////////////////////////////////////////////////////////////////////
// EngineInternal interface implementation.

// PublicAPI returns ACL-enforcing API.
func (e *engineImpl) PublicAPI() Engine {
	return e
}

// GetAllProjects returns projects that have at least one enabled job.
//
// Part of the internal interface, doesn't check ACLs.
func (e *engineImpl) GetAllProjects(c context.Context) ([]string, error) {
	q := ds.NewQuery("Job").
		Eq("Enabled", true).
		Project("ProjectID").
		Distinct(true)
	entities := []Job{}
	if err := ds.GetAll(c, q, &entities); err != nil {
		return nil, transient.Tag.Apply(err)
	}
	// Filter out duplicates, sort.
	projects := stringset.New(len(entities))
	for _, ent := range entities {
		projects.Add(ent.ProjectID)
	}
	out := projects.ToSlice()
	sort.Strings(out)
	return out, nil
}

// UpdateProjectJobs adds new, removes old and updates existing jobs.
//
// Part of the internal interface, doesn't check ACLs.
func (e *engineImpl) UpdateProjectJobs(c context.Context, projectID string, defs []catalog.Definition) error {
	// JobID -> *Job map.
	existing, err := e.getProjectJobs(c, projectID)
	if err != nil {
		return err
	}
	// JobID -> new definition revision map.
	updated := make(map[string]string, len(defs))
	for _, def := range defs {
		updated[def.JobID] = def.Revision
	}
	// List of job ids to disable.
	var toDisable []string
	for id := range existing {
		if updated[id] == "" {
			toDisable = append(toDisable, id)
		}
	}

	wg := sync.WaitGroup{}

	// Add new jobs, update existing ones.
	updateErrs := errors.NewLazyMultiError(len(defs))
	for i, def := range defs {
		if ent := existing[def.JobID]; ent != nil {
			if ent.Enabled && ent.MatchesDefinition(def) {
				continue
			}
		}
		wg.Add(1)
		go func(i int, def catalog.Definition) {
			updateErrs.Assign(i, e.updateJob(c, def))
			wg.Done()
		}(i, def)
	}

	// Disable old jobs.
	disableErrs := errors.NewLazyMultiError(len(toDisable))
	for i, jobID := range toDisable {
		wg.Add(1)
		go func(i int, jobID string) {
			disableErrs.Assign(i, e.disableJob(c, jobID))
			wg.Done()
		}(i, jobID)
	}

	wg.Wait()
	if updateErrs.Get() == nil && disableErrs.Get() == nil {
		return nil
	}
	return transient.Tag.Apply(errors.NewMultiError(updateErrs.Get(), disableErrs.Get()))
}

// ResetAllJobsOnDevServer forcefully resets state of all enabled jobs.
//
// Part of the internal interface, doesn't check ACLs.
func (e *engineImpl) ResetAllJobsOnDevServer(c context.Context) error {
	if !info.IsDevAppServer(c) {
		return errors.New("ResetAllJobsOnDevServer must not be used in production")
	}
	q := ds.NewQuery("Job").Eq("Enabled", true)
	keys := []*ds.Key{}
	if err := ds.GetAll(c, q, &keys); err != nil {
		return transient.Tag.Apply(err)
	}
	wg := sync.WaitGroup{}
	errs := errors.NewLazyMultiError(len(keys))
	for i, key := range keys {
		wg.Add(1)
		go func(i int, key *ds.Key) {
			errs.Assign(i, e.resetJobOnDevServer(c, key.StringID()))
			wg.Done()
		}(i, key)
	}
	wg.Wait()
	return transient.Tag.Apply(errs.Get())
}

// ExecuteSerializedAction is called via a task queue to execute an action
// produced by job state machine transition.
//
// Part of the internal interface, doesn't check ACLs.
func (e *engineImpl) ExecuteSerializedAction(c context.Context, action []byte, retryCount int) error {
	payload := actionTaskPayload{}
	if err := json.Unmarshal(action, &payload); err != nil {
		return err
	}
	if payload.InvID == 0 {
		return e.executeJobAction(c, &payload, retryCount)
	}
	return e.executeInvAction(c, &payload, retryCount)
}

// ProcessPubSubPush is called whenever incoming PubSub message is received.
//
// Part of the internal interface, doesn't check ACLs.
func (e *engineImpl) ProcessPubSubPush(c context.Context, body []byte) error {
	var pushBody struct {
		Message pubsub.PubsubMessage `json:"message"`
	}
	if err := json.Unmarshal(body, &pushBody); err != nil {
		return err
	}
	return e.handlePubSubMessage(c, &pushBody.Message)
}

// PullPubSubOnDevServer is called on dev server to pull messages from PubSub
// subscription associated with given publisher.
//
// Part of the internal interface, doesn't check ACLs.
func (e *engineImpl) PullPubSubOnDevServer(c context.Context, taskManagerName, publisher string) error {
	_, sub := e.genTopicAndSubNames(c, taskManagerName, publisher)
	msg, ack, err := pullSubcription(c, sub, "")
	if err != nil {
		return err
	}
	if msg == nil {
		logging.Infof(c, "No new PubSub messages")
		return nil
	}
	err = e.handlePubSubMessage(c, msg)
	if err == nil || !transient.Tag.In(err) {
		ack() // ack only on success of fatal errors (to stop redelivery)
	}
	return err
}

////////////////////////////////////////////////////////////////////////////////
// Job related methods.

// txnCallback is passed to 'txn' and it modifies 'job' in place. 'txn' then
// puts it into datastore. The callback may return errSkipPut to instruct 'txn'
// not to call datastore 'Put'. The callback may do other transactional things
// using the context.
type txnCallback func(c context.Context, job *Job, isNew bool) error

// errSkipPut can be returned by txnCallback to cancel ds.Put call.
var errSkipPut = errors.New("errSkipPut")

// jobTxn reads Job entity, calls the callback, then dumps the modified entity
// back into datastore (unless the callback returns errSkipPut).
func (e *engineImpl) jobTxn(c context.Context, jobID string, callback txnCallback) error {
	c = logging.SetField(c, "JobID", jobID)
	return runTxn(c, func(c context.Context) error {
		stored := Job{JobID: jobID}
		err := ds.Get(c, &stored)
		if err != nil && err != ds.ErrNoSuchEntity {
			return transient.Tag.Apply(err)
		}
		modified := stored // make a copy of Job struct
		switch err = callback(c, &modified, err == ds.ErrNoSuchEntity); {
		case err == errSkipPut:
			return nil // asked to skip the update
		case err != nil:
			return err // a real error (transient or fatal)
		case !modified.IsEqual(&stored):
			return transient.Tag.Apply(ds.Put(c, &modified))
		}
		return nil
	})
}

// rollSM is called under transaction to perform a single state machine step.
//
// It sets up StateMachine instance, calls the callback, mutates job.State in
// place (with a new state) and enqueues all emitted actions to task queues.
func (e *engineImpl) rollSM(c context.Context, job *Job, cb func(*StateMachine) error) error {
	assertInTransaction(c)
	sched, err := job.ParseSchedule()
	if err != nil {
		return fmt.Errorf("bad schedule %q - %s", job.EffectiveSchedule(), err)
	}
	now := clock.Now(c).UTC()
	rnd := mathrand.Get(c)
	sm := StateMachine{
		State:    job.State,
		Now:      now,
		Schedule: sched,
		Nonce:    func() int64 { return rnd.Int63() + 1 },
		Context:  c,
	}
	// All errors returned by state machine transition changes are transient.
	// Fatal errors (when we have them) should be reflected as a state changing
	// into "BROKEN" state.
	if err := sm.Pre(); err != nil {
		return err
	}
	if err := cb(&sm); err != nil {
		return transient.Tag.Apply(err)
	}
	sm.Post()
	if len(sm.Actions) != 0 {
		if err := e.enqueueJobActions(c, job.JobID, sm.Actions); err != nil {
			return err
		}
	}
	if sm.State.State != job.State.State {
		logging.Infof(c, "%s -> %s", job.State.State, sm.State.State)
	}
	job.State = sm.State
	return nil
}

// isV2Job returns true if the given job is using v2 scheduler engine.
func (e *engineImpl) isV2Job(jobID string) bool {
	return strings.HasSuffix(jobID, "-v2")
}

// jobController returns an appropriate implementation of the jobController
// depending on a version of the engine the job is using (v1 or v2).
func (e *engineImpl) jobController(jobID string) jobController {
	if e.isV2Job(jobID) {
		return &jobControllerV2{eng: e}
	}
	return &jobControllerV1{eng: e}
}

// getJob returns a job if it exists or nil if not.
//
// Doesn't check ACLs.
func (e *engineImpl) getJob(c context.Context, jobID string) (*Job, error) {
	job := &Job{JobID: jobID}
	switch err := ds.Get(c, job); {
	case err == nil:
		return job, nil
	case err == ds.ErrNoSuchEntity:
		return nil, nil
	default:
		return nil, transient.Tag.Apply(err)
	}
}

// getProjectJobs fetches from ds all enabled jobs belonging to a given
// project.
func (e *engineImpl) getProjectJobs(c context.Context, projectID string) (map[string]*Job, error) {
	q := ds.NewQuery("Job").
		Eq("Enabled", true).
		Eq("ProjectID", projectID)
	entities := []*Job{}
	if err := ds.GetAll(c, q, &entities); err != nil {
		return nil, transient.Tag.Apply(err)
	}
	out := make(map[string]*Job, len(entities))
	for _, job := range entities {
		if job.Enabled && job.ProjectID == projectID {
			out[job.JobID] = job
		}
	}
	return out, nil
}

// queryEnabledVisibleJobs fetches all jobs from the query and keeps only ones
// that are enabled and visible by the current caller.
func (e *engineImpl) queryEnabledVisibleJobs(c context.Context, q *ds.Query) ([]*Job, error) {
	entities := []*Job{}
	if err := ds.GetAll(c, q, &entities); err != nil {
		return nil, transient.Tag.Apply(err)
	}
	// Non-ancestor query used, need to recheck filters.
	enabled := make([]*Job, 0, len(entities))
	for _, job := range entities {
		if job.Enabled {
			enabled = append(enabled, job)
		}
	}
	// Keep only ones visible to the caller.
	return e.filterForRole(c, enabled, acl.Reader)
}

// filterForRole returns jobs for which caller has the given role.
//
// May return transient errors.
func (e *engineImpl) filterForRole(c context.Context, jobs []*Job, role acl.Role) ([]*Job, error) {
	// TODO(tandrii): improve batch ACLs check here to take advantage of likely
	// shared ACLs between most jobs of the same project.
	filtered := make([]*Job, 0, len(jobs))
	for _, job := range jobs {
		switch err := job.CheckRole(c, role); {
		case err == nil:
			filtered = append(filtered, job)
		case err != ErrNoPermission:
			return nil, err // a transient error when checking
		}
	}
	return filtered, nil
}

// setJobPausedFlag is implementation of PauseJob/ResumeJob.
//
// Doesn't check ACLs, assumes the check was done already.
func (e *engineImpl) setJobPausedFlag(c context.Context, job *Job, paused bool, who identity.Identity) error {
	return e.jobTxn(c, job.JobID, func(c context.Context, job *Job, isNew bool) error {
		switch {
		case isNew || !job.Enabled:
			return ErrNoSuchJob
		case job.Paused == paused:
			return errSkipPut
		}
		if paused {
			logging.Warningf(c, "Job is paused by %s", who)
		} else {
			logging.Warningf(c, "Job is resumed by %s", who)
		}
		job.Paused = paused
		return e.jobController(job.JobID).onJobScheduleChange(c, job)
	})
}

// updateJob updates an existing job if its definition has changed, adds
// a completely new job or enables a previously disabled job.
func (e *engineImpl) updateJob(c context.Context, def catalog.Definition) error {
	return e.jobTxn(c, def.JobID, func(c context.Context, job *Job, isNew bool) error {
		if !isNew && job.Enabled && job.MatchesDefinition(def) {
			return errSkipPut
		}
		if isNew {
			// JobID is <projectID>/<name>, it's ensured by Catalog.
			chunks := strings.Split(def.JobID, "/")
			if len(chunks) != 2 {
				return fmt.Errorf("unexpected jobID format: %s", def.JobID)
			}
			*job = Job{
				JobID:           def.JobID,
				ProjectID:       chunks[0],
				Flavor:          def.Flavor,
				Enabled:         false, // to trigger 'if !oldEnabled' below
				Schedule:        def.Schedule,
				Task:            def.Task,
				State:           JobState{State: JobStateDisabled},
				TriggeredJobIDs: def.TriggeredJobIDs,
			}
		}
		oldEnabled := job.Enabled
		oldEffectiveSchedule := job.EffectiveSchedule()

		// Update the job in full before running any state changes.
		job.Flavor = def.Flavor
		job.Revision = def.Revision
		job.RevisionURL = def.RevisionURL
		job.Acls = def.Acls
		job.Enabled = true
		job.Schedule = def.Schedule
		job.Task = def.Task
		job.TriggeredJobIDs = def.TriggeredJobIDs

		// Kick off task queue tasks.
		ctl := e.jobController(job.JobID)
		if !oldEnabled {
			if err := ctl.onJobEnabled(c, job); err != nil {
				return err
			}
		}
		if job.EffectiveSchedule() != oldEffectiveSchedule {
			logging.Infof(c, "Job's schedule changed: %q -> %q",
				job.EffectiveSchedule(), oldEffectiveSchedule)
			if err := ctl.onJobScheduleChange(c, job); err != nil {
				return err
			}
		}
		return nil
	})
}

// disableJob moves a job to disabled state.
func (e *engineImpl) disableJob(c context.Context, jobID string) error {
	return e.jobTxn(c, jobID, func(c context.Context, job *Job, isNew bool) error {
		if isNew || !job.Enabled {
			return errSkipPut
		}
		job.Enabled = false
		return e.jobController(jobID).onJobDisabled(c, job)
	})
}

// resetJobOnDevServer sends "off" signal followed by "on" signal.
//
// It effectively cancels any pending actions and schedules new ones. Used only
// on dev server.
func (e *engineImpl) resetJobOnDevServer(c context.Context, jobID string) error {
	return e.jobTxn(c, jobID, func(c context.Context, job *Job, isNew bool) error {
		if isNew || !job.Enabled {
			return errSkipPut
		}
		logging.Infof(c, "Resetting job")
		ctl := e.jobController(jobID)
		if err := ctl.onJobDisabled(c, job); err != nil {
			return err
		}
		return ctl.onJobEnabled(c, job)
	})
}

////////////////////////////////////////////////////////////////////////////////
// Invocations related methods.

// getInvocation returns an existing invocation or ErrNoSuchInvocation error.
//
// Supports both v1 and v2 invocations.
func (e *engineImpl) getInvocation(c context.Context, jobID string, invID int64) (*Invocation, error) {
	isV2 := e.isV2Job(jobID)
	inv := &Invocation{ID: invID}
	if !isV2 {
		inv.JobKey = ds.NewKey(c, "Job", jobID, 0, nil)
	}
	switch err := ds.Get(c, inv); {
	case err == nil:
		if isV2 && inv.JobID != jobID {
			logging.Errorf(c,
				"Invocation %d is associated with job %q, not %q. Treating it as missing.",
				invID, inv.JobID, jobID)
			return nil, ErrNoSuchInvocation
		}
		return inv, nil
	case err == ds.ErrNoSuchEntity:
		return nil, ErrNoSuchInvocation
	default:
		return nil, transient.Tag.Apply(err)
	}
}

// enqueueInvocations allocated a bunch of Invocation entities, adds them to
// ActiveInvocations list of the job and enqueues a tq task that kicks off their
// execution.
//
// Must be called within a Job transaction, but creates Invocation entities
// outside the transaction (since they are in different entity groups). If the
// transaction fails, these entities may keep hanging unreferenced by anything
// as garbage. This is fine, since they are not discoverable by any queries.
func (e *engineImpl) enqueueInvocations(c context.Context, job *Job, req []task.Request) ([]*Invocation, error) {
	assertInTransaction(c)

	// Create N new Invocation entities in Starting state.
	invs, err := e.allocateInvocations(c, job, req)
	if err != nil {
		return nil, err
	}

	// Enqueue a task that eventually calls 'launchInvocationTask' for each new
	// invocation.
	invIDs := make([]int64, len(invs))
	for i, inv := range invs {
		invIDs[i] = inv.ID
	}
	if err := e.kickLaunchInvocationsBatchTask(c, job.JobID, invIDs); err != nil {
		cleanupUnreferencedInvocations(c, invs)
		return nil, err
	}

	// Make the job know that there are invocations pending. This will make them
	// show up in UI and API after the current transaction lands. If it doesn't
	// land, new invocations will remain hanging as garbage, not referenced by
	// anything.
	job.ActiveInvocations = append(job.ActiveInvocations, invIDs...)
	return invs, nil
}

// allocateInvocation creates new Invocation entity in a separate transaction.
//
// Supports only v2 invocations!
func (e *engineImpl) allocateInvocation(c context.Context, job *Job, req task.Request) (*Invocation, error) {
	var inv *Invocation
	err := runIsolatedTxn(c, func(c context.Context) (err error) {
		inv, err = e.initInvocation(c, job.JobID, &Invocation{
			Started:         clock.Now(c).UTC(),
			Revision:        job.Revision,
			RevisionURL:     job.RevisionURL,
			Task:            job.Task,
			TriggeredJobIDs: job.TriggeredJobIDs,
			Status:          task.StatusStarting,
		}, &req)
		if err != nil {
			return
		}
		// TODO(vadimsh): Remove once InvocationNonce is gone. We need it for now
		// since task controller use InvocationNonce as dedup key.
		inv.InvocationNonce = inv.ID
		inv.debugLog(c, "New invocation initialized")
		if req.TriggeredBy != "" {
			inv.debugLog(c, "Triggered by %s", req.TriggeredBy)
		}
		return transient.Tag.Apply(ds.Put(c, inv))
	})
	if err != nil {
		return nil, err
	}
	return inv, nil
}

// allocateInvocations is a batch version of allocateInvocation.
//
// It launches N independent transactions in parallel to create N invocations.
func (e *engineImpl) allocateInvocations(c context.Context, job *Job, req []task.Request) ([]*Invocation, error) {
	wg := sync.WaitGroup{}
	wg.Add(len(req))

	invs := make([]*Invocation, len(req))
	merr := errors.NewLazyMultiError(len(req))
	for i := range req {
		go func(i int) {
			defer wg.Done()
			inv, err := e.allocateInvocation(c, job, req[i])
			invs[i] = inv
			merr.Assign(i, err)
			if err != nil {
				logging.WithError(err).Errorf(c, "Failed to create invocation with %d triggers", len(req[i].IncomingTriggers))
			}
		}(i)
	}

	wg.Wait()

	// Bail if any of them failed. Try best effort cleanup.
	if err := merr.Get(); err != nil {
		cleanupUnreferencedInvocations(c, invs)
		return nil, transient.Tag.Apply(err)
	}

	return invs, nil
}

// initInvocation populates fields of Invocation struct.
//
// It allocates invocation ID and populates related fields: ID, JobKey, JobID.
// It also copies data from given task.Request object into corresponding fields
// of the invocation (so they can be indexed etc).
//
// On success returns exact same 'inv' for convenience. It doesn't Put it into
// the datastore.
//
// Must be called within a transaction, since it verifies an allocated ID is
// not used yet.
//
// Supports both v1 and v2 invocations.
func (e *engineImpl) initInvocation(c context.Context, jobID string, inv *Invocation, req *task.Request) (*Invocation, error) {
	assertInTransaction(c)
	isV2 := e.isV2Job(jobID)
	var jobKey *ds.Key
	if !isV2 {
		jobKey = ds.NewKey(c, "Job", jobID, 0, nil)
	}
	invID, err := generateInvocationID(c, jobKey)
	if err != nil {
		return nil, errors.Annotate(err, "failed to generate invocation ID").Err()
	}
	inv.ID = invID
	if isV2 {
		inv.JobID = jobID
	} else {
		inv.JobKey = jobKey
	}
	if req != nil {
		if err := putRequestIntoInv(inv, req); err != nil {
			return nil, errors.Annotate(err, "failed to serialize task request").Err()
		}
		if req.DebugLog != "" {
			inv.DebugLog += "Debug output from the triage procedure:\n"
			inv.DebugLog += "---------------------------------------\n"
			inv.DebugLog += req.DebugLog
			inv.DebugLog += "---------------------------------------\n\n"
			inv.trimDebugLog() // in case it is HUGE
			inv.DebugLog += "Debug output from the invocation itself:\n"
		}
	}
	return inv, nil
}

// abortInvocation marks some invocation as aborted.
//
// Supports both v1 and v2 invocations.
func (e *engineImpl) abortInvocation(c context.Context, jobID string, invID int64) error {
	return e.withController(c, jobID, invID, "manual abort", func(c context.Context, ctl *taskController) error {
		ctl.DebugLog("Invocation is manually aborted by %s", auth.CurrentIdentity(c))
		if err := ctl.manager.AbortTask(c, ctl); err != nil {
			logging.WithError(err).Errorf(c, "Failed to abort the task")
			return err
		}
		ctl.State().Status = task.StatusAborted
		return nil
	})
}

////////////////////////////////////////////////////////////////////////////////
// Job related task queue messages and routing.

// actionTaskPayload is payload for task queue jobs emitted by the engine.
//
// Serialized as JSON, produced by enqueueJobActions, enqueueInvTimers, and
// enqueueTriggers, used as inputs in ExecuteSerializedAction.
//
// Union of all possible payloads for simplicity.
type actionTaskPayload struct {
	JobID string `json:",omitempty"` // ID of relevant Job
	InvID int64  `json:",omitempty"` // ID of relevant Invocation

	// For Job actions and timers (InvID == 0).
	Kind                string           `json:",omitempty"` // defines what fields below to examine
	TickNonce           int64            `json:",omitempty"` // valid for "TickLaterAction" kind
	InvocationNonce     int64            `json:",omitempty"` // valid for "StartInvocationAction" kind
	TriggeredBy         string           `json:",omitempty"` // valid for "StartInvocationAction" kind
	Triggers            triggersJSONList `json:",omitempty"` // valid for "StartInvocationAction", "EnqueueTriggersAction", and "EnqueueBatchOfTriggersAction"
	TriggeredJobIDs     []string         `json:",omitempty"` // Valid for "EnqueueBatchOfTriggersAction" kind only.
	Overruns            int              `json:",omitempty"` // valid for "RecordOverrunAction" kind
	RunningInvocationID int64            `json:",omitempty"` // valid for "RecordOverrunAction" kind

	// For Invocation actions and timers (InvID != 0).
	InvTimer *invocationTimer `json:",omitempty"` // used for AddTimer calls
}

// triggersJSONList is JSON-serializable list of triggers.
type triggersJSONList []*internal.Trigger

var (
	jsonPBMarshaller   = &jsonpb.Marshaler{}
	jsonPBUnmarshaller = &jsonpb.Unmarshaler{AllowUnknownFields: true}
)

func (l triggersJSONList) MarshalJSON() ([]byte, error) {
	out, err := jsonPBMarshaller.MarshalToString(&internal.TriggerList{Triggers: l})
	if err != nil {
		return nil, err
	}
	return []byte(out), nil
}

func (l *triggersJSONList) UnmarshalJSON(data []byte) error {
	list := internal.TriggerList{}
	err := jsonPBUnmarshaller.Unmarshal(bytes.NewReader(data), &list)
	if err != nil {
		return err
	}
	*l = list.Triggers
	return nil
}

// invocationTimer is carried as part of task queue task payload for tasks
// created by AddTimer calls.
//
// It will be serialized to JSON, so all fields are public.
type invocationTimer struct {
	Delay   time.Duration
	Name    string
	Payload []byte
}

// enqueueJobActions submits all actions emitted by a job state transition by
// adding corresponding tasks to task queues.
//
// See executeJobAction for a place where these actions are interpreted.
func (e *engineImpl) enqueueJobActions(c context.Context, jobID string, actions []Action) error {
	// AddMulti can't put tasks into multiple queues at once, split by queue name.
	qs := map[string][]*taskqueue.Task{}
	for _, a := range actions {
		switch a := a.(type) {
		case TickLaterAction:
			payload, err := json.Marshal(actionTaskPayload{
				JobID:     jobID,
				Kind:      "TickLaterAction",
				TickNonce: a.TickNonce,
			})
			if err != nil {
				return err
			}
			logging.Infof(c, "Scheduling tick %d after %.1f sec", a.TickNonce, a.When.Sub(time.Now()).Seconds())
			qs[e.cfg.TimersQueueName] = append(qs[e.cfg.TimersQueueName], &taskqueue.Task{
				Path:    e.cfg.TimersQueuePath,
				ETA:     a.When,
				Payload: payload,
			})
		case StartInvocationAction:
			payload, err := json.Marshal(actionTaskPayload{
				JobID:           jobID,
				Kind:            "StartInvocationAction",
				InvocationNonce: a.InvocationNonce,
				TriggeredBy:     string(a.TriggeredBy),
				Triggers:        a.Triggers,
			})
			if err != nil {
				return err
			}
			qs[e.cfg.InvocationsQueueName] = append(qs[e.cfg.InvocationsQueueName], &taskqueue.Task{
				Path:    e.cfg.InvocationsQueuePath,
				Delay:   time.Second, // give the transaction time to land
				Payload: payload,
				RetryOptions: &taskqueue.RetryOptions{
					// Give 5 attempts to mark the job as failed. See 'startInvocation'.
					RetryLimit: invocationRetryLimit + 5,
					MinBackoff: time.Second,
					MaxBackoff: maxInvocationRetryBackoff,
					AgeLimit:   time.Duration(invocationRetryLimit+5) * maxInvocationRetryBackoff,
				},
			})
		case EnqueueBatchOfTriggersAction:
			payload, err := json.Marshal(actionTaskPayload{
				JobID:           jobID,
				Kind:            "EnqueueBatchOfTriggersAction",
				Triggers:        a.Triggers,
				TriggeredJobIDs: a.TriggeredJobIDs,
			})
			if err != nil {
				return err
			}
			qs[e.cfg.InvocationsQueueName] = append(qs[e.cfg.InvocationsQueueName], &taskqueue.Task{
				Path:    e.cfg.InvocationsQueuePath,
				Payload: payload,
			})
		case EnqueueTriggersAction:
			payload, err := json.Marshal(actionTaskPayload{
				JobID:    jobID,
				Kind:     "EnqueueTriggersAction",
				Triggers: a.Triggers,
			})
			if err != nil {
				return err
			}
			qs[e.cfg.InvocationsQueueName] = append(qs[e.cfg.InvocationsQueueName], &taskqueue.Task{
				Path:    e.cfg.InvocationsQueuePath,
				Payload: payload,
			})
		case RecordOverrunAction:
			payload, err := json.Marshal(actionTaskPayload{
				JobID:               jobID,
				Kind:                "RecordOverrunAction",
				Overruns:            a.Overruns,
				RunningInvocationID: a.RunningInvocationID,
			})
			if err != nil {
				return err
			}
			qs[e.cfg.InvocationsQueueName] = append(qs[e.cfg.InvocationsQueueName], &taskqueue.Task{
				Path:    e.cfg.InvocationsQueuePath,
				Delay:   time.Second, // give the transaction time to land
				Payload: payload,
			})
		default:
			logging.Errorf(c, "Unexpected action type %T, skipping", a)
		}
	}
	wg := sync.WaitGroup{}
	errs := errors.NewLazyMultiError(len(qs))
	i := 0
	for queueName, tasks := range qs {
		wg.Add(1)
		go func(i int, queueName string, tasks []*taskqueue.Task) {
			errs.Assign(i, taskqueue.Add(c, queueName, tasks...))
			wg.Done()
		}(i, queueName, tasks)
		i++
	}
	wg.Wait()
	return transient.Tag.Apply(errs.Get())
}

// executeJobAction routes an action that targets a job.
func (e *engineImpl) executeJobAction(c context.Context, payload *actionTaskPayload, retryCount int) error {
	switch payload.Kind {
	case "TickLaterAction":
		return e.jobTimerTick(c, payload.JobID, payload.TickNonce)
	case "StartInvocationAction":
		return e.startInvocation(
			c, payload.JobID, payload.InvocationNonce,
			identity.Identity(payload.TriggeredBy), payload.Triggers, retryCount)
	case "RecordOverrunAction":
		return e.recordOverrun(c, payload.JobID, payload.Overruns, payload.RunningInvocationID)
	case "EnqueueBatchOfTriggersAction":
		return e.newBatchOfTriggers(c, payload.TriggeredJobIDs, payload.Triggers)
	case "EnqueueTriggersAction":
		return e.newTriggers(c, payload.JobID, payload.Triggers)
	default:
		return fmt.Errorf("unexpected job action kind %q", payload.Kind)
	}
}

// jobTimerTick is invoked via task queue in a task with some ETA. It what makes
// cron tick.
//
// Used by v1 engine only!
func (e *engineImpl) jobTimerTick(c context.Context, jobID string, tickNonce int64) error {
	return e.jobTxn(c, jobID, func(c context.Context, job *Job, isNew bool) error {
		if isNew {
			logging.Errorf(c, "Scheduled job is unexpectedly gone")
			return errSkipPut
		}
		logging.Infof(c, "Tick %d has arrived", tickNonce)
		return e.rollSM(c, job, func(sm *StateMachine) error { return sm.OnTimerTick(tickNonce) })
	})
}

// recordOverrun is invoked via task queue when a job should have been started,
// but previous invocation was still running.
//
// It creates new phony Invocation entity (in 'FAILED' state) in the datastore,
// to keep record of all overruns. Doesn't modify Job entity.
//
// Supports both v1 and v2 invocations.
func (e *engineImpl) recordOverrun(c context.Context, jobID string, overruns int, runningInvID int64) error {
	var inv *Invocation
	err := runTxn(c, func(c context.Context) error {
		now := clock.Now(c).UTC()
		var initError error
		inv, initError = e.initInvocation(c, jobID, &Invocation{
			Started:  now,
			Finished: now,
			Status:   task.StatusOverrun,
		}, nil)
		if initError != nil {
			return initError
		}
		if runningInvID == 0 {
			inv.debugLog(c, "New invocation should be starting now, but previous one is still starting")
		} else {
			inv.debugLog(c, "New invocation should be starting now, but previous one is still running: %d", runningInvID)
		}
		inv.debugLog(c, "Total overruns thus far: %d", overruns)
		return transient.Tag.Apply(ds.Put(c, inv))
	})
	if err == nil {
		inv.reportOverrunMetrics(c)
	}
	return err
}

// newBatchOfTriggers splits a batch of triggers into individual per-job async tasks by means
// of EnqueueTriggersAction.
//
// It should run outside of a transaction, since transaction limits
// us to just 5 task queue items hereby limiting len(triggeredJobID) to 5.
//
// Used by v1 engine only!
func (e *engineImpl) newBatchOfTriggers(c context.Context, triggeredJobIDs []string, triggers []*internal.Trigger) error {
	wg := sync.WaitGroup{}
	errs := errors.NewLazyMultiError(len(triggeredJobIDs))
	for i, jobID := range triggeredJobIDs {
		wg.Add(1)
		go func(i int, jobID string) {
			defer wg.Done()
			errs.Assign(i, e.enqueueJobActions(c, jobID, []Action{EnqueueTriggersAction{Triggers: triggers}}))
		}(i, jobID)
	}
	wg.Wait()
	return transient.Tag.Apply(errs.Get())
}

// newTriggers is invoked via task queue when a job receives new Triggers.
//
// It adds these triggers to job's state. If job isn't yet running, this may
// result in StartInvocationAction being emitted.
//
// Used by v1 engine only!
func (e *engineImpl) newTriggers(c context.Context, jobID string, triggers []*internal.Trigger) error {
	return e.jobTxn(c, jobID, func(c context.Context, job *Job, isNew bool) error {
		switch {
		case isNew:
			logging.Errorf(c, "Triggered job is unexpectedly gone")
			return errSkipPut
		case !job.Enabled:
			logging.Warningf(c, "Skipping %d incoming triggers: the job is disabled", len(triggers))
			return errSkipPut
		case job.Paused:
			logging.Warningf(c, "Skipping %d incoming triggers: the job is paused", len(triggers))
			return errSkipPut
		}
		logging.Infof(c, "Triggered %d times", len(triggers))
		return e.rollSM(c, job, func(sm *StateMachine) error {
			sm.OnNewTriggers(triggers)
			return nil
		})
	})
}

////////////////////////////////////////////////////////////////////////////////
// Invocation related task queue routing.

// executeInvAction routes an action that targets an invocation.
func (e *engineImpl) executeInvAction(c context.Context, payload *actionTaskPayload, retryCount int) error {
	switch {
	case payload.InvTimer != nil:
		return e.invocationTimerTick(c, payload.JobID, payload.InvID, payload.InvTimer)
	default:
		return fmt.Errorf("unexpected invocation action kind %q", payload)
	}
}

// invocationTimerTick is called via Task Queue to handle AddTimer callbacks.
//
// See also handlePubSubMessage, it is quite similar.
//
// v1 only! See timerTask for v2, it is slightly different.
func (e *engineImpl) invocationTimerTick(c context.Context, jobID string, invID int64, timer *invocationTimer) error {
	action := fmt.Sprintf("timer %q tick", timer.Name)
	return e.withController(c, jobID, invID, action, func(c context.Context, ctl *taskController) error {
		err := ctl.manager.HandleTimer(c, ctl, timer.Name, timer.Payload)
		switch {
		case err == nil:
			return nil // success! save the invocation
		case transient.Tag.In(err):
			return err // ask for redelivery on transient errors, don't touch the invocation
		}
		// On fatal errors, move the invocation to failed state (if not already).
		if ctl.State().Status != task.StatusFailed {
			ctl.DebugLog("Fatal error when handling timer, aborting invocation - %s", err)
			ctl.State().Status = task.StatusFailed
		}
		return nil // need to save the invocation, even on fatal errors
	})
}

////////////////////////////////////////////////////////////////////////////////
// Task controller and invocation launch.

const (
	// invocationRetryLimit is how many times to retry an invocation before giving
	// up and resuming the job's schedule.
	invocationRetryLimit = 5

	// maxInvocationRetryBackoff is how long to wait before retrying a failed
	// invocation.
	maxInvocationRetryBackoff = 10 * time.Second
)

var (
	// errRetryingLaunch is returned by launchTask if the task failed to start and
	// the launch attempt should be tried again.
	errRetryingLaunch = errors.New("task failed to start, retrying", transient.Tag)
)

// withController fetches the invocation, instantiates the task controller,
// calls the callback, and saves back the modified invocation state, initiating
// all necessary engine transitions along the way.
//
// Does nothing and returns nil if the invocation is already in a final state.
// The callback is not called in this case at all.
//
// Skips saving the invocation if the callback returns non-nil.
//
// 'action' is used exclusively for logging. It's a human readable cause of why
// the controller is instantiated.
func (e *engineImpl) withController(c context.Context, jobID string, invID int64, action string, cb func(context.Context, *taskController) error) error {
	c = logging.SetField(c, "JobID", jobID)
	c = logging.SetField(c, "InvID", invID)

	logging.Infof(c, "Handling %s", action)

	inv, err := e.getInvocation(c, jobID, invID)
	switch {
	case err != nil:
		logging.WithError(err).Errorf(c, "Failed to fetch the invocation")
		return err
	case inv.Status.Final():
		logging.Infof(c, "Skipping %s, the invocation is in final state %q", action, inv.Status)
		return nil
	}

	ctl, err := controllerForInvocation(c, e, inv)
	if err != nil {
		logging.WithError(err).Errorf(c, "Cannot get the controller")
		return err
	}

	if err := cb(c, ctl); err != nil {
		logging.WithError(err).Errorf(c, "Failed to perform %s, skipping saving the invocation", action)
		return err
	}

	if err := ctl.Save(c); err != nil {
		logging.WithError(err).Errorf(c, "Error when saving the invocation")
		return err
	}

	return nil
}

// startInvocation is called via task queue to start running a job.
//
// This call may be retried by task queue service.
//
// Used by v1 engine only!
func (e *engineImpl) startInvocation(c context.Context, jobID string, invocationNonce int64,
	triggeredBy identity.Identity, triggers []*internal.Trigger, retryCount int) error {

	c = logging.SetField(c, "JobID", jobID)
	c = logging.SetField(c, "InvNonce", invocationNonce)
	c = logging.SetField(c, "Attempt", retryCount)

	// TODO(vadimsh): This works only for v1 jobs currently.
	if e.isV2Job(jobID) {
		panic("must be v1 job")
	}

	// Figure out parameters of the invocation based on passed triggers.
	req := requestFromTriggers(c, triggers)
	if triggeredBy != "" {
		req.TriggeredBy = triggeredBy
	}

	// Create new Invocation entity in StatusStarting state and associated it with
	// Job entity.
	//
	// Task queue guarantees not to execute same task concurrently (i.e. retry
	// happens only if previous attempt finished already).
	// There are 4 possibilities here:
	// 1) It is a first attempt. In that case we generate new Invocation in
	//    state STARTING and update Job with a reference to it.
	// 2) It is a retry and previous attempt is still starting (indicated by
	//    IsExpectingInvocation returning true). Assume it failed to start
	//    and launch a new one. Mark old one as obsolete.
	// 3) It is a retry and previous attempt has already started (in this case
	//    the job is in RUNNING state and IsExpectingInvocation returns
	//    false). Assume this retry was unnecessary and skip it.
	// 4) It is a retry and we have retried too many times. Mark the invocation
	//    as failed and resume the job's schedule. This branch executes only if
	//    retryCount check at the end of this function failed to run (e.g. request
	//    handler crashed).
	inv := Invocation{}
	skipRunning := false
	err := e.jobTxn(c, jobID, func(c context.Context, job *Job, isNew bool) error {
		if isNew {
			logging.Errorf(c, "Queued job is unexpectedly gone")
			skipRunning = true
			return errSkipPut
		}
		if !job.State.IsExpectingInvocation(invocationNonce) {
			logging.Errorf(c, "No longer need to start invocation with nonce %d", invocationNonce)
			skipRunning = true
			return errSkipPut
		}
		inv = Invocation{
			Started:         clock.Now(c).UTC(),
			InvocationNonce: invocationNonce,
			Revision:        job.Revision,
			RevisionURL:     job.RevisionURL,
			Task:            job.Task,
			TriggeredJobIDs: job.TriggeredJobIDs,
			RetryCount:      int64(retryCount),
			Status:          task.StatusStarting,
		}
		if _, err := e.initInvocation(c, job.JobID, &inv, &req); err != nil {
			return err
		}
		inv.debugLog(c, "Invocation initiated (attempt %d)", retryCount+1)
		if req.TriggeredBy != "" {
			inv.debugLog(c, "Triggered by %s", req.TriggeredBy)
		}
		if retryCount >= invocationRetryLimit {
			logging.Errorf(c, "Too many attempts, giving up")
			inv.debugLog(c, "Too many attempts, giving up")
			inv.Status = task.StatusFailed
			inv.Finished = clock.Now(c).UTC()
			skipRunning = true
		}
		if err := ds.Put(c, &inv); err != nil {
			return transient.Tag.Apply(err)
		}
		// Move previous invocation (if any) to failed state. It has failed to
		// start.
		if job.State.InvocationID != 0 {
			prev := Invocation{
				ID:     job.State.InvocationID,
				JobKey: inv.JobKey, // works only for v1!
			}
			err := ds.Get(c, &prev)
			if err != nil && err != ds.ErrNoSuchEntity {
				return transient.Tag.Apply(err)
			}
			if err == nil && !prev.Status.Final() {
				prev.debugLog(c, "New invocation is starting (%d), marking this one as failed.", inv.ID)
				// TODO(tandrii): maybe report this special case to monitoring?
				prev.Status = task.StatusFailed
				prev.Finished = clock.Now(c).UTC()
				prev.MutationsCount++
				prev.trimDebugLog()
				if err := ds.Put(c, &prev); err != nil {
					return transient.Tag.Apply(err)
				}
			}
		}
		// Store the reference to the new invocation ID. Unblock the job if we are
		// giving up on retrying.
		return e.rollSM(c, job, func(sm *StateMachine) error {
			sm.OnInvocationStarting(invocationNonce, inv.ID, retryCount)
			if inv.Status == task.StatusFailed {
				sm.OnInvocationStarted(inv.ID)
				sm.OnInvocationDone(inv.ID)
			}
			return nil
		})
	})

	if err != nil {
		logging.WithError(err).Errorf(c, "Failed to update job state")
		return err
	}

	if skipRunning {
		logging.Warningf(c, "No need to start the invocation anymore")
		return nil
	}

	c = logging.SetField(c, "InvID", inv.ID)
	return e.launchTask(c, &inv)
}

// launchTask instantiates an invocation controller and calls its LaunchTask
// method, saving the invocation state when its done.
//
// It returns a transient error if the launch attempt should be retried.
//
// Supports both v1 and v2 invocations.
func (e *engineImpl) launchTask(c context.Context, inv *Invocation) error {
	// Now we have a new Invocation entity in the datastore in StatusStarting
	// state. Grab corresponding TaskManager and launch task through it, keeping
	// track of the progress in created Invocation entity.
	ctl, err := controllerForInvocation(c, e, inv)
	if err != nil {
		// Note: controllerForInvocation returns both ctl and err on errors, with
		// ctl not fully initialized (but good enough for what's done below).
		ctl.DebugLog("Failed to initialize task controller - %s", err)
		ctl.State().Status = task.StatusFailed
		return ctl.Save(c)
	}

	// Ask the manager to start the task. If it returns no errors, it should also
	// move the invocation out of StatusStarting state (a failure to do so is a
	// fatal error). If it returns an error, the invocation is forcefully moved to
	// StatusRetrying or StatusFailed state (depending on whether the error is
	// transient or not and how many retries are left). In either case, invocation
	// never ends up in StatusStarting state.
	err = ctl.manager.LaunchTask(c, ctl)
	if err != nil {
		logging.WithError(err).Errorf(c, "Failed to LaunchTask")
	}
	if ctl.State().Status == task.StatusStarting && err == nil {
		err = fmt.Errorf("LaunchTask didn't move invocation out of StatusStarting")
	}
	if transient.Tag.In(err) && inv.RetryCount+1 >= invocationRetryLimit {
		err = fmt.Errorf("Too many retries, giving up (original error - %s)", err)
	}

	// The task must always end up in a non-starting state. Do it on behalf of the
	// controller if necessary.
	if ctl.State().Status == task.StatusStarting {
		if transient.Tag.In(err) {
			// Note: in v1 version of the engine this status is changed into
			// StatusFailed when new invocation attempt starts (see the transaction
			// above, in particular "New invocation is starting" part). In v2 this
			// will be handled differently (v2 will reuse same Invocation object for
			// retries).
			ctl.State().Status = task.StatusRetrying
		} else {
			ctl.State().Status = task.StatusFailed
		}
	}

	// We MUST commit the state of the invocation. A failure to save the state
	// may cause the job state machine to get stuck. If we can't save it, we need
	// to retry the whole launch attempt from scratch (redoing all the work,
	// a properly implemented LaunchTask should be idempotent).
	if err := ctl.Save(c); err != nil {
		logging.WithError(err).Errorf(c, "Failed to save invocation state")
		return err
	}

	// Task retries happen via the task queue, need to explicitly trigger a retry
	// by returning a transient error.
	if ctl.State().Status == task.StatusRetrying {
		return errRetryingLaunch
	}

	return nil
}

// enqueueInvTimers submits all timers emitted by an invocation manager by
// adding corresponding tasks to the task queue.
//
// Called from a transaction around corresponding Invocation entity.
//
// See executeInvAction for a place where these actions are interpreted.
//
// Used by v1 engine only!
func (e *engineImpl) enqueueInvTimers(c context.Context, jobID string, invID int64, timers []invocationTimer) error {
	assertInTransaction(c)
	tasks := make([]*taskqueue.Task, len(timers))
	for i, timer := range timers {
		payload, err := json.Marshal(actionTaskPayload{
			JobID:    jobID,
			InvID:    invID,
			InvTimer: &timer,
		})
		if err != nil {
			return err
		}
		tasks[i] = &taskqueue.Task{
			Path:    e.cfg.TimersQueuePath,
			ETA:     clock.Now(c).Add(timer.Delay),
			Payload: payload,
		}
	}
	return transient.Tag.Apply(taskqueue.Add(c, e.cfg.TimersQueueName, tasks...))
}

// enqueueTriggers submits all triggers emitted by an invocation manager by
// adding corresponding tasks to the task queue.
//
// Called from a transaction around corresponding Invocation entity.
//
// See executeJobAction for a place where these actions are interpreted.
//
// Used by v1 engine only!
func (e *engineImpl) enqueueTriggers(c context.Context, triggeredJobIDs []string, triggers []*internal.Trigger) error {
	assertInTransaction(c)
	return transient.Tag.Apply(e.enqueueJobActions(c, "", []Action{EnqueueBatchOfTriggersAction{
		Triggers:        triggers,
		TriggeredJobIDs: triggeredJobIDs,
	}}))
}

////////////////////////////////////////////////////////////////////////////////
// Task queue handlers for v2 engine.

// kickLaunchInvocationsBatchTask enqueues LaunchInvocationsBatchTask that
// eventually launches new invocations.
func (e *engineImpl) kickLaunchInvocationsBatchTask(c context.Context, jobID string, invIDs []int64) error {
	payload := &internal.LaunchInvocationsBatchTask{
		Tasks: make([]*internal.LaunchInvocationTask, 0, len(invIDs)),
	}
	for _, invID := range invIDs {
		payload.Tasks = append(payload.Tasks, &internal.LaunchInvocationTask{
			JobId: jobID,
			InvId: invID,
		})
	}
	return e.cfg.Dispatcher.AddTask(c, &tq.Task{
		Payload: payload,
		Delay:   time.Second, // give some time to land Invocation transactions
	})
}

// execLaunchInvocationsBatchTask handles LaunchInvocationsBatchTask by fanning
// out the tasks.
//
// It is the entry point into starting new invocations. Even if the batch
// contains only one task, it still MUST come through LaunchInvocationsBatchTask
// since this is where we "gate" all launches (for example, we can pause the
// corresponding GAE task queue to shutdown new launches during an emergency).
func (e *engineImpl) execLaunchInvocationsBatchTask(c context.Context, tqTask proto.Message) error {
	batch := tqTask.(*internal.LaunchInvocationsBatchTask)

	tasks := []*tq.Task{}
	for _, subtask := range batch.Tasks {
		tasks = append(tasks, &tq.Task{
			DeduplicationKey: fmt.Sprintf("inv:%s:%d", subtask.JobId, subtask.InvId),
			Payload:          subtask,
		})
	}

	return e.cfg.Dispatcher.AddTask(c, tasks...)
}

// execLaunchInvocationTask handles LaunchInvocationTask.
//
// It can be redelivered a bunch of times in case the invocation fails to start.
func (e *engineImpl) execLaunchInvocationTask(c context.Context, tqTask proto.Message) error {
	msg := tqTask.(*internal.LaunchInvocationTask)

	c = logging.SetField(c, "JobID", msg.JobId)
	c = logging.SetField(c, "InvID", msg.InvId)

	hdrs, err := tq.RequestHeaders(c)
	if err != nil {
		return err
	}
	retryCount := hdrs.TaskExecutionCount // 0 for the first attempt
	if retryCount != 0 {
		logging.Warningf(c, "This is a retry (attempt %d)!", retryCount+1)
	}

	// Fetch up-to-date state of the invocation, verify we still need to start it.
	// Log that we are about to do it. We MUST write something to the datastore
	// before attempting the launch to make sure that if the datastore is in read
	// only mode (that happens), we don't spam LaunchTask retries when failing to
	// Save() the state in the end (better to fail now, before LaunchTask call).
	var skipLaunch bool
	var lastInvState Invocation
	logging.Infof(c, "Opening the invocation transaction")
	err = runTxn(c, func(c context.Context) error {
		skipLaunch = false // reset in case the transaction is retried

		// Grab up-to-date invocation state.
		inv := Invocation{ID: msg.InvId}
		switch err := ds.Get(c, &inv); {
		case err == ds.ErrNoSuchEntity:
			// This generally should not happen.
			logging.Warningf(c, "The invocation is unexpectedly gone")
			skipLaunch = true
			return nil
		case err != nil:
			return transient.Tag.Apply(err)
		case !inv.Status.Initial():
			logging.Warningf(c, "The invocation is already running or finished: %s", inv.Status)
			skipLaunch = true
			return nil
		}

		// The invocation is still starting or being retried now. Update its state
		// to indicate we are about to work with it. 'lastInvState' is later passed
		// to the task controller.
		lastInvState = inv
		lastInvState.RetryCount = retryCount
		lastInvState.MutationsCount++
		if retryCount >= invocationRetryLimit {
			logging.Errorf(c, "Too many attempts, giving up")
			lastInvState.debugLog(c, "Too many attempts, giving up")
			lastInvState.Status = task.StatusFailed
			lastInvState.Finished = clock.Now(c).UTC()
			skipLaunch = true
		} else {
			lastInvState.debugLog(c, "Starting the invocation (attempt %d)", retryCount+1)
		}

		// Notify the job controller about the invocation state change. It may
		// decide to update the corresponding job, e.g. if the invocation moves to
		// StatusFailed state.
		if err := e.jobController(msg.JobId).onInvUpdating(c, &inv, &lastInvState, nil, nil); err != nil {
			return err
		}

		// Store the updated invocation.
		lastInvState.trimDebugLog()
		return transient.Tag.Apply(ds.Put(c, &lastInvState))
	})

	switch {
	case err != nil:
		logging.WithError(err).Errorf(c, "Failed to update the invocation")
		return err
	case skipLaunch:
		logging.Warningf(c, "No need to start the invocation anymore")
		return nil
	}

	logging.Infof(c, "Actually launching the task")
	return e.launchTask(c, &lastInvState)
}

// execInvocationFinishedTask handles invocation completion notification.
//
// It is emitted by jobControllerV2.onInvUpdating when invocation switches into
// a final state.
//
// It adds the invocation ID to the set of recently finished invocations and
// kicks off a job triage task that eventually updates Job.ActiveInvocations set
// and moves the cron state machine.
//
// Note that we can't just open a Job transaction right here, since the rate
// of invocation finish events is not controllable and can easily be over 1 QPS
// limit, overwhelming the Job entity group.
//
// If the invocation emitted some triggers when it was finishing, we route them
// here as well.
func (e *engineImpl) execInvocationFinishedTask(c context.Context, tqTask proto.Message) error {
	msg := tqTask.(*internal.InvocationFinishedTask)

	c = logging.SetField(c, "JobID", msg.JobId)
	c = logging.SetField(c, "InvID", msg.InvId)

	if err := recentlyFinishedSet(c, msg.JobId).Add(c, []int64{msg.InvId}); err != nil {
		logging.WithError(err).Errorf(c, "Failed to update recently finished invocations set")
		return err
	}

	// Kick the triage task and fan out the emitted triggers in parallel. Retry
	// the whole thing if any of these operations fail. Everything that happens in
	// this handler is idempotent (including recentlyFinishedSet modification
	// above).

	wg := sync.WaitGroup{}
	errs := errors.MultiError{nil, nil}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if errs[0] = e.kickTriageNow(c, msg.JobId); errs[0] != nil {
			logging.WithError(errs[0]).Errorf(c, "Failed to kick job triage task")
		}
	}()

	if msg.Triggers != nil {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if errs[1] = e.execFanOutTriggersTask(c, msg.Triggers); errs[1] != nil {
				logging.WithError(errs[1]).Errorf(c, "Failed to fan out triggers")
			}
		}()
	}

	wg.Wait()

	if errs.First() != nil {
		return transient.Tag.Apply(errs)
	}

	return nil
}

////////////////////////////////////////////////////////////////////////////////
// Triggers handling (v2 engine).

// execFanOutTriggersTask handles a batch enqueue of triggers.
//
// It is enqueued transactionally by the invocation, and results in a bunch of
// non-transactional EnqueueTriggersTask tasks.
func (e *engineImpl) execFanOutTriggersTask(c context.Context, tqTask proto.Message) error {
	msg := tqTask.(*internal.FanOutTriggersTask)

	tasks := make([]*tq.Task, len(msg.JobIds))
	for i, jobID := range msg.JobIds {
		tasks[i] = &tq.Task{
			Payload: &internal.EnqueueTriggersTask{
				JobId:    jobID,
				Triggers: msg.Triggers,
			},
		}
	}

	return e.cfg.Dispatcher.AddTask(c, tasks...)
}

// execEnqueueTriggersTask adds a bunch of triggers to job's pending triggers
// set and kicks the triage process to process them.
//
// Note: it is invoked through TQ, and also directly from EmitTriggers RPC
// handler.
func (e *engineImpl) execEnqueueTriggersTask(c context.Context, tqTask proto.Message) error {
	msg := tqTask.(*internal.EnqueueTriggersTask)

	c = logging.SetField(c, "JobID", msg.JobId)

	logTriggers := func() {
		for _, t := range msg.Triggers {
			logging.Infof(c, "  %s (emitted by %q, inv %d)", t.Id, t.JobId, t.InvocationId)
		}
	}

	// Don't even bother if the job is paused or disabled. Note that if the job
	// became inactive after this check, the triage will get rid of pending
	// triggers itself. Thus the check here is just an optimization.
	job, err := e.getJob(c, msg.JobId)
	if err != nil {
		logging.WithError(err).Errorf(c, "Failed to grab Job entity")
		return err // transient error getting the job
	}
	if job == nil || !job.Enabled || job.Paused {
		logging.Warningf(c, "Discarding the following triggers since the job is inactive")
		logTriggers()
		return nil
	}

	logging.Infof(c, "Adding the following triggers to the pending triggers set")
	logTriggers()
	if err := pendingTriggersSet(c, msg.JobId).Add(c, msg.Triggers); err != nil {
		logging.WithError(err).Errorf(c, "Failed to update pending triggers set")
		return err
	}

	return e.kickTriageNow(c, msg.JobId)
}

////////////////////////////////////////////////////////////////////////////////
// Timers handling (v2 engine).

// execScheduleTimersTask adds a bunch of TimerTask tasks.
//
// It is emitted by Invocation transaction when it wants to schedule multiple
// timers.
func (e *engineImpl) execScheduleTimersTask(c context.Context, tqTask proto.Message) error {
	msg := tqTask.(*internal.ScheduleTimersTask)

	tasks := make([]*tq.Task, len(msg.Timers))
	for i, timer := range msg.Timers {
		tasks[i] = &tq.Task{
			ETA: google.TimeFromProto(timer.Eta),
			Payload: &internal.TimerTask{
				JobId: msg.JobId,
				InvId: msg.InvId,
				Timer: timer,
			},
		}
	}

	return e.cfg.Dispatcher.AddTask(c, tasks...)
}

// execTimerTask corresponds to a tick of a timer added via AddTimer.
func (e *engineImpl) execTimerTask(c context.Context, tqTask proto.Message) error {
	msg := tqTask.(*internal.TimerTask)
	timer := msg.Timer
	action := fmt.Sprintf("timer %q (%s)", timer.Title, timer.Id)

	return e.withController(c, msg.JobId, msg.InvId, action, func(c context.Context, ctl *taskController) error {
		// Pop the timer from the pending set, if it is still there. Return a fatal
		// error if it isn't to stop this task from being redelivered.
		switch consumed, err := ctl.consumeTimer(timer.Id); {
		case err != nil:
			return err
		case !consumed:
			return fmt.Errorf("no such timer: %s", timer.Id)
		}

		// Let the task manager handle the timer. It may add new timers.
		ctl.DebugLog("Handling timer %q (%s)", timer.Title, timer.Id)
		err := ctl.manager.HandleTimer(c, ctl, timer.Title, timer.Payload)
		switch {
		case err == nil:
			return nil // success! save the invocation
		case transient.Tag.In(err):
			return err // ask for redelivery on transient errors, don't touch the invocation
		}

		// On fatal errors, move the invocation to failed state (if not already).
		if ctl.State().Status != task.StatusFailed {
			ctl.DebugLog("Fatal error when handling timer, aborting invocation - %s", err)
			ctl.State().Status = task.StatusFailed
		}

		// Need to save the invocation, even on fatal errors (to indicate that the
		// timer has been consumed). So return nil.
		return nil
	})
}

////////////////////////////////////////////////////////////////////////////////
// Cron handling (v2 engine).

// execCronTickTask corresponds to a delayed tick emitted by a cron state
// machine.
//
// See jobTimerTick for (deprecated) v1 equivalent.
func (e *engineImpl) execCronTickTask(c context.Context, tqTask proto.Message) error {
	msg := tqTask.(*internal.CronTickTask)
	return e.jobTxn(c, msg.JobId, func(c context.Context, job *Job, isNew bool) error {
		if isNew {
			logging.Errorf(c, "Scheduled job is unexpectedly gone")
			return errSkipPut
		}
		logging.Infof(c, "Tick %d has arrived", msg.TickNonce)
		return e.jobController(msg.JobId).onJobCronTick(c, job, msg)
	})
}

////////////////////////////////////////////////////////////////////////////////
// Triage procedure (v2 engine).

// kickTriageNow enqueues a task to perform a triage for some job, if no such
// task was enqueued recently.
//
// Does it even if the job no longer exists or has been disabled. Such triage
// will just be skipped later.
//
// Uses named tasks and memcache internally, thus can't be part of a
// transaction. If you want to kick the triage transactionally, use
// kickTriageLater().
func (e *engineImpl) kickTriageNow(c context.Context, jobID string) error {
	assertNotInTransaction(c)

	c = logging.SetField(c, "JobID", jobID)

	// Throttle to once per 2 sec (and make sure it is always in the future).
	eta := clock.Now(c).Unix()
	eta = (eta/2 + 1) * 2
	dedupKey := fmt.Sprintf("triage:%s:%d", jobID, eta)

	// Use cheaper but crappier memcache as a first dedup check.
	itm := memcache.NewItem(c, dedupKey).SetExpiration(time.Minute)
	if memcache.Get(c, itm) == nil {
		logging.Infof(c, "The triage task has already been scheduled")
		return nil
	}

	// Enqueue the triage task, if not already there. This is rock solid, but slow
	// second dedup check.
	err := e.cfg.Dispatcher.AddTask(c, &tq.Task{
		DeduplicationKey: dedupKey,
		ETA:              time.Unix(eta, 0),
		Payload:          &internal.TriageJobStateTask{JobId: jobID},
	})
	if err != nil {
		return err
	}
	logging.Infof(c, "Scheduled the triage task")

	// Best effort in setting dedup memcache flag. No big deal if it fails.
	if err := memcache.Set(c, itm); err != nil {
		logging.WithError(err).Warningf(c, "Failed to set memcache triage flag")
	}

	return nil
}

// kickTriageLater schedules a triage to be kicked later.
//
// Unlike kickTriageNow, this just posts a single TQ task, and thus can be
// used inside transactions.
func (e *engineImpl) kickTriageLater(c context.Context, jobID string, delay time.Duration) error {
	c = logging.SetField(c, "JobID", jobID)
	return e.cfg.Dispatcher.AddTask(c, &tq.Task{
		Payload: &internal.KickTriageTask{JobId: jobID},
		Delay:   delay,
	})
}

// execKickTriageTask handles delayed KickTriageTask by scheduling a triage.
func (e *engineImpl) execKickTriageTask(c context.Context, tqTask proto.Message) error {
	return e.kickTriageNow(c, tqTask.(*internal.KickTriageTask).JobId)
}

// execTriageJobStateTask performs the triage of a job.
//
// It is throttled to run at most once per 2 seconds.
//
// It looks at pending triggers and recently finished invocations and launches
// new invocations (or schedules timers to do it later).
func (e *engineImpl) execTriageJobStateTask(c context.Context, tqTask proto.Message) error {
	jobID := tqTask.(*internal.TriageJobStateTask).JobId

	c = logging.SetField(c, "JobID", jobID)

	startedTs := clock.Now(c)
	defer func() {
		logging.Infof(c, "Triage took %s", clock.Now(c).Sub(startedTs))
	}()

	// There's error logging inside of triageOp already.
	op := triageOp{
		jobID:            jobID,
		dispatcher:       e.cfg.Dispatcher,
		triggeringPolicy: e.triggeringPolicy,
		enqueueInvocations: func(c context.Context, job *Job, req []task.Request) error {
			_, err := e.enqueueInvocations(c, job, req)
			return err
		},
	}
	if err := op.prepare(c); err != nil {
		return err
	}

	err := e.jobTxn(c, jobID, func(c context.Context, job *Job, isNew bool) error {
		if isNew {
			logging.Warningf(c, "The job is unexpectedly gone")
			return errSkipPut
		}
		return op.transaction(c, job)
	})
	if err != nil {
		logging.WithError(err).Errorf(c, "Failed to perform triage transaction")
		return err
	}

	// Best effort cleanup.
	op.finalize(c)
	return nil
}

// triggeringPolicy decides how to convert a set of pending triggers into
// a bunch of new invocations.
//
// Called within a job transaction. Must not do any expensive calls.
func (e *engineImpl) triggeringPolicy(c context.Context, job *Job, triggers []*internal.Trigger) ([]task.Request, error) {
	// TODO(vadimsh): This policy matches v1 behavior:
	//  * Don't start anything new if some invocation is already running.
	//  * Otherwise consume all pending triggers at once.
	if len(job.ActiveInvocations) != 0 {
		return nil, nil
	}
	return []task.Request{requestFromTriggers(c, triggers)}, nil
}

////////////////////////////////////////////////////////////////////////////////
// PubSub related methods.

// topicParams is passed to prepareTopic by task.Controller.
type topicParams struct {
	jobID     string       // the job invocation belongs to
	invID     int64        // ID of the invocation itself
	manager   task.Manager // task manager for the invocation
	publisher string       // name of publisher to add to PubSub topic.
}

// pubsubAuthToken describes how to generate HMAC protected tokens used to
// authenticate PubSub messages.
var pubsubAuthToken = tokens.TokenKind{
	Algo:       tokens.TokenAlgoHmacSHA256,
	Expiration: 48 * time.Hour,
	SecretKey:  "pubsub_auth_token",
	Version:    1,
}

// handlePubSubMessage routes the pubsub message to the invocation.
//
// See also invocationTimerTick, it is quite similar.
func (e *engineImpl) handlePubSubMessage(c context.Context, msg *pubsub.PubsubMessage) error {
	logging.Infof(c, "Received PubSub message %q", msg.MessageId)

	// Extract Job and Invocation ID from validated auth_token.
	var jobID string
	var invID int64
	data, err := pubsubAuthToken.Validate(c, msg.Attributes["auth_token"], nil)
	if err != nil {
		logging.WithError(err).Errorf(c, "Bad auth_token attribute")
		return err
	}
	jobID = data["job"]
	if invID, err = strconv.ParseInt(data["inv"], 10, 64); err != nil {
		logging.WithError(err).Errorf(c, "Could not parse 'inv' %q", data["inv"])
		return err
	}

	// Hand the message to the controller.
	action := fmt.Sprintf("pubsub message %q", msg.MessageId)
	return e.withController(c, jobID, invID, action, func(c context.Context, ctl *taskController) error {
		err := ctl.manager.HandleNotification(c, ctl, msg)
		switch {
		case err == nil:
			return nil // success! save the invocation
		case transient.Tag.In(err):
			return err // ask for redelivery on transient errors, don't touch the invocation
		}
		// On fatal errors, move the invocation to failed state (if not already).
		if ctl.State().Status != task.StatusFailed {
			ctl.DebugLog("Fatal error when handling PubSub notification, aborting invocation - %s", err)
			ctl.State().Status = task.StatusFailed
		}
		return nil // need to save the invocation, even on fatal errors
	})
}

// genTopicAndSubNames derives PubSub topic and subscription names to use for
// notifications from given publisher.
func (e *engineImpl) genTopicAndSubNames(c context.Context, manager, publisher string) (topic string, sub string) {
	// Avoid accidental override of the topic when running on dev server.
	prefix := "scheduler"
	if info.IsDevAppServer(c) {
		prefix = "dev-scheduler"
	}

	// Each publisher gets its own topic (and subscription), so it's clearer from
	// logs and PubSub console who's calling what. PubSub topics can't have "@" in
	// them, so replace "@" with "~". URL encoding could have been used too, but
	// Cloud Console confuses %40 with its own URL encoding and doesn't display
	// all pages correctly.
	id := fmt.Sprintf("%s+%s+%s",
		prefix,
		manager,
		strings.Replace(publisher, "@", "~", -1))

	appID := info.AppID(c)
	topic = fmt.Sprintf("projects/%s/topics/%s", appID, id)
	sub = fmt.Sprintf("projects/%s/subscriptions/%s", appID, id)
	return
}

// prepareTopic creates a pubsub topic that can be used to pass task related
// messages back to the task.Manager that handles the task.
//
// It returns full topic name, as well as a token that securely identifies the
// task. It should be put into 'auth_token' attribute of PubSub messages by
// whoever publishes them.
func (e *engineImpl) prepareTopic(c context.Context, params *topicParams) (topic string, tok string, err error) {
	// If given URL, ask the service for name of its default service account.
	// FetchServiceInfo implements efficient cache internally, so it's fine to
	// call it often.
	if strings.HasPrefix(params.publisher, "https://") {
		logging.Infof(c, "Fetching info about %q", params.publisher)
		serviceInfo, err := signing.FetchServiceInfoFromLUCIService(c, params.publisher)
		if err != nil {
			logging.Errorf(c, "Failed to fetch info about %q - %s", params.publisher, err)
			return "", "", err
		}
		logging.Infof(c, "%q is using %q", params.publisher, serviceInfo.ServiceAccountName)
		params.publisher = serviceInfo.ServiceAccountName
	}

	topic, sub := e.genTopicAndSubNames(c, params.manager.Name(), params.publisher)

	// Put same parameters in push URL to make them visible in logs. On dev server
	// use pull based subscription, since localhost push URL is not valid.
	pushURL := ""
	if !info.IsDevAppServer(c) {
		urlParams := url.Values{}
		urlParams.Add("kind", params.manager.Name())
		urlParams.Add("publisher", params.publisher)
		pushURL = fmt.Sprintf(
			"https://%s%s?%s", info.DefaultVersionHostname(c), e.cfg.PubSubPushPath, urlParams.Encode())
	}

	// Create and configure the topic. Do it only once.
	err = e.opsCache.Do(c, fmt.Sprintf("prepareTopic:v1:%s", topic), func() error {
		if e.configureTopic != nil {
			return e.configureTopic(c, topic, sub, pushURL, params.publisher)
		}
		return configureTopic(c, topic, sub, pushURL, params.publisher, "")
	})
	if err != nil {
		return "", "", err
	}

	// Encode full invocation identifier (job key + invocation ID) into HMAC
	// protected token.
	tok, err = pubsubAuthToken.Generate(c, nil, map[string]string{
		"job": params.jobID,
		"inv": fmt.Sprintf("%d", params.invID),
	}, 0)
	if err != nil {
		return "", "", err
	}

	return topic, tok, nil
}
