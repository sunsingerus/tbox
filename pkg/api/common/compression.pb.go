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

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: api/common/compression.proto

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

// Compression describes compression of the object
type Compression struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type specifies type of compression
	Type int32 `protobuf:"varint,100,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Compression) Reset() {
	*x = Compression{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_compression_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (*Compression) ProtoMessage() {}

func (x *Compression) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_compression_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Compression.ProtoReflect.Descriptor instead.
func (*Compression) Descriptor() ([]byte, []int) {
	return file_api_common_compression_proto_rawDescGZIP(), []int{0}
}

func (x *Compression) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

var File_api_common_compression_proto protoreflect.FileDescriptor

var file_api_common_compression_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x21, 0x0a, 0x0b, 0x43, 0x6f,
	0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x2c, 0x5a,
	0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x6e, 0x61,
	0x72, 0x6c, 0x79, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_api_common_compression_proto_rawDescOnce sync.Once
	file_api_common_compression_proto_rawDescData = file_api_common_compression_proto_rawDesc
)

func file_api_common_compression_proto_rawDescGZIP() []byte {
	file_api_common_compression_proto_rawDescOnce.Do(func() {
		file_api_common_compression_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_common_compression_proto_rawDescData)
	})
	return file_api_common_compression_proto_rawDescData
}

var file_api_common_compression_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_common_compression_proto_goTypes = []interface{}{
	(*Compression)(nil), // 0: api.common.Compression
}
var file_api_common_compression_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_common_compression_proto_init() }
func file_api_common_compression_proto_init() {
	if File_api_common_compression_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_common_compression_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Compression); i {
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
			RawDescriptor: file_api_common_compression_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_common_compression_proto_goTypes,
		DependencyIndexes: file_api_common_compression_proto_depIdxs,
		MessageInfos:      file_api_common_compression_proto_msgTypes,
	}.Build()
	File_api_common_compression_proto = out.File
	file_api_common_compression_proto_rawDesc = nil
	file_api_common_compression_proto_goTypes = nil
	file_api_common_compression_proto_depIdxs = nil
}
