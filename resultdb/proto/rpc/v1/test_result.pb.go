// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/resultdb/proto/rpc/v1/test_result.proto

package rpcpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_type "go.chromium.org/luci/resultdb/proto/type"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// Machine-readable status of a test result.
type TestStatus int32

const (
	// Status was not specified.
	// Not to be used in actual test results; serves as a default value for an
	// unset field.
	TestStatus_STATUS_UNSPECIFIED TestStatus = 0
	// The test case has passed.
	TestStatus_PASS TestStatus = 1
	// The test case has failed.
	// Suggests that the code under test is incorrect, but it is also possible
	// that the test is incorrect or it is a flake.
	TestStatus_FAIL TestStatus = 2
	// The test case has crashed during execution.
	// The outcome is inconclusive: the code under test might or might not be
	// correct, but the test+code is incorrect.
	TestStatus_CRASH TestStatus = 3
	// The test case has started, but was aborted before finishing.
	// A common reason: timeout.
	TestStatus_ABORT TestStatus = 4
	// The test case did not execute.
	// Examples:
	// - The execution of the collection of test cases, such as a test
	//   binary, was aborted prematurely and execution of some test cases was
	//   skipped.
	// - The test harness configuration specified that the test case MUST be
	//   skipped.
	TestStatus_SKIP TestStatus = 5
)

var TestStatus_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "PASS",
	2: "FAIL",
	3: "CRASH",
	4: "ABORT",
	5: "SKIP",
}

var TestStatus_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"PASS":               1,
	"FAIL":               2,
	"CRASH":              3,
	"ABORT":              4,
	"SKIP":               5,
}

func (x TestStatus) String() string {
	return proto.EnumName(TestStatus_name, int32(x))
}

func (TestStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8253a2b79929a2cf, []int{0}
}

// A result of a functional test case.
// Often a single test case is executed multiple times and has multiple results,
// a single test suite has multiple test cases,
// and the same test suite can be executed in different variants
// (OS, GPU, compile flags, etc).
//
// This message does not specify the test id.
// It should be available in the message that embeds this message.
type TestResult struct {
	// Can be used to refer to this test result, e.g. in ResultDB.GetTestResult
	// RPC.
	// Format:
	// "invocations/{INVOCATION_ID}/tests/{URL_ESCAPED_TEST_ID}/results/{RESULT_ID}".
	// URL_ESCAPED_TEST_ID is test_id escaped with
	// https://golang.org/pkg/net/url/#PathEscape See also https://aip.dev/122.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Test id, a unique identifier of the test in a LUCI project.
	// Regex: ^[[::print::]]{1,256}$
	//
	// If two tests have a common test id prefix that ends with a
	// non-alphanumeric character, they considered a part of a group. Examples:
	// - "a/b/c"
	// - "a/b/d"
	// - "a/b/e:x"
	// - "a/b/e:y"
	// - "a/f"
	// This defines the following groups:
	// - All items belong to one group because of the common prefix "a/"
	// - Within that group, the first 4 form a sub-group because of the common
	//   prefix "a/b/"
	// - Within that group, "a/b/e:x" and "a/b/e:y" form a sub-group because of
	//   the common prefix "a/b/e:".
	// This can be used in UI.
	// LUCI does not interpret test ids in any other way.
	TestId string `protobuf:"bytes,2,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
	// Identifies a test result in a given invocation and test id.
	// Regex: ^[[:ascii:]]{1,32}$
	ResultId string `protobuf:"bytes,3,opt,name=result_id,json=resultId,proto3" json:"result_id,omitempty"`
	// Description of one specific way of running the test,
	// e.g. a specific bucket, builder and a test suite.
	Variant *_type.Variant `protobuf:"bytes,4,opt,name=variant,proto3" json:"variant,omitempty"`
	// Whether the result of test case execution is expected.
	// In a typical Chromium CL, 99%+ of test results are expected.
	// Users are typically interested only in the unexpected results.
	//
	// An unexpected result != test case failure. There are test cases that are
	// expected to fail/skip/crash. The test harness compares the actual status
	// with the expected one(s) and this field is the result of the comparison.
	Expected bool `protobuf:"varint,5,opt,name=expected,proto3" json:"expected,omitempty"`
	// Machine-readable status of the test case.
	// MUST NOT be STATUS_UNSPECIFIED.
	Status TestStatus `protobuf:"varint,6,opt,name=status,proto3,enum=luci.resultdb.rpc.v1.TestStatus" json:"status,omitempty"`
	// Human-readable explanation of the result, in HTML.
	// MUST be sanitized before rendering in the browser.
	SummaryHtml string `protobuf:"bytes,7,opt,name=summary_html,json=summaryHtml,proto3" json:"summary_html,omitempty"`
	// The point in time when the test case started to execute.
	StartTime *timestamp.Timestamp `protobuf:"bytes,8,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Duration of the test case execution.
	Duration *duration.Duration `protobuf:"bytes,9,opt,name=duration,proto3" json:"duration,omitempty"`
	// Metadata for this test result.
	// It might describe this particular execution or the test case.
	Tags []*_type.StringPair `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
	// Artifacts consumed by this test result.
	//
	// Example: building a Chrome OS image is expensive and non-deterministic, so
	// they are retained and used as input artifact to a test case.
	InputArtifacts []*Artifact `protobuf:"bytes,11,rep,name=input_artifacts,json=inputArtifacts,proto3" json:"input_artifacts,omitempty"`
	// Artifacts produced by this test result.
	// Examples: traces, logs, screenshots, memory dumps, profiler output.
	OutputArtifacts      []*Artifact `protobuf:"bytes,12,rep,name=output_artifacts,json=outputArtifacts,proto3" json:"output_artifacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TestResult) Reset()         { *m = TestResult{} }
func (m *TestResult) String() string { return proto.CompactTextString(m) }
func (*TestResult) ProtoMessage()    {}
func (*TestResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_8253a2b79929a2cf, []int{0}
}

func (m *TestResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestResult.Unmarshal(m, b)
}
func (m *TestResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestResult.Marshal(b, m, deterministic)
}
func (m *TestResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResult.Merge(m, src)
}
func (m *TestResult) XXX_Size() int {
	return xxx_messageInfo_TestResult.Size(m)
}
func (m *TestResult) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResult.DiscardUnknown(m)
}

