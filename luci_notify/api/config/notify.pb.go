// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/luci_notify/api/config/notify.proto

package config

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	proto1 "go.chromium.org/luci/buildbucket/proto"
	_ "go.chromium.org/luci/common/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// ProjectConfig is a luci-notify configuration for a particular project.
type ProjectConfig struct {
	// Notifiers is a list of Notifiers which watch builders and send
	// notifications for this project.
	Notifiers []*Notifier `protobuf:"bytes,1,rep,name=notifiers,proto3" json:"notifiers,omitempty"`
	// If false, then LUCI-Notify won't actually close trees, only log what
	// actions it would have taken.
	TreeClosingEnabled   bool     `protobuf:"varint,2,opt,name=tree_closing_enabled,json=treeClosingEnabled,proto3" json:"tree_closing_enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectConfig) Reset()         { *m = ProjectConfig{} }
func (m *ProjectConfig) String() string { return proto.CompactTextString(m) }
func (*ProjectConfig) ProtoMessage()    {}
func (*ProjectConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a6945a7af0ec43b, []int{0}
}

func (m *ProjectConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectConfig.Unmarshal(m, b)
}
func (m *ProjectConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectConfig.Marshal(b, m, deterministic)
}
func (m *ProjectConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectConfig.Merge(m, src)
}
func (m *ProjectConfig) XXX_Size() int {
	return xxx_messageInfo_ProjectConfig.Size(m)
}
func (m *ProjectConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectConfig proto.InternalMessageInfo

func (m *ProjectConfig) GetNotifiers() []*Notifier {
	if m != nil {
		return m.Notifiers
	}
	return nil
}

func (m *ProjectConfig) GetTreeClosingEnabled() bool {
	if m != nil {
		return m.TreeClosingEnabled
	}
	return false
}

// Notifier contains a set of notification configurations (which specify
// triggers to send notifications on) and a set of builders that will be
// watched for these triggers.
type Notifier struct {
	// Name is an identifier for the notifier which must be unique within a
	// project.
	//
	// Name must additionally match ^[a-z\-]+$, meaning it must only
	// use an alphabet of lowercase characters and hyphens.
	//
	// Required.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Notifications is a list of notification configurations.
	Notifications []*Notification `protobuf:"bytes,2,rep,name=notifications,proto3" json:"notifications,omitempty"`
	// Builders is a list of buildbucket builders this Notifier should watch.
	Builders []*Builder `protobuf:"bytes,3,rep,name=builders,proto3" json:"builders,omitempty"`
	// A list of tree closing rules to execute for this notifier.
	TreeClosers          []*TreeCloser `protobuf:"bytes,4,rep,name=tree_closers,json=treeClosers,proto3" json:"tree_closers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Notifier) Reset()         { *m = Notifier{} }
func (m *Notifier) String() string { return proto.CompactTextString(m) }
func (*Notifier) ProtoMessage()    {}
func (*Notifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a6945a7af0ec43b, []int{1}
}

func (m *Notifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notifier.Unmarshal(m, b)
}
func (m *Notifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notifier.Marshal(b, m, deterministic)
}
func (m *Notifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notifier.Merge(m, src)
}
func (m *Notifier) XXX_Size() int {
	return xxx_messageInfo_Notifier.Size(m)
}
func (m *Notifier) XXX_DiscardUnknown() {
	xxx_messageInfo_Notifier.DiscardUnknown(m)
}

var xxx_messageInfo_Notifier proto.InternalMessageInfo

func (m *Notifier) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Notifier) GetNotifications() []*Notification {
	if m != nil {
		return m.Notifications
	}
	return nil
}

func (m *Notifier) GetBuilders() []*Builder {
	if m != nil {
		return m.Builders
	}
	return nil
}

func (m *Notifier) GetTreeClosers() []*TreeCloser {
	if m != nil {
		return m.TreeClosers
	}
	return nil
}

