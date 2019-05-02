// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/gce/api/config/v1/config.proto

package config

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	dayofweek "google.golang.org/genproto/googleapis/type/dayofweek"
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

// A network access config type.
// GCE's default and only value is 1:1 NAT.
type AccessConfigType int32

const (
	// 1:1 network address translation.
	AccessConfigType_ONE_TO_ONE_NAT AccessConfigType = 0
)

var AccessConfigType_name = map[int32]string{
	0: "ONE_TO_ONE_NAT",
}

var AccessConfigType_value = map[string]int32{
	"ONE_TO_ONE_NAT": 0,
}

func (x AccessConfigType) String() string {
	return proto.EnumName(AccessConfigType_name, int32(x))
}

func (AccessConfigType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_56a5743e153a90ec, []int{0}
}

// A description of a service account.
type ServiceAccount struct {
	// The email address of this service account.
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	// The scopes available for this service account.
	Scope                []string `protobuf:"bytes,2,rep,name=scope,proto3" json:"scope,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceAccount) Reset()         { *m = ServiceAccount{} }
func (m *ServiceAccount) String() string { return proto.CompactTextString(m) }
func (*ServiceAccount) ProtoMessage()    {}
func (*ServiceAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a5743e153a90ec, []int{0}
}

func (m *ServiceAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceAccount.Unmarshal(m, b)
}
func (m *ServiceAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceAccount.Marshal(b, m, deterministic)
}
func (m *ServiceAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceAccount.Merge(m, src)
}
func (m *ServiceAccount) XXX_Size() int {
	return xxx_messageInfo_ServiceAccount.Size(m)
}
func (m *ServiceAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceAccount proto.InternalMessageInfo

func (m *ServiceAccount) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ServiceAccount) GetScope() []string {
	if m != nil {
		return m.Scope
	}
	return nil
}

// A description of a network access config.
type AccessConfig struct {
	// The type of config this is.
	Type                 AccessConfigType `protobuf:"varint,1,opt,name=type,proto3,enum=config.AccessConfigType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *AccessConfig) Reset()         { *m = AccessConfig{} }
func (m *AccessConfig) String() string { return proto.CompactTextString(m) }
func (*AccessConfig) ProtoMessage()    {}
func (*AccessConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a5743e153a90ec, []int{1}
}

func (m *AccessConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessConfig.Unmarshal(m, b)
}
func (m *AccessConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessConfig.Marshal(b, m, deterministic)
}
func (m *AccessConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessConfig.Merge(m, src)
}
func (m *AccessConfig) XXX_Size() int {
	return xxx_messageInfo_AccessConfig.Size(m)
}
func (m *AccessConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AccessConfig proto.InternalMessageInfo

func (m *AccessConfig) GetType() AccessConfigType {
	if m != nil {
		return m.Type
	}
	return AccessConfigType_ONE_TO_ONE_NAT
}

// A description of a network interface.
type NetworkInterface struct {
	// The access configurations for this interface.
	// Required to enable external internet access.
	AccessConfig []*AccessConfig `protobuf:"bytes,1,rep,name=access_config,json=accessConfig,proto3" json:"access_config,omitempty"`
	// The name of a network to use for this interface.
	// https://cloud.google.com/compute/docs/reference/rest/v1/networks/list.
	Network              string   `protobuf:"bytes,2,opt,name=network,proto3" json:"network,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkInterface) Reset()         { *m = NetworkInterface{} }
func (m *NetworkInterface) String() string { return proto.CompactTextString(m) }
func (*NetworkInterface) ProtoMessage()    {}
func (*NetworkInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a5743e153a90ec, []int{2}
}

func (m *NetworkInterface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInterface.Unmarshal(m, b)
}
func (m *NetworkInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInterface.Marshal(b, m, deterministic)
}
func (m *NetworkInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInterface.Merge(m, src)
}
func (m *NetworkInterface) XXX_Size() int {
	return xxx_messageInfo_NetworkInterface.Size(m)
}
func (m *NetworkInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInterface.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInterface proto.InternalMessageInfo

func (m *NetworkInterface) GetAccessConfig() []*AccessConfig {
	if m != nil {
		return m.AccessConfig
	}
	return nil
}

func (m *NetworkInterface) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

// A description of a disk.
// https://cloud.google.com/compute/docs/reference/rest/v1/disks.
type Disk struct {
	// The name of an image to use to create this disk.
	// https://cloud.google.com/compute/docs/reference/rest/v1/images/list.
	Image string `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	// The size of this disk in GiB.
	Size int64 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	// The name of a disk type to use for this disk.
	// https://cloud.google.com/compute/docs/reference/rest/v1/diskTypes/list.
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Disk) Reset()         { *m = Disk{} }
func (m *Disk) String() string { return proto.CompactTextString(m) }
func (*Disk) ProtoMessage()    {}
func (*Disk) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a5743e153a90ec, []int{3}
}

func (m *Disk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Disk.Unmarshal(m, b)
}
func (m *Disk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Disk.Marshal(b, m, deterministic)
}
func (m *Disk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Disk.Merge(m, src)
}
func (m *Disk) XXX_Size() int {
	return xxx_messageInfo_Disk.Size(m)
}
func (m *Disk) XXX_DiscardUnknown() {
	xxx_messageInfo_Disk.DiscardUnknown(m)
}

var xxx_messageInfo_Disk proto.InternalMessageInfo

func (m *Disk) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *Disk) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Disk) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

