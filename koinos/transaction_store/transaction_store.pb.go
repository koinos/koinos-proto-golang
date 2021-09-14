// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: koinos/transaction_store/transaction_store.proto

package transaction_store

import (
	protocol "github.com/koinos/koinos-proto-golang/koinos/protocol"
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

type TransactionItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction      *protocol.Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	ContainingBlocks [][]byte              `protobuf:"bytes,2,rep,name=containing_blocks,json=containingBlocks,proto3" json:"containing_blocks,omitempty"`
}

func (x *TransactionItem) Reset() {
	*x = TransactionItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_transaction_store_transaction_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionItem) ProtoMessage() {}

func (x *TransactionItem) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_transaction_store_transaction_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionItem.ProtoReflect.Descriptor instead.
func (*TransactionItem) Descriptor() ([]byte, []int) {
	return file_koinos_transaction_store_transaction_store_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionItem) GetTransaction() *protocol.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *TransactionItem) GetContainingBlocks() [][]byte {
	if x != nil {
		return x.ContainingBlocks
	}
	return nil
}

var File_koinos_transaction_store_transaction_store_proto protoreflect.FileDescriptor

var file_koinos_transaction_store_transaction_store_proto_rawDesc = []byte{
	0x0a, 0x30, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x18, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x1e, 0x6b, 0x6f,
	0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a, 0x10,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x12, 0x3e, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x10, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x42, 0x40, 0x5a,
	0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x69, 0x6e,
	0x6f, 0x73, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_koinos_transaction_store_transaction_store_proto_rawDescOnce sync.Once
	file_koinos_transaction_store_transaction_store_proto_rawDescData = file_koinos_transaction_store_transaction_store_proto_rawDesc
)

func file_koinos_transaction_store_transaction_store_proto_rawDescGZIP() []byte {
	file_koinos_transaction_store_transaction_store_proto_rawDescOnce.Do(func() {
		file_koinos_transaction_store_transaction_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_koinos_transaction_store_transaction_store_proto_rawDescData)
	})
	return file_koinos_transaction_store_transaction_store_proto_rawDescData
}

var file_koinos_transaction_store_transaction_store_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_koinos_transaction_store_transaction_store_proto_goTypes = []interface{}{
	(*TransactionItem)(nil),      // 0: koinos.transaction_store.transaction_item
	(*protocol.Transaction)(nil), // 1: koinos.protocol.transaction
}
var file_koinos_transaction_store_transaction_store_proto_depIdxs = []int32{
	1, // 0: koinos.transaction_store.transaction_item.transaction:type_name -> koinos.protocol.transaction
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_koinos_transaction_store_transaction_store_proto_init() }
func file_koinos_transaction_store_transaction_store_proto_init() {
	if File_koinos_transaction_store_transaction_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_koinos_transaction_store_transaction_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionItem); i {
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
			RawDescriptor: file_koinos_transaction_store_transaction_store_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_koinos_transaction_store_transaction_store_proto_goTypes,
		DependencyIndexes: file_koinos_transaction_store_transaction_store_proto_depIdxs,
		MessageInfos:      file_koinos_transaction_store_transaction_store_proto_msgTypes,
	}.Build()
	File_koinos_transaction_store_transaction_store_proto = out.File
	file_koinos_transaction_store_transaction_store_proto_rawDesc = nil
	file_koinos_transaction_store_transaction_store_proto_goTypes = nil
	file_koinos_transaction_store_transaction_store_proto_depIdxs = nil
}