// Notification specifies the triggers to watch for and send
// notifications on. It also specifies email recipients.
//
// Next ID: 13.
type Notification struct {
	// Deprecated. Notify on each build success.
	OnSuccess bool `protobuf:"varint,1,opt,name=on_success,json=onSuccess,proto3" json:"on_success,omitempty"`
	// Deprecated. Notify on each build failure.
	OnFailure bool `protobuf:"varint,2,opt,name=on_failure,json=onFailure,proto3" json:"on_failure,omitempty"`
	// Deprecated. Notify on each build status different than the previous one.
	OnChange bool `protobuf:"varint,3,opt,name=on_change,json=onChange,proto3" json:"on_change,omitempty"`
	// Deprecated. Notify on each build failure unless the previous build was a
	// failure.
	OnNewFailure bool `protobuf:"varint,7,opt,name=on_new_failure,json=onNewFailure,proto3" json:"on_new_failure,omitempty"`
	// Notify on each build with a specified status.
	OnOccurrence []proto1.Status `protobuf:"varint,9,rep,packed,name=on_occurrence,json=onOccurrence,proto3,enum=buildbucket.v2.Status" json:"on_occurrence,omitempty"`
	// Notify on each build with a specified status different than the previous
	// one.
	OnNewStatus []proto1.Status `protobuf:"varint,10,rep,packed,name=on_new_status,json=onNewStatus,proto3,enum=buildbucket.v2.Status" json:"on_new_status,omitempty"`
	// Notify only on builds which had a failing step matching this regular
	// expression. Mutually exclusive with "on_new_status".
	FailedStepRegexp string `protobuf:"bytes,11,opt,name=failed_step_regexp,json=failedStepRegexp,proto3" json:"failed_step_regexp,omitempty"`
	// Notify only on builds which don't have a failing step matching this regular
	// expression. May be combined with "failed_step_regexp", in which case it
	// must also have a failed step matching that regular expression. Mutually
	// exclusive with "on_new_status".
	FailedStepRegexpExclude string `protobuf:"bytes,12,opt,name=failed_step_regexp_exclude,json=failedStepRegexpExclude,proto3" json:"failed_step_regexp_exclude,omitempty"`
	// Email is the set of email addresses to notify.
	//
	// Optional.
	Email *Notification_Email `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	// Refers to which project template name to use to format this email.
	// If not present, "default" will be used.
	//
	// Optional.
	Template string `protobuf:"bytes,5,opt,name=template,proto3" json:"template,omitempty"`
	// NotifyBlamelist specifies whether to notify the computed blamelist for a
	// given build.
	//
	// If set, this notification will be sent to the blamelist of a build. Note
	// that if this is set in multiple notifications pertaining to the same
	// builder, the blamelist may receive multiple emails.
	//
	// Optional.
	NotifyBlamelist      *Notification_Blamelist `protobuf:"bytes,6,opt,name=notify_blamelist,json=notifyBlamelist,proto3" json:"notify_blamelist,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a6945a7af0ec43b, []int{2}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetOnSuccess() bool {
	if m != nil {
		return m.OnSuccess
	}
	return false
}

func (m *Notification) GetOnFailure() bool {
	if m != nil {
		return m.OnFailure
	}
	return false
}

func (m *Notification) GetOnChange() bool {
	if m != nil {
		return m.OnChange
	}
	return false
}

func (m *Notification) GetOnNewFailure() bool {
	if m != nil {
		return m.OnNewFailure
	}
	return false
}

func (m *Notification) GetOnOccurrence() []proto1.Status {
	if m != nil {
		return m.OnOccurrence
	}
	return nil
}

func (m *Notification) GetOnNewStatus() []proto1.Status {
	if m != nil {
		return m.OnNewStatus
	}
	return nil
}

func (m *Notification) GetFailedStepRegexp() string {
	if m != nil {
		return m.FailedStepRegexp
	}
	return ""
}

func (m *Notification) GetFailedStepRegexpExclude() string {
	if m != nil {
		return m.FailedStepRegexpExclude
	}
	return ""
}

func (m *Notification) GetEmail() *Notification_Email {
	if m != nil {
		return m.Email
	}
	return nil
}

func (m *Notification) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

func (m *Notification) GetNotifyBlamelist() *Notification_Blamelist {
	if m != nil {
		return m.NotifyBlamelist
	}
	return nil
}

