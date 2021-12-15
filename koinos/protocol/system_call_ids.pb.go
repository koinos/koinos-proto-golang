// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: koinos/protocol/system_call_ids.proto

package protocol

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

type SystemCallId int32

const (
	SystemCallId_reserved_id SystemCallId = 0
	// General Blockchain Management
	SystemCallId_get_head_info                       SystemCallId = 1
	SystemCallId_apply_block                         SystemCallId = 2
	SystemCallId_apply_transaction                   SystemCallId = 3
	SystemCallId_apply_upload_contract_operation     SystemCallId = 4
	SystemCallId_apply_call_contract_operation       SystemCallId = 5
	SystemCallId_apply_set_system_call_operation     SystemCallId = 6
	SystemCallId_apply_set_system_contract_operation SystemCallId = 7
	// System Helpers
	SystemCallId_process_block_signature     SystemCallId = 101
	SystemCallId_get_transaction_payer       SystemCallId = 102
	SystemCallId_get_transaction_rc_limit    SystemCallId = 103
	SystemCallId_get_transaction_signature   SystemCallId = 104
	SystemCallId_get_last_irreversible_block SystemCallId = 105
	SystemCallId_get_account_nonce           SystemCallId = 106
	// Resource Subsystem
	SystemCallId_get_account_rc          SystemCallId = 201
	SystemCallId_consume_account_rc      SystemCallId = 202
	SystemCallId_get_resource_limits     SystemCallId = 203
	SystemCallId_consume_block_resources SystemCallId = 204
	// Database
	SystemCallId_put_object      SystemCallId = 301
	SystemCallId_remove_object   SystemCallId = 302
	SystemCallId_get_object      SystemCallId = 303
	SystemCallId_get_next_object SystemCallId = 304
	SystemCallId_get_prev_object SystemCallId = 305
	// Logging
	SystemCallId_prints SystemCallId = 401
	SystemCallId_event  SystemCallId = 402
	// Cryptography
	SystemCallId_hash               SystemCallId = 501
	SystemCallId_recover_public_key SystemCallId = 502
	SystemCallId_verify_merkle_root SystemCallId = 503
	// Contract Management
	SystemCallId_call_contract               SystemCallId = 601
	SystemCallId_get_entry_point             SystemCallId = 602
	SystemCallId_get_contract_arguments_size SystemCallId = 603
	SystemCallId_get_contract_arguments      SystemCallId = 604
	SystemCallId_set_contract_result         SystemCallId = 605
	SystemCallId_exit_contract               SystemCallId = 606
	SystemCallId_get_contract_id             SystemCallId = 607
	SystemCallId_get_caller                  SystemCallId = 608
	SystemCallId_require_authority           SystemCallId = 609
)

