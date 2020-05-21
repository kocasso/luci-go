// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/buildbucket/proto/common.proto

package buildbucketpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// Status of a build or a step.
type Status int32

const (
	// Unspecified state. Meaning depends on the context.
	Status_STATUS_UNSPECIFIED Status = 0
	// Build was scheduled, but did not start or end yet.
	Status_SCHEDULED Status = 1
	// Build/step has started.
	Status_STARTED Status = 2
	// A union of all terminal statuses.
	// Can be used in BuildPredicate.status.
	// A concrete build/step cannot have this status.
	// Can be used as a bitmask to check that a build/step ended.
	Status_ENDED_MASK Status = 4
	// A build/step ended successfully.
	// This is a terminal status. It may not transition to another status.
	Status_SUCCESS Status = 12
	// A build/step ended unsuccessfully due to its Build.Input,
	// e.g. tests failed, and NOT due to a build infrastructure failure.
	// This is a terminal status. It may not transition to another status.
	Status_FAILURE Status = 20
	// A build/step ended unsuccessfully due to a failure independent of the
	// input, e.g. swarming failed, not enough capacity or the recipe was unable
	// to read the patch from gerrit.
	// start_time is not required for this status.
	// This is a terminal status. It may not transition to another status.
	Status_INFRA_FAILURE Status = 36
	// A build was cancelled explicitly, e.g. via an RPC.
	// This is a terminal status. It may not transition to another status.
	Status_CANCELED Status = 68
)

var Status_name = map[int32]string{
	0:  "STATUS_UNSPECIFIED",
	1:  "SCHEDULED",
	2:  "STARTED",
	4:  "ENDED_MASK",
	12: "SUCCESS",
	20: "FAILURE",
	36: "INFRA_FAILURE",
	68: "CANCELED",
}

var Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"SCHEDULED":          1,
	"STARTED":            2,
	"ENDED_MASK":         4,
	"SUCCESS":            12,
	"FAILURE":            20,
	"INFRA_FAILURE":      36,
	"CANCELED":           68,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{0}
}

// A boolean with an undefined value.
type Trinary int32

const (
	Trinary_UNSET Trinary = 0
	Trinary_YES   Trinary = 1
	Trinary_NO    Trinary = 2
)

var Trinary_name = map[int32]string{
	0: "UNSET",
	1: "YES",
	2: "NO",
}

var Trinary_value = map[string]int32{
	"UNSET": 0,
	"YES":   1,
	"NO":    2,
}

func (x Trinary) String() string {
	return proto.EnumName(Trinary_name, int32(x))
}

func (Trinary) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{1}
}