// Email is a message representing a set of mail recipients.
type Notification_Email struct {
	// Recipients is a list of email addresses to notify.
	Recipients []string `protobuf:"bytes,1,rep,name=recipients,proto3" json:"recipients,omitempty"`
	// A list of rotations, for each of which we should notify the currently
	// active member.
	RotaNgRotations      []string `protobuf:"bytes,2,rep,name=rota_ng_rotations,json=rotaNgRotations,proto3" json:"rota_ng_rotations,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Notification_Email) Reset()         { *m = Notification_Email{} }
func (m *Notification_Email) String() string { return proto.CompactTextString(m) }
func (*Notification_Email) ProtoMessage()    {}
func (*Notification_Email) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a6945a7af0ec43b, []int{2, 0}
}

func (m *Notification_Email) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification_Email.Unmarshal(m, b)
}
func (m *Notification_Email) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification_Email.Marshal(b, m, deterministic)
}
func (m *Notification_Email) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification_Email.Merge(m, src)
}
func (m *Notification_Email) XXX_Size() int {
	return xxx_messageInfo_Notification_Email.Size(m)
}
func (m *Notification_Email) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification_Email.DiscardUnknown(m)
}

var xxx_messageInfo_Notification_Email proto.InternalMessageInfo

func (m *Notification_Email) GetRecipients() []string {
	if m != nil {
		return m.Recipients
	}
	return nil
}

func (m *Notification_Email) GetRotaNgRotations() []string {
	if m != nil {
		return m.RotaNgRotations
	}
	return nil
}

// Blamelist is a message representing configuration for notifying the
// blamelist.
type Notification_Blamelist struct {
	// A list of repositories which we are allowed to be included as part of the
	// blamelist. If unset, a blamelist will be computed based on a Builder's
	// repository field. If set, however luci-notify computes the blamelist for
	// all commits related to a build (which may span multiple repositories)
	// which are part of repository in this repository whitelist.
	//
	// Repositories should be valid Gerrit/Gitiles repository URLs, such as
	// https://chromium.googlesource.com/chromium/src
	//
	// Optional.
	RepositoryWhitelist  []string `protobuf:"bytes,1,rep,name=repository_whitelist,json=repositoryWhitelist,proto3" json:"repository_whitelist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Notification_Blamelist) Reset()         { *m = Notification_Blamelist{} }
func (m *Notification_Blamelist) String() string { return proto.CompactTextString(m) }
func (*Notification_Blamelist) ProtoMessage()    {}
func (*Notification_Blamelist) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a6945a7af0ec43b, []int{2, 1}
}

func (m *Notification_Blamelist) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification_Blamelist.Unmarshal(m, b)
}
func (m *Notification_Blamelist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification_Blamelist.Marshal(b, m, deterministic)
}
func (m *Notification_Blamelist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification_Blamelist.Merge(m, src)
}
func (m *Notification_Blamelist) XXX_Size() int {
	return xxx_messageInfo_Notification_Blamelist.Size(m)
}
func (m *Notification_Blamelist) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification_Blamelist.DiscardUnknown(m)
}

var xxx_messageInfo_Notification_Blamelist proto.InternalMessageInfo

func (m *Notification_Blamelist) GetRepositoryWhitelist() []string {
	if m != nil {
		return m.RepositoryWhitelist
	}
	return nil
}

