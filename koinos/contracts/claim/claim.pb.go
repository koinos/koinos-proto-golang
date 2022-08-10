// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: koinos/contracts/claim/claim.proto

package claim

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

type ClaimInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalEthAccounts   uint32 `protobuf:"varint,1,opt,name=total_eth_accounts,json=totalEthAccounts,proto3" json:"total_eth_accounts,omitempty"`
	EthAccountsClaimed uint32 `protobuf:"varint,2,opt,name=eth_accounts_claimed,json=ethAccountsClaimed,proto3" json:"eth_accounts_claimed,omitempty"`
	TotalKoin          uint64 `protobuf:"varint,3,opt,name=total_koin,json=totalKoin,proto3" json:"total_koin,omitempty"`
	KoinClaimed        uint64 `protobuf:"varint,4,opt,name=koin_claimed,json=koinClaimed,proto3" json:"koin_claimed,omitempty"`
}

func (x *ClaimInfo) Reset() {
	*x = ClaimInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_claim_claim_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimInfo) ProtoMessage() {}

func (x *ClaimInfo) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_claim_claim_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimInfo.ProtoReflect.Descriptor instead.
func (*ClaimInfo) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_claim_claim_proto_rawDescGZIP(), []int{0}
}

func (x *ClaimInfo) GetTotalEthAccounts() uint32 {
	if x != nil {
		return x.TotalEthAccounts
	}
	return 0
}

func (x *ClaimInfo) GetEthAccountsClaimed() uint32 {
	if x != nil {
		return x.EthAccountsClaimed
	}
	return 0
}

func (x *ClaimInfo) GetTotalKoin() uint64 {
	if x != nil {
		return x.TotalKoin
	}
	return 0
}

func (x *ClaimInfo) GetKoinClaimed() uint64 {
	if x != nil {
		return x.KoinClaimed
	}
	return 0
}

type ClaimStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenAmount uint64 `protobuf:"varint,1,opt,name=token_amount,json=tokenAmount,proto3" json:"token_amount,omitempty"`
	Claimed     bool   `protobuf:"varint,2,opt,name=claimed,proto3" json:"claimed,omitempty"`
}

func (x *ClaimStatus) Reset() {
	*x = ClaimStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_claim_claim_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimStatus) ProtoMessage() {}

func (x *ClaimStatus) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_claim_claim_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimStatus.ProtoReflect.Descriptor instead.
func (*ClaimStatus) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_claim_claim_proto_rawDescGZIP(), []int{1}
}

func (x *ClaimStatus) GetTokenAmount() uint64 {
	if x != nil {
		return x.TokenAmount
	}
	return 0
}

func (x *ClaimStatus) GetClaimed() bool {
	if x != nil {
		return x.Claimed
	}
	return false
}

type ClaimArguments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EthAddress  []byte `protobuf:"bytes,1,opt,name=eth_address,json=ethAddress,proto3" json:"eth_address,omitempty"`
	KoinAddress []byte `protobuf:"bytes,2,opt,name=koin_address,json=koinAddress,proto3" json:"koin_address,omitempty"`
	Signature   []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *ClaimArguments) Reset() {
	*x = ClaimArguments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_claim_claim_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimArguments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimArguments) ProtoMessage() {}

func (x *ClaimArguments) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_claim_claim_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimArguments.ProtoReflect.Descriptor instead.
func (*ClaimArguments) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_claim_claim_proto_rawDescGZIP(), []int{2}
}

func (x *ClaimArguments) GetEthAddress() []byte {
	if x != nil {
		return x.EthAddress
	}
	return nil
}

func (x *ClaimArguments) GetKoinAddress() []byte {
	if x != nil {
		return x.KoinAddress
	}
	return nil
}

func (x *ClaimArguments) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type ClaimResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClaimResult) Reset() {
	*x = ClaimResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_claim_claim_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClaimResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClaimResult) ProtoMessage() {}

func (x *ClaimResult) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_claim_claim_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClaimResult.ProtoReflect.Descriptor instead.
func (*ClaimResult) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_claim_claim_proto_rawDescGZIP(), []int{3}
}

type GetInfoArguments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetInfoArguments) Reset() {
	*x = GetInfoArguments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_claim_claim_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoArguments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoArguments) ProtoMessage() {}

func (x *GetInfoArguments) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_claim_claim_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoArguments.ProtoReflect.Descriptor instead.
func (*GetInfoArguments) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_claim_claim_proto_rawDescGZIP(), []int{4}
}

type GetInfoResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *ClaimInfo `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetInfoResult) Reset() {
	*x = GetInfoResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_claim_claim_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoResult) ProtoMessage() {}

func (x *GetInfoResult) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_claim_claim_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoResult.ProtoReflect.Descriptor instead.
func (*GetInfoResult) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_claim_claim_proto_rawDescGZIP(), []int{5}
}

func (x *GetInfoResult) GetValue() *ClaimInfo {
	if x != nil {
		return x.Value
	}
	return nil
}

type CheckClaimArguments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EthAddress []byte `protobuf:"bytes,1,opt,name=eth_address,json=ethAddress,proto3" json:"eth_address,omitempty"`
}

func (x *CheckClaimArguments) Reset() {
	*x = CheckClaimArguments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_claim_claim_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckClaimArguments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckClaimArguments) ProtoMessage() {}

func (x *CheckClaimArguments) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_claim_claim_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckClaimArguments.ProtoReflect.Descriptor instead.
func (*CheckClaimArguments) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_claim_claim_proto_rawDescGZIP(), []int{6}
}

func (x *CheckClaimArguments) GetEthAddress() []byte {
	if x != nil {
		return x.EthAddress
	}
	return nil
}

type CheckClaimResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *ClaimStatus `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CheckClaimResult) Reset() {
	*x = CheckClaimResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_claim_claim_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckClaimResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckClaimResult) ProtoMessage() {}

