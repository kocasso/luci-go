// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/resultdb/proto/rpc/v1/test_result.proto

package resultspb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
// This message does not specify the test path.
// It should be available in the message that embeds this message.
type TestResult struct {
	// Resource name.
	// Format: "invocations/{INVOCATION_ID}/tests/{URL_ESCAPED_TEST_PATH}/results/{RESULT_ID}".
	// URL_ESCAPED_TEST_PATH is test_path escaped with https://golang.org/pkg/net/url/#PathEscape
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Test path, a unique identifier of the test in a LUCI project.
	// Regex: ^[[::print::]]+$.
	//
	// If two tests have a common test path prefix that ends with a
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
	// LUCI does not interpret test paths in any other way.
	TestPath string `protobuf:"bytes,2,opt,name=test_path,json=testPath,proto3" json:"test_path,omitempty"`
	// Identifies a test result in a given invocation and test path.
	// Regex: ^[[:ascii:]]{1,32}$.
	ResultId string `protobuf:"bytes,3,opt,name=result_id,json=resultId,proto3" json:"result_id,omitempty"`
	// Added to Invocation.base_test_variant_def of the parent invocation.
	// The complete variant definition of this test result is the result of
	// addition.
	//
	// MUST NOT have keys present in Invocation.test_variant_def, or MUST
	// have the same value.
	ExtraVariantPairs *VariantDef `protobuf:"bytes,4,opt,name=extra_variant_pairs,json=extraVariantPairs,proto3" json:"extra_variant_pairs,omitempty"`
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
	Status TestStatus `protobuf:"varint,6,opt,name=status,proto3,enum=luci.resultdb.TestStatus" json:"status,omitempty"`
	// Human-readable explanation of the result.
	// Markdown spec: https://spec.commonmark.org/0.29/
	SummaryMarkdown string `protobuf:"bytes,7,opt,name=summary_markdown,json=summaryMarkdown,proto3" json:"summary_markdown,omitempty"`
	// The point in time when the test case started to execute.
	StartTime *timestamp.Timestamp `protobuf:"bytes,8,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Duration of the test case execution.
	Duration *duration.Duration `protobuf:"bytes,9,opt,name=duration,proto3" json:"duration,omitempty"`
	// Metadata for this test result.
	// It might describe this particular execution or the test case.
	Tags []*StringPair `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty"`
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

func (m *TestResult) GetTestPath() string {
	if m != nil {
		return m.TestPath
	}
	return ""
}

func (m *TestResult) GetResultId() string {
	if m != nil {
		return m.ResultId
	}
	return ""
}

func (m *TestResult) GetExtraVariantPairs() *VariantDef {
	if m != nil {
		return m.ExtraVariantPairs
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

func (m *TestResult) GetSummaryMarkdown() string {
	if m != nil {
		return m.SummaryMarkdown
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

func (m *TestResult) GetTags() []*StringPair {
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
	// Example: "traces/a.txt".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Machine-readable URL to fetch the contents of the artifact.
	// Valid schemes: "isolate", "gs", "logdog", "rbe-cas".
	FetchUrl string `protobuf:"bytes,2,opt,name=fetch_url,json=fetchUrl,proto3" json:"fetch_url,omitempty"`
	// Human-consumable URL to the file content.
	// Typically a URL of a page where the user can view/download the arficact.
	ViewUrl string `protobuf:"bytes,3,opt,name=view_url,json=viewUrl,proto3" json:"view_url,omitempty"`
	// Media type of the artifact.
	// Logs are typically "plain/text" and screenshots are typically "image/png".
	ContentType string `protobuf:"bytes,4,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// Size of the file, in bytes.
	// Can be used in UI to decide whether to fetch an artifact and display it
	// inline, or only show a link if it is too large.
	Size int64 `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	// Contents of the artifact if it is stored inline with the test result.
	// Empty for artifacts stored elsewhere. To fetch such artifacts, use
	// fetch_url.
	// Size MUST be <= 8KB.
	Contents             []byte   `protobuf:"bytes,6,opt,name=contents,proto3" json:"contents,omitempty"`
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

func (m *Artifact) GetViewUrl() string {
	if m != nil {
		return m.ViewUrl
	}
	return ""
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

func (m *Artifact) GetContents() []byte {
	if m != nil {
		return m.Contents
	}
	return nil
}

// A pair of test path and a variant definition, identifying a variant of a
// test.
type TestVariant struct {
	// Test identifier, see TestResult.test_path.
	TestPath string `protobuf:"bytes,1,opt,name=test_path,json=testPath,proto3" json:"test_path,omitempty"`
	// Description of the variant of the test, see VariantDef.
	Variant              *VariantDef `protobuf:"bytes,2,opt,name=variant,proto3" json:"variant,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TestVariant) Reset()         { *m = TestVariant{} }
func (m *TestVariant) String() string { return proto.CompactTextString(m) }
func (*TestVariant) ProtoMessage()    {}
func (*TestVariant) Descriptor() ([]byte, []int) {
	return fileDescriptor_8253a2b79929a2cf, []int{2}
}

func (m *TestVariant) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestVariant.Unmarshal(m, b)
}
func (m *TestVariant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestVariant.Marshal(b, m, deterministic)
}
func (m *TestVariant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestVariant.Merge(m, src)
}
func (m *TestVariant) XXX_Size() int {
	return xxx_messageInfo_TestVariant.Size(m)
}
func (m *TestVariant) XXX_DiscardUnknown() {
	xxx_messageInfo_TestVariant.DiscardUnknown(m)
}

var xxx_messageInfo_TestVariant proto.InternalMessageInfo

func (m *TestVariant) GetTestPath() string {
	if m != nil {
		return m.TestPath
	}
	return ""
}

func (m *TestVariant) GetVariant() *VariantDef {
	if m != nil {
		return m.Variant
	}
	return nil
}

// Indicates the test subject (e.g. a CL) is absolved from blame
// for an unexpected result of a test variant.
// For example, the test variant fails both with and without CL, so it is not
// CL's fault.
type TestExoneration struct {
	// The resource name of the exoneration.
	// Format: invocations/{INVOCATION_ID}/testExonerations/{EXONERATION_ID}.
	// EXONERATION_ID is server-generated.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Failure of this test variant is forgiven.
	TestVariant *TestVariant `protobuf:"bytes,2,opt,name=test_variant,json=testVariant,proto3" json:"test_variant,omitempty"`
	// Reasoning behind the exoneration, in markdown.
	// Markdown spec: https://spec.commonmark.org/0.29/
	ExplanationMarkdown  string   `protobuf:"bytes,3,opt,name=explanation_markdown,json=explanationMarkdown,proto3" json:"explanation_markdown,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestExoneration) Reset()         { *m = TestExoneration{} }
func (m *TestExoneration) String() string { return proto.CompactTextString(m) }
func (*TestExoneration) ProtoMessage()    {}
func (*TestExoneration) Descriptor() ([]byte, []int) {
	return fileDescriptor_8253a2b79929a2cf, []int{3}
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

func (m *TestExoneration) GetTestVariant() *TestVariant {
	if m != nil {
		return m.TestVariant
	}
	return nil
}

func (m *TestExoneration) GetExplanationMarkdown() string {
	if m != nil {
		return m.ExplanationMarkdown
	}
	return ""
}

func init() {
	proto.RegisterEnum("luci.resultdb.TestStatus", TestStatus_name, TestStatus_value)
	proto.RegisterType((*TestResult)(nil), "luci.resultdb.TestResult")
	proto.RegisterType((*Artifact)(nil), "luci.resultdb.Artifact")
	proto.RegisterType((*TestVariant)(nil), "luci.resultdb.TestVariant")
	proto.RegisterType((*TestExoneration)(nil), "luci.resultdb.TestExoneration")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/resultdb/proto/rpc/v1/test_result.proto", fileDescriptor_8253a2b79929a2cf)
}

var fileDescriptor_8253a2b79929a2cf = []byte{
	// 708 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdb, 0x6e, 0xd3, 0x4a,
	0x14, 0x3d, 0x6e, 0x2e, 0x75, 0x76, 0x72, 0x1a, 0x9f, 0xe9, 0xd1, 0xa9, 0x4f, 0x90, 0xda, 0x28,
	0x48, 0x28, 0xe2, 0xc1, 0x16, 0xad, 0x54, 0x24, 0x84, 0x90, 0xdc, 0x36, 0x85, 0x08, 0x28, 0x51,
	0x9c, 0x22, 0xde, 0xac, 0x89, 0x33, 0x49, 0x46, 0xc4, 0x17, 0x8d, 0xc7, 0x69, 0xca, 0x17, 0xf0,
	0x2b, 0x7c, 0x09, 0xbf, 0xd1, 0x4f, 0x41, 0x73, 0x31, 0x4d, 0x4a, 0x55, 0xc1, 0x5b, 0xbc, 0xd7,
	0x5a, 0xfb, 0x36, 0x7b, 0x05, 0x5e, 0xce, 0x12, 0x27, 0x9c, 0xb3, 0x24, 0xa2, 0x79, 0xe4, 0x24,
	0x6c, 0xe6, 0x2e, 0xf2, 0x90, 0xba, 0x8c, 0x64, 0xf9, 0x82, 0x4f, 0xc6, 0x6e, 0xca, 0x12, 0x9e,
	0xb8, 0x2c, 0x0d, 0xdd, 0xe5, 0x33, 0x97, 0x93, 0x8c, 0x07, 0x0a, 0x72, 0x24, 0x80, 0xfe, 0x16,
	0x6c, 0xa7, 0x60, 0xb7, 0x0e, 0x66, 0x49, 0x32, 0x5b, 0x10, 0x17, 0xa7, 0xd4, 0x9d, 0x52, 0xb2,
	0x98, 0x04, 0x63, 0x32, 0xc7, 0x4b, 0x9a, 0x30, 0xc5, 0x6f, 0xed, 0x6b, 0x82, 0xfc, 0x1a, 0xe7,
	0x53, 0x77, 0x92, 0x33, 0xcc, 0x69, 0x12, 0x6b, 0xfc, 0xe0, 0x2e, 0xce, 0x69, 0x44, 0x32, 0x8e,
	0xa3, 0x54, 0x13, 0x9e, 0xff, 0x41, 0xbb, 0x61, 0x12, 0x45, 0x45, 0xe6, 0xce, 0xd7, 0x0a, 0xc0,
	0x88, 0x64, 0x7c, 0x28, 0xa9, 0xa8, 0x05, 0xe5, 0x18, 0x47, 0xc4, 0x36, 0xda, 0x46, 0xb7, 0x76,
	0x52, 0xbd, 0xf1, 0x4a, 0x37, 0x5e, 0x65, 0x28, 0x63, 0xa8, 0x0d, 0x35, 0x39, 0x69, 0x8a, 0xf9,
	0xdc, 0xde, 0x92, 0x04, 0x89, 0x9a, 0x22, 0x3a, 0xc0, 0x7c, 0x8e, 0x1e, 0x43, 0x4d, 0x95, 0x0c,
	0xe8, 0xc4, 0x2e, 0x15, 0x29, 0x2a, 0x37, 0xde, 0xd6, 0xd0, 0x54, 0x40, 0x7f, 0x82, 0x2e, 0x60,
	0x97, 0xac, 0x38, 0xc3, 0xc1, 0x12, 0x33, 0x8a, 0x63, 0x91, 0x8f, 0xb2, 0xcc, 0x2e, 0xb7, 0x8d,
	0x6e, 0xfd, 0xf0, 0x7f, 0x67, 0x63, 0x73, 0xce, 0x47, 0xc5, 0x39, 0x23, 0x53, 0x55, 0xeb, 0x1f,
	0x29, 0xd5, 0xd1, 0x81, 0x10, 0xa2, 0x03, 0x30, 0xc9, 0x2a, 0x25, 0x21, 0x27, 0x13, 0xbb, 0xd2,
	0x36, 0xba, 0xa6, 0xee, 0xaa, 0x08, 0xa2, 0x63, 0xa8, 0x66, 0x1c, 0xf3, 0x3c, 0xb3, 0xab, 0x6d,
	0xa3, 0xbb, 0xf3, 0x4b, 0x0d, 0x31, 0xbe, 0x2f, 0x09, 0x4a, 0xa9, 0xd9, 0xc8, 0x01, 0x2b, 0xcb,
	0xa3, 0x08, 0xb3, 0xeb, 0x20, 0xc2, 0xec, 0xf3, 0x24, 0xb9, 0x8a, 0xed, 0xed, 0xdb, 0xb1, 0x9b,
	0x1a, 0x7c, 0xaf, 0x31, 0xf4, 0x0a, 0x20, 0xe3, 0x98, 0xf1, 0x40, 0x3c, 0x8e, 0x6d, 0xca, 0x79,
	0x5a, 0x8e, 0x7a, 0x39, 0xa7, 0x78, 0x39, 0x67, 0x54, 0xbc, 0x9c, 0xca, 0x52, 0x93, 0x12, 0x11,
	0x44, 0x2f, 0xc0, 0x2c, 0x9e, 0xdd, 0xae, 0xe9, 0x6d, 0xdc, 0x55, 0x9f, 0x69, 0x82, 0x9e, 0xb1,
	0xe0, 0xa3, 0x43, 0x28, 0x73, 0x3c, 0xcb, 0x6c, 0x68, 0x97, 0xee, 0xd9, 0xa2, 0xcf, 0x19, 0x8d,
	0x67, 0x62, 0x5d, 0x4a, 0x27, 0xb9, 0xa8, 0x07, 0x4d, 0x1a, 0xa7, 0x39, 0x0f, 0x30, 0xe3, 0x74,
	0x8a, 0x43, 0x9e, 0xd9, 0x75, 0x29, 0xdf, 0xbb, 0x23, 0xf7, 0x34, 0xae, 0xc4, 0x3b, 0x52, 0x54,
	0xc4, 0x32, 0xf4, 0x1a, 0xac, 0x24, 0xe7, 0x9b, 0x79, 0x1a, 0xbf, 0x91, 0xa7, 0xa9, 0x54, 0x3f,
	0x13, 0x75, 0xbe, 0x1b, 0x60, 0x16, 0x5f, 0x68, 0x6f, 0xe3, 0x10, 0x37, 0xaf, 0x70, 0x4a, 0x78,
	0x38, 0x0f, 0x72, 0xb6, 0xd8, 0xb8, 0x42, 0x19, 0xbd, 0x64, 0x0b, 0xb4, 0x0f, 0xe6, 0x92, 0x92,
	0x2b, 0x49, 0x28, 0xdd, 0x12, 0xb6, 0x45, 0x50, 0xe0, 0x4f, 0xa0, 0x11, 0x26, 0x31, 0x27, 0x31,
	0x0f, 0xf8, 0x75, 0x4a, 0xe4, 0xe5, 0x69, 0x4e, 0x5d, 0x03, 0xa3, 0xeb, 0x94, 0x88, 0x16, 0x32,
	0xfa, 0x85, 0xc8, 0xa3, 0x2a, 0xe9, 0x16, 0x44, 0x40, 0x5c, 0x9c, 0xe6, 0xa9, 0x93, 0x6a, 0xe8,
	0x0e, 0x8a, 0x60, 0x27, 0x80, 0xba, 0x38, 0x2a, 0x7d, 0xa6, 0xe8, 0xd1, 0xba, 0x71, 0xe4, 0x40,
	0x6b, 0x9e, 0x39, 0x82, 0x6d, 0x6d, 0x04, 0x39, 0xcd, 0x43, 0x16, 0x18, 0x16, 0xcc, 0xce, 0x37,
	0x03, 0x9a, 0xa2, 0x42, 0x6f, 0x95, 0xc4, 0x44, 0x9f, 0xc0, 0x43, 0xd6, 0x3d, 0x81, 0x86, 0xec,
	0x60, 0xb3, 0x52, 0xeb, 0x1e, 0x23, 0xe8, 0x6a, 0x7a, 0x1d, 0x7c, 0x6d, 0x8a, 0x63, 0xf8, 0x97,
	0xac, 0xd2, 0x05, 0x8e, 0x65, 0xb9, 0x5b, 0x4b, 0xac, 0xad, 0x78, 0x77, 0x8d, 0x50, 0xd8, 0xe2,
	0xe9, 0x27, 0xf5, 0x07, 0xa3, 0x1c, 0x86, 0xfe, 0x03, 0xe4, 0x8f, 0xbc, 0xd1, 0xa5, 0x1f, 0x5c,
	0x5e, 0xf8, 0x83, 0xde, 0x69, 0xff, 0xbc, 0xdf, 0x3b, 0xb3, 0xfe, 0x42, 0x26, 0x94, 0x07, 0x9e,
	0xef, 0x5b, 0x86, 0xf8, 0x75, 0xee, 0xf5, 0xdf, 0x59, 0x5b, 0xa8, 0x06, 0x95, 0xd3, 0xa1, 0xe7,
	0xbf, 0xb1, 0x4a, 0xe2, 0xa7, 0x77, 0xf2, 0x61, 0x38, 0xb2, 0xca, 0x02, 0xf7, 0xdf, 0xf6, 0x07,
	0x56, 0x65, 0x5c, 0x95, 0xb6, 0x38, 0xfa, 0x11, 0x00, 0x00, 0xff, 0xff, 0x1a, 0xbb, 0x46, 0x54,
	0xac, 0x05, 0x00, 0x00,
}
