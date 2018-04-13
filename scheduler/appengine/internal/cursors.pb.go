// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/scheduler/appengine/internal/cursors.proto

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	go.chromium.org/luci/scheduler/appengine/internal/cursors.proto
	go.chromium.org/luci/scheduler/appengine/internal/db.proto
	go.chromium.org/luci/scheduler/appengine/internal/timers.proto
	go.chromium.org/luci/scheduler/appengine/internal/tq.proto
	go.chromium.org/luci/scheduler/appengine/internal/triggers.proto

It has these top-level messages:
	InvocationsCursor
	FinishedInvocation
	FinishedInvocationList
	Timer
	TimerList
	ReadProjectConfigTask
	LaunchInvocationTask
	LaunchInvocationsBatchTask
	TriageJobStateTask
	KickTriageTask
	InvocationFinishedTask
	FanOutTriggersTask
	EnqueueTriggersTask
	ScheduleTimersTask
	TimerTask
	CronTickTask
	Trigger
	TriggerList
*/
package internal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// InvocationsCursor is used to paginate results of GetInvocations RPC call.
//
// It is serialized in base64 and sent to the clients. There's no integrity
// protection: we assume broken cursors are rejected down the call stack
// (which is the case with the datastore cursors).
//
// The internal structure of the cursor is implementation detail and clients
// must not depend on it.
type InvocationsCursor struct {
	DsCursor []byte `protobuf:"bytes,1,opt,name=ds_cursor,json=dsCursor,proto3" json:"ds_cursor,omitempty"`
}

func (m *InvocationsCursor) Reset()                    { *m = InvocationsCursor{} }
func (m *InvocationsCursor) String() string            { return proto.CompactTextString(m) }
func (*InvocationsCursor) ProtoMessage()               {}
func (*InvocationsCursor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InvocationsCursor) GetDsCursor() []byte {
	if m != nil {
		return m.DsCursor
	}
	return nil
}

func init() {
	proto.RegisterType((*InvocationsCursor)(nil), "internal.cursors.InvocationsCursor")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/scheduler/appengine/internal/cursors.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4f, 0xcf, 0xd7, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c, 0xcd, 0xd5, 0xcb, 0x2f, 0x4a, 0xd7, 0xcf, 0x29, 0x4d, 0xce,
	0xd4, 0x2f, 0x4e, 0xce, 0x48, 0x4d, 0x29, 0xcd, 0x49, 0x2d, 0xd2, 0x4f, 0x2c, 0x28, 0x48, 0xcd,
	0x4b, 0xcf, 0xcc, 0x4b, 0xd5, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x4f, 0x2e,
	0x2d, 0x2a, 0xce, 0x2f, 0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x80, 0x89, 0xeb,
	0x41, 0xc5, 0x95, 0x0c, 0xb8, 0x04, 0x3d, 0xf3, 0xca, 0xf2, 0x93, 0x13, 0x4b, 0x32, 0xf3, 0xf3,
	0x8a, 0x9d, 0xc1, 0xa2, 0x42, 0xd2, 0x5c, 0x9c, 0x29, 0xc5, 0xf1, 0x10, 0x25, 0x12, 0x8c, 0x0a,
	0x8c, 0x1a, 0x3c, 0x41, 0x1c, 0x29, 0x50, 0x49, 0x27, 0xae, 0x28, 0x0e, 0x98, 0x29, 0x49, 0x6c,
	0x60, 0x63, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x1a, 0xf3, 0xf8, 0x99, 0x00, 0x00,
	0x00,
}
