// Copyright The TBox Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//*
// Address is an abstraction over domain-specific addresses.
// Represents all types of addresses in the system.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: api/common/trace.proto

package common

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

// Trace represents an object trace
type Trace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TraceUid specifies UUID of the trace
	TraceUid *UUID `protobuf:"bytes,100,opt,name=trace_uid,json=traceUid,proto3,oneof" json:"trace_uid,omitempty"`
	// NodeUid specifies id the the node in the trace
	NodeUid *string `protobuf:"bytes,200,opt,name=node_uid,json=nodeUid,proto3,oneof" json:"node_uid,omitempty"`
}

func (x *Trace) Reset() {
	*x = Trace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_trace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (*Trace) ProtoMessage() {}

func (x *Trace) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_trace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trace.ProtoReflect.Descriptor instead.
func (*Trace) Descriptor() ([]byte, []int) {
	return file_api_common_trace_proto_rawDescGZIP(), []int{0}
}

func (x *Trace) GetTraceUid() *UUID {
	if x != nil {
		return x.TraceUid
	}
	return nil
}

func (x *Trace) GetNodeUid() string {
	if x != nil && x.NodeUid != nil {
		return *x.NodeUid
	}
	return ""
}

var File_api_common_trace_proto protoreflect.FileDescriptor

var file_api_common_trace_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x75, 0x75, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x05, 0x54,
	0x72, 0x61, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x75, 0x69,
	0x64, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x48, 0x00, 0x52, 0x08, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x55, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x75, 0x69, 0x64, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x6e,
	0x6f, 0x64, 0x65, 0x55, 0x69, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x5f, 0x75, 0x69, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x75, 0x69, 0x64, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x6c, 0x79, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x74,
	0x6c, 0x61, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_common_trace_proto_rawDescOnce sync.Once
	file_api_common_trace_proto_rawDescData = file_api_common_trace_proto_rawDesc
)

func file_api_common_trace_proto_rawDescGZIP() []byte {
	file_api_common_trace_proto_rawDescOnce.Do(func() {
		file_api_common_trace_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_common_trace_proto_rawDescData)
	})
	return file_api_common_trace_proto_rawDescData
}

var file_api_common_trace_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_common_trace_proto_goTypes = []interface{}{
	(*Trace)(nil), // 0: api.common.Trace
	(*UUID)(nil),  // 1: api.common.UUID
}
var file_api_common_trace_proto_depIdxs = []int32{
	1, // 0: api.common.Trace.trace_uid:type_name -> api.common.UUID
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_common_trace_proto_init() }
func file_api_common_trace_proto_init() {
	if File_api_common_trace_proto != nil {
		return
	}
	file_api_common_uuid_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_common_trace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Trace); i {
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
	file_api_common_trace_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_common_trace_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_common_trace_proto_goTypes,
		DependencyIndexes: file_api_common_trace_proto_depIdxs,
		MessageInfos:      file_api_common_trace_proto_msgTypes,
	}.Build()
	File_api_common_trace_proto = out.File
	file_api_common_trace_proto_rawDesc = nil
	file_api_common_trace_proto_goTypes = nil
	file_api_common_trace_proto_depIdxs = nil
}
