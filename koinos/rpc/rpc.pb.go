// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: koinos/rpc/rpc.proto

package rpc

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

type ReservedRpc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReservedRpc) Reset() {
	*x = ReservedRpc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_rpc_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservedRpc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservedRpc) ProtoMessage() {}

func (x *ReservedRpc) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_rpc_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservedRpc.ProtoReflect.Descriptor instead.
func (*ReservedRpc) Descriptor() ([]byte, []int) {
	return file_koinos_rpc_rpc_proto_rawDescGZIP(), []int{0}
}

type ErrorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Data    string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ErrorResponse) Reset() {
	*x = ErrorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_rpc_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorResponse) ProtoMessage() {}

func (x *ErrorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_rpc_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorResponse.ProtoReflect.Descriptor instead.
func (*ErrorResponse) Descriptor() ([]byte, []int) {
	return file_koinos_rpc_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *ErrorResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ErrorResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_koinos_rpc_rpc_proto protoreflect.FileDescriptor

var file_koinos_rpc_rpc_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72,
	0x70, 0x63, 0x22, 0x0e, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x72,
	0x70, 0x63, 0x22, 0x3e, 0x0a, 0x0e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6b, 0x6f, 0x69, 0x6e,
	0x6f, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_koinos_rpc_rpc_proto_rawDescOnce sync.Once
	file_koinos_rpc_rpc_proto_rawDescData = file_koinos_rpc_rpc_proto_rawDesc
)

func file_koinos_rpc_rpc_proto_rawDescGZIP() []byte {
	file_koinos_rpc_rpc_proto_rawDescOnce.Do(func() {
		file_koinos_rpc_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_koinos_rpc_rpc_proto_rawDescData)
	})
	return file_koinos_rpc_rpc_proto_rawDescData
}

var file_koinos_rpc_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_koinos_rpc_rpc_proto_goTypes = []interface{}{
	(*ReservedRpc)(nil),   // 0: koinos.rpc.reserved_rpc
	(*ErrorResponse)(nil), // 1: koinos.rpc.error_response
}
var file_koinos_rpc_rpc_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_koinos_rpc_rpc_proto_init() }
func file_koinos_rpc_rpc_proto_init() {
	if File_koinos_rpc_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_koinos_rpc_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservedRpc); i {
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
		file_koinos_rpc_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorResponse); i {
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
			RawDescriptor: file_koinos_rpc_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_koinos_rpc_rpc_proto_goTypes,
		DependencyIndexes: file_koinos_rpc_rpc_proto_depIdxs,
		MessageInfos:      file_koinos_rpc_rpc_proto_msgTypes,
	}.Build()
	File_koinos_rpc_rpc_proto = out.File
	file_koinos_rpc_rpc_proto_rawDesc = nil
	file_koinos_rpc_rpc_proto_goTypes = nil
	file_koinos_rpc_rpc_proto_depIdxs = nil
}
