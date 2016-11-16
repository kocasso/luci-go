// Code generated by protoc-gen-go.
// source: github.com/luci/luci-go/scheduler/appengine/messages/cron.proto
// DO NOT EDIT!

/*
Package messages is a generated protocol buffer package.

It is generated from these files:
	github.com/luci/luci-go/scheduler/appengine/messages/cron.proto

It has these top-level messages:
	Job
	Trigger
	NoopTask
	UrlFetchTask
	SwarmingTask
	BuildbucketTask
	ProjectConfig
	TaskDefWrapper
*/
package messages

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Job specifies a single regular job belonging to a project.
//
// Such jobs runs on a schedule or can be triggered by some trigger.
type Job struct {
	// Id is a name of the job (unique for the project).
	//
	// Must match '^[0-9A-Za-z_\-\.]{1,100}$'.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Schedule describes when to run the job.
	//
	// Supported kinds of schedules (illustrated by examples):
	//   - "* 0 * * * *": standard cron-like expression. Cron engine will attempt
	//     to start a job at specified moments in time (based on UTC clock). If
	//     when triggering a job, previous invocation is still running, an overrun
	//     will be recorded (and next attempt to start a job happens based on the
	//     schedule, not when the previous invocation finishes). This is absolute
	//     schedule (i.e. doesn't depend on job state).
	//   - "with 10s interval": runs invocations in a loop, waiting 10s after
	//     finishing invocation before starting a new one. This is relative
	//     schedule. Overruns are not possible.
	//   - "continuously" is alias for "with 0s interval", meaning the job will
	//     run in a loop without any pauses.
	//   - "triggered" schedule indicates that job is always started via "Run now"
	//     button or via a trigger.
	//
	// Default is "triggered".
	Schedule string `protobuf:"bytes,2,opt,name=schedule" json:"schedule,omitempty"`
	// Disabled is true to disable this job.
	Disabled bool `protobuf:"varint,3,opt,name=disabled" json:"disabled,omitempty"`
	// Task defines what exactly to execute.
	//
	// TODO(vadimsh): Remove this field once all configs are updated not to
	// use it.
	Task *TaskDefWrapper `protobuf:"bytes,4,opt,name=task" json:"task,omitempty"`
	// Noop is used for testing. It is "do nothing" task.
	Noop *NoopTask `protobuf:"bytes,100,opt,name=noop" json:"noop,omitempty"`
	// UrlFetch can be used to make a simple HTTP call.
	UrlFetch *UrlFetchTask `protobuf:"bytes,101,opt,name=url_fetch,json=urlFetch" json:"url_fetch,omitempty"`
	// SwarmingTask can be used to schedule swarming job.
	Swarming *SwarmingTask `protobuf:"bytes,102,opt,name=swarming" json:"swarming,omitempty"`
	// BuildbucketTask can be used to schedule buildbucket job.
	Buildbucket *BuildbucketTask `protobuf:"bytes,103,opt,name=buildbucket" json:"buildbucket,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Job) GetTask() *TaskDefWrapper {
	if m != nil {
		return m.Task
	}
	return nil
}

func (m *Job) GetNoop() *NoopTask {
	if m != nil {
		return m.Noop
	}
	return nil
}

func (m *Job) GetUrlFetch() *UrlFetchTask {
	if m != nil {
		return m.UrlFetch
	}
	return nil
}

func (m *Job) GetSwarming() *SwarmingTask {
	if m != nil {
		return m.Swarming
	}
	return nil
}

func (m *Job) GetBuildbucket() *BuildbucketTask {
	if m != nil {
		return m.Buildbucket
	}
	return nil
}

// Trigger specifies a job that triggers other jobs.
//
// It is a special kind of job that periodically checks the state of the world
// and triggers other jobs.
type Trigger struct {
	// Id is a name of the job (unique for the project).
	//
	// Must match '^[0-9A-Za-z_\-\.]{1,100}$'. It's in the same namespace as
	// regular jobs.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Schedule describes when to run this triggering job.
	//
	// See Job.schedule fro more info. Default is "with 30s interval".
	Schedule string `protobuf:"bytes,2,opt,name=schedule" json:"schedule,omitempty"`
	// Disabled is true to disable this job.
	Disabled bool `protobuf:"varint,3,opt,name=disabled" json:"disabled,omitempty"`
	// Noop is used for testing. It is "do nothing" trigger.
	Noop *NoopTask `protobuf:"bytes,100,opt,name=noop" json:"noop,omitempty"`
}

func (m *Trigger) Reset()                    { *m = Trigger{} }
func (m *Trigger) String() string            { return proto.CompactTextString(m) }
func (*Trigger) ProtoMessage()               {}
func (*Trigger) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Trigger) GetNoop() *NoopTask {
	if m != nil {
		return m.Noop
	}
	return nil
}

// NoopTask is used for testing. It is "do nothing" task.
type NoopTask struct {
}

func (m *NoopTask) Reset()                    { *m = NoopTask{} }
func (m *NoopTask) String() string            { return proto.CompactTextString(m) }
func (*NoopTask) ProtoMessage()               {}
func (*NoopTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// UrlFetchTask specifies parameters for simple HTTP call.
type UrlFetchTask struct {
	// Method is HTTP method to use, such as "GET" or "POST". Default is "GET".
	Method string `protobuf:"bytes,1,opt,name=method" json:"method,omitempty"`
	// Url to send the request to.
	Url string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	// Timeout is how long to wait for request to complete. Default is 60 sec.
	TimeoutSec int32 `protobuf:"varint,3,opt,name=timeout_sec,json=timeoutSec" json:"timeout_sec,omitempty"`
}

func (m *UrlFetchTask) Reset()                    { *m = UrlFetchTask{} }
func (m *UrlFetchTask) String() string            { return proto.CompactTextString(m) }
func (*UrlFetchTask) ProtoMessage()               {}
func (*UrlFetchTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// SwarmingTask specifies parameters of Swarming-based jobs.
type SwarmingTask struct {
	// Server is URL of the swarming service to use.
	Server string `protobuf:"bytes,1,opt,name=server" json:"server,omitempty"`
	// What to run. Only one of 'command' or 'isolated_ref' must be given.
	Command     []string                  `protobuf:"bytes,2,rep,name=command" json:"command,omitempty"`
	IsolatedRef *SwarmingTask_IsolatedRef `protobuf:"bytes,3,opt,name=isolated_ref,json=isolatedRef" json:"isolated_ref,omitempty"`
	// Additional arguments to pass to isolated command.
	ExtraArgs []string `protobuf:"bytes,4,rep,name=extra_args,json=extraArgs" json:"extra_args,omitempty"`
	// List of "key=value" pairs with additional OS environment variables.
	Env []string `protobuf:"bytes,5,rep,name=env" json:"env,omitempty"`
	// Where to run it. List of "key:value" pairs.
	Dimensions []string `protobuf:"bytes,6,rep,name=dimensions" json:"dimensions,omitempty"`
	// Tags is a list of tags (as "key:value" pairs) to assign to the task.
	Tags []string `protobuf:"bytes,7,rep,name=tags" json:"tags,omitempty"`
	// Priority is task priority (or niceness, lower value - higher priority).
	Priority int32 `protobuf:"varint,8,opt,name=priority" json:"priority,omitempty"`
	// Timeouts. All optional. The scheduler will set reasonable default values.
	ExecutionTimeoutSecs int32 `protobuf:"varint,9,opt,name=execution_timeout_secs,json=executionTimeoutSecs" json:"execution_timeout_secs,omitempty"`
	GracePeriodSecs      int32 `protobuf:"varint,10,opt,name=grace_period_secs,json=gracePeriodSecs" json:"grace_period_secs,omitempty"`
	IoTimeoutSecs        int32 `protobuf:"varint,11,opt,name=io_timeout_secs,json=ioTimeoutSecs" json:"io_timeout_secs,omitempty"`
}

func (m *SwarmingTask) Reset()                    { *m = SwarmingTask{} }
func (m *SwarmingTask) String() string            { return proto.CompactTextString(m) }
func (*SwarmingTask) ProtoMessage()               {}
func (*SwarmingTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SwarmingTask) GetIsolatedRef() *SwarmingTask_IsolatedRef {
	if m != nil {
		return m.IsolatedRef
	}
	return nil
}

// IsolatedRef defines a data tree reference, normally a reference to
// an .isolated file
type SwarmingTask_IsolatedRef struct {
	Isolated       string `protobuf:"bytes,1,opt,name=isolated" json:"isolated,omitempty"`
	IsolatedServer string `protobuf:"bytes,2,opt,name=isolated_server,json=isolatedServer" json:"isolated_server,omitempty"`
	Namespace      string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *SwarmingTask_IsolatedRef) Reset()                    { *m = SwarmingTask_IsolatedRef{} }
func (m *SwarmingTask_IsolatedRef) String() string            { return proto.CompactTextString(m) }
func (*SwarmingTask_IsolatedRef) ProtoMessage()               {}
func (*SwarmingTask_IsolatedRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

// BuildbucketTask specifies parameters of Buildbucket-based jobs.
type BuildbucketTask struct {
	// Server is URL of the bulildbucket service to use.
	Server string `protobuf:"bytes,1,opt,name=server" json:"server,omitempty"`
	// Bucket defines what bucket to add the task to.
	Bucket string `protobuf:"bytes,2,opt,name=bucket" json:"bucket,omitempty"`
	// Builder defines what to run.
	Builder string `protobuf:"bytes,3,opt,name=builder" json:"builder,omitempty"`
	// Properties is arbitrary "key:value" pairs describing the task.
	Properties []string `protobuf:"bytes,4,rep,name=properties" json:"properties,omitempty"`
	// Tags is a list of tags (as "key:value" pairs) to assign to the task.
	Tags []string `protobuf:"bytes,5,rep,name=tags" json:"tags,omitempty"`
}

func (m *BuildbucketTask) Reset()                    { *m = BuildbucketTask{} }
func (m *BuildbucketTask) String() string            { return proto.CompactTextString(m) }
func (*BuildbucketTask) ProtoMessage()               {}
func (*BuildbucketTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// ProjectConfig defines a schema for config file that describe jobs belonging
// to some project.
type ProjectConfig struct {
	// Job is a set of jobs defined in the project.
	Job []*Job `protobuf:"bytes,1,rep,name=job" json:"job,omitempty"`
	// Trigger is a set of triggering jobs defined in the project.
	Trigger []*Trigger `protobuf:"bytes,2,rep,name=trigger" json:"trigger,omitempty"`
}

func (m *ProjectConfig) Reset()                    { *m = ProjectConfig{} }
func (m *ProjectConfig) String() string            { return proto.CompactTextString(m) }
func (*ProjectConfig) ProtoMessage()               {}
func (*ProjectConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ProjectConfig) GetJob() []*Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *ProjectConfig) GetTrigger() []*Trigger {
	if m != nil {
		return m.Trigger
	}
	return nil
}

// TaskDefWrapper is a union type of all possible tasks known to the scheduler.
//
// It is used internally when storing jobs in the datastore.
//
// TODO(vadimsh): Remove '_task' suffixes once TaskDefWrapper is no longer
// a part of 'Job' proto.
type TaskDefWrapper struct {
	Noop            *NoopTask        `protobuf:"bytes,1,opt,name=noop" json:"noop,omitempty"`
	UrlFetch        *UrlFetchTask    `protobuf:"bytes,2,opt,name=url_fetch,json=urlFetch" json:"url_fetch,omitempty"`
	SwarmingTask    *SwarmingTask    `protobuf:"bytes,3,opt,name=swarming_task,json=swarmingTask" json:"swarming_task,omitempty"`
	BuildbucketTask *BuildbucketTask `protobuf:"bytes,4,opt,name=buildbucket_task,json=buildbucketTask" json:"buildbucket_task,omitempty"`
}

func (m *TaskDefWrapper) Reset()                    { *m = TaskDefWrapper{} }
func (m *TaskDefWrapper) String() string            { return proto.CompactTextString(m) }
func (*TaskDefWrapper) ProtoMessage()               {}
func (*TaskDefWrapper) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TaskDefWrapper) GetNoop() *NoopTask {
	if m != nil {
		return m.Noop
	}
	return nil
}

func (m *TaskDefWrapper) GetUrlFetch() *UrlFetchTask {
	if m != nil {
		return m.UrlFetch
	}
	return nil
}

func (m *TaskDefWrapper) GetSwarmingTask() *SwarmingTask {
	if m != nil {
		return m.SwarmingTask
	}
	return nil
}

func (m *TaskDefWrapper) GetBuildbucketTask() *BuildbucketTask {
	if m != nil {
		return m.BuildbucketTask
	}
	return nil
}

func init() {
	proto.RegisterType((*Job)(nil), "messages.Job")
	proto.RegisterType((*Trigger)(nil), "messages.Trigger")
	proto.RegisterType((*NoopTask)(nil), "messages.NoopTask")
	proto.RegisterType((*UrlFetchTask)(nil), "messages.UrlFetchTask")
	proto.RegisterType((*SwarmingTask)(nil), "messages.SwarmingTask")
	proto.RegisterType((*SwarmingTask_IsolatedRef)(nil), "messages.SwarmingTask.IsolatedRef")
	proto.RegisterType((*BuildbucketTask)(nil), "messages.BuildbucketTask")
	proto.RegisterType((*ProjectConfig)(nil), "messages.ProjectConfig")
	proto.RegisterType((*TaskDefWrapper)(nil), "messages.TaskDefWrapper")
}

func init() {
	proto.RegisterFile("github.com/luci/luci-go/scheduler/appengine/messages/cron.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 728 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0x4b, 0x6f, 0xdb, 0x46,
	0x10, 0x06, 0x25, 0x59, 0x8f, 0xa1, 0x64, 0xd9, 0x8b, 0x42, 0xd8, 0x1a, 0x6d, 0x2d, 0xf0, 0xe0,
	0x0a, 0x7d, 0x48, 0x80, 0xdc, 0x9b, 0x0f, 0x45, 0x1b, 0x27, 0x40, 0x7c, 0x08, 0x0c, 0xda, 0x41,
	0x90, 0x43, 0x40, 0xf0, 0x31, 0xa2, 0xd6, 0x26, 0xb9, 0xc4, 0xee, 0xd2, 0xb1, 0xef, 0xb9, 0xe7,
	0x77, 0xe5, 0xdf, 0xe4, 0x27, 0x04, 0x5c, 0x3e, 0x6d, 0x24, 0x86, 0x03, 0xe4, 0x62, 0x70, 0xbe,
	0xef, 0x9b, 0xd9, 0xf5, 0x7c, 0x3b, 0x23, 0xf8, 0x37, 0x64, 0x6a, 0x9b, 0x79, 0x4b, 0x9f, 0xc7,
	0xab, 0x28, 0xf3, 0x99, 0xfe, 0xf3, 0x77, 0xc8, 0x57, 0xd2, 0xdf, 0x62, 0x90, 0x45, 0x28, 0x56,
	0x6e, 0x9a, 0x62, 0x12, 0xb2, 0x04, 0x57, 0x31, 0x4a, 0xe9, 0x86, 0x28, 0x57, 0xbe, 0xe0, 0xc9,
	0x32, 0x15, 0x5c, 0x71, 0x32, 0xac, 0x40, 0xeb, 0x53, 0x07, 0xba, 0x67, 0xdc, 0x23, 0xbb, 0xd0,
	0x61, 0x01, 0x35, 0xe6, 0xc6, 0x62, 0x64, 0x77, 0x58, 0x40, 0x0e, 0x60, 0x58, 0x15, 0xa3, 0x1d,
	0x8d, 0xd6, 0x71, 0xce, 0x05, 0x4c, 0xba, 0x5e, 0x84, 0x01, 0xed, 0xce, 0x8d, 0xc5, 0xd0, 0xae,
	0x63, 0xf2, 0x17, 0xf4, 0x94, 0x2b, 0xaf, 0x69, 0x6f, 0x6e, 0x2c, 0xcc, 0x35, 0x5d, 0x56, 0x07,
	0x2d, 0x2f, 0x5d, 0x79, 0x7d, 0x8a, 0x9b, 0x37, 0x22, 0xbf, 0x99, 0xb0, 0xb5, 0x8a, 0x1c, 0x41,
	0x2f, 0xe1, 0x3c, 0xa5, 0x81, 0x56, 0x93, 0x46, 0xfd, 0x8a, 0xf3, 0x34, 0xcf, 0xb0, 0x35, 0x4f,
	0x8e, 0x61, 0x94, 0x89, 0xc8, 0xd9, 0xa0, 0xf2, 0xb7, 0x14, 0xb5, 0x78, 0xd6, 0x88, 0x5f, 0x8b,
	0xe8, 0x45, 0xce, 0xe8, 0x84, 0x61, 0x56, 0x46, 0x64, 0x0d, 0x43, 0xf9, 0xde, 0x15, 0x31, 0x4b,
	0x42, 0xba, 0x79, 0x98, 0x73, 0x51, 0x32, 0x45, 0x4e, 0xa5, 0x23, 0x27, 0x60, 0x7a, 0x19, 0x8b,
	0x02, 0x2f, 0xf3, 0xaf, 0x51, 0xd1, 0x50, 0xa7, 0xfd, 0xdc, 0xa4, 0xfd, 0xdf, 0x90, 0x3a, 0xb3,
	0xad, 0xb6, 0xee, 0x60, 0x70, 0x29, 0x58, 0x18, 0xa2, 0xf8, 0x61, 0xed, 0x7c, 0x62, 0x83, 0x2c,
	0x80, 0x61, 0x85, 0x58, 0x6f, 0x61, 0xdc, 0xee, 0x08, 0x99, 0x41, 0x3f, 0x46, 0xb5, 0xe5, 0xd5,
	0x7d, 0xca, 0x88, 0xec, 0x41, 0x37, 0x13, 0x51, 0x79, 0x9d, 0xfc, 0x93, 0x1c, 0x82, 0xa9, 0x58,
	0x8c, 0x3c, 0x53, 0x8e, 0x44, 0x5f, 0x5f, 0x66, 0xc7, 0x86, 0x12, 0xba, 0x40, 0xdf, 0xfa, 0xd0,
	0x83, 0x71, 0xbb, 0x73, 0x79, 0x6d, 0x89, 0xe2, 0x06, 0x45, 0x55, 0xbb, 0x88, 0x08, 0x85, 0x81,
	0xcf, 0xe3, 0xd8, 0x4d, 0x02, 0xda, 0x99, 0x77, 0x17, 0x23, 0xbb, 0x0a, 0xc9, 0x73, 0x18, 0x33,
	0xc9, 0x23, 0x57, 0x61, 0xe0, 0x08, 0xdc, 0xe8, 0x43, 0xcc, 0xb5, 0xf5, 0x75, 0x67, 0x96, 0x2f,
	0x4b, 0xa9, 0x8d, 0x1b, 0xdb, 0x64, 0x4d, 0x40, 0x7e, 0x05, 0xc0, 0x5b, 0x25, 0x5c, 0xc7, 0x15,
	0xa1, 0xa4, 0x3d, 0x7d, 0xc6, 0x48, 0x23, 0xff, 0x89, 0x50, 0xe6, 0xff, 0x1b, 0x26, 0x37, 0x74,
	0x47, 0xe3, 0xf9, 0x27, 0xf9, 0x0d, 0x20, 0x60, 0x31, 0x26, 0x92, 0xf1, 0x44, 0xd2, 0xbe, 0x26,
	0x5a, 0x08, 0x21, 0xf9, 0xc3, 0x0d, 0x25, 0x1d, 0x68, 0x46, 0x7f, 0xe7, 0xce, 0xa4, 0x82, 0x71,
	0xc1, 0xd4, 0x1d, 0x1d, 0xea, 0x66, 0xd4, 0x31, 0xf9, 0x07, 0x66, 0x78, 0x8b, 0x7e, 0xa6, 0x18,
	0x4f, 0x9c, 0x56, 0xd7, 0x24, 0x1d, 0x69, 0xe5, 0x4f, 0x35, 0x7b, 0x59, 0xf7, 0x4f, 0x92, 0x3f,
	0x60, 0x3f, 0x14, 0xae, 0x8f, 0x4e, 0x8a, 0x82, 0xf1, 0xa0, 0x48, 0x00, 0x9d, 0x30, 0xd5, 0xc4,
	0xb9, 0xc6, 0xb5, 0xf6, 0x08, 0xa6, 0x8c, 0xdf, 0x2f, 0x6d, 0x6a, 0xe5, 0x84, 0xf1, 0x56, 0xcd,
	0x83, 0x14, 0xcc, 0x56, 0x9b, 0xf2, 0x4b, 0x57, 0x8d, 0x2a, 0x4d, 0xa9, 0x63, 0xf2, 0x3b, 0x4c,
	0xeb, 0xe6, 0x97, 0xbe, 0x15, 0xf6, 0xef, 0x56, 0xf0, 0x45, 0xe1, 0xdf, 0x2f, 0x30, 0x4a, 0xdc,
	0x18, 0x65, 0xea, 0xfa, 0xa8, 0x2d, 0x1a, 0xd9, 0x0d, 0x60, 0x7d, 0x34, 0x60, 0xfa, 0x60, 0x12,
	0xbe, 0xf9, 0x12, 0x66, 0xd0, 0x2f, 0x87, 0xa9, 0x38, 0xa9, 0x8c, 0xf2, 0x17, 0xa2, 0x67, 0x07,
	0x45, 0x59, 0xbf, 0x0a, 0x73, 0xa7, 0x52, 0xc1, 0x53, 0x14, 0x8a, 0x61, 0x65, 0x6d, 0x0b, 0xa9,
	0x9d, 0xda, 0x69, 0x9c, 0xb2, 0xde, 0xc1, 0xe4, 0x5c, 0xf0, 0x2b, 0xf4, 0xd5, 0x33, 0x9e, 0x6c,
	0x58, 0x48, 0x0e, 0xa1, 0x7b, 0xc5, 0x3d, 0x6a, 0xcc, 0xbb, 0x0b, 0x73, 0x3d, 0x69, 0x5e, 0xd7,
	0x19, 0xf7, 0xec, 0x9c, 0x21, 0x7f, 0xc2, 0x40, 0x15, 0xc3, 0xaa, 0x5f, 0xa8, 0xb9, 0xde, 0x6f,
	0xed, 0xaa, 0x82, 0xb0, 0x2b, 0x85, 0xf5, 0xd9, 0x80, 0xdd, 0xfb, 0x0b, 0xac, 0x9e, 0x4c, 0xe3,
	0x7b, 0x56, 0x57, 0xe7, 0x89, 0xab, 0xeb, 0x04, 0x26, 0xd5, 0x4a, 0x72, 0xf4, 0x3a, 0xed, 0x3e,
	0xba, 0xbf, 0xc6, 0xb2, 0x3d, 0x93, 0xa7, 0xb0, 0xd7, 0xda, 0x4a, 0x4e, 0x6b, 0x1d, 0x3f, 0xb2,
	0xc8, 0xa6, 0xde, 0x7d, 0xc0, 0xeb, 0xeb, 0x5f, 0x8a, 0xe3, 0x2f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xdb, 0x73, 0xc9, 0xfc, 0x6c, 0x06, 0x00, 0x00,
}
