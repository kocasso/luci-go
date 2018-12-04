// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/cipd/api/admin/v1/admin.proto

package api

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	mapper "go.chromium.org/luci/appengine/mapper"
	grpc "google.golang.org/grpc"
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

// Supported mapping jobs.
type MapperKind int32

const (
	MapperKind_MAPPER_KIND_UNSPECIFIED MapperKind = 0
	// Dump names of all packages to GAE logs, to test mapping jobs framework.
	MapperKind_ENUMERATE_PACKAGES MapperKind = 1
	// Find tags that don't pass ValidateInstanceTag and marks them.
	MapperKind_FIND_MALFORMED_TAGS MapperKind = 2
)

var MapperKind_name = map[int32]string{
	0: "MAPPER_KIND_UNSPECIFIED",
	1: "ENUMERATE_PACKAGES",
	2: "FIND_MALFORMED_TAGS",
}

var MapperKind_value = map[string]int32{
	"MAPPER_KIND_UNSPECIFIED": 0,
	"ENUMERATE_PACKAGES":      1,
	"FIND_MALFORMED_TAGS":     2,
}

func (x MapperKind) String() string {
	return proto.EnumName(MapperKind_name, int32(x))
}

func (MapperKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d3c583be32ae6c76, []int{0}
}

// Defines what a mapping job should do.
type JobConfig struct {
	Kind                 MapperKind `protobuf:"varint,1,opt,name=kind,proto3,enum=cipd.MapperKind" json:"kind,omitempty"`
	Comment              string     `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"`
	DryRun               bool       `protobuf:"varint,3,opt,name=dry_run,json=dryRun,proto3" json:"dry_run,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *JobConfig) Reset()         { *m = JobConfig{} }
func (m *JobConfig) String() string { return proto.CompactTextString(m) }
func (*JobConfig) ProtoMessage()    {}
func (*JobConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3c583be32ae6c76, []int{0}
}

func (m *JobConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobConfig.Unmarshal(m, b)
}
func (m *JobConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobConfig.Marshal(b, m, deterministic)
}
func (m *JobConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobConfig.Merge(m, src)
}
func (m *JobConfig) XXX_Size() int {
	return xxx_messageInfo_JobConfig.Size(m)
}
func (m *JobConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_JobConfig.DiscardUnknown(m)
}

var xxx_messageInfo_JobConfig proto.InternalMessageInfo

func (m *JobConfig) GetKind() MapperKind {
	if m != nil {
		return m.Kind
	}
	return MapperKind_MAPPER_KIND_UNSPECIFIED
}

func (m *JobConfig) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func (m *JobConfig) GetDryRun() bool {
	if m != nil {
		return m.DryRun
	}
	return false
}

// Identifies an instance of a mapping job.
type JobID struct {
	JobId                int64    `protobuf:"varint,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobID) Reset()         { *m = JobID{} }
func (m *JobID) String() string { return proto.CompactTextString(m) }
func (*JobID) ProtoMessage()    {}
func (*JobID) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3c583be32ae6c76, []int{1}
}

func (m *JobID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobID.Unmarshal(m, b)
}
func (m *JobID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobID.Marshal(b, m, deterministic)
}
func (m *JobID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobID.Merge(m, src)
}
func (m *JobID) XXX_Size() int {
	return xxx_messageInfo_JobID.Size(m)
}
func (m *JobID) XXX_DiscardUnknown() {
	xxx_messageInfo_JobID.DiscardUnknown(m)
}

var xxx_messageInfo_JobID proto.InternalMessageInfo

func (m *JobID) GetJobId() int64 {
	if m != nil {
		return m.JobId
	}
	return 0
}

