// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/common/proto/gerrit/gerrit.proto

package gerrit

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

// Specifies what extra information to include in the response.
//
// Source of truth:
// https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#query-options
type QueryOption int32

const (
	QueryOption_OPTION_UNSPECIFIED QueryOption = 0
	// A summary of each label required for submit, and approvers that have
	// granted (or rejected) with that label.
	QueryOption_LABELS QueryOption = 1
	// Detailed label information, including numeric values of all existing
	// approvals, recognized label values, values permitted to be set by the
	// current user, all reviewers by state, and reviewers that may be removed by
	// the current user.
	QueryOption_DETAILED_LABELS QueryOption = 2
	// Describe the current revision (patch set) of the change, including the
	// commit SHA-1 and URLs to fetch from.
	QueryOption_CURRENT_REVISION QueryOption = 4
	// Describe all revisions, not just current.
	QueryOption_ALL_REVISIONS QueryOption = 8
	// Parse and output all header fields from the commit object, including
	// message. Only valid when the CURRENT_REVISION or ALL_REVISIONS option is
	// selected.
	QueryOption_CURRENT_COMMIT QueryOption = 16
	// Parse and output all header fields from the output revisions. If only
	// CURRENT_REVISION was requested then only the current revision’s commit data
	// will be output.
	QueryOption_ALL_COMMITS QueryOption = 32
	// List files modified by the commit and magic files, including basic line
	// counts inserted/deleted per file. Only valid when the CURRENT_REVISION or
	// ALL_REVISIONS option is selected.
	QueryOption_CURRENT_FILES QueryOption = 64
	// List files modified by the commit and magic files, including basic line
	// counts inserted/deleted per file. If only the CURRENT_REVISION was
	// requested then only that commit’s modified files will be output.
	QueryOption_ALL_FILES QueryOption = 128
	// Include _account_id, email and username fields when referencing accounts.
	QueryOption_DETAILED_ACCOUNTS QueryOption = 256
	// Include updates to reviewers set as ReviewerUpdateInfo entities.
	QueryOption_REVIEWER_UPDATES QueryOption = 512
	// Include messages associated with the change.
	QueryOption_MESSAGES QueryOption = 1024
	// Include information on available actions for the change and its current
	// revision. Ignored if the caller is not authenticated.
	QueryOption_CURRENT_ACTIONS QueryOption = 2048
	// Include information on available change actions for the change. Ignored if
	// the caller is not authenticated.
	QueryOption_CHANGE_ACTIONS QueryOption = 4096
	// Include the reviewed field if all of the following are true:
	// - the change is open
	// - the caller is authenticated
	// - the caller has commented on the change more recently than the last update
	//   from the change owner, i.e. this change would show up in the results of
	//   reviewedby:self.
	QueryOption_REVIEWED QueryOption = 8192
	// Skip the mergeable field in ChangeInfo. For fast moving projects, this
	// field must be recomputed often, which is slow for projects with big trees.
	QueryOption_SKIP_MERGEABLE QueryOption = 16384
	// Include the submittable field in ChangeInfo, which can be used to tell if
	// the change is reviewed and ready for submit.
	QueryOption_SUBMITTABLE QueryOption = 32768
	// Include the web_links field in CommitInfo, therefore only valid in
	// combination with CURRENT_COMMIT or ALL_COMMITS.
	QueryOption_WEB_LINKS QueryOption = 65536
	// Include potential problems with the change.
	QueryOption_CHECK QueryOption = 131072
	// Include the full commit message with Gerrit-specific commit footers in the
	// RevisionInfo.
	QueryOption_COMMIT_FOOTERS QueryOption = 262144
	// Include push certificate information in the RevisionInfo. Ignored if signed
	// push is not enabled on the server.
	QueryOption_PUSH_CERTIFICATES QueryOption = 524288
	// Include references to external tracking systems as TrackingIdInfo.
	QueryOption_TRACKING_IDS QueryOption = 1048576
	// Include the commands field in the FetchInfo for revisions. Only valid when
	// the CURRENT_REVISION or ALL_REVISIONS option is selected.
	QueryOption_DOWNLOAD_COMMANDS QueryOption = 2097152
)

