// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.17.3
// source: koinos/chain/error.proto

package chain

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

type ErrorCode int32

const (
	ErrorCode_success ErrorCode = 0
	// Reversions
	ErrorCode_reversion                    ErrorCode = 1
	ErrorCode_internal_error               ErrorCode = 100
	ErrorCode_system_authorization_failure ErrorCode = 101
	ErrorCode_invalid_contract             ErrorCode = 102
	ErrorCode_insufficient_privileges      ErrorCode = 103
	ErrorCode_insufficient_rc              ErrorCode = 104
	ErrorCode_insufficient_return_buffer   ErrorCode = 105
	ErrorCode_unknown_thunk                ErrorCode = 106
	ErrorCode_unknown_operation            ErrorCode = 107
	ErrorCode_read_only_context            ErrorCode = 108
	// Failures
	ErrorCode_failure                ErrorCode = -1
	ErrorCode_field_not_found        ErrorCode = -100
	ErrorCode_unknown_hash_code      ErrorCode = -101
	ErrorCode_unknown_dsa            ErrorCode = -102
	ErrorCode_unknown_system_call    ErrorCode = -103
	ErrorCode_operation_not_found    ErrorCode = -104
	ErrorCode_authorization_failure  ErrorCode = -200
	ErrorCode_invalid_nonce          ErrorCode = -201
	ErrorCode_invalid_signature      ErrorCode = -202
	ErrorCode_malformed_block        ErrorCode = -203
	ErrorCode_malformed_transaction  ErrorCode = -204
	ErrorCode_block_resource_failure ErrorCode = -205
	// Framework failures
	ErrorCode_unknown_backend                  ErrorCode = -1000
	ErrorCode_unexpected_state                 ErrorCode = -1001
	ErrorCode_missing_required_arguments       ErrorCode = -1002
	ErrorCode_unknown_previous_block           ErrorCode = -1003
	ErrorCode_unexpected_height                ErrorCode = -1004
	ErrorCode_block_state_error                ErrorCode = -1005
	ErrorCode_state_merkle_mismatch            ErrorCode = -1006
	ErrorCode_unexpected_receipt               ErrorCode = -1007
	ErrorCode_rpc_failure                      ErrorCode = -1008
	ErrorCode_pending_state_error              ErrorCode = -1009
	ErrorCode_timestamp_out_of_bounds          ErrorCode = -1010
	ErrorCode_indexer_failure                  ErrorCode = -1011
	ErrorCode_network_bandwidth_limit_exceeded ErrorCode = -1012
	ErrorCode_compute_bandwidth_limit_exceeded ErrorCode = -1013
	ErrorCode_disk_storage_limit_exceeded      ErrorCode = -1014
	ErrorCode_pre_irreversibility_block        ErrorCode = -1015
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0:     "success",
		1:     "reversion",
		100:   "internal_error",
		101:   "system_authorization_failure",
		102:   "invalid_contract",
		103:   "insufficient_privileges",
		104:   "insufficient_rc",
		105:   "insufficient_return_buffer",
		106:   "unknown_thunk",
		107:   "unknown_operation",
		108:   "read_only_context",
		-1:    "failure",
		-100:  "field_not_found",
		-101:  "unknown_hash_code",
		-102:  "unknown_dsa",
		-103:  "unknown_system_call",
		-104:  "operation_not_found",
		-200:  "authorization_failure",
		-201:  "invalid_nonce",
		-202:  "invalid_signature",
		-203:  "malformed_block",
		-204:  "malformed_transaction",
		-205:  "block_resource_failure",
		-1000: "unknown_backend",
		-1001: "unexpected_state",
		-1002: "missing_required_arguments",
		-1003: "unknown_previous_block",
		-1004: "unexpected_height",
		-1005: "block_state_error",
		-1006: "state_merkle_mismatch",
		-1007: "unexpected_receipt",
		-1008: "rpc_failure",
		-1009: "pending_state_error",
		-1010: "timestamp_out_of_bounds",
		-1011: "indexer_failure",
		-1012: "network_bandwidth_limit_exceeded",
		-1013: "compute_bandwidth_limit_exceeded",
		-1014: "disk_storage_limit_exceeded",
		-1015: "pre_irreversibility_block",
	}
	ErrorCode_value = map[string]int32{
		"success":                          0,
		"reversion":                        1,
		"internal_error":                   100,
		"system_authorization_failure":     101,
		"invalid_contract":                 102,
		"insufficient_privileges":          103,
		"insufficient_rc":                  104,
		"insufficient_return_buffer":       105,
		"unknown_thunk":                    106,
		"unknown_operation":                107,
		"read_only_context":                108,
		"failure":                          -1,
		"field_not_found":                  -100,
		"unknown_hash_code":                -101,
		"unknown_dsa":                      -102,
		"unknown_system_call":              -103,
		"operation_not_found":              -104,
		"authorization_failure":            -200,
		"invalid_nonce":                    -201,
		"invalid_signature":                -202,
		"malformed_block":                  -203,
		"malformed_transaction":            -204,
		"block_resource_failure":           -205,
		"unknown_backend":                  -1000,
		"unexpected_state":                 -1001,
		"missing_required_arguments":       -1002,
		"unknown_previous_block":           -1003,
		"unexpected_height":                -1004,
		"block_state_error":                -1005,
		"state_merkle_mismatch":            -1006,
		"unexpected_receipt":               -1007,
		"rpc_failure":                      -1008,
		"pending_state_error":              -1009,
		"timestamp_out_of_bounds":          -1010,
		"indexer_failure":                  -1011,
		"network_bandwidth_limit_exceeded": -1012,
		"compute_bandwidth_limit_exceeded": -1013,
		"disk_storage_limit_exceeded":      -1014,
		"pre_irreversibility_block":        -1015,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_koinos_chain_error_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_koinos_chain_error_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_koinos_chain_error_proto_rawDescGZIP(), []int{0}
}

type ErrorDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Logs []string `protobuf:"bytes,2,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *ErrorDetails) Reset() {
	*x = ErrorDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_koinos_chain_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorDetails) ProtoMessage() {}

func (x *ErrorDetails) ProtoReflect() protoreflect.Message {
	mi := &file_koinos_chain_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorDetails.ProtoReflect.Descriptor instead.
func (*ErrorDetails) Descriptor() ([]byte, []int) {
	return file_koinos_chain_error_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorDetails) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ErrorDetails) GetLogs() []string {
	if x != nil {
		return x.Logs
	}
	return nil
}

var File_koinos_chain_error_proto protoreflect.FileDescriptor

var file_koinos_chain_error_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6b, 0x6f, 0x69, 0x6e,
	0x6f, 0x73, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x22, 0x37, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67,
	0x73, 0x2a, 0xc0, 0x09, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00, 0x12, 0x0d, 0x0a,
	0x09, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x64,
	0x12, 0x20, 0x0a, 0x1c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x10, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x10, 0x66, 0x12, 0x1b, 0x0a, 0x17, 0x69, 0x6e, 0x73, 0x75,
	0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65,
	0x67, 0x65, 0x73, 0x10, 0x67, 0x12, 0x13, 0x0a, 0x0f, 0x69, 0x6e, 0x73, 0x75, 0x66, 0x66, 0x69,
	0x63, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x63, 0x10, 0x68, 0x12, 0x1e, 0x0a, 0x1a, 0x69, 0x6e,
	0x73, 0x75, 0x66, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x10, 0x69, 0x12, 0x11, 0x0a, 0x0d, 0x75, 0x6e,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x74, 0x68, 0x75, 0x6e, 0x6b, 0x10, 0x6a, 0x12, 0x15, 0x0a,
	0x11, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x10, 0x6b, 0x12, 0x15, 0x0a, 0x11, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e, 0x6c,
	0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x10, 0x6c, 0x12, 0x14, 0x0a, 0x07, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0x01, 0x12, 0x1c, 0x0a, 0x0f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x10, 0x9c, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12,
	0x1e, 0x0a, 0x11, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x10, 0x9b, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12,
	0x18, 0x0a, 0x0b, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x64, 0x73, 0x61, 0x10, 0x9a,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x20, 0x0a, 0x13, 0x75, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x61, 0x6c, 0x6c,
	0x10, 0x99, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x20, 0x0a, 0x13, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x10, 0x98, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x22, 0x0a,
	0x15, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x10, 0xb8, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0x01, 0x12, 0x1a, 0x0a, 0x0d, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x10, 0xb7, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x1e, 0x0a,
	0x11, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x10, 0xb6, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x1c, 0x0a,
	0x0f, 0x6d, 0x61, 0x6c, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x10, 0xb5, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x22, 0x0a, 0x15, 0x6d,
	0x61, 0x6c, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x10, 0xb4, 0xfe, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12,
	0x23, 0x0a, 0x16, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x10, 0xb3, 0xfe, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0x01, 0x12, 0x1c, 0x0a, 0x0f, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x10, 0x98, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0x01, 0x12, 0x1d, 0x0a, 0x10, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x10, 0x97, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0x01, 0x12, 0x27, 0x0a, 0x1a, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x10,
	0x96, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x23, 0x0a, 0x16, 0x75, 0x6e,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x10, 0x95, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12,
	0x1e, 0x0a, 0x11, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x68, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x10, 0x94, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12,
	0x1e, 0x0a, 0x11, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x10, 0x93, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12,
	0x22, 0x0a, 0x15, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f,
	0x6d, 0x69, 0x73, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x10, 0x92, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0x01, 0x12, 0x1f, 0x0a, 0x12, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x10, 0x91, 0xf8, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0x01, 0x12, 0x18, 0x0a, 0x0b, 0x72, 0x70, 0x63, 0x5f, 0x66, 0x61, 0x69, 0x6c,
	0x75, 0x72, 0x65, 0x10, 0x90, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x20,
	0x0a, 0x13, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x8f, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01,
	0x12, 0x24, 0x0a, 0x17, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6f, 0x75,
	0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x73, 0x10, 0x8e, 0xf8, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x1c, 0x0a, 0x0f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x65,
	0x72, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x10, 0x8d, 0xf8, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0x01, 0x12, 0x2d, 0x0a, 0x20, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f,
	0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f,
	0x65, 0x78, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x10, 0x8c, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0x01, 0x12, 0x2d, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x62,
	0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x65,
	0x78, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x10, 0x8b, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0x01, 0x12, 0x28, 0x0a, 0x1b, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x65, 0x78, 0x63, 0x65, 0x65, 0x64, 0x65,
	0x64, 0x10, 0x8a, 0xf8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x26, 0x0a, 0x19,
	0x70, 0x72, 0x65, 0x5f, 0x69, 0x72, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x10, 0x89, 0xf8, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73,
	0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x76, 0x32,
	0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_koinos_chain_error_proto_rawDescOnce sync.Once
	file_koinos_chain_error_proto_rawDescData = file_koinos_chain_error_proto_rawDesc
)

func file_koinos_chain_error_proto_rawDescGZIP() []byte {
	file_koinos_chain_error_proto_rawDescOnce.Do(func() {
		file_koinos_chain_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_koinos_chain_error_proto_rawDescData)
	})
	return file_koinos_chain_error_proto_rawDescData
}

var file_koinos_chain_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_koinos_chain_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_koinos_chain_error_proto_goTypes = []interface{}{
	(ErrorCode)(0),       // 0: koinos.chain.error_code
	(*ErrorDetails)(nil), // 1: koinos.chain.error_details
}
var file_koinos_chain_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_koinos_chain_error_proto_init() }
func file_koinos_chain_error_proto_init() {
	if File_koinos_chain_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_koinos_chain_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorDetails); i {
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
			RawDescriptor: file_koinos_chain_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_koinos_chain_error_proto_goTypes,
		DependencyIndexes: file_koinos_chain_error_proto_depIdxs,
		EnumInfos:         file_koinos_chain_error_proto_enumTypes,
		MessageInfos:      file_koinos_chain_error_proto_msgTypes,
	}.Build()
	File_koinos_chain_error_proto = out.File
	file_koinos_chain_error_proto_rawDesc = nil
	file_koinos_chain_error_proto_goTypes = nil
	file_koinos_chain_error_proto_depIdxs = nil
}
