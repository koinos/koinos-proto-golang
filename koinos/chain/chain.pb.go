// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: koinos/chain/chain.proto

package chain

import (
	koinos "github.com/koinos/koinos-proto-golang/koinos"
	_ "github.com/koinos/koinos-proto-golang/koinos/protocol"
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

type Privilege int32

const (
	Privilege_kernel_mode Privilege = 0
	Privilege_user_mode   Privilege = 1
)

// Enum value maps for Privilege.
var (
	Privilege_name = map[int32]string{
		0: "kernel_mode",
		1: "user_mode",
	}
	Privilege_value = map[string]int32{
		"kernel_mode": 0,
		"user_mode":   1,
	}
)

func (x Privilege) Enum() *Privilege {
	p := new(Privilege)
	*p = x
	return p
}

func (x Privilege) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Privilege) Descriptor() protoreflect.EnumDescriptor {
	return file_koinos_chain_chain_proto_enumTypes[0].Descriptor()
}

func (Privilege) Type() protoreflect.EnumType {
	return &file_koinos_chain_chain_proto_enumTypes[0]
}

func (x Privilege) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Privilege.Descriptor instead.
func (Privilege) EnumDescriptor() ([]byte, []int) {
	return file_koinos_chain_chain_proto_rawDescGZIP(), []int{0}
}

type ObjectSpace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	System bool   `protobuf:"varint,1,opt,name=system,proto3" json:"system,omitempty"`
	Zone   []byte `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty"`
	Id     uint32 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ObjectSpace) Reset() {
	*x = ObjectSpace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_chain_chain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ObjectSpace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObjectSpace) ProtoMessage() {}

func (x *ObjectSpace) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_chain_chain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObjectSpace.ProtoReflect.Descriptor instead.
func (*ObjectSpace) Descriptor() ([]byte, []int) {
	return file_koinos_chain_chain_proto_rawDescGZIP(), []int{0}
}

func (x *ObjectSpace) GetSystem() bool {
	if x != nil {
		return x.System
	}
	return false
}

func (x *ObjectSpace) GetZone() []byte {
	if x != nil {
		return x.Zone
	}
	return nil
}

