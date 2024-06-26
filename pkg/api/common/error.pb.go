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
// source: api/common/error.proto

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

// ErrorDetail specifies optional error details
type ErrorDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name specifies error details name
	Name *string `protobuf:"bytes,100,opt,name=name,proto3,oneof" json:"name,omitempty"`
	// Text specifies error details text
	Text *string `protobuf:"bytes,200,opt,name=text,proto3,oneof" json:"text,omitempty"`
}

func (x *ErrorDetail) Reset() {
	*x = ErrorDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (*ErrorDetail) ProtoMessage() {}

func (x *ErrorDetail) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorDetail.ProtoReflect.Descriptor instead.
func (*ErrorDetail) Descriptor() ([]byte, []int) {
	return file_api_common_error_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorDetail) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ErrorDetail) GetText() string {
	if x != nil && x.Text != nil {
		return *x.Text
	}
	return ""
}

// Error specifies general error
type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Code specifies error code
	Code int64 `protobuf:"varint,100,opt,name=code,proto3" json:"code,omitempty"`
	// Msg specifies error message
	Msg string `protobuf:"bytes,200,opt,name=msg,proto3" json:"msg,omitempty"`
	// Details specifies multiple error details of the error
	Details []*ErrorDetail `protobuf:"bytes,300,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_common_error_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_api_common_error_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_api_common_error_proto_rawDescGZIP(), []int{1}
}

func (x *Error) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *Error) GetDetails() []*ErrorDetail {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_api_common_error_proto protoreflect.FileDescriptor

var file_api_common_error_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x52, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x22, 0x62, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x11, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0xc8, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x32, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0xac, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x2c, 0x5a, 0x2a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x6e, 0x61, 0x72,
	0x6c, 0x79, 0x2d, 0x69, 0x6f, 0x2f, 0x61, 0x74, 0x6c, 0x61, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_common_error_proto_rawDescOnce sync.Once
	file_api_common_error_proto_rawDescData = file_api_common_error_proto_rawDesc
)

func file_api_common_error_proto_rawDescGZIP() []byte {
	file_api_common_error_proto_rawDescOnce.Do(func() {
		file_api_common_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_common_error_proto_rawDescData)
	})
	return file_api_common_error_proto_rawDescData
}

var file_api_common_error_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_common_error_proto_goTypes = []interface{}{
	(*ErrorDetail)(nil), // 0: api.common.ErrorDetail
	(*Error)(nil),       // 1: api.common.Error
}
var file_api_common_error_proto_depIdxs = []int32{
	0, // 0: api.common.Error.details:type_name -> api.common.ErrorDetail
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_common_error_proto_init() }
func file_api_common_error_proto_init() {
	if File_api_common_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_common_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorDetail); i {
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
		file_api_common_error_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
	file_api_common_error_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_common_error_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_common_error_proto_goTypes,
		DependencyIndexes: file_api_common_error_proto_depIdxs,
		MessageInfos:      file_api_common_error_proto_msgTypes,
	}.Build()
	File_api_common_error_proto = out.File
	file_api_common_error_proto_rawDesc = nil
	file_api_common_error_proto_goTypes = nil
	file_api_common_error_proto_depIdxs = nil
}