// A description of instance metadata.
type Metadata struct {
	// Types that are valid to be assigned to Metadata:
	//	*Metadata_FromText
	//	*Metadata_FromFile
	Metadata             isMetadata_Metadata `protobuf_oneof:"metadata"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a5743e153a90ec, []int{4}
}

func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (m *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(m, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

type isMetadata_Metadata interface {
	isMetadata_Metadata()
}

type Metadata_FromText struct {
	FromText string `protobuf:"bytes,1,opt,name=from_text,json=fromText,proto3,oneof"`
}

type Metadata_FromFile struct {
	FromFile string `protobuf:"bytes,2,opt,name=from_file,json=fromFile,proto3,oneof"`
}

func (*Metadata_FromText) isMetadata_Metadata() {}

func (*Metadata_FromFile) isMetadata_Metadata() {}

func (m *Metadata) GetMetadata() isMetadata_Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Metadata) GetFromText() string {
	if x, ok := m.GetMetadata().(*Metadata_FromText); ok {
		return x.FromText
	}
	return ""
}

func (m *Metadata) GetFromFile() string {
	if x, ok := m.GetMetadata().(*Metadata_FromFile); ok {
		return x.FromFile
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Metadata) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Metadata_FromText)(nil),
		(*Metadata_FromFile)(nil),
	}
}

// A description of a VM.
// https://cloud.google.com/compute/docs/reference/rest/v1/instances.
type VM struct {
	// The disks to attach to this VM.
	Disk []*Disk `protobuf:"bytes,1,rep,name=disk,proto3" json:"disk,omitempty"`
	// The name of a machine type to use for this VM.
	// https://cloud.google.com/compute/docs/reference/rest/v1/machineTypes/list.
	MachineType string `protobuf:"bytes,2,opt,name=machine_type,json=machineType,proto3" json:"machine_type,omitempty"`
	// The metadata to attach to this VM.
	Metadata []*Metadata `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty"`
	// The minimum CPU platform to use for this VM.
	// https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform.
	MinCpuPlatform string `protobuf:"bytes,4,opt,name=min_cpu_platform,json=minCpuPlatform,proto3" json:"min_cpu_platform,omitempty"`
	// The network interfaces to configure for this VM.
	NetworkInterface []*NetworkInterface `protobuf:"bytes,5,rep,name=network_interface,json=networkInterface,proto3" json:"network_interface,omitempty"`
	// The name of a GCP project to create this VM in.
	Project string `protobuf:"bytes,6,opt,name=project,proto3" json:"project,omitempty"`
	// The service accounts to make available to this VM.
	ServiceAccount []*ServiceAccount `protobuf:"bytes,7,rep,name=service_account,json=serviceAccount,proto3" json:"service_account,omitempty"`
	// The tags to attach to this VM.
	Tag []string `protobuf:"bytes,8,rep,name=tag,proto3" json:"tag,omitempty"`
	// The name of a zone to create this VM in.
	// https://cloud.google.com/compute/docs/reference/rest/v1/zones/list.
	Zone                 string   `protobuf:"bytes,9,opt,name=zone,proto3" json:"zone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VM) Reset()         { *m = VM{} }
func (m *VM) String() string { return proto.CompactTextString(m) }
func (*VM) ProtoMessage()    {}
func (*VM) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a5743e153a90ec, []int{5}
}

func (m *VM) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VM.Unmarshal(m, b)
}
func (m *VM) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VM.Marshal(b, m, deterministic)
}
func (m *VM) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VM.Merge(m, src)
}
func (m *VM) XXX_Size() int {
	return xxx_messageInfo_VM.Size(m)
}
func (m *VM) XXX_DiscardUnknown() {
	xxx_messageInfo_VM.DiscardUnknown(m)
}

var xxx_messageInfo_VM proto.InternalMessageInfo

func (m *VM) GetDisk() []*Disk {
	if m != nil {
		return m.Disk
	}
	return nil
}

func (m *VM) GetMachineType() string {
	if m != nil {
		return m.MachineType
	}
	return ""
}

func (m *VM) GetMetadata() []*Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *VM) GetMinCpuPlatform() string {
	if m != nil {
		return m.MinCpuPlatform
	}
	return ""
}

func (m *VM) GetNetworkInterface() []*NetworkInterface {
	if m != nil {
		return m.NetworkInterface
	}
	return nil
}

func (m *VM) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *VM) GetServiceAccount() []*ServiceAccount {
	if m != nil {
		return m.ServiceAccount
	}
	return nil
}

func (m *VM) GetTag() []string {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *VM) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

// A length of time.
type TimePeriod struct {
	// Types that are valid to be assigned to Time:
	//	*TimePeriod_Duration
	//	*TimePeriod_Seconds
	Time                 isTimePeriod_Time `protobuf_oneof:"time"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TimePeriod) Reset()         { *m = TimePeriod{} }
