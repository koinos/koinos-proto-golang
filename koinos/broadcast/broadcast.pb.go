// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: koinos/broadcast/broadcast.proto

package broadcast

import (
	koinos "github.com/koinos/koinos-proto-golang/koinos"
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

type TransactionAccepted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction       *protocol.Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Payer             []byte                `protobuf:"bytes,2,opt,name=payer,proto3" json:"payer,omitempty"`
	MaxPayerResources uint64                `protobuf:"varint,3,opt,name=max_payer_resources,json=maxPayerResources,proto3" json:"max_payer_resources,omitempty"`
	TrxResourceLimit  uint64                `protobuf:"varint,4,opt,name=trx_resource_limit,json=trxResourceLimit,proto3" json:"trx_resource_limit,omitempty"`
	Height            uint64                `protobuf:"varint,5,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *TransactionAccepted) Reset() {
	*x = TransactionAccepted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_broadcast_broadcast_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionAccepted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionAccepted) ProtoMessage() {}

func (x *TransactionAccepted) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_broadcast_broadcast_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionAccepted.ProtoReflect.Descriptor instead.
func (*TransactionAccepted) Descriptor() ([]byte, []int) {
	return file_koinos_broadcast_broadcast_proto_rawDescGZIP(), []int{0}
}

func (x *TransactionAccepted) GetTransaction() *protocol.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *TransactionAccepted) GetPayer() []byte {
	if x != nil {
		return x.Payer
	}
	return nil
}

func (x *TransactionAccepted) GetMaxPayerResources() uint64 {
	if x != nil {
		return x.MaxPayerResources
	}
	return 0
}

func (x *TransactionAccepted) GetTrxResourceLimit() uint64 {
	if x != nil {
		return x.TrxResourceLimit
	}
	return 0
}

func (x *TransactionAccepted) GetHeight() uint64 {
	if x != nil {
		return x.Height
	}
	return 0
}

type BlockAccepted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block *protocol.Block `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *BlockAccepted) Reset() {
	*x = BlockAccepted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_broadcast_broadcast_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockAccepted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockAccepted) ProtoMessage() {}

func (x *BlockAccepted) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_broadcast_broadcast_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockAccepted.ProtoReflect.Descriptor instead.
func (*BlockAccepted) Descriptor() ([]byte, []int) {
	return file_koinos_broadcast_broadcast_proto_rawDescGZIP(), []int{1}
}

func (x *BlockAccepted) GetBlock() *protocol.Block {
	if x != nil {
		return x.Block
	}
	return nil
}

type BlockIrreversible struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topology *koinos.BlockTopology `protobuf:"bytes,1,opt,name=topology,proto3" json:"topology,omitempty"`
}

func (x *BlockIrreversible) Reset() {
	*x = BlockIrreversible{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_broadcast_broadcast_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockIrreversible) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockIrreversible) ProtoMessage() {}

func (x *BlockIrreversible) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_broadcast_broadcast_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockIrreversible.ProtoReflect.Descriptor instead.
func (*BlockIrreversible) Descriptor() ([]byte, []int) {
	return file_koinos_broadcast_broadcast_proto_rawDescGZIP(), []int{2}
}

func (x *BlockIrreversible) GetTopology() *koinos.BlockTopology {
	if x != nil {
		return x.Topology
	}
	return nil
}

type ForkHeads struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastIrreversibleBlock *koinos.BlockTopology   `protobuf:"bytes,1,opt,name=last_irreversible_block,json=lastIrreversibleBlock,proto3" json:"last_irreversible_block,omitempty"`
	Heads                 []*koinos.BlockTopology `protobuf:"bytes,2,rep,name=heads,proto3" json:"heads,omitempty"`
}

func (x *ForkHeads) Reset() {
	*x = ForkHeads{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_broadcast_broadcast_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForkHeads) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForkHeads) ProtoMessage() {}

func (x *ForkHeads) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_broadcast_broadcast_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForkHeads.ProtoReflect.Descriptor instead.
func (*ForkHeads) Descriptor() ([]byte, []int) {
	return file_koinos_broadcast_broadcast_proto_rawDescGZIP(), []int{3}
}

func (x *ForkHeads) GetLastIrreversibleBlock() *koinos.BlockTopology {
	if x != nil {
		return x.LastIrreversibleBlock
	}
	return nil
}

func (x *ForkHeads) GetHeads() []*koinos.BlockTopology {
	if x != nil {
		return x.Heads
	}
	return nil
}

var File_koinos_broadcast_broadcast_proto protoreflect.FileDescriptor