func (x *CheckClaimResult) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_claim_claim_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckClaimResult.ProtoReflect.Descriptor instead.
func (*CheckClaimResult) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_claim_claim_proto_rawDescGZIP(), []int{7}
}

func (x *CheckClaimResult) GetValue() *ClaimStatus {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_koinos_contracts_claim_claim_proto protoreflect.FileDescriptor

var file_koinos_contracts_claim_claim_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x73, 0x2f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x2f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x1a, 0x14, 0x6b, 0x6f,
	0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x65, 0x74, 0x68, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x45, 0x74, 0x68, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12,
	0x30, 0x0a, 0x14, 0x65, 0x74, 0x68, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x5f,
	0x63, 0x6c, 0x61, 0x69, 0x6d, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x65,
	0x74, 0x68, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x65,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x6b, 0x6f, 0x69, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4b, 0x6f, 0x69, 0x6e,
	0x12, 0x21, 0x0a, 0x0c, 0x6b, 0x6f, 0x69, 0x6e, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6b, 0x6f, 0x69, 0x6e, 0x43, 0x6c, 0x61, 0x69,
	0x6d, 0x65, 0x64, 0x22, 0x4f, 0x0a, 0x0c, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x0b, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c,
	0x61, 0x69, 0x6d, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x6c, 0x61,
	0x69, 0x6d, 0x65, 0x64, 0x22, 0x7f, 0x0a, 0x0f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x61, 0x72,
	0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x0b, 0x65, 0x74, 0x68, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x04, 0x80, 0xb5,
	0x18, 0x02, 0x52, 0x0a, 0x65, 0x74, 0x68, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x27,
	0x0a, 0x0c, 0x6b, 0x6f, 0x69, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x42, 0x04, 0x80, 0xb5, 0x18, 0x06, 0x52, 0x0b, 0x6b, 0x6f, 0x69, 0x6e,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x5f, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x4b, 0x0a, 0x0f, 0x67,
	0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x38,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73,
	0x2e, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x2e, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3e, 0x0a, 0x15, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x25, 0x0a, 0x0b, 0x65, 0x74, 0x68, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x04, 0x80, 0xb5, 0x18, 0x02, 0x52, 0x0a, 0x65, 0x74,
	0x68, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x50, 0x0a, 0x12, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x5f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x3a,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73,
	0x2e, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x2e, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f,
	0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x73, 0x2f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_koinos_contracts_claim_claim_proto_rawDescOnce sync.Once
	file_koinos_contracts_claim_claim_proto_rawDescData = file_koinos_contracts_claim_claim_proto_rawDesc
)

func file_koinos_contracts_claim_claim_proto_rawDescGZIP() []byte {
	file_koinos_contracts_claim_claim_proto_rawDescOnce.Do(func() {
		file_koinos_contracts_claim_claim_proto_rawDescData = protoimpl.X.CompressGZIP(file_koinos_contracts_claim_claim_proto_rawDescData)
	})
	return file_koinos_contracts_claim_claim_proto_rawDescData
}

var file_koinos_contracts_claim_claim_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_koinos_contracts_claim_claim_proto_goTypes = []interface{}{
	(*ClaimInfo)(nil),           // 0: koinos.contracts.claim.claim_info
	(*ClaimStatus)(nil),         // 1: koinos.contracts.claim.claim_status
	(*ClaimArguments)(nil),      // 2: koinos.contracts.claim.claim_arguments
	(*ClaimResult)(nil),         // 3: koinos.contracts.claim.claim_result
	(*GetInfoArguments)(nil),    // 4: koinos.contracts.claim.get_info_arguments
	(*GetInfoResult)(nil),       // 5: koinos.contracts.claim.get_info_result
	(*CheckClaimArguments)(nil), // 6: koinos.contracts.claim.check_claim_arguments
	(*CheckClaimResult)(nil),    // 7: koinos.contracts.claim.check_claim_result
}
var file_koinos_contracts_claim_claim_proto_depIdxs = []int32{
	0, // 0: koinos.contracts.claim.get_info_result.value:type_name -> koinos.contracts.claim.claim_info
	1, // 1: koinos.contracts.claim.check_claim_result.value:type_name -> koinos.contracts.claim.claim_status
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_koinos_contracts_claim_claim_proto_init() }
func file_koinos_contracts_claim_claim_proto_init() {
	if File_koinos_contracts_claim_claim_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_koinos_contracts_claim_claim_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimInfo); i {
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
		file_koinos_contracts_claim_claim_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimStatus); i {
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
		file_koinos_contracts_claim_claim_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimArguments); i {
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
		file_koinos_contracts_claim_claim_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClaimResult); i {
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
		file_koinos_contracts_claim_claim_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoArguments); i {
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
		file_koinos_contracts_claim_claim_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoResult); i {
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
		file_koinos_contracts_claim_claim_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckClaimArguments); i {
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
		file_koinos_contracts_claim_claim_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckClaimResult); i {
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
			RawDescriptor: file_koinos_contracts_claim_claim_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_koinos_contracts_claim_claim_proto_goTypes,
		DependencyIndexes: file_koinos_contracts_claim_claim_proto_depIdxs,
		MessageInfos:      file_koinos_contracts_claim_claim_proto_msgTypes,
	}.Build()
	File_koinos_contracts_claim_claim_proto = out.File
	file_koinos_contracts_claim_claim_proto_rawDesc = nil
	file_koinos_contracts_claim_claim_proto_goTypes = nil
	file_koinos_contracts_claim_claim_proto_depIdxs = nil
}
