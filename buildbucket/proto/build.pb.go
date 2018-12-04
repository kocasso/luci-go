// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/buildbucket/proto/build.proto

package buildbucketpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// A single build, identified by an int64 id.
// Belongs to a builder.
//
// RPC: see Builds service for build creation and retrieval.
// Some Build fields are marked as excluded from responses by default.
// Use build_fields request field to specify that a field must be included.
//
// BigQuery: this message also defines schema of a BigQuery table of completed builds.
// A BigQuery row is inserted soon after build ends, i.e. a row represents a state of
// a build at completion time and does not change after that.
// All fields are included.
type Build struct {
	// Identifier of the build, unique per LUCI deployment.
	// IDs are monotonically decreasing.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Required. The builder this build belongs to.
	//
	// Tuple (builder.project, builder.bucket) defines build ACL
	// which may change after build has ended.
	Builder *BuilderID `protobuf:"bytes,2,opt,name=builder,proto3" json:"builder,omitempty"`
	// Human-oriented identifier of the build with the following properties:
	// - unique within the builder
	// - a monotonically increasing number
	// - mostly contiguous
	// - much shorter than id
	//
	// Caution: populated (positive number) iff build numbers were enabled
	// in the builder configuration at the time of build creation.
	//
	// Caution: Build numbers are not guaranteed to be contiguous.
	// There may be gaps during outages.
	//
	// Caution: Build numbers, while monotonically increasing, do not
	// necessarily reflect source-code order. For example, force builds
	// or rebuilds can allocate new, higher, numbers, but build an older-
	// than-HEAD version of the source.
	Number int32 `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	// Verified identity which created this build.
	CreatedBy string `protobuf:"bytes,4,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	// When the build was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// When the build started.
	// Required iff status is STARTED, SUCCESS or FAILURE.
	StartTime *timestamp.Timestamp `protobuf:"bytes,7,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// When the build ended.
	// Present iff status is terminal.
	// MUST NOT be before start_time.
	EndTime *timestamp.Timestamp `protobuf:"bytes,8,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	// When the build was most recently updated.
	//
	// RPC: can be > end_time if, e.g. new tags were attached to a completed
	// build.
	UpdateTime *timestamp.Timestamp `protobuf:"bytes,9,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Status of the build.
	// Must be specified, i.e. not STATUS_UNSPECIFIED.
	//
	// RPC: Responses have most current status.
	//
	// BigQuery: Final status of the build. Cannot be SCHEDULED or STARTED.
	Status Status `protobuf:"varint,12,opt,name=status,proto3,enum=buildbucket.v2.Status" json:"status,omitempty"`
	// Explanation of the current status.
	//
	// Types that are valid to be assigned to StatusReason:
	//	*Build_InfraFailureReason
	//	*Build_CancelReason
	StatusReason isBuild_StatusReason `protobuf_oneof:"status_reason"`
	// Input to the build script / recipe.
	Input *Build_Input `protobuf:"bytes,15,opt,name=input,proto3" json:"input,omitempty"`
	// Output of the build script / recipe.
	// SHOULD depend only on input field and NOT other fields.
	//
	// RPC: By default, this field is excluded from responses.
	// Updated while the build is running and finalized when the build ends.
	Output *Build_Output `protobuf:"bytes,16,opt,name=output,proto3" json:"output,omitempty"`
	// Current list of build steps.
	// Updated as build runs.
	//
	// RPC: By default, this field is excluded from responses.
	Steps []*Step `protobuf:"bytes,17,rep,name=steps,proto3" json:"steps,omitempty"`
	// Build infrastructure used by the build.
	//
	// RPC: By default, this field is excluded from responses.
	Infra *BuildInfra `protobuf:"bytes,18,opt,name=infra,proto3" json:"infra,omitempty"`
	// Arbitrary annotations for the build.
	// One key may have multiple values, which is why this is not a map<string,string>.
	// Indexed by the server, see also BuildFilter.tags.
	Tags                 []*StringPair `protobuf:"bytes,19,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Build) Reset()         { *m = Build{} }
func (m *Build) String() string { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()    {}
func (*Build) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3b43e011b474f4, []int{0}
}

func (m *Build) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Build.Unmarshal(m, b)
}
func (m *Build) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Build.Marshal(b, m, deterministic)
}
func (m *Build) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Build.Merge(m, src)
}
func (m *Build) XXX_Size() int {
	return xxx_messageInfo_Build.Size(m)
}
func (m *Build) XXX_DiscardUnknown() {
	xxx_messageInfo_Build.DiscardUnknown(m)
}

var xxx_messageInfo_Build proto.InternalMessageInfo

func (m *Build) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Build) GetBuilder() *BuilderID {
	if m != nil {
		return m.Builder
	}
	return nil
}

func (m *Build) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Build) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *Build) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Build) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Build) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *Build) GetUpdateTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpdateTime
	}
	return nil
}

func (m *Build) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_STATUS_UNSPECIFIED
}

type isBuild_StatusReason interface {
	isBuild_StatusReason()
}

type Build_InfraFailureReason struct {
	InfraFailureReason *InfraFailureReason `protobuf:"bytes,13,opt,name=infra_failure_reason,json=infraFailureReason,proto3,oneof"`
}

