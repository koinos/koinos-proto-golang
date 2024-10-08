// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.17.3
// source: koinos/rpc/account_history/account_history_rpc.proto

package account_history

import (
	_ "github.com/koinos/koinos-proto-golang/v2/koinos"
	account_history "github.com/koinos/koinos-proto-golang/v2/koinos/account_history"
	rpc "github.com/koinos/koinos-proto-golang/v2/koinos/rpc"
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

type AccountHistoryEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SeqNum uint64 `protobuf:"varint,1,opt,name=seq_num,json=seqNum,proto3" json:"seq_num,omitempty"`
	// Types that are assignable to Record:
	//
	//	*AccountHistoryEntry_Trx
	//	*AccountHistoryEntry_Block
	Record isAccountHistoryEntry_Record `protobuf_oneof:"record"`
}

func (x *AccountHistoryEntry) Reset() {
	*x = AccountHistoryEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_rpc_account_history_account_history_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountHistoryEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountHistoryEntry) ProtoMessage() {}

func (x *AccountHistoryEntry) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_rpc_account_history_account_history_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountHistoryEntry.ProtoReflect.Descriptor instead.
func (*AccountHistoryEntry) Descriptor() ([]byte, []int) {
	return file_koinos_rpc_account_history_account_history_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *AccountHistoryEntry) GetSeqNum() uint64 {
	if x != nil {
		return x.SeqNum
	}
	return 0
}

func (m *AccountHistoryEntry) GetRecord() isAccountHistoryEntry_Record {
	if m != nil {
		return m.Record
	}
	return nil
}

func (x *AccountHistoryEntry) GetTrx() *account_history.TransactionRecord {
	if x, ok := x.GetRecord().(*AccountHistoryEntry_Trx); ok {
		return x.Trx
	}
	return nil
}

func (x *AccountHistoryEntry) GetBlock() *account_history.BlockRecord {
	if x, ok := x.GetRecord().(*AccountHistoryEntry_Block); ok {
		return x.Block
	}
	return nil
}

type isAccountHistoryEntry_Record interface {
	isAccountHistoryEntry_Record()
}

type AccountHistoryEntry_Trx struct {
	Trx *account_history.TransactionRecord `protobuf:"bytes,2,opt,name=trx,proto3,oneof"`
}

type AccountHistoryEntry_Block struct {
	Block *account_history.BlockRecord `protobuf:"bytes,3,opt,name=block,proto3,oneof"`
}

func (*AccountHistoryEntry_Trx) isAccountHistoryEntry_Record() {}

func (*AccountHistoryEntry_Block) isAccountHistoryEntry_Record() {}

type GetAccountHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address      []byte  `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	SeqNum       *uint64 `protobuf:"varint,2,opt,name=seq_num,json=seqNum,proto3,oneof" json:"seq_num,omitempty"`
	Limit        uint64  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Ascending    bool    `protobuf:"varint,4,opt,name=ascending,proto3" json:"ascending,omitempty"`
	Irreversible bool    `protobuf:"varint,5,opt,name=irreversible,proto3" json:"irreversible,omitempty"`
}

func (x *GetAccountHistoryRequest) Reset() {
	*x = GetAccountHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_rpc_account_history_account_history_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountHistoryRequest) ProtoMessage() {}

func (x *GetAccountHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_rpc_account_history_account_history_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountHistoryRequest.ProtoReflect.Descriptor instead.
func (*GetAccountHistoryRequest) Descriptor() ([]byte, []int) {
	return file_koinos_rpc_account_history_account_history_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *GetAccountHistoryRequest) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *GetAccountHistoryRequest) GetSeqNum() uint64 {
	if x != nil && x.SeqNum != nil {
		return *x.SeqNum
	}
	return 0
}

func (x *GetAccountHistoryRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAccountHistoryRequest) GetAscending() bool {
	if x != nil {
		return x.Ascending
	}
	return false
}

func (x *GetAccountHistoryRequest) GetIrreversible() bool {
	if x != nil {
		return x.Irreversible
	}
	return false
}

type GetAccountHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*AccountHistoryEntry `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *GetAccountHistoryResponse) Reset() {
	*x = GetAccountHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_rpc_account_history_account_history_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountHistoryResponse) ProtoMessage() {}