// An executable to run when the build is ready to start.
//
// Please refer to go.chromium.org/luci/luciexe for the protocol this executable
// is expected to implement.
//
// In addition to the "Host Application" responsibilities listed there,
// buildbucket will also ensure that $CWD points to an empty directory when it
// starts the build.
type Executable struct {
	// The CIPD package containing the executable.
	//
	// See the `cmd` field below for how the executable will be located within the
	// package.
	CipdPackage string `protobuf:"bytes,1,opt,name=cipd_package,json=cipdPackage,proto3" json:"cipd_package,omitempty"`
	// The CIPD version to fetch.
	//
	// Optional. If omitted, this defaults to `latest`.
	CipdVersion string `protobuf:"bytes,2,opt,name=cipd_version,json=cipdVersion,proto3" json:"cipd_version,omitempty"`
	// The command to invoke within the package.
	//
	// The 0th argument is taken as relative to the cipd_package root (a.k.a.
	// BBAgentArgs.payload_path), so "foo" would invoke the binary called "foo" in
	// the root of the package. On Windows, this will automatically look
	// first for ".exe" and ".bat" variants. Similarly, "subdir/foo" would
	// look for "foo" in "subdir" of the CIPD package.
	//
	// The other arguments are passed verbatim to the executable.
	//
	// The 'build.proto' binary message will always be passed to stdin, even when
	// this command has arguments (see go.chromium.org/luci/luciexe).
	//
	// RECOMMENDATION: It's advised to rely on the build.proto's Input.Properties
	// field for passing task-specific data. Properties are JSON-typed and can be
	// modeled with a protobuf (using JSONPB). However, supplying additional args
	// can be useful to, e.g., increase logging verbosity, or similar
	// 'system level' settings within the binary.
	//
	// Optional. If omitted, defaults to `['luciexe']`.
	Cmd                  []string `protobuf:"bytes,3,rep,name=cmd,proto3" json:"cmd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Executable) Reset()         { *m = Executable{} }
func (m *Executable) String() string { return proto.CompactTextString(m) }
func (*Executable) ProtoMessage()    {}
func (*Executable) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{0}
}

func (m *Executable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Executable.Unmarshal(m, b)
}
func (m *Executable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Executable.Marshal(b, m, deterministic)
}
func (m *Executable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Executable.Merge(m, src)
}
func (m *Executable) XXX_Size() int {
	return xxx_messageInfo_Executable.Size(m)
}
func (m *Executable) XXX_DiscardUnknown() {
	xxx_messageInfo_Executable.DiscardUnknown(m)
}

var xxx_messageInfo_Executable proto.InternalMessageInfo

func (m *Executable) GetCipdPackage() string {
	if m != nil {
		return m.CipdPackage
	}
	return ""
}

func (m *Executable) GetCipdVersion() string {
	if m != nil {
		return m.CipdVersion
	}
	return ""
}

func (m *Executable) GetCmd() []string {
	if m != nil {
		return m.Cmd
	}
	return nil
}

// Machine-readable details of a status.
// Human-readble details are present in a sibling summary_markdown field.
type StatusDetails struct {
	// If set, indicates that the failure was due to a resource exhaustion / quota
	// denial.
	// Applicable in FAILURE and INFRA_FAILURE statuses.
	ResourceExhaustion *StatusDetails_ResourceExhaustion `protobuf:"bytes,3,opt,name=resource_exhaustion,json=resourceExhaustion,proto3" json:"resource_exhaustion,omitempty"`
	// If set, indicates that the failure was due to a timeout.
	// Applicable in FAILURE and INFRA_FAILURE statuses.
	Timeout              *StatusDetails_Timeout `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *StatusDetails) Reset()         { *m = StatusDetails{} }
func (m *StatusDetails) String() string { return proto.CompactTextString(m) }
func (*StatusDetails) ProtoMessage()    {}
func (*StatusDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{1}
}

func (m *StatusDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusDetails.Unmarshal(m, b)
}
func (m *StatusDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusDetails.Marshal(b, m, deterministic)
}
func (m *StatusDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusDetails.Merge(m, src)
}
func (m *StatusDetails) XXX_Size() int {
	return xxx_messageInfo_StatusDetails.Size(m)
}
func (m *StatusDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusDetails.DiscardUnknown(m)
}

var xxx_messageInfo_StatusDetails proto.InternalMessageInfo

func (m *StatusDetails) GetResourceExhaustion() *StatusDetails_ResourceExhaustion {
	if m != nil {
		return m.ResourceExhaustion
	}
	return nil
}

func (m *StatusDetails) GetTimeout() *StatusDetails_Timeout {
	if m != nil {
		return m.Timeout
	}
	return nil
}

type StatusDetails_ResourceExhaustion struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusDetails_ResourceExhaustion) Reset()         { *m = StatusDetails_ResourceExhaustion{} }
func (m *StatusDetails_ResourceExhaustion) String() string { return proto.CompactTextString(m) }
func (*StatusDetails_ResourceExhaustion) ProtoMessage()    {}
func (*StatusDetails_ResourceExhaustion) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{1, 0}
}

func (m *StatusDetails_ResourceExhaustion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusDetails_ResourceExhaustion.Unmarshal(m, b)
}
func (m *StatusDetails_ResourceExhaustion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusDetails_ResourceExhaustion.Marshal(b, m, deterministic)
}
func (m *StatusDetails_ResourceExhaustion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusDetails_ResourceExhaustion.Merge(m, src)
}
func (m *StatusDetails_ResourceExhaustion) XXX_Size() int {
	return xxx_messageInfo_StatusDetails_ResourceExhaustion.Size(m)
}
func (m *StatusDetails_ResourceExhaustion) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusDetails_ResourceExhaustion.DiscardUnknown(m)
}

