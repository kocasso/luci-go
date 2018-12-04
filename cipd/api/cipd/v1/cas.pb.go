// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/cipd/api/cipd/v1/cas.proto

package api

import prpc "go.chromium.org/luci/grpc/prpc"

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// Supported hashing algorithms used by the content-addressable storage.
//
// Literal names are important, since they are used to construct Google Storage
// paths internally.
type HashAlgo int32

const (
	HashAlgo_HASH_ALGO_UNSPECIFIED HashAlgo = 0
	HashAlgo_SHA1                  HashAlgo = 1
	HashAlgo_SHA256                HashAlgo = 2
)

var HashAlgo_name = map[int32]string{
	0: "HASH_ALGO_UNSPECIFIED",
	1: "SHA1",
	2: "SHA256",
}

var HashAlgo_value = map[string]int32{
	"HASH_ALGO_UNSPECIFIED": 0,
	"SHA1":                  1,
	"SHA256":                2,
}

func (x HashAlgo) String() string {
	return proto.EnumName(HashAlgo_name, int32(x))
}

func (HashAlgo) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_df64270f060b5e81, []int{0}
}

type UploadStatus int32

const (
	UploadStatus_UPLOAD_STATUS_UNSPECIFIED UploadStatus = 0
	UploadStatus_UPLOADING                 UploadStatus = 1
	UploadStatus_VERIFYING                 UploadStatus = 2
	UploadStatus_PUBLISHED                 UploadStatus = 3
	UploadStatus_ERRORED                   UploadStatus = 4
	UploadStatus_CANCELED                  UploadStatus = 5
)

var UploadStatus_name = map[int32]string{
	0: "UPLOAD_STATUS_UNSPECIFIED",
	1: "UPLOADING",
	2: "VERIFYING",
	3: "PUBLISHED",
	4: "ERRORED",
	5: "CANCELED",
}

var UploadStatus_value = map[string]int32{
	"UPLOAD_STATUS_UNSPECIFIED": 0,
	"UPLOADING":                 1,
	"VERIFYING":                 2,
	"PUBLISHED":                 3,
	"ERRORED":                   4,
	"CANCELED":                  5,
}

func (x UploadStatus) String() string {
	return proto.EnumName(UploadStatus_name, int32(x))
}

func (UploadStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_df64270f060b5e81, []int{1}
}

