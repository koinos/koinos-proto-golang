// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.17.3
// source: koinos/chain/value.proto

package chain

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ValueType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//	*ValueType_MessageValue
	//	*ValueType_Int32Value
	//	*ValueType_Int64Value
	//	*ValueType_Uint32Value
	//	*ValueType_Uint64Value
	//	*ValueType_Sint32Value
	//	*ValueType_Sint64Value
	//	*ValueType_Fixed32Value
	//	*ValueType_Fixed64Value
	//	*ValueType_Sfixed32Value
	//	*ValueType_Sfixed64Value
	//	*ValueType_BoolValue
	//	*ValueType_StringValue
	//	*ValueType_BytesValue
	Kind isValueType_Kind `protobuf_oneof:"kind"`
}

func (x *ValueType) Reset() {
	*x = ValueType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_chain_value_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueType) ProtoMessage() {}

func (x *ValueType) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_chain_value_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueType.ProtoReflect.Descriptor instead.
func (*ValueType) Descriptor() ([]byte, []int) {
	return file_koinos_chain_value_proto_rawDescGZIP(), []int{0}
}

func (m *ValueType) GetKind() isValueType_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *ValueType) GetMessageValue() *anypb.Any {
	if x, ok := x.GetKind().(*ValueType_MessageValue); ok {
		return x.MessageValue
	}
	return nil
}

func (x *ValueType) GetInt32Value() int32 {
	if x, ok := x.GetKind().(*ValueType_Int32Value); ok {
		return x.Int32Value
	}
	return 0
}