var QueryOption_name = map[int32]string{
	0:       "OPTION_UNSPECIFIED",
	1:       "LABELS",
	2:       "DETAILED_LABELS",
	4:       "CURRENT_REVISION",
	8:       "ALL_REVISIONS",
	16:      "CURRENT_COMMIT",
	32:      "ALL_COMMITS",
	64:      "CURRENT_FILES",
	128:     "ALL_FILES",
	256:     "DETAILED_ACCOUNTS",
	512:     "REVIEWER_UPDATES",
	1024:    "MESSAGES",
	2048:    "CURRENT_ACTIONS",
	4096:    "CHANGE_ACTIONS",
	8192:    "REVIEWED",
	16384:   "SKIP_MERGEABLE",
	32768:   "SUBMITTABLE",
	65536:   "WEB_LINKS",
	131072:  "CHECK",
	262144:  "COMMIT_FOOTERS",
	524288:  "PUSH_CERTIFICATES",
	1048576: "TRACKING_IDS",
	2097152: "DOWNLOAD_COMMANDS",
}

var QueryOption_value = map[string]int32{
	"OPTION_UNSPECIFIED": 0,
	"LABELS":             1,
	"DETAILED_LABELS":    2,
	"CURRENT_REVISION":   4,
	"ALL_REVISIONS":      8,
	"CURRENT_COMMIT":     16,
	"ALL_COMMITS":        32,
	"CURRENT_FILES":      64,
	"ALL_FILES":          128,
	"DETAILED_ACCOUNTS":  256,
	"REVIEWER_UPDATES":   512,
	"MESSAGES":           1024,
	"CURRENT_ACTIONS":    2048,
	"CHANGE_ACTIONS":     4096,
	"REVIEWED":           8192,
	"SKIP_MERGEABLE":     16384,
	"SUBMITTABLE":        32768,
	"WEB_LINKS":          65536,
	"CHECK":              131072,
	"COMMIT_FOOTERS":     262144,
	"PUSH_CERTIFICATES":  524288,
	"TRACKING_IDS":       1048576,
	"DOWNLOAD_COMMANDS":  2097152,
}

func (x QueryOption) String() string {
	return proto.EnumName(QueryOption_name, int32(x))
}

func (QueryOption) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3c6e096860f6adc3, []int{0}
}

type GetChangeRequest struct {
	// Change number.
	Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	// What to include in the response.
	Options              []QueryOption `protobuf:"varint,2,rep,packed,name=options,proto3,enum=gerrit.QueryOption" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetChangeRequest) Reset()         { *m = GetChangeRequest{} }
func (m *GetChangeRequest) String() string { return proto.CompactTextString(m) }
func (*GetChangeRequest) ProtoMessage()    {}
func (*GetChangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c6e096860f6adc3, []int{0}
}

func (m *GetChangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetChangeRequest.Unmarshal(m, b)
}
func (m *GetChangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetChangeRequest.Marshal(b, m, deterministic)
}
func (m *GetChangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetChangeRequest.Merge(m, src)
}
func (m *GetChangeRequest) XXX_Size() int {
	return xxx_messageInfo_GetChangeRequest.Size(m)
}
func (m *GetChangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetChangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetChangeRequest proto.InternalMessageInfo

func (m *GetChangeRequest) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *GetChangeRequest) GetOptions() []QueryOption {
	if m != nil {
		return m.Options
	}
	return nil
}

// Information about an account.
// Source of truth: https://gerrit-review.googlesource.com/Documentation/rest-api-accounts.html#account-info
type AccountInfo struct {
	// The full name of the user.
	// Only set if detailed account information is requested.
	// See option DETAILED_ACCOUNTS for change queries
	// and option DETAILS for account queries.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The email address the user prefers to be contacted through.
	// Only set if detailed account information is requested.
	// See option DETAILED_ACCOUNTS for change queries
	// and options DETAILS and ALL_EMAILS for account queries.
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	// A list of the secondary email addresses of the user.
	// Only set for account queries when the ALL_EMAILS option or the suggest
	// parameter is set. Secondary emails are only included if the calling user
	// has the Modify Account, and hence is allowed to see secondary emails of
	// other users.
	SecondaryEmails []string `protobuf:"bytes,3,rep,name=secondary_emails,json=secondaryEmails,proto3" json:"secondary_emails,omitempty"`
	// The username of the user.
	// Only set if detailed account information is requested.
	// See option DETAILED_ACCOUNTS for change queries
	// and option DETAILS for account queries.
	Username             string   `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountInfo) Reset()         { *m = AccountInfo{} }