func (m *TimePeriod) String() string { return proto.CompactTextString(m) }
func (*TimePeriod) ProtoMessage()    {}
func (*TimePeriod) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a5743e153a90ec, []int{6}
}

func (m *TimePeriod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimePeriod.Unmarshal(m, b)
}
func (m *TimePeriod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimePeriod.Marshal(b, m, deterministic)
}
func (m *TimePeriod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimePeriod.Merge(m, src)
}
func (m *TimePeriod) XXX_Size() int {
	return xxx_messageInfo_TimePeriod.Size(m)
}
func (m *TimePeriod) XXX_DiscardUnknown() {
	xxx_messageInfo_TimePeriod.DiscardUnknown(m)
}

var xxx_messageInfo_TimePeriod proto.InternalMessageInfo

type isTimePeriod_Time interface {
	isTimePeriod_Time()
}

type TimePeriod_Duration struct {
	Duration string `protobuf:"bytes,1,opt,name=duration,proto3,oneof"`
}

type TimePeriod_Seconds struct {
	Seconds int64 `protobuf:"varint,2,opt,name=seconds,proto3,oneof"`
}

func (*TimePeriod_Duration) isTimePeriod_Time() {}

func (*TimePeriod_Seconds) isTimePeriod_Time() {}

func (m *TimePeriod) GetTime() isTimePeriod_Time {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *TimePeriod) GetDuration() string {
	if x, ok := m.GetTime().(*TimePeriod_Duration); ok {
		return x.Duration
	}
	return ""
}

