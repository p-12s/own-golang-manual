// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.2
// source: chatpb/chat.proto

package chatpb

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

type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Id   int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chatpb_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chatpb_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_chatpb_chat_proto_rawDescGZIP(), []int{0}
}

func (x *ChatMessage) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ChatMessage) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_chatpb_chat_proto protoreflect.FileDescriptor

var file_chatpb_chat_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x37, 0x0a, 0x0b, 0x43, 0x68, 0x61,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x4a, 0x04, 0x08, 0x03,
	0x10, 0x04, 0x32, 0x42, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x12, 0x33, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x11, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x1a, 0x11, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x63, 0x68, 0x61, 0x74,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chatpb_chat_proto_rawDescOnce sync.Once
	file_chatpb_chat_proto_rawDescData = file_chatpb_chat_proto_rawDesc
)

func file_chatpb_chat_proto_rawDescGZIP() []byte {
	file_chatpb_chat_proto_rawDescOnce.Do(func() {
		file_chatpb_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chatpb_chat_proto_rawDescData)
	})
	return file_chatpb_chat_proto_rawDescData
}

var file_chatpb_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_chatpb_chat_proto_goTypes = []interface{}{
	(*ChatMessage)(nil), // 0: chat.ChatMessage
}
var file_chatpb_chat_proto_depIdxs = []int32{
	0, // 0: chat.ChatExample.sendMessage:input_type -> chat.ChatMessage
	0, // 1: chat.ChatExample.sendMessage:output_type -> chat.ChatMessage
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chatpb_chat_proto_init() }
func file_chatpb_chat_proto_init() {
	if File_chatpb_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chatpb_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessage); i {
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
			RawDescriptor: file_chatpb_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chatpb_chat_proto_goTypes,
		DependencyIndexes: file_chatpb_chat_proto_depIdxs,
		MessageInfos:      file_chatpb_chat_proto_msgTypes,
	}.Build()
	File_chatpb_chat_proto = out.File
	file_chatpb_chat_proto_rawDesc = nil
	file_chatpb_chat_proto_goTypes = nil
	file_chatpb_chat_proto_depIdxs = nil
}