// Details about a mapping job.
type JobState struct {
	// Original job config, exactly as it was submitted to LaunchJob.
	Config *JobConfig `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	// Current state of the job and all its shards.
	Info                 *mapper.JobInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *JobState) Reset()         { *m = JobState{} }
func (m *JobState) String() string { return proto.CompactTextString(m) }
func (*JobState) ProtoMessage()    {}
func (*JobState) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3c583be32ae6c76, []int{2}
}

func (m *JobState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobState.Unmarshal(m, b)
}
func (m *JobState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobState.Marshal(b, m, deterministic)
}
func (m *JobState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobState.Merge(m, src)
}
func (m *JobState) XXX_Size() int {
	return xxx_messageInfo_JobState.Size(m)
}
func (m *JobState) XXX_DiscardUnknown() {
	xxx_messageInfo_JobState.DiscardUnknown(m)
}

var xxx_messageInfo_JobState proto.InternalMessageInfo

func (m *JobState) GetConfig() *JobConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *JobState) GetInfo() *mapper.JobInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

// Result of running FixMarkedTags.
type TagFixReport struct {
	Fixed                []*TagFixReport_Tag `protobuf:"bytes,1,rep,name=fixed,proto3" json:"fixed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *TagFixReport) Reset()         { *m = TagFixReport{} }
func (m *TagFixReport) String() string { return proto.CompactTextString(m) }
func (*TagFixReport) ProtoMessage()    {}
func (*TagFixReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3c583be32ae6c76, []int{3}
}

func (m *TagFixReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagFixReport.Unmarshal(m, b)
}
func (m *TagFixReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagFixReport.Marshal(b, m, deterministic)
}
func (m *TagFixReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagFixReport.Merge(m, src)
}
func (m *TagFixReport) XXX_Size() int {
	return xxx_messageInfo_TagFixReport.Size(m)
}
func (m *TagFixReport) XXX_DiscardUnknown() {
	xxx_messageInfo_TagFixReport.DiscardUnknown(m)
}

var xxx_messageInfo_TagFixReport proto.InternalMessageInfo

func (m *TagFixReport) GetFixed() []*TagFixReport_Tag {
	if m != nil {
		return m.Fixed
	}
	return nil
}

type TagFixReport_Tag struct {
	Pkg                  string   `protobuf:"bytes,1,opt,name=pkg,proto3" json:"pkg,omitempty"`
	Instance             string   `protobuf:"bytes,2,opt,name=instance,proto3" json:"instance,omitempty"`
	BrokenTag            string   `protobuf:"bytes,3,opt,name=broken_tag,json=brokenTag,proto3" json:"broken_tag,omitempty"`
	FixedTag             string   `protobuf:"bytes,4,opt,name=fixed_tag,json=fixedTag,proto3" json:"fixed_tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TagFixReport_Tag) Reset()         { *m = TagFixReport_Tag{} }
func (m *TagFixReport_Tag) String() string { return proto.CompactTextString(m) }
func (*TagFixReport_Tag) ProtoMessage()    {}
func (*TagFixReport_Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3c583be32ae6c76, []int{3, 0}
}

func (m *TagFixReport_Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagFixReport_Tag.Unmarshal(m, b)
}
func (m *TagFixReport_Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagFixReport_Tag.Marshal(b, m, deterministic)
}
func (m *TagFixReport_Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagFixReport_Tag.Merge(m, src)
}
func (m *TagFixReport_Tag) XXX_Size() int {
	return xxx_messageInfo_TagFixReport_Tag.Size(m)
}
func (m *TagFixReport_Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_TagFixReport_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_TagFixReport_Tag proto.InternalMessageInfo

func (m *TagFixReport_Tag) GetPkg() string {
	if m != nil {
		return m.Pkg
	}
	return ""
}

func (m *TagFixReport_Tag) GetInstance() string {
	if m != nil {
		return m.Instance
	}
	return ""
}

func (m *TagFixReport_Tag) GetBrokenTag() string {
	if m != nil {
		return m.BrokenTag
	}
	return ""
}

func (m *TagFixReport_Tag) GetFixedTag() string {
	if m != nil {
		return m.FixedTag
	}
	return ""
}

func init() {
	proto.RegisterEnum("cipd.MapperKind", MapperKind_name, MapperKind_value)
	proto.RegisterType((*JobConfig)(nil), "cipd.JobConfig")
	proto.RegisterType((*JobID)(nil), "cipd.JobID")
	proto.RegisterType((*JobState)(nil), "cipd.JobState")
	proto.RegisterType((*TagFixReport)(nil), "cipd.TagFixReport")
	proto.RegisterType((*TagFixReport_Tag)(nil), "cipd.TagFixReport.Tag")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/cipd/api/admin/v1/admin.proto", fileDescriptor_d3c583be32ae6c76)
}

var fileDescriptor_d3c583be32ae6c76 = []byte{
	// 550 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xfd, 0xdc, 0xfc, 0x34, 0x9e, 0x7c, 0x94, 0x68, 0x11, 0x6d, 0xe4, 0x0a, 0x54, 0x45, 0x48,
	0x84, 0x0a, 0xd9, 0x60, 0x10, 0xf7, 0xa6, 0x71, 0x2a, 0xa7, 0x75, 0x89, 0x36, 0xe9, 0x4d, 0x6f,
	0xac, 0xb5, 0xbd, 0x71, 0xb7, 0xa9, 0x77, 0x2d, 0xc7, 0x46, 0xcd, 0x43, 0xf1, 0x10, 0xbc, 0x19,
	0xda, 0xdd, 0xd6, 0xb4, 0xc0, 0xdd, 0xcc, 0x9c, 0x33, 0x33, 0x67, 0x7e, 0xc0, 0xcd, 0x84, 0x9d,
	0x5c, 0x97, 0x22, 0x67, 0x75, 0x6e, 0x8b, 0x32, 0x73, 0x6e, 0xeb, 0x84, 0x39, 0x09, 0x2b, 0x52,
	0x87, 0x14, 0xcc, 0x21, 0x69, 0xce, 0xb8, 0xf3, 0xfd, 0xa3, 0x36, 0xec, 0xa2, 0x14, 0x95, 0x40,
	0x6d, 0x09, 0x5b, 0x87, 0x99, 0x10, 0xd9, 0x2d, 0x75, 0x54, 0x2c, 0xae, 0x57, 0x0e, 0xcd, 0x8b,
	0x6a, 0xab, 0x29, 0xd6, 0xe7, 0x7f, 0x96, 0x25, 0x45, 0x41, 0x79, 0xc6, 0x38, 0x75, 0x72, 0x69,
	0x96, 0x4e, 0x4e, 0x37, 0x1b, 0x92, 0xd1, 0x8d, 0xce, 0x1a, 0xa5, 0x60, 0xce, 0x44, 0x7c, 0x22,
	0xf8, 0x8a, 0x65, 0xe8, 0x0d, 0xb4, 0xd7, 0x8c, 0xa7, 0x43, 0xe3, 0xc8, 0x18, 0xef, 0xb9, 0x03,
	0x5b, 0x36, 0xb5, 0x43, 0x95, 0x77, 0xc6, 0x78, 0x8a, 0x15, 0x8a, 0x86, 0xb0, 0x9b, 0x88, 0x3c,
	0xa7, 0xbc, 0x1a, 0xee, 0x1c, 0x19, 0x63, 0x13, 0x3f, 0xb8, 0xe8, 0x00, 0x76, 0xd3, 0x72, 0x1b,
	0x95, 0x35, 0x1f, 0xb6, 0x8e, 0x8c, 0x71, 0x0f, 0x77, 0xd3, 0x72, 0x8b, 0x6b, 0x3e, 0x7a, 0x0d,
	0x9d, 0x99, 0x88, 0x83, 0x09, 0x7a, 0x09, 0xdd, 0x1b, 0x11, 0x47, 0x4c, 0xf7, 0x68, 0xe1, 0xce,
	0x8d, 0x88, 0x83, 0x74, 0xb4, 0x86, 0xde, 0x4c, 0xc4, 0x8b, 0x8a, 0x54, 0x14, 0xbd, 0x85, 0x6e,
	0xa2, 0xe4, 0x28, 0x4a, 0xdf, 0x7d, 0xae, 0x65, 0x34, 0x2a, 0xf1, 0x3d, 0x8c, 0xbe, 0x40, 0x9b,
	0xf1, 0x95, 0x50, 0x22, 0xfa, 0xee, 0xc8, 0x6e, 0x46, 0xb5, 0xf5, 0xa8, 0x76, 0x33, 0xaa, 0xec,
	0xcd, 0x57, 0x02, 0x2b, 0xfe, 0xe8, 0x87, 0x01, 0xff, 0x2f, 0x49, 0x36, 0x65, 0x77, 0x98, 0x16,
	0xa2, 0xac, 0xd0, 0x7b, 0xe8, 0xac, 0xd8, 0x1d, 0x95, 0x9a, 0x5a, 0xe3, 0xbe, 0xbb, 0xaf, 0x1b,
	0x3e, 0xa6, 0x48, 0x07, 0x6b, 0x92, 0x25, 0xa0, 0xb5, 0x24, 0x19, 0x1a, 0x40, 0xab, 0x58, 0x6b,
	0x8d, 0x26, 0x96, 0x26, 0xb2, 0xa0, 0xc7, 0xf8, 0xa6, 0x22, 0x3c, 0xa1, 0xf7, 0x8b, 0x69, 0x7c,
	0xf4, 0x0a, 0x20, 0x2e, 0xc5, 0x9a, 0xf2, 0xa8, 0x22, 0x99, 0x5a, 0x8e, 0x89, 0x4d, 0x1d, 0x91,
	0xc5, 0x0e, 0xc1, 0x54, 0xc5, 0x15, 0xda, 0xd6, 0xb9, 0x2a, 0xb0, 0x24, 0xd9, 0xf1, 0x15, 0xc0,
	0xef, 0x1b, 0xa0, 0x43, 0x38, 0x08, 0xbd, 0xf9, 0xdc, 0xc7, 0xd1, 0x59, 0x70, 0x31, 0x89, 0x2e,
	0x2f, 0x16, 0x73, 0xff, 0x24, 0x98, 0x06, 0xfe, 0x64, 0xf0, 0x1f, 0xda, 0x07, 0xe4, 0x5f, 0x5c,
	0x86, 0x3e, 0xf6, 0x96, 0x7e, 0x34, 0xf7, 0x4e, 0xce, 0xbc, 0x53, 0x7f, 0x31, 0x30, 0xd0, 0x01,
	0xbc, 0x98, 0x4a, 0x76, 0xe8, 0x9d, 0x4f, 0xbf, 0xe1, 0xd0, 0x9f, 0x44, 0x4b, 0xef, 0x74, 0x31,
	0xd8, 0x71, 0x7f, 0x1a, 0xd0, 0xf1, 0xe4, 0x9f, 0xa1, 0x77, 0x60, 0x9e, 0x93, 0x9a, 0x27, 0xd7,
	0x33, 0x11, 0xa3, 0x3f, 0x77, 0x6e, 0xf5, 0x9b, 0x40, 0x30, 0x41, 0x0e, 0xf4, 0xbc, 0x58, 0x94,
	0x95, 0x64, 0x3e, 0x06, 0xac, 0x7d, 0x5b, 0x3f, 0xa8, 0xfd, 0xf0, 0xa0, 0xb6, 0x2f, 0x1f, 0x14,
	0x1d, 0x43, 0xff, 0x94, 0x56, 0xcd, 0x85, 0x9f, 0xe4, 0xec, 0x35, 0x8e, 0x06, 0x3f, 0xc0, 0xb3,
	0x29, 0xbb, 0x0b, 0x49, 0xb9, 0x56, 0xd3, 0x6f, 0x9e, 0xb2, 0xd1, 0xdf, 0xb7, 0xf9, 0xda, 0xb9,
	0x6a, 0x91, 0x82, 0xc5, 0x5d, 0xd5, 0xf4, 0xd3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb0, 0x79,
	0x5c, 0x9b, 0x5f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminClient interface {
	// Launches a mapping job that examines and/or fixes datastore entities.
	LaunchJob(ctx context.Context, in *JobConfig, opts ...grpc.CallOption) (*JobID, error)
	// Initiates an abort of a mapping job.
	AbortJob(ctx context.Context, in *JobID, opts ...grpc.CallOption) (*empty.Empty, error)
	// Returns state of a mapping job.
	GetJobState(ctx context.Context, in *JobID, opts ...grpc.CallOption) (*JobState, error)
	// Fixes (right inside the handler) tags marked by the given mapper job.
	FixMarkedTags(ctx context.Context, in *JobID, opts ...grpc.CallOption) (*TagFixReport, error)
}
type adminPRPCClient struct {
	client *prpc.Client
}

func NewAdminPRPCClient(client *prpc.Client) AdminClient {
	return &adminPRPCClient{client}
}

func (c *adminPRPCClient) LaunchJob(ctx context.Context, in *JobConfig, opts ...grpc.CallOption) (*JobID, error) {
	out := new(JobID)
	err := c.client.Call(ctx, "cipd.Admin", "LaunchJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminPRPCClient) AbortJob(ctx context.Context, in *JobID, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.client.Call(ctx, "cipd.Admin", "AbortJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminPRPCClient) GetJobState(ctx context.Context, in *JobID, opts ...grpc.CallOption) (*JobState, error) {
	out := new(JobState)
	err := c.client.Call(ctx, "cipd.Admin", "GetJobState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminPRPCClient) FixMarkedTags(ctx context.Context, in *JobID, opts ...grpc.CallOption) (*TagFixReport, error) {
	out := new(TagFixReport)
	err := c.client.Call(ctx, "cipd.Admin", "FixMarkedTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type adminClient struct {
	cc *grpc.ClientConn
}

func NewAdminClient(cc *grpc.ClientConn) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) LaunchJob(ctx context.Context, in *JobConfig, opts ...grpc.CallOption) (*JobID, error) {
	out := new(JobID)
	err := c.cc.Invoke(ctx, "/cipd.Admin/LaunchJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) AbortJob(ctx context.Context, in *JobID, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/cipd.Admin/AbortJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) GetJobState(ctx context.Context, in *JobID, opts ...grpc.CallOption) (*JobState, error) {
	out := new(JobState)
	err := c.cc.Invoke(ctx, "/cipd.Admin/GetJobState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminClient) FixMarkedTags(ctx context.Context, in *JobID, opts ...grpc.CallOption) (*TagFixReport, error) {
	out := new(TagFixReport)
	err := c.cc.Invoke(ctx, "/cipd.Admin/FixMarkedTags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
type AdminServer interface {
	// Launches a mapping job that examines and/or fixes datastore entities.
	LaunchJob(context.Context, *JobConfig) (*JobID, error)
	// Initiates an abort of a mapping job.
	AbortJob(context.Context, *JobID) (*empty.Empty, error)
	// Returns state of a mapping job.
	GetJobState(context.Context, *JobID) (*JobState, error)
	// Fixes (right inside the handler) tags marked by the given mapper job.
	FixMarkedTags(context.Context, *JobID) (*TagFixReport, error)
}

func RegisterAdminServer(s prpc.Registrar, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_LaunchJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).LaunchJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cipd.Admin/LaunchJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).LaunchJob(ctx, req.(*JobConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_AbortJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).AbortJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cipd.Admin/AbortJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).AbortJob(ctx, req.(*JobID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_GetJobState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).GetJobState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cipd.Admin/GetJobState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).GetJobState(ctx, req.(*JobID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Admin_FixMarkedTags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).FixMarkedTags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cipd.Admin/FixMarkedTags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).FixMarkedTags(ctx, req.(*JobID))
	}
	return interceptor(ctx, in, info, handler)
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cipd.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LaunchJob",
			Handler:    _Admin_LaunchJob_Handler,
		},
		{
			MethodName: "AbortJob",
			Handler:    _Admin_AbortJob_Handler,
		},
		{
			MethodName: "GetJobState",
			Handler:    _Admin_GetJobState_Handler,
		},
		{
			MethodName: "FixMarkedTags",
			Handler:    _Admin_FixMarkedTags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/cipd/api/admin/v1/admin.proto",
}
