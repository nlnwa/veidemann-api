// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.2
// source: uricanonicalizer/v1/uricanonicalizer.proto

package uricanonicalizer

import (
	v1 "github.com/nlnwa/veidemann-api/go/commons/v1"
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

type CanonicalizeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *CanonicalizeRequest) Reset() {
	*x = CanonicalizeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uricanonicalizer_v1_uricanonicalizer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CanonicalizeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanonicalizeRequest) ProtoMessage() {}

func (x *CanonicalizeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_uricanonicalizer_v1_uricanonicalizer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanonicalizeRequest.ProtoReflect.Descriptor instead.
func (*CanonicalizeRequest) Descriptor() ([]byte, []int) {
	return file_uricanonicalizer_v1_uricanonicalizer_proto_rawDescGZIP(), []int{0}
}

func (x *CanonicalizeRequest) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

type CanonicalizeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uri *v1.ParsedUri `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
}

func (x *CanonicalizeResponse) Reset() {
	*x = CanonicalizeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_uricanonicalizer_v1_uricanonicalizer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CanonicalizeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CanonicalizeResponse) ProtoMessage() {}

func (x *CanonicalizeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_uricanonicalizer_v1_uricanonicalizer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CanonicalizeResponse.ProtoReflect.Descriptor instead.
func (*CanonicalizeResponse) Descriptor() ([]byte, []int) {
	return file_uricanonicalizer_v1_uricanonicalizer_proto_rawDescGZIP(), []int{1}
}

func (x *CanonicalizeResponse) GetUri() *v1.ParsedUri {
	if x != nil {
		return x.Uri
	}
	return nil
}

var File_uricanonicalizer_v1_uricanonicalizer_proto protoreflect.FileDescriptor

var file_uricanonicalizer_v1_uricanonicalizer_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x75, 0x72, 0x69, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x72, 0x69, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x76, 0x65,
	0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x72, 0x69, 0x63,
	0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a,
	0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x13, 0x43,
	0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x69, 0x22, 0x4d, 0x0a, 0x14, 0x43, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x03,
	0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x76, 0x65, 0x69, 0x64,
	0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x64, 0x55, 0x72, 0x69, 0x52, 0x03,
	0x75, 0x72, 0x69, 0x32, 0x9d, 0x01, 0x0a, 0x17, 0x55, 0x72, 0x69, 0x43, 0x61, 0x6e, 0x6f, 0x6e,
	0x69, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x81, 0x01, 0x0a, 0x0c, 0x43, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x12, 0x36, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x75, 0x72, 0x69, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65,
	0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x72, 0x69, 0x63, 0x61, 0x6e, 0x6f,
	0x6e, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e,
	0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x89, 0x01, 0x0a, 0x2b, 0x6e, 0x6f, 0x2e, 0x6e, 0x62, 0x2e, 0x6e, 0x6e,
	0x61, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x75, 0x72, 0x69, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x42, 0x10, 0x55, 0x72, 0x69, 0x43, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x72, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6c, 0x6e, 0x77, 0x61, 0x2f, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d,
	0x61, 0x6e, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x75, 0x72, 0x69, 0x63, 0x61,
	0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x75,
	0x72, 0x69, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_uricanonicalizer_v1_uricanonicalizer_proto_rawDescOnce sync.Once
	file_uricanonicalizer_v1_uricanonicalizer_proto_rawDescData = file_uricanonicalizer_v1_uricanonicalizer_proto_rawDesc
)

func file_uricanonicalizer_v1_uricanonicalizer_proto_rawDescGZIP() []byte {
	file_uricanonicalizer_v1_uricanonicalizer_proto_rawDescOnce.Do(func() {
		file_uricanonicalizer_v1_uricanonicalizer_proto_rawDescData = protoimpl.X.CompressGZIP(file_uricanonicalizer_v1_uricanonicalizer_proto_rawDescData)
	})
	return file_uricanonicalizer_v1_uricanonicalizer_proto_rawDescData
}

var file_uricanonicalizer_v1_uricanonicalizer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_uricanonicalizer_v1_uricanonicalizer_proto_goTypes = []interface{}{
	(*CanonicalizeRequest)(nil),  // 0: veidemann.api.uricanonicalizer.v1.CanonicalizeRequest
	(*CanonicalizeResponse)(nil), // 1: veidemann.api.uricanonicalizer.v1.CanonicalizeResponse
	(*v1.ParsedUri)(nil),         // 2: veidemann.api.commons.v1.ParsedUri
}
var file_uricanonicalizer_v1_uricanonicalizer_proto_depIdxs = []int32{
	2, // 0: veidemann.api.uricanonicalizer.v1.CanonicalizeResponse.uri:type_name -> veidemann.api.commons.v1.ParsedUri
	0, // 1: veidemann.api.uricanonicalizer.v1.UriCanonicalizerService.Canonicalize:input_type -> veidemann.api.uricanonicalizer.v1.CanonicalizeRequest
	1, // 2: veidemann.api.uricanonicalizer.v1.UriCanonicalizerService.Canonicalize:output_type -> veidemann.api.uricanonicalizer.v1.CanonicalizeResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_uricanonicalizer_v1_uricanonicalizer_proto_init() }
func file_uricanonicalizer_v1_uricanonicalizer_proto_init() {
	if File_uricanonicalizer_v1_uricanonicalizer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_uricanonicalizer_v1_uricanonicalizer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CanonicalizeRequest); i {
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
		file_uricanonicalizer_v1_uricanonicalizer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CanonicalizeResponse); i {
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
			RawDescriptor: file_uricanonicalizer_v1_uricanonicalizer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_uricanonicalizer_v1_uricanonicalizer_proto_goTypes,
		DependencyIndexes: file_uricanonicalizer_v1_uricanonicalizer_proto_depIdxs,
		MessageInfos:      file_uricanonicalizer_v1_uricanonicalizer_proto_msgTypes,
	}.Build()
	File_uricanonicalizer_v1_uricanonicalizer_proto = out.File
	file_uricanonicalizer_v1_uricanonicalizer_proto_rawDesc = nil
	file_uricanonicalizer_v1_uricanonicalizer_proto_goTypes = nil
	file_uricanonicalizer_v1_uricanonicalizer_proto_depIdxs = nil
}
