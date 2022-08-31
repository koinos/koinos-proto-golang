// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: koinos/contracts/vhp/vhp.proto

package vhp

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

type EffectiveBalanceOfArguments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner []byte `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (x *EffectiveBalanceOfArguments) Reset() {
	*x = EffectiveBalanceOfArguments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_vhp_vhp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EffectiveBalanceOfArguments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EffectiveBalanceOfArguments) ProtoMessage() {}

func (x *EffectiveBalanceOfArguments) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_vhp_vhp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EffectiveBalanceOfArguments.ProtoReflect.Descriptor instead.
func (*EffectiveBalanceOfArguments) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_vhp_vhp_proto_rawDescGZIP(), []int{0}
}

func (x *EffectiveBalanceOfArguments) GetOwner() []byte {
	if x != nil {
		return x.Owner
	}
	return nil
}

type EffectiveBalanceOfResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *EffectiveBalanceOfResult) Reset() {
	*x = EffectiveBalanceOfResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_vhp_vhp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EffectiveBalanceOfResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EffectiveBalanceOfResult) ProtoMessage() {}

func (x *EffectiveBalanceOfResult) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_vhp_vhp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EffectiveBalanceOfResult.ProtoReflect.Descriptor instead.
func (*EffectiveBalanceOfResult) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_vhp_vhp_proto_rawDescGZIP(), []int{1}
}

func (x *EffectiveBalanceOfResult) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type BalanceEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockHeight uint64 `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	Balance     uint64 `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *BalanceEntry) Reset() {
	*x = BalanceEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_vhp_vhp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BalanceEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BalanceEntry) ProtoMessage() {}

func (x *BalanceEntry) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_vhp_vhp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BalanceEntry.ProtoReflect.Descriptor instead.
func (*BalanceEntry) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_vhp_vhp_proto_rawDescGZIP(), []int{2}
}

func (x *BalanceEntry) GetBlockHeight() uint64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

func (x *BalanceEntry) GetBalance() uint64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type EffectiveBalanceObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentBalance uint64          `protobuf:"varint,1,opt,name=current_balance,json=currentBalance,proto3" json:"current_balance,omitempty"`
	PastBalances   []*BalanceEntry `protobuf:"bytes,2,rep,name=past_balances,json=pastBalances,proto3" json:"past_balances,omitempty"`
}

func (x *EffectiveBalanceObject) Reset() {
	*x = EffectiveBalanceObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_vhp_vhp_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EffectiveBalanceObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EffectiveBalanceObject) ProtoMessage() {}

func (x *EffectiveBalanceObject) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_vhp_vhp_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EffectiveBalanceObject.ProtoReflect.Descriptor instead.
func (*EffectiveBalanceObject) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_vhp_vhp_proto_rawDescGZIP(), []int{3}
}

func (x *EffectiveBalanceObject) GetCurrentBalance() uint64 {
	if x != nil {
		return x.CurrentBalance
	}
	return 0
}

func (x *EffectiveBalanceObject) GetPastBalances() []*BalanceEntry {
	if x != nil {
		return x.PastBalances
	}
	return nil
}

var File_koinos_contracts_vhp_vhp_proto protoreflect.FileDescriptor

var file_koinos_contracts_vhp_vhp_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x73, 0x2f, 0x76, 0x68, 0x70, 0x2f, 0x76, 0x68, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x73, 0x2e, 0x76, 0x68, 0x70, 0x1a, 0x14, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3c, 0x0a, 0x1e,
	0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1a,
	0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x04, 0x80,
	0xb5, 0x18, 0x06, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x37, 0x0a, 0x1b, 0x65, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x6f, 0x66, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x54, 0x0a, 0x0d, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x25, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x0b,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1c, 0x0a, 0x07, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01,
	0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x18, 0x65, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x2b, 0x0a, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x02, 0x30, 0x01, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x70, 0x61, 0x73, 0x74, 0x5f, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x6f, 0x69,
	0x6e, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x68,
	0x70, 0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0c, 0x70, 0x61, 0x73, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x42, 0x3c, 0x5a,
	0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x69, 0x6e,
	0x6f, 0x73, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x68, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_koinos_contracts_vhp_vhp_proto_rawDescOnce sync.Once
	file_koinos_contracts_vhp_vhp_proto_rawDescData = file_koinos_contracts_vhp_vhp_proto_rawDesc
)

func file_koinos_contracts_vhp_vhp_proto_rawDescGZIP() []byte {
	file_koinos_contracts_vhp_vhp_proto_rawDescOnce.Do(func() {
		file_koinos_contracts_vhp_vhp_proto_rawDescData = protoimpl.X.CompressGZIP(file_koinos_contracts_vhp_vhp_proto_rawDescData)
	})
	return file_koinos_contracts_vhp_vhp_proto_rawDescData
}

var file_koinos_contracts_vhp_vhp_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_koinos_contracts_vhp_vhp_proto_goTypes = []interface{}{
	(*EffectiveBalanceOfArguments)(nil), // 0: koinos.contracts.vhp.effective_balance_of_arguments
	(*EffectiveBalanceOfResult)(nil),    // 1: koinos.contracts.vhp.effective_balance_of_result
	(*BalanceEntry)(nil),                // 2: koinos.contracts.vhp.balance_entry
	(*EffectiveBalanceObject)(nil),      // 3: koinos.contracts.vhp.effective_balance_object
}
var file_koinos_contracts_vhp_vhp_proto_depIdxs = []int32{
	2, // 0: koinos.contracts.vhp.effective_balance_object.past_balances:type_name -> koinos.contracts.vhp.balance_entry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_koinos_contracts_vhp_vhp_proto_init() }
func file_koinos_contracts_vhp_vhp_proto_init() {
	if File_koinos_contracts_vhp_vhp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_koinos_contracts_vhp_vhp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EffectiveBalanceOfArguments); i {
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
		file_koinos_contracts_vhp_vhp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EffectiveBalanceOfResult); i {
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
		file_koinos_contracts_vhp_vhp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BalanceEntry); i {
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
		file_koinos_contracts_vhp_vhp_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EffectiveBalanceObject); i {
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
			RawDescriptor: file_koinos_contracts_vhp_vhp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_koinos_contracts_vhp_vhp_proto_goTypes,
		DependencyIndexes: file_koinos_contracts_vhp_vhp_proto_depIdxs,
		MessageInfos:      file_koinos_contracts_vhp_vhp_proto_msgTypes,
	}.Build()
	File_koinos_contracts_vhp_vhp_proto = out.File
	file_koinos_contracts_vhp_vhp_proto_rawDesc = nil
	file_koinos_contracts_vhp_vhp_proto_goTypes = nil
	file_koinos_contracts_vhp_vhp_proto_depIdxs = nil
}
