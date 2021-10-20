// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: koinos/contracts/token/token.proto

package token

import (
	_ "github.com/koinos/koinos-proto-golang/koinos"
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

type NameArguments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NameArguments) Reset() {
	*x = NameArguments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_token_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameArguments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameArguments) ProtoMessage() {}

func (x *NameArguments) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_token_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameArguments.ProtoReflect.Descriptor instead.
func (*NameArguments) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_token_token_proto_rawDescGZIP(), []int{0}
}

type NameResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *NameResult) Reset() {
	*x = NameResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_token_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameResult) ProtoMessage() {}

func (x *NameResult) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_token_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NameResult.ProtoReflect.Descriptor instead.
func (*NameResult) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_token_token_proto_rawDescGZIP(), []int{1}
}

func (x *NameResult) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SymbolArguments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SymbolArguments) Reset() {
	*x = SymbolArguments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_token_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SymbolArguments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SymbolArguments) ProtoMessage() {}

func (x *SymbolArguments) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_token_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SymbolArguments.ProtoReflect.Descriptor instead.
func (*SymbolArguments) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_token_token_proto_rawDescGZIP(), []int{2}
}

type SymbolResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SymbolResult) Reset() {
	*x = SymbolResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_token_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SymbolResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SymbolResult) ProtoMessage() {}

func (x *SymbolResult) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_token_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SymbolResult.ProtoReflect.Descriptor instead.
func (*SymbolResult) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_token_token_proto_rawDescGZIP(), []int{3}
}

func (x *SymbolResult) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type DecimalsArguments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DecimalsArguments) Reset() {
	*x = DecimalsArguments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_token_token_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecimalsArguments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecimalsArguments) ProtoMessage() {}

func (x *DecimalsArguments) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_token_token_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecimalsArguments.ProtoReflect.Descriptor instead.
func (*DecimalsArguments) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_token_token_proto_rawDescGZIP(), []int{4}
}

type DecimalsResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DecimalsResult) Reset() {
	*x = DecimalsResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_token_token_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecimalsResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecimalsResult) ProtoMessage() {}

func (x *DecimalsResult) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_token_token_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecimalsResult.ProtoReflect.Descriptor instead.
func (*DecimalsResult) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_token_token_proto_rawDescGZIP(), []int{5}
}

func (x *DecimalsResult) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type TotalSupplyArguments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TotalSupplyArguments) Reset() {
	*x = TotalSupplyArguments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_token_token_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TotalSupplyArguments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TotalSupplyArguments) ProtoMessage() {}

func (x *TotalSupplyArguments) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_token_token_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TotalSupplyArguments.ProtoReflect.Descriptor instead.
func (*TotalSupplyArguments) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_token_token_proto_rawDescGZIP(), []int{6}
}

type TotalSupplyResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TotalSupplyResult) Reset() {
	*x = TotalSupplyResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_token_token_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TotalSupplyResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TotalSupplyResult) ProtoMessage() {}

func (x *TotalSupplyResult) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_token_token_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TotalSupplyResult.ProtoReflect.Descriptor instead.
func (*TotalSupplyResult) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_token_token_proto_rawDescGZIP(), []int{7}
}

func (x *TotalSupplyResult) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type BalanceOfArguments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner []byte `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *BalanceOfArguments) Reset() {
	*x = BalanceOfArguments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_token_token_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceOfArguments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceOfArguments) ProtoMessage() {}

func (x *BalanceOfArguments) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_token_token_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceOfArguments.ProtoReflect.Descriptor instead.
func (*BalanceOfArguments) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_token_token_proto_rawDescGZIP(), []int{8}
}

func (x *BalanceOfArguments) GetOwner() []byte {
	if x != nil {
		return x.Owner
	}
	return nil
}

type BalanceOfResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *BalanceOfResult) Reset() {
	*x = BalanceOfResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_token_token_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceOfResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceOfResult) ProtoMessage() {}

func (x *BalanceOfResult) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_token_token_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceOfResult.ProtoReflect.Descriptor instead.
func (*BalanceOfResult) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_token_token_proto_rawDescGZIP(), []int{9}
}

func (x *BalanceOfResult) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type TransferArguments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From  []byte `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To    []byte `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Value uint64 `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TransferArguments) Reset() {
	*x = TransferArguments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_token_token_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferArguments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferArguments) ProtoMessage() {}

