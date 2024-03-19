// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.17.3
// source: koinos/transaction_store/rpc.proto

package transaction_store

import (
	koinos "github.com/koinos/koinos-proto-golang/koinos"
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

type GetTransactionsByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionIds [][]byte `protobuf:"bytes,1,rep,name=transaction_ids,json=transactionIds,proto3" json:"transaction_ids,omitempty"`
}

func (x *GetTransactionsByIdRequest) Reset() {
	*x = GetTransactionsByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_transaction_store_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionsByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionsByIdRequest) ProtoMessage() {}

func (x *GetTransactionsByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_transaction_store_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionsByIdRequest.ProtoReflect.Descriptor instead.
func (*GetTransactionsByIdRequest) Descriptor() ([]byte, []int) {
	return file_koinos_transaction_store_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *GetTransactionsByIdRequest) GetTransactionIds() [][]byte {
	if x != nil {
		return x.TransactionIds
	}
	return nil
}

type GetTransactionsByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*TransactionItem `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *GetTransactionsByIdResponse) Reset() {
	*x = GetTransactionsByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_transaction_store_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionsByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionsByIdResponse) ProtoMessage() {}

func (x *GetTransactionsByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_transaction_store_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionsByIdResponse.ProtoReflect.Descriptor instead.
func (*GetTransactionsByIdResponse) Descriptor() ([]byte, []int) {
	return file_koinos_transaction_store_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *GetTransactionsByIdResponse) GetTransactions() []*TransactionItem {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type TransactionStoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//	*TransactionStoreRequest_Reserved
	//	*TransactionStoreRequest_GetTransactionsById
	Request isTransactionStoreRequest_Request `protobuf_oneof:"request"`
}

func (x *TransactionStoreRequest) Reset() {
	*x = TransactionStoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_transaction_store_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionStoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionStoreRequest) ProtoMessage() {}

func (x *TransactionStoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_transaction_store_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionStoreRequest.ProtoReflect.Descriptor instead.
func (*TransactionStoreRequest) Descriptor() ([]byte, []int) {
	return file_koinos_transaction_store_rpc_proto_rawDescGZIP(), []int{2}
}

func (m *TransactionStoreRequest) GetRequest() isTransactionStoreRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *TransactionStoreRequest) GetReserved() *koinos.ReservedRpc {
	if x, ok := x.GetRequest().(*TransactionStoreRequest_Reserved); ok {
		return x.Reserved
	}
	return nil
}

func (x *TransactionStoreRequest) GetGetTransactionsById() *GetTransactionsByIdRequest {
	if x, ok := x.GetRequest().(*TransactionStoreRequest_GetTransactionsById); ok {
		return x.GetTransactionsById
	}
	return nil
}

type isTransactionStoreRequest_Request interface {
	isTransactionStoreRequest_Request()
}

type TransactionStoreRequest_Reserved struct {
	Reserved *koinos.ReservedRpc `protobuf:"bytes,1,opt,name=reserved,proto3,oneof"`
}

type TransactionStoreRequest_GetTransactionsById struct {
	GetTransactionsById *GetTransactionsByIdRequest `protobuf:"bytes,2,opt,name=get_transactions_by_id,json=getTransactionsById,proto3,oneof"`
}

func (*TransactionStoreRequest_Reserved) isTransactionStoreRequest_Request() {}

func (*TransactionStoreRequest_GetTransactionsById) isTransactionStoreRequest_Request() {}

type TransactionStoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Response:
	//	*TransactionStoreResponse_Reserved
	//	*TransactionStoreResponse_Error
	//	*TransactionStoreResponse_GetTransactionsById
	Response isTransactionStoreResponse_Response `protobuf_oneof:"response"`
}

func (x *TransactionStoreResponse) Reset() {
	*x = TransactionStoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_transaction_store_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionStoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionStoreResponse) ProtoMessage() {}

func (x *TransactionStoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_transaction_store_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionStoreResponse.ProtoReflect.Descriptor instead.
func (*TransactionStoreResponse) Descriptor() ([]byte, []int) {
	return file_koinos_transaction_store_rpc_proto_rawDescGZIP(), []int{3}
}

func (m *TransactionStoreResponse) GetResponse() isTransactionStoreResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *TransactionStoreResponse) GetReserved() *koinos.ReservedRpc {
	if x, ok := x.GetResponse().(*TransactionStoreResponse_Reserved); ok {
		return x.Reserved
	}
	return nil
}

func (x *TransactionStoreResponse) GetError() *koinos.ErrorStatus {
	if x, ok := x.GetResponse().(*TransactionStoreResponse_Error); ok {
		return x.Error
	}
	return nil
}

func (x *TransactionStoreResponse) GetGetTransactionsById() *GetTransactionsByIdResponse {
	if x, ok := x.GetResponse().(*TransactionStoreResponse_GetTransactionsById); ok {
		return x.GetTransactionsById
	}
	return nil
}

type isTransactionStoreResponse_Response interface {
	isTransactionStoreResponse_Response()
}

type TransactionStoreResponse_Reserved struct {
	Reserved *koinos.ReservedRpc `protobuf:"bytes,1,opt,name=reserved,proto3,oneof"`
}

type TransactionStoreResponse_Error struct {
	Error *koinos.ErrorStatus `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

type TransactionStoreResponse_GetTransactionsById struct {
	GetTransactionsById *GetTransactionsByIdResponse `protobuf:"bytes,3,opt,name=get_transactions_by_id,json=getTransactionsById,proto3,oneof"`
}

func (*TransactionStoreResponse_Reserved) isTransactionStoreResponse_Response() {}

func (*TransactionStoreResponse_Error) isTransactionStoreResponse_Response() {}

func (*TransactionStoreResponse_GetTransactionsById) isTransactionStoreResponse_Response() {}

var File_koinos_transaction_store_rpc_proto protoreflect.FileDescriptor

var file_koinos_transaction_store_rpc_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x1a, 0x14,
	0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x1e, 0x67, 0x65, 0x74, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x62, 0x79, 0x5f,
	0x69, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0c, 0x42, 0x04, 0x80, 0xb5, 0x18, 0x04, 0x52, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x22, 0x71, 0x0a, 0x1f, 0x67, 0x65, 0x74,
	0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x62, 0x79,
	0x5f, 0x69, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x0c,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x52, 0x0c,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xcb, 0x01, 0x0a,
	0x19, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x72, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b,
	0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x72,
	0x70, 0x63, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x12, 0x6f,
	0x0a, 0x16, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38,
	0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x00, 0x52, 0x13, 0x67, 0x65, 0x74, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x49, 0x64, 0x42,
	0x09, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xfc, 0x01, 0x0a, 0x1a, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b, 0x6f,
	0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x72, 0x70,
	0x63, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x12, 0x2c, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6b,
	0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x70, 0x0a, 0x16, 0x67,
	0x65, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f,
	0x62, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x6b, 0x6f,
	0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x13, 0x67, 0x65, 0x74, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x79, 0x49, 0x64, 0x42, 0x0a, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x6b,
	0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_koinos_transaction_store_rpc_proto_rawDescOnce sync.Once
	file_koinos_transaction_store_rpc_proto_rawDescData = file_koinos_transaction_store_rpc_proto_rawDesc
)

