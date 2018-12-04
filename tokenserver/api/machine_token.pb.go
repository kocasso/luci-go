// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/tokenserver/api/machine_token.proto

package tokenserver

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

// The kinds of machine tokens the token server can mint.
//
// Passed to MintMachineToken and InspectMachineToken.
//
// Reserved: 1.
type MachineTokenType int32

const (
	MachineTokenType_UNKNOWN_TYPE       MachineTokenType = 0
	MachineTokenType_LUCI_MACHINE_TOKEN MachineTokenType = 2
)

var MachineTokenType_name = map[int32]string{
	0: "UNKNOWN_TYPE",
	2: "LUCI_MACHINE_TOKEN",
}

var MachineTokenType_value = map[string]int32{
	"UNKNOWN_TYPE":       0,
	"LUCI_MACHINE_TOKEN": 2,
}

func (x MachineTokenType) String() string {
	return proto.EnumName(MachineTokenType_name, int32(x))
}

func (MachineTokenType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_dff8dcc9d4458b55, []int{0}
}

// MachineTokenBody describes internal structure of the machine token.
//
// The token will be put in HTTP headers and its body shouldn't be too large.
// For that reason we use unix timestamps instead of google.protobuf.Timestamp
// (no need for microsecond precision), and assume certificate serial numbers
// are smallish uint64 integers (not random blobs).
type MachineTokenBody struct {
	// Machine identity this token conveys (machine FQDN).
	//
	// It is extracted from a Common Name of a certificate used as a basis for
	// the token.
	MachineFqdn string `protobuf:"bytes,1,opt,name=machine_fqdn,json=machineFqdn,proto3" json:"machine_fqdn,omitempty"`
	// Service account email that signed this token.
	//
	// When verifying the token backends will check that the issuer is in
	// "auth-token-servers" group.
	IssuedBy string `protobuf:"bytes,2,opt,name=issued_by,json=issuedBy,proto3" json:"issued_by,omitempty"`
	// Unix timestamp in seconds when this token was issued. Required.
	IssuedAt uint64 `protobuf:"varint,3,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	// Number of seconds the token is considered valid.
	//
	// Usually 3600. Set by the token server. Required.
	Lifetime uint64 `protobuf:"varint,4,opt,name=lifetime,proto3" json:"lifetime,omitempty"`
	// Id of a CA that issued machine certificate used to make this token.
	//
	// These IDs are defined in token server config (via unique_id field).
	CaId int64 `protobuf:"varint,5,opt,name=ca_id,json=caId,proto3" json:"ca_id,omitempty"`
	// Serial number of the machine certificate used to make this token.
	//
	// ca_id and cert_sn together uniquely identify the certificate, and can be
	// used to check for certificate revocation (by asking token server whether
	// the given certificate is in CRL). Revocation checks are optional, most
	// callers can rely on expiration checks only.
	CertSn               uint64   `protobuf:"varint,6,opt,name=cert_sn,json=certSn,proto3" json:"cert_sn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MachineTokenBody) Reset()         { *m = MachineTokenBody{} }
func (m *MachineTokenBody) String() string { return proto.CompactTextString(m) }
func (*MachineTokenBody) ProtoMessage()    {}
func (*MachineTokenBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_dff8dcc9d4458b55, []int{0}
}

func (m *MachineTokenBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MachineTokenBody.Unmarshal(m, b)
}
func (m *MachineTokenBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MachineTokenBody.Marshal(b, m, deterministic)
}
func (m *MachineTokenBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MachineTokenBody.Merge(m, src)
}
func (m *MachineTokenBody) XXX_Size() int {
	return xxx_messageInfo_MachineTokenBody.Size(m)
}
func (m *MachineTokenBody) XXX_DiscardUnknown() {
	xxx_messageInfo_MachineTokenBody.DiscardUnknown(m)
}

var xxx_messageInfo_MachineTokenBody proto.InternalMessageInfo

func (m *MachineTokenBody) GetMachineFqdn() string {
	if m != nil {
		return m.MachineFqdn
	}
	return ""
}

func (m *MachineTokenBody) GetIssuedBy() string {
	if m != nil {
		return m.IssuedBy
	}
	return ""
}

func (m *MachineTokenBody) GetIssuedAt() uint64 {
	if m != nil {
		return m.IssuedAt
	}
	return 0
}

func (m *MachineTokenBody) GetLifetime() uint64 {
	if m != nil {
		return m.Lifetime
	}
	return 0
}

func (m *MachineTokenBody) GetCaId() int64 {
	if m != nil {
		return m.CaId
	}
	return 0
}

func (m *MachineTokenBody) GetCertSn() uint64 {
	if m != nil {
		return m.CertSn
	}
	return 0
}

// MachineTokenEnvelope is what is actually being serialized and represented
// as a machine token (after being encoded using base64 standard raw encoding).
//
// Resulting token (including base64 encoding) is usually ~500 bytes long.
type MachineTokenEnvelope struct {
	TokenBody            []byte   `protobuf:"bytes,1,opt,name=token_body,json=tokenBody,proto3" json:"token_body,omitempty"`
	KeyId                string   `protobuf:"bytes,2,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	RsaSha256            []byte   `protobuf:"bytes,3,opt,name=rsa_sha256,json=rsaSha256,proto3" json:"rsa_sha256,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MachineTokenEnvelope) Reset()         { *m = MachineTokenEnvelope{} }
func (m *MachineTokenEnvelope) String() string { return proto.CompactTextString(m) }
func (*MachineTokenEnvelope) ProtoMessage()    {}
func (*MachineTokenEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_dff8dcc9d4458b55, []int{1}
}

func (m *MachineTokenEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MachineTokenEnvelope.Unmarshal(m, b)
}
func (m *MachineTokenEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MachineTokenEnvelope.Marshal(b, m, deterministic)
}
func (m *MachineTokenEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MachineTokenEnvelope.Merge(m, src)
}
func (m *MachineTokenEnvelope) XXX_Size() int {
	return xxx_messageInfo_MachineTokenEnvelope.Size(m)
}
func (m *MachineTokenEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_MachineTokenEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_MachineTokenEnvelope proto.InternalMessageInfo

func (m *MachineTokenEnvelope) GetTokenBody() []byte {
	if m != nil {
		return m.TokenBody
	}
	return nil
}

func (m *MachineTokenEnvelope) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *MachineTokenEnvelope) GetRsaSha256() []byte {
	if m != nil {
		return m.RsaSha256
	}
	return nil
}

func init() {
	proto.RegisterEnum("tokenserver.MachineTokenType", MachineTokenType_name, MachineTokenType_value)
	proto.RegisterType((*MachineTokenBody)(nil), "tokenserver.MachineTokenBody")
	proto.RegisterType((*MachineTokenEnvelope)(nil), "tokenserver.MachineTokenEnvelope")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/tokenserver/api/machine_token.proto", fileDescriptor_dff8dcc9d4458b55)
}

var fileDescriptor_dff8dcc9d4458b55 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x4d, 0x4f, 0xf2, 0x40,
	0x14, 0x85, 0xdf, 0xf2, 0xd1, 0x97, 0x0e, 0x5d, 0x34, 0xe3, 0x57, 0xa3, 0x31, 0x41, 0x56, 0xc4,
	0x05, 0x4d, 0x34, 0x1a, 0x17, 0x6e, 0x80, 0xd4, 0xd8, 0x20, 0xc5, 0x94, 0x12, 0xe3, 0x6a, 0x32,
	0x74, 0x06, 0x98, 0x94, 0x76, 0xca, 0xb4, 0x90, 0xcc, 0x4f, 0xf3, 0xdf, 0x99, 0x4e, 0xc5, 0xe0,
	0xf2, 0x3e, 0xe7, 0xdc, 0xdc, 0x73, 0x0f, 0x78, 0x5a, 0xf1, 0x7e, 0xb4, 0x16, 0x3c, 0x61, 0xbb,
	0xa4, 0xcf, 0xc5, 0xca, 0xd9, 0xec, 0x22, 0xe6, 0x14, 0x3c, 0xa6, 0x69, 0x4e, 0xc5, 0x9e, 0x0a,
	0x07, 0x67, 0xcc, 0x49, 0x70, 0xb4, 0x66, 0x29, 0x45, 0x8a, 0xf7, 0x33, 0xc1, 0x0b, 0x0e, 0xdb,
	0x47, 0xa6, 0xee, 0x97, 0x06, 0xac, 0x49, 0x65, 0x0a, 0x4b, 0x3c, 0xe4, 0x44, 0xc2, 0x1b, 0x60,
	0x1e, 0x16, 0x97, 0x5b, 0x92, 0xda, 0x5a, 0x47, 0xeb, 0x19, 0x41, 0xfb, 0x87, 0xbd, 0x6c, 0x49,
	0x0a, 0xaf, 0x80, 0xc1, 0xf2, 0x7c, 0x47, 0x09, 0x5a, 0x48, 0xbb, 0xa6, 0xf4, 0x56, 0x05, 0x86,
	0xf2, 0x48, 0xc4, 0x85, 0x5d, 0xef, 0x68, 0xbd, 0xc6, 0x41, 0x1c, 0x14, 0xf0, 0x12, 0xb4, 0x36,
	0x6c, 0x49, 0x0b, 0x96, 0x50, 0xbb, 0x51, 0x69, 0x87, 0x19, 0x9e, 0x80, 0x66, 0x84, 0x11, 0x23,
	0x76, 0xb3, 0xa3, 0xf5, 0xea, 0x41, 0x23, 0xc2, 0x1e, 0x81, 0x17, 0xe0, 0x7f, 0x44, 0x45, 0x81,
	0xf2, 0xd4, 0xd6, 0x95, 0x5f, 0x2f, 0xc7, 0x59, 0xda, 0x8d, 0xc1, 0xe9, 0x71, 0x74, 0x37, 0xdd,
	0xd3, 0x0d, 0xcf, 0x28, 0xbc, 0x06, 0x40, 0xbd, 0x88, 0x16, 0x9c, 0x48, 0x15, 0xde, 0x0c, 0x8c,
	0xe2, 0xf7, 0xbb, 0x33, 0xa0, 0xc7, 0x54, 0x96, 0x57, 0xaa, 0xdc, 0xcd, 0x98, 0x4a, 0x8f, 0x94,
	0x5b, 0x22, 0xc7, 0x28, 0x5f, 0xe3, 0xbb, 0x87, 0x47, 0x95, 0xda, 0x0c, 0x0c, 0x91, 0xe3, 0x99,
	0x02, 0xb7, 0xcf, 0x7f, 0x7b, 0x0a, 0x65, 0x46, 0xa1, 0x05, 0xcc, 0xb9, 0x3f, 0xf6, 0xa7, 0x1f,
	0x3e, 0x0a, 0x3f, 0xdf, 0x5d, 0xeb, 0x1f, 0x3c, 0x07, 0xf0, 0x6d, 0x3e, 0xf2, 0xd0, 0x64, 0x30,
	0x7a, 0xf5, 0x7c, 0x17, 0x85, 0xd3, 0xb1, 0xeb, 0x5b, 0xb5, 0x85, 0xae, 0xaa, 0xbf, 0xff, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0x6c, 0x8d, 0xf8, 0x14, 0xb6, 0x01, 0x00, 0x00,
}
