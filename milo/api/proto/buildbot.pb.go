// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/milo/api/proto/buildbot.proto

package milo

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// The request containing the name of the master.
type MasterRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// if true, exclude response data that the foundation team is actively trying
	// to deprecate:
	// - slave info
	ExcludeDeprecated bool `protobuf:"varint,10,opt,name=exclude_deprecated,json=excludeDeprecated,proto3" json:"exclude_deprecated,omitempty"`
	// If true, turn off emulation mode.
	NoEmulation          bool     `protobuf:"varint,11,opt,name=no_emulation,json=noEmulation,proto3" json:"no_emulation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MasterRequest) Reset()         { *m = MasterRequest{} }
func (m *MasterRequest) String() string { return proto.CompactTextString(m) }
func (*MasterRequest) ProtoMessage()    {}
func (*MasterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_42752ec01fa9d3bf, []int{0}
}

func (m *MasterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MasterRequest.Unmarshal(m, b)
}
func (m *MasterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MasterRequest.Marshal(b, m, deterministic)
}
func (m *MasterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MasterRequest.Merge(m, src)
}
func (m *MasterRequest) XXX_Size() int {
	return xxx_messageInfo_MasterRequest.Size(m)
}
func (m *MasterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MasterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MasterRequest proto.InternalMessageInfo

func (m *MasterRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MasterRequest) GetExcludeDeprecated() bool {
	if m != nil {
		return m.ExcludeDeprecated
	}
	return false
}

func (m *MasterRequest) GetNoEmulation() bool {
	if m != nil {
		return m.NoEmulation
	}
	return false
}

// The response message containing master information.
type CompressedMasterJSON struct {
	// Whether the master is internal or not.
	Internal bool `protobuf:"varint,1,opt,name=internal,proto3" json:"internal,omitempty"`
	// Timestamp of the freshness of the master data.
	Modified *timestamp.Timestamp `protobuf:"bytes,2,opt,name=modified,proto3" json:"modified,omitempty"`
	// Gzipped json data of the master.
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompressedMasterJSON) Reset()         { *m = CompressedMasterJSON{} }
func (m *CompressedMasterJSON) String() string { return proto.CompactTextString(m) }
func (*CompressedMasterJSON) ProtoMessage()    {}
func (*CompressedMasterJSON) Descriptor() ([]byte, []int) {
	return fileDescriptor_42752ec01fa9d3bf, []int{1}
}

func (m *CompressedMasterJSON) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompressedMasterJSON.Unmarshal(m, b)
}
func (m *CompressedMasterJSON) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompressedMasterJSON.Marshal(b, m, deterministic)
}
func (m *CompressedMasterJSON) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompressedMasterJSON.Merge(m, src)
}
func (m *CompressedMasterJSON) XXX_Size() int {
	return xxx_messageInfo_CompressedMasterJSON.Size(m)
}
func (m *CompressedMasterJSON) XXX_DiscardUnknown() {
	xxx_messageInfo_CompressedMasterJSON.DiscardUnknown(m)
}

var xxx_messageInfo_CompressedMasterJSON proto.InternalMessageInfo

func (m *CompressedMasterJSON) GetInternal() bool {
	if m != nil {
		return m.Internal
	}
	return false
}

func (m *CompressedMasterJSON) GetModified() *timestamp.Timestamp {
	if m != nil {
		return m.Modified
	}
	return nil
}

func (m *CompressedMasterJSON) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// The request for a specific build.
type BuildbotBuildRequest struct {
	Master   string `protobuf:"bytes,1,opt,name=master,proto3" json:"master,omitempty"`
	Builder  string `protobuf:"bytes,2,opt,name=builder,proto3" json:"builder,omitempty"`
	BuildNum int64  `protobuf:"varint,3,opt,name=build_num,json=buildNum,proto3" json:"build_num,omitempty"`
	// if true, exclude response data that the foundation team is actively trying
	// to deprecate:
	// - slave info
	ExcludeDeprecated    bool     `protobuf:"varint,10,opt,name=exclude_deprecated,json=excludeDeprecated,proto3" json:"exclude_deprecated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildbotBuildRequest) Reset()         { *m = BuildbotBuildRequest{} }