var xxx_messageInfo_TestResult proto.InternalMessageInfo

func (m *TestResult) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TestResult) GetTestId() string {
	if m != nil {
		return m.TestId
	}
	return ""
}

func (m *TestResult) GetResultId() string {
	if m != nil {
		return m.ResultId
	}
	return ""
}

func (m *TestResult) GetVariant() *_type.Variant {
	if m != nil {
		return m.Variant
	}
	return nil
}

func (m *TestResult) GetExpected() bool {
	if m != nil {
		return m.Expected
	}
	return false
}

func (m *TestResult) GetStatus() TestStatus {
	if m != nil {
		return m.Status
	}
	return TestStatus_STATUS_UNSPECIFIED
}

func (m *TestResult) GetSummaryHtml() string {
	if m != nil {
		return m.SummaryHtml
	}
	return ""
}

func (m *TestResult) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *TestResult) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *TestResult) GetTags() []*_type.StringPair {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *TestResult) GetInputArtifacts() []*Artifact {
	if m != nil {
		return m.InputArtifacts
	}
	return nil
}

func (m *TestResult) GetOutputArtifacts() []*Artifact {
	if m != nil {
		return m.OutputArtifacts
	}
	return nil
}

// A file produced/consumed by a test case.
// See TestResult.output_artifacts for examples.
type Artifact struct {
	// A slash-separated relative path, identifies the artifact.
	// Regex: ^[[:word:]]([[:print:]]{0,254}[[:word:]])?$
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Machine-readable URL to fetch the contents of the artifact.
	// This must be a plain (curlable) HTTPS URL.
	//
	// Internally, this may have format "isolate://{host}/{ns}/{digest}" at the
	// storage level, but it is converted to an HTTP URL before serving.
	FetchUrl string `protobuf:"bytes,2,opt,name=fetch_url,json=fetchUrl,proto3" json:"fetch_url,omitempty"`
	// When fetch_url expires. If expired, re-request this Artifact.
	// If fetch_url does not expire, this is unset.
	FetchUrlExpiration *timestamp.Timestamp `protobuf:"bytes,3,opt,name=fetch_url_expiration,json=fetchUrlExpiration,proto3" json:"fetch_url_expiration,omitempty"`
	// Media type of the artifact.
	// Logs are typically "text/plain" and screenshots are typically "image/png".
	ContentType string `protobuf:"bytes,4,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// Size of the file, in bytes.
	// Can be used in UI to decide between displaying the artifact inline or only
	// showing a link if it is too large.
	Size                 int64    `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Artifact) Reset()         { *m = Artifact{} }
func (m *Artifact) String() string { return proto.CompactTextString(m) }
func (*Artifact) ProtoMessage()    {}
func (*Artifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_8253a2b79929a2cf, []int{1}
}