var file_koinos_broadcast_broadcast_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x10, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x62, 0x72, 0x6f, 0x61, 0x64,
	0x63, 0x61, 0x73, 0x74, 0x1a, 0x13, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x14, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x65, 0x64, 0x12, 0x3e, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x70, 0x61, 0x79, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x13, 0x6d, 0x61, 0x78, 0x5f,
	0x70, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x50, 0x61,
	0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x12,
	0x74, 0x72, 0x78, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x10, 0x74, 0x72,
	0x78, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1a,
	0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02,
	0x30, 0x01, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x3e, 0x0a, 0x0e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x05,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x6f,
	0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x48, 0x0a, 0x12, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x72, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x62, 0x6c, 0x65,
	0x12, 0x32, 0x0a, 0x08, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52, 0x08, 0x74, 0x6f, 0x70, 0x6f,
	0x6c, 0x6f, 0x67, 0x79, 0x22, 0x8a, 0x01, 0x0a, 0x0a, 0x66, 0x6f, 0x72, 0x6b, 0x5f, 0x68, 0x65,
	0x61, 0x64, 0x73, 0x12, 0x4e, 0x0a, 0x17, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x72, 0x72, 0x65,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52, 0x15, 0x6c, 0x61,
	0x73, 0x74, 0x49, 0x72, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x2c, 0x0a, 0x05, 0x68, 0x65, 0x61, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x74, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52, 0x05, 0x68, 0x65, 0x61, 0x64,
	0x73, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f,
	0x73, 0x2f, 0x62, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_koinos_broadcast_broadcast_proto_rawDescOnce sync.Once
	file_koinos_broadcast_broadcast_proto_rawDescData = file_koinos_broadcast_broadcast_proto_rawDesc
)

func file_koinos_broadcast_broadcast_proto_rawDescGZIP() []byte {
	file_koinos_broadcast_broadcast_proto_rawDescOnce.Do(func() {
		file_koinos_broadcast_broadcast_proto_rawDescData = protoimpl.X.CompressGZIP(file_koinos_broadcast_broadcast_proto_rawDescData)
	})
	return file_koinos_broadcast_broadcast_proto_rawDescData
}

var file_koinos_broadcast_broadcast_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_koinos_broadcast_broadcast_proto_goTypes = []interface{}{
	(*TransactionAccepted)(nil),  // 0: koinos.broadcast.transaction_accepted
	(*BlockAccepted)(nil),        // 1: koinos.broadcast.block_accepted
	(*BlockIrreversible)(nil),    // 2: koinos.broadcast.block_irreversible
	(*ForkHeads)(nil),            // 3: koinos.broadcast.fork_heads
	(*protocol.Transaction)(nil), // 4: koinos.protocol.transaction
	(*protocol.Block)(nil),       // 5: koinos.protocol.block
	(*koinos.BlockTopology)(nil), // 6: koinos.block_topology
}
var file_koinos_broadcast_broadcast_proto_depIdxs = []int32{
	4, // 0: koinos.broadcast.transaction_accepted.transaction:type_name -> koinos.protocol.transaction
	5, // 1: koinos.broadcast.block_accepted.block:type_name -> koinos.protocol.block
	6, // 2: koinos.broadcast.block_irreversible.topology:type_name -> koinos.block_topology
	6, // 3: koinos.broadcast.fork_heads.last_irreversible_block:type_name -> koinos.block_topology
	6, // 4: koinos.broadcast.fork_heads.heads:type_name -> koinos.block_topology
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_koinos_broadcast_broadcast_proto_init() }
func file_koinos_broadcast_broadcast_proto_init() {
	if File_koinos_broadcast_broadcast_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_koinos_broadcast_broadcast_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionAccepted); i {
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
		file_koinos_broadcast_broadcast_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockAccepted); i {
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
		file_koinos_broadcast_broadcast_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockIrreversible); i {
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
		file_koinos_broadcast_broadcast_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForkHeads); i {
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
			RawDescriptor: file_koinos_broadcast_broadcast_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_koinos_broadcast_broadcast_proto_goTypes,
		DependencyIndexes: file_koinos_broadcast_broadcast_proto_depIdxs,
		MessageInfos:      file_koinos_broadcast_broadcast_proto_msgTypes,
	}.Build()
	File_koinos_broadcast_broadcast_proto = out.File
	file_koinos_broadcast_broadcast_proto_rawDesc = nil
	file_koinos_broadcast_broadcast_proto_goTypes = nil
	file_koinos_broadcast_broadcast_proto_depIdxs = nil
}