func (m *BuildbotBuildRequest) String() string { return proto.CompactTextString(m) }
func (*BuildbotBuildRequest) ProtoMessage()    {}
func (*BuildbotBuildRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_42752ec01fa9d3bf, []int{2}
}

func (m *BuildbotBuildRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildbotBuildRequest.Unmarshal(m, b)
}
func (m *BuildbotBuildRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildbotBuildRequest.Marshal(b, m, deterministic)
}
func (m *BuildbotBuildRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildbotBuildRequest.Merge(m, src)
}
func (m *BuildbotBuildRequest) XXX_Size() int {
	return xxx_messageInfo_BuildbotBuildRequest.Size(m)
}
func (m *BuildbotBuildRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildbotBuildRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BuildbotBuildRequest proto.InternalMessageInfo

func (m *BuildbotBuildRequest) GetMaster() string {
	if m != nil {
		return m.Master
	}
	return ""
}

func (m *BuildbotBuildRequest) GetBuilder() string {
	if m != nil {
		return m.Builder
	}
	return ""
}

func (m *BuildbotBuildRequest) GetBuildNum() int64 {
	if m != nil {
		return m.BuildNum
	}
	return 0
}

func (m *BuildbotBuildRequest) GetExcludeDeprecated() bool {
	if m != nil {
		return m.ExcludeDeprecated
	}
	return false
}

// The response message for a specific build.
type BuildbotBuildJSON struct {
	// Json data of the build.
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildbotBuildJSON) Reset()         { *m = BuildbotBuildJSON{} }
func (m *BuildbotBuildJSON) String() string { return proto.CompactTextString(m) }
func (*BuildbotBuildJSON) ProtoMessage()    {}
func (*BuildbotBuildJSON) Descriptor() ([]byte, []int) {
	return fileDescriptor_42752ec01fa9d3bf, []int{3}
}

func (m *BuildbotBuildJSON) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildbotBuildJSON.Unmarshal(m, b)
}
func (m *BuildbotBuildJSON) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildbotBuildJSON.Marshal(b, m, deterministic)
}
func (m *BuildbotBuildJSON) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildbotBuildJSON.Merge(m, src)
}
func (m *BuildbotBuildJSON) XXX_Size() int {
	return xxx_messageInfo_BuildbotBuildJSON.Size(m)
}
func (m *BuildbotBuildJSON) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildbotBuildJSON.DiscardUnknown(m)
}

var xxx_messageInfo_BuildbotBuildJSON proto.InternalMessageInfo