// Enum value maps for SystemCallId.
var (
	SystemCallId_name = map[int32]string{
		0:   "reserved_id",
		1:   "get_head_info",
		2:   "apply_block",
		3:   "apply_transaction",
		4:   "apply_upload_contract_operation",
		5:   "apply_call_contract_operation",
		6:   "apply_set_system_call_operation",
		7:   "apply_set_system_contract_operation",
		101: "process_block_signature",
		102: "get_transaction_payer",
		103: "get_transaction_rc_limit",
		104: "get_transaction_signature",
		105: "get_last_irreversible_block",
		106: "get_account_nonce",
		201: "get_account_rc",
		202: "consume_account_rc",
		203: "get_resource_limits",
		204: "consume_block_resources",
		301: "put_object",
		302: "remove_object",
		303: "get_object",
		304: "get_next_object",
		305: "get_prev_object",
		401: "prints",
		402: "event",
		501: "hash",
		502: "recover_public_key",
		503: "verify_merkle_root",
		601: "call_contract",
		602: "get_entry_point",
		603: "get_contract_arguments_size",
		604: "get_contract_arguments",
		605: "set_contract_result",
		606: "exit_contract",
		607: "get_contract_id",
		608: "get_caller",
		609: "require_authority",
	}
	SystemCallId_value = map[string]int32{
		"reserved_id":                         0,
		"get_head_info":                       1,
		"apply_block":                         2,
		"apply_transaction":                   3,
		"apply_upload_contract_operation":     4,
		"apply_call_contract_operation":       5,
		"apply_set_system_call_operation":     6,
		"apply_set_system_contract_operation": 7,
		"process_block_signature":             101,
		"get_transaction_payer":               102,
		"get_transaction_rc_limit":            103,
		"get_transaction_signature":           104,
		"get_last_irreversible_block":         105,
		"get_account_nonce":                   106,
		"get_account_rc":                      201,
		"consume_account_rc":                  202,
		"get_resource_limits":                 203,
		"consume_block_resources":             204,
		"put_object":                          301,
		"remove_object":                       302,
		"get_object":                          303,
		"get_next_object":                     304,
		"get_prev_object":                     305,
		"prints":                              401,
		"event":                               402,
		"hash":                                501,
		"recover_public_key":                  502,
		"verify_merkle_root":                  503,
		"call_contract":                       601,
		"get_entry_point":                     602,
		"get_contract_arguments_size":         603,
		"get_contract_arguments":              604,
		"set_contract_result":                 605,
		"exit_contract":                       606,
		"get_contract_id":                     607,
		"get_caller":                          608,
		"require_authority":                   609,
	}
)

func (x SystemCallId) Enum() *SystemCallId {
	p := new(SystemCallId)
	*p = x
	return p
}

func (x SystemCallId) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SystemCallId) Descriptor() protoreflect.EnumDescriptor {
	return file_koinos_protocol_system_call_ids_proto_enumTypes[0].Descriptor()
}

func (SystemCallId) Type() protoreflect.EnumType {
	return &file_koinos_protocol_system_call_ids_proto_enumTypes[0]
}

func (x SystemCallId) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SystemCallId.Descriptor instead.
func (SystemCallId) EnumDescriptor() ([]byte, []int) {
	return file_koinos_protocol_system_call_ids_proto_rawDescGZIP(), []int{0}
}

var File_koinos_protocol_system_call_ids_proto protoreflect.FileDescriptor