var xxx_messageInfo_StatusDetails_ResourceExhaustion proto.InternalMessageInfo

type StatusDetails_Timeout struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusDetails_Timeout) Reset()         { *m = StatusDetails_Timeout{} }
func (m *StatusDetails_Timeout) String() string { return proto.CompactTextString(m) }
func (*StatusDetails_Timeout) ProtoMessage()    {}
func (*StatusDetails_Timeout) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{1, 1}
}

func (m *StatusDetails_Timeout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusDetails_Timeout.Unmarshal(m, b)
}
func (m *StatusDetails_Timeout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusDetails_Timeout.Marshal(b, m, deterministic)
}
func (m *StatusDetails_Timeout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusDetails_Timeout.Merge(m, src)
}
func (m *StatusDetails_Timeout) XXX_Size() int {
	return xxx_messageInfo_StatusDetails_Timeout.Size(m)
}
func (m *StatusDetails_Timeout) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusDetails_Timeout.DiscardUnknown(m)
}

var xxx_messageInfo_StatusDetails_Timeout proto.InternalMessageInfo

// A named log of a step or build.
type Log struct {
	// Log name, standard ("stdout", "stderr") or custom (e.g. "json.output").
	// Unique within the containing message (step or build).
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// URL of a Human-readable page that displays log contents.
	ViewUrl string `protobuf:"bytes,2,opt,name=view_url,json=viewUrl,proto3" json:"view_url,omitempty"`
	// URL of the log content.
	// As of 2018-09-06, the only supported scheme is "logdog".
	// Typically it has form
	// "logdog://<host>/<project>/<prefix>/+/<stream_name>".
	// See also
	// https://godoc.org/go.chromium.org/luci/logdog/common/types#ParseURL
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Log) Reset()         { *m = Log{} }
func (m *Log) String() string { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()    {}
func (*Log) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{2}
}

func (m *Log) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Log.Unmarshal(m, b)
}
func (m *Log) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Log.Marshal(b, m, deterministic)
}
func (m *Log) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Log.Merge(m, src)
}
func (m *Log) XXX_Size() int {
	return xxx_messageInfo_Log.Size(m)
}
func (m *Log) XXX_DiscardUnknown() {
	xxx_messageInfo_Log.DiscardUnknown(m)
}

var xxx_messageInfo_Log proto.InternalMessageInfo

func (m *Log) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Log) GetViewUrl() string {
	if m != nil {
		return m.ViewUrl
	}
	return ""
}

func (m *Log) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// A Gerrit patchset.
type GerritChange struct {
	// Gerrit hostname, e.g. "chromium-review.googlesource.com".
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Gerrit project, e.g. "chromium/src".
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Change number, e.g. 12345.
	Change int64 `protobuf:"varint,3,opt,name=change,proto3" json:"change,omitempty"`
	// Patch set number, e.g. 1.
	Patchset             int64    `protobuf:"varint,4,opt,name=patchset,proto3" json:"patchset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GerritChange) Reset()         { *m = GerritChange{} }
func (m *GerritChange) String() string { return proto.CompactTextString(m) }
func (*GerritChange) ProtoMessage()    {}
func (*GerritChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{3}
}

func (m *GerritChange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GerritChange.Unmarshal(m, b)
}
func (m *GerritChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GerritChange.Marshal(b, m, deterministic)
}
func (m *GerritChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GerritChange.Merge(m, src)
}
func (m *GerritChange) XXX_Size() int {
	return xxx_messageInfo_GerritChange.Size(m)
}
func (m *GerritChange) XXX_DiscardUnknown() {
	xxx_messageInfo_GerritChange.DiscardUnknown(m)
}

var xxx_messageInfo_GerritChange proto.InternalMessageInfo

func (m *GerritChange) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *GerritChange) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *GerritChange) GetChange() int64 {
	if m != nil {
		return m.Change
	}
	return 0
}

func (m *GerritChange) GetPatchset() int64 {
	if m != nil {
		return m.Patchset
	}
	return 0
}

// A landed Git commit hosted on Gitiles.
type GitilesCommit struct {
	// Gitiles hostname, e.g. "chromium.googlesource.com".
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// Repository name on the host, e.g. "chromium/src".
	Project string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	// Commit HEX SHA1.
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// Commit ref, e.g. "refs/heads/master".
	// NOT a branch name: if specified, must start with "refs/".
	Ref string `protobuf:"bytes,4,opt,name=ref,proto3" json:"ref,omitempty"`
	// Defines a total order of commits on the ref. Requires ref field.
	// Typically 1-based, monotonically increasing, contiguous integer
	// defined by a Gerrit plugin, goto.google.com/git-numberer.
	// TODO(tandrii): make it a public doc.
	Position             uint32   `protobuf:"varint,5,opt,name=position,proto3" json:"position,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GitilesCommit) Reset()         { *m = GitilesCommit{} }
func (m *GitilesCommit) String() string { return proto.CompactTextString(m) }
func (*GitilesCommit) ProtoMessage()    {}
func (*GitilesCommit) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{4}
}