func (m *BuildbotBuildJSON) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// The request for multiple build on a builder.
type BuildbotBuildsRequest struct {
	Master  string `protobuf:"bytes,1,opt,name=master,proto3" json:"master,omitempty"`
	Builder string `protobuf:"bytes,2,opt,name=builder,proto3" json:"builder,omitempty"`
	// Limit to the number of builds to return (default: 20).
	Limit int32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// Include ongoing builds (default: false).
	IncludeCurrent bool `protobuf:"varint,4,opt,name=include_current,json=includeCurrent,proto3" json:"include_current,omitempty"`
	// if true, exclude response data that the foundation team is actively trying
	// to deprecate:
	// - slave info
	ExcludeDeprecated bool `protobuf:"varint,10,opt,name=exclude_deprecated,json=excludeDeprecated,proto3" json:"exclude_deprecated,omitempty"`
	// If true, turn off emulation mode.
	NoEmulation          bool     `protobuf:"varint,11,opt,name=no_emulation,json=noEmulation,proto3" json:"no_emulation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildbotBuildsRequest) Reset()         { *m = BuildbotBuildsRequest{} }
func (m *BuildbotBuildsRequest) String() string { return proto.CompactTextString(m) }
func (*BuildbotBuildsRequest) ProtoMessage()    {}
func (*BuildbotBuildsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_42752ec01fa9d3bf, []int{4}
}

func (m *BuildbotBuildsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildbotBuildsRequest.Unmarshal(m, b)
}
func (m *BuildbotBuildsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildbotBuildsRequest.Marshal(b, m, deterministic)
}
func (m *BuildbotBuildsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildbotBuildsRequest.Merge(m, src)
}
func (m *BuildbotBuildsRequest) XXX_Size() int {
	return xxx_messageInfo_BuildbotBuildsRequest.Size(m)
}
func (m *BuildbotBuildsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildbotBuildsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BuildbotBuildsRequest proto.InternalMessageInfo

func (m *BuildbotBuildsRequest) GetMaster() string {
	if m != nil {
		return m.Master
	}
	return ""
}

func (m *BuildbotBuildsRequest) GetBuilder() string {
	if m != nil {
		return m.Builder
	}
	return ""
}

func (m *BuildbotBuildsRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *BuildbotBuildsRequest) GetIncludeCurrent() bool {
	if m != nil {
		return m.IncludeCurrent
	}
	return false
}

func (m *BuildbotBuildsRequest) GetExcludeDeprecated() bool {
	if m != nil {
		return m.ExcludeDeprecated
	}
	return false
}

func (m *BuildbotBuildsRequest) GetNoEmulation() bool {
	if m != nil {
		return m.NoEmulation
	}
	return false
}

// The response message for multiple builds in a builder.
type BuildbotBuildsJSON struct {
	// builds is the list of builds resulting from the builds request.
	Builds               []*BuildbotBuildJSON `protobuf:"bytes,1,rep,name=builds,proto3" json:"builds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BuildbotBuildsJSON) Reset()         { *m = BuildbotBuildsJSON{} }
func (m *BuildbotBuildsJSON) String() string { return proto.CompactTextString(m) }
func (*BuildbotBuildsJSON) ProtoMessage()    {}
func (*BuildbotBuildsJSON) Descriptor() ([]byte, []int) {
	return fileDescriptor_42752ec01fa9d3bf, []int{5}
}

func (m *BuildbotBuildsJSON) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildbotBuildsJSON.Unmarshal(m, b)
}
func (m *BuildbotBuildsJSON) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildbotBuildsJSON.Marshal(b, m, deterministic)
}
func (m *BuildbotBuildsJSON) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildbotBuildsJSON.Merge(m, src)
}
func (m *BuildbotBuildsJSON) XXX_Size() int {
	return xxx_messageInfo_BuildbotBuildsJSON.Size(m)
}
func (m *BuildbotBuildsJSON) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildbotBuildsJSON.DiscardUnknown(m)
}

var xxx_messageInfo_BuildbotBuildsJSON proto.InternalMessageInfo

func (m *BuildbotBuildsJSON) GetBuilds() []*BuildbotBuildJSON {
	if m != nil {
		return m.Builds
	}
	return nil
}

func init() {
	proto.RegisterType((*MasterRequest)(nil), "milo.MasterRequest")
	proto.RegisterType((*CompressedMasterJSON)(nil), "milo.CompressedMasterJSON")
	proto.RegisterType((*BuildbotBuildRequest)(nil), "milo.BuildbotBuildRequest")
	proto.RegisterType((*BuildbotBuildJSON)(nil), "milo.BuildbotBuildJSON")
	proto.RegisterType((*BuildbotBuildsRequest)(nil), "milo.BuildbotBuildsRequest")
	proto.RegisterType((*BuildbotBuildsJSON)(nil), "milo.BuildbotBuildsJSON")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/milo/api/proto/buildbot.proto", fileDescriptor_42752ec01fa9d3bf)
}

