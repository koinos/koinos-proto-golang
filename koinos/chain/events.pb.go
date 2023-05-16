// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.17.3
// source: koinos/chain/events.proto

package chain

import (
	_ "github.com/koinos/koinos-proto-golang/koinos"
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

type SetSystemCallEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallId uint32                     `protobuf:"varint,1,opt,name=call_id,json=callId,proto3" json:"call_id,omitempty"`
	Target *protocol.SystemCallTarget `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *SetSystemCallEvent) Reset() {
	*x = SetSystemCallEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_chain_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetSystemCallEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSystemCallEvent) ProtoMessage() {}

func (x *SetSystemCallEvent) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_chain_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSystemCallEvent.ProtoReflect.Descriptor instead.
func (*SetSystemCallEvent) Descriptor() ([]byte, []int) {
	return file_koinos_chain_events_proto_rawDescGZIP(), []int{0}
}

func (x *SetSystemCallEvent) GetCallId() uint32 {
	if x != nil {
		return x.CallId
	}
	return 0
}

func (x *SetSystemCallEvent) GetTarget() *protocol.SystemCallTarget {
	if x != nil {
		return x.Target
	}
	return nil
}

type SetSystemContractEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractId     []byte `protobuf:"bytes,1,opt,name=contract_id,json=contractId,proto3" json:"contract_id,omitempty"`
	SystemContract bool   `protobuf:"varint,2,opt,name=system_contract,json=systemContract,proto3" json:"system_contract,omitempty"`
}

func (x *SetSystemContractEvent) Reset() {
	*x = SetSystemContractEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_chain_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetSystemContractEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSystemContractEvent) ProtoMessage() {}

func (x *SetSystemContractEvent) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_chain_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSystemContractEvent.ProtoReflect.Descriptor instead.
func (*SetSystemContractEvent) Descriptor() ([]byte, []int) {
	return file_koinos_chain_events_proto_rawDescGZIP(), []int{1}
}

func (x *SetSystemContractEvent) GetContractId() []byte {
	if x != nil {
		return x.ContractId
	}
	return nil
}

func (x *SetSystemContractEvent) GetSystemContract() bool {
	if x != nil {
		return x.SystemContract
	}
	return false
}

var File_koinos_chain_events_proto protoreflect.FileDescriptor

var file_koinos_chain_events_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6b, 0x6f, 0x69,
	0x6e, 0x6f, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x1a, 0x14, 0x6b, 0x6f, 0x69, 0x6e, 0x6f,
	0x73, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6d, 0x0a, 0x15, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x61,
	0x6c, 0x6c, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x6c, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x61, 0x6c, 0x6c, 0x49,
	0x64, 0x12, 0x3b, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x6b,
	0x0a, 0x19, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0b, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x42, 0x04, 0x80, 0xb5, 0x18, 0x05, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x42, 0x34, 0x5a, 0x32, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73,
	0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_koinos_chain_events_proto_rawDescOnce sync.Once
	file_koinos_chain_events_proto_rawDescData = file_koinos_chain_events_proto_rawDesc
)

func file_koinos_chain_events_proto_rawDescGZIP() []byte {
	file_koinos_chain_events_proto_rawDescOnce.Do(func() {
		file_koinos_chain_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_koinos_chain_events_proto_rawDescData)
	})
	return file_koinos_chain_events_proto_rawDescData
}

var file_koinos_chain_events_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_koinos_chain_events_proto_goTypes = []interface{}{
	(*SetSystemCallEvent)(nil),        // 0: koinos.chain.set_system_call_event
	(*SetSystemContractEvent)(nil),    // 1: koinos.chain.set_system_contract_event
	(*protocol.SystemCallTarget)(nil), // 2: koinos.protocol.system_call_target
}
var file_koinos_chain_events_proto_depIdxs = []int32{
	2, // 0: koinos.chain.set_system_call_event.target:type_name -> koinos.protocol.system_call_target
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_koinos_chain_events_proto_init() }
func file_koinos_chain_events_proto_init() {
	if File_koinos_chain_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_koinos_chain_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetSystemCallEvent); i {
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
		file_koinos_chain_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetSystemContractEvent); i {
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
			RawDescriptor: file_koinos_chain_events_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_koinos_chain_events_proto_goTypes,
		DependencyIndexes: file_koinos_chain_events_proto_depIdxs,
		MessageInfos:      file_koinos_chain_events_proto_msgTypes,
	}.Build()
	File_koinos_chain_events_proto = out.File
	file_koinos_chain_events_proto_rawDesc = nil
	file_koinos_chain_events_proto_goTypes = nil
	file_koinos_chain_events_proto_depIdxs = nil
}