func (m *GitilesCommit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GitilesCommit.Unmarshal(m, b)
}
func (m *GitilesCommit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GitilesCommit.Marshal(b, m, deterministic)
}
func (m *GitilesCommit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GitilesCommit.Merge(m, src)
}
func (m *GitilesCommit) XXX_Size() int {
	return xxx_messageInfo_GitilesCommit.Size(m)
}
func (m *GitilesCommit) XXX_DiscardUnknown() {
	xxx_messageInfo_GitilesCommit.DiscardUnknown(m)
}

var xxx_messageInfo_GitilesCommit proto.InternalMessageInfo

func (m *GitilesCommit) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *GitilesCommit) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *GitilesCommit) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GitilesCommit) GetRef() string {
	if m != nil {
		return m.Ref
	}
	return ""
}

func (m *GitilesCommit) GetPosition() uint32 {
	if m != nil {
		return m.Position
	}
	return 0
}

// A key-value pair of strings.
type StringPair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringPair) Reset()         { *m = StringPair{} }
func (m *StringPair) String() string { return proto.CompactTextString(m) }
func (*StringPair) ProtoMessage()    {}
func (*StringPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{5}
}

func (m *StringPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringPair.Unmarshal(m, b)
}
func (m *StringPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringPair.Marshal(b, m, deterministic)
}
func (m *StringPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringPair.Merge(m, src)
}
func (m *StringPair) XXX_Size() int {
	return xxx_messageInfo_StringPair.Size(m)
}
func (m *StringPair) XXX_DiscardUnknown() {
	xxx_messageInfo_StringPair.DiscardUnknown(m)
}

var xxx_messageInfo_StringPair proto.InternalMessageInfo

func (m *StringPair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StringPair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// Half-open time range.
type TimeRange struct {
	// Inclusive lower boundary. Optional.
	StartTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Exclusive upper boundary. Optional.
	EndTime              *timestamp.Timestamp `protobuf:"bytes,2,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TimeRange) Reset()         { *m = TimeRange{} }
func (m *TimeRange) String() string { return proto.CompactTextString(m) }
func (*TimeRange) ProtoMessage()    {}
func (*TimeRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{6}
}

func (m *TimeRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeRange.Unmarshal(m, b)
}
func (m *TimeRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeRange.Marshal(b, m, deterministic)
}
func (m *TimeRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeRange.Merge(m, src)
}
func (m *TimeRange) XXX_Size() int {
	return xxx_messageInfo_TimeRange.Size(m)
}
func (m *TimeRange) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeRange.DiscardUnknown(m)
}

var xxx_messageInfo_TimeRange proto.InternalMessageInfo

func (m *TimeRange) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TimeRange) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

// A requested dimension. Looks like StringPair, but also has an expiration.
type RequestedDimension struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// If set, ignore this dimension after this duration.
	Expiration           *duration.Duration `protobuf:"bytes,3,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RequestedDimension) Reset()         { *m = RequestedDimension{} }
func (m *RequestedDimension) String() string { return proto.CompactTextString(m) }
func (*RequestedDimension) ProtoMessage()    {}
func (*RequestedDimension) Descriptor() ([]byte, []int) {
	return fileDescriptor_a1a0c34bd7fcf0dc, []int{7}
}

func (m *RequestedDimension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestedDimension.Unmarshal(m, b)
}
func (m *RequestedDimension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestedDimension.Marshal(b, m, deterministic)
}
func (m *RequestedDimension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestedDimension.Merge(m, src)
}
func (m *RequestedDimension) XXX_Size() int {
	return xxx_messageInfo_RequestedDimension.Size(m)
}
func (m *RequestedDimension) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestedDimension.DiscardUnknown(m)
}

var xxx_messageInfo_RequestedDimension proto.InternalMessageInfo

func (m *RequestedDimension) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RequestedDimension) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *RequestedDimension) GetExpiration() *duration.Duration {
	if m != nil {
		return m.Expiration
	}
	return nil
}

