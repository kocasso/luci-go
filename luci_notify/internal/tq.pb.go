// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/luci_notify/internal/tq.proto

package internal

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

// EmailTask represents a single email notification to be dispatched.
type EmailTask struct {
	// Recipients is a list of email addresses to send the email to.
	// TODO(nodir): make it non-repeated.
	Recipients []string `protobuf:"bytes,1,rep,name=recipients,proto3" json:"recipients,omitempty"`
	// Subject is the subject line of the email to be sent.
	Subject string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	// DEPRECATED. See body_gzip.
	Body string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	// Gzipped, HTML-formatted string containing the body of the email
	// to be sent.
	BodyGzip             []byte   `protobuf:"bytes,4,opt,name=body_gzip,json=bodyGzip,proto3" json:"body_gzip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailTask) Reset()         { *m = EmailTask{} }
func (m *EmailTask) String() string { return proto.CompactTextString(m) }
func (*EmailTask) ProtoMessage()    {}
func (*EmailTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_82c51ac555dce9eb, []int{0}
}

func (m *EmailTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailTask.Unmarshal(m, b)
}
func (m *EmailTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailTask.Marshal(b, m, deterministic)
}
func (m *EmailTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailTask.Merge(m, src)
}
func (m *EmailTask) XXX_Size() int {
	return xxx_messageInfo_EmailTask.Size(m)
}
func (m *EmailTask) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailTask.DiscardUnknown(m)
}

var xxx_messageInfo_EmailTask proto.InternalMessageInfo

func (m *EmailTask) GetRecipients() []string {
	if m != nil {
		return m.Recipients
	}
	return nil
}

func (m *EmailTask) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *EmailTask) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *EmailTask) GetBodyGzip() []byte {
	if m != nil {
		return m.BodyGzip
	}
	return nil
}

func init() {
	proto.RegisterType((*EmailTask)(nil), "internal.EmailTask")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/luci_notify/internal/tq.proto", fileDescriptor_82c51ac555dce9eb)
}

var fileDescriptor_82c51ac555dce9eb = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xce, 0x41, 0xeb, 0x82, 0x40,
	0x10, 0x05, 0x70, 0xfc, 0x2b, 0xff, 0xdc, 0xa1, 0xd3, 0x9e, 0x16, 0x82, 0x90, 0x4e, 0x9e, 0x14,
	0xea, 0x33, 0x44, 0x77, 0xe9, 0x2e, 0xba, 0x6d, 0x36, 0xa5, 0x3b, 0xdb, 0xba, 0x1b, 0xe8, 0xa7,
	0x0f, 0x17, 0x84, 0x2e, 0x33, 0xef, 0xfd, 0x4e, 0x0f, 0x8e, 0x1d, 0x15, 0xf2, 0x61, 0x69, 0x40,
	0x3f, 0x14, 0x64, 0xbb, 0xb2, 0xf7, 0x12, 0xc3, 0xa9, 0x35, 0x39, 0xbc, 0x4f, 0x25, 0x6a, 0xa7,
	0xac, 0x6e, 0xfa, 0xd2, 0xbd, 0x0b, 0x63, 0xc9, 0x11, 0x4f, 0x57, 0x3a, 0x7c, 0x80, 0x9d, 0x87,
	0x06, 0xfb, 0x6b, 0x33, 0xbe, 0xf8, 0x1e, 0xc0, 0x2a, 0x89, 0x06, 0x95, 0x76, 0xa3, 0x88, 0xb2,
	0x38, 0x67, 0xd5, 0x8f, 0x70, 0x01, 0x9b, 0xd1, 0xb7, 0x4f, 0x25, 0x9d, 0xf8, 0xcb, 0xa2, 0x9c,
	0x55, 0x6b, 0xe5, 0x1c, 0x92, 0x96, 0x6e, 0x93, 0x88, 0x03, 0x87, 0xcc, 0x77, 0xc0, 0x96, 0x5f,
	0x77, 0x33, 0x1a, 0x91, 0x64, 0x51, 0xbe, 0xad, 0xd2, 0x05, 0x2e, 0x33, 0x9a, 0xf6, 0x3f, 0x0c,
	0x39, 0x7d, 0x03, 0x00, 0x00, 0xff, 0xff, 0xad, 0xb9, 0xd8, 0xa9, 0xbe, 0x00, 0x00, 0x00,
}