type Build_CancelReason struct {
	CancelReason *CancelReason `protobuf:"bytes,14,opt,name=cancel_reason,json=cancelReason,proto3,oneof"`
}

func (*Build_InfraFailureReason) isBuild_StatusReason() {}

func (*Build_CancelReason) isBuild_StatusReason() {}

func (m *Build) GetStatusReason() isBuild_StatusReason {
	if m != nil {
		return m.StatusReason
	}
	return nil
}

func (m *Build) GetInfraFailureReason() *InfraFailureReason {
	if x, ok := m.GetStatusReason().(*Build_InfraFailureReason); ok {
		return x.InfraFailureReason
	}
	return nil
}

func (m *Build) GetCancelReason() *CancelReason {
	if x, ok := m.GetStatusReason().(*Build_CancelReason); ok {
		return x.CancelReason
	}
	return nil
}

func (m *Build) GetInput() *Build_Input {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *Build) GetOutput() *Build_Output {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *Build) GetSteps() []*Step {
	if m != nil {
		return m.Steps
	}
	return nil
}

func (m *Build) GetInfra() *BuildInfra {
	if m != nil {
		return m.Infra
	}
	return nil
}

func (m *Build) GetTags() []*StringPair {
	if m != nil {
		return m.Tags
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Build) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Build_InfraFailureReason)(nil),
		(*Build_CancelReason)(nil),
	}
}

// Defines what to build/test.
type Build_Input struct {
	// Arbitrary JSON object. Available at build run time.
	//
	// RPC: By default, this field is excluded from responses.
	//
	// V1 equivalent: corresponds to "properties" key in "parameters_json".
	Properties *_struct.Struct `protobuf:"bytes,1,opt,name=properties,proto3" json:"properties,omitempty"`
	// The Gitiles commit to run against.
	// Usually present in CI builds, set by LUCI Scheduler.
	// If not present, the build may checkout "refs/heads/master".
	// NOT a blamelist.
	//
	// V1 equivalent: supersedes "revision" property and "buildset"
	// tag that starts with "commit/gitiles/".
	GitilesCommit *GitilesCommit `protobuf:"bytes,2,opt,name=gitiles_commit,json=gitilesCommit,proto3" json:"gitiles_commit,omitempty"`
	// Gerrit patchsets to run against.
	// Usually present in tryjobs, set by CQ, Gerrit, git-cl-try.
	// Applied on top of gitiles_commit if specified, otherwise tip of the tree.
	//
	// V1 equivalent: supersedes patch_* properties and "buildset"
	// tag that starts with "patch/gerrit/".
	GerritChanges []*GerritChange `protobuf:"bytes,3,rep,name=gerrit_changes,json=gerritChanges,proto3" json:"gerrit_changes,omitempty"`
	// If true, the build does not affect prod. In recipe land, runtime.is_experimental will
	// return true and recipes should not make prod-visible side effects.
	// By default, experimental builds are not surfaced in RPCs, PubSub
	// notifications (unless it is callback), and reported in metrics / BigQuery tables
	// under a different name.
	// See also include_experimental fields to in request messages.
	Experimental         bool     `protobuf:"varint,5,opt,name=experimental,proto3" json:"experimental,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Build_Input) Reset()         { *m = Build_Input{} }
func (m *Build_Input) String() string { return proto.CompactTextString(m) }
func (*Build_Input) ProtoMessage()    {}
func (*Build_Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3b43e011b474f4, []int{0, 0}
}

func (m *Build_Input) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Build_Input.Unmarshal(m, b)
}
func (m *Build_Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Build_Input.Marshal(b, m, deterministic)
}
func (m *Build_Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Build_Input.Merge(m, src)
}
func (m *Build_Input) XXX_Size() int {
	return xxx_messageInfo_Build_Input.Size(m)
}
func (m *Build_Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Build_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Build_Input proto.InternalMessageInfo

func (m *Build_Input) GetProperties() *_struct.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *Build_Input) GetGitilesCommit() *GitilesCommit {
	if m != nil {
		return m.GitilesCommit
	}
	return nil
}

func (m *Build_Input) GetGerritChanges() []*GerritChange {
	if m != nil {
		return m.GerritChanges
	}
	return nil
}

func (m *Build_Input) GetExperimental() bool {
	if m != nil {
		return m.Experimental
	}
	return false
}

// Output of the build script / recipe.
type Build_Output struct {
	// Arbitrary JSON object produced by the build.
	//
	// V1 equivalent: corresponds to "properties" key in
	// "result_details_json".
	// In V1 output properties are not populated until build ends.
	Properties *_struct.Struct `protobuf:"bytes,1,opt,name=properties,proto3" json:"properties,omitempty"`
	// Human-oriented summary of the build provided by the build itself,
	// in Markdown format (https://spec.commonmark.org/0.28/).
	//
	// BigQuery: excluded from rows.
	SummaryMarkdown string `protobuf:"bytes,2,opt,name=summary_markdown,json=summaryMarkdown,proto3" json:"summary_markdown,omitempty"`
	// Build checked out and executed on this commit.
	//
	// Should correspond to Build.Input.gitiles_commit.
	// May be present even if Build.Input.gitiles_commit is not set, for example
	// in cron builders.
	//
	// V1 equivalent: this supersedes all got_revision output property.
	GitilesCommit        *GitilesCommit `protobuf:"bytes,3,opt,name=gitiles_commit,json=gitilesCommit,proto3" json:"gitiles_commit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Build_Output) Reset()         { *m = Build_Output{} }
func (m *Build_Output) String() string { return proto.CompactTextString(m) }
func (*Build_Output) ProtoMessage()    {}
func (*Build_Output) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3b43e011b474f4, []int{0, 1}
}

