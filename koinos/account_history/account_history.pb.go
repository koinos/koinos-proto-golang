// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.17.3
// source: koinos/account_history/account_history.proto

package account_history

import (
	_ "github.com/koinos/koinos-proto-golang/v2/koinos"
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

type AccountMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeqNum uint64 `protobuf:"varint,1,opt,name=seq_num,json=seqNum,proto3" json:"seq_num,omitempty"`
}

func (x *AccountMetadata) Reset() {
	*x = AccountMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_account_history_account_history_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountMetadata) ProtoMessage() {}

func (x *AccountMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_account_history_account_history_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountMetadata.ProtoReflect.Descriptor instead.
func (*AccountMetadata) Descriptor() ([]byte, []int) {
	return file_koinos_account_history_account_history_proto_rawDescGZIP(), []int{0}
}

func (x *AccountMetadata) GetSeqNum() uint64 {
	if x != nil {
		return x.SeqNum
	}
	return 0
}

type HistoryIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	SeqNum  uint64 `protobuf:"varint,2,opt,name=seq_num,json=seqNum,proto3" json:"seq_num,omitempty"`
}

func (x *HistoryIndex) Reset() {
	*x = HistoryIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_account_history_account_history_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryIndex) ProtoMessage() {}

func (x *HistoryIndex) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_account_history_account_history_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryIndex.ProtoReflect.Descriptor instead.
func (*HistoryIndex) Descriptor() ([]byte, []int) {
	return file_koinos_account_history_account_history_proto_rawDescGZIP(), []int{1}
}

func (x *HistoryIndex) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *HistoryIndex) GetSeqNum() uint64 {
	if x != nil {
		return x.SeqNum
	}
	return 0
}

type TransactionRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *protocol.Transaction        `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Receipt     *protocol.TransactionReceipt `protobuf:"bytes,2,opt,name=receipt,proto3" json:"receipt,omitempty"`
}

func (x *TransactionRecord) Reset() {
	*x = TransactionRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_account_history_account_history_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionRecord) ProtoMessage() {}

func (x *TransactionRecord) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_account_history_account_history_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionRecord.ProtoReflect.Descriptor instead.
func (*TransactionRecord) Descriptor() ([]byte, []int) {
	return file_koinos_account_history_account_history_proto_rawDescGZIP(), []int{2}
}

func (x *TransactionRecord) GetTransaction() *protocol.Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *TransactionRecord) GetReceipt() *protocol.TransactionReceipt {
	if x != nil {
		return x.Receipt
	}
	return nil
}

type BlockRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header  *protocol.BlockHeader  `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Receipt *protocol.BlockReceipt `protobuf:"bytes,2,opt,name=receipt,proto3" json:"receipt,omitempty"`
}

func (x *BlockRecord) Reset() {
	*x = BlockRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_account_history_account_history_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockRecord) ProtoMessage() {}

func (x *BlockRecord) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_account_history_account_history_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockRecord.ProtoReflect.Descriptor instead.
func (*BlockRecord) Descriptor() ([]byte, []int) {
	return file_koinos_account_history_account_history_proto_rawDescGZIP(), []int{3}
}

func (x *BlockRecord) GetHeader() *protocol.BlockHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *BlockRecord) GetReceipt() *protocol.BlockReceipt {
	if x != nil {
		return x.Receipt
	}
	return nil
}

type HistoryRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Record:
	//
	//	*HistoryRecord_Trx
	//	*HistoryRecord_Block
	Record isHistoryRecord_Record `protobuf_oneof:"record"`
}

func (x *HistoryRecord) Reset() {
	*x = HistoryRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_account_history_account_history_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryRecord) ProtoMessage() {}

func (x *HistoryRecord) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_account_history_account_history_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryRecord.ProtoReflect.Descriptor instead.
func (*HistoryRecord) Descriptor() ([]byte, []int) {
	return file_koinos_account_history_account_history_proto_rawDescGZIP(), []int{4}
}

func (m *HistoryRecord) GetRecord() isHistoryRecord_Record {
	if m != nil {
		return m.Record
	}
	return nil
}

func (x *HistoryRecord) GetTrx() *TransactionRecord {
	if x, ok := x.GetRecord().(*HistoryRecord_Trx); ok {
		return x.Trx
	}
	return nil
}

func (x *HistoryRecord) GetBlock() *BlockRecord {
	if x, ok := x.GetRecord().(*HistoryRecord_Block); ok {
		return x.Block
	}
	return nil
}

type isHistoryRecord_Record interface {
	isHistoryRecord_Record()
}

type HistoryRecord_Trx struct {
	Trx *TransactionRecord `protobuf:"bytes,1,opt,name=trx,proto3,oneof"`
}

type HistoryRecord_Block struct {
	Block *BlockRecord `protobuf:"bytes,2,opt,name=block,proto3,oneof"`
}

