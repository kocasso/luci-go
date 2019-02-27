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
	// Failure on HW tests warns only (does not generate error).
	WarnOnly bool `protobuf:"varint,3,opt,name=warn_only,json=warnOnly,proto3" json:"warn_only,omitempty"`
	// Usually we consider structural failures here as OK.
	Critical bool `protobuf:"varint,4,opt,name=critical,proto3" json:"critical,omitempty"`
	// Should we file bugs if a test fails in a suite run.
	FileBugs bool `protobuf:"varint,5,opt,name=file_bugs,json=fileBugs,proto3" json:"file_bugs,omitempty"`
	// Minimum number of DUTs required for testing in the hw lab.
	MinimumDuts int32 `protobuf:"varint,6,opt,name=minimum_duts,json=minimumDuts,proto3" json:"minimum_duts,omitempty"`
	// Whether we should retry tests that fail in a suite run.
	Retry bool `protobuf:"varint,7,opt,name=retry,proto3" json:"retry,omitempty"`
	// Maximum job retries allowed at suite level. 0 for no max.
	MaxRetries int32 `protobuf:"varint,8,opt,name=max_retries,json=maxRetries,proto3" json:"max_retries,omitempty"`
	// Preferred minimum duts. Lab will prioritize on getting such
	// number of duts even if the suite is competing with
	// other suites that have higher priority.
	SuiteMinDuts int32 `protobuf:"varint,9,opt,name=suite_min_duts,json=suiteMinDuts,proto3" json:"suite_min_duts,omitempty"`
	// Only offload failed tests to Google Storage.
	OffloadFailuresOnly  bool     `protobuf:"varint,10,opt,name=offload_failures_only,json=offloadFailuresOnly,proto3" json:"offload_failures_only,omitempty"`
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