func (m *Build_Output) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Build_Output.Unmarshal(m, b)
}
func (m *Build_Output) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Build_Output.Marshal(b, m, deterministic)
}
func (m *Build_Output) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Build_Output.Merge(m, src)
}
func (m *Build_Output) XXX_Size() int {
	return xxx_messageInfo_Build_Output.Size(m)
}
func (m *Build_Output) XXX_DiscardUnknown() {
	xxx_messageInfo_Build_Output.DiscardUnknown(m)
}

var xxx_messageInfo_Build_Output proto.InternalMessageInfo

func (m *Build_Output) GetProperties() *_struct.Struct {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *Build_Output) GetSummaryMarkdown() string {
	if m != nil {
		return m.SummaryMarkdown
	}
	return ""
}

func (m *Build_Output) GetGitilesCommit() *GitilesCommit {
	if m != nil {
		return m.GitilesCommit
	}
	return nil
}

// Explains why status is CANCELED.
type CancelReason struct {
	// Human-oriented reasoning.
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// Verified identity who canceled this build.
	CanceledBy           string   `protobuf:"bytes,2,opt,name=canceled_by,json=canceledBy,proto3" json:"canceled_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelReason) Reset()         { *m = CancelReason{} }
func (m *CancelReason) String() string { return proto.CompactTextString(m) }
func (*CancelReason) ProtoMessage()    {}
func (*CancelReason) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3b43e011b474f4, []int{1}
}

func (m *CancelReason) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelReason.Unmarshal(m, b)
}
func (m *CancelReason) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelReason.Marshal(b, m, deterministic)
}
func (m *CancelReason) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelReason.Merge(m, src)
}
func (m *CancelReason) XXX_Size() int {
	return xxx_messageInfo_CancelReason.Size(m)
}
func (m *CancelReason) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelReason.DiscardUnknown(m)
}

var xxx_messageInfo_CancelReason proto.InternalMessageInfo

func (m *CancelReason) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CancelReason) GetCanceledBy() string {
	if m != nil {
		return m.CanceledBy
	}
	return ""
}

// Explains why status is INFRA_FAILURE.
type InfraFailureReason struct {
	// Human-oriented explanation of the infrastructure failure.
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// Indicates that the failure was due to a resource exhaustion / quota denial.
	ResourceExhaustion   bool     `protobuf:"varint,2,opt,name=resource_exhaustion,json=resourceExhaustion,proto3" json:"resource_exhaustion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InfraFailureReason) Reset()         { *m = InfraFailureReason{} }
func (m *InfraFailureReason) String() string { return proto.CompactTextString(m) }
func (*InfraFailureReason) ProtoMessage()    {}
func (*InfraFailureReason) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3b43e011b474f4, []int{2}
}

func (m *InfraFailureReason) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InfraFailureReason.Unmarshal(m, b)
}
func (m *InfraFailureReason) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InfraFailureReason.Marshal(b, m, deterministic)
}
func (m *InfraFailureReason) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InfraFailureReason.Merge(m, src)
}
func (m *InfraFailureReason) XXX_Size() int {
	return xxx_messageInfo_InfraFailureReason.Size(m)
}
func (m *InfraFailureReason) XXX_DiscardUnknown() {
	xxx_messageInfo_InfraFailureReason.DiscardUnknown(m)
}

var xxx_messageInfo_InfraFailureReason proto.InternalMessageInfo

func (m *InfraFailureReason) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *InfraFailureReason) GetResourceExhaustion() bool {
	if m != nil {
		return m.ResourceExhaustion
	}
	return false
}

