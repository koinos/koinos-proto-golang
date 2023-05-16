// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.17.3
// source: koinos/contracts/pow/pow.proto

package pow

import (
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

type DifficultyMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target              []byte `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	LastBlockTime       uint64 `protobuf:"varint,2,opt,name=last_block_time,json=lastBlockTime,proto3" json:"last_block_time,omitempty"`
	Difficulty          []byte `protobuf:"bytes,3,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	TargetBlockInterval uint64 `protobuf:"varint,4,opt,name=target_block_interval,json=targetBlockInterval,proto3" json:"target_block_interval,omitempty"`
}

func (x *DifficultyMetadata) Reset() {
	*x = DifficultyMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_pow_pow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DifficultyMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DifficultyMetadata) ProtoMessage() {}

func (x *DifficultyMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_pow_pow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DifficultyMetadata.ProtoReflect.Descriptor instead.
func (*DifficultyMetadata) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_pow_pow_proto_rawDescGZIP(), []int{0}
}

func (x *DifficultyMetadata) GetTarget() []byte {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *DifficultyMetadata) GetLastBlockTime() uint64 {
	if x != nil {
		return x.LastBlockTime
	}
	return 0
}

func (x *DifficultyMetadata) GetDifficulty() []byte {
	if x != nil {
		return x.Difficulty
	}
	return nil
}

func (x *DifficultyMetadata) GetTargetBlockInterval() uint64 {
	if x != nil {
		return x.TargetBlockInterval
	}
	return 0
}

type GetDifficultyMetadataArguments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDifficultyMetadataArguments) Reset() {
	*x = GetDifficultyMetadataArguments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_pow_pow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDifficultyMetadataArguments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDifficultyMetadataArguments) ProtoMessage() {}

func (x *GetDifficultyMetadataArguments) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_pow_pow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDifficultyMetadataArguments.ProtoReflect.Descriptor instead.
func (*GetDifficultyMetadataArguments) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_pow_pow_proto_rawDescGZIP(), []int{1}
}

type GetDifficultyMetadataResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *DifficultyMetadata `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetDifficultyMetadataResult) Reset() {
	*x = GetDifficultyMetadataResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_pow_pow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDifficultyMetadataResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDifficultyMetadataResult) ProtoMessage() {}

func (x *GetDifficultyMetadataResult) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_pow_pow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDifficultyMetadataResult.ProtoReflect.Descriptor instead.
func (*GetDifficultyMetadataResult) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_pow_pow_proto_rawDescGZIP(), []int{2}
}

func (x *GetDifficultyMetadataResult) GetValue() *DifficultyMetadata {
	if x != nil {
		return x.Value
	}
	return nil
}

type PowSignatureData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nonce                []byte `protobuf:"bytes,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	RecoverableSignature []byte `protobuf:"bytes,2,opt,name=recoverable_signature,json=recoverableSignature,proto3" json:"recoverable_signature,omitempty"`
}

func (x *PowSignatureData) Reset() {
	*x = PowSignatureData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_contracts_pow_pow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PowSignatureData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PowSignatureData) ProtoMessage() {}

func (x *PowSignatureData) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_contracts_pow_pow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PowSignatureData.ProtoReflect.Descriptor instead.
func (*PowSignatureData) Descriptor() ([]byte, []int) {
	return file_koinos_contracts_pow_pow_proto_rawDescGZIP(), []int{3}
}

func (x *PowSignatureData) GetNonce() []byte {
	if x != nil {
		return x.Nonce
	}
	return nil
}

func (x *PowSignatureData) GetRecoverableSignature() []byte {
	if x != nil {
		return x.RecoverableSignature
	}
	return nil
}

var File_koinos_contracts_pow_pow_proto protoreflect.FileDescriptor

var file_koinos_contracts_pow_pow_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x73, 0x2f, 0x70, 0x6f, 0x77, 0x2f, 0x70, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x73, 0x2e, 0x70, 0x6f, 0x77, 0x22, 0xb1, 0x01, 0x0a, 0x13, 0x64, 0x69, 0x66, 0x66, 0x69,
	0x63, 0x75, 0x6c, 0x74, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x2a, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x02, 0x30, 0x01, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c,
	0x74, 0x79, 0x12, 0x36, 0x0a, 0x15, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x13, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x22, 0x23, 0x0a, 0x21, 0x67, 0x65,
	0x74, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x61, 0x0a, 0x1e, 0x67, 0x65, 0x74, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74,
	0x79, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x3f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x73, 0x2e, 0x70, 0x6f, 0x77, 0x2e, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c,
	0x74, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x5f, 0x0a, 0x12, 0x70, 0x6f, 0x77, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x33,
	0x0a, 0x15, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x14, 0x72,
	0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6b, 0x6f, 0x69,
	0x6e, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x70, 0x6f,
	0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_koinos_contracts_pow_pow_proto_rawDescOnce sync.Once
	file_koinos_contracts_pow_pow_proto_rawDescData = file_koinos_contracts_pow_pow_proto_rawDesc
)

func file_koinos_contracts_pow_pow_proto_rawDescGZIP() []byte {
	file_koinos_contracts_pow_pow_proto_rawDescOnce.Do(func() {
		file_koinos_contracts_pow_pow_proto_rawDescData = protoimpl.X.CompressGZIP(file_koinos_contracts_pow_pow_proto_rawDescData)
	})
	return file_koinos_contracts_pow_pow_proto_rawDescData
}

var file_koinos_contracts_pow_pow_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_koinos_contracts_pow_pow_proto_goTypes = []interface{}{
	(*DifficultyMetadata)(nil),             // 0: koinos.contracts.pow.difficulty_metadata
	(*GetDifficultyMetadataArguments)(nil), // 1: koinos.contracts.pow.get_difficulty_metadata_arguments
	(*GetDifficultyMetadataResult)(nil),    // 2: koinos.contracts.pow.get_difficulty_metadata_result
	(*PowSignatureData)(nil),               // 3: koinos.contracts.pow.pow_signature_data
}
var file_koinos_contracts_pow_pow_proto_depIdxs = []int32{
	0, // 0: koinos.contracts.pow.get_difficulty_metadata_result.value:type_name -> koinos.contracts.pow.difficulty_metadata
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_koinos_contracts_pow_pow_proto_init() }
func file_koinos_contracts_pow_pow_proto_init() {
	if File_koinos_contracts_pow_pow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_koinos_contracts_pow_pow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DifficultyMetadata); i {
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
		file_koinos_contracts_pow_pow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDifficultyMetadataArguments); i {
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
		file_koinos_contracts_pow_pow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDifficultyMetadataResult); i {
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
		file_koinos_contracts_pow_pow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PowSignatureData); i {
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
			RawDescriptor: file_koinos_contracts_pow_pow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_koinos_contracts_pow_pow_proto_goTypes,
		DependencyIndexes: file_koinos_contracts_pow_pow_proto_depIdxs,
		MessageInfos:      file_koinos_contracts_pow_pow_proto_msgTypes,
	}.Build()
	File_koinos_contracts_pow_pow_proto = out.File
	file_koinos_contracts_pow_pow_proto_rawDesc = nil
	file_koinos_contracts_pow_pow_proto_goTypes = nil
	file_koinos_contracts_pow_pow_proto_depIdxs = nil
}