func (x *ValueType) GetInt64Value() int64 {
	if x, ok := x.GetKind().(*ValueType_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (x *ValueType) GetUint32Value() uint32 {
	if x, ok := x.GetKind().(*ValueType_Uint32Value); ok {
		return x.Uint32Value
	}
	return 0
}

func (x *ValueType) GetUint64Value() uint64 {
	if x, ok := x.GetKind().(*ValueType_Uint64Value); ok {
		return x.Uint64Value
	}
	return 0
}

func (x *ValueType) GetSint32Value() int32 {
	if x, ok := x.GetKind().(*ValueType_Sint32Value); ok {
		return x.Sint32Value
	}
	return 0
}

func (x *ValueType) GetSint64Value() int64 {
	if x, ok := x.GetKind().(*ValueType_Sint64Value); ok {
		return x.Sint64Value
	}
	return 0
}

func (x *ValueType) GetFixed32Value() uint32 {
	if x, ok := x.GetKind().(*ValueType_Fixed32Value); ok {
		return x.Fixed32Value
	}
	return 0
}

func (x *ValueType) GetFixed64Value() uint64 {
	if x, ok := x.GetKind().(*ValueType_Fixed64Value); ok {
		return x.Fixed64Value
	}
	return 0
}

func (x *ValueType) GetSfixed32Value() int32 {
	if x, ok := x.GetKind().(*ValueType_Sfixed32Value); ok {
		return x.Sfixed32Value
	}
	return 0
}

func (x *ValueType) GetSfixed64Value() int64 {
	if x, ok := x.GetKind().(*ValueType_Sfixed64Value); ok {
		return x.Sfixed64Value
	}
	return 0
}

func (x *ValueType) GetBoolValue() bool {
	if x, ok := x.GetKind().(*ValueType_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (x *ValueType) GetStringValue() string {
	if x, ok := x.GetKind().(*ValueType_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *ValueType) GetBytesValue() []byte {
	if x, ok := x.GetKind().(*ValueType_BytesValue); ok {
		return x.BytesValue
	}
	return nil
}

type isValueType_Kind interface {
	isValueType_Kind()
}

type ValueType_MessageValue struct {
	MessageValue *anypb.Any `protobuf:"bytes,1,opt,name=message_value,json=messageValue,proto3,oneof"`
}

type ValueType_Int32Value struct {
	Int32Value int32 `protobuf:"varint,2,opt,name=int32_value,json=int32Value,proto3,oneof"`
}

type ValueType_Int64Value struct {
	Int64Value int64 `protobuf:"varint,3,opt,name=int64_value,json=int64Value,proto3,oneof"`
}

type ValueType_Uint32Value struct {
	Uint32Value uint32 `protobuf:"varint,4,opt,name=uint32_value,json=uint32Value,proto3,oneof"`
}

type ValueType_Uint64Value struct {
	Uint64Value uint64 `protobuf:"varint,5,opt,name=uint64_value,json=uint64Value,proto3,oneof"`
}

type ValueType_Sint32Value struct {
	Sint32Value int32 `protobuf:"zigzag32,6,opt,name=sint32_value,json=sint32Value,proto3,oneof"`
}

type ValueType_Sint64Value struct {
	Sint64Value int64 `protobuf:"zigzag64,7,opt,name=sint64_value,json=sint64Value,proto3,oneof"`
}

type ValueType_Fixed32Value struct {
	Fixed32Value uint32 `protobuf:"fixed32,8,opt,name=fixed32_value,json=fixed32Value,proto3,oneof"`
}

type ValueType_Fixed64Value struct {
	Fixed64Value uint64 `protobuf:"fixed64,9,opt,name=fixed64_value,json=fixed64Value,proto3,oneof"`
}

type ValueType_Sfixed32Value struct {
	Sfixed32Value int32 `protobuf:"fixed32,10,opt,name=sfixed32_value,json=sfixed32Value,proto3,oneof"`
}

type ValueType_Sfixed64Value struct {
	Sfixed64Value int64 `protobuf:"fixed64,11,opt,name=sfixed64_value,json=sfixed64Value,proto3,oneof"`
}

type ValueType_BoolValue struct {
	BoolValue bool `protobuf:"varint,12,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type ValueType_StringValue struct {
	StringValue string `protobuf:"bytes,13,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type ValueType_BytesValue struct {
	BytesValue []byte `protobuf:"bytes,14,opt,name=bytes_value,json=bytesValue,proto3,oneof"`
}

func (*ValueType_MessageValue) isValueType_Kind() {}

func (*ValueType_Int32Value) isValueType_Kind() {}

func (*ValueType_Int64Value) isValueType_Kind() {}

func (*ValueType_Uint32Value) isValueType_Kind() {}

func (*ValueType_Uint64Value) isValueType_Kind() {}

func (*ValueType_Sint32Value) isValueType_Kind() {}

func (*ValueType_Sint64Value) isValueType_Kind() {}

func (*ValueType_Fixed32Value) isValueType_Kind() {}

func (*ValueType_Fixed64Value) isValueType_Kind() {}

func (*ValueType_Sfixed32Value) isValueType_Kind() {}

func (*ValueType_Sfixed64Value) isValueType_Kind() {}

func (*ValueType_BoolValue) isValueType_Kind() {}

func (*ValueType_StringValue) isValueType_Kind() {}

func (*ValueType_BytesValue) isValueType_Kind() {}

type EnumType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Number int32  `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *EnumType) Reset() {
	*x = EnumType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_chain_value_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnumType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnumType) ProtoMessage() {}

func (x *EnumType) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_chain_value_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnumType.ProtoReflect.Descriptor instead.
func (*EnumType) Descriptor() ([]byte, []int) {
	return file_koinos_chain_value_proto_rawDescGZIP(), []int{1}
}

func (x *EnumType) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EnumType) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type ListType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*ValueType `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *ListType) Reset() {
	*x = ListType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_chain_value_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListType) ProtoMessage() {}

func (x *ListType) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_chain_value_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListType.ProtoReflect.Descriptor instead.
func (*ListType) Descriptor() ([]byte, []int) {
	return file_koinos_chain_value_proto_rawDescGZIP(), []int{2}
}

func (x *ListType) GetValues() []*ValueType {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_koinos_chain_value_proto protoreflect.FileDescriptor

var file_koinos_chain_value_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6b, 0x6f, 0x69, 0x6e,
	0x6f, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x04, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x3b, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48,
	0x00, 0x52, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x21, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x01, 0x48, 0x00, 0x52, 0x0a, 0x69,
	0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x75, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x48,
	0x00, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x27,
	0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x11, 0x48, 0x00, 0x52,
	0x0b, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x27, 0x0a, 0x0c,
	0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x12, 0x42, 0x02, 0x30, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x25, 0x0a, 0x0d, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x07, 0x48, 0x00, 0x52, 0x0c,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x29, 0x0a, 0x0d,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x06, 0x42, 0x02, 0x30, 0x01, 0x48, 0x00, 0x52, 0x0c, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x27, 0x0a, 0x0e, 0x73, 0x66, 0x69, 0x78, 0x65,
	0x64, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0f, 0x48,
	0x00, 0x52, 0x0d, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x2b, 0x0a, 0x0e, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x10, 0x42, 0x02, 0x30, 0x01, 0x48, 0x00, 0x52, 0x0d,
	0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a,
	0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x00, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23,
	0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x37,
	0x0a, 0x09, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x3d, 0x0a, 0x09, 0x6c, 0x69, 0x73, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x52, 0x06,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x6b, 0x6f, 0x69, 0x6e,
	0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f,
	0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_koinos_chain_value_proto_rawDescOnce sync.Once
	file_koinos_chain_value_proto_rawDescData = file_koinos_chain_value_proto_rawDesc
)

func file_koinos_chain_value_proto_rawDescGZIP() []byte {
	file_koinos_chain_value_proto_rawDescOnce.Do(func() {
		file_koinos_chain_value_proto_rawDescData = protoimpl.X.CompressGZIP(file_koinos_chain_value_proto_rawDescData)
	})
	return file_koinos_chain_value_proto_rawDescData
}

var file_koinos_chain_value_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_koinos_chain_value_proto_goTypes = []interface{}{
	(*ValueType)(nil), // 0: koinos.chain.value_type
	(*EnumType)(nil),  // 1: koinos.chain.enum_type
	(*ListType)(nil),  // 2: koinos.chain.list_type
	(*anypb.Any)(nil), // 3: google.protobuf.Any
}
var file_koinos_chain_value_proto_depIdxs = []int32{
	3, // 0: koinos.chain.value_type.message_value:type_name -> google.protobuf.Any
	0, // 1: koinos.chain.list_type.values:type_name -> koinos.chain.value_type
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_koinos_chain_value_proto_init() }
func file_koinos_chain_value_proto_init() {
	if File_koinos_chain_value_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_koinos_chain_value_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueType); i {
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
		file_koinos_chain_value_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnumType); i {
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
		file_koinos_chain_value_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListType); i {
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
	file_koinos_chain_value_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ValueType_MessageValue)(nil),
		(*ValueType_Int32Value)(nil),
		(*ValueType_Int64Value)(nil),
		(*ValueType_Uint32Value)(nil),
		(*ValueType_Uint64Value)(nil),
		(*ValueType_Sint32Value)(nil),
		(*ValueType_Sint64Value)(nil),
		(*ValueType_Fixed32Value)(nil),
		(*ValueType_Fixed64Value)(nil),
		(*ValueType_Sfixed32Value)(nil),
		(*ValueType_Sfixed64Value)(nil),
		(*ValueType_BoolValue)(nil),
		(*ValueType_StringValue)(nil),
		(*ValueType_BytesValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_koinos_chain_value_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_koinos_chain_value_proto_goTypes,
		DependencyIndexes: file_koinos_chain_value_proto_depIdxs,
		MessageInfos:      file_koinos_chain_value_proto_msgTypes,
	}.Build()
	File_koinos_chain_value_proto = out.File
	file_koinos_chain_value_proto_rawDesc = nil
	file_koinos_chain_value_proto_goTypes = nil
	file_koinos_chain_value_proto_depIdxs = nil
}