// A reference to an object in the content-addressable storage.
type ObjectRef struct {
	HashAlgo             HashAlgo `protobuf:"varint,1,opt,name=hash_algo,json=hashAlgo,proto3,enum=cipd.HashAlgo" json:"hash_algo,omitempty"`
	HexDigest            string   `protobuf:"bytes,2,opt,name=hex_digest,json=hexDigest,proto3" json:"hex_digest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectRef) Reset()         { *m = ObjectRef{} }
func (m *ObjectRef) String() string { return proto.CompactTextString(m) }
func (*ObjectRef) ProtoMessage()    {}
func (*ObjectRef) Descriptor() ([]byte, []int) {
	return fileDescriptor_df64270f060b5e81, []int{0}
}

func (m *ObjectRef) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectRef.Unmarshal(m, b)
}
func (m *ObjectRef) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectRef.Marshal(b, m, deterministic)
}
func (m *ObjectRef) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectRef.Merge(m, src)
}
func (m *ObjectRef) XXX_Size() int {
	return xxx_messageInfo_ObjectRef.Size(m)
}
func (m *ObjectRef) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectRef.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectRef proto.InternalMessageInfo

func (m *ObjectRef) GetHashAlgo() HashAlgo {
	if m != nil {
		return m.HashAlgo
	}
	return HashAlgo_HASH_ALGO_UNSPECIFIED
}

func (m *ObjectRef) GetHexDigest() string {
	if m != nil {
		return m.HexDigest
	}
	return ""
}

type GetObjectURLRequest struct {
	// A reference to the object the client wants to fetch.
	Object *ObjectRef `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	// If present, the returned URL will be served with Content-Disposition header
	// that includes the given filename. It makes browsers save the file under the
	// given name.
	DownloadFilename     string   `protobuf:"bytes,2,opt,name=download_filename,json=downloadFilename,proto3" json:"download_filename,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetObjectURLRequest) Reset()         { *m = GetObjectURLRequest{} }
func (m *GetObjectURLRequest) String() string { return proto.CompactTextString(m) }
func (*GetObjectURLRequest) ProtoMessage()    {}
func (*GetObjectURLRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df64270f060b5e81, []int{1}
}

func (m *GetObjectURLRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetObjectURLRequest.Unmarshal(m, b)
}
func (m *GetObjectURLRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetObjectURLRequest.Marshal(b, m, deterministic)
}
func (m *GetObjectURLRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetObjectURLRequest.Merge(m, src)
}
func (m *GetObjectURLRequest) XXX_Size() int {
	return xxx_messageInfo_GetObjectURLRequest.Size(m)
}
func (m *GetObjectURLRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetObjectURLRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetObjectURLRequest proto.InternalMessageInfo

func (m *GetObjectURLRequest) GetObject() *ObjectRef {
	if m != nil {
		return m.Object
	}
	return nil
}

func (m *GetObjectURLRequest) GetDownloadFilename() string {
	if m != nil {
		return m.DownloadFilename
	}
	return ""
}

type ObjectURL struct {
	// A signed HTTPS URL to the object's body.
	//
	// Fetching it doesn't require authentication. Expires after some unspecified
	// short amount of time. It is expected that callers will use it immediately.
	//
	// The URL isn't guaranteed to have any particular internal structure. Do not
	// attempt to parse it.
	SignedUrl            string   `protobuf:"bytes,1,opt,name=signed_url,json=signedUrl,proto3" json:"signed_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectURL) Reset()         { *m = ObjectURL{} }
func (m *ObjectURL) String() string { return proto.CompactTextString(m) }
func (*ObjectURL) ProtoMessage()    {}
func (*ObjectURL) Descriptor() ([]byte, []int) {
	return fileDescriptor_df64270f060b5e81, []int{2}
}

func (m *ObjectURL) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectURL.Unmarshal(m, b)
}
func (m *ObjectURL) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectURL.Marshal(b, m, deterministic)
}
func (m *ObjectURL) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectURL.Merge(m, src)
}
func (m *ObjectURL) XXX_Size() int {
	return xxx_messageInfo_ObjectURL.Size(m)
}
func (m *ObjectURL) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectURL.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectURL proto.InternalMessageInfo

func (m *ObjectURL) GetSignedUrl() string {
	if m != nil {
		return m.SignedUrl
	}
	return ""
}

