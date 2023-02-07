// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: koinos/rpc/services.proto

package services

import (
	account_history "github.com/koinos/koinos-proto-golang/koinos/rpc/account_history"
	block_store "github.com/koinos/koinos-proto-golang/koinos/rpc/block_store"
	chain "github.com/koinos/koinos-proto-golang/koinos/rpc/chain"
	contract_meta_store "github.com/koinos/koinos-proto-golang/koinos/rpc/contract_meta_store"
	mempool "github.com/koinos/koinos-proto-golang/koinos/rpc/mempool"
	p2p "github.com/koinos/koinos-proto-golang/koinos/rpc/p2p"
	transaction_store "github.com/koinos/koinos-proto-golang/koinos/rpc/transaction_store"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_koinos_rpc_services_proto protoreflect.FileDescriptor

var file_koinos_rpc_services_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x6b, 0x6f, 0x69,
	0x6e, 0x6f, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x34, 0x6b, 0x6f,
	0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2c, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x3c, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x24, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6d, 0x65, 0x6d,
	0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x6d, 0x65, 0x6d, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x72, 0x70, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x32, 0x70, 0x2f, 0x70, 0x32, 0x70, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9d,
	0x11, 0x0a, 0x06, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x12, 0x8a, 0x01, 0x0a, 0x13, 0x67, 0x65,
	0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x37, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x67,
	0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x6b, 0x6f, 0x69,
	0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5f, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x79, 0x0a, 0x10, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x12, 0x30, 0x2e, 0x6b, 0x6f, 0x69,
	0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x5f, 0x62,
	0x79, 0x5f, 0x69, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6b,
	0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x85, 0x01, 0x0a, 0x14, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x5f, 0x62, 0x79, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x34, 0x2e, 0x6b, 0x6f, 0x69,
	0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x5f, 0x62,
	0x79, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x35, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7c, 0x0a, 0x11, 0x67, 0x65, 0x74,
	0x5f, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x31,
	0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x68, 0x69, 0x67, 0x68,
	0x65, 0x73, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x32, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x68,
	0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x26, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x73, 0x0a, 0x12, 0x73, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2c, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d,
	0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2e, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x64, 0x0a, 0x0d, 0x67, 0x65, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x12, 0x27, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6b, 0x6f, 0x69, 0x6e,
	0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x65, 0x74,
	0x5f, 0x68, 0x65, 0x61, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x0c, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x5f, 0x69, 0x64, 0x12, 0x26, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x2e, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x0e, 0x67, 0x65, 0x74, 0x5f,
	0x66, 0x6f, 0x72, 0x6b, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x73, 0x12, 0x28, 0x2e, 0x6b, 0x6f, 0x69,
	0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x65,
	0x74, 0x5f, 0x66, 0x6f, 0x72, 0x6b, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x73, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x6b,
	0x5f, 0x68, 0x65, 0x61, 0x64, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x64, 0x0a, 0x0d, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x12, 0x27, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6b, 0x6f,
	0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x70, 0x0a, 0x11, 0x67, 0x65, 0x74, 0x5f, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x2b, 0x2e, 0x6b,
	0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x67, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6b, 0x6f, 0x69, 0x6e,
	0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x65, 0x74,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x67, 0x0a, 0x0e, 0x67, 0x65, 0x74,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x63, 0x12, 0x28, 0x2e, 0x6b, 0x6f,
	0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67,
	0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x63, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72,
	0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x63, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x76, 0x0a, 0x13, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x2d, 0x2e, 0x6b, 0x6f, 0x69, 0x6e,
	0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x65, 0x74,
	0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f,
	0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x67, 0x65, 0x74, 0x5f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8c, 0x01, 0x0a, 0x11, 0x67,
	0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x12, 0x39, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x6d,
	0x65, 0x74, 0x61, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x6b, 0x6f,
	0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x67, 0x65, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x89, 0x01, 0x0a, 0x18, 0x67, 0x65,
	0x74, 0x5f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x34, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x6d, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x67, 0x65, 0x74, 0x5f,
	0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x6b,
	0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x6d, 0x70, 0x6f, 0x6f,
	0x6c, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x9e, 0x01, 0x0a, 0x1f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f,
	0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x3b, 0x2e, 0x6b, 0x6f, 0x69, 0x6e,
	0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x6d, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x5f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3c, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x6d, 0x65, 0x6d, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x5f, 0x70, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6c, 0x0a, 0x11, 0x67, 0x65, 0x74, 0x5f, 0x67, 0x6f,
	0x73, 0x73, 0x69, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x29, 0x2e, 0x6b, 0x6f,
	0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x67, 0x65, 0x74,
	0x5f, 0x67, 0x6f, 0x73, 0x73, 0x69, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x32, 0x70, 0x2e, 0x67, 0x65, 0x74, 0x5f, 0x67, 0x6f, 0x73, 0x73,
	0x69, 0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x97, 0x01, 0x0a, 0x16, 0x67, 0x65, 0x74, 0x5f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x12,
	0x3c, 0x2e, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x67,
	0x65, 0x74, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f,
	0x62, 0x79, 0x5f, 0x69, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e,
	0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x67, 0x65, 0x74,
	0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x62, 0x79,
	0x5f, 0x69, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x37,
	0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6f, 0x69,
	0x6e, 0x6f, 0x73, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6b, 0x6f, 0x69, 0x6e, 0x6f, 0x73, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_koinos_rpc_services_proto_goTypes = []interface{}{
	(*account_history.GetAccountHistoryRequest)(nil),      // 0: koinos.rpc.account_history.get_account_history_request
	(*block_store.GetBlocksByIdRequest)(nil),              // 1: koinos.rpc.block_store.get_blocks_by_id_request
	(*block_store.GetBlocksByHeightRequest)(nil),          // 2: koinos.rpc.block_store.get_blocks_by_height_request
	(*block_store.GetHighestBlockRequest)(nil),            // 3: koinos.rpc.block_store.get_highest_block_request
	(*chain.SubmitBlockRequest)(nil),                      // 4: koinos.rpc.chain.submit_block_request
	(*chain.SubmitTransactionRequest)(nil),                // 5: koinos.rpc.chain.submit_transaction_request
	(*chain.GetHeadInfoRequest)(nil),                      // 6: koinos.rpc.chain.get_head_info_request
	(*chain.GetChainIdRequest)(nil),                       // 7: koinos.rpc.chain.get_chain_id_request
	(*chain.GetForkHeadsRequest)(nil),                     // 8: koinos.rpc.chain.get_fork_heads_request
	(*chain.ReadContractRequest)(nil),                     // 9: koinos.rpc.chain.read_contract_request
	(*chain.GetAccountNonceRequest)(nil),                  // 10: koinos.rpc.chain.get_account_nonce_request
	(*chain.GetAccountRcRequest)(nil),                     // 11: koinos.rpc.chain.get_account_rc_request
	(*chain.GetResourceLimitsRequest)(nil),                // 12: koinos.rpc.chain.get_resource_limits_request
	(*contract_meta_store.GetContractMetaRequest)(nil),    // 13: koinos.rpc.contract_meta_store.get_contract_meta_request
	(*mempool.GetPendingTransactionsRequest)(nil),         // 14: koinos.rpc.mempool.get_pending_transactions_request
	(*mempool.CheckPendingAccountResourcesRequest)(nil),   // 15: koinos.rpc.mempool.check_pending_account_resources_request
	(*p2p.GetGossipStatusRequest)(nil),                    // 16: koinos.rpc.p2p.get_gossip_status_request
	(*transaction_store.GetTransactionsByIdRequest)(nil),  // 17: koinos.rpc.transaction_store.get_transactions_by_id_request
	(*account_history.GetAccountHistoryResponse)(nil),     // 18: koinos.rpc.account_history.get_account_history_response
	(*block_store.GetBlocksByIdResponse)(nil),             // 19: koinos.rpc.block_store.get_blocks_by_id_response
	(*block_store.GetBlocksByHeightResponse)(nil),         // 20: koinos.rpc.block_store.get_blocks_by_height_response
	(*block_store.GetHighestBlockResponse)(nil),           // 21: koinos.rpc.block_store.get_highest_block_response
	(*chain.SubmitBlockResponse)(nil),                     // 22: koinos.rpc.chain.submit_block_response
	(*chain.SubmitTransactionResponse)(nil),               // 23: koinos.rpc.chain.submit_transaction_response
	(*chain.GetHeadInfoResponse)(nil),                     // 24: koinos.rpc.chain.get_head_info_response
	(*chain.GetChainIdResponse)(nil),                      // 25: koinos.rpc.chain.get_chain_id_response
	(*chain.GetForkHeadsResponse)(nil),                    // 26: koinos.rpc.chain.get_fork_heads_response
	(*chain.ReadContractResponse)(nil),                    // 27: koinos.rpc.chain.read_contract_response
	(*chain.GetAccountNonceResponse)(nil),                 // 28: koinos.rpc.chain.get_account_nonce_response
	(*chain.GetAccountRcResponse)(nil),                    // 29: koinos.rpc.chain.get_account_rc_response
	(*chain.GetResourceLimitsResponse)(nil),               // 30: koinos.rpc.chain.get_resource_limits_response
	(*contract_meta_store.GetContractMetaResponse)(nil),   // 31: koinos.rpc.contract_meta_store.get_contract_meta_response
	(*mempool.GetPendingTransactionsResponse)(nil),        // 32: koinos.rpc.mempool.get_pending_transactions_response
	(*mempool.CheckPendingAccountResourcesResponse)(nil),  // 33: koinos.rpc.mempool.check_pending_account_resources_response
	(*p2p.GetGossipStatusResponse)(nil),                   // 34: koinos.rpc.p2p.get_gossip_status_response
	(*transaction_store.GetTransactionsByIdResponse)(nil), // 35: koinos.rpc.transaction_store.get_transactions_by_id_response
}
var file_koinos_rpc_services_proto_depIdxs = []int32{
	0,  // 0: koinos.services.koinos.get_account_history:input_type -> koinos.rpc.account_history.get_account_history_request
	1,  // 1: koinos.services.koinos.get_blocks_by_id:input_type -> koinos.rpc.block_store.get_blocks_by_id_request
	2,  // 2: koinos.services.koinos.get_blocks_by_height:input_type -> koinos.rpc.block_store.get_blocks_by_height_request
	3,  // 3: koinos.services.koinos.get_highest_block:input_type -> koinos.rpc.block_store.get_highest_block_request
	4,  // 4: koinos.services.koinos.submit_block:input_type -> koinos.rpc.chain.submit_block_request
	5,  // 5: koinos.services.koinos.submit_transaction:input_type -> koinos.rpc.chain.submit_transaction_request
	6,  // 6: koinos.services.koinos.get_head_info:input_type -> koinos.rpc.chain.get_head_info_request
	7,  // 7: koinos.services.koinos.get_chain_id:input_type -> koinos.rpc.chain.get_chain_id_request
	8,  // 8: koinos.services.koinos.get_fork_heads:input_type -> koinos.rpc.chain.get_fork_heads_request
	9,  // 9: koinos.services.koinos.read_contract:input_type -> koinos.rpc.chain.read_contract_request
	10, // 10: koinos.services.koinos.get_account_nonce:input_type -> koinos.rpc.chain.get_account_nonce_request
	11, // 11: koinos.services.koinos.get_account_rc:input_type -> koinos.rpc.chain.get_account_rc_request
	12, // 12: koinos.services.koinos.get_resource_limits:input_type -> koinos.rpc.chain.get_resource_limits_request
	13, // 13: koinos.services.koinos.get_contract_meta:input_type -> koinos.rpc.contract_meta_store.get_contract_meta_request
	14, // 14: koinos.services.koinos.get_pending_transactions:input_type -> koinos.rpc.mempool.get_pending_transactions_request
	15, // 15: koinos.services.koinos.check_pending_account_resources:input_type -> koinos.rpc.mempool.check_pending_account_resources_request
	16, // 16: koinos.services.koinos.get_gossip_status:input_type -> koinos.rpc.p2p.get_gossip_status_request
	17, // 17: koinos.services.koinos.get_transactions_by_id:input_type -> koinos.rpc.transaction_store.get_transactions_by_id_request
	18, // 18: koinos.services.koinos.get_account_history:output_type -> koinos.rpc.account_history.get_account_history_response
	19, // 19: koinos.services.koinos.get_blocks_by_id:output_type -> koinos.rpc.block_store.get_blocks_by_id_response
	20, // 20: koinos.services.koinos.get_blocks_by_height:output_type -> koinos.rpc.block_store.get_blocks_by_height_response
	21, // 21: koinos.services.koinos.get_highest_block:output_type -> koinos.rpc.block_store.get_highest_block_response
	22, // 22: koinos.services.koinos.submit_block:output_type -> koinos.rpc.chain.submit_block_response
	23, // 23: koinos.services.koinos.submit_transaction:output_type -> koinos.rpc.chain.submit_transaction_response
	24, // 24: koinos.services.koinos.get_head_info:output_type -> koinos.rpc.chain.get_head_info_response
	25, // 25: koinos.services.koinos.get_chain_id:output_type -> koinos.rpc.chain.get_chain_id_response
	26, // 26: koinos.services.koinos.get_fork_heads:output_type -> koinos.rpc.chain.get_fork_heads_response
	27, // 27: koinos.services.koinos.read_contract:output_type -> koinos.rpc.chain.read_contract_response
	28, // 28: koinos.services.koinos.get_account_nonce:output_type -> koinos.rpc.chain.get_account_nonce_response
	29, // 29: koinos.services.koinos.get_account_rc:output_type -> koinos.rpc.chain.get_account_rc_response
	30, // 30: koinos.services.koinos.get_resource_limits:output_type -> koinos.rpc.chain.get_resource_limits_response
	31, // 31: koinos.services.koinos.get_contract_meta:output_type -> koinos.rpc.contract_meta_store.get_contract_meta_response
	32, // 32: koinos.services.koinos.get_pending_transactions:output_type -> koinos.rpc.mempool.get_pending_transactions_response
	33, // 33: koinos.services.koinos.check_pending_account_resources:output_type -> koinos.rpc.mempool.check_pending_account_resources_response
	34, // 34: koinos.services.koinos.get_gossip_status:output_type -> koinos.rpc.p2p.get_gossip_status_response
	35, // 35: koinos.services.koinos.get_transactions_by_id:output_type -> koinos.rpc.transaction_store.get_transactions_by_id_response
	18, // [18:36] is the sub-list for method output_type
	0,  // [0:18] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_koinos_rpc_services_proto_init() }
func file_koinos_rpc_services_proto_init() {
	if File_koinos_rpc_services_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_koinos_rpc_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_koinos_rpc_services_proto_goTypes,
		DependencyIndexes: file_koinos_rpc_services_proto_depIdxs,
	}.Build()
	File_koinos_rpc_services_proto = out.File
	file_koinos_rpc_services_proto_rawDesc = nil
	file_koinos_rpc_services_proto_goTypes = nil
	file_koinos_rpc_services_proto_depIdxs = nil
}
