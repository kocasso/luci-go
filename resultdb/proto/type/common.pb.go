// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/resultdb/proto/type/common.proto

package typepb

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

// A key-value map describing one variant of a test case.
//
// The same test case can be executed in different ways, for example on
// different OS, GPUs, with different compile options, runtime flags or even
// with different values of the test parameter (for parameterized tests).
// A variant definition captures one variant.
// A test case with a specific variant definition is called test variant.
//
// Guidelines for variant definition design:
// - This rule guides what keys MUST be present in the definition.
//   A single expected result of a given test variant is enough to consider it
//   passing (potentially flakily). If it is important to differentiate across
//   a certain dimension (e.g. whether web tests are executed with or without
//   site per process isolation), then there MUST be a key that captures the
//   dimension (e.g. a name from test_suites.pyl).
//   Otherwise, a pass in one variant will hide a failure of another one.
//
// - This rule guides what keys MUST NOT be present in the definition.
//   A change in the key-value set essentially resets the test result history.
//   For example, if GN args are among variant key-value pairs, then adding a
//   new GN arg changes the identity of the test variant and resets its history.
//
// In Chromium, typical variant keys are:
// - bucket: the LUCI bucket, e.g. "ci"
// - builder: the LUCI builder, e.g. "linux-rel"
// - test_suite: a name from
//   https://cs.chromium.org/chromium/src/testing/buildbot/test_suites.pyl
type Variant struct {
	// The definition of the variant.
	// Key and values must be valid StringPair keys and values, see their
	// constraints.
	Def                  map[string]string `protobuf:"bytes,1,rep,name=def,proto3" json:"def,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Variant) Reset()         { *m = Variant{} }
func (m *Variant) String() string { return proto.CompactTextString(m) }
func (*Variant) ProtoMessage()    {}
func (*Variant) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad57c4fb5ae8d219, []int{0}
}

func (m *Variant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Variant.Unmarshal(m, b)
}
func (m *Variant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Variant.Marshal(b, m, deterministic)
}
func (m *Variant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Variant.Merge(m, src)
}
func (m *Variant) XXX_Size() int {
	return xxx_messageInfo_Variant.Size(m)
}
func (m *Variant) XXX_DiscardUnknown() {
	xxx_messageInfo_Variant.DiscardUnknown(m)
}

var xxx_messageInfo_Variant proto.InternalMessageInfo

func (m *Variant) GetDef() map[string]string {
	if m != nil {
		return m.Def
	}
	return nil
}

// A string key-value pair. Typically used for tagging, see Invocation.tags
type StringPair struct {
	// Regex: ^[a-z][a-z0-9_]*(/[a-z][a-z0-9_]*)*$
	// Max length: 64.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Max length: 256.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringPair) Reset()         { *m = StringPair{} }
func (m *StringPair) String() string { return proto.CompactTextString(m) }
func (*StringPair) ProtoMessage()    {}
func (*StringPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad57c4fb5ae8d219, []int{1}
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

func init() {
	proto.RegisterType((*Variant)(nil), "luci.resultdb.type.Variant")
	proto.RegisterMapType((map[string]string)(nil), "luci.resultdb.type.Variant.DefEntry")
	proto.RegisterType((*StringPair)(nil), "luci.resultdb.type.StringPair")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/resultdb/proto/type/common.proto", fileDescriptor_ad57c4fb5ae8d219)
}

var fileDescriptor_ad57c4fb5ae8d219 = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4d, 0xcf, 0xd7, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c, 0xcd, 0xd5, 0xcb, 0x2f, 0x4a, 0xd7, 0xcf, 0x29, 0x4d, 0xce,
	0xd4, 0x2f, 0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0x49, 0x49, 0xd2, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7,
	0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x03, 0x8b, 0x08, 0x09,
	0x81, 0x94, 0xe9, 0xc1, 0x94, 0xe9, 0x81, 0x14, 0x28, 0x55, 0x72, 0xb1, 0x87, 0x25, 0x16, 0x65,
	0x26, 0xe6, 0x95, 0x08, 0x99, 0x71, 0x31, 0xa7, 0xa4, 0xa6, 0x49, 0x30, 0x2a, 0x30, 0x6b, 0x70,
	0x1b, 0xa9, 0xe8, 0x61, 0x2a, 0xd6, 0x83, 0xaa, 0xd4, 0x73, 0x49, 0x4d, 0x73, 0xcd, 0x2b, 0x29,
	0xaa, 0x0c, 0x02, 0x69, 0x90, 0x32, 0xe3, 0xe2, 0x80, 0x09, 0x08, 0x09, 0x70, 0x31, 0x67, 0xa7,
	0x56, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x89,
	0x39, 0xa5, 0xa9, 0x12, 0x4c, 0x60, 0x31, 0x08, 0xc7, 0x8a, 0xc9, 0x82, 0x51, 0xc9, 0x84, 0x8b,
	0x2b, 0xb8, 0xa4, 0x28, 0x33, 0x2f, 0x3d, 0x20, 0x31, 0xb3, 0x88, 0x58, 0x9d, 0x4e, 0x86, 0x51,
	0xfa, 0xc4, 0xfa, 0xde, 0x1a, 0x44, 0x14, 0x24, 0x25, 0xb1, 0x81, 0x85, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x95, 0x5e, 0x65, 0xa5, 0x37, 0x01, 0x00, 0x00,
}