func file_koinos_transaction_store_rpc_proto_rawDescGZIP() []byte {
	file_koinos_transaction_store_rpc_proto_rawDescOnce.Do(func() {
		file_koinos_transaction_store_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_koinos_transaction_store_rpc_proto_rawDescData)
	})
	return file_koinos_transaction_store_rpc_proto_rawDescData
}

var file_koinos_transaction_store_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_koinos_transaction_store_rpc_proto_goTypes = []interface{}{
	(*GetTransactionsByIdRequest)(nil),  // 0: koinos.transaction_store.get_transactions_by_id_request
	(*GetTransactionsByIdResponse)(nil), // 1: koinos.transaction_store.get_transactions_by_id_response
	(*TransactionStoreRequest)(nil),     // 2: koinos.transaction_store.transaction_store_request
	(*TransactionStoreResponse)(nil),    // 3: koinos.transaction_store.transaction_store_response
	(*TransactionItem)(nil),             // 4: koinos.transaction_store.transaction_item
	(*koinos.ReservedRpc)(nil),          // 5: koinos.reserved_rpc
	(*koinos.ErrorStatus)(nil),          // 6: koinos.error_status
}
var file_koinos_transaction_store_rpc_proto_depIdxs = []int32{
	4, // 0: koinos.transaction_store.get_transactions_by_id_response.transactions:type_name -> koinos.transaction_store.transaction_item
	5, // 1: koinos.transaction_store.transaction_store_request.reserved:type_name -> koinos.reserved_rpc
	0, // 2: koinos.transaction_store.transaction_store_request.get_transactions_by_id:type_name -> koinos.transaction_store.get_transactions_by_id_request
	5, // 3: koinos.transaction_store.transaction_store_response.reserved:type_name -> koinos.reserved_rpc
	6, // 4: koinos.transaction_store.transaction_store_response.error:type_name -> koinos.error_status
	1, // 5: koinos.transaction_store.transaction_store_response.get_transactions_by_id:type_name -> koinos.transaction_store.get_transactions_by_id_response
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_koinos_transaction_store_rpc_proto_init() }
func file_koinos_transaction_store_rpc_proto_init() {
	if File_koinos_transaction_store_rpc_proto != nil {
		return
	}
	file_koinos_transaction_store_transaction_store_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_koinos_transaction_store_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionsByIdRequest); i {
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
		file_koinos_transaction_store_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionsByIdResponse); i {
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
		file_koinos_transaction_store_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionStoreRequest); i {
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
		file_koinos_transaction_store_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionStoreResponse); i {
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
	file_koinos_transaction_store_rpc_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*TransactionStoreRequest_Reserved)(nil),
		(*TransactionStoreRequest_GetTransactionsById)(nil),
	}
	file_koinos_transaction_store_rpc_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*TransactionStoreResponse_Reserved)(nil),
		(*TransactionStoreResponse_Error)(nil),
		(*TransactionStoreResponse_GetTransactionsById)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_koinos_transaction_store_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_koinos_transaction_store_rpc_proto_goTypes,
		DependencyIndexes: file_koinos_transaction_store_rpc_proto_depIdxs,
		MessageInfos:      file_koinos_transaction_store_rpc_proto_msgTypes,
	}.Build()
	File_koinos_transaction_store_rpc_proto = out.File
	file_koinos_transaction_store_rpc_proto_rawDesc = nil
	file_koinos_transaction_store_rpc_proto_goTypes = nil
	file_koinos_transaction_store_rpc_proto_depIdxs = nil
}