func (x *TransferArguments) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_token_token_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferArguments.ProtoReflect.Descriptor instead.
func (*TransferArguments) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_token_token_proto_rawDescGZIP(), []int{10}
}

func (x *TransferArguments) GetFrom() []byte {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *TransferArguments) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *TransferArguments) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type TransferResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TransferResult) Reset() {
	*x = TransferResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_token_token_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferResult) ProtoMessage() {}

func (x *TransferResult) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_token_token_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferResult.ProtoReflect.Descriptor instead.
func (*TransferResult) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_token_token_proto_rawDescGZIP(), []int{11}
}

func (x *TransferResult) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type MintArguments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To    []byte `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Value uint64 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MintArguments) Reset() {
	*x = MintArguments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_token_token_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MintArguments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MintArguments) ProtoMessage() {}

func (x *MintArguments) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_token_token_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MintArguments.ProtoReflect.Descriptor instead.
func (*MintArguments) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_token_token_proto_rawDescGZIP(), []int{12}
}

func (x *MintArguments) GetTo() []byte {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *MintArguments) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type MintResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MintResult) Reset() {
	*x = MintResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_token_token_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MintResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MintResult) ProtoMessage() {}

func (x *MintResult) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_token_token_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MintResult.ProtoReflect.Descriptor instead.
func (*MintResult) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_token_token_proto_rawDescGZIP(), []int{13}
}

func (x *MintResult) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type BalanceObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *BalanceObject) Reset() {
	*x = BalanceObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_token_token_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceObject) ProtoMessage() {}

func (x *BalanceObject) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_token_token_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceObject.ProtoReflect.Descriptor instead.
func (*BalanceObject) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_token_token_proto_rawDescGZIP(), []int{14}
}

func (x *BalanceObject) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type ManaBalanceObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance        uint64 `protobuf:"varint,1,opt,name=balance,proto3" json:"balance,omitempty"`
	Mana           uint64 `protobuf:"varint,2,opt,name=mana,proto3" json:"mana,omitempty"`
	LastManaUpdate uint64 `protobuf:"varint,3,opt,name=last_mana_update,json=lastManaUpdate,proto3" json:"last_mana_update,omitempty"`
}

func (x *ManaBalanceObject) Reset() {
	*x = ManaBalanceObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_token_token_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManaBalanceObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManaBalanceObject) ProtoMessage() {}

func (x *ManaBalanceObject) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_token_token_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManaBalanceObject.ProtoReflect.Descriptor instead.
func (*ManaBalanceObject) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_token_token_proto_rawDescGZIP(), []int{15}
}

func (x *ManaBalanceObject) GetBalance() uint64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *ManaBalanceObject) GetMana() uint64 {
	if x != nil {
		return x.Mana
	}
	return 0
}

func (x *ManaBalanceObject) GetLastManaUpdate() uint64 {
	if x != nil {
		return x.LastManaUpdate
	}
	return 0
}

var File_koinos_contracts_token_token_proto protoreflect.FileDescriptor

var file_koinos_contracts_token_token_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x73, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x14, 0x6b, 0x6f,
	0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x10, 0x0a, 0x0e, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x61, 0x72, 0x67, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x23, 0x0a, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x5f, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x25, 0x0a,
	0x0d, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73,
	0x5f, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x27, 0x0a, 0x0f, 0x64, 0x65,
	0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x75, 0x70,
	0x70, 0x6c, 0x79, 0x5f, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x2f, 0x0a,
	0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x32,
	0x0a, 0x14, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x61, 0x72, 0x67,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x04, 0x80, 0xb5, 0x18, 0x06, 0x52, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x22, 0x2d, 0x0a, 0x11, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6f, 0x66,
	0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x5e, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x61, 0x72,
	0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x18, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x04, 0x80, 0xb5, 0x18, 0x06, 0x52, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x12, 0x14, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x04, 0x80,
	0xb5, 0x18, 0x06, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x18, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x27, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x40, 0x0a, 0x0e, 0x6d, 0x69,
	0x6e, 0x74, 0x5f, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x02,
	0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x04, 0x80, 0xb5, 0x18, 0x06, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x18, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x23, 0x0a, 0x0b,
	0x6d, 0x69, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x2a, 0x0a, 0x0e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x79, 0x0a,
	0x13, 0x6d, 0x61, 0x6e, 0x61, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x16, 0x0a, 0x04, 0x6d, 0x61, 0x6e, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x02, 0x30, 0x01, 0x52, 0x04, 0x6d, 0x61, 0x6e, 0x61, 0x12, 0x2c, 0x0a, 0x10, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x0e, 0x6c, 0x61, 0x73, 0x74, 0x4d, 0x61,
	0x6e, 0x61, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x6b, 0x6f,
	0x69, 0x6e, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x73, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_koinos_contracts_token_token_proto_rawDescOnce sync.Once
	file_koinos_contracts_token_token_proto_rawDescData = file_koinos_contracts_token_token_proto_rawDesc
)

func file_koinos_contracts_token_token_proto_rawDescGZIP() []byte {
	file_koinos_contracts_token_token_proto_rawDescOnce.Do(func() {
		file_koinos_contracts_token_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_koinos_contracts_token_token_proto_rawDescData)
	})
	return file_koinos_contracts_token_token_proto_rawDescData
}

var file_koinos_contracts_token_token_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_koinos_contracts_token_token_proto_goTypes = []interface{}{
	(*NameArguments)(nil),        // 0: koinos.contracts.token.name_arguments
	(*NameResult)(nil),           // 1: koinos.contracts.token.name_result
	(*SymbolArguments)(nil),      // 2: koinos.contracts.token.symbol_arguments
	(*SymbolResult)(nil),         // 3: koinos.contracts.token.symbol_result
	(*DecimalsArguments)(nil),    // 4: koinos.contracts.token.decimals_arguments
	(*DecimalsResult)(nil),       // 5: koinos.contracts.token.decimals_result
	(*TotalSupplyArguments)(nil), // 6: koinos.contracts.token.total_supply_arguments
	(*TotalSupplyResult)(nil),    // 7: koinos.contracts.token.total_supply_result
	(*BalanceOfArguments)(nil),   // 8: koinos.contracts.token.balance_of_arguments
	(*BalanceOfResult)(nil),      // 9: koinos.contracts.token.balance_of_result
	(*TransferArguments)(nil),    // 10: koinos.contracts.token.transfer_arguments
	(*TransferResult)(nil),       // 11: koinos.contracts.token.transfer_result
	(*MintArguments)(nil),        // 12: koinos.contracts.token.mint_arguments
	(*MintResult)(nil),           // 13: koinos.contracts.token.mint_result
	(*BalanceObject)(nil),        // 14: koinos.contracts.token.balance_object
	(*ManaBalanceObject)(nil),    // 15: koinos.contracts.token.mana_balance_object
}
var file_koinos_contracts_token_token_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_koinos_contracts_token_token_proto_init() }
func file_koinos_contracts_token_token_proto_init() {
	if File_koinos_contracts_token_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_koinos_contracts_token_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameArguments); i {
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
		file_koinos_contracts_token_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NameResult); i {
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
		file_koinos_contracts_token_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SymbolArguments); i {
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
		file_koinos_contracts_token_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SymbolResult); i {
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
		file_koinos_contracts_token_token_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecimalsArguments); i {
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
		file_koinos_contracts_token_token_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecimalsResult); i {
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
		file_koinos_contracts_token_token_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TotalSupplyArguments); i {
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
		file_koinos_contracts_token_token_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TotalSupplyResult); i {
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
		file_koinos_contracts_token_token_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceOfArguments); i {
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
		file_koinos_contracts_token_token_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceOfResult); i {
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
		file_koinos_contracts_token_token_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferArguments); i {
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
		file_koinos_contracts_token_token_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferResult); i {
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
		file_koinos_contracts_token_token_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MintArguments); i {
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
		file_koinos_contracts_token_token_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MintResult); i {
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
		file_koinos_contracts_token_token_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceObject); i {
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
		file_koinos_contracts_token_token_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManaBalanceObject); i {
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
			RawDescriptor: file_koinos_contracts_token_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_koinos_contracts_token_token_proto_goTypes,
		DependencyIndexes: file_koinos_contracts_token_token_proto_depIdxs,
		MessageInfos:      file_koinos_contracts_token_token_proto_msgTypes,
	}.Build()
	File_koinos_contracts_token_token_proto = out.File
	file_koinos_contracts_token_token_proto_rawDesc = nil
	file_koinos_contracts_token_token_proto_goTypes = nil
	file_koinos_contracts_token_token_proto_depIdxs = nil
}