var file_koinos_protocol_system_call_ids_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x64,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2a, 0x93, 0x07, 0x0a, 0x0e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x12, 0x0f, 0x0a, 0x0b, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d,
	0x67, 0x65, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x10, 0x01, 0x12,
	0x0f, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x10, 0x02,
	0x12, 0x15, 0x0a, 0x11, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x03, 0x12, 0x23, 0x0a, 0x1f, 0x61, 0x70, 0x70, 0x6c, 0x79,
	0x5f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x04, 0x12, 0x21, 0x0a, 0x1d,
	0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x05, 0x12,
	0x23, 0x0a, 0x1f, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x5f, 0x63, 0x61, 0x6c, 0x6c, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x10, 0x06, 0x12, 0x27, 0x0a, 0x23, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x73, 0x65,
	0x74, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x07, 0x12, 0x1b, 0x0a,
	0x17, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x10, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x67, 0x65,
	0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61,
	0x79, 0x65, 0x72, 0x10, 0x66, 0x12, 0x1c, 0x0a, 0x18, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x63, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x10, 0x67, 0x12, 0x1d, 0x0a, 0x19, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x10, 0x68, 0x12, 0x1f, 0x0a, 0x1b, 0x67, 0x65, 0x74, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69,
	0x72, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x62, 0x6c, 0x65, 0x5f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x10, 0x69, 0x12, 0x15, 0x0a, 0x11, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x10, 0x6a, 0x12, 0x13, 0x0a, 0x0e, 0x67, 0x65,
	0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x63, 0x10, 0xc9, 0x01, 0x12,
	0x17, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x72, 0x63, 0x10, 0xca, 0x01, 0x12, 0x18, 0x0a, 0x13, 0x67, 0x65, 0x74, 0x5f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x10,
	0xcb, 0x01, 0x12, 0x1c, 0x0a, 0x17, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x5f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x10, 0xcc, 0x01,
	0x12, 0x0f, 0x0a, 0x0a, 0x70, 0x75, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x10, 0xad,
	0x02, 0x12, 0x12, 0x0a, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x10, 0xae, 0x02, 0x12, 0x0f, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x5f, 0x6f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x10, 0xaf, 0x02, 0x12, 0x14, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x5f, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x10, 0xb0, 0x02, 0x12, 0x14, 0x0a, 0x0f,
	0x67, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x10,
	0xb1, 0x02, 0x12, 0x0b, 0x0a, 0x06, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x73, 0x10, 0x91, 0x03, 0x12,
	0x0a, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x10, 0x92, 0x03, 0x12, 0x09, 0x0a, 0x04, 0x68,
	0x61, 0x73, 0x68, 0x10, 0xf5, 0x03, 0x12, 0x17, 0x0a, 0x12, 0x72, 0x65, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x10, 0xf6, 0x03, 0x12,
	0x17, 0x0a, 0x12, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65,
	0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x10, 0xf7, 0x03, 0x12, 0x12, 0x0a, 0x0d, 0x63, 0x61, 0x6c, 0x6c,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x10, 0xd9, 0x04, 0x12, 0x14, 0x0a, 0x0f,
	0x67, 0x65, 0x74, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x10,
	0xda, 0x04, 0x12, 0x20, 0x0a, 0x1b, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x5f, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x10, 0xdb, 0x04, 0x12, 0x1b, 0x0a, 0x16, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x10, 0xdc,
	0x04, 0x12, 0x18, 0x0a, 0x13, 0x73, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x10, 0xdd, 0x04, 0x12, 0x12, 0x0a, 0x0d, 0x65,
	0x78, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x10, 0xde, 0x04, 0x12,
	0x14, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x10, 0xdf, 0x04, 0x12, 0x0f, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x61, 0x6c,
	0x6c, 0x65, 0x72, 0x10, 0xe0, 0x04, 0x12, 0x16, 0x0a, 0x11, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x10, 0xe1, 0x04, 0x42, 0x37,
	0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x69,
	0x6e, 0x6f, 0x73, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_koinos_protocol_system_call_ids_proto_rawDescOnce sync.Once
	file_koinos_protocol_system_call_ids_proto_rawDescData = file_koinos_protocol_system_call_ids_proto_rawDesc
)

func file_koinos_protocol_system_call_ids_proto_rawDescGZIP() []byte {
	file_koinos_protocol_system_call_ids_proto_rawDescOnce.Do(func() {
		file_koinos_protocol_system_call_ids_proto_rawDescData = protoimpl.X.CompressGZIP(file_koinos_protocol_system_call_ids_proto_rawDescData)
	})
	return file_koinos_protocol_system_call_ids_proto_rawDescData
}

var file_koinos_protocol_system_call_ids_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_koinos_protocol_system_call_ids_proto_goTypes = []interface{}{
	(SystemCallId)(0), // 0: koinos.protocol.system_call_id
}
var file_koinos_protocol_system_call_ids_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_koinos_protocol_system_call_ids_proto_init() }
func file_koinos_protocol_system_call_ids_proto_init() {
	if File_koinos_protocol_system_call_ids_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_koinos_protocol_system_call_ids_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_koinos_protocol_system_call_ids_proto_goTypes,
		DependencyIndexes: file_koinos_protocol_system_call_ids_proto_depIdxs,
		EnumInfos:         file_koinos_protocol_system_call_ids_proto_enumTypes,
	}.Build()
	File_koinos_protocol_system_call_ids_proto = out.File
	file_koinos_protocol_system_call_ids_proto_rawDesc = nil
	file_koinos_protocol_system_call_ids_proto_goTypes = nil
	file_koinos_protocol_system_call_ids_proto_depIdxs = nil
}