func (m *AccountInfo) String() string { return proto.CompactTextString(m) }
func (*AccountInfo) ProtoMessage()    {}
func (*AccountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c6e096860f6adc3, []int{1}
}

func (m *AccountInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountInfo.Unmarshal(m, b)
}
func (m *AccountInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountInfo.Marshal(b, m, deterministic)
}
func (m *AccountInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountInfo.Merge(m, src)
}
func (m *AccountInfo) XXX_Size() int {
	return xxx_messageInfo_AccountInfo.Size(m)
}
func (m *AccountInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AccountInfo proto.InternalMessageInfo

func (m *AccountInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AccountInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *AccountInfo) GetSecondaryEmails() []string {
	if m != nil {
		return m.SecondaryEmails
	}
	return nil
}

func (m *AccountInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

// Information about a change.
// Source of truth: https://gerrit-review.googlesource.com/Documentation/rest-api-changes.html#change-info
type ChangeInfo struct {
	// The change number.
	Number int64 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	// The owner of the change.
	Owner *AccountInfo `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// The project of this change. For example, "chromium/src".
	Project              string   `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeInfo) Reset()         { *m = ChangeInfo{} }
func (m *ChangeInfo) String() string { return proto.CompactTextString(m) }
func (*ChangeInfo) ProtoMessage()    {}
func (*ChangeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3c6e096860f6adc3, []int{2}
}

func (m *ChangeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeInfo.Unmarshal(m, b)
}
func (m *ChangeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeInfo.Marshal(b, m, deterministic)
}
func (m *ChangeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeInfo.Merge(m, src)
}
func (m *ChangeInfo) XXX_Size() int {
	return xxx_messageInfo_ChangeInfo.Size(m)
}
func (m *ChangeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeInfo proto.InternalMessageInfo

func (m *ChangeInfo) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *ChangeInfo) GetOwner() *AccountInfo {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *ChangeInfo) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func init() {
	proto.RegisterEnum("gerrit.QueryOption", QueryOption_name, QueryOption_value)
	proto.RegisterType((*GetChangeRequest)(nil), "gerrit.GetChangeRequest")
	proto.RegisterType((*AccountInfo)(nil), "gerrit.AccountInfo")
	proto.RegisterType((*ChangeInfo)(nil), "gerrit.ChangeInfo")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/common/proto/gerrit/gerrit.proto", fileDescriptor_3c6e096860f6adc3)
}

var fileDescriptor_3c6e096860f6adc3 = []byte{
	// 596 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0x2d, 0x1f, 0x21, 0x30, 0x34, 0xb0, 0x4c, 0x68, 0x6a, 0xe5, 0x84, 0x72, 0x4a, 0x2a, 0x15,
	0xa4, 0x54, 0x3d, 0xf4, 0x16, 0xb3, 0x5e, 0xc8, 0x0a, 0x63, 0xd3, 0xb5, 0x69, 0xd4, 0x93, 0x45,
	0x5c, 0x97, 0x50, 0xc5, 0xde, 0xd4, 0xd8, 0xaa, 0x72, 0xa9, 0xfc, 0x1f, 0xfb, 0x87, 0x2a, 0xaf,
	0x31, 0x44, 0x95, 0x7a, 0xb2, 0xdf, 0xdb, 0xb7, 0x6f, 0xde, 0xac, 0x66, 0xe0, 0xe3, 0x5a, 0x0e,
	0xfd, 0x87, 0x58, 0x86, 0x9b, 0x34, 0x1c, 0xca, 0x78, 0x3d, 0x7a, 0x4c, 0xfd, 0xcd, 0xc8, 0x97,
	0x61, 0x28, 0xa3, 0xd1, 0x53, 0x2c, 0x13, 0x39, 0x5a, 0x07, 0x71, 0xbc, 0x49, 0x76, 0x9f, 0xa1,
	0xe2, 0xb0, 0x51, 0xa0, 0x8b, 0xaf, 0x40, 0xa6, 0x41, 0x42, 0x1f, 0x56, 0xd1, 0x3a, 0x10, 0xc1,
	0xcf, 0x34, 0xd8, 0x26, 0x78, 0x06, 0x8d, 0x28, 0x0d, 0xef, 0x83, 0x58, 0xab, 0x0c, 0x2a, 0x97,
	0x35, 0xb1, 0x43, 0xf8, 0x1e, 0x8e, 0xe5, 0x53, 0xb2, 0x91, 0xd1, 0x56, 0xab, 0x0e, 0x6a, 0x97,
	0x9d, 0xeb, 0xd3, 0xe1, 0xce, 0xf3, 0x73, 0x1a, 0xc4, 0xcf, 0xb6, 0x3a, 0x13, 0xa5, 0xe6, 0xe2,
	0x37, 0xb4, 0x75, 0xdf, 0x97, 0x69, 0x94, 0xf0, 0xe8, 0xbb, 0x44, 0x84, 0x7a, 0xb4, 0x0a, 0x03,
	0xe5, 0xd9, 0x12, 0xea, 0x1f, 0xfb, 0x70, 0x14, 0x84, 0xab, 0xcd, 0xa3, 0x56, 0x55, 0x64, 0x01,
	0xf0, 0x0a, 0xc8, 0x36, 0xf0, 0x65, 0xf4, 0x6d, 0x15, 0x3f, 0x7b, 0x8a, 0xda, 0x6a, 0xb5, 0x41,
	0xed, 0xb2, 0x25, 0xba, 0x7b, 0x9e, 0x29, 0x1a, 0xcf, 0xa1, 0x99, 0x6e, 0x83, 0x58, 0x19, 0xd7,
	0x95, 0xc7, 0x1e, 0x5f, 0x6c, 0x00, 0x8a, 0xbe, 0x54, 0xf9, 0xff, 0x35, 0x75, 0x05, 0x47, 0xf2,
	0x57, 0x14, 0xc4, 0x2a, 0x42, 0xfb, 0xd0, 0xd2, 0x8b, 0xe8, 0xa2, 0x50, 0xa0, 0x06, 0xc7, 0x4f,
	0xb1, 0xfc, 0x11, 0xf8, 0x89, 0x56, 0x53, 0xb5, 0x4a, 0xf8, 0xee, 0x4f, 0x0d, 0xda, 0x2f, 0xde,
	0x00, 0xcf, 0x00, 0xed, 0x85, 0xcb, 0x6d, 0xcb, 0x5b, 0x5a, 0xce, 0x82, 0x51, 0x3e, 0xe1, 0xcc,
	0x20, 0xaf, 0x10, 0xa0, 0x61, 0xea, 0x63, 0x66, 0x3a, 0xa4, 0x82, 0xa7, 0xd0, 0x35, 0x98, 0xab,
	0x73, 0x93, 0x19, 0xde, 0x8e, 0xac, 0x62, 0x1f, 0x08, 0x5d, 0x0a, 0xc1, 0x2c, 0xd7, 0x13, 0xec,
	0x0b, 0x77, 0xb8, 0x6d, 0x91, 0x3a, 0xf6, 0xe0, 0x44, 0x37, 0xcd, 0x3d, 0xe3, 0x90, 0x26, 0x22,
	0x74, 0x4a, 0x21, 0xb5, 0xe7, 0x73, 0xee, 0x12, 0x82, 0x5d, 0x68, 0xe7, 0xb2, 0x02, 0x3b, 0x64,
	0x90, 0xdf, 0x2b, 0x45, 0x13, 0x6e, 0x32, 0x87, 0xdc, 0x60, 0x07, 0x5a, 0xb9, 0xa6, 0x80, 0x59,
	0x05, 0xcf, 0xa0, 0xb7, 0x4f, 0xa1, 0x53, 0x6a, 0x2f, 0x2d, 0xd7, 0x21, 0x59, 0x15, 0xdf, 0x00,
	0xc9, 0xcb, 0xb1, 0x3b, 0x26, 0xbc, 0xe5, 0xc2, 0xd0, 0xdd, 0x5c, 0x5e, 0xc7, 0x13, 0x68, 0xce,
	0x99, 0xe3, 0xe8, 0xd3, 0x1c, 0x36, 0xb1, 0x0f, 0xdd, 0xb2, 0x80, 0x4e, 0x5d, 0x15, 0x2d, 0x23,
	0x78, 0x0a, 0x1d, 0x7a, 0xab, 0x5b, 0x53, 0x76, 0x20, 0x07, 0xf9, 0xcd, 0x9d, 0xa1, 0x41, 0xb2,
	0x1b, 0xec, 0x43, 0xc7, 0x99, 0xf1, 0x85, 0x37, 0x67, 0x62, 0xca, 0xf4, 0xb1, 0xc9, 0x48, 0x96,
	0x55, 0xb0, 0x07, 0x6d, 0x67, 0x39, 0x9e, 0x73, 0xd7, 0xdd, 0x51, 0x55, 0xec, 0x42, 0xeb, 0x8e,
	0x8d, 0x3d, 0x93, 0x5b, 0x33, 0x87, 0x64, 0x59, 0x1d, 0xdb, 0x70, 0x44, 0x6f, 0x19, 0x9d, 0x91,
	0x4c, 0x05, 0xe8, 0x14, 0xed, 0x7a, 0x13, 0xdb, 0x76, 0x99, 0xc8, 0x25, 0x04, 0xdf, 0x42, 0x6f,
	0xb1, 0x74, 0x6e, 0x3d, 0xca, 0x84, 0xcb, 0x27, 0x9c, 0x16, 0xe9, 0xb3, 0x01, 0x22, 0xbc, 0x76,
	0x85, 0x4e, 0x67, 0xdc, 0x9a, 0x7a, 0xdc, 0xc8, 0xb9, 0x1b, 0xd4, 0xa0, 0x67, 0xd8, 0x77, 0x96,
	0x69, 0xeb, 0x86, 0x7a, 0x3a, 0xdd, 0x52, 0x07, 0x59, 0xe5, 0x9a, 0x42, 0x63, 0xaa, 0x86, 0x01,
	0x3f, 0x41, 0x6b, 0xbf, 0x25, 0xa8, 0x95, 0x23, 0xf2, 0xef, 0xe2, 0x9c, 0x63, 0x79, 0x72, 0x98,
	0xbb, 0xfb, 0x86, 0xda, 0xb7, 0x0f, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x17, 0xaa, 0x86, 0x8f,
	0xa8, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GerritClient is the client API for Gerrit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GerritClient interface {
	// Loads a change by id.
	GetChange(ctx context.Context, in *GetChangeRequest, opts ...grpc.CallOption) (*ChangeInfo, error)
}
type gerritPRPCClient struct {
	client *prpc.Client
}

func NewGerritPRPCClient(client *prpc.Client) GerritClient {
	return &gerritPRPCClient{client}
}

func (c *gerritPRPCClient) GetChange(ctx context.Context, in *GetChangeRequest, opts ...grpc.CallOption) (*ChangeInfo, error) {
	out := new(ChangeInfo)
	err := c.client.Call(ctx, "gerrit.Gerrit", "GetChange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type gerritClient struct {
	cc *grpc.ClientConn
}

func NewGerritClient(cc *grpc.ClientConn) GerritClient {
	return &gerritClient{cc}
}

func (c *gerritClient) GetChange(ctx context.Context, in *GetChangeRequest, opts ...grpc.CallOption) (*ChangeInfo, error) {
	out := new(ChangeInfo)
	err := c.cc.Invoke(ctx, "/gerrit.Gerrit/GetChange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GerritServer is the server API for Gerrit service.
type GerritServer interface {
	// Loads a change by id.
	GetChange(context.Context, *GetChangeRequest) (*ChangeInfo, error)
}

func RegisterGerritServer(s prpc.Registrar, srv GerritServer) {
	s.RegisterService(&_Gerrit_serviceDesc, srv)
}

func _Gerrit_GetChange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GerritServer).GetChange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gerrit.Gerrit/GetChange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GerritServer).GetChange(ctx, req.(*GetChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gerrit_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gerrit.Gerrit",
	HandlerType: (*GerritServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChange",
			Handler:    _Gerrit_GetChange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "go.chromium.org/luci/common/proto/gerrit/gerrit.proto",
}