// Build infrastructure that was used for a particular build.
type BuildInfra struct {
	Buildbucket          *BuildInfra_Buildbucket `protobuf:"bytes,1,opt,name=buildbucket,proto3" json:"buildbucket,omitempty"`
	Swarming             *BuildInfra_Swarming    `protobuf:"bytes,2,opt,name=swarming,proto3" json:"swarming,omitempty"`
	Logdog               *BuildInfra_LogDog      `protobuf:"bytes,3,opt,name=logdog,proto3" json:"logdog,omitempty"`
	Recipe               *BuildInfra_Recipe      `protobuf:"bytes,4,opt,name=recipe,proto3" json:"recipe,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *BuildInfra) Reset()         { *m = BuildInfra{} }
func (m *BuildInfra) String() string { return proto.CompactTextString(m) }
func (*BuildInfra) ProtoMessage()    {}
func (*BuildInfra) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3b43e011b474f4, []int{3}
}

func (m *BuildInfra) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildInfra.Unmarshal(m, b)
}
func (m *BuildInfra) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildInfra.Marshal(b, m, deterministic)
}
func (m *BuildInfra) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildInfra.Merge(m, src)
}
func (m *BuildInfra) XXX_Size() int {
	return xxx_messageInfo_BuildInfra.Size(m)
}
func (m *BuildInfra) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildInfra.DiscardUnknown(m)
}

var xxx_messageInfo_BuildInfra proto.InternalMessageInfo

func (m *BuildInfra) GetBuildbucket() *BuildInfra_Buildbucket {
	if m != nil {
		return m.Buildbucket
	}
	return nil
}

func (m *BuildInfra) GetSwarming() *BuildInfra_Swarming {
	if m != nil {
		return m.Swarming
	}
	return nil
}

func (m *BuildInfra) GetLogdog() *BuildInfra_LogDog {
	if m != nil {
		return m.Logdog
	}
	return nil
}

func (m *BuildInfra) GetRecipe() *BuildInfra_Recipe {
	if m != nil {
		return m.Recipe
	}
	return nil
}

// Buildbucket-specific information, captured at the build creation time.
type BuildInfra_Buildbucket struct {
	// Version of swarming task template. Defines
	// versions of kitchen, git, git wrapper, python, vpython, etc.
	ServiceConfigRevision string `protobuf:"bytes,2,opt,name=service_config_revision,json=serviceConfigRevision,proto3" json:"service_config_revision,omitempty"`
	// Whether canary version of the swarming task template was used for this
	// build.
	Canary               bool     `protobuf:"varint,4,opt,name=canary,proto3" json:"canary,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildInfra_Buildbucket) Reset()         { *m = BuildInfra_Buildbucket{} }
func (m *BuildInfra_Buildbucket) String() string { return proto.CompactTextString(m) }
func (*BuildInfra_Buildbucket) ProtoMessage()    {}
func (*BuildInfra_Buildbucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3b43e011b474f4, []int{3, 0}
}

func (m *BuildInfra_Buildbucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildInfra_Buildbucket.Unmarshal(m, b)
}
func (m *BuildInfra_Buildbucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildInfra_Buildbucket.Marshal(b, m, deterministic)
}
func (m *BuildInfra_Buildbucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildInfra_Buildbucket.Merge(m, src)
}
func (m *BuildInfra_Buildbucket) XXX_Size() int {
	return xxx_messageInfo_BuildInfra_Buildbucket.Size(m)
}
func (m *BuildInfra_Buildbucket) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildInfra_Buildbucket.DiscardUnknown(m)
}

var xxx_messageInfo_BuildInfra_Buildbucket proto.InternalMessageInfo

func (m *BuildInfra_Buildbucket) GetServiceConfigRevision() string {
	if m != nil {
		return m.ServiceConfigRevision
	}
	return ""
}

func (m *BuildInfra_Buildbucket) GetCanary() bool {
	if m != nil {
		return m.Canary
	}
	return false
}

// Swarming-specific information.
type BuildInfra_Swarming struct {
	// Swarming hostname, e.g. "chromium-swarm.appspot.com".
	// Populated at the build creation time.
	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// Swarming task id.
	// Not guaranteed to be populated at the build creation time.
	TaskId string `protobuf:"bytes,2,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// Task service account email address.
	// This is the service account used for all authenticated requests by the
	// build.
	TaskServiceAccount string `protobuf:"bytes,3,opt,name=task_service_account,json=taskServiceAccount,proto3" json:"task_service_account,omitempty"`
	// Priority of the task. The lower the more important.
	// Valid values are [1..255].
	Priority int32 `protobuf:"varint,4,opt,name=priority,proto3" json:"priority,omitempty"`
	// Swarming dimensions for the task.
	TaskDimensions []*StringPair `protobuf:"bytes,5,rep,name=task_dimensions,json=taskDimensions,proto3" json:"task_dimensions,omitempty"`
	// Swarming dimensions of the bot used for the task.
	BotDimensions        []*StringPair `protobuf:"bytes,6,rep,name=bot_dimensions,json=botDimensions,proto3" json:"bot_dimensions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *BuildInfra_Swarming) Reset()         { *m = BuildInfra_Swarming{} }
func (m *BuildInfra_Swarming) String() string { return proto.CompactTextString(m) }
func (*BuildInfra_Swarming) ProtoMessage()    {}
func (*BuildInfra_Swarming) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3b43e011b474f4, []int{3, 1}
}

func (m *BuildInfra_Swarming) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildInfra_Swarming.Unmarshal(m, b)
}
func (m *BuildInfra_Swarming) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildInfra_Swarming.Marshal(b, m, deterministic)
}
func (m *BuildInfra_Swarming) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildInfra_Swarming.Merge(m, src)
}
func (m *BuildInfra_Swarming) XXX_Size() int {
	return xxx_messageInfo_BuildInfra_Swarming.Size(m)
}
func (m *BuildInfra_Swarming) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildInfra_Swarming.DiscardUnknown(m)
}

var xxx_messageInfo_BuildInfra_Swarming proto.InternalMessageInfo

func (m *BuildInfra_Swarming) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *BuildInfra_Swarming) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *BuildInfra_Swarming) GetTaskServiceAccount() string {
	if m != nil {
		return m.TaskServiceAccount
	}
	return ""
}

func (m *BuildInfra_Swarming) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *BuildInfra_Swarming) GetTaskDimensions() []*StringPair {
	if m != nil {
		return m.TaskDimensions
	}
	return nil
}

