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
// source: api/common/data_packet.proto

//*
// DataChunk represents one chunk (block,single piece) of data send used by DataChunks() function in Control Plane

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

// DataPacket is a chunk of data transferred with additional data.
// Can be part of bigger data, transferred by smaller chunks.
// Main difference with DataChunk is that Packet has additional data.
type DataPacket struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DataChunk is the purpose of the whole packet type.
	DataChunk *DataChunk `protobuf:"bytes,100,opt,name=data_chunk,json=dataChunk,proto3" json:"data_chunk,omitempty"`
	// StreamOptions is an optional transport-level information, describing whole data chunk stream,
	// such as: encoding, compression, etc... [Optional].
	StreamOptions *PresentationOptions `protobuf:"bytes,300,opt,name=stream_options,json=streamOptions,proto3,oneof" json:"stream_options,omitempty"`
	// PayloadMetadata provides additional metadata, which describes payload. [Optional].
	PayloadMetadata *Metadata `protobuf:"bytes,400,opt,name=payload_metadata,json=payloadMetadata,proto3,oneof" json:"payload_metadata,omitempty"`
}

func (x *DataPacket) Reset() {
	*x = DataPacket{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_data_packet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (*DataPacket) ProtoMessage() {}

func (x *DataPacket) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_data_packet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataPacket.ProtoReflect.Descriptor instead.
func (*DataPacket) Descriptor() ([]byte, []int) {
	return file_api_common_data_packet_proto_rawDescGZIP(), []int{0}
}

func (x *DataPacket) GetDataChunk() *DataChunk {
	if x != nil {
		return x.DataChunk
	}
	return nil
}

func (x *DataPacket) GetStreamOptions() *PresentationOptions {
	if x != nil {
		return x.StreamOptions
	}
	return nil
}

func (x *DataPacket) GetPayloadMetadata() *Metadata {
	if x != nil {
		return x.PayloadMetadata
	}
	return nil
}

var File_api_common_data_packet_proto protoreflect.FileDescriptor

var file_api_common_data_packet_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x1b, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xff, 0x01, 0x0a, 0x0a, 0x44, 0x61,
	0x74, 0x61, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x34, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x4c,
	0x0a, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xac, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x88, 0x01, 0x01, 0x12, 0x45, 0x0a, 0x10,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x90, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x01, 0x52,
	0x0f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x2c, 0x5a, 0x2a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x6c,
	0x79, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_common_data_packet_proto_rawDescOnce sync.Once
	file_api_common_data_packet_proto_rawDescData = file_api_common_data_packet_proto_rawDesc
)

func file_api_common_data_packet_proto_rawDescGZIP() []byte {
	file_api_common_data_packet_proto_rawDescOnce.Do(func() {
		file_api_common_data_packet_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_common_data_packet_proto_rawDescData)
	})
	return file_api_common_data_packet_proto_rawDescData
}

var file_api_common_data_packet_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_common_data_packet_proto_goTypes = []interface{}{
	(*DataPacket)(nil),          // 0: api.common.DataPacket
	(*DataChunk)(nil),           // 1: api.common.DataChunk
	(*PresentationOptions)(nil), // 2: api.common.PresentationOptions
	(*Metadata)(nil),            // 3: api.common.Metadata
}
var file_api_common_data_packet_proto_depIdxs = []int32{
	1, // 0: api.common.DataPacket.data_chunk:type_name -> api.common.DataChunk
	2, // 1: api.common.DataPacket.stream_options:type_name -> api.common.PresentationOptions
	3, // 2: api.common.DataPacket.payload_metadata:type_name -> api.common.Metadata
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_common_data_packet_proto_init() }
func file_api_common_data_packet_proto_init() {
	if File_api_common_data_packet_proto != nil {
		return
	}
	file_api_common_data_chunk_proto_init()
	file_api_common_metadata_proto_init()
	file_api_common_presentation_options_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_common_data_packet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataPacket); i {
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
	file_api_common_data_packet_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_common_data_packet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_common_data_packet_proto_goTypes,
		DependencyIndexes: file_api_common_data_packet_proto_depIdxs,
		MessageInfos:      file_api_common_data_packet_proto_msgTypes,
	}.Build()
	File_api_common_data_packet_proto = out.File
	file_api_common_data_packet_proto_rawDesc = nil
	file_api_common_data_packet_proto_goTypes = nil
	file_api_common_data_packet_proto_depIdxs = nil
}
