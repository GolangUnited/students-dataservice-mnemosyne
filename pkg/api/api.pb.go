// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: api.proto

package api

import (
	common "github.com/NEKETSKY/mnemosyne/pkg/api/common"
	helloworld "github.com/NEKETSKY/mnemosyne/pkg/api/helloworld"
	interview "github.com/NEKETSKY/mnemosyne/pkg/api/interview"
	user "github.com/NEKETSKY/mnemosyne/pkg/api/user"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65,
	0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf8, 0x07, 0x0a,
	0x09, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x73, 0x79, 0x6e, 0x65, 0x12, 0x53, 0x0a, 0x08, 0x53, 0x61,
	0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x18, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f,
	0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x73, 0x79, 0x6e, 0x65, 0x12,
	0x34, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0a, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x49, 0x64, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x3a, 0x01, 0x2a, 0x22, 0x05,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x1a, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x22, 0x0d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x33, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49,
	0x64, 0x1a, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x12, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x12, 0x42, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x16, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x3a, 0x01, 0x2a, 0x1a, 0x05,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x64, 0x1a, 0x16, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x2a, 0x0a, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x3c, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x08, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x49,
	0x64, 0x1a, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x4d, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x1a, 0x16, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77,
	0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x1a, 0x0d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x4d, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x12, 0x14, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x1a,
	0x0d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x49, 0x64, 0x22, 0x15,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a, 0x22, 0x0a, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x69, 0x65, 0x77, 0x12, 0x48, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65,
	0x77, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x22, 0x12, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x12,
	0x4c, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x12,
	0x0d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x49, 0x64, 0x1a, 0x14,
	0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x77, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5b, 0x0a,
	0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77,
	0x12, 0x14, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x1a, 0x16, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x1a, 0x0f, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x48, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x08, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x2a, 0x0d, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x51, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x12, 0x0d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x69, 0x65, 0x77, 0x2e, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x2a, 0x0f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69,
	0x65, 0x77, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x9b, 0x04, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x45, 0x4b, 0x45, 0x54, 0x53, 0x4b, 0x59, 0x2f,
	0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x73, 0x79, 0x6e, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x92, 0x41, 0xf0, 0x03, 0x12, 0x45, 0x0a, 0x11, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x73, 0x79,
	0x6e, 0x65, 0x20, 0x41, 0x70, 0x70, 0x20, 0x41, 0x50, 0x49, 0x12, 0x24, 0x41, 0x50, 0x49, 0x20,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x4d, 0x6e, 0x65, 0x6d, 0x6f,
	0x73, 0x79, 0x6e, 0x65, 0x20, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x05, 0x0a, 0x03, 0x42, 0x61, 0x72, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x1a, 0x0e, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x68, 0x6f, 0x73, 0x74, 0x3a, 0x38, 0x30, 0x30, 0x30, 0x2a, 0x02, 0x01, 0x02,
	0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x50, 0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x49, 0x0a, 0x47, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x68,
	0x61, 0x76, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x74,
	0x6f, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x3b, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x34, 0x0a,
	0x2a, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x64, 0x6f, 0x65, 0x73,
	0x20, 0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x2e, 0x12, 0x06, 0x0a, 0x04, 0x9a,
	0x02, 0x01, 0x07, 0x5a, 0xbe, 0x01, 0x0a, 0xaa, 0x01, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65,
	0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x9b, 0x01, 0x08, 0x02, 0x1a, 0x09, 0x58, 0x2d, 0x41, 0x50,
	0x49, 0x2d, 0x4b, 0x65, 0x79, 0x20, 0x02, 0x4a, 0x60, 0x0a, 0x1e, 0x78, 0x2d, 0x61, 0x6d, 0x61,
	0x7a, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2d, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x3e, 0x2a, 0x3c, 0x0a, 0x29, 0x0a,
	0x1c, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x54, 0x74, 0x6c, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x09, 0x11,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x4e, 0x40, 0x0a, 0x0f, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x07, 0x1a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4a, 0x28, 0x0a, 0x1c, 0x78, 0x2d, 0x61,
	0x6d, 0x61, 0x7a, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2d, 0x61, 0x75, 0x74, 0x68, 0x74, 0x79, 0x70, 0x65, 0x12, 0x08, 0x1a, 0x06, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x32, 0x0a, 0x0f, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x02, 0x08, 0x01, 0x62, 0x1f, 0x0a, 0x0e, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x00, 0x0a, 0x0d, 0x0a, 0x09, 0x42, 0x61, 0x73, 0x69, 0x63, 0x41,
	0x75, 0x74, 0x68, 0x12, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_proto_goTypes = []interface{}{
	(*helloworld.HelloRequest)(nil), // 0: helloworld.HelloRequest
	(*user.User)(nil),               // 1: user.User
	(*user.Options)(nil),            // 2: user.Options
	(*user.Id)(nil),                 // 3: user.Id
	(*user.Contact)(nil),            // 4: user.Contact
	(*interview.Interview)(nil),     // 5: interview.Interview
	(*common.Empty)(nil),            // 6: common.Empty
	(*interview.Id)(nil),            // 7: interview.Id
	(*helloworld.HelloReply)(nil),   // 8: helloworld.HelloReply
	(*user.Users)(nil),              // 9: user.Users
}
var file_api_proto_depIdxs = []int32{
	0,  // 0: api.Mnemosyne.SayHello:input_type -> helloworld.HelloRequest
	1,  // 1: api.Mnemosyne.CreateUser:input_type -> user.User
	2,  // 2: api.Mnemosyne.GetUsers:input_type -> user.Options
	3,  // 3: api.Mnemosyne.GetUser:input_type -> user.Id
	1,  // 4: api.Mnemosyne.UpdateUser:input_type -> user.User
	3,  // 5: api.Mnemosyne.DeleteUser:input_type -> user.Id
	3,  // 6: api.Mnemosyne.GetContact:input_type -> user.Id
	4,  // 7: api.Mnemosyne.UpdateContact:input_type -> user.Contact
	5,  // 8: api.Mnemosyne.CreateInterview:input_type -> interview.Interview
	6,  // 9: api.Mnemosyne.GetInterviews:input_type -> common.Empty
	7,  // 10: api.Mnemosyne.GetInterview:input_type -> interview.Id
	5,  // 11: api.Mnemosyne.UpdateInterview:input_type -> interview.Interview
	3,  // 12: api.Mnemosyne.DeleteContact:input_type -> user.Id
	7,  // 13: api.Mnemosyne.DeleteInterview:input_type -> interview.Id
	8,  // 14: api.Mnemosyne.SayHello:output_type -> helloworld.HelloReply
	3,  // 15: api.Mnemosyne.CreateUser:output_type -> user.Id
	9,  // 16: api.Mnemosyne.GetUsers:output_type -> user.Users
	1,  // 17: api.Mnemosyne.GetUser:output_type -> user.User
	8,  // 18: api.Mnemosyne.UpdateUser:output_type -> helloworld.HelloReply
	8,  // 19: api.Mnemosyne.DeleteUser:output_type -> helloworld.HelloReply
	4,  // 20: api.Mnemosyne.GetContact:output_type -> user.Contact
	8,  // 21: api.Mnemosyne.UpdateContact:output_type -> helloworld.HelloReply
	7,  // 22: api.Mnemosyne.CreateInterview:output_type -> interview.Id
	5,  // 23: api.Mnemosyne.GetInterviews:output_type -> interview.Interview
	5,  // 24: api.Mnemosyne.GetInterview:output_type -> interview.Interview
	8,  // 25: api.Mnemosyne.UpdateInterview:output_type -> helloworld.HelloReply
	8,  // 26: api.Mnemosyne.DeleteContact:output_type -> helloworld.HelloReply
	8,  // 27: api.Mnemosyne.DeleteInterview:output_type -> helloworld.HelloReply
	14, // [14:28] is the sub-list for method output_type
	0,  // [0:14] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
