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
// source: api/service/base_control_plane.proto

package service

import (
	common "github.com/sunsingerus/tbox/pkg/api/common"
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

var File_api_service_base_control_plane_proto protoreflect.FileDescriptor

var file_api_service_base_control_plane_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x1a, 0x17, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x78, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6c,
	0x61, 0x6e, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x10, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x10,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x35, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x12, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x1a, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x22, 0x00, 0x28, 0x01, 0x42, 0x2d, 0x5a,
	0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x6e, 0x61,
	0x72, 0x6c, 0x79, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_api_service_base_control_plane_proto_goTypes = []interface{}{
	(*common.Task)(nil),   // 0: api.common.Task
	(*common.Metric)(nil), // 1: api.common.Metric
}
var file_api_service_base_control_plane_proto_depIdxs = []int32{
	0, // 0: api.service.ControlPlane.Tasks:input_type -> api.common.Task
	1, // 1: api.service.ControlPlane.Metrics:input_type -> api.common.Metric
	0, // 2: api.service.ControlPlane.Tasks:output_type -> api.common.Task
	1, // 3: api.service.ControlPlane.Metrics:output_type -> api.common.Metric
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_service_base_control_plane_proto_init() }
func file_api_service_base_control_plane_proto_init() {
	if File_api_service_base_control_plane_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_service_base_control_plane_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_service_base_control_plane_proto_goTypes,
		DependencyIndexes: file_api_service_base_control_plane_proto_depIdxs,
	}.Build()
	File_api_service_base_control_plane_proto = out.File
	file_api_service_base_control_plane_proto_rawDesc = nil
	file_api_service_base_control_plane_proto_goTypes = nil
	file_api_service_base_control_plane_proto_depIdxs = nil
}