type BeginUploadRequest struct {
	// A reference to the object the client wants to put in the storage, if known.
	//
	// If such object already exists, RPC will finish with ALREADY_EXISTS status
	// right away.
	//
	// If this field is missing (in case the client doesn't know the hash yet),
	// the client MUST supply hash_algo field, to let the backend know what
	// hashing algorithm it should use for calculating object's hash.
	//
	// The calculated hash will be returned back to the client as part of
	// UploadOperation ('object' field) by FinishUpload call.
	Object *ObjectRef `protobuf:"bytes,1,opt,name=object,proto3" json:"object,omitempty"`
	// An algorithm to use to derive object's name during uploads when the final
	// hash of the object is not yet known.
	//
	// Optional if 'object' is present.
	//
	// If both 'object' and 'hash_algo' are present, 'object.hash_algo' MUST match
	// 'hash_algo'.
	HashAlgo             HashAlgo `protobuf:"varint,2,opt,name=hash_algo,json=hashAlgo,proto3,enum=cipd.HashAlgo" json:"hash_algo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BeginUploadRequest) Reset()         { *m = BeginUploadRequest{} }
func (m *BeginUploadRequest) String() string { return proto.CompactTextString(m) }
func (*BeginUploadRequest) ProtoMessage()    {}
func (*BeginUploadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df64270f060b5e81, []int{3}
}

func (m *BeginUploadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BeginUploadRequest.Unmarshal(m, b)
}
func (m *BeginUploadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BeginUploadRequest.Marshal(b, m, deterministic)
}
func (m *BeginUploadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeginUploadRequest.Merge(m, src)
}
func (m *BeginUploadRequest) XXX_Size() int {
	return xxx_messageInfo_BeginUploadRequest.Size(m)
}
func (m *BeginUploadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BeginUploadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BeginUploadRequest proto.InternalMessageInfo

func (m *BeginUploadRequest) GetObject() *ObjectRef {
	if m != nil {
		return m.Object
	}
	return nil
}

func (m *BeginUploadRequest) GetHashAlgo() HashAlgo {
	if m != nil {
		return m.HashAlgo
	}
	return HashAlgo_HASH_ALGO_UNSPECIFIED
}

type FinishUploadRequest struct {
	// An identifier of an upload operation returned by BeginUpload RPC.
	UploadOperationId string `protobuf:"bytes,1,opt,name=upload_operation_id,json=uploadOperationId,proto3" json:"upload_operation_id,omitempty"`
	// If set, instructs Storage to skip the hash verification and just assume the
	// uploaded object has the given hash.
	//
	// This is used internally by the service as an optimization for cases when
	// it trusts the uploaded data (for example, when it upload it itself).
	//
	// External callers are denied usage of this field. Attempt to use it results
	// in PERMISSION_DENIED.
	ForceHash            *ObjectRef `protobuf:"bytes,2,opt,name=force_hash,json=forceHash,proto3" json:"force_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *FinishUploadRequest) Reset()         { *m = FinishUploadRequest{} }
func (m *FinishUploadRequest) String() string { return proto.CompactTextString(m) }
func (*FinishUploadRequest) ProtoMessage()    {}
func (*FinishUploadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df64270f060b5e81, []int{4}
}

func (m *FinishUploadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FinishUploadRequest.Unmarshal(m, b)
}
func (m *FinishUploadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FinishUploadRequest.Marshal(b, m, deterministic)
}
func (m *FinishUploadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FinishUploadRequest.Merge(m, src)
}
func (m *FinishUploadRequest) XXX_Size() int {
	return xxx_messageInfo_FinishUploadRequest.Size(m)
}
func (m *FinishUploadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FinishUploadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FinishUploadRequest proto.InternalMessageInfo

func (m *FinishUploadRequest) GetUploadOperationId() string {
	if m != nil {
		return m.UploadOperationId
	}
	return ""
}

func (m *FinishUploadRequest) GetForceHash() *ObjectRef {
	if m != nil {
		return m.ForceHash
	}
	return nil
}

type CancelUploadRequest struct {
	// An identifier of an upload operation returned by BeginUpload RPC.
	UploadOperationId    string   `protobuf:"bytes,1,opt,name=upload_operation_id,json=uploadOperationId,proto3" json:"upload_operation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelUploadRequest) Reset()         { *m = CancelUploadRequest{} }
func (m *CancelUploadRequest) String() string { return proto.CompactTextString(m) }
func (*CancelUploadRequest) ProtoMessage()    {}
func (*CancelUploadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_df64270f060b5e81, []int{5}
}

func (m *CancelUploadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelUploadRequest.Unmarshal(m, b)
}
func (m *CancelUploadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelUploadRequest.Marshal(b, m, deterministic)
}
func (m *CancelUploadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelUploadRequest.Merge(m, src)
}
func (m *CancelUploadRequest) XXX_Size() int {
	return xxx_messageInfo_CancelUploadRequest.Size(m)
}
func (m *CancelUploadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelUploadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CancelUploadRequest proto.InternalMessageInfo

func (m *CancelUploadRequest) GetUploadOperationId() string {
	if m != nil {
		return m.UploadOperationId
	}
	return ""
}

type UploadOperation struct {
	// An opaque string that identifies this upload operation.
	//
	// It acts as a temporary authorization token for FinishUpload RPC. Treat it
	// as a secret.
	OperationId string `protobuf:"bytes,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	// URL the client should use in Google Storage Resumable Upload protocol to
	// upload the object's body.
	//
	// No authentication is required to upload data to this URL, so it also should
	// be treated as a secret.
	UploadUrl string `protobuf:"bytes,2,opt,name=upload_url,json=uploadUrl,proto3" json:"upload_url,omitempty"`
	// Status of the upload operation.
	Status UploadStatus `protobuf:"varint,3,opt,name=status,proto3,enum=cipd.UploadStatus" json:"status,omitempty"`
	// For PUBLISHED status, the reference to the published object.
	//
	// This is in particular useful for uploads when the hash of the object is not
	// known until the upload is finished.
	Object *ObjectRef `protobuf:"bytes,4,opt,name=object,proto3" json:"object,omitempty"`
	// For ERRORED status, a human readable error message.
	ErrorMessage         string   `protobuf:"bytes,5,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadOperation) Reset()         { *m = UploadOperation{} }
func (m *UploadOperation) String() string { return proto.CompactTextString(m) }
func (*UploadOperation) ProtoMessage()    {}
func (*UploadOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_df64270f060b5e81, []int{6}
}

func (m *UploadOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadOperation.Unmarshal(m, b)
}
func (m *UploadOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadOperation.Marshal(b, m, deterministic)
}
func (m *UploadOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadOperation.Merge(m, src)
}
func (m *UploadOperation) XXX_Size() int {
	return xxx_messageInfo_UploadOperation.Size(m)
}
func (m *UploadOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadOperation.DiscardUnknown(m)
}

var xxx_messageInfo_UploadOperation proto.InternalMessageInfo

func (m *UploadOperation) GetOperationId() string {
	if m != nil {
		return m.OperationId
	}
	return ""
}

func (m *UploadOperation) GetUploadUrl() string {
	if m != nil {
		return m.UploadUrl
	}
	return ""
}

func (m *UploadOperation) GetStatus() UploadStatus {
	if m != nil {
		return m.Status
	}
	return UploadStatus_UPLOAD_STATUS_UNSPECIFIED
}

func (m *UploadOperation) GetObject() *ObjectRef {
	if m != nil {
		return m.Object
	}
	return nil
}

func (m *UploadOperation) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func init() {
	proto.RegisterEnum("cipd.HashAlgo", HashAlgo_name, HashAlgo_value)
	proto.RegisterEnum("cipd.UploadStatus", UploadStatus_name, UploadStatus_value)
	proto.RegisterType((*ObjectRef)(nil), "cipd.ObjectRef")
	proto.RegisterType((*GetObjectURLRequest)(nil), "cipd.GetObjectURLRequest")
	proto.RegisterType((*ObjectURL)(nil), "cipd.ObjectURL")
	proto.RegisterType((*BeginUploadRequest)(nil), "cipd.BeginUploadRequest")
	proto.RegisterType((*FinishUploadRequest)(nil), "cipd.FinishUploadRequest")
	proto.RegisterType((*CancelUploadRequest)(nil), "cipd.CancelUploadRequest")
	proto.RegisterType((*UploadOperation)(nil), "cipd.UploadOperation")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/cipd/api/cipd/v1/cas.proto", fileDescriptor_df64270f060b5e81)
}

var fileDescriptor_df64270f060b5e81 = []byte{
	// 610 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0xa5, 0x5d, 0xd7, 0x35, 0xb7, 0xdd, 0x96, 0xb9, 0x9a, 0xd4, 0x4d, 0x9a, 0x34, 0xc2, 0x03,
	0x53, 0x27, 0xa5, 0x5a, 0x11, 0x3c, 0x80, 0x84, 0xc8, 0xda, 0x74, 0xad, 0x54, 0xd6, 0xc9, 0x59,
	0x40, 0xf0, 0x12, 0x79, 0x89, 0x97, 0x78, 0xa4, 0x71, 0x71, 0x12, 0xd8, 0x2f, 0xf0, 0x71, 0xfc,
	0x13, 0x8a, 0x93, 0x42, 0x06, 0xa5, 0x0f, 0xf0, 0x16, 0x9f, 0x73, 0x7b, 0xce, 0xf1, 0xbd, 0xd7,
	0x85, 0x9e, 0xcf, 0x75, 0x37, 0x10, 0x7c, 0xce, 0xd2, 0xb9, 0xce, 0x85, 0xdf, 0x0b, 0x53, 0x97,
	0xf5, 0x5c, 0xb6, 0xf0, 0x7a, 0x64, 0x51, 0x7c, 0x7c, 0x39, 0xeb, 0xb9, 0x24, 0xd6, 0x17, 0x82,
	0x27, 0x1c, 0xd5, 0x32, 0x48, 0x7b, 0x0f, 0xca, 0xec, 0xe6, 0x8e, 0xba, 0x09, 0xa6, 0xb7, 0xe8,
	0x14, 0x94, 0x80, 0xc4, 0x81, 0x43, 0x42, 0x9f, 0x77, 0x2a, 0xc7, 0x95, 0x93, 0x9d, 0xfe, 0x8e,
	0x9e, 0x95, 0xe9, 0x63, 0x12, 0x07, 0x46, 0xe8, 0x73, 0xdc, 0x08, 0x8a, 0x2f, 0x74, 0x04, 0x10,
	0xd0, 0x7b, 0xc7, 0x63, 0x3e, 0x8d, 0x93, 0x4e, 0xf5, 0xb8, 0x72, 0xa2, 0x60, 0x25, 0xa0, 0xf7,
	0x43, 0x09, 0x68, 0x9f, 0xa0, 0x7d, 0x41, 0x93, 0x5c, 0xdb, 0xc6, 0x53, 0x4c, 0x3f, 0xa7, 0x34,
	0x4e, 0xd0, 0x53, 0xa8, 0x73, 0x89, 0x49, 0xfd, 0x66, 0x7f, 0x37, 0xd7, 0xff, 0x99, 0x01, 0x17,
	0x34, 0x3a, 0x85, 0x3d, 0x8f, 0x7f, 0x8d, 0x42, 0x4e, 0x3c, 0xe7, 0x96, 0x85, 0x34, 0x22, 0x73,
	0x5a, 0xb8, 0xa8, 0x4b, 0x62, 0x54, 0xe0, 0x5a, 0x77, 0x79, 0x0b, 0x1b, 0x4f, 0xb3, 0x60, 0x31,
	0xf3, 0x23, 0xea, 0x39, 0xa9, 0x08, 0xa5, 0x8d, 0x82, 0x95, 0x1c, 0xb1, 0x45, 0xa8, 0xdd, 0x01,
	0x3a, 0xa7, 0x3e, 0x8b, 0xec, 0x45, 0x26, 0xf1, 0x0f, 0xb9, 0x4a, 0x3d, 0xaa, 0xae, 0xef, 0x91,
	0x96, 0x42, 0x7b, 0xc4, 0x22, 0x16, 0x07, 0x0f, 0xcd, 0x74, 0x68, 0xa7, 0x12, 0x70, 0xf8, 0x82,
	0x0a, 0x92, 0x30, 0x1e, 0x39, 0xcc, 0x2b, 0xa2, 0xee, 0xe5, 0xd4, 0x6c, 0xc9, 0x4c, 0x3c, 0xa4,
	0x03, 0xdc, 0x72, 0xe1, 0x52, 0x27, 0x13, 0x96, 0xa6, 0x2b, 0x02, 0x2a, 0xb2, 0x24, 0x0b, 0xa1,
	0x99, 0xd0, 0x1e, 0x90, 0xc8, 0xa5, 0xe1, 0x7f, 0xd9, 0x6a, 0xdf, 0x2b, 0xb0, 0x6b, 0x3f, 0x44,
	0xd1, 0x63, 0x68, 0xad, 0xf8, 0x71, 0x93, 0x97, 0xd2, 0x1e, 0x01, 0x14, 0x36, 0x59, 0xff, 0x8b,
	0xc5, 0xc8, 0x11, 0x5b, 0x84, 0xa8, 0x0b, 0xf5, 0x38, 0x21, 0x49, 0x1a, 0x77, 0x36, 0x64, 0xf7,
	0x50, 0x7e, 0x91, 0xdc, 0xc8, 0x92, 0x0c, 0x2e, 0x2a, 0x4a, 0x53, 0xa9, 0xad, 0x9f, 0xca, 0x13,
	0xd8, 0xa6, 0x42, 0x70, 0xe1, 0xcc, 0x69, 0x1c, 0x13, 0x9f, 0x76, 0x36, 0xa5, 0x6d, 0x4b, 0x82,
	0x6f, 0x73, 0xac, 0xfb, 0x0a, 0x1a, 0xcb, 0x19, 0xa1, 0x03, 0xd8, 0x1f, 0x1b, 0xd6, 0xd8, 0x31,
	0xa6, 0x17, 0x33, 0xc7, 0xbe, 0xb4, 0xae, 0xcc, 0xc1, 0x64, 0x34, 0x31, 0x87, 0xea, 0x23, 0xd4,
	0x80, 0x9a, 0x35, 0x36, 0xce, 0xd4, 0x0a, 0x02, 0xa8, 0x5b, 0x63, 0xa3, 0xff, 0xfc, 0x85, 0x5a,
	0xed, 0xa6, 0xd0, 0x2a, 0x47, 0x44, 0x47, 0x70, 0x60, 0x5f, 0x4d, 0x67, 0xc6, 0xd0, 0xb1, 0xae,
	0x8d, 0x6b, 0xdb, 0xfa, 0x4d, 0x64, 0x1b, 0x94, 0x9c, 0x9e, 0x5c, 0x5e, 0xa8, 0x95, 0xec, 0xf8,
	0xce, 0xc4, 0x93, 0xd1, 0x87, 0xec, 0x58, 0xcd, 0x8e, 0x57, 0xf6, 0xf9, 0x74, 0x62, 0x8d, 0xcd,
	0xa1, 0xba, 0x81, 0x9a, 0xb0, 0x65, 0x62, 0x3c, 0xc3, 0xe6, 0x50, 0xad, 0xa1, 0x16, 0x34, 0x06,
	0xc6, 0xe5, 0xc0, 0x9c, 0x9a, 0x43, 0x75, 0xb3, 0xff, 0xad, 0x0a, 0x5b, 0x56, 0xc2, 0x05, 0xf1,
	0x29, 0x7a, 0x09, 0xad, 0xf2, 0x93, 0x42, 0x07, 0x79, 0x37, 0x56, 0x3c, 0xb3, 0xc3, 0x07, 0x8d,
	0xca, 0x6a, 0x5f, 0x43, 0xb3, 0xb4, 0xf5, 0xa8, 0x93, 0xf3, 0x7f, 0x3e, 0x84, 0xc3, 0xfd, 0xf2,
	0x38, 0x7e, 0xcd, 0xfd, 0x0d, 0xb4, 0xca, 0x9b, 0xbc, 0xf4, 0x5e, 0xb1, 0xdd, 0x6b, 0x14, 0xca,
	0x4b, 0xb9, 0x54, 0x58, 0xb1, 0xa8, 0x7f, 0x51, 0x38, 0xdf, 0xfc, 0xb8, 0x41, 0x16, 0xec, 0xa6,
	0x2e, 0xff, 0xbf, 0x9e, 0xfd, 0x08, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xc2, 0x37, 0x4b, 0xf2, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StorageClient is the client API for Storage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StorageClient interface {
	// Produces a signed URL that can be used to fetch an object.
	//
	// Returns NOT_FOUND status code if there's no such object.
	GetObjectURL(ctx context.Context, in *GetObjectURLRequest, opts ...grpc.CallOption) (*ObjectURL, error)
	// Initiates an upload operation.
	//
	// Once the upload is initiated the client is responsible for uploading the
	// object to the temporary location (provided via 'upload_url' which should be
	// used as an upload URL in Google Storage Resumable Upload protocol) and
	// finishing the upload with a call to FinishUpload, which will launch
	// the verification of the object's hash on the server side.
	//
	// If the client knows the hash of the object it wants to upload already, it
	// can provide it via 'object' field. In that case Storage may reply right
	// away that such object already exists by retuning ALREADY_EXISTS status
	// code.
	//
	// If the client doesn't know the hash yet (perhaps if the object's body is
	// generated on the fly), it still can open an upload operation and start
	// streaming the data. When finalizing the upload the backend will calculate
	// and return the resulting hash of the object.
	//
	// An UploadOperation returned by this method contains tokens that grant
	// direct upload access to whoever possesses them, so it should be treated as
	// a secret. See UploadOperation for more info.
	BeginUpload(ctx context.Context, in *BeginUploadRequest, opts ...grpc.CallOption) (*UploadOperation, error)
	// Finishes the pending upload operation, returning its new status.
	//
	// Clients are expected to finish Google Storage Resumable protocol first
	// before calling FinishUpload. Failure to do so will cause the upload
	// operation to end up in ERROR state.
	//
	// This call is idempotent and it is expected that clients will keep polling
	// it if they want to wait for the server to verify the hash of the uploaded
	// object.
	//
	// Returns NOT_FOUND if the provided upload operation doesn't exist.
	//
	// Errors related to the uploaded file body are communicated through 'status'
	// field of the upload operation, since they are not directly related to this
	// RPC call, but rather to the upload operation itself.
	FinishUpload(ctx context.Context, in *FinishUploadRequest, opts ...grpc.CallOption) (*UploadOperation, error)
	// CancelUpload aborts the pending upload operation.
	//
	// It moves it to CANCELED state if it was in UPLOADING state and cleans up
	// any temporary garbage. Returns the most recent state of the upload
	// operation (whatever it may be).
	//
	// Does nothing if the operation is already canceled or failed.
	//
	// Returns:
	//   NOT_FOUND if the provided upload operation doesn't exist.
	//   FAILED_PRECONDITION if the upload operation is in PUBLISHED or VERIFYING
	//      state (i.e. finished or being finalized now).
	CancelUpload(ctx context.Context, in *CancelUploadRequest, opts ...grpc.CallOption) (*UploadOperation, error)
}
type storagePRPCClient struct {
	client *prpc.Client
}

func NewStoragePRPCClient(client *prpc.Client) StorageClient {
	return &storagePRPCClient{client}
}

func (c *storagePRPCClient) GetObjectURL(ctx context.Context, in *GetObjectURLRequest, opts ...grpc.CallOption) (*ObjectURL, error) {
	out := new(ObjectURL)
	err := c.client.Call(ctx, "cipd.Storage", "GetObjectURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storagePRPCClient) BeginUpload(ctx context.Context, in *BeginUploadRequest, opts ...grpc.CallOption) (*UploadOperation, error) {
	out := new(UploadOperation)
	err := c.client.Call(ctx, "cipd.Storage", "BeginUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storagePRPCClient) FinishUpload(ctx context.Context, in *FinishUploadRequest, opts ...grpc.CallOption) (*UploadOperation, error) {
	out := new(UploadOperation)
	err := c.client.Call(ctx, "cipd.Storage", "FinishUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storagePRPCClient) CancelUpload(ctx context.Context, in *CancelUploadRequest, opts ...grpc.CallOption) (*UploadOperation, error) {
	out := new(UploadOperation)
	err := c.client.Call(ctx, "cipd.Storage", "CancelUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type storageClient struct {
	cc *grpc.ClientConn
}

func NewStorageClient(cc *grpc.ClientConn) StorageClient {
	return &storageClient{cc}
}

func (c *storageClient) GetObjectURL(ctx context.Context, in *GetObjectURLRequest, opts ...grpc.CallOption) (*ObjectURL, error) {
	out := new(ObjectURL)
	err := c.cc.Invoke(ctx, "/cipd.Storage/GetObjectURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) BeginUpload(ctx context.Context, in *BeginUploadRequest, opts ...grpc.CallOption) (*UploadOperation, error) {
	out := new(UploadOperation)
	err := c.cc.Invoke(ctx, "/cipd.Storage/BeginUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) FinishUpload(ctx context.Context, in *FinishUploadRequest, opts ...grpc.CallOption) (*UploadOperation, error) {
	out := new(UploadOperation)
	err := c.cc.Invoke(ctx, "/cipd.Storage/FinishUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageClient) CancelUpload(ctx context.Context, in *CancelUploadRequest, opts ...grpc.CallOption) (*UploadOperation, error) {
	out := new(UploadOperation)
	err := c.cc.Invoke(ctx, "/cipd.Storage/CancelUpload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageServer is the server API for Storage service.
type StorageServer interface {
	// Produces a signed URL that can be used to fetch an object.
	//
	// Returns NOT_FOUND status code if there's no such object.
	GetObjectURL(context.Context, *GetObjectURLRequest) (*ObjectURL, error)
	// Initiates an upload operation.
	//
	// Once the upload is initiated the client is responsible for uploading the
	// object to the temporary location (provided via 'upload_url' which should be
	// used as an upload URL in Google Storage Resumable Upload protocol) and
	// finishing the upload with a call to FinishUpload, which will launch
	// the verification of the object's hash on the server side.
	//
	// If the client knows the hash of the object it wants to upload already, it
	// can provide it via 'object' field. In that case Storage may reply right
	// away that such object already exists by retuning ALREADY_EXISTS status
	// code.
	//
	// If the client doesn't know the hash yet (perhaps if the object's body is
	// generated on the fly), it still can open an upload operation and start
	// streaming the data. When finalizing the upload the backend will calculate
	// and return the resulting hash of the object.
	//
	// An UploadOperation returned by this method contains tokens that grant
	// direct upload access to whoever possesses them, so it should be treated as
	// a secret. See UploadOperation for more info.
	BeginUpload(context.Context, *BeginUploadRequest) (*UploadOperation, error)
	// Finishes the pending upload operation, returning its new status.
	//
	// Clients are expected to finish Google Storage Resumable protocol first
	// before calling FinishUpload. Failure to do so will cause the upload
	// operation to end up in ERROR state.
	//
	// This call is idempotent and it is expected that clients will keep polling
	// it if they want to wait for the server to verify the hash of the uploaded
	// object.
	//
	// Returns NOT_FOUND if the provided upload operation doesn't exist.
	//
	// Errors related to the uploaded file body are communicated through 'status'
	// field of the upload operation, since they are not directly related to this
	// RPC call, but rather to the upload operation itself.
	FinishUpload(context.Context, *FinishUploadRequest) (*UploadOperation, error)
	// CancelUpload aborts the pending upload operation.
	//
	// It moves it to CANCELED state if it was in UPLOADING state and cleans up
	// any temporary garbage. Returns the most recent state of the upload
	// operation (whatever it may be).
	//
	// Does nothing if the operation is already canceled or failed.
	//
	// Returns:
	//   NOT_FOUND if the provided upload operation doesn't exist.
	//   FAILED_PRECONDITION if the upload operation is in PUBLISHED or VERIFYING
	//      state (i.e. finished or being finalized now).
	CancelUpload(context.Context, *CancelUploadRequest) (*UploadOperation, error)
}

func RegisterStorageServer(s prpc.Registrar, srv StorageServer) {
	s.RegisterService(&_Storage_serviceDesc, srv)
}

func _Storage_GetObjectURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetObjectURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).GetObjectURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cipd.Storage/GetObjectURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).GetObjectURL(ctx, req.(*GetObjectURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_BeginUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BeginUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).BeginUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cipd.Storage/BeginUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).BeginUpload(ctx, req.(*BeginUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_FinishUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).FinishUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cipd.Storage/FinishUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).FinishUpload(ctx, req.(*FinishUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Storage_CancelUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelUploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageServer).CancelUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cipd.Storage/CancelUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageServer).CancelUpload(ctx, req.(*CancelUploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Storage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cipd.Storage",
	HandlerType: (*StorageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetObjectURL",
			Handler:    _Storage_GetObjectURL_Handler,
		},
		{
			MethodName: "BeginUpload",
			Handler:    _Storage_BeginUpload_Handler,
		},
		{
			MethodName: "FinishUpload",
			Handler:    _Storage_FinishUpload_Handler,
		},
		{
			MethodName: "CancelUpload",
			Handler:    _Storage_CancelUpload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/cipd/api/cipd/v1/cas.proto",
}