func (x *ObjectSpace) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DatabaseKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Space *ObjectSpace `protobuf:"bytes,1,opt,name=space,proto3" json:"space,omitempty"`
	Key   []byte       `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DatabaseKey) Reset() {
	*x = DatabaseKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_chain_chain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatabaseKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatabaseKey) ProtoMessage() {}

func (x *DatabaseKey) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_chain_chain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatabaseKey.ProtoReflect.Descriptor instead.
func (*DatabaseKey) Descriptor() ([]byte, []int) {
	return file_koinos_chain_chain_proto_rawDescGZIP(), []int{1}
}

func (x *DatabaseKey) GetSpace() *ObjectSpace {
	if x != nil {
		return x.Space
	}
	return nil
}

func (x *DatabaseKey) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type MaxAccountResources struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MaxAccountResources) Reset() {
	*x = MaxAccountResources{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_chain_chain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaxAccountResources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaxAccountResources) ProtoMessage() {}

func (x *MaxAccountResources) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_chain_chain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaxAccountResources.ProtoReflect.Descriptor instead.
func (*MaxAccountResources) Descriptor() ([]byte, []int) {
	return file_koinos_chain_chain_proto_rawDescGZIP(), []int{2}
}

func (x *MaxAccountResources) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type HeadInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeadTopology          *koinos.BlockTopology `protobuf:"bytes,1,opt,name=head_topology,json=headTopology,proto3" json:"head_topology,omitempty"`
	HeadBlockTime         uint64                `protobuf:"varint,2,opt,name=head_block_time,json=headBlockTime,proto3" json:"head_block_time,omitempty"`
	LastIrreversibleBlock uint64                `protobuf:"varint,3,opt,name=last_irreversible_block,json=lastIrreversibleBlock,proto3" json:"last_irreversible_block,omitempty"`
}

func (x *HeadInfo) Reset() {
	*x = HeadInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_chain_chain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeadInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeadInfo) ProtoMessage() {}

func (x *HeadInfo) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_chain_chain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeadInfo.ProtoReflect.Descriptor instead.
func (*HeadInfo) Descriptor() ([]byte, []int) {
	return file_koinos_chain_chain_proto_rawDescGZIP(), []int{3}
}

func (x *HeadInfo) GetHeadTopology() *koinos.BlockTopology {
	if x != nil {
		return x.HeadTopology
	}
	return nil
}

func (x *HeadInfo) GetHeadBlockTime() uint64 {
	if x != nil {
		return x.HeadBlockTime
	}
	return 0
}

func (x *HeadInfo) GetLastIrreversibleBlock() uint64 {
	if x != nil {
		return x.LastIrreversibleBlock
	}
	return 0
}

type CallerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Caller          []byte    `protobuf:"bytes,1,opt,name=caller,proto3" json:"caller,omitempty"`
	CallerPrivilege Privilege `protobuf:"varint,2,opt,name=caller_privilege,json=callerPrivilege,proto3,enum=koinos.chain.Privilege" json:"caller_privilege,omitempty"`
}

func (x *CallerData) Reset() {
	*x = CallerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_chain_chain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallerData) ProtoMessage() {}

func (x *CallerData) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_chain_chain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallerData.ProtoReflect.Descriptor instead.
func (*CallerData) Descriptor() ([]byte, []int) {
	return file_koinos_chain_chain_proto_rawDescGZIP(), []int{4}
}

func (x *CallerData) GetCaller() []byte {
	if x != nil {
		return x.Caller
	}
	return nil
}

func (x *CallerData) GetCallerPrivilege() Privilege {
	if x != nil {
		return x.CallerPrivilege
	}
	return Privilege_kernel_mode
}

type ResourceLimitData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiskStorageLimit      uint64 `protobuf:"varint,1,opt,name=disk_storage_limit,json=diskStorageLimit,proto3" json:"disk_storage_limit,omitempty"`
	DiskStorageCost       uint64 `protobuf:"varint,2,opt,name=disk_storage_cost,json=diskStorageCost,proto3" json:"disk_storage_cost,omitempty"`
	NetworkBandwidthLimit uint64 `protobuf:"varint,3,opt,name=network_bandwidth_limit,json=networkBandwidthLimit,proto3" json:"network_bandwidth_limit,omitempty"`
	NetworkBandwidthCost  uint64 `protobuf:"varint,4,opt,name=network_bandwidth_cost,json=networkBandwidthCost,proto3" json:"network_bandwidth_cost,omitempty"`
	ComputeBandwidthLimit uint64 `protobuf:"varint,5,opt,name=compute_bandwidth_limit,json=computeBandwidthLimit,proto3" json:"compute_bandwidth_limit,omitempty"`
	ComputeBandwidthCost  uint64 `protobuf:"varint,6,opt,name=compute_bandwidth_cost,json=computeBandwidthCost,proto3" json:"compute_bandwidth_cost,omitempty"`
}

func (x *ResourceLimitData) Reset() {
	*x = ResourceLimitData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_chain_chain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceLimitData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceLimitData) ProtoMessage() {}

func (x *ResourceLimitData) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_chain_chain_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceLimitData.ProtoReflect.Descriptor instead.
func (*ResourceLimitData) Descriptor() ([]byte, []int) {
	return file_koinos_chain_chain_proto_rawDescGZIP(), []int{5}
}

func (x *ResourceLimitData) GetDiskStorageLimit() uint64 {
	if x != nil {
		return x.DiskStorageLimit
	}
	return 0
}

func (x *ResourceLimitData) GetDiskStorageCost() uint64 {
	if x != nil {
		return x.DiskStorageCost
	}
	return 0
}

func (x *ResourceLimitData) GetNetworkBandwidthLimit() uint64 {
	if x != nil {
		return x.NetworkBandwidthLimit
	}
	return 0
}

func (x *ResourceLimitData) GetNetworkBandwidthCost() uint64 {
	if x != nil {
		return x.NetworkBandwidthCost
	}
	return 0
}

func (x *ResourceLimitData) GetComputeBandwidthLimit() uint64 {
	if x != nil {
		return x.ComputeBandwidthLimit
	}
	return 0
}

func (x *ResourceLimitData) GetComputeBandwidthCost() uint64 {
	if x != nil {
		return x.ComputeBandwidthCost
	}
	return 0
}

type ContractMetadataObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash                     []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	System                   bool   `protobuf:"varint,2,opt,name=system,proto3" json:"system,omitempty"`
	AuthorizesCallContract   bool   `protobuf:"varint,3,opt,name=authorizes_call_contract,json=authorizesCallContract,proto3" json:"authorizes_call_contract,omitempty"`
	AuthorizesUseRc          bool   `protobuf:"varint,4,opt,name=authorizes_use_rc,json=authorizesUseRc,proto3" json:"authorizes_use_rc,omitempty"`
	AuthorizesUploadContract bool   `protobuf:"varint,5,opt,name=authorizes_upload_contract,json=authorizesUploadContract,proto3" json:"authorizes_upload_contract,omitempty"`
}

func (x *ContractMetadataObject) Reset() {
	*x = ContractMetadataObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_chain_chain_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractMetadataObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractMetadataObject) ProtoMessage() {}

func (x *ContractMetadataObject) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_chain_chain_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractMetadataObject.ProtoReflect.Descriptor instead.
func (*ContractMetadataObject) Descriptor() ([]byte, []int) {
	return file_koinos_chain_chain_proto_rawDescGZIP(), []int{6}
}

func (x *ContractMetadataObject) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *ContractMetadataObject) GetSystem() bool {
	if x != nil {
		return x.System
	}
	return false
}

func (x *ContractMetadataObject) GetAuthorizesCallContract() bool {
	if x != nil {
		return x.AuthorizesCallContract
	}
	return false
}

func (x *ContractMetadataObject) GetAuthorizesUseRc() bool {
	if x != nil {
		return x.AuthorizesUseRc
	}
	return false
}

func (x *ContractMetadataObject) GetAuthorizesUploadContract() bool {
	if x != nil {
		return x.AuthorizesUploadContract
	}
	return false
}

type ComputeBandwidthEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Compute uint64 `protobuf:"varint,2,opt,name=compute,proto3" json:"compute,omitempty"`
}

func (x *ComputeBandwidthEntry) Reset() {
	*x = ComputeBandwidthEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_chain_chain_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeBandwidthEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeBandwidthEntry) ProtoMessage() {}

func (x *ComputeBandwidthEntry) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_chain_chain_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeBandwidthEntry.ProtoReflect.Descriptor instead.
func (*ComputeBandwidthEntry) Descriptor() ([]byte, []int) {
	return file_koinos_chain_chain_proto_rawDescGZIP(), []int{7}
}

func (x *ComputeBandwidthEntry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ComputeBandwidthEntry) GetCompute() uint64 {
	if x != nil {
		return x.Compute
	}
	return 0
}

type ComputeBandwidthRegistry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*ComputeBandwidthEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *ComputeBandwidthRegistry) Reset() {
	*x = ComputeBandwidthRegistry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_chain_chain_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeBandwidthRegistry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeBandwidthRegistry) ProtoMessage() {}

func (x *ComputeBandwidthRegistry) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_chain_chain_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeBandwidthRegistry.ProtoReflect.Descriptor instead.
func (*ComputeBandwidthRegistry) Descriptor() ([]byte, []int) {
	return file_koinos_chain_chain_proto_rawDescGZIP(), []int{8}
}

func (x *ComputeBandwidthRegistry) GetEntries() []*ComputeBandwidthEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type GenesisEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Space *ObjectSpace `protobuf:"bytes,1,opt,name=space,proto3" json:"space,omitempty"`
	Key   []byte       `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte       `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GenesisEntry) Reset() {
	*x = GenesisEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_chain_chain_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisEntry) ProtoMessage() {}

func (x *GenesisEntry) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_chain_chain_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisEntry.ProtoReflect.Descriptor instead.
func (*GenesisEntry) Descriptor() ([]byte, []int) {
	return file_koinos_chain_chain_proto_rawDescGZIP(), []int{9}
}

func (x *GenesisEntry) GetSpace() *ObjectSpace {
	if x != nil {
		return x.Space
	}
	return nil
}

func (x *GenesisEntry) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *GenesisEntry) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type GenesisData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*GenesisEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *GenesisData) Reset() {
	*x = GenesisData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_chain_chain_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisData) ProtoMessage() {}

func (x *GenesisData) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_chain_chain_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenesisData.ProtoReflect.Descriptor instead.
func (*GenesisData) Descriptor() ([]byte, []int) {
	return file_koinos_chain_chain_proto_rawDescGZIP(), []int{10}
}

func (x *GenesisData) GetEntries() []*GenesisEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

var File_koinos_chain_chain_proto protoreflect.FileDescriptor

var file_koinos_chain_chain_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6b, 0x6f, 0x69, 0x6e,
	0x6f, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x1a, 0x14, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x4a, 0x0a, 0x0c, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x7a,
	0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x7a, 0x6f, 0x6e, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x52, 0x0a, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6b, 0x65, 0x79, 0x12,
	0x30, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x05, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x22, 0x31, 0x0a, 0x15, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xb0, 0x01, 0x0a, 0x09, 0x68, 0x65, 0x61, 0x64, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x12, 0x3b, 0x0a, 0x0d, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x6f, 0x70,
	0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6b, 0x6f,
	0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x6f, 0x70, 0x6f, 0x6c,
	0x6f, 0x67, 0x79, 0x52, 0x0c, 0x68, 0x65, 0x61, 0x64, 0x54, 0x6f, 0x70, 0x6f, 0x6c, 0x6f, 0x67,
	0x79, 0x12, 0x2a, 0x0a, 0x0f, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x0d,
	0x68, 0x65, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3a, 0x0a,
	0x17, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x72, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x62,
	0x6c, 0x65, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02,
	0x30, 0x01, 0x52, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x72, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x62, 0x6c, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x6f, 0x0a, 0x0b, 0x63, 0x61, 0x6c,
	0x6c, 0x65, 0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x06, 0x63, 0x61, 0x6c, 0x6c,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x04, 0x80, 0xb5, 0x18, 0x06, 0x52, 0x06,
	0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x10, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72,
	0x5f, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x17, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x52, 0x0f, 0x63, 0x61, 0x6c, 0x6c, 0x65,
	0x72, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x22, 0xe3, 0x02, 0x0a, 0x13, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x30, 0x0a, 0x12, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02,
	0x30, 0x01, 0x52, 0x10, 0x64, 0x69, 0x73, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x12, 0x2e, 0x0a, 0x11, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x02, 0x30, 0x01, 0x52, 0x0f, 0x64, 0x69, 0x73, 0x6b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x17, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f,
	0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x15, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x38, 0x0a, 0x16, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x62, 0x61, 0x6e, 0x64,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x02, 0x30, 0x01, 0x52, 0x14, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x61, 0x6e,
	0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x17, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52,
	0x15, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74,
	0x68, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x38, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x5f, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x63, 0x6f, 0x73, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x14, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x43, 0x6f, 0x73, 0x74,
	0x22, 0xea, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x38, 0x0a, 0x18, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x73, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x73, 0x43, 0x61, 0x6c, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x73, 0x5f, 0x75, 0x73, 0x65, 0x5f, 0x72, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x73, 0x55, 0x73, 0x65, 0x52, 0x63, 0x12,
	0x3c, 0x0a, 0x1a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x73, 0x5f, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x18, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x73, 0x55,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x22, 0x47, 0x0a,
	0x17, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x22, 0x5d, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x5f, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x12, 0x3f, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x62, 0x61, 0x6e,
	0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x69, 0x0a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73,
	0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2e, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x52, 0x05, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x45, 0x0a, 0x0c, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x5f, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x35, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07,
	0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2a, 0x2b, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x76, 0x69,
	0x6c, 0x65, 0x67, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x10, 0x01, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6b, 0x6f,
	0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_koinos_chain_chain_proto_rawDescOnce sync.Once
	file_koinos_chain_chain_proto_rawDescData = file_koinos_chain_chain_proto_rawDesc
)

func file_koinos_chain_chain_proto_rawDescGZIP() []byte {
	file_koinos_chain_chain_proto_rawDescOnce.Do(func() {
		file_koinos_chain_chain_proto_rawDescData = protoimpl.X.CompressGZIP(file_koinos_chain_chain_proto_rawDescData)
	})
	return file_koinos_chain_chain_proto_rawDescData
}

var file_koinos_chain_chain_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_koinos_chain_chain_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_koinos_chain_chain_proto_goTypes = []interface{}{
	(Privilege)(0),                   // 0: koinos.chain.privilege
	(*ObjectSpace)(nil),              // 1: koinos.chain.object_space
	(*DatabaseKey)(nil),              // 2: koinos.chain.database_key
	(*MaxAccountResources)(nil),      // 3: koinos.chain.max_account_resources
	(*HeadInfo)(nil),                 // 4: koinos.chain.head_info
	(*CallerData)(nil),               // 5: koinos.chain.caller_data
	(*ResourceLimitData)(nil),        // 6: koinos.chain.resource_limit_data
	(*ContractMetadataObject)(nil),   // 7: koinos.chain.contract_metadata_object
	(*ComputeBandwidthEntry)(nil),    // 8: koinos.chain.compute_bandwidth_entry
	(*ComputeBandwidthRegistry)(nil), // 9: koinos.chain.compute_bandwidth_registry
	(*GenesisEntry)(nil),             // 10: koinos.chain.genesis_entry
	(*GenesisData)(nil),              // 11: koinos.chain.genesis_data
	(*koinos.BlockTopology)(nil),     // 12: koinos.block_topology
}
var file_koinos_chain_chain_proto_depIdxs = []int32{
	1,  // 0: koinos.chain.database_key.space:type_name -> koinos.chain.object_space
	12, // 1: koinos.chain.head_info.head_topology:type_name -> koinos.block_topology
	0,  // 2: koinos.chain.caller_data.caller_privilege:type_name -> koinos.chain.privilege
	8,  // 3: koinos.chain.compute_bandwidth_registry.entries:type_name -> koinos.chain.compute_bandwidth_entry
	1,  // 4: koinos.chain.genesis_entry.space:type_name -> koinos.chain.object_space
	10, // 5: koinos.chain.genesis_data.entries:type_name -> koinos.chain.genesis_entry
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_koinos_chain_chain_proto_init() }
func file_koinos_chain_chain_proto_init() {
	if File_koinos_chain_chain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_koinos_chain_chain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ObjectSpace); i {
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
		file_koinos_chain_chain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DatabaseKey); i {
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
		file_koinos_chain_chain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaxAccountResources); i {
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
		file_koinos_chain_chain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeadInfo); i {
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
		file_koinos_chain_chain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallerData); i {
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
		file_koinos_chain_chain_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceLimitData); i {
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
		file_koinos_chain_chain_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractMetadataObject); i {
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
		file_koinos_chain_chain_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeBandwidthEntry); i {
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
		file_koinos_chain_chain_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeBandwidthRegistry); i {
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
		file_koinos_chain_chain_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisEntry); i {
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
		file_koinos_chain_chain_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisData); i {
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
			RawDescriptor: file_koinos_chain_chain_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_koinos_chain_chain_proto_goTypes,
		DependencyIndexes: file_koinos_chain_chain_proto_depIdxs,
		EnumInfos:         file_koinos_chain_chain_proto_enumTypes,
		MessageInfos:      file_koinos_chain_chain_proto_msgTypes,
	}.Build()
	File_koinos_chain_chain_proto = out.File
	file_koinos_chain_chain_proto_rawDesc = nil
	file_koinos_chain_chain_proto_goTypes = nil
	file_koinos_chain_chain_proto_depIdxs = nil
}