// TreeCloser represents an action which closes a tree, by interfacing with an
// instance of the tree-status app.
type TreeCloser struct {
	// The hostname of the tree-status instance which this rule opens and closes.
	TreeStatusHost string `protobuf:"bytes,1,opt,name=tree_status_host,json=treeStatusHost,proto3" json:"tree_status_host,omitempty"`
	// Close the tree only on builds which had a failing step matching this
	// regular expression.
	FailedStepRegexp string `protobuf:"bytes,2,opt,name=failed_step_regexp,json=failedStepRegexp,proto3" json:"failed_step_regexp,omitempty"`
	// Close the tree only on builds which don't have a failing step matching this
	// regular expression. May be combined with "failed_step_regexp", in which
	// case it must also have a failed step matching that regular expression.
	FailedStepRegexpExclude string `protobuf:"bytes,3,opt,name=failed_step_regexp_exclude,json=failedStepRegexpExclude,proto3" json:"failed_step_regexp_exclude,omitempty"`
	// Refers to which project template name to use to format this email.
	// If not present, "default_tree_status" will be used.
	Template             string   `protobuf:"bytes,4,opt,name=template,proto3" json:"template,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TreeCloser) Reset()         { *m = TreeCloser{} }
func (m *TreeCloser) String() string { return proto.CompactTextString(m) }
func (*TreeCloser) ProtoMessage()    {}
func (*TreeCloser) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a6945a7af0ec43b, []int{3}
}

func (m *TreeCloser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TreeCloser.Unmarshal(m, b)
}
func (m *TreeCloser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TreeCloser.Marshal(b, m, deterministic)
}
func (m *TreeCloser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TreeCloser.Merge(m, src)
}
func (m *TreeCloser) XXX_Size() int {
	return xxx_messageInfo_TreeCloser.Size(m)
}
func (m *TreeCloser) XXX_DiscardUnknown() {
	xxx_messageInfo_TreeCloser.DiscardUnknown(m)
}

var xxx_messageInfo_TreeCloser proto.InternalMessageInfo

func (m *TreeCloser) GetTreeStatusHost() string {
	if m != nil {
		return m.TreeStatusHost
	}
	return ""
}

func (m *TreeCloser) GetFailedStepRegexp() string {
	if m != nil {
		return m.FailedStepRegexp
	}
	return ""
}

func (m *TreeCloser) GetFailedStepRegexpExclude() string {
	if m != nil {
		return m.FailedStepRegexpExclude
	}
	return ""
}

func (m *TreeCloser) GetTemplate() string {
	if m != nil {
		return m.Template
	}
	return ""
}

// Builder references a buildbucket builder in the current project.
type Builder struct {
	// Bucket is the buildbucket bucket that the builder is a part of.
	//
	// Required.
	Bucket string `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// Name is the name of the buildbucket builder.
	//
	// Required.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Repository is the git repository associated with this particular builder.
	//
	// The repository should look like a URL, e.g.
	// https://chromium.googlesource.com/src
	//
	// Currently, luci-notify only supports Gerrit-like URLs since it checks
	// against gitiles commits, so the URL's path (e.g. "src" in the above
	// example) should map directly to a Gerrit project.
	//
	// Builds attached to the history of this repository will use this
	// repository's git history to determine the order between two builds for the
	// OnChange notification.
	//
	// Optional.
	//
	// If not set, OnChange notifications will derive their notion of
	// "previous" build solely from build creation time, which is potentially
	// less reliable.
	Repository           string   `protobuf:"bytes,3,opt,name=repository,proto3" json:"repository,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Builder) Reset()         { *m = Builder{} }
func (m *Builder) String() string { return proto.CompactTextString(m) }
func (*Builder) ProtoMessage()    {}
func (*Builder) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a6945a7af0ec43b, []int{4}
}

func (m *Builder) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Builder.Unmarshal(m, b)
}
func (m *Builder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Builder.Marshal(b, m, deterministic)
}
func (m *Builder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Builder.Merge(m, src)
}
func (m *Builder) XXX_Size() int {
	return xxx_messageInfo_Builder.Size(m)
}
func (m *Builder) XXX_DiscardUnknown() {
	xxx_messageInfo_Builder.DiscardUnknown(m)
}

var xxx_messageInfo_Builder proto.InternalMessageInfo

func (m *Builder) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *Builder) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Builder) GetRepository() string {
	if m != nil {
		return m.Repository
	}
	return ""
}

// Notifications encapsulates a list of notifications as a proto so code for
// storing it in the datastore may be generated.
type Notifications struct {
	// Notifications is a list of notification configurations.
	Notifications        []*Notification `protobuf:"bytes,1,rep,name=notifications,proto3" json:"notifications,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Notifications) Reset()         { *m = Notifications{} }
func (m *Notifications) String() string { return proto.CompactTextString(m) }
func (*Notifications) ProtoMessage()    {}
func (*Notifications) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a6945a7af0ec43b, []int{5}
}

func (m *Notifications) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notifications.Unmarshal(m, b)
}
func (m *Notifications) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notifications.Marshal(b, m, deterministic)
}
func (m *Notifications) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notifications.Merge(m, src)
}
func (m *Notifications) XXX_Size() int {
	return xxx_messageInfo_Notifications.Size(m)
}
func (m *Notifications) XXX_DiscardUnknown() {
	xxx_messageInfo_Notifications.DiscardUnknown(m)
}

var xxx_messageInfo_Notifications proto.InternalMessageInfo

