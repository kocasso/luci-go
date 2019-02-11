// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/lucicfg/external/crostesting/proto/config/target_test_requirements_config.proto

package config

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

type GceTestCfg struct {
	GceTest              []*GceTestCfg_GceTest `protobuf:"bytes,1,rep,name=gce_test,json=gceTest,proto3" json:"gce_test,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GceTestCfg) Reset()         { *m = GceTestCfg{} }
func (m *GceTestCfg) String() string { return proto.CompactTextString(m) }
func (*GceTestCfg) ProtoMessage()    {}
func (*GceTestCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_efae681839d29e51, []int{0}
}

func (m *GceTestCfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GceTestCfg.Unmarshal(m, b)
}
func (m *GceTestCfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GceTestCfg.Marshal(b, m, deterministic)
}
func (m *GceTestCfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GceTestCfg.Merge(m, src)
}
func (m *GceTestCfg) XXX_Size() int {
	return xxx_messageInfo_GceTestCfg.Size(m)
}
func (m *GceTestCfg) XXX_DiscardUnknown() {
	xxx_messageInfo_GceTestCfg.DiscardUnknown(m)
}

var xxx_messageInfo_GceTestCfg proto.InternalMessageInfo

func (m *GceTestCfg) GetGceTest() []*GceTestCfg_GceTest {
	if m != nil {
		return m.GceTest
	}
	return nil
}

type GceTestCfg_GceTest struct {
	// Test type to be run.
	TestType string `protobuf:"bytes,1,opt,name=test_type,json=testType,proto3" json:"test_type,omitempty"`
	// Test suite to be run in GCETest.
	TestSuite string `protobuf:"bytes,2,opt,name=test_suite,json=testSuite,proto3" json:"test_suite,omitempty"`
	// Number of seconds to wait before timing out waiting for results.
	TimeoutSec int32 `protobuf:"varint,3,opt,name=timeout_sec,json=timeoutSec,proto3" json:"timeout_sec,omitempty"`
	// Use the old ctest code path rather than the new chromite one.
	UseCtest             bool     `protobuf:"varint,4,opt,name=use_ctest,json=useCtest,proto3" json:"use_ctest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GceTestCfg_GceTest) Reset()         { *m = GceTestCfg_GceTest{} }
func (m *GceTestCfg_GceTest) String() string { return proto.CompactTextString(m) }
func (*GceTestCfg_GceTest) ProtoMessage()    {}
func (*GceTestCfg_GceTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_efae681839d29e51, []int{0, 0}
}

func (m *GceTestCfg_GceTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GceTestCfg_GceTest.Unmarshal(m, b)
}
func (m *GceTestCfg_GceTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GceTestCfg_GceTest.Marshal(b, m, deterministic)
}
func (m *GceTestCfg_GceTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GceTestCfg_GceTest.Merge(m, src)
}
func (m *GceTestCfg_GceTest) XXX_Size() int {
	return xxx_messageInfo_GceTestCfg_GceTest.Size(m)
}
func (m *GceTestCfg_GceTest) XXX_DiscardUnknown() {
	xxx_messageInfo_GceTestCfg_GceTest.DiscardUnknown(m)
}

var xxx_messageInfo_GceTestCfg_GceTest proto.InternalMessageInfo

func (m *GceTestCfg_GceTest) GetTestType() string {
	if m != nil {
		return m.TestType
	}
	return ""
}

func (m *GceTestCfg_GceTest) GetTestSuite() string {
	if m != nil {
		return m.TestSuite
	}
	return ""
}

func (m *GceTestCfg_GceTest) GetTimeoutSec() int32 {
	if m != nil {
		return m.TimeoutSec
	}
	return 0
}

func (m *GceTestCfg_GceTest) GetUseCtest() bool {
	if m != nil {
		return m.UseCtest
	}
	return false
}

type HwTestCfg struct {
	HwTest               []*HwTestCfg_HwTest `protobuf:"bytes,1,rep,name=hw_test,json=hwTest,proto3" json:"hw_test,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *HwTestCfg) Reset()         { *m = HwTestCfg{} }
func (m *HwTestCfg) String() string { return proto.CompactTextString(m) }
func (*HwTestCfg) ProtoMessage()    {}
func (*HwTestCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_efae681839d29e51, []int{1}
}

func (m *HwTestCfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HwTestCfg.Unmarshal(m, b)
}
func (m *HwTestCfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HwTestCfg.Marshal(b, m, deterministic)
}
func (m *HwTestCfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HwTestCfg.Merge(m, src)
}
func (m *HwTestCfg) XXX_Size() int {
	return xxx_messageInfo_HwTestCfg.Size(m)
}
func (m *HwTestCfg) XXX_DiscardUnknown() {
	xxx_messageInfo_HwTestCfg.DiscardUnknown(m)
}

var xxx_messageInfo_HwTestCfg proto.InternalMessageInfo

func (m *HwTestCfg) GetHwTest() []*HwTestCfg_HwTest {
	if m != nil {
		return m.HwTest
	}
	return nil
}

// Configuration for a hardware test suite.
type HwTestCfg_HwTest struct {
	// Name of the test suite to run.
	Suite string `protobuf:"bytes,1,opt,name=suite,proto3" json:"suite,omitempty"`
	// Number of seconds to wait before timing out waiting for results.
	TimeoutSec int32 `protobuf:"varint,2,opt,name=timeout_sec,json=timeoutSec,proto3" json:"timeout_sec,omitempty"`
	// Pool to use for hw testing.
	Pool string `protobuf:"bytes,3,opt,name=pool,proto3" json:"pool,omitempty"`
	// Note, if you want multiple suites to block other suites but run
	// in parallel, you should only mark the last one scheduled as
	// blocking (it effectively serves as a thread/process join).
	Blocking bool `protobuf:"varint,4,opt,name=blocking,proto3" json:"blocking,omitempty"`
	// Fire-and-forget suite.
	Async bool `protobuf:"varint,5,opt,name=async,proto3" json:"async,omitempty"`
	// Failure on HW tests warns only (does not generate error).
	WarnOnly bool `protobuf:"varint,6,opt,name=warn_only,json=warnOnly,proto3" json:"warn_only,omitempty"`
	// Usually we consider structural failures here as OK.
	Critical bool `protobuf:"varint,7,opt,name=critical,proto3" json:"critical,omitempty"`
	// Priority at which tests in the suite will be scheduled in the hw lab.
	Priority string `protobuf:"bytes,8,opt,name=priority,proto3" json:"priority,omitempty"`
	// Should we file bugs if a test fails in a suite run.
	FileBugs bool `protobuf:"varint,9,opt,name=file_bugs,json=fileBugs,proto3" json:"file_bugs,omitempty"`
	// Minimum number of DUTs required for testing in the hw lab.
	MinimumDuts int32 `protobuf:"varint,10,opt,name=minimum_duts,json=minimumDuts,proto3" json:"minimum_duts,omitempty"`
	// Whether we should retry tests that fail in a suite run.
	Retry bool `protobuf:"varint,11,opt,name=retry,proto3" json:"retry,omitempty"`
	// Maximum job retries allowed at suite level. 0 for no max.
	MaxRetries int32 `protobuf:"varint,12,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	// Preferred minimum duts. Lab will prioritize on getting such
	// number of duts even if the suite is competing with
	// other suites that have higher priority.
	SuiteMinDuts int32 `protobuf:"varint,13,opt,name=suite_min_duts,json=suiteMinDuts,proto3" json:"suite_min_duts,omitempty"`
	// Only offload failed tests to Google Storage.
	OffloadFailuresOnly  bool     `protobuf:"varint,14,opt,name=offload_failures_only,json=offloadFailuresOnly,proto3" json:"offload_failures_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HwTestCfg_HwTest) Reset()         { *m = HwTestCfg_HwTest{} }
func (m *HwTestCfg_HwTest) String() string { return proto.CompactTextString(m) }
func (*HwTestCfg_HwTest) ProtoMessage()    {}
func (*HwTestCfg_HwTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_efae681839d29e51, []int{1, 0}
}

func (m *HwTestCfg_HwTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HwTestCfg_HwTest.Unmarshal(m, b)
}
func (m *HwTestCfg_HwTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HwTestCfg_HwTest.Marshal(b, m, deterministic)
}
func (m *HwTestCfg_HwTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HwTestCfg_HwTest.Merge(m, src)
}
func (m *HwTestCfg_HwTest) XXX_Size() int {
	return xxx_messageInfo_HwTestCfg_HwTest.Size(m)
}
func (m *HwTestCfg_HwTest) XXX_DiscardUnknown() {
	xxx_messageInfo_HwTestCfg_HwTest.DiscardUnknown(m)
}

var xxx_messageInfo_HwTestCfg_HwTest proto.InternalMessageInfo

func (m *HwTestCfg_HwTest) GetSuite() string {
	if m != nil {
		return m.Suite
	}
	return ""
}

func (m *HwTestCfg_HwTest) GetTimeoutSec() int32 {
	if m != nil {
		return m.TimeoutSec
	}
	return 0
}

func (m *HwTestCfg_HwTest) GetPool() string {
	if m != nil {
		return m.Pool
	}
	return ""
}

func (m *HwTestCfg_HwTest) GetBlocking() bool {
	if m != nil {
		return m.Blocking
	}
	return false
}

func (m *HwTestCfg_HwTest) GetAsync() bool {
	if m != nil {
		return m.Async
	}
	return false
}

func (m *HwTestCfg_HwTest) GetWarnOnly() bool {
	if m != nil {
		return m.WarnOnly
	}
	return false
}

func (m *HwTestCfg_HwTest) GetCritical() bool {
	if m != nil {
		return m.Critical
	}
	return false
}

func (m *HwTestCfg_HwTest) GetPriority() string {
	if m != nil {
		return m.Priority
	}
	return ""
}

func (m *HwTestCfg_HwTest) GetFileBugs() bool {
	if m != nil {
		return m.FileBugs
	}
	return false
}

func (m *HwTestCfg_HwTest) GetMinimumDuts() int32 {
	if m != nil {
		return m.MinimumDuts
	}
	return 0
}

func (m *HwTestCfg_HwTest) GetRetry() bool {
	if m != nil {
		return m.Retry
	}
	return false
}

func (m *HwTestCfg_HwTest) GetMaxRetries() int32 {
	if m != nil {
		return m.MaxRetries
	}
	return 0
}

func (m *HwTestCfg_HwTest) GetSuiteMinDuts() int32 {
	if m != nil {
		return m.SuiteMinDuts
	}
	return 0
}

func (m *HwTestCfg_HwTest) GetOffloadFailuresOnly() bool {
	if m != nil {
		return m.OffloadFailuresOnly
	}
	return false
}

type MoblabVmTestCfg struct {
	MoblabTest           []*MoblabVmTestCfg_MoblabTest `protobuf:"bytes,1,rep,name=moblab_test,json=moblabTest,proto3" json:"moblab_test,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *MoblabVmTestCfg) Reset()         { *m = MoblabVmTestCfg{} }
func (m *MoblabVmTestCfg) String() string { return proto.CompactTextString(m) }
func (*MoblabVmTestCfg) ProtoMessage()    {}
func (*MoblabVmTestCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_efae681839d29e51, []int{2}
}

func (m *MoblabVmTestCfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoblabVmTestCfg.Unmarshal(m, b)
}
func (m *MoblabVmTestCfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoblabVmTestCfg.Marshal(b, m, deterministic)
}
func (m *MoblabVmTestCfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoblabVmTestCfg.Merge(m, src)
}
func (m *MoblabVmTestCfg) XXX_Size() int {
	return xxx_messageInfo_MoblabVmTestCfg.Size(m)
}
func (m *MoblabVmTestCfg) XXX_DiscardUnknown() {
	xxx_messageInfo_MoblabVmTestCfg.DiscardUnknown(m)
}

var xxx_messageInfo_MoblabVmTestCfg proto.InternalMessageInfo

func (m *MoblabVmTestCfg) GetMoblabTest() []*MoblabVmTestCfg_MoblabTest {
	if m != nil {
		return m.MoblabTest
	}
	return nil
}

type MoblabVmTestCfg_MoblabTest struct {
	// Test type to be run.
	TestType string `protobuf:"bytes,1,opt,name=test_type,json=testType,proto3" json:"test_type,omitempty"`
	// Number of seconds to wait before timing out waiting for results.
	TimeoutSec           int32    `protobuf:"varint,2,opt,name=timeout_sec,json=timeoutSec,proto3" json:"timeout_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MoblabVmTestCfg_MoblabTest) Reset()         { *m = MoblabVmTestCfg_MoblabTest{} }
func (m *MoblabVmTestCfg_MoblabTest) String() string { return proto.CompactTextString(m) }
func (*MoblabVmTestCfg_MoblabTest) ProtoMessage()    {}
func (*MoblabVmTestCfg_MoblabTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_efae681839d29e51, []int{2, 0}
}

func (m *MoblabVmTestCfg_MoblabTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoblabVmTestCfg_MoblabTest.Unmarshal(m, b)
}
func (m *MoblabVmTestCfg_MoblabTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoblabVmTestCfg_MoblabTest.Marshal(b, m, deterministic)
}
func (m *MoblabVmTestCfg_MoblabTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoblabVmTestCfg_MoblabTest.Merge(m, src)
}
func (m *MoblabVmTestCfg_MoblabTest) XXX_Size() int {
	return xxx_messageInfo_MoblabVmTestCfg_MoblabTest.Size(m)
}
func (m *MoblabVmTestCfg_MoblabTest) XXX_DiscardUnknown() {
	xxx_messageInfo_MoblabVmTestCfg_MoblabTest.DiscardUnknown(m)
}

var xxx_messageInfo_MoblabVmTestCfg_MoblabTest proto.InternalMessageInfo

func (m *MoblabVmTestCfg_MoblabTest) GetTestType() string {
	if m != nil {
		return m.TestType
	}
	return ""
}

func (m *MoblabVmTestCfg_MoblabTest) GetTimeoutSec() int32 {
	if m != nil {
		return m.TimeoutSec
	}
	return 0
}

type VmTestCfg struct {
	VmTest               []*VmTestCfg_VmTest     `protobuf:"bytes,1,rep,name=vm_test,json=vmTest,proto3" json:"vm_test,omitempty"`
	TastVmTest           []*VmTestCfg_TastVmTest `protobuf:"bytes,2,rep,name=tast_vm_test,json=tastVmTest,proto3" json:"tast_vm_test,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *VmTestCfg) Reset()         { *m = VmTestCfg{} }
func (m *VmTestCfg) String() string { return proto.CompactTextString(m) }
func (*VmTestCfg) ProtoMessage()    {}
func (*VmTestCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_efae681839d29e51, []int{3}
}

func (m *VmTestCfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VmTestCfg.Unmarshal(m, b)
}
func (m *VmTestCfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VmTestCfg.Marshal(b, m, deterministic)
}
func (m *VmTestCfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VmTestCfg.Merge(m, src)
}
func (m *VmTestCfg) XXX_Size() int {
	return xxx_messageInfo_VmTestCfg.Size(m)
}
func (m *VmTestCfg) XXX_DiscardUnknown() {
	xxx_messageInfo_VmTestCfg.DiscardUnknown(m)
}

var xxx_messageInfo_VmTestCfg proto.InternalMessageInfo

func (m *VmTestCfg) GetVmTest() []*VmTestCfg_VmTest {
	if m != nil {
		return m.VmTest
	}
	return nil
}

func (m *VmTestCfg) GetTastVmTest() []*VmTestCfg_TastVmTest {
	if m != nil {
		return m.TastVmTest
	}
	return nil
}

type VmTestCfg_VmTest struct {
	// Test type to be run.
	TestType string `protobuf:"bytes,1,opt,name=test_type,json=testType,proto3" json:"test_type,omitempty"`
	// Test suite to be run in VMTest.
	TestSuite string `protobuf:"bytes,2,opt,name=test_suite,json=testSuite,proto3" json:"test_suite,omitempty"`
	// Number of seconds to wait before timing out waiting for results.
	TimeoutSec int32 `protobuf:"varint,3,opt,name=timeout_sec,json=timeoutSec,proto3" json:"timeout_sec,omitempty"`
	// Whether we should retry tests that fail in a suite run.
	Retry bool `protobuf:"varint,4,opt,name=retry,proto3" json:"retry,omitempty"`
	// Maximum job retries allowed at suite level. 0 for no max.
	MaxRetries int32 `protobuf:"varint,5,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	// Failure on VM tests warns only.
	WarnOnly bool `protobuf:"varint,6,opt,name=warn_only,json=warnOnly,proto3" json:"warn_only,omitempty"`
	// Use the old ctest code path rather than the new chromite one.
	UseCtest             bool     `protobuf:"varint,7,opt,name=use_ctest,json=useCtest,proto3" json:"use_ctest,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VmTestCfg_VmTest) Reset()         { *m = VmTestCfg_VmTest{} }
func (m *VmTestCfg_VmTest) String() string { return proto.CompactTextString(m) }
func (*VmTestCfg_VmTest) ProtoMessage()    {}
func (*VmTestCfg_VmTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_efae681839d29e51, []int{3, 0}
}

func (m *VmTestCfg_VmTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VmTestCfg_VmTest.Unmarshal(m, b)
}
func (m *VmTestCfg_VmTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VmTestCfg_VmTest.Marshal(b, m, deterministic)
}
func (m *VmTestCfg_VmTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VmTestCfg_VmTest.Merge(m, src)
}
func (m *VmTestCfg_VmTest) XXX_Size() int {
	return xxx_messageInfo_VmTestCfg_VmTest.Size(m)
}
func (m *VmTestCfg_VmTest) XXX_DiscardUnknown() {
	xxx_messageInfo_VmTestCfg_VmTest.DiscardUnknown(m)
}

var xxx_messageInfo_VmTestCfg_VmTest proto.InternalMessageInfo

func (m *VmTestCfg_VmTest) GetTestType() string {
	if m != nil {
		return m.TestType
	}
	return ""
}

func (m *VmTestCfg_VmTest) GetTestSuite() string {
	if m != nil {
		return m.TestSuite
	}
	return ""
}

func (m *VmTestCfg_VmTest) GetTimeoutSec() int32 {
	if m != nil {
		return m.TimeoutSec
	}
	return 0
}

func (m *VmTestCfg_VmTest) GetRetry() bool {
	if m != nil {
		return m.Retry
	}
	return false
}

func (m *VmTestCfg_VmTest) GetMaxRetries() int32 {
	if m != nil {
		return m.MaxRetries
	}
	return 0
}

func (m *VmTestCfg_VmTest) GetWarnOnly() bool {
	if m != nil {
		return m.WarnOnly
	}
	return false
}

func (m *VmTestCfg_VmTest) GetUseCtest() bool {
	if m != nil {
		return m.UseCtest
	}
	return false
}

type VmTestCfg_TastTestExpr struct {
	// A single tast test expression. See https://goo.gl/UPNEgT
	TestExpr             string   `protobuf:"bytes,1,opt,name=test_expr,json=testExpr,proto3" json:"test_expr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VmTestCfg_TastTestExpr) Reset()         { *m = VmTestCfg_TastTestExpr{} }
func (m *VmTestCfg_TastTestExpr) String() string { return proto.CompactTextString(m) }
func (*VmTestCfg_TastTestExpr) ProtoMessage()    {}
func (*VmTestCfg_TastTestExpr) Descriptor() ([]byte, []int) {
	return fileDescriptor_efae681839d29e51, []int{3, 1}
}

func (m *VmTestCfg_TastTestExpr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VmTestCfg_TastTestExpr.Unmarshal(m, b)
}
func (m *VmTestCfg_TastTestExpr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VmTestCfg_TastTestExpr.Marshal(b, m, deterministic)
}
func (m *VmTestCfg_TastTestExpr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VmTestCfg_TastTestExpr.Merge(m, src)
}
func (m *VmTestCfg_TastTestExpr) XXX_Size() int {
	return xxx_messageInfo_VmTestCfg_TastTestExpr.Size(m)
}
func (m *VmTestCfg_TastTestExpr) XXX_DiscardUnknown() {
	xxx_messageInfo_VmTestCfg_TastTestExpr.DiscardUnknown(m)
}

var xxx_messageInfo_VmTestCfg_TastTestExpr proto.InternalMessageInfo

func (m *VmTestCfg_TastTestExpr) GetTestExpr() string {
	if m != nil {
		return m.TestExpr
	}
	return ""
}

type VmTestCfg_TastVmTest struct {
	// String containing short human-readable name describing test suite.
	SuiteName string `protobuf:"bytes,1,opt,name=suite_name,json=suiteName,proto3" json:"suite_name,omitempty"`
	// List of string expressions describing which tests to run; this
	// is passed directly to the 'tast run' command. See
	// https://goo.gl/UPNEgT for info about test expressions.
	TastTestExpr []*VmTestCfg_TastTestExpr `protobuf:"bytes,2,rep,name=tast_test_expr,json=tastTestExpr,proto3" json:"tast_test_expr,omitempty"`
	// Number of seconds to wait before timing out waiting for results.
	TimeoutSec           int32    `protobuf:"varint,3,opt,name=timeout_sec,json=timeoutSec,proto3" json:"timeout_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VmTestCfg_TastVmTest) Reset()         { *m = VmTestCfg_TastVmTest{} }
func (m *VmTestCfg_TastVmTest) String() string { return proto.CompactTextString(m) }
func (*VmTestCfg_TastVmTest) ProtoMessage()    {}
func (*VmTestCfg_TastVmTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_efae681839d29e51, []int{3, 2}
}

func (m *VmTestCfg_TastVmTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VmTestCfg_TastVmTest.Unmarshal(m, b)
}
func (m *VmTestCfg_TastVmTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VmTestCfg_TastVmTest.Marshal(b, m, deterministic)
}
func (m *VmTestCfg_TastVmTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VmTestCfg_TastVmTest.Merge(m, src)
}
func (m *VmTestCfg_TastVmTest) XXX_Size() int {
	return xxx_messageInfo_VmTestCfg_TastVmTest.Size(m)
}
func (m *VmTestCfg_TastVmTest) XXX_DiscardUnknown() {
	xxx_messageInfo_VmTestCfg_TastVmTest.DiscardUnknown(m)
}

var xxx_messageInfo_VmTestCfg_TastVmTest proto.InternalMessageInfo

func (m *VmTestCfg_TastVmTest) GetSuiteName() string {
	if m != nil {
		return m.SuiteName
	}
	return ""
}

func (m *VmTestCfg_TastVmTest) GetTastTestExpr() []*VmTestCfg_TastTestExpr {
	if m != nil {
		return m.TastTestExpr
	}
	return nil
}

func (m *VmTestCfg_TastVmTest) GetTimeoutSec() int32 {
	if m != nil {
		return m.TimeoutSec
	}
	return 0
}

// Specifies a CrOS build, either by reference design or by a specific build
// target.
type BuildCriteria struct {
	// Types that are valid to be assigned to TargetType:
	//	*BuildCriteria_ReferenceDesign
	//	*BuildCriteria_BuildTarget
	TargetType           isBuildCriteria_TargetType `protobuf_oneof:"target_type"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *BuildCriteria) Reset()         { *m = BuildCriteria{} }
func (m *BuildCriteria) String() string { return proto.CompactTextString(m) }
func (*BuildCriteria) ProtoMessage()    {}
func (*BuildCriteria) Descriptor() ([]byte, []int) {
	return fileDescriptor_efae681839d29e51, []int{4}
}

func (m *BuildCriteria) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildCriteria.Unmarshal(m, b)
}
func (m *BuildCriteria) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildCriteria.Marshal(b, m, deterministic)
}
func (m *BuildCriteria) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildCriteria.Merge(m, src)
}
func (m *BuildCriteria) XXX_Size() int {
	return xxx_messageInfo_BuildCriteria.Size(m)
}
func (m *BuildCriteria) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildCriteria.DiscardUnknown(m)
}

var xxx_messageInfo_BuildCriteria proto.InternalMessageInfo

type isBuildCriteria_TargetType interface {
	isBuildCriteria_TargetType()
}

type BuildCriteria_ReferenceDesign struct {
	ReferenceDesign string `protobuf:"bytes,1,opt,name=reference_design,json=referenceDesign,proto3,oneof"`
}

type BuildCriteria_BuildTarget struct {
	BuildTarget string `protobuf:"bytes,2,opt,name=build_target,json=buildTarget,proto3,oneof"`
}

func (*BuildCriteria_ReferenceDesign) isBuildCriteria_TargetType() {}

func (*BuildCriteria_BuildTarget) isBuildCriteria_TargetType() {}

func (m *BuildCriteria) GetTargetType() isBuildCriteria_TargetType {
	if m != nil {
		return m.TargetType
	}
	return nil
}

func (m *BuildCriteria) GetReferenceDesign() string {
	if x, ok := m.GetTargetType().(*BuildCriteria_ReferenceDesign); ok {
		return x.ReferenceDesign
	}
	return ""
}

func (m *BuildCriteria) GetBuildTarget() string {
	if x, ok := m.GetTargetType().(*BuildCriteria_BuildTarget); ok {
		return x.BuildTarget
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BuildCriteria) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BuildCriteria_ReferenceDesign)(nil),
		(*BuildCriteria_BuildTarget)(nil),
	}
}

// Details which testing is needed for a single CrOS build target.
type PerTargetTestRequirements struct {
	// Specifies the builds to which these testing requirements should be applied.
	BuildCriteria *BuildCriteria `protobuf:"bytes,1,opt,name=build_criteria,json=buildCriteria,proto3" json:"build_criteria,omitempty"`
	// These configure what testing is needed for these BuildCriteria.
	GceTestCfg           *GceTestCfg      `protobuf:"bytes,2,opt,name=gce_test_cfg,json=gceTestCfg,proto3" json:"gce_test_cfg,omitempty"`
	HwTestCfg            *HwTestCfg       `protobuf:"bytes,3,opt,name=hw_test_cfg,json=hwTestCfg,proto3" json:"hw_test_cfg,omitempty"`
	MoblabVmTestCfg      *MoblabVmTestCfg `protobuf:"bytes,4,opt,name=moblab_vm_test_cfg,json=moblabVmTestCfg,proto3" json:"moblab_vm_test_cfg,omitempty"`
	VmTestCfg            *VmTestCfg       `protobuf:"bytes,5,opt,name=vm_test_cfg,json=vmTestCfg,proto3" json:"vm_test_cfg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PerTargetTestRequirements) Reset()         { *m = PerTargetTestRequirements{} }
func (m *PerTargetTestRequirements) String() string { return proto.CompactTextString(m) }
func (*PerTargetTestRequirements) ProtoMessage()    {}
func (*PerTargetTestRequirements) Descriptor() ([]byte, []int) {
	return fileDescriptor_efae681839d29e51, []int{5}
}

func (m *PerTargetTestRequirements) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PerTargetTestRequirements.Unmarshal(m, b)
}
func (m *PerTargetTestRequirements) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PerTargetTestRequirements.Marshal(b, m, deterministic)
}
func (m *PerTargetTestRequirements) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerTargetTestRequirements.Merge(m, src)
}
func (m *PerTargetTestRequirements) XXX_Size() int {
	return xxx_messageInfo_PerTargetTestRequirements.Size(m)
}
func (m *PerTargetTestRequirements) XXX_DiscardUnknown() {
	xxx_messageInfo_PerTargetTestRequirements.DiscardUnknown(m)
}

var xxx_messageInfo_PerTargetTestRequirements proto.InternalMessageInfo

func (m *PerTargetTestRequirements) GetBuildCriteria() *BuildCriteria {
	if m != nil {
		return m.BuildCriteria
	}
	return nil
}

func (m *PerTargetTestRequirements) GetGceTestCfg() *GceTestCfg {
	if m != nil {
		return m.GceTestCfg
	}
	return nil
}

func (m *PerTargetTestRequirements) GetHwTestCfg() *HwTestCfg {
	if m != nil {
		return m.HwTestCfg
	}
	return nil
}

func (m *PerTargetTestRequirements) GetMoblabVmTestCfg() *MoblabVmTestCfg {
	if m != nil {
		return m.MoblabVmTestCfg
	}
	return nil
}

func (m *PerTargetTestRequirements) GetVmTestCfg() *VmTestCfg {
	if m != nil {
		return m.VmTestCfg
	}
	return nil
}

// A listing of all testing that should be done for all CrOS builds.
type TargetTestRequirementsCfg struct {
	// The testing that should be performed for a single CrOS build target or
	// reference design.
	PerTargetTestRequirements []*PerTargetTestRequirements `protobuf:"bytes,1,rep,name=per_target_test_requirements,json=perTargetTestRequirements,proto3" json:"per_target_test_requirements,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                     `json:"-"`
	XXX_unrecognized          []byte                       `json:"-"`
	XXX_sizecache             int32                        `json:"-"`
}

func (m *TargetTestRequirementsCfg) Reset()         { *m = TargetTestRequirementsCfg{} }
func (m *TargetTestRequirementsCfg) String() string { return proto.CompactTextString(m) }
func (*TargetTestRequirementsCfg) ProtoMessage()    {}
func (*TargetTestRequirementsCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_efae681839d29e51, []int{6}
}

func (m *TargetTestRequirementsCfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetTestRequirementsCfg.Unmarshal(m, b)
}
func (m *TargetTestRequirementsCfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetTestRequirementsCfg.Marshal(b, m, deterministic)
}
func (m *TargetTestRequirementsCfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetTestRequirementsCfg.Merge(m, src)
}
func (m *TargetTestRequirementsCfg) XXX_Size() int {
	return xxx_messageInfo_TargetTestRequirementsCfg.Size(m)
}
func (m *TargetTestRequirementsCfg) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetTestRequirementsCfg.DiscardUnknown(m)
}

var xxx_messageInfo_TargetTestRequirementsCfg proto.InternalMessageInfo

func (m *TargetTestRequirementsCfg) GetPerTargetTestRequirements() []*PerTargetTestRequirements {
	if m != nil {
		return m.PerTargetTestRequirements
	}
	return nil
}

func init() {
	proto.RegisterType((*GceTestCfg)(nil), "crostesting.GceTestCfg")
	proto.RegisterType((*GceTestCfg_GceTest)(nil), "crostesting.GceTestCfg.GceTest")
	proto.RegisterType((*HwTestCfg)(nil), "crostesting.HwTestCfg")
	proto.RegisterType((*HwTestCfg_HwTest)(nil), "crostesting.HwTestCfg.HwTest")
	proto.RegisterType((*MoblabVmTestCfg)(nil), "crostesting.MoblabVmTestCfg")
	proto.RegisterType((*MoblabVmTestCfg_MoblabTest)(nil), "crostesting.MoblabVmTestCfg.MoblabTest")
	proto.RegisterType((*VmTestCfg)(nil), "crostesting.VmTestCfg")
	proto.RegisterType((*VmTestCfg_VmTest)(nil), "crostesting.VmTestCfg.VmTest")
	proto.RegisterType((*VmTestCfg_TastTestExpr)(nil), "crostesting.VmTestCfg.TastTestExpr")
	proto.RegisterType((*VmTestCfg_TastVmTest)(nil), "crostesting.VmTestCfg.TastVmTest")
	proto.RegisterType((*BuildCriteria)(nil), "crostesting.BuildCriteria")
	proto.RegisterType((*PerTargetTestRequirements)(nil), "crostesting.PerTargetTestRequirements")
	proto.RegisterType((*TargetTestRequirementsCfg)(nil), "crostesting.TargetTestRequirementsCfg")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/lucicfg/external/crostesting/proto/config/target_test_requirements_config.proto", fileDescriptor_efae681839d29e51)
}

var fileDescriptor_efae681839d29e51 = []byte{
	// 894 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x5f, 0x6f, 0xe3, 0x44,
	0x10, 0x3f, 0x37, 0x7f, 0x3d, 0x4e, 0x5b, 0x64, 0xfe, 0xb9, 0xe1, 0x4e, 0xd7, 0xcb, 0x21, 0xa8,
	0x74, 0x52, 0x22, 0x15, 0xa9, 0x12, 0xbc, 0xd1, 0x1e, 0xd0, 0x43, 0x3a, 0x40, 0xbe, 0x8a, 0x07,
	0x5e, 0x2c, 0x67, 0x33, 0x76, 0x57, 0x78, 0x6d, 0xb3, 0xbb, 0x6e, 0x93, 0x67, 0xf8, 0x0e, 0x7c,
	0x04, 0x90, 0xf8, 0x2e, 0x3c, 0x21, 0xbe, 0x07, 0xdf, 0x00, 0xed, 0xac, 0x9d, 0x3f, 0xbd, 0x26,
	0xf7, 0x74, 0x2f, 0xc9, 0xcc, 0xfc, 0x76, 0x66, 0x7f, 0x33, 0x3b, 0x99, 0x09, 0xcc, 0xd2, 0x62,
	0xcc, 0xae, 0x65, 0x21, 0x78, 0x25, 0xc6, 0x85, 0x4c, 0x27, 0x59, 0xc5, 0x38, 0x7d, 0xb0, 0x24,
	0x9d, 0xe0, 0x5c, 0xa3, 0xcc, 0xe3, 0x6c, 0xc2, 0x64, 0xa1, 0x34, 0x2a, 0xcd, 0xf3, 0x74, 0x52,
	0xca, 0x42, 0x17, 0x13, 0x56, 0xe4, 0x09, 0x4f, 0x27, 0x3a, 0x96, 0x29, 0xea, 0xc8, 0x60, 0x91,
	0xc4, 0x5f, 0x2a, 0x2e, 0x51, 0x60, 0xae, 0x55, 0x64, 0xf1, 0x31, 0x1d, 0xf6, 0xbd, 0x35, 0xff,
	0xd1, 0xdf, 0x0e, 0xc0, 0x37, 0x0c, 0xaf, 0x50, 0xe9, 0x8b, 0x24, 0xf5, 0xbf, 0x80, 0x7e, 0xca,
	0x90, 0x22, 0x04, 0xce, 0x71, 0xeb, 0xc4, 0x3b, 0x7d, 0x3c, 0x5e, 0x3b, 0x3e, 0x5e, 0x1d, 0x6d,
	0xc4, 0xb0, 0x97, 0x5a, 0x61, 0xf8, 0xab, 0x03, 0xbd, 0xda, 0xe8, 0x7f, 0x04, 0x2e, 0xb1, 0xd0,
	0x8b, 0x12, 0x03, 0xe7, 0xd8, 0x39, 0x71, 0xc3, 0xbe, 0x31, 0x5c, 0x2d, 0x4a, 0xf4, 0x1f, 0x01,
	0x10, 0xa8, 0x2a, 0xae, 0x31, 0xd8, 0x23, 0x94, 0x8e, 0xbf, 0x32, 0x06, 0xff, 0x31, 0x78, 0x9a,
	0x0b, 0x2c, 0x2a, 0x1d, 0x29, 0x64, 0x41, 0xeb, 0xd8, 0x39, 0xe9, 0x84, 0x50, 0x9b, 0x5e, 0x21,
	0x33, 0xc1, 0x2b, 0x85, 0x11, 0x23, 0x96, 0xed, 0x63, 0xe7, 0xa4, 0x1f, 0xf6, 0x2b, 0x85, 0x17,
	0x46, 0x1f, 0xfd, 0xd7, 0x02, 0xf7, 0xf2, 0xb6, 0xc9, 0xe7, 0x0c, 0x7a, 0xd7, 0xb7, 0xeb, 0xe9,
	0x3c, 0xda, 0x48, 0x67, 0x79, 0xb0, 0x96, 0xc2, 0xee, 0x35, 0x7d, 0x0f, 0xff, 0x6a, 0x41, 0xd7,
	0x9a, 0xfc, 0xf7, 0xa0, 0x63, 0x89, 0xda, 0x34, 0xac, 0x72, 0x97, 0xe4, 0xde, 0x6b, 0x24, 0x7d,
	0x68, 0x97, 0x45, 0x91, 0x11, 0x7d, 0x37, 0x24, 0xd9, 0x1f, 0x42, 0x7f, 0x9a, 0x15, 0xec, 0x67,
	0x9e, 0xa7, 0x0d, 0xef, 0x46, 0x37, 0xd7, 0xc4, 0x6a, 0x91, 0xb3, 0xa0, 0x43, 0x80, 0x55, 0x4c,
	0xaa, 0xb7, 0xb1, 0xcc, 0xa3, 0x22, 0xcf, 0x16, 0x41, 0xd7, 0xba, 0x18, 0xc3, 0xf7, 0x79, 0xb6,
	0x30, 0xe1, 0x98, 0xe4, 0x9a, 0xb3, 0x38, 0x0b, 0x7a, 0x16, 0x6b, 0x74, 0x83, 0x95, 0x92, 0x17,
	0x92, 0xeb, 0x45, 0xd0, 0xb7, 0xf5, 0x6f, 0x74, 0x13, 0x34, 0xe1, 0x19, 0x46, 0xd3, 0x2a, 0x55,
	0x81, 0x6b, 0x1d, 0x8d, 0xe1, 0xbc, 0x4a, 0x95, 0xff, 0x04, 0x06, 0x82, 0xe7, 0x5c, 0x54, 0x22,
	0x9a, 0x55, 0x5a, 0x05, 0x40, 0x99, 0x79, 0xb5, 0xed, 0x79, 0xa5, 0x95, 0xa1, 0x2a, 0x51, 0xcb,
	0x45, 0xe0, 0x59, 0xaa, 0xa4, 0x98, 0x8a, 0x88, 0x78, 0x1e, 0x19, 0x85, 0xa3, 0x0a, 0x06, 0xb6,
	0x22, 0x22, 0x9e, 0x87, 0xd6, 0xe2, 0x7f, 0x0c, 0x07, 0x54, 0xbb, 0x48, 0xf0, 0xdc, 0xc6, 0xde,
	0xa7, 0x33, 0x03, 0xb2, 0xbe, 0xe4, 0x39, 0x05, 0x3f, 0x85, 0xf7, 0x8b, 0x24, 0xc9, 0x8a, 0x78,
	0x16, 0x25, 0x31, 0xcf, 0x2a, 0x89, 0xca, 0x66, 0x7f, 0x40, 0x97, 0xbd, 0x5b, 0x83, 0x5f, 0xd7,
	0x98, 0x29, 0xc4, 0xe8, 0x0f, 0x07, 0x0e, 0x5f, 0x16, 0xd3, 0x2c, 0x9e, 0xfe, 0x28, 0x9a, 0x97,
	0xbf, 0x04, 0x4f, 0x90, 0x69, 0xfd, 0xf5, 0x3f, 0xdd, 0x78, 0xfd, 0x3b, 0x2e, 0xb5, 0x4e, 0x7d,
	0x00, 0x62, 0x29, 0x0f, 0xbf, 0x05, 0x58, 0x21, 0xbb, 0x3b, 0xfb, 0x4d, 0x5d, 0x31, 0xfa, 0xb3,
	0x0d, 0xee, 0x8a, 0xe3, 0x19, 0xf4, 0x6e, 0xc4, 0xf6, 0xee, 0x5c, 0x31, 0xb3, 0x52, 0xd8, 0xbd,
	0xa1, 0x6f, 0xff, 0x02, 0x06, 0x3a, 0x56, 0x3a, 0x6a, 0x9c, 0xf7, 0xc8, 0xf9, 0xc9, 0x16, 0xe7,
	0xab, 0x58, 0xe9, 0x3a, 0x00, 0xe8, 0xa5, 0x3c, 0xfc, 0xd7, 0x81, 0xae, 0x15, 0xdf, 0xee, 0xaf,
	0x75, 0xd9, 0x2d, 0xed, 0x1d, 0xdd, 0xd2, 0x79, 0xad, 0x5b, 0x76, 0x76, 0xfe, 0xc6, 0x04, 0xe8,
	0x6d, 0x4e, 0x80, 0xe1, 0x33, 0x18, 0x98, 0x94, 0x4d, 0x66, 0x5f, 0xcd, 0x4b, 0xb9, 0xcc, 0x0e,
	0xe7, 0xa5, 0x5c, 0xcf, 0xce, 0x80, 0xc3, 0xdf, 0x1d, 0x80, 0x55, 0x81, 0x4c, 0xb2, 0xb6, 0x47,
	0xf3, 0x58, 0x34, 0xa5, 0x70, 0xc9, 0xf2, 0x5d, 0x2c, 0xd0, 0x7f, 0x01, 0x07, 0x54, 0xf8, 0x55,
	0x3c, 0x5b, 0xfa, 0xa7, 0x3b, 0x4a, 0xdf, 0xf0, 0x08, 0xe9, 0xcd, 0x96, 0xac, 0xde, 0x54, 0xb7,
	0x51, 0x09, 0xfb, 0xe7, 0x15, 0xcf, 0x66, 0x17, 0x92, 0x6b, 0x94, 0x3c, 0xf6, 0x9f, 0xc1, 0x3b,
	0x12, 0x13, 0x94, 0x98, 0x33, 0x8c, 0x66, 0xa8, 0x78, 0x9a, 0x5b, 0x86, 0x97, 0x0f, 0xc2, 0xc3,
	0x25, 0xf2, 0x9c, 0x00, 0xff, 0x29, 0x0c, 0xa6, 0xc6, 0x3b, 0xb2, 0x3b, 0xc1, 0xbe, 0xdb, 0xe5,
	0x83, 0xd0, 0x23, 0xeb, 0x15, 0x19, 0xcf, 0xf7, 0xc1, 0x6b, 0x56, 0xc6, 0xa2, 0xc4, 0xd1, 0x3f,
	0x7b, 0x70, 0xf4, 0x03, 0x4a, 0x0b, 0x52, 0xbf, 0xac, 0xed, 0x10, 0xff, 0x4b, 0x38, 0xb0, 0x11,
	0x59, 0x4d, 0x88, 0x2e, 0xf7, 0x4e, 0x87, 0x1b, 0xb9, 0x6f, 0x50, 0x0e, 0xf7, 0xa7, 0x1b, 0x19,
	0x7c, 0x0e, 0x83, 0x66, 0xbb, 0x44, 0x2c, 0x49, 0x89, 0x94, 0x77, 0xfa, 0xe1, 0x96, 0x0d, 0x13,
	0x42, 0xba, 0x5a, 0x4c, 0x67, 0xe0, 0xd5, 0x83, 0x9c, 0x3c, 0x5b, 0xe4, 0xf9, 0xc1, 0xfd, 0xc3,
	0x3c, 0x74, 0xaf, 0x97, 0x0b, 0xe0, 0x05, 0xf8, 0xf5, 0x18, 0xa8, 0x7f, 0x2c, 0xe4, 0xde, 0x26,
	0xf7, 0x87, 0xbb, 0xa6, 0x41, 0x78, 0x28, 0xee, 0x4c, 0x94, 0x33, 0xf0, 0xd6, 0x63, 0x74, 0xee,
	0xa1, 0xb0, 0xf2, 0x76, 0x6f, 0x1a, 0x71, 0xf4, 0x9b, 0x03, 0x47, 0xf7, 0xd7, 0xd4, 0x44, 0x4d,
	0xe1, 0x61, 0x89, 0x32, 0xda, 0xb6, 0xba, 0xeb, 0xc1, 0xf0, 0xc9, 0xc6, 0x35, 0x5b, 0x1f, 0x29,
	0x3c, 0x2a, 0xb7, 0x41, 0xe7, 0xfd, 0x9f, 0xba, 0xf6, 0x6f, 0xc0, 0xb4, 0x4b, 0xff, 0x03, 0x3e,
	0xfb, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xd1, 0xa2, 0xd7, 0x74, 0x6f, 0x08, 0x00, 0x00,
}
