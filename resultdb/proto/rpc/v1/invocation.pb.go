// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/resultdb/proto/rpc/v1/invocation.proto

package resultpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Invocation_State int32

const (
	// The default value. This value is used if the state is omitted.
	Invocation_STATE_UNSPECIFIED Invocation_State = 0
	// The invocation was created and accepts new results.
	Invocation_ACTIVE Invocation_State = 1
	// The invocation is in the process of transitioning into FINALIZED state.
	// This will happen automatically soon after all of its directly or
	// indirectly included invocations become inactive.
	Invocation_FINALIZING Invocation_State = 2
	// The invocation is immutable and no longer accepts new results nor
	// inclusions directly or indirectly.
	Invocation_FINALIZED Invocation_State = 3
)

var Invocation_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "ACTIVE",
	2: "FINALIZING",
	3: "FINALIZED",
}

var Invocation_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"ACTIVE":            1,
	"FINALIZING":        2,
	"FINALIZED":         3,
}

func (x Invocation_State) String() string {
	return proto.EnumName(Invocation_State_name, int32(x))
}

func (Invocation_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4005c8951497aaef, []int{0, 0}
}

// A conceptual container of results. Immutable once finalized.
// It represents all results of some computation; examples: swarming task,
// buildbucket build, CQ attempt.
// Composable: can include other invocations, see inclusion.proto.
//
// Next id: 12.
type Invocation struct {
	// Can be used to refer to this invocation, e.g. in ResultDB.GetInvocation
	// RPC.
	// Format: invocations/{INVOCATION_ID}
	// See also https://aip.dev/122.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Current state of the invocation.
	State Invocation_State `protobuf:"varint,2,opt,name=state,proto3,enum=luci.resultdb.rpc.v1.Invocation_State" json:"state,omitempty"`
	// When the invocation was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Invocation-level string key-value pairs.
	// A key can be repeated.
	Tags []*_type.StringPair `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	// When the invocation was finalized, i.e. transitioned to FINALIZED state.
	// If this field is set, implies that the invocation is finalized.
	FinalizeTime *timestamp.Timestamp `protobuf:"bytes,6,opt,name=finalize_time,json=finalizeTime,proto3" json:"finalize_time,omitempty"`
	// Timestamp when the invocation will be forcefully finalized.
	// Can be extended with UpdateInvocation until finalized.
	Deadline *timestamp.Timestamp `protobuf:"bytes,7,opt,name=deadline,proto3" json:"deadline,omitempty"`
	// Names of invocations included into this one. Overall results of this
	// invocation is a UNION of results directly included into this invocation
	// and results from the included invocations, recursively.
	// For example, a Buildbucket build invocation may include invocations of its
	// child swarming tasks and represent overall result of the build,
	// encapsulating the internal structure of the build.
	//
	// The graph is directed.
	// There can be at most one edge between a given pair of invocations.
	// The shape of the graph does not matter. What matters is only the set of
	// reachable invocations. Thus cycles are allowed and are noop.
	//
	// QueryTestResults returns test results from the transitive closure of
	// invocations.
	//
	// Use Recorder.Include RPC to modify this field.
	IncludedInvocations []string `protobuf:"bytes,8,rep,name=included_invocations,json=includedInvocations,proto3" json:"included_invocations,omitempty"`
	// bigquery_exports indicates what BigQuery table(s) that results in this
	// invocation should export to.
	BigqueryExports []*BigQueryExport `protobuf:"bytes,9,rep,name=bigquery_exports,json=bigqueryExports,proto3" json:"bigquery_exports,omitempty"`
	// LUCI identity (e.g. "user:<email>") who created the invocation.
	// Typically, a LUCI service account (e.g.
	// "user:cr-buildbucket@appspot.gserviceaccount.com"), but can also be a user
	// (e.g. "user:johndoe@example.com").
	CreatedBy string `protobuf:"bytes,10,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	// Full name of the resource that produced results in this invocation.
	// See also https://aip.dev/122#full-resource-names
	// Typical examples:
	// - Swarming task: "//chromium-swarm.appspot.com/tasks/deadbeef"
	// - Buildbucket build: "//cr-buildbucket.appspot.com/builds/1234567890".
	ProducerResource     string   `protobuf:"bytes,11,opt,name=producer_resource,json=producerResource,proto3" json:"producer_resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Invocation) Reset()         { *m = Invocation{} }
func (m *Invocation) String() string { return proto.CompactTextString(m) }
func (*Invocation) ProtoMessage()    {}
func (*Invocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_4005c8951497aaef, []int{0}
}

func (m *Invocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invocation.Unmarshal(m, b)
}
func (m *Invocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invocation.Marshal(b, m, deterministic)
}
func (m *Invocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invocation.Merge(m, src)
}
func (m *Invocation) XXX_Size() int {
	return xxx_messageInfo_Invocation.Size(m)
}
func (m *Invocation) XXX_DiscardUnknown() {
	xxx_messageInfo_Invocation.DiscardUnknown(m)
}

var xxx_messageInfo_Invocation proto.InternalMessageInfo

func (m *Invocation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Invocation) GetState() Invocation_State {
	if m != nil {
		return m.State
	}
	return Invocation_STATE_UNSPECIFIED
}

func (m *Invocation) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Invocation) GetTags() []*_type.StringPair {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Invocation) GetFinalizeTime() *timestamp.Timestamp {
	if m != nil {
		return m.FinalizeTime
	}
	return nil
}

func (m *Invocation) GetDeadline() *timestamp.Timestamp {
	if m != nil {
		return m.Deadline
	}
	return nil
}

func (m *Invocation) GetIncludedInvocations() []string {
	if m != nil {
		return m.IncludedInvocations
	}
	return nil
}

func (m *Invocation) GetBigqueryExports() []*BigQueryExport {
	if m != nil {
		return m.BigqueryExports
	}
	return nil
}

func (m *Invocation) GetCreatedBy() string {
	if m != nil {
		return m.CreatedBy
	}
	return ""
}

func (m *Invocation) GetProducerResource() string {
	if m != nil {
		return m.ProducerResource
	}
	return ""
}

// BigQueryExport indicates that results in this invocation should be exported
// to BigQuery after finalization.
type BigQueryExport struct {
	// Name of the BigQuery project.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Name of the BigQuery Dataset.
	Dataset string `protobuf:"bytes,2,opt,name=dataset,proto3" json:"dataset,omitempty"`
	// Name of the BigQuery Table.
	Table                string                      `protobuf:"bytes,3,opt,name=table,proto3" json:"table,omitempty"`
	TestResults          *BigQueryExport_TestResults `protobuf:"bytes,4,opt,name=test_results,json=testResults,proto3" json:"test_results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *BigQueryExport) Reset()         { *m = BigQueryExport{} }
func (m *BigQueryExport) String() string { return proto.CompactTextString(m) }
func (*BigQueryExport) ProtoMessage()    {}
func (*BigQueryExport) Descriptor() ([]byte, []int) {
	return fileDescriptor_4005c8951497aaef, []int{1}
}

func (m *BigQueryExport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BigQueryExport.Unmarshal(m, b)
}
func (m *BigQueryExport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BigQueryExport.Marshal(b, m, deterministic)
}
func (m *BigQueryExport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BigQueryExport.Merge(m, src)
}
func (m *BigQueryExport) XXX_Size() int {
	return xxx_messageInfo_BigQueryExport.Size(m)
}
func (m *BigQueryExport) XXX_DiscardUnknown() {
	xxx_messageInfo_BigQueryExport.DiscardUnknown(m)
}

var xxx_messageInfo_BigQueryExport proto.InternalMessageInfo

func (m *BigQueryExport) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *BigQueryExport) GetDataset() string {
	if m != nil {
		return m.Dataset
	}
	return ""
}

