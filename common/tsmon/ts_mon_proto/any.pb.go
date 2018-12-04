// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/common/tsmon/ts_mon_proto/any.proto

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

type Any struct {
	TypeUrl              *string  `protobuf:"bytes,1,opt,name=type_url,json=typeUrl" json:"type_url,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Any) Reset()         { *m = Any{} }
func (m *Any) String() string { return proto.CompactTextString(m) }
func (*Any) ProtoMessage()    {}
func (*Any) Descriptor() ([]byte, []int) {
	return fileDescriptor_58f91f4e40f17744, []int{0}
}

func (m *Any) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Any.Unmarshal(m, b)
}
func (m *Any) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Any.Marshal(b, m, deterministic)
}
func (m *Any) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Any.Merge(m, src)
}
func (m *Any) XXX_Size() int {
	return xxx_messageInfo_Any.Size(m)
}
func (m *Any) XXX_DiscardUnknown() {
	xxx_messageInfo_Any.DiscardUnknown(m)
}

var xxx_messageInfo_Any proto.InternalMessageInfo

func (m *Any) GetTypeUrl() string {
	if m != nil && m.TypeUrl != nil {
		return *m.TypeUrl
	}
	return ""
}

func (m *Any) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Any)(nil), "ts_mon.proto.Any")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/common/tsmon/ts_mon_proto/any.proto", fileDescriptor_58f91f4e40f17744)
}

var fileDescriptor_58f91f4e40f17744 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x48, 0xcf, 0xd7, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c, 0xcd, 0xd5, 0xcb, 0x2f, 0x4a, 0xd7, 0xcf, 0x29, 0x4d, 0xce,
	0xd4, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x2f, 0x29, 0x86, 0x90, 0xf1, 0xb9, 0xf9, 0x79,
	0xf1, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0xfa, 0x89, 0x79, 0x95, 0x7a, 0x60, 0x96, 0x10, 0x0f, 0x44,
	0x1c, 0xc2, 0x53, 0xb2, 0xe3, 0x62, 0x76, 0xcc, 0xab, 0x14, 0x92, 0xe5, 0xe2, 0x28, 0xa9, 0x2c,
	0x48, 0x8d, 0x2f, 0x2d, 0xca, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x74, 0x62, 0xe2, 0x60, 0x0a,
	0x62, 0x07, 0x89, 0x85, 0x16, 0xe5, 0x08, 0x49, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a,
	0x30, 0x29, 0x30, 0x6a, 0xf0, 0x38, 0x31, 0x71, 0x30, 0x06, 0x41, 0x04, 0x00, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x4c, 0x2e, 0xf4, 0xf8, 0x88, 0x00, 0x00, 0x00,
}