func init() {
	proto.RegisterEnum("buildbucket.v2.Status", Status_name, Status_value)
	proto.RegisterEnum("buildbucket.v2.Trinary", Trinary_name, Trinary_value)
	proto.RegisterType((*Executable)(nil), "buildbucket.v2.Executable")
	proto.RegisterType((*StatusDetails)(nil), "buildbucket.v2.StatusDetails")
	proto.RegisterType((*StatusDetails_ResourceExhaustion)(nil), "buildbucket.v2.StatusDetails.ResourceExhaustion")
	proto.RegisterType((*StatusDetails_Timeout)(nil), "buildbucket.v2.StatusDetails.Timeout")
	proto.RegisterType((*Log)(nil), "buildbucket.v2.Log")
	proto.RegisterType((*GerritChange)(nil), "buildbucket.v2.GerritChange")
	proto.RegisterType((*GitilesCommit)(nil), "buildbucket.v2.GitilesCommit")
	proto.RegisterType((*StringPair)(nil), "buildbucket.v2.StringPair")
	proto.RegisterType((*TimeRange)(nil), "buildbucket.v2.TimeRange")
	proto.RegisterType((*RequestedDimension)(nil), "buildbucket.v2.RequestedDimension")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/buildbucket/proto/common.proto", fileDescriptor_a1a0c34bd7fcf0dc)
}