func (m *TimePeriod) GetSeconds() int64 {
	if x, ok := m.GetTime().(*TimePeriod_Seconds); ok {
		return x.Seconds
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TimePeriod) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TimePeriod_Duration)(nil),
		(*TimePeriod_Seconds)(nil),
	}
}

// A time of day.
type TimeOfDay struct {
	// The day of the week the time applies to.
	Day dayofweek.DayOfWeek `protobuf:"varint,1,opt,name=day,proto3,enum=google.type.DayOfWeek" json:"day,omitempty"`
	// The location the time should be interpreted in.
	// https://en.wikipedia.org/wiki/List_of_tz_database_time_zones.
	Location string `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	// The time in 24-hour <hour>:<minute>.
	Time                 string   `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeOfDay) Reset()         { *m = TimeOfDay{} }
func (m *TimeOfDay) String() string { return proto.CompactTextString(m) }
func (*TimeOfDay) ProtoMessage()    {}
func (*TimeOfDay) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a5743e153a90ec, []int{7}
}

func (m *TimeOfDay) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeOfDay.Unmarshal(m, b)
}
func (m *TimeOfDay) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeOfDay.Marshal(b, m, deterministic)
}
func (m *TimeOfDay) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeOfDay.Merge(m, src)
}
func (m *TimeOfDay) XXX_Size() int {
	return xxx_messageInfo_TimeOfDay.Size(m)
}
func (m *TimeOfDay) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeOfDay.DiscardUnknown(m)
}

var xxx_messageInfo_TimeOfDay proto.InternalMessageInfo

func (m *TimeOfDay) GetDay() dayofweek.DayOfWeek {
	if m != nil {
		return m.Day
	}
	return dayofweek.DayOfWeek_DAY_OF_WEEK_UNSPECIFIED
}

func (m *TimeOfDay) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *TimeOfDay) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

// An amount of VMs for particular days of the week.
type Schedule struct {
	// The amount of VMs.
	Amount int32 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// The length of time the amount is in effect.
	// With start, this creates a half-open interval.
	// During [start, start+length) the amount will apply.
	Length *TimePeriod `protobuf:"bytes,2,opt,name=length,proto3" json:"length,omitempty"`
	// The start times when this amount goes into effect.
	Start                *TimeOfDay `protobuf:"bytes,3,opt,name=start,proto3" json:"start,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Schedule) Reset()         { *m = Schedule{} }
func (m *Schedule) String() string { return proto.CompactTextString(m) }
func (*Schedule) ProtoMessage()    {}
func (*Schedule) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a5743e153a90ec, []int{8}
}

func (m *Schedule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Schedule.Unmarshal(m, b)
}
func (m *Schedule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Schedule.Marshal(b, m, deterministic)
}
func (m *Schedule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Schedule.Merge(m, src)
}
func (m *Schedule) XXX_Size() int {
	return xxx_messageInfo_Schedule.Size(m)
}
func (m *Schedule) XXX_DiscardUnknown() {
	xxx_messageInfo_Schedule.DiscardUnknown(m)
}

var xxx_messageInfo_Schedule proto.InternalMessageInfo

func (m *Schedule) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Schedule) GetLength() *TimePeriod {
	if m != nil {
		return m.Length
	}
	return nil
}

func (m *Schedule) GetStart() *TimeOfDay {
	if m != nil {
		return m.Start
	}
	return nil
}

// An amount of VMs.
type Amount struct {
	// The default amount to use outside scheduled hours.
	Default int32 `protobuf:"varint,1,opt,name=default,proto3" json:"default,omitempty"`
	// The amount of VMs to have at a particular time.
	Change               []*Schedule `protobuf:"bytes,2,rep,name=change,proto3" json:"change,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Amount) Reset()         { *m = Amount{} }
func (m *Amount) String() string { return proto.CompactTextString(m) }
func (*Amount) ProtoMessage()    {}
func (*Amount) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a5743e153a90ec, []int{9}
}

