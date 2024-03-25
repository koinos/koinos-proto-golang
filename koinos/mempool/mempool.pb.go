// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.17.3
// source: koinos/mempool/mempool.proto

package mempool

import (
	protocol "github.com/koinos/koinos-proto-golang/v2/koinos/protocol"
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

type MempoolMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeqNum uint64 `protobuf:"varint,1,opt,name=seq_num,json=seqNum,proto3" json:"seq_num,omitempty"`
}

func (x *MempoolMetadata) Reset() {
	*x = MempoolMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_mempool_mempool_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MempoolMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MempoolMetadata) ProtoMessage() {}

func (x *MempoolMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_mempool_mempool_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MempoolMetadata.ProtoReflect.Descriptor instead.
func (*MempoolMetadata) Descriptor() ([]byte, []int) {
	return file_koinos_mempool_mempool_proto_rawDescGZIP(), []int{0}
}

func (x *MempoolMetadata) GetSeqNum() uint64 {
	if x != nil {
		return x.SeqNum
	}
	return 0
}

type AddressResourceRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxRc     uint64 `protobuf:"varint,1,opt,name=max_rc,json=maxRc,proto3" json:"max_rc,omitempty"`
	CurrentRc uint64 `protobuf:"varint,2,opt,name=current_rc,json=currentRc,proto3" json:"current_rc,omitempty"`
}

func (x *AddressResourceRecord) Reset() {
	*x = AddressResourceRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_mempool_mempool_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressResourceRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressResourceRecord) ProtoMessage() {}

func (x *AddressResourceRecord) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_mempool_mempool_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressResourceRecord.ProtoReflect.Descriptor instead.
func (*AddressResourceRecord) Descriptor() ([]byte, []int) {
	return file_koinos_mempool_mempool_proto_rawDescGZIP(), []int{1}
}

func (x *AddressResourceRecord) GetMaxRc() uint64 {
	if x != nil {
		return x.MaxRc
	}
	return 0
}

func (x *AddressResourceRecord) GetCurrentRc() uint64 {
	if x != nil {
		return x.CurrentRc
	}
	return 0
}

type PendingTransactionRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction          *protocol.Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Timestamp            uint64                `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	DiskStorageUsed      uint64                `protobuf:"varint,3,opt,name=disk_storage_used,json=diskStorageUsed,proto3" json:"disk_storage_used,omitempty"`
	NetworkBandwidthUsed uint64                `protobuf:"varint,4,opt,name=network_bandwidth_used,json=networkBandwidthUsed,proto3" json:"network_bandwidth_used,omitempty"`
	ComputeBandwidthUsed uint64                `protobuf:"varint,5,opt,name=compute_bandwidth_used,json=computeBandwidthUsed,proto3" json:"compute_bandwidth_used,omitempty"`
}

func (x *PendingTransactionRecord) Reset() {
	*x = PendingTransactionRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_mempool_mempool_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PendingTransactionRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PendingTransactionRecord) ProtoMessage() {}

func (x *PendingTransactionRecord) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_mempool_mempool_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PendingTransactionRecord.ProtoReflect.Descriptor instead.
func (*PendingTransactionRecord) Descriptor() ([]byte, []int) {
	return file_koinos_mempool_mempool_proto_rawDescGZIP(), []int{2}
}

func (x *PendingTransactionRecord) GetTransaction() *protocol.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *PendingTransactionRecord) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *PendingTransactionRecord) GetDiskStorageUsed() uint64 {
	if x != nil {
		return x.DiskStorageUsed
	}
	return 0
}

func (x *PendingTransactionRecord) GetNetworkBandwidthUsed() uint64 {
	if x != nil {
		return x.NetworkBandwidthUsed
	}
	return 0
}

func (x *PendingTransactionRecord) GetComputeBandwidthUsed() uint64 {
	if x != nil {
		return x.ComputeBandwidthUsed
	}
	return 0
}

var File_koinos_mempool_mempool_proto protoreflect.FileDescriptor

var file_koinos_mempool_mempool_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x6d, 0x70, 0x6f, 0x6f, 0x6c,
	0x2f, 0x6d, 0x65, 0x6d, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x6d, 0x65, 0x6d, 0x70, 0x6f, 0x6f, 0x6c, 0x1a, 0x1e,
	0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b,
	0x0a, 0x10, 0x6d, 0x65, 0x6d, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x71, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x65, 0x71, 0x4e, 0x75, 0x6d, 0x22, 0x4f, 0x0a, 0x17, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6d, 0x61, 0x78, 0x52, 0x63, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x63, 0x22, 0x92, 0x02, 0x0a,
	0x1a, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x3e, 0x0a, 0x0b, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x69, 0x73,
	0x6b, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x55, 0x73, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x16, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x5f, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x61,
	0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x55, 0x73, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x16, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68,
	0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x55, 0x73, 0x65,
	0x64, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x6b, 0x6f,
	0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x6d, 0x70, 0x6f, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_koinos_mempool_mempool_proto_rawDescOnce sync.Once
	file_koinos_mempool_mempool_proto_rawDescData = file_koinos_mempool_mempool_proto_rawDesc
)

func file_koinos_mempool_mempool_proto_rawDescGZIP() []byte {
	file_koinos_mempool_mempool_proto_rawDescOnce.Do(func() {
		file_koinos_mempool_mempool_proto_rawDescData = protoimpl.X.CompressGZIP(file_koinos_mempool_mempool_proto_rawDescData)
	})
	return file_koinos_mempool_mempool_proto_rawDescData
}

var file_koinos_mempool_mempool_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_koinos_mempool_mempool_proto_goTypes = []interface{}{
	(*MempoolMetadata)(nil),          // 0: koinos.mempool.mempool_metadata
	(*AddressResourceRecord)(nil),    // 1: koinos.mempool.address_resource_record
	(*PendingTransactionRecord)(nil), // 2: koinos.mempool.pending_transaction_record
	(*protocol.Transaction)(nil),     // 3: koinos.protocol.transaction
}
var file_koinos_mempool_mempool_proto_depIdxs = []int32{
	3, // 0: koinos.mempool.pending_transaction_record.transaction:type_name -> koinos.protocol.transaction
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_koinos_mempool_mempool_proto_init() }
func file_koinos_mempool_mempool_proto_init() {
	if File_koinos_mempool_mempool_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_koinos_mempool_mempool_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MempoolMetadata); i {
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
		file_koinos_mempool_mempool_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressResourceRecord); i {
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
		file_koinos_mempool_mempool_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PendingTransactionRecord); i {
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
			RawDescriptor: file_koinos_mempool_mempool_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_koinos_mempool_mempool_proto_goTypes,
		DependencyIndexes: file_koinos_mempool_mempool_proto_depIdxs,
		MessageInfos:      file_koinos_mempool_mempool_proto_msgTypes,
	}.Build()
	File_koinos_mempool_mempool_proto = out.File
	file_koinos_mempool_mempool_proto_rawDesc = nil
	file_koinos_mempool_mempool_proto_goTypes = nil
	file_koinos_mempool_mempool_proto_depIdxs = nil
}
