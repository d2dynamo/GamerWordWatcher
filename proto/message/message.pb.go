// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.12.4
// source: proto/message/message.proto

package message

import (
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

type TextMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content    string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	GuildDID   string `protobuf:"bytes,2,opt,name=guildDID,proto3" json:"guildDID,omitempty"`
	ChannelDID string `protobuf:"bytes,3,opt,name=channelDID,proto3" json:"channelDID,omitempty"`
	MessageDID string `protobuf:"bytes,4,opt,name=messageDID,proto3" json:"messageDID,omitempty"`
}

func (x *TextMessage) Reset() {
	*x = TextMessage{}
	mi := &file_proto_message_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TextMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextMessage) ProtoMessage() {}

func (x *TextMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextMessage.ProtoReflect.Descriptor instead.
func (*TextMessage) Descriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{0}
}

func (x *TextMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *TextMessage) GetGuildDID() string {
	if x != nil {
		return x.GuildDID
	}
	return ""
}

func (x *TextMessage) GetChannelDID() string {
	if x != nil {
		return x.ChannelDID
	}
	return ""
}

func (x *TextMessage) GetMessageDID() string {
	if x != nil {
		return x.MessageDID
	}
	return ""
}

type WatcherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Acknowledged bool `protobuf:"varint,1,opt,name=acknowledged,proto3" json:"acknowledged,omitempty"`
	IsGamerWord  bool `protobuf:"varint,2,opt,name=isGamerWord,proto3" json:"isGamerWord,omitempty"`
}

func (x *WatcherResponse) Reset() {
	*x = WatcherResponse{}
	mi := &file_proto_message_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WatcherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatcherResponse) ProtoMessage() {}

func (x *WatcherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatcherResponse.ProtoReflect.Descriptor instead.
func (*WatcherResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{1}
}

func (x *WatcherResponse) GetAcknowledged() bool {
	if x != nil {
		return x.Acknowledged
	}
	return false
}

func (x *WatcherResponse) GetIsGamerWord() bool {
	if x != nil {
		return x.IsGamerWord
	}
	return false
}

var File_proto_message_message_proto protoreflect.FileDescriptor

var file_proto_message_message_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x0b, 0x54, 0x65, 0x78, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x67, 0x75, 0x69, 0x6c, 0x64, 0x44, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x44, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x44, 0x49, 0x44, 0x22, 0x57, 0x0a, 0x0f,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x61, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x73, 0x47, 0x61, 0x6d, 0x65, 0x72, 0x57, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x47, 0x61, 0x6d, 0x65,
	0x72, 0x57, 0x6f, 0x72, 0x64, 0x32, 0x56, 0x0a, 0x10, 0x47, 0x61, 0x6d, 0x65, 0x72, 0x57, 0x6f,
	0x72, 0x64, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x0e, 0x57, 0x61, 0x74,
	0x63, 0x68, 0x47, 0x61, 0x6d, 0x65, 0x72, 0x57, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x18, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x34, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x32, 0x64, 0x79,
	0x6e, 0x61, 0x6d, 0x6f, 0x2f, 0x47, 0x61, 0x6d, 0x65, 0x72, 0x57, 0x6f, 0x72, 0x64, 0x57, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_message_message_proto_rawDescOnce sync.Once
	file_proto_message_message_proto_rawDescData = file_proto_message_message_proto_rawDesc
)

func file_proto_message_message_proto_rawDescGZIP() []byte {
	file_proto_message_message_proto_rawDescOnce.Do(func() {
		file_proto_message_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_message_message_proto_rawDescData)
	})
	return file_proto_message_message_proto_rawDescData
}

var file_proto_message_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_message_message_proto_goTypes = []any{
	(*TextMessage)(nil),     // 0: message.TextMessage
	(*WatcherResponse)(nil), // 1: message.WatcherResponse
}
var file_proto_message_message_proto_depIdxs = []int32{
	0, // 0: message.GamerWordWatcher.WatchGamerWord:input_type -> message.TextMessage
	1, // 1: message.GamerWordWatcher.WatchGamerWord:output_type -> message.WatcherResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_message_message_proto_init() }
func file_proto_message_message_proto_init() {
	if File_proto_message_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_message_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_message_message_proto_goTypes,
		DependencyIndexes: file_proto_message_message_proto_depIdxs,
		MessageInfos:      file_proto_message_message_proto_msgTypes,
	}.Build()
	File_proto_message_message_proto = out.File
	file_proto_message_message_proto_rawDesc = nil
	file_proto_message_message_proto_goTypes = nil
	file_proto_message_message_proto_depIdxs = nil
}