var fileDescriptor_42752ec01fa9d3bf = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x9e, 0xd7, 0xb4, 0xa4, 0xa7, 0x03, 0x36, 0xd3, 0xb1, 0x28, 0xbb, 0x20, 0xe4, 0x66, 0xb9,
	0x21, 0x91, 0x8a, 0xc4, 0x03, 0x30, 0xd0, 0xa4, 0xc2, 0x06, 0x32, 0xdc, 0x57, 0x69, 0x72, 0x56,
	0x2c, 0xc5, 0x71, 0x70, 0x6c, 0x89, 0x2b, 0x9e, 0x82, 0x07, 0xe4, 0x35, 0xb8, 0x43, 0xb5, 0x93,
	0x8a, 0xb2, 0x70, 0x01, 0xe2, 0x2a, 0xe7, 0xe7, 0xcb, 0xf1, 0xf7, 0x7d, 0x3e, 0x86, 0xc5, 0x46,
	0xa6, 0xc5, 0x27, 0x25, 0x05, 0x37, 0x22, 0x95, 0x6a, 0x93, 0x55, 0xa6, 0xe0, 0x99, 0xe0, 0x95,
	0xcc, 0xf2, 0x86, 0x67, 0x8d, 0x92, 0x5a, 0x66, 0x6b, 0xc3, 0xab, 0x72, 0x2d, 0x75, 0x6a, 0x53,
	0xea, 0x6d, 0xdb, 0xe1, 0x93, 0x8d, 0x94, 0x9b, 0x0a, 0x1d, 0x64, 0x6d, 0x6e, 0x33, 0xcd, 0x05,
	0xb6, 0x3a, 0x17, 0x8d, 0x83, 0xc5, 0x06, 0xee, 0x5f, 0xe7, 0xad, 0x46, 0xc5, 0xf0, 0xb3, 0xc1,
	0x56, 0x53, 0x0a, 0x5e, 0x9d, 0x0b, 0x0c, 0x48, 0x44, 0x92, 0x29, 0xb3, 0x31, 0x7d, 0x06, 0x14,
	0xbf, 0x14, 0x95, 0x29, 0x71, 0x55, 0x62, 0xa3, 0xb0, 0xc8, 0x35, 0x96, 0x01, 0x44, 0x24, 0xf1,
	0xd9, 0x49, 0xd7, 0x79, 0xb5, 0x6b, 0xd0, 0xa7, 0x70, 0x54, 0xcb, 0x15, 0x0a, 0x53, 0xe5, 0x9a,
	0xcb, 0x3a, 0x98, 0x59, 0xe0, 0xac, 0x96, 0xaf, 0xfb, 0x52, 0xfc, 0x15, 0xe6, 0x97, 0x52, 0x34,
	0x0a, 0xdb, 0x16, 0x4b, 0x47, 0x60, 0xf9, 0xe1, 0xdd, 0x0d, 0x0d, 0xc1, 0xe7, 0xb5, 0x46, 0x55,
	0xe7, 0x95, 0x65, 0xe0, 0xb3, 0x5d, 0x4e, 0x5f, 0x80, 0x2f, 0x64, 0xc9, 0x6f, 0x39, 0x96, 0xc1,
	0x61, 0x44, 0x92, 0xd9, 0x22, 0x4c, 0x9d, 0xbc, 0xb4, 0x97, 0x97, 0x7e, 0xec, 0xe5, 0xb1, 0x1d,
	0x76, 0xab, 0xa8, 0xcc, 0x75, 0x1e, 0x8c, 0x22, 0x92, 0x1c, 0x31, 0x1b, 0xc7, 0xdf, 0x08, 0xcc,
	0x5f, 0x76, 0x86, 0xd9, 0x6f, 0x2f, 0xff, 0x31, 0x4c, 0x84, 0xa5, 0xd3, 0x19, 0xd0, 0x65, 0x34,
	0x80, 0x7b, 0xd6, 0x60, 0x54, 0xf6, 0xec, 0x29, 0xeb, 0x53, 0x7a, 0x0e, 0x53, 0x1b, 0xae, 0x6a,
	0x23, 0xec, 0x19, 0x23, 0xe6, 0xdb, 0xc2, 0x8d, 0x11, 0x7f, 0xe9, 0x5c, 0x7c, 0x01, 0x27, 0x7b,
	0xac, 0xac, 0x27, 0x3d, 0x7f, 0xf2, 0x0b, 0xff, 0xef, 0x04, 0x4e, 0xf7, 0x90, 0xed, 0xbf, 0x0b,
	0x98, 0xc3, 0xb8, 0xe2, 0x82, 0x6b, 0x4b, 0x7e, 0xcc, 0x5c, 0x42, 0x2f, 0xe0, 0x21, 0xaf, 0x1d,
	0xf3, 0xc2, 0x28, 0x85, 0xb5, 0x0e, 0x3c, 0x4b, 0xfb, 0x41, 0x57, 0xbe, 0x74, 0xd5, 0xff, 0xbf,
	0x1c, 0x4b, 0xcf, 0x1f, 0x1f, 0x4f, 0xe2, 0x37, 0x40, 0xf7, 0x15, 0x5a, 0x33, 0x32, 0x98, 0x58,
	0xde, 0x6d, 0x40, 0xa2, 0x51, 0x32, 0x5b, 0x9c, 0xa5, 0xdb, 0x3d, 0x4f, 0xef, 0xb8, 0xc6, 0x3a,
	0xd8, 0xd2, 0xf3, 0x0f, 0x8f, 0x47, 0x8b, 0x1f, 0x04, 0xfc, 0x1e, 0x43, 0xdf, 0xc2, 0xd9, 0x15,
	0xea, 0xc1, 0xfd, 0x7b, 0xe4, 0xc6, 0xed, 0x3d, 0x89, 0x30, 0x74, 0xc5, 0xa1, 0x1f, 0xe2, 0x03,
	0x7a, 0x0d, 0xf3, 0x2b, 0xd4, 0x77, 0xaf, 0x2d, 0x1c, 0x60, 0xd6, 0x4f, 0xfc, 0x13, 0xeb, 0xf8,
	0x80, 0xbe, 0x87, 0xd3, 0xdf, 0xc7, 0x39, 0xe5, 0xe7, 0x03, 0xff, 0xf4, 0xb7, 0x1e, 0x06, 0x43,
	0x4d, 0x37, 0x71, 0x3d, 0xb1, 0xaf, 0xe3, 0xf9, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x28, 0x2d,
	0x6f, 0x96, 0x46, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BuildbotClient is the client API for Buildbot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BuildbotClient interface {
	GetCompressedMasterJSON(ctx context.Context, in *MasterRequest, opts ...grpc.CallOption) (*CompressedMasterJSON, error)
	GetBuildbotBuildJSON(ctx context.Context, in *BuildbotBuildRequest, opts ...grpc.CallOption) (*BuildbotBuildJSON, error)
	GetBuildbotBuildsJSON(ctx context.Context, in *BuildbotBuildsRequest, opts ...grpc.CallOption) (*BuildbotBuildsJSON, error)
}
type buildbotPRPCClient struct {
	client *prpc.Client
}