func (*HistoryRecord_Trx) isHistoryRecord_Record() {}

func (*HistoryRecord_Block) isHistoryRecord_Record() {}

var File_koinos_account_history_account_history_proto protoreflect.FileDescriptor

var file_koinos_account_history_account_history_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0x14, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6b, 0x6f,
	0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x10,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x71, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x73, 0x65, 0x71, 0x4e, 0x75, 0x6d, 0x22, 0x48, 0x0a, 0x0d, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1e, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x04, 0x80, 0xb5, 0x18,
	0x06, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x65,
	0x71, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x65, 0x71,
	0x4e, 0x75, 0x6d, 0x22, 0x94, 0x01, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x3e, 0x0a, 0x0b, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x07, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6b, 0x6f,
	0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x52, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x22, 0x7f, 0x0a, 0x0c, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x35, 0x0a, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6b, 0x6f, 0x69,
	0x6e, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x38, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x70, 0x74, 0x52, 0x07, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x22, 0x98, 0x01, 0x0a, 0x0e,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x3e,
	0x0a, 0x03, 0x74, 0x72, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6b, 0x6f,
	0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x48, 0x00, 0x52, 0x03, 0x74, 0x72, 0x78, 0x12, 0x3c,
	0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x48, 0x00, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x08, 0x0a, 0x06,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x6b, 0x6f, 0x69, 0x6e,
	0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f,
	0x76, 0x32, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_koinos_account_history_account_history_proto_rawDescOnce sync.Once
	file_koinos_account_history_account_history_proto_rawDescData = file_koinos_account_history_account_history_proto_rawDesc
)

func file_koinos_account_history_account_history_proto_rawDescGZIP() []byte {
	file_koinos_account_history_account_history_proto_rawDescOnce.Do(func() {
		file_koinos_account_history_account_history_proto_rawDescData = protoimpl.X.CompressGZIP(file_koinos_account_history_account_history_proto_rawDescData)
	})
	return file_koinos_account_history_account_history_proto_rawDescData
}

var file_koinos_account_history_account_history_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_koinos_account_history_account_history_proto_goTypes = []interface{}{
	(*AccountMetadata)(nil),             // 0: koinos.account_history.account_metadata
	(*HistoryIndex)(nil),                // 1: koinos.account_history.history_index
	(*TransactionRecord)(nil),           // 2: koinos.account_history.transaction_record
	(*BlockRecord)(nil),                 // 3: koinos.account_history.block_record
	(*HistoryRecord)(nil),               // 4: koinos.account_history.history_record
	(*protocol.Transaction)(nil),        // 5: koinos.protocol.transaction
	(*protocol.TransactionReceipt)(nil), // 6: koinos.protocol.transaction_receipt
	(*protocol.BlockHeader)(nil),        // 7: koinos.protocol.block_header
	(*protocol.BlockReceipt)(nil),       // 8: koinos.protocol.block_receipt
}
var file_koinos_account_history_account_history_proto_depIdxs = []int32{
	5, // 0: koinos.account_history.transaction_record.transaction:type_name -> koinos.protocol.transaction
	6, // 1: koinos.account_history.transaction_record.receipt:type_name -> koinos.protocol.transaction_receipt
	7, // 2: koinos.account_history.block_record.header:type_name -> koinos.protocol.block_header
	8, // 3: koinos.account_history.block_record.receipt:type_name -> koinos.protocol.block_receipt
	2, // 4: koinos.account_history.history_record.trx:type_name -> koinos.account_history.transaction_record
	3, // 5: koinos.account_history.history_record.block:type_name -> koinos.account_history.block_record
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_koinos_account_history_account_history_proto_init() }
func file_koinos_account_history_account_history_proto_init() {
	if File_koinos_account_history_account_history_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_koinos_account_history_account_history_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountMetadata); i {
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
		file_koinos_account_history_account_history_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryIndex); i {
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
		file_koinos_account_history_account_history_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionRecord); i {
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
		file_koinos_account_history_account_history_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockRecord); i {
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
		file_koinos_account_history_account_history_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryRecord); i {
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
	file_koinos_account_history_account_history_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*HistoryRecord_Trx)(nil),
		(*HistoryRecord_Block)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_koinos_account_history_account_history_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_koinos_account_history_account_history_proto_goTypes,
		DependencyIndexes: file_koinos_account_history_account_history_proto_depIdxs,
		MessageInfos:      file_koinos_account_history_account_history_proto_msgTypes,
	}.Build()
	File_koinos_account_history_account_history_proto = out.File
	file_koinos_account_history_account_history_proto_rawDesc = nil
	file_koinos_account_history_account_history_proto_goTypes = nil
	file_koinos_account_history_account_history_proto_depIdxs = nil
}