func (m *Artifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Artifact.Unmarshal(m, b)
}
func (m *Artifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Artifact.Marshal(b, m, deterministic)
}
func (m *Artifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Artifact.Merge(m, src)
}
func (m *Artifact) XXX_Size() int {
	return xxx_messageInfo_Artifact.Size(m)
}
func (m *Artifact) XXX_DiscardUnknown() {
	xxx_messageInfo_Artifact.DiscardUnknown(m)
}

var xxx_messageInfo_Artifact proto.InternalMessageInfo

func (m *Artifact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Artifact) GetFetchUrl() string {
	if m != nil {
		return m.FetchUrl
	}
	return ""
}

func (m *Artifact) GetFetchUrlExpiration() *timestamp.Timestamp {
	if m != nil {
		return m.FetchUrlExpiration
	}
	return nil
}

func (m *Artifact) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *Artifact) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

// Indicates the test subject (e.g. a CL) is absolved from blame
// for an unexpected result of a test variant.
// For example, the test variant fails both with and without CL, so it is not
// CL's fault.
type TestExoneration struct {
	// Can be used to refer to this test exoneration, e.g. in
	// ResultDB.GetTestExoneration RPC.
	// Format:
	// invocations/{INVOCATION_ID}/tests/{URL_ESCAPED_TEST_ID}/exonerations/{EXONERATION_ID}.
	// URL_ESCAPED_TEST_ID is test_variant.test_id escaped with
	// https://golang.org/pkg/net/url/#PathEscape See also https://aip.dev/122.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Test identifier, see TestResult.test_id.
	TestId string `protobuf:"bytes,2,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
	// Description of the variant of the test, see Variant type.
	// Unlike TestResult.extra_variant_pairs, this one must be a full definition
	// of the variant, i.e. it is not combined with Invocation.base_test_variant.
	Variant *_type.Variant `protobuf:"bytes,3,opt,name=variant,proto3" json:"variant,omitempty"`
	// Identifies an exoneration in a given invocation and test id.
	// It is server-generated.
	ExonerationId string `protobuf:"bytes,4,opt,name=exoneration_id,json=exonerationId,proto3" json:"exoneration_id,omitempty"`
	// Reasoning behind the exoneration, in HTML.
	// MUST be sanitized before rendering in the browser.
	ExplanationHtml      string   `protobuf:"bytes,5,opt,name=explanation_html,json=explanationHtml,proto3" json:"explanation_html,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestExoneration) Reset()         { *m = TestExoneration{} }
func (m *TestExoneration) String() string { return proto.CompactTextString(m) }
func (*TestExoneration) ProtoMessage()    {}
func (*TestExoneration) Descriptor() ([]byte, []int) {
	return fileDescriptor_8253a2b79929a2cf, []int{2}
}

func (m *TestExoneration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestExoneration.Unmarshal(m, b)
}
func (m *TestExoneration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestExoneration.Marshal(b, m, deterministic)
}
func (m *TestExoneration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestExoneration.Merge(m, src)
}
func (m *TestExoneration) XXX_Size() int {
	return xxx_messageInfo_TestExoneration.Size(m)
}
func (m *TestExoneration) XXX_DiscardUnknown() {
	xxx_messageInfo_TestExoneration.DiscardUnknown(m)
}

var xxx_messageInfo_TestExoneration proto.InternalMessageInfo

func (m *TestExoneration) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TestExoneration) GetTestId() string {
	if m != nil {
		return m.TestId
	}
	return ""
}

func (m *TestExoneration) GetVariant() *_type.Variant {
	if m != nil {
		return m.Variant
	}
	return nil
}

func (m *TestExoneration) GetExonerationId() string {
	if m != nil {
		return m.ExonerationId
	}
	return ""
}

func (m *TestExoneration) GetExplanationHtml() string {
	if m != nil {
		return m.ExplanationHtml
	}
	return ""
}

func init() {
	proto.RegisterEnum("luci.resultdb.rpc.v1.TestStatus", TestStatus_name, TestStatus_value)
	proto.RegisterType((*TestResult)(nil), "luci.resultdb.rpc.v1.TestResult")
	proto.RegisterType((*Artifact)(nil), "luci.resultdb.rpc.v1.Artifact")
	proto.RegisterType((*TestExoneration)(nil), "luci.resultdb.rpc.v1.TestExoneration")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/resultdb/proto/rpc/v1/test_result.proto", fileDescriptor_8253a2b79929a2cf)
}