func NewBuildbotPRPCClient(client *prpc.Client) BuildbotClient {
	return &buildbotPRPCClient{client}
}

func (c *buildbotPRPCClient) GetCompressedMasterJSON(ctx context.Context, in *MasterRequest, opts ...grpc.CallOption) (*CompressedMasterJSON, error) {
	out := new(CompressedMasterJSON)
	err := c.client.Call(ctx, "milo.Buildbot", "GetCompressedMasterJSON", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildbotPRPCClient) GetBuildbotBuildJSON(ctx context.Context, in *BuildbotBuildRequest, opts ...grpc.CallOption) (*BuildbotBuildJSON, error) {
	out := new(BuildbotBuildJSON)
	err := c.client.Call(ctx, "milo.Buildbot", "GetBuildbotBuildJSON", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildbotPRPCClient) GetBuildbotBuildsJSON(ctx context.Context, in *BuildbotBuildsRequest, opts ...grpc.CallOption) (*BuildbotBuildsJSON, error) {
	out := new(BuildbotBuildsJSON)
	err := c.client.Call(ctx, "milo.Buildbot", "GetBuildbotBuildsJSON", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type buildbotClient struct {
	cc *grpc.ClientConn
}

func NewBuildbotClient(cc *grpc.ClientConn) BuildbotClient {
	return &buildbotClient{cc}
}

func (c *buildbotClient) GetCompressedMasterJSON(ctx context.Context, in *MasterRequest, opts ...grpc.CallOption) (*CompressedMasterJSON, error) {
	out := new(CompressedMasterJSON)
	err := c.cc.Invoke(ctx, "/milo.Buildbot/GetCompressedMasterJSON", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildbotClient) GetBuildbotBuildJSON(ctx context.Context, in *BuildbotBuildRequest, opts ...grpc.CallOption) (*BuildbotBuildJSON, error) {
	out := new(BuildbotBuildJSON)
	err := c.cc.Invoke(ctx, "/milo.Buildbot/GetBuildbotBuildJSON", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildbotClient) GetBuildbotBuildsJSON(ctx context.Context, in *BuildbotBuildsRequest, opts ...grpc.CallOption) (*BuildbotBuildsJSON, error) {
	out := new(BuildbotBuildsJSON)
	err := c.cc.Invoke(ctx, "/milo.Buildbot/GetBuildbotBuildsJSON", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuildbotServer is the server API for Buildbot service.
type BuildbotServer interface {
	GetCompressedMasterJSON(context.Context, *MasterRequest) (*CompressedMasterJSON, error)
	GetBuildbotBuildJSON(context.Context, *BuildbotBuildRequest) (*BuildbotBuildJSON, error)
	GetBuildbotBuildsJSON(context.Context, *BuildbotBuildsRequest) (*BuildbotBuildsJSON, error)
}

func RegisterBuildbotServer(s prpc.Registrar, srv BuildbotServer) {
	s.RegisterService(&_Buildbot_serviceDesc, srv)
}

func _Buildbot_GetCompressedMasterJSON_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MasterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildbotServer).GetCompressedMasterJSON(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milo.Buildbot/GetCompressedMasterJSON",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildbotServer).GetCompressedMasterJSON(ctx, req.(*MasterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Buildbot_GetBuildbotBuildJSON_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildbotBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildbotServer).GetBuildbotBuildJSON(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milo.Buildbot/GetBuildbotBuildJSON",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildbotServer).GetBuildbotBuildJSON(ctx, req.(*BuildbotBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Buildbot_GetBuildbotBuildsJSON_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildbotBuildsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildbotServer).GetBuildbotBuildsJSON(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/milo.Buildbot/GetBuildbotBuildsJSON",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildbotServer).GetBuildbotBuildsJSON(ctx, req.(*BuildbotBuildsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Buildbot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "milo.Buildbot",
	HandlerType: (*BuildbotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCompressedMasterJSON",
			Handler:    _Buildbot_GetCompressedMasterJSON_Handler,
		},
		{
			MethodName: "GetBuildbotBuildJSON",
			Handler:    _Buildbot_GetBuildbotBuildJSON_Handler,
		},
		{
			MethodName: "GetBuildbotBuildsJSON",
			Handler:    _Buildbot_GetBuildbotBuildsJSON_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/milo/api/proto/buildbot.proto",
}