func (m *Notifications) GetNotifications() []*Notification {
	if m != nil {
		return m.Notifications
	}
	return nil
}

// A collection of landed Git commits hosted on Gitiles.
type GitilesCommits struct {
	// The Gitiles commits in this collection.
	Commits              []*proto1.GitilesCommit `protobuf:"bytes,1,rep,name=commits,proto3" json:"commits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *GitilesCommits) Reset()         { *m = GitilesCommits{} }
func (m *GitilesCommits) String() string { return proto.CompactTextString(m) }
func (*GitilesCommits) ProtoMessage()    {}
func (*GitilesCommits) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a6945a7af0ec43b, []int{6}
}

func (m *GitilesCommits) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GitilesCommits.Unmarshal(m, b)
}
func (m *GitilesCommits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GitilesCommits.Marshal(b, m, deterministic)
}
func (m *GitilesCommits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GitilesCommits.Merge(m, src)
}
func (m *GitilesCommits) XXX_Size() int {
	return xxx_messageInfo_GitilesCommits.Size(m)
}
func (m *GitilesCommits) XXX_DiscardUnknown() {
	xxx_messageInfo_GitilesCommits.DiscardUnknown(m)
}

var xxx_messageInfo_GitilesCommits proto.InternalMessageInfo

func (m *GitilesCommits) GetCommits() []*proto1.GitilesCommit {
	if m != nil {
		return m.Commits
	}
	return nil
}

// Input to an email template.
type TemplateInput struct {
	// Buildbucket hostname, e.g. "cr-buildbucket.appspot.com".
	BuildbucketHostname string `protobuf:"bytes,1,opt,name=buildbucket_hostname,json=buildbucketHostname,proto3" json:"buildbucket_hostname,omitempty"`
	// The completed build.
	Build *proto1.Build `protobuf:"bytes,2,opt,name=build,proto3" json:"build,omitempty"`
	// State of the previous build in this builder.
	OldStatus proto1.Status `protobuf:"varint,3,opt,name=old_status,json=oldStatus,proto3,enum=buildbucket.v2.Status" json:"old_status,omitempty"`
	// The failed steps that passed the given regexes (see the fields
	// "failed_step_regexp" and "failed_step_regexp_exclude" above). If that field
	// wasn't supplied, this will be empty.
	MatchingFailedSteps  []*proto1.Step `protobuf:"bytes,4,rep,name=matching_failed_steps,json=matchingFailedSteps,proto3" json:"matching_failed_steps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *TemplateInput) Reset()         { *m = TemplateInput{} }
func (m *TemplateInput) String() string { return proto.CompactTextString(m) }
func (*TemplateInput) ProtoMessage()    {}
func (*TemplateInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a6945a7af0ec43b, []int{7}
}

func (m *TemplateInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemplateInput.Unmarshal(m, b)
}
func (m *TemplateInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemplateInput.Marshal(b, m, deterministic)
}
func (m *TemplateInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemplateInput.Merge(m, src)
}
func (m *TemplateInput) XXX_Size() int {
	return xxx_messageInfo_TemplateInput.Size(m)
}
func (m *TemplateInput) XXX_DiscardUnknown() {
	xxx_messageInfo_TemplateInput.DiscardUnknown(m)
}

var xxx_messageInfo_TemplateInput proto.InternalMessageInfo

func (m *TemplateInput) GetBuildbucketHostname() string {
	if m != nil {
		return m.BuildbucketHostname
	}
	return ""
}

func (m *TemplateInput) GetBuild() *proto1.Build {
	if m != nil {
		return m.Build
	}
	return nil
}

func (m *TemplateInput) GetOldStatus() proto1.Status {
	if m != nil {
		return m.OldStatus
	}
	return proto1.Status_STATUS_UNSPECIFIED
}

func (m *TemplateInput) GetMatchingFailedSteps() []*proto1.Step {
	if m != nil {
		return m.MatchingFailedSteps
	}
	return nil
}

func init() {
	proto.RegisterType((*ProjectConfig)(nil), "notify.ProjectConfig")
	proto.RegisterType((*Notifier)(nil), "notify.Notifier")
	proto.RegisterType((*Notification)(nil), "notify.Notification")
	proto.RegisterType((*Notification_Email)(nil), "notify.Notification.Email")
	proto.RegisterType((*Notification_Blamelist)(nil), "notify.Notification.Blamelist")
	proto.RegisterType((*TreeCloser)(nil), "notify.TreeCloser")
	proto.RegisterType((*Builder)(nil), "notify.Builder")
	proto.RegisterType((*Notifications)(nil), "notify.Notifications")
	proto.RegisterType((*GitilesCommits)(nil), "notify.GitilesCommits")
	proto.RegisterType((*TemplateInput)(nil), "notify.TemplateInput")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/luci_notify/api/config/notify.proto", fileDescriptor_9a6945a7af0ec43b)
}

var fileDescriptor_9a6945a7af0ec43b = []byte{
	// 881 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xef, 0x6e, 0xe3, 0x44,
	0x10, 0x97, 0x9b, 0x3f, 0x8d, 0x27, 0x4d, 0x1b, 0xb6, 0xbd, 0xc3, 0x0a, 0xba, 0x53, 0x14, 0xf8,
	0x10, 0x71, 0xe0, 0xdc, 0xe5, 0x74, 0x02, 0xf5, 0x24, 0x84, 0x5a, 0xf5, 0x68, 0x41, 0x0a, 0xc8,
	0x3d, 0x84, 0xc4, 0x17, 0xcb, 0xd9, 0x6c, 0x9d, 0x05, 0x7b, 0xd7, 0x78, 0xd7, 0xf4, 0xee, 0x75,
	0xf8, 0xc2, 0x53, 0xc0, 0x33, 0xf1, 0x04, 0x08, 0x79, 0x76, 0xed, 0x3a, 0x6d, 0xee, 0x54, 0xf1,
	0x25, 0xf1, 0xce, 0xef, 0xf7, 0x9b, 0x9d, 0x9d, 0x99, 0x9d, 0x85, 0x2f, 0x63, 0xe9, 0xd3, 0x75,
	0x2e, 0x53, 0x5e, 0xa4, 0xbe, 0xcc, 0xe3, 0x59, 0x52, 0x50, 0x8e, 0x3f, 0xa1, 0x90, 0x9a, 0x5f,
	0xbd, 0x9d, 0x45, 0x19, 0x9f, 0x51, 0x29, 0xae, 0x78, 0x3c, 0x33, 0x16, 0x3f, 0xcb, 0xa5, 0x96,
	0xa4, 0x6b, 0x56, 0xa3, 0xf9, 0x56, 0x0f, 0xcb, 0x82, 0x27, 0xab, 0x65, 0x41, 0x7f, 0x65, 0x7a,
	0x86, 0x7c, 0x63, 0x31, 0xda, 0xd1, 0xf3, 0x7b, 0x6a, 0xa8, 0x4c, 0x53, 0x29, 0xac, 0xe8, 0xd9,
	0x3d, 0x45, 0x4a, 0xb3, 0xcc, 0x4a, 0x66, 0x5b, 0x25, 0xc6, 0xab, 0x65, 0xcb, 0x4c, 0x73, 0x29,
	0x94, 0x11, 0x4c, 0x7e, 0x83, 0xc1, 0x0f, 0xb9, 0xfc, 0x85, 0x51, 0x7d, 0x8a, 0x47, 0x26, 0x3e,
	0xb8, 0x78, 0x4e, 0xce, 0x72, 0xe5, 0x39, 0xe3, 0xd6, 0xb4, 0x3f, 0x1f, 0xfa, 0x36, 0x0f, 0x0b,
	0x0b, 0x04, 0x37, 0x14, 0xf2, 0x14, 0x8e, 0x74, 0xce, 0x58, 0x48, 0x13, 0xa9, 0xb8, 0x88, 0x43,
	0x26, 0xa2, 0x65, 0xc2, 0x56, 0xde, 0xce, 0xd8, 0x99, 0xf6, 0x02, 0x52, 0x62, 0xa7, 0x06, 0x3a,
	0x33, 0xc8, 0xe4, 0x6f, 0x07, 0x7a, 0x95, 0x27, 0x42, 0xa0, 0x2d, 0xa2, 0x94, 0x79, 0xce, 0xd8,
	0x99, 0xba, 0x01, 0x7e, 0x93, 0x63, 0x18, 0x18, 0xff, 0x34, 0xc2, 0x50, 0xbd, 0x1d, 0x0c, 0xe3,
	0x68, 0x33, 0x0c, 0x03, 0x06, 0x9b, 0x54, 0xf2, 0x04, 0x7a, 0x98, 0xa0, 0x32, 0xfa, 0x16, 0xca,
	0x0e, 0x2a, 0xd9, 0x89, 0xb1, 0x07, 0x35, 0x81, 0xbc, 0x80, 0xbd, 0x3a, 0xf6, 0x52, 0xd0, 0x46,
	0x01, 0xa9, 0x04, 0xaf, 0x6d, 0xec, 0x2c, 0x0f, 0xfa, 0xba, 0xfe, 0x56, 0x93, 0x3f, 0x3b, 0xb0,
	0xd7, 0x8c, 0x81, 0x3c, 0x02, 0x90, 0x22, 0x54, 0x05, 0xa5, 0x4c, 0x29, 0x3c, 0x4a, 0x2f, 0x70,
	0xa5, 0xb8, 0x34, 0x06, 0x0b, 0x5f, 0x45, 0x3c, 0x29, 0x72, 0x66, 0x13, 0xe3, 0x4a, 0xf1, 0xca,
	0x18, 0xc8, 0x47, 0xe0, 0x4a, 0x11, 0xd2, 0x75, 0x24, 0x62, 0xe6, 0xb5, 0x10, 0xed, 0x49, 0x71,
	0x8a, 0x6b, 0xf2, 0x09, 0xec, 0x4b, 0x11, 0x0a, 0x76, 0x5d, 0xeb, 0x77, 0x91, 0xb1, 0x27, 0xc5,
	0x82, 0x5d, 0x57, 0x2e, 0x5e, 0xc2, 0x40, 0x8a, 0x50, 0x52, 0x5a, 0xe4, 0x39, 0x13, 0x94, 0x79,
	0xee, 0xb8, 0x35, 0xdd, 0x9f, 0x3f, 0xf4, 0x1b, 0xcd, 0xe2, 0xff, 0x3e, 0xf7, 0x2f, 0x75, 0xa4,
	0x0b, 0x55, 0x8a, 0xbf, 0xaf, 0xb9, 0x65, 0xba, 0xed, 0x16, 0x0a, 0x61, 0x0f, 0xde, 0x2b, 0xee,
	0xe3, 0xce, 0x66, 0x41, 0x3e, 0x03, 0x52, 0xc6, 0xc5, 0x56, 0x61, 0xd9, 0x84, 0x61, 0xce, 0x62,
	0xf6, 0x26, 0xf3, 0xfa, 0x58, 0xcc, 0xa1, 0x41, 0x2e, 0x35, 0xcb, 0x02, 0xb4, 0x93, 0x97, 0x30,
	0xba, 0xcb, 0x0e, 0xd9, 0x1b, 0x9a, 0x14, 0x2b, 0xe6, 0xed, 0xa1, 0xea, 0xc3, 0xdb, 0xaa, 0x33,
	0x03, 0x93, 0xa7, 0xd0, 0x61, 0x69, 0xc4, 0x13, 0xaf, 0x3d, 0x76, 0xa6, 0xfd, 0xf9, 0x68, 0x5b,
	0x37, 0xf8, 0x67, 0x25, 0x23, 0x30, 0x44, 0x32, 0x82, 0x9e, 0x66, 0x69, 0x96, 0x44, 0x9a, 0x79,
	0x1d, 0x74, 0x5e, 0xaf, 0xc9, 0x05, 0x0c, 0x8d, 0x3e, 0x5c, 0x26, 0x51, 0xca, 0x12, 0xae, 0xb4,
	0xd7, 0x45, 0xc7, 0x8f, 0xb7, 0x3a, 0x3e, 0xa9, 0x58, 0xc1, 0x81, 0x81, 0x6b, 0xc3, 0xe8, 0x12,
	0x3a, 0xb8, 0x2d, 0x79, 0x0c, 0x90, 0x33, 0xca, 0x33, 0xce, 0x84, 0x36, 0x77, 0xc7, 0x0d, 0x1a,
	0x16, 0xf2, 0x29, 0x7c, 0x90, 0x4b, 0x1d, 0x85, 0x22, 0x0e, 0xcb, 0xff, 0x9b, 0xde, 0x76, 0x83,
	0x83, 0xd2, 0xb0, 0x88, 0x83, 0xca, 0x3c, 0xfa, 0x0a, 0xdc, 0x7a, 0x07, 0xf2, 0x0c, 0x8e, 0x72,
	0x96, 0x49, 0xc5, 0xb5, 0xcc, 0xdf, 0x86, 0xd7, 0x6b, 0xae, 0x4d, 0xc0, 0x66, 0x8b, 0xc3, 0x1b,
	0xec, 0xa7, 0x0a, 0xfa, 0xb6, 0xdd, 0xeb, 0x0d, 0xdd, 0xc9, 0x5f, 0x0e, 0xc0, 0x4d, 0x17, 0x93,
	0x29, 0x0c, 0xb1, 0xdf, 0x4d, 0x9d, 0xc3, 0xb5, 0x44, 0x1f, 0x65, 0x62, 0xf6, 0x4b, 0xbb, 0xa9,
	0xe9, 0xb9, 0x54, 0xfa, 0x1d, 0x75, 0xdd, 0xf9, 0x5f, 0x75, 0x6d, 0xbd, 0xbf, 0xae, 0xcd, 0x2a,
	0xb5, 0x37, 0xab, 0x34, 0xf9, 0x11, 0x76, 0xed, 0xad, 0x25, 0x0f, 0xa1, 0x6b, 0x5a, 0xd1, 0x46,
	0x6c, 0x57, 0xf5, 0x00, 0xd9, 0x69, 0x0c, 0x10, 0x2c, 0x44, 0x95, 0x13, 0xbb, 0x7f, 0xc3, 0x32,
	0xf9, 0x0e, 0x06, 0x8b, 0x8d, 0xa9, 0x71, 0x67, 0xe2, 0x38, 0xf7, 0x9e, 0x38, 0x93, 0x0b, 0xd8,
	0xff, 0x86, 0x6b, 0x9e, 0x30, 0x75, 0x2a, 0xd3, 0x94, 0x6b, 0x45, 0xbe, 0x80, 0x5d, 0x6a, 0x3e,
	0xad, 0x9f, 0x47, 0xb7, 0xaf, 0xd2, 0x86, 0x20, 0xa8, 0xd8, 0x93, 0x7f, 0x1c, 0x18, 0xbc, 0xb6,
	0x67, 0xbf, 0x10, 0x59, 0x81, 0x95, 0x6f, 0x48, 0xb1, 0x62, 0x8d, 0x71, 0x79, 0xd8, 0xc0, 0xce,
	0x2d, 0x44, 0x9e, 0x40, 0x07, 0xcd, 0x98, 0x91, 0xfe, 0xfc, 0xc1, 0xed, 0xbd, 0x31, 0xa1, 0x81,
	0xe1, 0x90, 0x17, 0x00, 0x32, 0x59, 0x55, 0x17, 0xbf, 0xcc, 0xd4, 0xbb, 0x2f, 0xbe, 0x2b, 0x93,
	0x95, 0xbd, 0xf6, 0xe7, 0xf0, 0x20, 0x8d, 0x34, 0x5d, 0x97, 0x03, 0xbf, 0x51, 0xf9, 0x6a, 0x82,
	0x1e, 0xdd, 0xf5, 0xc0, 0xb2, 0xe0, 0xb0, 0x92, 0xbc, 0xaa, 0x3b, 0x41, 0x9d, 0x2c, 0xfe, 0xf8,
	0xf7, 0xe3, 0x13, 0xf8, 0x7a, 0xad, 0x75, 0xa6, 0x8e, 0x67, 0xf8, 0x5c, 0x7d, 0x6e, 0x1e, 0x5f,
	0x3f, 0xca, 0x32, 0x95, 0x49, 0xed, 0x53, 0x99, 0xce, 0x14, 0x5d, 0xb3, 0x34, 0x52, 0xe5, 0x13,
	0x56, 0xbe, 0x55, 0xea, 0x18, 0x89, 0xb6, 0x3a, 0xf4, 0x2a, 0xfe, 0xb9, 0x6b, 0x44, 0xcb, 0x2e,
	0x3e, 0x6b, 0xcf, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x41, 0xed, 0x76, 0xe7, 0x07, 0x00,
	0x00,
}