func (m *BuildInfra_Swarming) GetBotDimensions() []*StringPair {
	if m != nil {
		return m.BotDimensions
	}
	return nil
}

// LogDog-specific information.
type BuildInfra_LogDog struct {
	// LogDog hostname, e.g. "logs.chromium.org".
	Hostname string `protobuf:"bytes,1,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// LogDog project, e.g. "chromium".
	// Typically matches Build.builder.project.
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// A slash-separated path prefix shared by all logs and artifacts of this
	// build.
	// No other build can have the same prefix.
	// Can be used to discover logs and/or load log contents.
	//
	// Can be used to construct URL of a page that displays stdout/stderr of all
	// steps of a build. In pseudo-JS:
	//   q_stdout = `${project}/${prefix}/+/**/stdout`;
	//   q_stderr = `${project}/${prefix}/+/**/stderr`;
	//   url = `https://${host}/v/?s=${urlquote(q_stdout)}&s=${urlquote(q_stderr)}`;
	Prefix               string   `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildInfra_LogDog) Reset()         { *m = BuildInfra_LogDog{} }
func (m *BuildInfra_LogDog) String() string { return proto.CompactTextString(m) }
func (*BuildInfra_LogDog) ProtoMessage()    {}
func (*BuildInfra_LogDog) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3b43e011b474f4, []int{3, 2}
}

func (m *BuildInfra_LogDog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildInfra_LogDog.Unmarshal(m, b)
}
func (m *BuildInfra_LogDog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildInfra_LogDog.Marshal(b, m, deterministic)
}
func (m *BuildInfra_LogDog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildInfra_LogDog.Merge(m, src)
}
func (m *BuildInfra_LogDog) XXX_Size() int {
	return xxx_messageInfo_BuildInfra_LogDog.Size(m)
}
func (m *BuildInfra_LogDog) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildInfra_LogDog.DiscardUnknown(m)
}

var xxx_messageInfo_BuildInfra_LogDog proto.InternalMessageInfo

func (m *BuildInfra_LogDog) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *BuildInfra_LogDog) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *BuildInfra_LogDog) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

// Recipe-specific information.
type BuildInfra_Recipe struct {
	// CIPD package name containing the recipe used to run this build.
	CipdPackage string `protobuf:"bytes,1,opt,name=cipd_package,json=cipdPackage,proto3" json:"cipd_package,omitempty"`
	// Name of the recipe used to run this build.
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildInfra_Recipe) Reset()         { *m = BuildInfra_Recipe{} }
func (m *BuildInfra_Recipe) String() string { return proto.CompactTextString(m) }
func (*BuildInfra_Recipe) ProtoMessage()    {}
func (*BuildInfra_Recipe) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3b43e011b474f4, []int{3, 3}
}

func (m *BuildInfra_Recipe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildInfra_Recipe.Unmarshal(m, b)
}
func (m *BuildInfra_Recipe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildInfra_Recipe.Marshal(b, m, deterministic)
}
func (m *BuildInfra_Recipe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildInfra_Recipe.Merge(m, src)
}
func (m *BuildInfra_Recipe) XXX_Size() int {
	return xxx_messageInfo_BuildInfra_Recipe.Size(m)
}
func (m *BuildInfra_Recipe) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildInfra_Recipe.DiscardUnknown(m)
}

var xxx_messageInfo_BuildInfra_Recipe proto.InternalMessageInfo

func (m *BuildInfra_Recipe) GetCipdPackage() string {
	if m != nil {
		return m.CipdPackage
	}
	return ""
}

func (m *BuildInfra_Recipe) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Identifies a builder.
// Canonical string representation: “{project}/{bucket}/{builder}”.
type BuilderID struct {
	// Project ID, e.g. "chromium". Unique within a LUCI deployment.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Bucket name, e.g. "try". Unique within the project.
	// Together with project, defines an ACL.
	Bucket string `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// Builder name, e.g. "linux-rel". Unique within the bucket.
	Builder              string   `protobuf:"bytes,3,opt,name=builder,proto3" json:"builder,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuilderID) Reset()         { *m = BuilderID{} }
func (m *BuilderID) String() string { return proto.CompactTextString(m) }
func (*BuilderID) ProtoMessage()    {}
func (*BuilderID) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc3b43e011b474f4, []int{4}
}

func (m *BuilderID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuilderID.Unmarshal(m, b)
}
func (m *BuilderID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuilderID.Marshal(b, m, deterministic)
}
func (m *BuilderID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuilderID.Merge(m, src)
}
func (m *BuilderID) XXX_Size() int {
	return xxx_messageInfo_BuilderID.Size(m)
}
func (m *BuilderID) XXX_DiscardUnknown() {
	xxx_messageInfo_BuilderID.DiscardUnknown(m)
}

var xxx_messageInfo_BuilderID proto.InternalMessageInfo

func (m *BuilderID) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *BuilderID) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *BuilderID) GetBuilder() string {
	if m != nil {
		return m.Builder
	}
	return ""
}

func init() {
	proto.RegisterType((*Build)(nil), "buildbucket.v2.Build")
	proto.RegisterType((*Build_Input)(nil), "buildbucket.v2.Build.Input")
	proto.RegisterType((*Build_Output)(nil), "buildbucket.v2.Build.Output")
	proto.RegisterType((*CancelReason)(nil), "buildbucket.v2.CancelReason")
	proto.RegisterType((*InfraFailureReason)(nil), "buildbucket.v2.InfraFailureReason")
	proto.RegisterType((*BuildInfra)(nil), "buildbucket.v2.BuildInfra")
	proto.RegisterType((*BuildInfra_Buildbucket)(nil), "buildbucket.v2.BuildInfra.Buildbucket")
	proto.RegisterType((*BuildInfra_Swarming)(nil), "buildbucket.v2.BuildInfra.Swarming")
	proto.RegisterType((*BuildInfra_LogDog)(nil), "buildbucket.v2.BuildInfra.LogDog")
	proto.RegisterType((*BuildInfra_Recipe)(nil), "buildbucket.v2.BuildInfra.Recipe")
	proto.RegisterType((*BuilderID)(nil), "buildbucket.v2.BuilderID")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/buildbucket/proto/build.proto", fileDescriptor_bc3b43e011b474f4)
}

var fileDescriptor_bc3b43e011b474f4 = []byte{
	// 1043 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x61, 0x6f, 0xdb, 0x36,
	0x13, 0x7e, 0x9d, 0xc4, 0xb2, 0x7d, 0x8e, 0x9d, 0xbc, 0x6c, 0xd6, 0x68, 0x5a, 0x8b, 0xba, 0x1e,
	0x30, 0x78, 0xfb, 0xa0, 0xb4, 0x4e, 0xd7, 0xa1, 0xc8, 0x87, 0xa2, 0x76, 0xb6, 0x36, 0xc3, 0x86,
	0x15, 0xcc, 0xd0, 0x01, 0x03, 0x06, 0x81, 0xa6, 0x68, 0x85, 0x8b, 0x25, 0x0a, 0x24, 0x95, 0x26,
	0x3f, 0x64, 0x1f, 0xf7, 0x27, 0xf6, 0xc7, 0x86, 0xfd, 0x83, 0x41, 0x24, 0xe5, 0xa8, 0x71, 0x96,
	0x64, 0xfb, 0xa6, 0xbb, 0x7b, 0x9e, 0xbb, 0xe3, 0xc3, 0x3b, 0x49, 0x30, 0x4e, 0x44, 0x48, 0x4f,
	0xa4, 0x48, 0x79, 0x91, 0x86, 0x42, 0x26, 0x7b, 0x8b, 0x82, 0xf2, 0xbd, 0x59, 0xc1, 0x17, 0xf1,
	0xac, 0xa0, 0xa7, 0x4c, 0xef, 0xe5, 0x52, 0x68, 0x61, 0x3d, 0xa1, 0x79, 0x46, 0xfd, 0x5a, 0x38,
	0x3c, 0x1b, 0x07, 0x8f, 0x12, 0x21, 0x92, 0x05, 0xb3, 0xc8, 0x59, 0x31, 0xdf, 0xd3, 0x3c, 0x65,
	0x4a, 0x93, 0x34, 0xb7, 0x84, 0xe0, 0xc1, 0x55, 0x80, 0xd2, 0xb2, 0xa0, 0xda, 0x45, 0xf7, 0xef,
	0xd8, 0x02, 0x15, 0x69, 0x2a, 0x32, 0x47, 0x7a, 0x7a, 0x47, 0x92, 0xd2, 0xcc, 0x75, 0x31, 0xfc,
	0xab, 0x03, 0xcd, 0x49, 0x09, 0x40, 0x7d, 0x58, 0xe3, 0xb1, 0xdf, 0x18, 0x34, 0x46, 0xeb, 0x78,
	0x8d, 0xc7, 0x68, 0x1f, 0x5a, 0x86, 0xc9, 0xa4, 0xbf, 0x36, 0x68, 0x8c, 0xba, 0xe3, 0x8f, 0xc3,
	0x0f, 0x8f, 0x18, 0x4e, 0x6c, 0xf8, 0xe8, 0x10, 0x57, 0x48, 0x74, 0x1f, 0xbc, 0xac, 0x48, 0x67,
	0x4c, 0xfa, 0xeb, 0x83, 0xc6, 0xa8, 0x89, 0x9d, 0x85, 0x1e, 0x02, 0x50, 0xc9, 0x88, 0x66, 0x71,
	0x34, 0xbb, 0xf0, 0x37, 0x06, 0x8d, 0x51, 0x07, 0x77, 0x9c, 0x67, 0x72, 0x81, 0x0e, 0xa0, 0x6b,
	0x8d, 0xa8, 0x54, 0xc9, 0xf7, 0x4c, 0xbd, 0x20, 0xb4, 0x0a, 0x85, 0x95, 0x42, 0xe1, 0x8f, 0x95,
	0x84, 0xd8, 0x65, 0x2b, 0x1d, 0xe8, 0x05, 0x80, 0xd2, 0x44, 0x6a, 0xcb, 0x6d, 0xdd, 0xca, 0xed,
	0x18, 0xb4, 0xa1, 0x7e, 0x09, 0x6d, 0x96, 0xc5, 0x96, 0xd8, 0xbe, 0x95, 0xd8, 0x62, 0x59, 0x6c,
	0x68, 0x07, 0xd0, 0x2d, 0xf2, 0x78, 0xd9, 0x6e, 0xe7, 0xf6, 0x76, 0x2d, 0xdc, 0x90, 0x43, 0xf0,
	0x94, 0x26, 0xba, 0x50, 0xfe, 0xe6, 0xa0, 0x31, 0xea, 0x8f, 0xef, 0x5f, 0x95, 0xf5, 0xd8, 0x44,
	0xb1, 0x43, 0xa1, 0x77, 0xb0, 0xc3, 0xb3, 0xb9, 0x24, 0xd1, 0x9c, 0xf0, 0x45, 0x21, 0x59, 0x24,
	0x19, 0x51, 0x22, 0xf3, 0x7b, 0xa6, 0xea, 0xf0, 0x2a, 0xfb, 0xa8, 0xc4, 0x7e, 0x63, 0xa1, 0xd8,
	0x20, 0xdf, 0xfc, 0x0f, 0x23, 0xbe, 0xe2, 0x45, 0x53, 0xe8, 0x51, 0x92, 0x51, 0xb6, 0xa8, 0x12,
	0xf6, 0x4d, 0xc2, 0x07, 0x57, 0x13, 0x4e, 0x0d, 0x68, 0x99, 0x6a, 0x93, 0xd6, 0x6c, 0xf4, 0x14,
	0x9a, 0x3c, 0xcb, 0x0b, 0xed, 0x6f, 0x19, 0xf2, 0x27, 0xd7, 0x8e, 0x48, 0x78, 0x54, 0x42, 0xb0,
	0x45, 0xa2, 0x67, 0xe0, 0x89, 0x42, 0x97, 0x9c, 0xed, 0xeb, 0x0b, 0x5a, 0xce, 0x0f, 0x06, 0x83,
	0x1d, 0x16, 0x7d, 0x01, 0xcd, 0x72, 0x6a, 0x95, 0xff, 0xff, 0xc1, 0xfa, 0xa8, 0x3b, 0xde, 0x59,
	0x15, 0x8d, 0xe5, 0xd8, 0x42, 0xd0, 0x93, 0xb2, 0xa9, 0xb9, 0x24, 0x3e, 0x72, 0x17, 0x73, 0x5d,
	0x01, 0xa3, 0x13, 0xb6, 0x40, 0x14, 0xc2, 0x86, 0x26, 0x89, 0xf2, 0xef, 0x99, 0xe4, 0xc1, 0x6a,
	0x72, 0xc9, 0xb3, 0xe4, 0x2d, 0xe1, 0x12, 0x1b, 0x5c, 0xf0, 0x67, 0x03, 0x9a, 0xe6, 0x50, 0xe8,
	0x2b, 0x80, 0x5c, 0x8a, 0x9c, 0x49, 0xcd, 0x99, 0x32, 0xdb, 0xd3, 0x1d, 0xef, 0xae, 0x4c, 0xc2,
	0xb1, 0x59, 0x6d, 0x5c, 0x83, 0xa2, 0x43, 0xe8, 0x27, 0x5c, 0xf3, 0x05, 0x53, 0x51, 0xb9, 0xc3,
	0x5c, 0xbb, 0x2d, 0x7b, 0x78, 0xb5, 0xf8, 0x6b, 0x8b, 0x9a, 0x1a, 0x10, 0xee, 0x25, 0x75, 0x13,
	0x4d, 0xa1, 0x9f, 0x30, 0x29, 0xb9, 0x8e, 0xe8, 0x09, 0xc9, 0x12, 0xa6, 0xfc, 0x75, 0x73, 0x84,
	0x15, 0x51, 0x5f, 0x1b, 0xd4, 0xd4, 0x80, 0x70, 0x2f, 0xa9, 0x59, 0x0a, 0x0d, 0x61, 0x93, 0x9d,
	0xe7, 0x4c, 0xf2, 0x94, 0x65, 0x9a, 0x2c, 0xfc, 0xe6, 0xa0, 0x31, 0x6a, 0xe3, 0x0f, 0x7c, 0xc1,
	0x1f, 0x0d, 0xf0, 0xec, 0x95, 0xfc, 0xf7, 0x23, 0x7f, 0x0e, 0xdb, 0xaa, 0x48, 0x53, 0x22, 0x2f,
	0xa2, 0x94, 0xc8, 0xd3, 0x58, 0xbc, 0xcf, 0xcc, 0xa1, 0x3b, 0x78, 0xcb, 0xf9, 0xbf, 0x77, 0xee,
	0x6b, 0xd4, 0x59, 0xff, 0xf7, 0xea, 0x4c, 0xb6, 0xa0, 0x67, 0x97, 0xc8, 0x8d, 0xf8, 0xb7, 0x1b,
	0xed, 0xe6, 0xb6, 0x37, 0x3c, 0x82, 0xcd, 0xfa, 0x50, 0x23, 0x1f, 0x5a, 0x29, 0x53, 0x8a, 0x24,
	0xcc, 0x9c, 0xa6, 0x83, 0x2b, 0x13, 0x3d, 0x82, 0xae, 0x1d, 0x77, 0xfb, 0xde, 0xb2, 0xcd, 0x42,
	0xe5, 0x9a, 0x5c, 0x0c, 0x23, 0x40, 0xab, 0x0b, 0x77, 0x43, 0xc2, 0x3d, 0xb8, 0x27, 0x99, 0x12,
	0x85, 0xa4, 0x2c, 0x62, 0xe7, 0x27, 0xa4, 0x50, 0x9a, 0x0b, 0xab, 0x42, 0x1b, 0xa3, 0x2a, 0xf4,
	0xf5, 0x32, 0x32, 0xfc, 0xdd, 0x03, 0xb8, 0x9c, 0x57, 0xf4, 0x06, 0xba, 0x35, 0x01, 0x9c, 0xf8,
	0x9f, 0xfd, 0xf3, 0x80, 0xdb, 0x47, 0x1b, 0xc1, 0x75, 0x2a, 0x7a, 0x09, 0x6d, 0xf5, 0x9e, 0xc8,
	0x94, 0x67, 0x89, 0x9b, 0xbc, 0x4f, 0x6f, 0x48, 0x73, 0xec, 0xa0, 0x78, 0x49, 0x42, 0x2f, 0xc0,
	0x5b, 0x88, 0x24, 0x16, 0x89, 0xbb, 0x9a, 0xc7, 0x37, 0xd0, 0xbf, 0x13, 0xc9, 0xa1, 0x48, 0xb0,
	0x23, 0x94, 0x54, 0xc9, 0x28, 0xcf, 0x99, 0xf9, 0x12, 0xdc, 0x4c, 0xc5, 0x06, 0x88, 0x1d, 0x21,
	0xf8, 0x05, 0xba, 0xb5, 0x23, 0xa1, 0xe7, 0xb0, 0xab, 0x98, 0x3c, 0xe3, 0x94, 0x45, 0x54, 0x64,
	0x73, 0x9e, 0x44, 0x92, 0x9d, 0x71, 0x55, 0x69, 0xda, 0xc1, 0x1f, 0xb9, 0xf0, 0xd4, 0x44, 0xb1,
	0x0b, 0x96, 0xdf, 0x29, 0x4a, 0x32, 0x22, 0xed, 0xb7, 0xa8, 0x8d, 0x9d, 0x15, 0xfc, 0xb6, 0x06,
	0xed, 0xea, 0xac, 0x28, 0x80, 0xf6, 0x89, 0x50, 0x3a, 0x23, 0x69, 0x75, 0x8f, 0x4b, 0x1b, 0xed,
	0x42, 0x4b, 0x13, 0x75, 0x1a, 0xf1, 0xd8, 0x15, 0xf2, 0x4a, 0xf3, 0x28, 0x46, 0x4f, 0x60, 0xc7,
	0x04, 0xaa, 0xb6, 0x08, 0xa5, 0xa2, 0xc8, 0xec, 0xfc, 0x76, 0x30, 0x2a, 0x63, 0xc7, 0x36, 0xf4,
	0xca, 0x46, 0xca, 0x32, 0xb9, 0xe4, 0x42, 0x72, 0x6d, 0xbb, 0x69, 0xe2, 0xa5, 0x8d, 0xa6, 0xb0,
	0x65, 0xb2, 0xc5, 0xe5, 0x1e, 0x96, 0x9d, 0x2b, 0xbf, 0x79, 0xeb, 0x3b, 0xaa, 0x5f, 0x52, 0x0e,
	0x97, 0x0c, 0xf4, 0x0a, 0xfa, 0x33, 0xa1, 0xeb, 0x39, 0xbc, 0x5b, 0x73, 0xf4, 0x66, 0x42, 0x5f,
	0xa6, 0x08, 0xde, 0x81, 0x67, 0xef, 0xf0, 0x46, 0x51, 0x7c, 0x68, 0xe5, 0x52, 0xfc, 0xca, 0xa8,
	0x76, 0xa2, 0x54, 0x66, 0xa9, 0x77, 0x2e, 0xd9, 0x9c, 0x9f, 0x3b, 0x1d, 0x9c, 0x15, 0xbc, 0x04,
	0xcf, 0x5e, 0x30, 0x7a, 0x0c, 0x9b, 0x94, 0xe7, 0x71, 0x94, 0x13, 0x7a, 0x7a, 0xb9, 0x38, 0xdd,
	0xd2, 0xf7, 0xd6, 0xba, 0x10, 0x82, 0x0d, 0x53, 0xd6, 0xe6, 0x36, 0xcf, 0xc3, 0x9f, 0xa0, 0xb3,
	0xfc, 0x0d, 0xa9, 0xd7, 0x6f, 0xac, 0xd4, 0x77, 0x2b, 0xe3, 0x6e, 0xcb, 0xcd, 0x8f, 0x7f, 0xf9,
	0x93, 0x63, 0x1b, 0xab, 0xcc, 0xc9, 0xf3, 0x9f, 0x9f, 0xdd, 0xed, 0x6f, 0xea, 0xa0, 0xe6, 0xc9,
	0x67, 0x33, 0xcf, 0x38, 0xf7, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x55, 0x80, 0x6d, 0xae, 0x44,
	0x0a, 0x00, 0x00,
}
