// Copyright 2018 The Swarming Authors. All rights reserved.
// Use of this source code is governed by the Apache v2.0 license that can be
// found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.1
// source: go.chromium.org/luci/buildbucket/proto/notification.proto

package buildbucketpb

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// A notification about a build.
type Notification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// When this notification was created.
	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Cloud Project ID of the Buildbucket instance that sent this notification,
	// e.g. "cr-buildbucket".
	// Useful if a service listens to both prod and dev instances of buildbucket.
	AppId string `protobuf:"bytes,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	// Buildbucket build id.
	// Use GetBuild rpc to load the contents.
	BuildId int64 `protobuf:"varint,3,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	// User-defined opaque blob specified in NotificationConfig.user_data.
	UserData []byte `protobuf:"bytes,4,opt,name=user_data,json=userData,proto3" json:"user_data,omitempty"`
}

func (x *Notification) Reset() {
	*x = Notification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_proto_notification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_proto_notification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDescGZIP(), []int{0}
}

func (x *Notification) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Notification) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *Notification) GetBuildId() int64 {
	if x != nil {
		return x.BuildId
	}
	return 0
}

func (x *Notification) GetUserData() []byte {
	if x != nil {
		return x.UserData
	}
	return nil
}

// Configuration of notifications.
type NotificationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Target Cloud PubSub topic.
	// Usually has format "projects/{cloud project}/topics/{topic name}".
	//
	// The PubSub message data is a Notification message in binary format.
	//
	// <buildbucket-app-id>@appspot.gserviceaccount.com must have
	// "pubsub.topics.publish" permissions on the topic, where
	// <buildbucket-app-id> is usually "cr-buildbucket."
	PubsubTopic string `protobuf:"bytes,1,opt,name=pubsub_topic,json=pubsubTopic,proto3" json:"pubsub_topic,omitempty"`
	// Will be available in Notification.user_data.
	// Max length: 4096.
	UserData []byte `protobuf:"bytes,2,opt,name=user_data,json=userData,proto3" json:"user_data,omitempty"`
}

func (x *NotificationConfig) Reset() {
	*x = NotificationConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_go_chromium_org_luci_buildbucket_proto_notification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotificationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationConfig) ProtoMessage() {}

func (x *NotificationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_go_chromium_org_luci_buildbucket_proto_notification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationConfig.ProtoReflect.Descriptor instead.
func (*NotificationConfig) Descriptor() ([]byte, []int) {
	return file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDescGZIP(), []int{1}
}

func (x *NotificationConfig) GetPubsubTopic() string {
	if x != nil {
		return x.PubsubTopic
	}
	return ""
}

func (x *NotificationConfig) GetUserData() []byte {
	if x != nil {
		return x.UserData
	}
	return nil
}

var File_go_chromium_org_luci_buildbucket_proto_notification_proto protoreflect.FileDescriptor

var file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDesc = []byte{
	0x0a, 0x39, 0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x6c, 0x75, 0x63, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x01, 0x0a,
	0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x22, 0x54, 0x0a, 0x12, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12,
	0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x42, 0x36, 0x5a, 0x34,
	0x67, 0x6f, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x6f, 0x72, 0x67, 0x2f,
	0x6c, 0x75, 0x63, 0x69, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDescOnce sync.Once
	file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDescData = file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDesc
)

func file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDescGZIP() []byte {
	file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDescOnce.Do(func() {
		file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDescData = protoimpl.X.CompressGZIP(file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDescData)
	})
	return file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDescData
}

var file_go_chromium_org_luci_buildbucket_proto_notification_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_go_chromium_org_luci_buildbucket_proto_notification_proto_goTypes = []interface{}{
	(*Notification)(nil),        // 0: buildbucket.v2.Notification
	(*NotificationConfig)(nil),  // 1: buildbucket.v2.NotificationConfig
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_go_chromium_org_luci_buildbucket_proto_notification_proto_depIdxs = []int32{
	2, // 0: buildbucket.v2.Notification.timestamp:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_go_chromium_org_luci_buildbucket_proto_notification_proto_init() }
func file_go_chromium_org_luci_buildbucket_proto_notification_proto_init() {
	if File_go_chromium_org_luci_buildbucket_proto_notification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_go_chromium_org_luci_buildbucket_proto_notification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Notification); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_go_chromium_org_luci_buildbucket_proto_notification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotificationConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_go_chromium_org_luci_buildbucket_proto_notification_proto_goTypes,
		DependencyIndexes: file_go_chromium_org_luci_buildbucket_proto_notification_proto_depIdxs,
		MessageInfos:      file_go_chromium_org_luci_buildbucket_proto_notification_proto_msgTypes,
	}.Build()
	File_go_chromium_org_luci_buildbucket_proto_notification_proto = out.File
	file_go_chromium_org_luci_buildbucket_proto_notification_proto_rawDesc = nil
	file_go_chromium_org_luci_buildbucket_proto_notification_proto_goTypes = nil
	file_go_chromium_org_luci_buildbucket_proto_notification_proto_depIdxs = nil
}