func (m *BigQueryExport) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *BigQueryExport) GetTestResults() *BigQueryExport_TestResults {
	if m != nil {
		return m.TestResults
	}
	return nil
}

// TestResultExport indicates that test results should be exported.
type BigQueryExport_TestResults struct {
	// Use predicate to query test results that should be exported to
	// BigQuery table.
	Predicate            *TestResultPredicate `protobuf:"bytes,1,opt,name=predicate,proto3" json:"predicate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BigQueryExport_TestResults) Reset()         { *m = BigQueryExport_TestResults{} }
func (m *BigQueryExport_TestResults) String() string { return proto.CompactTextString(m) }
func (*BigQueryExport_TestResults) ProtoMessage()    {}
func (*BigQueryExport_TestResults) Descriptor() ([]byte, []int) {
	return fileDescriptor_4005c8951497aaef, []int{1, 0}
}

func (m *BigQueryExport_TestResults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BigQueryExport_TestResults.Unmarshal(m, b)
}
func (m *BigQueryExport_TestResults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BigQueryExport_TestResults.Marshal(b, m, deterministic)
}
func (m *BigQueryExport_TestResults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BigQueryExport_TestResults.Merge(m, src)
}
func (m *BigQueryExport_TestResults) XXX_Size() int {
	return xxx_messageInfo_BigQueryExport_TestResults.Size(m)
}
func (m *BigQueryExport_TestResults) XXX_DiscardUnknown() {
	xxx_messageInfo_BigQueryExport_TestResults.DiscardUnknown(m)
}

var xxx_messageInfo_BigQueryExport_TestResults proto.InternalMessageInfo

func (m *BigQueryExport_TestResults) GetPredicate() *TestResultPredicate {
	if m != nil {
		return m.Predicate
	}
	return nil
}

func init() {
	proto.RegisterEnum("luci.resultdb.rpc.v1.Invocation_State", Invocation_State_name, Invocation_State_value)
	proto.RegisterType((*Invocation)(nil), "luci.resultdb.rpc.v1.Invocation")
	proto.RegisterType((*BigQueryExport)(nil), "luci.resultdb.rpc.v1.BigQueryExport")
	proto.RegisterType((*BigQueryExport_TestResults)(nil), "luci.resultdb.rpc.v1.BigQueryExport.TestResults")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/resultdb/proto/rpc/v1/invocation.proto", fileDescriptor_4005c8951497aaef)
}

var fileDescriptor_4005c8951497aaef = []byte{
	// 627 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x25, 0x71, 0x92, 0x36, 0x93, 0xb6, 0xa4, 0x4b, 0x91, 0x4c, 0x24, 0x20, 0x8a, 0x10, 0x0a,
	0x42, 0x5a, 0xd3, 0x54, 0xed, 0x81, 0x9e, 0x9c, 0x36, 0xad, 0x8c, 0x50, 0x29, 0x4e, 0xe8, 0xa1,
	0x17, 0x6b, 0x6d, 0x6f, 0xdd, 0x45, 0x8e, 0xd7, 0xac, 0xd7, 0x11, 0xe1, 0x1f, 0xb8, 0xf2, 0x7d,
	0xf9, 0x14, 0xe4, 0xb5, 0x1d, 0x03, 0xaa, 0xd4, 0x72, 0x9c, 0x37, 0xef, 0x79, 0xe6, 0xbd, 0x9d,
	0x04, 0x8e, 0x03, 0x8e, 0xbd, 0x5b, 0xc1, 0xe7, 0x2c, 0x9d, 0x63, 0x2e, 0x02, 0x23, 0x4c, 0x3d,
	0x66, 0x08, 0x9a, 0xa4, 0xa1, 0xf4, 0x5d, 0x23, 0x16, 0x5c, 0x72, 0x43, 0xc4, 0x9e, 0xb1, 0xd8,
	0x37, 0x58, 0xb4, 0xe0, 0x1e, 0x91, 0x8c, 0x47, 0x58, 0xe1, 0x68, 0x2f, 0x23, 0xe3, 0x92, 0x8c,
	0x45, 0xec, 0xe1, 0xc5, 0x7e, 0xef, 0x65, 0xc0, 0x79, 0x10, 0x52, 0x83, 0xc4, 0xcc, 0xb8, 0x61,
	0x34, 0xf4, 0x1d, 0x97, 0xde, 0x92, 0x05, 0xe3, 0x22, 0x97, 0xad, 0x09, 0xaa, 0x72, 0xd3, 0x1b,
	0x43, 0xb2, 0x39, 0x4d, 0x24, 0x99, 0xc7, 0x05, 0xe1, 0xfd, 0x7f, 0x2c, 0x15, 0x0b, 0xea, 0x33,
	0x8f, 0x48, 0x5a, 0x68, 0x0f, 0x1f, 0xa2, 0x95, 0xcb, 0x98, 0x1a, 0x1e, 0x9f, 0xcf, 0x4b, 0x2b,
	0x83, 0x9f, 0x4d, 0x00, 0x6b, 0xed, 0x0f, 0xf5, 0xa0, 0x11, 0x91, 0x39, 0xd5, 0x6b, 0xfd, 0xda,
	0xb0, 0x3d, 0x6e, 0xad, 0x4c, 0x6d, 0x65, 0x36, 0x6d, 0x85, 0x21, 0x13, 0x9a, 0x89, 0x24, 0x92,
	0xea, 0xf5, 0x7e, 0x6d, 0xb8, 0x33, 0x7a, 0x8d, 0xef, 0x4a, 0x01, 0x57, 0x1f, 0xc3, 0xd3, 0x8c,
	0x3d, 0xd6, 0x56, 0xa6, 0x66, 0xe7, 0x4a, 0x74, 0x02, 0x1d, 0x4f, 0x50, 0x22, 0xa9, 0x93, 0x59,
	0xd7, 0x1b, 0xfd, 0xda, 0xb0, 0x33, 0xea, 0xe1, 0x3c, 0x17, 0x5c, 0xe6, 0x82, 0x67, 0x65, 0x2e,
	0xeb, 0x0d, 0x20, 0x97, 0x65, 0x0d, 0x34, 0x82, 0x86, 0x24, 0x41, 0xa2, 0x37, 0xfb, 0xda, 0xb0,
	0x33, 0x7a, 0xf1, 0xcf, 0x1a, 0x99, 0x45, 0x3c, 0x95, 0x82, 0x45, 0xc1, 0x25, 0x61, 0xc2, 0x56,
	0x5c, 0x74, 0x0a, 0xdb, 0x37, 0x2c, 0x22, 0x21, 0xfb, 0x51, 0x8c, 0x6e, 0xdd, 0x3b, 0x5a, 0xed,
	0xbd, 0x55, 0xaa, 0xd4, 0xe4, 0x23, 0xd8, 0xf4, 0x29, 0xf1, 0x43, 0x16, 0x51, 0x7d, 0xe3, 0xbe,
	0x0f, 0xd8, 0x6b, 0x2e, 0x3a, 0x82, 0x3d, 0x16, 0x79, 0x61, 0xea, 0x53, 0xdf, 0xa9, 0x8e, 0x29,
	0xd1, 0x37, 0xfb, 0xda, 0xb0, 0x9d, 0x0f, 0x7a, 0x52, 0x12, 0xaa, 0xfc, 0x12, 0xf4, 0x09, 0xba,
	0x2e, 0x0b, 0xbe, 0xa5, 0x54, 0x2c, 0x1d, 0xfa, 0x3d, 0xe6, 0x42, 0x26, 0x7a, 0x5b, 0xb9, 0x7e,
	0x75, 0x77, 0xf8, 0x63, 0x16, 0x7c, 0xce, 0xd8, 0x13, 0x45, 0xb6, 0x1f, 0x97, 0xea, 0xbc, 0x4e,
	0xd0, 0x00, 0x8a, 0x20, 0x7d, 0xc7, 0x5d, 0xea, 0xa0, 0x1e, 0x59, 0x8d, 0x6f, 0x17, 0xf0, 0x78,
	0x89, 0xde, 0xc2, 0x6e, 0x2c, 0xb8, 0x9f, 0x7a, 0x54, 0x38, 0x82, 0x26, 0x3c, 0x15, 0x1e, 0xd5,
	0x3b, 0x19, 0xd5, 0xee, 0x96, 0x0d, 0xbb, 0xc0, 0x07, 0x16, 0x34, 0xd5, 0x2b, 0xa3, 0xa7, 0xb0,
	0x3b, 0x9d, 0x99, 0xb3, 0x89, 0xf3, 0xe5, 0x62, 0x7a, 0x39, 0x39, 0xb1, 0xce, 0xac, 0xc9, 0x69,
	0xf7, 0x11, 0x02, 0x68, 0x99, 0x27, 0x33, 0xeb, 0x6a, 0xd2, 0xad, 0xa1, 0x1d, 0x80, 0x33, 0xeb,
	0xc2, 0xfc, 0x68, 0x5d, 0x5b, 0x17, 0xe7, 0xdd, 0x3a, 0xda, 0x86, 0x76, 0x51, 0x4f, 0x4e, 0xbb,
	0xda, 0x87, 0xc6, 0xa6, 0xd6, 0x6d, 0x0c, 0x7e, 0xd5, 0x61, 0xe7, 0x6f, 0x17, 0xe8, 0x39, 0x6c,
	0xc4, 0x82, 0x7f, 0xa5, 0x9e, 0x2c, 0xce, 0x52, 0x5b, 0x99, 0x75, 0xbb, 0xc4, 0xb2, 0xb6, 0x4f,
	0x24, 0x49, 0xa8, 0x54, 0x87, 0x59, 0xb6, 0x0b, 0x0c, 0x3d, 0x83, 0xa6, 0x24, 0x6e, 0x48, 0x75,
	0xad, 0x6a, 0xe6, 0x08, 0x9a, 0xc2, 0x96, 0xa4, 0x89, 0x74, 0xf2, 0x14, 0x93, 0xe2, 0x1c, 0xdf,
	0x3d, 0x24, 0x5a, 0x3c, 0xa3, 0x89, 0xb4, 0x73, 0x9d, 0xdd, 0x91, 0x55, 0xd1, 0xbb, 0x82, 0xce,
	0x1f, 0x3d, 0x74, 0x0e, 0xed, 0xf5, 0x2f, 0x55, 0xad, 0xdf, 0x19, 0xbd, 0xb9, 0x7b, 0x40, 0xa5,
	0xba, 0x2c, 0x05, 0x76, 0xa5, 0x1d, 0x1f, 0x5e, 0x1f, 0x3c, 0xfc, 0xdf, 0xe1, 0x38, 0x47, 0x63,
	0xd7, 0x6d, 0x29, 0xf8, 0xe0, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x75, 0x98, 0xd6, 0x38, 0xf0,
	0x04, 0x00, 0x00,
}