func (x *GetAccountHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_rpc_account_history_account_history_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountHistoryResponse.ProtoReflect.Descriptor instead.
func (*GetAccountHistoryResponse) Descriptor() ([]byte, []int) {
	return file_koinos_rpc_account_history_account_history_rpc_proto_rawDescGZIP(), []int{2}
}

func (x *GetAccountHistoryResponse) GetValues() []*AccountHistoryEntry {
	if x != nil {
		return x.Values
	}
	return nil
}

type AccountHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//
	//	*AccountHistoryRequest_Reserved
	//	*AccountHistoryRequest_GetAccountHistory
	Request isAccountHistoryRequest_Request `protobuf_oneof:"request"`
}

func (x *AccountHistoryRequest) Reset() {
	*x = AccountHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_rpc_account_history_account_history_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountHistoryRequest) ProtoMessage() {}

func (x *AccountHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_rpc_account_history_account_history_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountHistoryRequest.ProtoReflect.Descriptor instead.
func (*AccountHistoryRequest) Descriptor() ([]byte, []int) {
	return file_koinos_rpc_account_history_account_history_rpc_proto_rawDescGZIP(), []int{3}
}

func (m *AccountHistoryRequest) GetRequest() isAccountHistoryRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *AccountHistoryRequest) GetReserved() *rpc.ReservedRpc {
	if x, ok := x.GetRequest().(*AccountHistoryRequest_Reserved); ok {
		return x.Reserved
	}
	return nil
}

func (x *AccountHistoryRequest) GetGetAccountHistory() *GetAccountHistoryRequest {
	if x, ok := x.GetRequest().(*AccountHistoryRequest_GetAccountHistory); ok {
		return x.GetAccountHistory
	}
	return nil
}

type isAccountHistoryRequest_Request interface {
	isAccountHistoryRequest_Request()
}

type AccountHistoryRequest_Reserved struct {
	Reserved *rpc.ReservedRpc `protobuf:"bytes,1,opt,name=reserved,proto3,oneof"`
}

type AccountHistoryRequest_GetAccountHistory struct {
	GetAccountHistory *GetAccountHistoryRequest `protobuf:"bytes,2,opt,name=get_account_history,json=getAccountHistory,proto3,oneof"`
}

func (*AccountHistoryRequest_Reserved) isAccountHistoryRequest_Request() {}

func (*AccountHistoryRequest_GetAccountHistory) isAccountHistoryRequest_Request() {}

type AccountHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//
	//	*AccountHistoryResponse_Reserved
	//	*AccountHistoryResponse_Error
	//	*AccountHistoryResponse_GetAccountHistory
	Response isAccountHistoryResponse_Response `protobuf_oneof:"response"`
}

func (x *AccountHistoryResponse) Reset() {
	*x = AccountHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_rpc_account_history_account_history_rpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountHistoryResponse) ProtoMessage() {}

func (x *AccountHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_rpc_account_history_account_history_rpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountHistoryResponse.ProtoReflect.Descriptor instead.
func (*AccountHistoryResponse) Descriptor() ([]byte, []int) {
	return file_koinos_rpc_account_history_account_history_rpc_proto_rawDescGZIP(), []int{4}
}

func (m *AccountHistoryResponse) GetResponse() isAccountHistoryResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *AccountHistoryResponse) GetReserved() *rpc.ReservedRpc {
	if x, ok := x.GetResponse().(*AccountHistoryResponse_Reserved); ok {
		return x.Reserved
	}
	return nil
}

func (x *AccountHistoryResponse) GetError() *rpc.ErrorStatus {
	if x, ok := x.GetResponse().(*AccountHistoryResponse_Error); ok {
		return x.Error
	}
	return nil
}