var fileDescriptor_8253a2b79929a2cf = []byte{
	// 686 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5b, 0x6b, 0xdb, 0x3a,
	0x1c, 0x3f, 0x6e, 0x2e, 0x75, 0x94, 0x9e, 0xc6, 0x88, 0x72, 0xea, 0x93, 0x1e, 0xda, 0xd0, 0x03,
	0x23, 0x0c, 0x26, 0xd3, 0x8e, 0x32, 0x28, 0xdb, 0xc0, 0x6d, 0x53, 0x6a, 0x56, 0xba, 0x60, 0xa7,
	0x63, 0x6f, 0x46, 0xb1, 0x95, 0x44, 0xe0, 0x1b, 0xb2, 0x1c, 0xd2, 0x3d, 0xee, 0x0b, 0xee, 0x0b,
	0xec, 0xa1, 0x1f, 0x65, 0x48, 0xb2, 0x93, 0xa6, 0x2b, 0x6b, 0xdf, 0xac, 0xff, 0xef, 0x22, 0xfd,
	0x6f, 0x06, 0xef, 0xa7, 0x29, 0x0a, 0x66, 0x2c, 0x8d, 0x69, 0x11, 0xa3, 0x94, 0x4d, 0xad, 0xa8,
	0x08, 0xa8, 0xc5, 0x48, 0x5e, 0x44, 0x3c, 0x1c, 0x5b, 0x19, 0x4b, 0x79, 0x6a, 0xb1, 0x2c, 0xb0,
	0xe6, 0x47, 0x16, 0x27, 0x39, 0xf7, 0x15, 0x84, 0x24, 0x00, 0x77, 0x04, 0x1b, 0x55, 0x6c, 0xc4,
	0xb2, 0x00, 0xcd, 0x8f, 0xba, 0x07, 0xd3, 0x34, 0x9d, 0x46, 0xc4, 0xc2, 0x19, 0xb5, 0x26, 0x94,
	0x44, 0xa1, 0x3f, 0x26, 0x33, 0x3c, 0xa7, 0x29, 0x53, 0xb2, 0xee, 0x7e, 0x49, 0x90, 0xa7, 0x71,
	0x31, 0xb1, 0xc2, 0x82, 0x61, 0x4e, 0xd3, 0xa4, 0xc4, 0x0f, 0x1e, 0xe3, 0x9c, 0xc6, 0x24, 0xe7,
	0x38, 0xce, 0x4a, 0xc2, 0xc9, 0x4b, 0x5e, 0xcd, 0xef, 0x32, 0x62, 0x05, 0x69, 0x1c, 0x57, 0xbe,
	0x87, 0xdf, 0x1b, 0x00, 0x8c, 0x48, 0xce, 0x5d, 0x49, 0x84, 0x5d, 0x50, 0x4f, 0x70, 0x4c, 0x4c,
	0xad, 0xa7, 0xf5, 0x5b, 0x67, 0xcd, 0x7b, 0xbb, 0x76, 0x6f, 0x37, 0x5c, 0x19, 0x83, 0xff, 0x81,
	0x4d, 0x99, 0x2e, 0x0d, 0xcd, 0x0d, 0x09, 0x4b, 0xac, 0x29, 0x62, 0x4e, 0x08, 0xff, 0x07, 0x2d,
	0x75, 0x99, 0xc0, 0x6b, 0x95, 0xbc, 0x71, 0x6f, 0x6f, 0xb8, 0xba, 0x02, 0x9c, 0x10, 0x9e, 0x82,
	0xcd, 0x39, 0x66, 0x14, 0x27, 0xdc, 0xac, 0xf7, 0xb4, 0x7e, 0xfb, 0x78, 0x0f, 0xad, 0x97, 0x4b,
	0x3c, 0x10, 0x7d, 0x51, 0x14, 0xe5, 0x5f, 0x09, 0xe0, 0x01, 0xd0, 0xc9, 0x22, 0x23, 0x01, 0x27,
	0xa1, 0xd9, 0xe8, 0x69, 0x7d, 0x5d, 0xe1, 0xcb, 0x20, 0xfc, 0x00, 0x9a, 0x39, 0xc7, 0xbc, 0xc8,
	0xcd, 0x66, 0x4f, 0xeb, 0x6f, 0x1f, 0xf7, 0xd0, 0x53, 0xad, 0x40, 0x22, 0x5b, 0x4f, 0xf2, 0xca,
	0x04, 0x94, 0x08, 0xbe, 0x02, 0x5b, 0x79, 0x11, 0xc7, 0x98, 0xdd, 0xf9, 0x33, 0x1e, 0x47, 0xe6,
	0xe6, 0x2a, 0xc7, 0x76, 0x09, 0x5c, 0xf1, 0x38, 0x82, 0x1f, 0x01, 0xc8, 0x39, 0x66, 0xdc, 0x17,
	0x1d, 0x30, 0x75, 0x99, 0x46, 0x17, 0xa9, 0xf6, 0xa0, 0xaa, 0x3d, 0x68, 0x54, 0xb5, 0x47, 0x39,
	0xb4, 0xa4, 0x44, 0x04, 0xe1, 0x29, 0xd0, 0xab, 0xde, 0x9a, 0x2d, 0xa9, 0xfe, 0xf7, 0x37, 0xf5,
	0x45, 0x49, 0x28, 0x53, 0xac, 0xf8, 0xf0, 0x1d, 0xa8, 0x73, 0x3c, 0xcd, 0x4d, 0xd0, 0xab, 0xf5,
	0xdb, 0xc7, 0xfb, 0x4f, 0x15, 0xcf, 0xe3, 0x8c, 0x26, 0xd3, 0x21, 0xa6, 0x4c, 0x89, 0xa5, 0x00,
	0x5e, 0x83, 0x0e, 0x4d, 0xb2, 0x82, 0xfb, 0x98, 0x71, 0x3a, 0xc1, 0x01, 0xcf, 0xcd, 0xf6, 0x93,
	0x1e, 0x65, 0x91, 0xec, 0x92, 0xa6, 0x3c, 0xb6, 0xa5, 0xb6, 0x8a, 0xe5, 0xf0, 0x06, 0x18, 0x69,
	0xc1, 0xd7, 0xed, 0xb6, 0x5e, 0x6e, 0xd7, 0x51, 0xe2, 0xa5, 0xdf, 0xe1, 0x0f, 0x0d, 0xe8, 0xd5,
	0x09, 0xee, 0xae, 0x8d, 0xe0, 0x83, 0xf9, 0xdb, 0x03, 0xad, 0x09, 0xe1, 0xc1, 0xcc, 0x2f, 0x58,
	0xa4, 0x26, 0xd0, 0xd5, 0x65, 0xe0, 0x96, 0x45, 0xf0, 0x1a, 0xec, 0x2c, 0x41, 0x9f, 0x2c, 0x32,
	0x5a, 0x56, 0xb8, 0xf6, 0x5c, 0x7f, 0x5c, 0x58, 0x79, 0x0c, 0x96, 0x2a, 0x31, 0x0b, 0x41, 0x9a,
	0x70, 0x92, 0x70, 0x5f, 0x14, 0x55, 0x0e, 0x6b, 0x35, 0x0b, 0x25, 0x30, 0xba, 0xcb, 0x88, 0x78,
	0x6b, 0x4e, 0xbf, 0x11, 0x39, 0x8f, 0xb5, 0xf2, 0xad, 0x22, 0x70, 0xf8, 0x53, 0x03, 0x1d, 0x31,
	0x68, 0x83, 0x45, 0x9a, 0x90, 0xd2, 0xf4, 0x4f, 0xbb, 0xb5, 0xfb, 0x68, 0xb7, 0x96, 0x6b, 0x75,
	0xb2, 0xda, 0x98, 0xda, 0xb3, 0x1b, 0xb3, 0x5a, 0x96, 0x37, 0x60, 0x9b, 0xac, 0xae, 0x16, 0xb6,
	0xf5, 0xb5, 0x5b, 0xff, 0x7e, 0x80, 0x3a, 0x21, 0x44, 0xc0, 0x20, 0x8b, 0x2c, 0xc2, 0x89, 0xa2,
	0xcb, 0xf9, 0x6f, 0xac, 0x72, 0xee, 0x3c, 0x00, 0xc5, 0x0e, 0xbc, 0xfe, 0xaa, 0x7e, 0x1a, 0x6a,
	0x8d, 0xe0, 0x3f, 0x00, 0x7a, 0x23, 0x7b, 0x74, 0xeb, 0xf9, 0xb7, 0x37, 0xde, 0x70, 0x70, 0xee,
	0x5c, 0x3a, 0x83, 0x0b, 0xe3, 0x2f, 0xa8, 0x83, 0xfa, 0xd0, 0xf6, 0x3c, 0x43, 0x13, 0x5f, 0x97,
	0xb6, 0x73, 0x6d, 0x6c, 0xc0, 0x16, 0x68, 0x9c, 0xbb, 0xb6, 0x77, 0x65, 0xd4, 0xc4, 0xa7, 0x7d,
	0xf6, 0xd9, 0x1d, 0x19, 0x75, 0x81, 0x7b, 0x9f, 0x9c, 0xa1, 0xd1, 0x18, 0x37, 0x65, 0x87, 0xde,
	0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x15, 0xe6, 0x54, 0x96, 0x85, 0x05, 0x00, 0x00,
}