func (m *Amount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Amount.Unmarshal(m, b)
}
func (m *Amount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Amount.Marshal(b, m, deterministic)
}
func (m *Amount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Amount.Merge(m, src)
}
func (m *Amount) XXX_Size() int {
	return xxx_messageInfo_Amount.Size(m)
}
func (m *Amount) XXX_DiscardUnknown() {
	xxx_messageInfo_Amount.DiscardUnknown(m)
}

var xxx_messageInfo_Amount proto.InternalMessageInfo

func (m *Amount) GetDefault() int32 {
	if m != nil {
		return m.Default
	}
	return 0
}

func (m *Amount) GetChange() []*Schedule {
	if m != nil {
		return m.Change
	}
	return nil
}

// A config for one type of VM.
type Config struct {
	// The amount of these VMs.
	Amount *Amount `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	// The attributes of these VMs.
	Attributes *VM `protobuf:"bytes,2,opt,name=attributes,proto3" json:"attributes,omitempty"`
	// The lifetime of these VMs.
	// At the end of their lifetime, each VM is deleted and replaced.
	Lifetime *TimePeriod `protobuf:"bytes,3,opt,name=lifetime,proto3" json:"lifetime,omitempty"`
	// The prefix to use when naming these VMs.
	Prefix string `protobuf:"bytes,4,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// Should only be set by the server. The revision of this config.
	Revision string `protobuf:"bytes,5,opt,name=revision,proto3" json:"revision,omitempty"`
	// The hostname of the Swarming server these VMs should connect to.
	Swarming string `protobuf:"bytes,6,opt,name=swarming,proto3" json:"swarming,omitempty"`
	// The timeout of these VMs.
	// If no Swarming bot has connected by the timeout,
	// the VM is deleted and replaced.
	Timeout              *TimePeriod `protobuf:"bytes,7,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a5743e153a90ec, []int{10}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetAmount() *Amount {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *Config) GetAttributes() *VM {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Config) GetLifetime() *TimePeriod {
	if m != nil {
		return m.Lifetime
	}
	return nil
}

func (m *Config) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *Config) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *Config) GetSwarming() string {
	if m != nil {
		return m.Swarming
	}
	return ""
}

func (m *Config) GetTimeout() *TimePeriod {
	if m != nil {
		return m.Timeout
	}
	return nil
}

// A config for several types of VMs.
type Configs struct {
	// The configs for different types of VMs.
	Vms                  []*Config `protobuf:"bytes,1,rep,name=vms,proto3" json:"vms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Configs) Reset()         { *m = Configs{} }
func (m *Configs) String() string { return proto.CompactTextString(m) }
func (*Configs) ProtoMessage()    {}
func (*Configs) Descriptor() ([]byte, []int) {
	return fileDescriptor_56a5743e153a90ec, []int{11}
}

func (m *Configs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configs.Unmarshal(m, b)
}
func (m *Configs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configs.Marshal(b, m, deterministic)
}
func (m *Configs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configs.Merge(m, src)
}
func (m *Configs) XXX_Size() int {
	return xxx_messageInfo_Configs.Size(m)
}
func (m *Configs) XXX_DiscardUnknown() {
	xxx_messageInfo_Configs.DiscardUnknown(m)
}

var xxx_messageInfo_Configs proto.InternalMessageInfo

func (m *Configs) GetVms() []*Config {
	if m != nil {
		return m.Vms
	}
	return nil
}

func init() {
	proto.RegisterEnum("config.AccessConfigType", AccessConfigType_name, AccessConfigType_value)
	proto.RegisterType((*ServiceAccount)(nil), "config.ServiceAccount")
	proto.RegisterType((*AccessConfig)(nil), "config.AccessConfig")
	proto.RegisterType((*NetworkInterface)(nil), "config.NetworkInterface")
	proto.RegisterType((*Disk)(nil), "config.Disk")
	proto.RegisterType((*Metadata)(nil), "config.Metadata")
	proto.RegisterType((*VM)(nil), "config.VM")
	proto.RegisterType((*TimePeriod)(nil), "config.TimePeriod")
	proto.RegisterType((*TimeOfDay)(nil), "config.TimeOfDay")
	proto.RegisterType((*Schedule)(nil), "config.Schedule")
	proto.RegisterType((*Amount)(nil), "config.Amount")
	proto.RegisterType((*Config)(nil), "config.Config")
	proto.RegisterType((*Configs)(nil), "config.Configs")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/gce/api/config/v1/config.proto", fileDescriptor_56a5743e153a90ec)
}

var fileDescriptor_56a5743e153a90ec = []byte{
	// 810 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x5b, 0x6f, 0xdb, 0x36,
	0x14, 0x4e, 0x7c, 0x91, 0x9d, 0xe3, 0xcc, 0x73, 0x89, 0x22, 0x10, 0xb2, 0x0d, 0xf0, 0xf4, 0xd0,
	0x19, 0x5d, 0x20, 0x63, 0xee, 0xd3, 0x80, 0x02, 0x43, 0xda, 0x74, 0xe8, 0x80, 0x25, 0x29, 0x54,
	0xa3, 0x7b, 0x14, 0x58, 0xea, 0x48, 0xe6, 0x2c, 0x89, 0x82, 0x44, 0x25, 0x71, 0x7f, 0xc2, 0xfe,
	0xe9, 0xfe, 0xc5, 0xc0, 0x9b, 0xec, 0x04, 0xdd, 0x93, 0x78, 0x2e, 0xfa, 0xce, 0xe5, 0xfb, 0x48,
	0x78, 0x95, 0x89, 0x90, 0x6d, 0x6a, 0x51, 0xf0, 0xb6, 0x08, 0x45, 0x9d, 0x2d, 0xf3, 0x96, 0xf1,
	0x65, 0xc6, 0x70, 0x49, 0x2b, 0xbe, 0x64, 0xa2, 0x4c, 0x79, 0xb6, 0xbc, 0xfb, 0xc5, 0x9e, 0xc2,
	0xaa, 0x16, 0x52, 0x10, 0xcf, 0x58, 0xe7, 0xdf, 0x65, 0x42, 0x64, 0x39, 0x2e, 0xe5, 0xae, 0xc2,
	0x65, 0x42, 0x77, 0x22, 0xbd, 0x47, 0xdc, 0x9a, 0xa4, 0xe0, 0x35, 0x4c, 0x3f, 0x62, 0x7d, 0xc7,
	0x19, 0x5e, 0x32, 0x26, 0xda, 0x52, 0x92, 0xe7, 0x30, 0xc4, 0x82, 0xf2, 0xdc, 0x3f, 0x9e, 0x1f,
	0x2f, 0x4e, 0x22, 0x63, 0x28, 0x6f, 0xc3, 0x44, 0x85, 0x7e, 0x6f, 0xde, 0x57, 0x5e, 0x6d, 0x04,
	0xaf, 0xe1, 0xf4, 0x92, 0x31, 0x6c, 0x9a, 0xb7, 0xba, 0x14, 0xb9, 0x80, 0x81, 0xaa, 0xa2, 0x7f,
	0x9d, 0xae, 0xfc, 0xd0, 0xf6, 0x73, 0x98, 0xb3, 0xde, 0x55, 0x18, 0xe9, 0xac, 0x20, 0x83, 0xd9,
	0x0d, 0xca, 0x7b, 0x51, 0x6f, 0xff, 0x28, 0x25, 0xd6, 0x29, 0x65, 0x48, 0x7e, 0x85, 0x6f, 0xa8,
	0xce, 0x8e, 0xcd, 0xbf, 0xfe, 0xf1, 0xbc, 0xbf, 0x98, 0xac, 0x9e, 0x7f, 0x0d, 0x2a, 0x3a, 0xa5,
	0x87, 0xc5, 0x7d, 0x18, 0x95, 0x06, 0xce, 0xef, 0xe9, 0xd6, 0x9d, 0x19, 0x5c, 0xc1, 0xe0, 0x8a,
	0x37, 0x5b, 0x35, 0x04, 0x2f, 0x68, 0x86, 0x6e, 0x34, 0x6d, 0x10, 0x02, 0x83, 0x86, 0x7f, 0x41,
	0xfd, 0x53, 0x3f, 0xd2, 0x67, 0xe5, 0xd3, 0x83, 0xf4, 0x75, 0xa2, 0x69, 0x77, 0x0d, 0xe3, 0x6b,
	0x94, 0x34, 0xa1, 0x92, 0x92, 0x1f, 0xe0, 0x24, 0xad, 0x45, 0x11, 0x4b, 0x7c, 0x90, 0x06, 0xed,
	0xfd, 0x51, 0x34, 0x56, 0xae, 0x35, 0x3e, 0xc8, 0x2e, 0x9c, 0xf2, 0xdc, 0xe0, 0x76, 0xe1, 0xdf,
	0x79, 0x8e, 0x6f, 0x00, 0xc6, 0x85, 0x45, 0x0a, 0xfe, 0xed, 0x41, 0xef, 0xd3, 0x35, 0x99, 0xc3,
	0x20, 0xe1, 0xcd, 0xd6, 0x8e, 0x7b, 0xea, 0xc6, 0x55, 0x6d, 0x47, 0x3a, 0x42, 0x7e, 0x84, 0xd3,
	0x82, 0xb2, 0x0d, 0x2f, 0x31, 0xd6, 0xad, 0x99, 0x19, 0x27, 0xd6, 0xa7, 0xd6, 0x4a, 0x2e, 0xf6,
	0xb8, 0x7e, 0x5f, 0x03, 0xcd, 0x1c, 0x90, 0xeb, 0x3c, 0xea, 0x32, 0xc8, 0x02, 0x66, 0x05, 0x2f,
	0x63, 0x56, 0xb5, 0x71, 0x95, 0x53, 0x99, 0x8a, 0xba, 0xf0, 0x07, 0x1a, 0x74, 0x5a, 0xf0, 0xf2,
	0x6d, 0xd5, 0x7e, 0xb0, 0x5e, 0xf2, 0x0e, 0x9e, 0xd9, 0x55, 0xc6, 0xdc, 0x31, 0xe5, 0x0f, 0x75,
	0x81, 0x8e, 0xe3, 0xa7, 0x4c, 0x46, 0xb3, 0xf2, 0x29, 0xb7, 0x3e, 0x8c, 0xaa, 0x5a, 0xfc, 0x8d,
	0x4c, 0xfa, 0x9e, 0x21, 0xc8, 0x9a, 0xe4, 0x37, 0xf8, 0xb6, 0x31, 0x2a, 0x8c, 0xa9, 0x91, 0xa1,
	0x3f, 0xd2, 0xf0, 0x67, 0x0e, 0xfe, 0xb1, 0x48, 0xa3, 0x69, 0xf3, 0x58, 0xb4, 0x33, 0xe8, 0x4b,
	0x9a, 0xf9, 0x63, 0x2d, 0x4e, 0x75, 0x54, 0x0c, 0x7e, 0x11, 0x25, 0xfa, 0x27, 0x86, 0x41, 0x75,
	0x0e, 0x6e, 0x00, 0xd6, 0xbc, 0xc0, 0x0f, 0x58, 0x73, 0x91, 0x90, 0xef, 0x61, 0x9c, 0xb4, 0x35,
	0x95, 0x5c, 0x94, 0x7b, 0x0a, 0x9d, 0x87, 0x9c, 0xc3, 0xa8, 0x41, 0x26, 0xca, 0xa4, 0x31, 0xc2,
	0x78, 0x7f, 0x14, 0x39, 0xc7, 0x1b, 0x0f, 0x06, 0x92, 0x17, 0x18, 0x20, 0x9c, 0x28, 0xbc, 0xdb,
	0xf4, 0x8a, 0xee, 0xc8, 0x02, 0xfa, 0x09, 0xdd, 0x59, 0xe9, 0x9f, 0x85, 0xe6, 0xd2, 0x85, 0x8a,
	0xaa, 0xf0, 0x8a, 0xee, 0x6e, 0xd3, 0xbf, 0x10, 0xb7, 0x91, 0x4a, 0x21, 0xe7, 0x30, 0xce, 0x05,
	0x33, 0x85, 0x0d, 0x8b, 0x9d, 0xad, 0x85, 0xc7, 0x8b, 0xbd, 0xf0, 0x54, 0x99, 0x7b, 0x18, 0x7f,
	0x64, 0x1b, 0x4c, 0xda, 0x1c, 0xc9, 0x19, 0x78, 0xb4, 0xd0, 0x0b, 0x52, 0x85, 0x86, 0x91, 0xb5,
	0xc8, 0x4b, 0xf0, 0x72, 0x2c, 0x33, 0xb9, 0xd1, 0x88, 0x93, 0x15, 0x71, 0x8b, 0xdb, 0x0f, 0x1c,
	0xd9, 0x0c, 0xf2, 0x13, 0x0c, 0x1b, 0x49, 0x6b, 0xa9, 0x8b, 0x4c, 0x56, 0xcf, 0x0e, 0x53, 0xf5,
	0x2c, 0x91, 0x89, 0x07, 0x7f, 0x82, 0x77, 0x69, 0xe0, 0x7d, 0x18, 0x25, 0x98, 0xd2, 0x36, 0x77,
	0x75, 0x9d, 0x49, 0x16, 0xe0, 0xb1, 0x0d, 0x2d, 0x33, 0xf3, 0x32, 0x1c, 0x28, 0xce, 0xb5, 0x1c,
	0xd9, 0x78, 0xf0, 0x4f, 0x0f, 0x3c, 0x7b, 0x55, 0x5f, 0x3c, 0x9a, 0x62, 0xb2, 0x9a, 0x76, 0xd7,
	0x5b, 0x7b, 0x0f, 0xa6, 0x02, 0x2a, 0x65, 0xcd, 0x3f, 0xb7, 0x12, 0x1b, 0x3b, 0x19, 0xb8, 0xdc,
	0x4f, 0xd7, 0xd1, 0x41, 0x94, 0x84, 0x30, 0xce, 0x79, 0x8a, 0xdd, 0xf6, 0xbe, 0xbe, 0x83, 0x2e,
	0x47, 0x6d, 0xb2, 0xaa, 0x31, 0xe5, 0x0f, 0x56, 0xf4, 0xd6, 0x52, 0xec, 0xd4, 0x78, 0xc7, 0x1b,
	0xc5, 0xce, 0xd0, 0xb0, 0xe3, 0x6c, 0x15, 0x6b, 0xee, 0x69, 0x5d, 0xf0, 0x32, 0xb3, 0x12, 0xee,
	0x6c, 0x72, 0x01, 0x23, 0x85, 0x2b, 0x5a, 0xa5, 0xdd, 0xff, 0x2b, 0xef, 0x52, 0x82, 0x9f, 0x61,
	0x64, 0x76, 0xd1, 0x90, 0x39, 0xf4, 0xef, 0x8a, 0xc6, 0xde, 0xfc, 0x6e, 0x13, 0xf6, 0x89, 0x53,
	0xa1, 0x97, 0x2f, 0x60, 0xf6, 0xf4, 0x09, 0x25, 0x04, 0xa6, 0xb7, 0x37, 0xef, 0xe2, 0xf5, 0x6d,
	0xac, 0x3e, 0x37, 0x97, 0xeb, 0xd9, 0xd1, 0x67, 0x4f, 0xbf, 0xe9, 0xaf, 0xfe, 0x0b, 0x00, 0x00,
	0xff, 0xff, 0xdc, 0x92, 0xf9, 0xf8, 0x2f, 0x06, 0x00, 0x00,
}