func (x *AccountHistoryResponse) GetGetAccountHistory() *GetAccountHistoryResponse {
	if x, ok := x.GetResponse().(*AccountHistoryResponse_GetAccountHistory); ok {
		return x.GetAccountHistory
	}
	return nil
}

type isAccountHistoryResponse_Response interface {
	isAccountHistoryResponse_Response()
}

type AccountHistoryResponse_Reserved struct {
	Reserved *rpc.ReservedRpc `protobuf:"bytes,1,opt,name=reserved,proto3,oneof"`
}

type AccountHistoryResponse_Error struct {
	Error *rpc.ErrorStatus `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

type AccountHistoryResponse_GetAccountHistory struct {
	GetAccountHistory *GetAccountHistoryResponse `protobuf:"bytes,3,opt,name=get_account_history,json=getAccountHistory,proto3,oneof"`
}

func (*AccountHistoryResponse_Reserved) isAccountHistoryResponse_Response() {}

func (*AccountHistoryResponse_Error) isAccountHistoryResponse_Response() {}

func (*AccountHistoryResponse_GetAccountHistory) isAccountHistoryResponse_Response() {}

var File_koinos_rpc_account_history_account_history_rpc_proto protoreflect.FileDescriptor

var file_koinos_rpc_account_history_account_history_rpc_proto_rawDesc = []byte{
	0x0a, 0x34, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x1a, 0x2c, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x72,
	0x70, 0x63, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a,
	0x15, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x07, 0x73, 0x65, 0x71, 0x5f, 0x6e, 0x75,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x06, 0x73, 0x65, 0x71,
	0x4e, 0x75, 0x6d, 0x12, 0x3e, 0x0a, 0x03, 0x74, 0x72, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x48, 0x00, 0x52, 0x03,
	0x74, 0x72, 0x78, 0x12, 0x3c, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x48, 0x00, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x22, 0xc7, 0x01, 0x0a, 0x1b,
	0x67, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x04, 0x80, 0xb5,
	0x18, 0x06, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x07, 0x73,
	0x65, 0x71, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01,
	0x48, 0x00, 0x52, 0x06, 0x73, 0x65, 0x71, 0x4e, 0x75, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x73, 0x63, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x61, 0x73, 0x63, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x72, 0x72, 0x65, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x72, 0x72,
	0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x73, 0x65,
	0x71, 0x5f, 0x6e, 0x75, 0x6d, 0x22, 0x69, 0x0a, 0x1c, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x22, 0xc7, 0x01, 0x0a, 0x17, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x64, 0x5f, 0x72, 0x70, 0x63, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x64, 0x12, 0x69, 0x0a, 0x13, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x37, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x67,
	0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x11, 0x67, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x42,
	0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xfc, 0x01, 0x0a, 0x18, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6b, 0x6f, 0x69, 0x6e,
	0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f,
	0x72, 0x70, 0x63, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x12,
	0x30, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x6a, 0x0a, 0x13, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38,
	0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x67, 0x65, 0x74, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x11, 0x67, 0x65, 0x74, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x0a, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x6b,
	0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2f, 0x76, 0x32, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_koinos_rpc_account_history_account_history_rpc_proto_rawDescOnce sync.Once
	file_koinos_rpc_account_history_account_history_rpc_proto_rawDescData = file_koinos_rpc_account_history_account_history_rpc_proto_rawDesc
)

func file_koinos_rpc_account_history_account_history_rpc_proto_rawDescGZIP() []byte {
	file_koinos_rpc_account_history_account_history_rpc_proto_rawDescOnce.Do(func() {
		file_koinos_rpc_account_history_account_history_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_koinos_rpc_account_history_account_history_rpc_proto_rawDescData)
	})
	return file_koinos_rpc_account_history_account_history_rpc_proto_rawDescData
}

var file_koinos_rpc_account_history_account_history_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_koinos_rpc_account_history_account_history_rpc_proto_goTypes = []interface{}{
	(*AccountHistoryEntry)(nil),               // 0: koinos.rpc.account_history.account_history_entry
	(*GetAccountHistoryRequest)(nil),          // 1: koinos.rpc.account_history.get_account_history_request
	(*GetAccountHistoryResponse)(nil),         // 2: koinos.rpc.account_history.get_account_history_response
	(*AccountHistoryRequest)(nil),             // 3: koinos.rpc.account_history.account_history_request
	(*AccountHistoryResponse)(nil),            // 4: koinos.rpc.account_history.account_history_response
	(*account_history.TransactionRecord)(nil), // 5: koinos.account_history.transaction_record
	(*account_history.BlockRecord)(nil),       // 6: koinos.account_history.block_record
	(*rpc.ReservedRpc)(nil),                   // 7: koinos.rpc.reserved_rpc
	(*rpc.ErrorStatus)(nil),                   // 8: koinos.rpc.error_status
}
var file_koinos_rpc_account_history_account_history_rpc_proto_depIdxs = []int32{
	5, // 0: koinos.rpc.account_history.account_history_entry.trx:type_name -> koinos.account_history.transaction_record
	6, // 1: koinos.rpc.account_history.account_history_entry.block:type_name -> koinos.account_history.block_record
	0, // 2: koinos.rpc.account_history.get_account_history_response.values:type_name -> koinos.rpc.account_history.account_history_entry
	7, // 3: koinos.rpc.account_history.account_history_request.reserved:type_name -> koinos.rpc.reserved_rpc
	1, // 4: koinos.rpc.account_history.account_history_request.get_account_history:type_name -> koinos.rpc.account_history.get_account_history_request
	7, // 5: koinos.rpc.account_history.account_history_response.reserved:type_name -> koinos.rpc.reserved_rpc
	8, // 6: koinos.rpc.account_history.account_history_response.error:type_name -> koinos.rpc.error_status
	2, // 7: koinos.rpc.account_history.account_history_response.get_account_history:type_name -> koinos.rpc.account_history.get_account_history_response
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_koinos_rpc_account_history_account_history_rpc_proto_init() }
func file_koinos_rpc_account_history_account_history_rpc_proto_init() {
	if File_koinos_rpc_account_history_account_history_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_koinos_rpc_account_history_account_history_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountHistoryEntry); i {
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
		file_koinos_rpc_account_history_account_history_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountHistoryRequest); i {
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
		file_koinos_rpc_account_history_account_history_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccountHistoryResponse); i {
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
		file_koinos_rpc_account_history_account_history_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountHistoryRequest); i {
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
		file_koinos_rpc_account_history_account_history_rpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountHistoryResponse); i {
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
	file_koinos_rpc_account_history_account_history_rpc_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*AccountHistoryEntry_Trx)(nil),
		(*AccountHistoryEntry_Block)(nil),
	}
	file_koinos_rpc_account_history_account_history_rpc_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_koinos_rpc_account_history_account_history_rpc_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*AccountHistoryRequest_Reserved)(nil),
		(*AccountHistoryRequest_GetAccountHistory)(nil),
	}
	file_koinos_rpc_account_history_account_history_rpc_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*AccountHistoryResponse_Reserved)(nil),
		(*AccountHistoryResponse_Error)(nil),
		(*AccountHistoryResponse_GetAccountHistory)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_koinos_rpc_account_history_account_history_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_koinos_rpc_account_history_account_history_rpc_proto_goTypes,
		DependencyIndexes: file_koinos_rpc_account_history_account_history_rpc_proto_depIdxs,
		MessageInfos:      file_koinos_rpc_account_history_account_history_rpc_proto_msgTypes,
	}.Build()
	File_koinos_rpc_account_history_account_history_rpc_proto = out.File
	file_koinos_rpc_account_history_account_history_rpc_proto_rawDesc = nil
	file_koinos_rpc_account_history_account_history_rpc_proto_goTypes = nil
	file_koinos_rpc_account_history_account_history_rpc_proto_depIdxs = nil
}
