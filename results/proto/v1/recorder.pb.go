// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/results/proto/v1/recorder.proto

package resultspb

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type UpdateInvocationRequest struct {
	// If a request with same (invocation.id, request_id) was processed successfully
	// before, then this request is a noop.
	// In other words, UpdateInvocation is idempotent.
	// Required.
	//
	// Internally, UpdateInvocation uses request id to generate test result keys:
	// <request_id>-<a number unique within the request>
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Invocation to update.
	//
	// Each test result is appended.
	// Overwriting or removing test results is not supported.
	//
	// If a test log has contents, they are stored inline with the test result.
	// Such log MUST be less than 8KB. If it is more, clients must supply log URL.
	//
	// If a test variant has an exoneration, it is ensured
	// on the server.
	// Removing exoneration is not supported.
	//
	// Includes are appended. An existing invocation can be marked inconsequential.
	// Removing includes is not supported.
	//
	// If invocation.is_final is true, finalizes the invocation.
	// If invocation.deadline is specified, overwrites the
	// server-stored value.
	//
	// Invocation.update_token is required and must match the token
	// returned by CreateInvocation.
	Invocation           *Invocation `protobuf:"bytes,2,opt,name=invocation,proto3" json:"invocation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdateInvocationRequest) Reset()         { *m = UpdateInvocationRequest{} }
func (m *UpdateInvocationRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateInvocationRequest) ProtoMessage()    {}
func (*UpdateInvocationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7445f3675a5ef248, []int{0}
}

func (m *UpdateInvocationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateInvocationRequest.Unmarshal(m, b)
}
func (m *UpdateInvocationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateInvocationRequest.Marshal(b, m, deterministic)
}
func (m *UpdateInvocationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateInvocationRequest.Merge(m, src)
}
func (m *UpdateInvocationRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateInvocationRequest.Size(m)
}
func (m *UpdateInvocationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateInvocationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateInvocationRequest proto.InternalMessageInfo

func (m *UpdateInvocationRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *UpdateInvocationRequest) GetInvocation() *Invocation {
	if m != nil {
		return m.Invocation
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateInvocationRequest)(nil), "luci.resultsdb.UpdateInvocationRequest")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/results/proto/v1/recorder.proto", fileDescriptor_7445f3675a5ef248)
}

var fileDescriptor_7445f3675a5ef248 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x49, 0xcf, 0xd7, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c, 0xcd, 0xd5, 0xcb, 0x2f, 0x4a, 0xd7, 0xcf, 0x29, 0x4d, 0xce,
	0xd4, 0x2f, 0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0x29, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f,
	0x33, 0xd4, 0x2f, 0x4a, 0x4d, 0xce, 0x2f, 0x4a, 0x49, 0x2d, 0xd2, 0x03, 0x8b, 0x08, 0xf1, 0x81,
	0x54, 0xe9, 0x41, 0x55, 0xa5, 0x24, 0x49, 0x49, 0xa7, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0x42, 0xd4,
	0x27, 0x95, 0xa6, 0xe9, 0xa7, 0xe6, 0x16, 0x94, 0x54, 0x42, 0x14, 0x4b, 0x99, 0x11, 0x67, 0x45,
	0x66, 0x5e, 0x59, 0x7e, 0x72, 0x62, 0x49, 0x66, 0x7e, 0x1e, 0x44, 0x9f, 0x52, 0x09, 0x97, 0x78,
	0x68, 0x41, 0x4a, 0x62, 0x49, 0xaa, 0x27, 0x5c, 0x26, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44,
	0x48, 0x96, 0x8b, 0xab, 0x08, 0xc2, 0x8c, 0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0xe2, 0x84, 0x8a, 0x78, 0xa6, 0x08, 0x59, 0x71, 0x71, 0x21, 0x4c, 0x93, 0x60, 0x52, 0x60, 0xd4,
	0xe0, 0x36, 0x92, 0xd2, 0x43, 0x75, 0xb3, 0x1e, 0x92, 0xa9, 0x48, 0xaa, 0x8d, 0xd6, 0x33, 0x72,
	0x71, 0x04, 0x41, 0x7d, 0x2b, 0xe4, 0xc3, 0x25, 0xe0, 0x99, 0x57, 0x9c, 0x5a, 0x54, 0x82, 0x50,
	0x2c, 0x84, 0xc7, 0x20, 0x29, 0x3c, 0x72, 0x4a, 0x0c, 0x42, 0xa1, 0x5c, 0x02, 0xe8, 0x1e, 0x12,
	0x52, 0x47, 0xd7, 0x81, 0xc3, 0xcb, 0x52, 0x62, 0x7a, 0x90, 0x30, 0xd6, 0x83, 0x85, 0xb1, 0x9e,
	0x2b, 0x28, 0x8c, 0x95, 0x18, 0x9c, 0x0c, 0xa3, 0xf4, 0x89, 0x0a, 0x61, 0x6b, 0xa8, 0x40, 0x41,
	0x52, 0x12, 0x1b, 0x58, 0xcc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x13, 0xd6, 0xe6, 0x83, 0xfe,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RecorderClient is the client API for Recorder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecorderClient interface {
	// InsertInvocation creates a new invocation.
	//
	// If invocation with the given ID already exists, returns ALREADY_EXISTS
	// error code.
	//
	// It is inspired by ResultStoreUpload.CreateInvocation. Notable differences:
	// - no CreateInvocationRequest: it is "inlined", i.e. Invocation used instead.
	// - no request_id because it is unnecessary. ResultStore requires invocation_id
	//   to be set if request_id is set. Invocation id must be unique, therefore
	//   request_id gains nothing in addition to invocation id
	// - no authorization_token: CreateInvocation handler will generate a token
	//   associated with the invocation ID and return in the response.
	//   Also "authorization_token" is abstract: authorizing what exactly?
	//   This design uses the term "update_token".
	//   READ permissions  are defined elsewhere.
	//
	// Impl note: transactionally inserts a new spanner row with invocation id
	// primary key. If insertion fails with a conflict, returns ALREADY_EXISTS.
	InsertInvocation(ctx context.Context, in *Invocation, opts ...grpc.CallOption) (*Invocation, error)
	// A request to update an existing non-final invocation.
	//
	// Compared to ResultStoreUpload:
	// - In a sense, combines UpdateInvocation, FinishInvocation,
	//   CreateTarget, UpdateTarget, CreateConfiguredTarget, UpdateConfiguredTarget,
	//   CreateAction, UpdateAction, CreateConfiguration, UpdateConfiguration
	//   to be more aligned with the intended usage of the API and because it is a
	//   single Spanner mutation underneath.
	// - Finalization is coarse: entire invocation, as opposed to per sub-entity.
	// - More aligned with the domain model: test results can be only inserted.
	//
	// Impl note: Transactionally inserts a (invocation_id, request_id) row with
	// the rest of the payload. If request insertion fails, exits successfully.
	// Request table cleanup will be performed out of band.
	UpdateInvocation(ctx context.Context, in *UpdateInvocationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}
type recorderPRPCClient struct {
	client *prpc.Client
}

func NewRecorderPRPCClient(client *prpc.Client) RecorderClient {
	return &recorderPRPCClient{client}
}

func (c *recorderPRPCClient) InsertInvocation(ctx context.Context, in *Invocation, opts ...grpc.CallOption) (*Invocation, error) {
	out := new(Invocation)
	err := c.client.Call(ctx, "luci.resultsdb.Recorder", "InsertInvocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recorderPRPCClient) UpdateInvocation(ctx context.Context, in *UpdateInvocationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.client.Call(ctx, "luci.resultsdb.Recorder", "UpdateInvocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type recorderClient struct {
	cc *grpc.ClientConn
}

func NewRecorderClient(cc *grpc.ClientConn) RecorderClient {
	return &recorderClient{cc}
}

func (c *recorderClient) InsertInvocation(ctx context.Context, in *Invocation, opts ...grpc.CallOption) (*Invocation, error) {
	out := new(Invocation)
	err := c.cc.Invoke(ctx, "/luci.resultsdb.Recorder/InsertInvocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recorderClient) UpdateInvocation(ctx context.Context, in *UpdateInvocationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/luci.resultsdb.Recorder/UpdateInvocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RecorderServer is the server API for Recorder service.
type RecorderServer interface {
	// InsertInvocation creates a new invocation.
	//
	// If invocation with the given ID already exists, returns ALREADY_EXISTS
	// error code.
	//
	// It is inspired by ResultStoreUpload.CreateInvocation. Notable differences:
	// - no CreateInvocationRequest: it is "inlined", i.e. Invocation used instead.
	// - no request_id because it is unnecessary. ResultStore requires invocation_id
	//   to be set if request_id is set. Invocation id must be unique, therefore
	//   request_id gains nothing in addition to invocation id
	// - no authorization_token: CreateInvocation handler will generate a token
	//   associated with the invocation ID and return in the response.
	//   Also "authorization_token" is abstract: authorizing what exactly?
	//   This design uses the term "update_token".
	//   READ permissions  are defined elsewhere.
	//
	// Impl note: transactionally inserts a new spanner row with invocation id
	// primary key. If insertion fails with a conflict, returns ALREADY_EXISTS.
	InsertInvocation(context.Context, *Invocation) (*Invocation, error)
	// A request to update an existing non-final invocation.
	//
	// Compared to ResultStoreUpload:
	// - In a sense, combines UpdateInvocation, FinishInvocation,
	//   CreateTarget, UpdateTarget, CreateConfiguredTarget, UpdateConfiguredTarget,
	//   CreateAction, UpdateAction, CreateConfiguration, UpdateConfiguration
	//   to be more aligned with the intended usage of the API and because it is a
	//   single Spanner mutation underneath.
	// - Finalization is coarse: entire invocation, as opposed to per sub-entity.
	// - More aligned with the domain model: test results can be only inserted.
	//
	// Impl note: Transactionally inserts a (invocation_id, request_id) row with
	// the rest of the payload. If request insertion fails, exits successfully.
	// Request table cleanup will be performed out of band.
	UpdateInvocation(context.Context, *UpdateInvocationRequest) (*empty.Empty, error)
}

// UnimplementedRecorderServer can be embedded to have forward compatible implementations.
type UnimplementedRecorderServer struct {
}

func (*UnimplementedRecorderServer) InsertInvocation(ctx context.Context, req *Invocation) (*Invocation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertInvocation not implemented")
}
func (*UnimplementedRecorderServer) UpdateInvocation(ctx context.Context, req *UpdateInvocationRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInvocation not implemented")
}

func RegisterRecorderServer(s prpc.Registrar, srv RecorderServer) {
	s.RegisterService(&_Recorder_serviceDesc, srv)
}

func _Recorder_InsertInvocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Invocation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecorderServer).InsertInvocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/luci.resultsdb.Recorder/InsertInvocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecorderServer).InsertInvocation(ctx, req.(*Invocation))
	}
	return interceptor(ctx, in, info, handler)
}

func _Recorder_UpdateInvocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInvocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RecorderServer).UpdateInvocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/luci.resultsdb.Recorder/UpdateInvocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RecorderServer).UpdateInvocation(ctx, req.(*UpdateInvocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Recorder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "luci.resultsdb.Recorder",
	HandlerType: (*RecorderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertInvocation",
			Handler:    _Recorder_InsertInvocation_Handler,
		},
		{
			MethodName: "UpdateInvocation",
			Handler:    _Recorder_UpdateInvocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/results/proto/v1/recorder.proto",
}