type TastVmTestCfg struct {
	TastVmTest           []*TastVmTestCfg_TastVmTest `protobuf:"bytes,1,rep,name=tast_vm_test,json=tastVmTest,proto3" json:"tast_vm_test,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *TastVmTestCfg) Reset()         { *m = TastVmTestCfg{} }
func (m *TastVmTestCfg) String() string { return proto.CompactTextString(m) }
func (*TastVmTestCfg) ProtoMessage()    {}
func (*TastVmTestCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_efae681839d29e51, []int{3}
}

func (m *TastVmTestCfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TastVmTestCfg.Unmarshal(m, b)
}
func (m *TastVmTestCfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TastVmTestCfg.Marshal(b, m, deterministic)
}
func (m *TastVmTestCfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TastVmTestCfg.Merge(m, src)
}
func (m *TastVmTestCfg) XXX_Size() int {
	return xxx_messageInfo_TastVmTestCfg.Size(m)
}
func (m *TastVmTestCfg) XXX_DiscardUnknown() {
	xxx_messageInfo_TastVmTestCfg.DiscardUnknown(m)
}

var xxx_messageInfo_TastVmTestCfg proto.InternalMessageInfo

func (m *TastVmTestCfg) GetTastVmTest() []*TastVmTestCfg_TastVmTest {
	if m != nil {
		return m.TastVmTest
	}
	return nil
}

type TastVmTestCfg_TastTestExpr struct {
	// A single tast test expression. See https://goo.gl/UPNEgT
	TestExpr             string   `protobuf:"bytes,1,opt,name=test_expr,json=testExpr,proto3" json:"test_expr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TastVmTestCfg_TastTestExpr) Reset()         { *m = TastVmTestCfg_TastTestExpr{} }
func (m *TastVmTestCfg_TastTestExpr) String() string { return proto.CompactTextString(m) }
func (*TastVmTestCfg_TastTestExpr) ProtoMessage()    {}
func (*TastVmTestCfg_TastTestExpr) Descriptor() ([]byte, []int) {
	return fileDescriptor_efae681839d29e51, []int{3, 0}
}

func (m *TastVmTestCfg_TastTestExpr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TastVmTestCfg_TastTestExpr.Unmarshal(m, b)
}
func (m *TastVmTestCfg_TastTestExpr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TastVmTestCfg_TastTestExpr.Marshal(b, m, deterministic)
}
func (m *TastVmTestCfg_TastTestExpr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TastVmTestCfg_TastTestExpr.Merge(m, src)
}
func (m *TastVmTestCfg_TastTestExpr) XXX_Size() int {
	return xxx_messageInfo_TastVmTestCfg_TastTestExpr.Size(m)
}
func (m *TastVmTestCfg_TastTestExpr) XXX_DiscardUnknown() {
	xxx_messageInfo_TastVmTestCfg_TastTestExpr.DiscardUnknown(m)
}

var xxx_messageInfo_TastVmTestCfg_TastTestExpr proto.InternalMessageInfo

func (m *TastVmTestCfg_TastTestExpr) GetTestExpr() string {
	if m != nil {
		return m.TestExpr
	}
	return ""
}

type TastVmTestCfg_TastVmTest struct {
	// String containing short human-readable name describing test suite.
	SuiteName string `protobuf:"bytes,1,opt,name=suite_name,json=suiteName,proto3" json:"suite_name,omitempty"`
	// List of string expressions describing which tests to run; this
	// is passed directly to the 'tast run' command. See
	// https://goo.gl/UPNEgT for info about test expressions.
	TastTestExpr []*TastVmTestCfg_TastTestExpr `protobuf:"bytes,2,rep,name=tast_test_expr,json=tastTestExpr,proto3" json:"tast_test_expr,omitempty"`
	// Number of seconds to wait before timing out waiting for results.
	TimeoutSec           int32    `protobuf:"varint,3,opt,name=timeout_sec,json=timeoutSec,proto3" json:"timeout_sec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TastVmTestCfg_TastVmTest) Reset()         { *m = TastVmTestCfg_TastVmTest{} }
func (m *TastVmTestCfg_TastVmTest) String() string { return proto.CompactTextString(m) }
func (*TastVmTestCfg_TastVmTest) ProtoMessage()    {}
func (*TastVmTestCfg_TastVmTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_efae681839d29e51, []int{3, 1}
}

func (m *TastVmTestCfg_TastVmTest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TastVmTestCfg_TastVmTest.Unmarshal(m, b)
}
func (m *TastVmTestCfg_TastVmTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TastVmTestCfg_TastVmTest.Marshal(b, m, deterministic)
}
func (m *TastVmTestCfg_TastVmTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TastVmTestCfg_TastVmTest.Merge(m, src)
}
func (m *TastVmTestCfg_TastVmTest) XXX_Size() int {
	return xxx_messageInfo_TastVmTestCfg_TastVmTest.Size(m)
}
func (m *TastVmTestCfg_TastVmTest) XXX_DiscardUnknown() {
	xxx_messageInfo_TastVmTestCfg_TastVmTest.DiscardUnknown(m)
}

var xxx_messageInfo_TastVmTestCfg_TastVmTest proto.InternalMessageInfo

func (m *TastVmTestCfg_TastVmTest) GetSuiteName() string {
	if m != nil {
		return m.SuiteName
	}
	return ""
}

func (m *TastVmTestCfg_TastVmTest) GetTastTestExpr() []*TastVmTestCfg_TastTestExpr {
	if m != nil {
		return m.TastTestExpr
	}
	return nil
}

func (m *TastVmTestCfg_TastVmTest) GetTimeoutSec() int32 {
	if m != nil {
		return m.TimeoutSec
	}
	return 0
}

type VmTestCfg struct {
	VmTest               []*VmTestCfg_VmTest `protobuf:"bytes,1,rep,name=vm_test,json=vmTest,proto3" json:"vm_test,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *VmTestCfg) Reset()         { *m = VmTestCfg{} }
func (m *VmTestCfg) String() string { return proto.CompactTextString(m) }
func (*VmTestCfg) ProtoMessage()    {}
func (*VmTestCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_efae681839d29e51, []int{4}
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
	return fileDescriptor_efae681839d29e51, []int{4, 0}
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

// Specifies a CrOS target, either by reference design or by a specific build
// target.
type TargetCriteria struct {
	// Types that are valid to be assigned to TargetType:
	//	*TargetCriteria_ReferenceDesign
	//	*TargetCriteria_BuildTarget
	TargetType           isTargetCriteria_TargetType `protobuf_oneof:"target_type"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *TargetCriteria) Reset()         { *m = TargetCriteria{} }
func (m *TargetCriteria) String() string { return proto.CompactTextString(m) }
func (*TargetCriteria) ProtoMessage()    {}
func (*TargetCriteria) Descriptor() ([]byte, []int) {
	return fileDescriptor_efae681839d29e51, []int{5}
}

func (m *TargetCriteria) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetCriteria.Unmarshal(m, b)
}
func (m *TargetCriteria) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetCriteria.Marshal(b, m, deterministic)
}
func (m *TargetCriteria) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetCriteria.Merge(m, src)
}
func (m *TargetCriteria) XXX_Size() int {
	return xxx_messageInfo_TargetCriteria.Size(m)
}
func (m *TargetCriteria) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetCriteria.DiscardUnknown(m)
}

var xxx_messageInfo_TargetCriteria proto.InternalMessageInfo

type isTargetCriteria_TargetType interface {
	isTargetCriteria_TargetType()
}

type TargetCriteria_ReferenceDesign struct {
	ReferenceDesign string `protobuf:"bytes,1,opt,name=reference_design,json=referenceDesign,proto3,oneof"`
}

type TargetCriteria_BuildTarget struct {
	BuildTarget string `protobuf:"bytes,2,opt,name=build_target,json=buildTarget,proto3,oneof"`
}

func (*TargetCriteria_ReferenceDesign) isTargetCriteria_TargetType() {}

func (*TargetCriteria_BuildTarget) isTargetCriteria_TargetType() {}

func (m *TargetCriteria) GetTargetType() isTargetCriteria_TargetType {
	if m != nil {
		return m.TargetType
	}
	return nil
}

func (m *TargetCriteria) GetReferenceDesign() string {
	if x, ok := m.GetTargetType().(*TargetCriteria_ReferenceDesign); ok {
		return x.ReferenceDesign
	}
	return ""
}

func (m *TargetCriteria) GetBuildTarget() string {
	if x, ok := m.GetTargetType().(*TargetCriteria_BuildTarget); ok {
		return x.BuildTarget
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TargetCriteria) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TargetCriteria_ReferenceDesign)(nil),
		(*TargetCriteria_BuildTarget)(nil),
	}
}

// Details which testing is needed for a single CrOS build target.
type PerTargetTestRequirements struct {
	// Specifies the reference design or build target to which these testing
	// requirements should be applied.
	TargetCriteria *TargetCriteria `protobuf:"bytes,1,opt,name=target_criteria,json=targetCriteria,proto3" json:"target_criteria,omitempty"`
	// These configure what testing is needed for these BuildCriteria.
	GceTestCfg           *GceTestCfg      `protobuf:"bytes,2,opt,name=gce_test_cfg,json=gceTestCfg,proto3" json:"gce_test_cfg,omitempty"`
	HwTestCfg            *HwTestCfg       `protobuf:"bytes,3,opt,name=hw_test_cfg,json=hwTestCfg,proto3" json:"hw_test_cfg,omitempty"`
	MoblabVmTestCfg      *MoblabVmTestCfg `protobuf:"bytes,4,opt,name=moblab_vm_test_cfg,json=moblabVmTestCfg,proto3" json:"moblab_vm_test_cfg,omitempty"`
	TastVmTestCfg        *TastVmTestCfg   `protobuf:"bytes,6,opt,name=tast_vm_test_cfg,json=tastVmTestCfg,proto3" json:"tast_vm_test_cfg,omitempty"`
	VmTestCfg            *VmTestCfg       `protobuf:"bytes,5,opt,name=vm_test_cfg,json=vmTestCfg,proto3" json:"vm_test_cfg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PerTargetTestRequirements) Reset()         { *m = PerTargetTestRequirements{} }
func (m *PerTargetTestRequirements) String() string { return proto.CompactTextString(m) }
func (*PerTargetTestRequirements) ProtoMessage()    {}
func (*PerTargetTestRequirements) Descriptor() ([]byte, []int) {
	return fileDescriptor_efae681839d29e51, []int{6}
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

func (m *PerTargetTestRequirements) GetTargetCriteria() *TargetCriteria {
	if m != nil {
		return m.TargetCriteria
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

func (m *PerTargetTestRequirements) GetTastVmTestCfg() *TastVmTestCfg {
	if m != nil {
		return m.TastVmTestCfg
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
	return fileDescriptor_efae681839d29e51, []int{7}
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
	proto.RegisterType((*TastVmTestCfg)(nil), "crostesting.TastVmTestCfg")
	proto.RegisterType((*TastVmTestCfg_TastTestExpr)(nil), "crostesting.TastVmTestCfg.TastTestExpr")
	proto.RegisterType((*TastVmTestCfg_TastVmTest)(nil), "crostesting.TastVmTestCfg.TastVmTest")
	proto.RegisterType((*VmTestCfg)(nil), "crostesting.VmTestCfg")
	proto.RegisterType((*VmTestCfg_VmTest)(nil), "crostesting.VmTestCfg.VmTest")
	proto.RegisterType((*TargetCriteria)(nil), "crostesting.TargetCriteria")
	proto.RegisterType((*PerTargetTestRequirements)(nil), "crostesting.PerTargetTestRequirements")
	proto.RegisterType((*TargetTestRequirementsCfg)(nil), "crostesting.TargetTestRequirementsCfg")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/lucicfg/external/crostesting/proto/config/target_test_requirements_config.proto", fileDescriptor_efae681839d29e51)
}

var fileDescriptor_efae681839d29e51 = []byte{
	// 877 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x4b, 0x6f, 0xdb, 0x46,
	0x10, 0x0e, 0x25, 0xeb, 0x35, 0x94, 0xed, 0x80, 0x7d, 0xc9, 0x4a, 0x82, 0xb8, 0xea, 0xcb, 0x40,
	0x00, 0x09, 0x70, 0x01, 0x03, 0x6d, 0x6f, 0xb6, 0x53, 0xbb, 0x05, 0xdc, 0x16, 0x8c, 0xd1, 0x43,
	0x2f, 0x04, 0x45, 0x0d, 0xe9, 0x05, 0xb8, 0xa4, 0xb2, 0xbb, 0xb4, 0xa5, 0x73, 0x7b, 0xef, 0x0f,
	0x68, 0x6e, 0x3d, 0xf4, 0xb7, 0xf4, 0xd2, 0x63, 0x7f, 0x4f, 0xb1, 0xb3, 0xa4, 0x48, 0x2a, 0x92,
	0x7d, 0x28, 0x72, 0xb1, 0x76, 0xe7, 0xb5, 0xdf, 0xb7, 0xc3, 0xf9, 0xd6, 0x30, 0x8b, 0xd2, 0x71,
	0x70, 0x23, 0x52, 0xce, 0x32, 0x3e, 0x4e, 0x45, 0x34, 0x89, 0xb3, 0x80, 0xd1, 0x9f, 0x20, 0x8c,
	0x26, 0xb8, 0x50, 0x28, 0x12, 0x3f, 0x9e, 0x04, 0x22, 0x95, 0x0a, 0xa5, 0x62, 0x49, 0x34, 0x99,
	0x8b, 0x54, 0xa5, 0x93, 0x20, 0x4d, 0x42, 0x16, 0x4d, 0x94, 0x2f, 0x22, 0x54, 0x9e, 0xf6, 0x79,
	0x02, 0x5f, 0x67, 0x4c, 0x20, 0xc7, 0x44, 0x49, 0xcf, 0xf8, 0xc7, 0x14, 0xec, 0xd8, 0x95, 0xfc,
	0xd1, 0x3f, 0x16, 0xc0, 0x45, 0x80, 0xd7, 0x28, 0xd5, 0x59, 0x18, 0x39, 0x5f, 0x43, 0x37, 0x0a,
	0x90, 0x2a, 0x0c, 0xac, 0xc3, 0xe6, 0x91, 0x7d, 0xfc, 0x7c, 0x5c, 0x09, 0x1f, 0x97, 0xa1, 0xc5,
	0xd2, 0xed, 0x44, 0x66, 0x31, 0xfc, 0xd5, 0x82, 0x4e, 0x6e, 0x74, 0x9e, 0x40, 0x8f, 0x50, 0xa8,
	0xe5, 0x1c, 0x07, 0xd6, 0xa1, 0x75, 0xd4, 0x73, 0xbb, 0xda, 0x70, 0xbd, 0x9c, 0xa3, 0xf3, 0x0c,
	0x80, 0x9c, 0x32, 0x63, 0x0a, 0x07, 0x0d, 0xf2, 0x52, 0xf8, 0x2b, 0x6d, 0x70, 0x9e, 0x83, 0xad,
	0x18, 0xc7, 0x34, 0x53, 0x9e, 0xc4, 0x60, 0xd0, 0x3c, 0xb4, 0x8e, 0x5a, 0x2e, 0xe4, 0xa6, 0x57,
	0x18, 0xe8, 0xe2, 0x99, 0x44, 0x2f, 0x20, 0x94, 0x3b, 0x87, 0xd6, 0x51, 0xd7, 0xed, 0x66, 0x12,
	0xcf, 0xf4, 0x7e, 0xf4, 0x7b, 0x13, 0x7a, 0x97, 0x77, 0x05, 0x9f, 0x13, 0xe8, 0xdc, 0xdc, 0x55,
	0xe9, 0x3c, 0xab, 0xd1, 0x59, 0x05, 0xe6, 0x2b, 0xb7, 0x7d, 0x43, 0xbf, 0xc3, 0xbf, 0x1b, 0xd0,
	0x36, 0x26, 0xe7, 0x7d, 0x68, 0x19, 0xa0, 0x86, 0x86, 0xd9, 0xac, 0x83, 0x6c, 0x6c, 0x02, 0x79,
	0xe7, 0x8b, 0xc4, 0x4b, 0x93, 0x78, 0x49, 0x1c, 0xba, 0x6e, 0x57, 0x1b, 0x7e, 0x4c, 0xe2, 0xa5,
	0x33, 0x84, 0x6e, 0x20, 0x98, 0x62, 0x81, 0x1f, 0x17, 0x04, 0x8a, 0xbd, 0x4e, 0x0c, 0x59, 0x8c,
	0xde, 0x34, 0x8b, 0xe4, 0xa0, 0x65, 0x9c, 0xda, 0x70, 0x9a, 0x45, 0xd2, 0xf9, 0x18, 0xfa, 0x9c,
	0x25, 0x8c, 0x67, 0xdc, 0x9b, 0x65, 0x4a, 0x0e, 0xda, 0x74, 0xae, 0x9d, 0xdb, 0xce, 0x33, 0x25,
	0x35, 0x5e, 0x81, 0x4a, 0x2c, 0x07, 0x1d, 0xca, 0x35, 0x1b, 0x8d, 0x97, 0xfb, 0x0b, 0x4f, 0x6f,
	0x18, 0xca, 0x41, 0xd7, 0xe0, 0xe5, 0xfe, 0xc2, 0x35, 0x16, 0xe7, 0x53, 0xd8, 0x23, 0x66, 0x1e,
	0x67, 0x89, 0xa9, 0xdd, 0xa3, 0x98, 0x3e, 0x59, 0xaf, 0x58, 0x42, 0xc5, 0x8f, 0xe1, 0x83, 0x34,
	0x0c, 0xe3, 0xd4, 0x9f, 0x79, 0xa1, 0xcf, 0xe2, 0x4c, 0xa0, 0x34, 0x0c, 0x81, 0x0e, 0x7b, 0x2f,
	0x77, 0x7e, 0x9b, 0xfb, 0x34, 0xd9, 0xd1, 0x5f, 0x16, 0xec, 0x5f, 0xa5, 0xd3, 0xd8, 0x9f, 0xfe,
	0xcc, 0x8b, 0xbe, 0x5c, 0x82, 0xcd, 0xc9, 0x54, 0xed, 0xcd, 0x17, 0xb5, 0xde, 0xac, 0xa5, 0xe4,
	0x7b, 0xea, 0x12, 0xf0, 0xd5, 0x7a, 0xf8, 0x3d, 0x40, 0xe9, 0xb9, 0xff, 0xbb, 0x7b, 0xa8, 0x67,
	0xa3, 0x3f, 0x1b, 0xb0, 0x7b, 0xed, 0x4b, 0x55, 0xe2, 0xbc, 0x80, 0xbe, 0xf2, 0xa5, 0xf2, 0x6e,
	0x79, 0x15, 0xe8, 0x67, 0x35, 0xa0, 0xb5, 0x8c, 0xca, 0xce, 0x05, 0xb5, 0x5a, 0x0f, 0x5f, 0x40,
	0x5f, 0x7b, 0xf4, 0xfa, 0xe5, 0x62, 0x2e, 0x56, 0x40, 0x71, 0x31, 0x17, 0x55, 0xa0, 0xda, 0x39,
	0xfc, 0xc3, 0x02, 0x28, 0xeb, 0xe8, 0x79, 0x31, 0xad, 0x49, 0x7c, 0x5e, 0xb0, 0xea, 0x91, 0xe5,
	0x07, 0x9f, 0xa3, 0x73, 0x05, 0x7b, 0x84, 0xb1, 0xac, 0xd7, 0xd8, 0x70, 0x9d, 0x6f, 0xa3, 0x2c,
	0xb0, 0xb8, 0x44, 0x71, 0x85, 0xec, 0xa1, 0xf1, 0x1b, 0xbd, 0x69, 0x40, 0xaf, 0xbc, 0xa1, 0x13,
	0xe8, 0xd4, 0x2f, 0xa7, 0x3e, 0x61, 0xe5, 0x91, 0xf9, 0xa5, 0xb4, 0x6f, 0xcd, 0x85, 0xfc, 0x6b,
	0x41, 0x3b, 0xe7, 0xf7, 0x4e, 0xc5, 0x62, 0x35, 0x0e, 0x3b, 0xf7, 0x8c, 0x43, 0xeb, 0xad, 0x71,
	0xa8, 0x8d, 0x6f, 0x7b, 0x6d, 0x7c, 0x6b, 0x02, 0xd4, 0x59, 0x13, 0xa0, 0xd7, 0xb0, 0x77, 0x4d,
	0x3a, 0x7c, 0x26, 0x98, 0x42, 0xc1, 0x7c, 0xe7, 0x05, 0x3c, 0x16, 0x18, 0xa2, 0xc0, 0x24, 0x40,
	0x6f, 0x86, 0x92, 0x45, 0x89, 0xa1, 0x79, 0xf9, 0xc8, 0xdd, 0x5f, 0x79, 0xce, 0xc9, 0xe1, 0x7c,
	0x02, 0xfd, 0x69, 0xc6, 0xe2, 0x99, 0x67, 0xc4, 0xdc, 0x30, 0xbe, 0x7c, 0xe4, 0xda, 0x64, 0x35,
	0x95, 0x4f, 0x77, 0xc1, 0x2e, 0xb4, 0x7e, 0x39, 0xc7, 0xd1, 0x9b, 0x26, 0x1c, 0xfc, 0x84, 0xc2,
	0x38, 0xe9, 0x96, 0x2b, 0xe2, 0xef, 0x9c, 0xc3, 0x7e, 0x1e, 0x1c, 0xe4, 0x88, 0xe8, 0x74, 0xfb,
	0xf8, 0xc9, 0xda, 0x07, 0x52, 0x05, 0xed, 0xee, 0xa9, 0x3a, 0x89, 0xaf, 0xa0, 0x5f, 0xbc, 0x0c,
	0x5e, 0x10, 0x46, 0x84, 0xcb, 0x3e, 0xfe, 0x68, 0xcb, 0xeb, 0xe0, 0x42, 0x54, 0x3e, 0x2a, 0x27,
	0x60, 0xe7, 0x22, 0x4c, 0x99, 0x4d, 0xca, 0xfc, 0x70, 0xb3, 0x10, 0xbb, 0xbd, 0x9b, 0x95, 0x78,
	0x7f, 0x07, 0x4e, 0x2e, 0x12, 0xf9, 0x17, 0x46, 0xe9, 0x3b, 0x94, 0xfe, 0xf4, 0x3e, 0xad, 0x70,
	0xf7, 0xf9, 0x9a, 0xde, 0x9c, 0xc1, 0xe3, 0xea, 0x1c, 0x53, 0xa1, 0x36, 0x15, 0x1a, 0x6e, 0x9f,
	0x12, 0x77, 0x57, 0xd5, 0xc4, 0xe0, 0x04, 0xec, 0x6a, 0x7e, 0x6b, 0x03, 0x8f, 0x32, 0xb7, 0x77,
	0x5b, 0x2c, 0x47, 0xbf, 0x59, 0x70, 0xb0, 0xb9, 0x37, 0xba, 0x6a, 0x04, 0x4f, 0xe7, 0x28, 0xbc,
	0x6d, 0x6f, 0x77, 0x3e, 0x55, 0x9f, 0xd7, 0x8e, 0xd9, 0xda, 0x6c, 0xf7, 0x60, 0xbe, 0xcd, 0x75,
	0x7a, 0xf1, 0xcb, 0xcb, 0xff, 0xf7, 0xff, 0xc5, 0x37, 0xe6, 0x67, 0xda, 0x26, 0xe3, 0x97, 0xff,
	0x05, 0x00, 0x00, 0xff, 0xff, 0x97, 0x2a, 0x9b, 0x2e, 0xaf, 0x08, 0x00, 0x00,
}