var fileDescriptor_a1a0c34bd7fcf0dc = []byte{
	// 714 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x51, 0x8f, 0xdb, 0x44,
	0x10, 0xae, 0xed, 0xdc, 0x39, 0x9e, 0x24, 0x27, 0xb3, 0x9c, 0xaa, 0x34, 0x0f, 0x10, 0x2c, 0x2a,
	0x45, 0x7d, 0x70, 0xd0, 0xb5, 0x20, 0x21, 0x1e, 0x50, 0xb0, 0x9d, 0x92, 0x72, 0x84, 0x93, 0x9d,
	0x20, 0xc1, 0x4b, 0xb4, 0xb1, 0xb7, 0xce, 0x72, 0xb6, 0xd7, 0xac, 0xd7, 0xe9, 0x55, 0x88, 0x67,
	0xfe, 0x26, 0x3f, 0x05, 0xed, 0xda, 0x0e, 0x57, 0x4e, 0x82, 0xf2, 0x36, 0xdf, 0xb7, 0xdf, 0x37,
	0xb3, 0xb3, 0x33, 0x36, 0x3c, 0x4f, 0x99, 0x1b, 0x1f, 0x38, 0xcb, 0x69, 0x9d, 0xbb, 0x8c, 0xa7,
	0xf3, 0xac, 0x8e, 0xe9, 0x7c, 0x5f, 0xd3, 0x2c, 0xd9, 0xd7, 0xf1, 0x2d, 0x11, 0xf3, 0x92, 0x33,
	0xc1, 0xe6, 0x31, 0xcb, 0x73, 0x56, 0xb8, 0x0a, 0xa0, 0x8b, 0x7b, 0xe7, 0xee, 0xf1, 0x6a, 0xf2,
	0x51, 0xca, 0x58, 0x9a, 0x91, 0x46, 0xba, 0xaf, 0x5f, 0xcf, 0x93, 0x9a, 0x63, 0x41, 0x3b, 0xfd,
	0xe4, 0xe3, 0x7f, 0x9e, 0x0b, 0x9a, 0x93, 0x4a, 0xe0, 0xbc, 0x6c, 0x04, 0x4e, 0x02, 0x10, 0xdc,
	0x91, 0xb8, 0x16, 0x78, 0x9f, 0x11, 0xf4, 0x09, 0x0c, 0x63, 0x5a, 0x26, 0xbb, 0x12, 0xc7, 0xb7,
	0x38, 0x25, 0x63, 0x6d, 0xaa, 0xcd, 0xac, 0x70, 0x20, 0xb9, 0x9b, 0x86, 0x3a, 0x49, 0x8e, 0x84,
	0x57, 0x94, 0x15, 0x63, 0xfd, 0x6f, 0xc9, 0x8f, 0x0d, 0x85, 0x6c, 0x30, 0xe2, 0x3c, 0x19, 0x1b,
	0x53, 0x63, 0x66, 0x85, 0x32, 0x74, 0xfe, 0xd4, 0x60, 0x14, 0x09, 0x2c, 0xea, 0xca, 0x27, 0x02,
	0xd3, 0xac, 0x42, 0x18, 0x3e, 0xe4, 0xa4, 0x62, 0x35, 0x8f, 0xc9, 0x8e, 0xdc, 0x1d, 0x70, 0x5d,
	0xc9, 0x5b, 0x8f, 0x8d, 0xa9, 0x36, 0x1b, 0x5c, 0x7d, 0xe6, 0xbe, 0xdb, 0xa6, 0xfb, 0x8e, 0xd7,
	0x0d, 0x5b, 0x63, 0x70, 0xf2, 0x85, 0x88, 0x3f, 0xe0, 0xd0, 0xd7, 0x60, 0xca, 0x6e, 0x59, 0x2d,
	0xc6, 0x3d, 0x95, 0xf6, 0xe9, 0xbf, 0xa7, 0xdd, 0x34, 0xe2, 0xb0, 0x73, 0x4d, 0x2e, 0x01, 0x3d,
	0x2c, 0x35, 0xb1, 0xc0, 0x6c, 0x95, 0xaf, 0x7a, 0x7d, 0xcd, 0xd6, 0x5f, 0xf5, 0xfa, 0xba, 0x6d,
	0x38, 0x4b, 0x30, 0xae, 0x59, 0x8a, 0x10, 0xf4, 0x0a, 0x9c, 0x77, 0x2f, 0xa7, 0x62, 0xf4, 0x04,
	0xfa, 0x47, 0x4a, 0xde, 0xec, 0x6a, 0x9e, 0xb5, 0xcf, 0x65, 0x4a, 0xbc, 0xe5, 0x99, 0x7c, 0x2a,
	0xc9, 0x1a, 0x8a, 0x95, 0xa1, 0x53, 0xc2, 0xf0, 0x25, 0xe1, 0x9c, 0x0a, 0xef, 0x80, 0x8b, 0x94,
	0xc8, 0x84, 0x07, 0x56, 0x89, 0x2e, 0xa1, 0x8c, 0xd1, 0x18, 0xcc, 0x92, 0xb3, 0x5f, 0x48, 0x2c,
	0xba, 0x7c, 0x2d, 0x44, 0x8f, 0xe1, 0x3c, 0x56, 0x3e, 0x95, 0xd2, 0x08, 0x5b, 0x84, 0x26, 0xd0,
	0x2f, 0xb1, 0x88, 0x0f, 0x15, 0x69, 0x1e, 0xc3, 0x08, 0x4f, 0xd8, 0xf9, 0x0d, 0x46, 0x2f, 0xa9,
	0xa0, 0x19, 0xa9, 0x3c, 0x96, 0xe7, 0x54, 0xfc, 0xcf, 0x92, 0x17, 0xa0, 0xd3, 0xa4, 0xed, 0x40,
	0xa7, 0x89, 0x6c, 0x89, 0x93, 0xd7, 0xaa, 0x8a, 0x15, 0xca, 0x50, 0x15, 0x67, 0x15, 0x55, 0x03,
	0x3e, 0x9b, 0x6a, 0xb3, 0x51, 0x78, 0xc2, 0xce, 0x0b, 0x80, 0x48, 0x70, 0x5a, 0xa4, 0x37, 0x98,
	0x72, 0xe9, 0xbd, 0x25, 0x6f, 0xdb, 0xc2, 0x32, 0x44, 0x97, 0x70, 0x76, 0xc4, 0x59, 0x4d, 0xda,
	0xaa, 0x0d, 0x70, 0x7e, 0x07, 0x4b, 0xce, 0x20, 0x54, 0xbd, 0x7d, 0x09, 0x50, 0x09, 0xcc, 0xc5,
	0x4e, 0xce, 0x4d, 0x79, 0x07, 0x57, 0x13, 0xb7, 0x59, 0x7c, 0xb7, 0x5b, 0x7c, 0x35, 0x5d, 0xb5,
	0xf8, 0xa1, 0xa5, 0xd4, 0x12, 0xa3, 0xcf, 0xa1, 0x4f, 0x8a, 0xa4, 0x31, 0xea, 0xff, 0x69, 0x34,
	0x49, 0x91, 0x48, 0xe4, 0xbc, 0x91, 0x8b, 0xf1, 0x6b, 0x4d, 0x2a, 0x41, 0x12, 0x9f, 0xe6, 0xa4,
	0xe8, 0xd6, 0xfe, 0x7d, 0x2e, 0x2f, 0xef, 0x4b, 0xee, 0x4a, 0xda, 0x7c, 0xa7, 0xed, 0xc6, 0x3f,
	0x79, 0x50, 0xd6, 0x6f, 0x3f, 0xe4, 0xf0, 0x9e, 0xf8, 0xd9, 0x1f, 0x1a, 0x9c, 0x37, 0x4b, 0x8b,
	0x1e, 0x03, 0x8a, 0x36, 0x8b, 0xcd, 0x36, 0xda, 0x6d, 0xd7, 0xd1, 0x4d, 0xe0, 0xad, 0x96, 0xab,
	0xc0, 0xb7, 0x1f, 0xa1, 0x11, 0x58, 0x91, 0xf7, 0x6d, 0xe0, 0x6f, 0xaf, 0x03, 0xdf, 0xd6, 0xd0,
	0x00, 0xcc, 0x68, 0xb3, 0x08, 0x37, 0x81, 0x6f, 0xeb, 0xe8, 0x02, 0x20, 0x58, 0xfb, 0x81, 0xbf,
	0xfb, 0x7e, 0x11, 0x7d, 0x67, 0xf7, 0xd4, 0xe1, 0xd6, 0xf3, 0x82, 0x28, 0xb2, 0x87, 0x12, 0x2c,
	0x17, 0xab, 0xeb, 0x6d, 0x18, 0xd8, 0x97, 0xe8, 0x03, 0x18, 0xad, 0xd6, 0xcb, 0x70, 0xb1, 0xeb,
	0xa8, 0x4f, 0xd1, 0x10, 0xfa, 0xde, 0x62, 0xed, 0x05, 0x32, 0xaf, 0xff, 0xec, 0x29, 0x98, 0x1b,
	0x4e, 0x0b, 0xcc, 0xdf, 0x22, 0x0b, 0xce, 0xb6, 0xeb, 0x28, 0xd8, 0xd8, 0x8f, 0x90, 0x09, 0xc6,
	0x4f, 0x41, 0x64, 0x6b, 0xe8, 0x1c, 0xf4, 0xf5, 0x0f, 0xb6, 0xfe, 0xcd, 0x17, 0x3f, 0xbf, 0x78,
	0xbf, 0xdf, 0xdc, 0x57, 0xf7, 0x98, 0x72, 0xbf, 0x3f, 0x57, 0xe4, 0xf3, 0xbf, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x22, 0x56, 0xb5, 0xc0, 0x25, 0x05, 0x00, 0x00,
}
