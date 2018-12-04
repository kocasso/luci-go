// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/common/tsmon/ts_mon_proto/endpoint.proto

package ts_mon_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Request struct {
	Payload              *MetricsPayload `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_933117f6dff2e889, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetPayload() *MetricsPayload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "ts_mon.proto.Request")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/common/tsmon/ts_mon_proto/endpoint.proto", fileDescriptor_933117f6dff2e889)
}

var fileDescriptor_933117f6dff2e889 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4d, 0xcf, 0xd7, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c, 0xcd, 0xd5, 0xcb, 0x2f, 0x4a, 0xd7, 0xcf, 0x29, 0x4d, 0xce,
	0xd4, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x2f, 0x29, 0x86, 0x90, 0xf1, 0xb9, 0xf9, 0x79,
	0xf1, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0xfa, 0xa9, 0x79, 0x29, 0x05, 0xf9, 0x99, 0x79, 0x25, 0x7a,
	0x60, 0xae, 0x10, 0x0f, 0x44, 0x12, 0xc2, 0x93, 0xb2, 0x21, 0xd1, 0xb0, 0xdc, 0xd4, 0x92, 0xa2,
	0xcc, 0xe4, 0x62, 0x88, 0x6e, 0x25, 0x47, 0x2e, 0xf6, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12,
	0x21, 0x33, 0x2e, 0xf6, 0x82, 0xc4, 0xca, 0x9c, 0xfc, 0xc4, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0x6e, 0x23, 0x19, 0x3d, 0x64, 0x8b, 0xf4, 0x7c, 0x21, 0x1a, 0x03, 0x20, 0x6a, 0x82, 0x60, 0x8a,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xb9, 0xdb, 0x01, 0xce, 0x00, 0x00, 0x00,
}
