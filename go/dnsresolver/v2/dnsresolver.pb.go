// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.2
// source: dnsresolver/v2/dnsresolver.proto

package dnsresolver

import (
	v2 "github.com/nlnwa/veidemann-api/go/commons/v2"
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

type ResolveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *ResolveRequest) Reset() {
	*x = ResolveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dnsresolver_v2_dnsresolver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveRequest) ProtoMessage() {}

func (x *ResolveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dnsresolver_v2_dnsresolver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveRequest.ProtoReflect.Descriptor instead.
func (*ResolveRequest) Descriptor() ([]byte, []int) {
	return file_dnsresolver_v2_dnsresolver_proto_rawDescGZIP(), []int{0}
}

func (x *ResolveRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

type ResolveReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host      string    `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	TextualIp string    `protobuf:"bytes,2,opt,name=textual_ip,json=textualIp,proto3" json:"textual_ip,omitempty"`
	RawIp     []byte    `protobuf:"bytes,3,opt,name=raw_ip,json=rawIp,proto3" json:"raw_ip,omitempty"`
	Error     *v2.Error `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *ResolveReply) Reset() {
	*x = ResolveReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dnsresolver_v2_dnsresolver_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveReply) ProtoMessage() {}

func (x *ResolveReply) ProtoReflect() protoreflect.Message {
	mi := &file_dnsresolver_v2_dnsresolver_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveReply.ProtoReflect.Descriptor instead.
func (*ResolveReply) Descriptor() ([]byte, []int) {
	return file_dnsresolver_v2_dnsresolver_proto_rawDescGZIP(), []int{1}
}

func (x *ResolveReply) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ResolveReply) GetTextualIp() string {
	if x != nil {
		return x.TextualIp
	}
	return ""
}

func (x *ResolveReply) GetRawIp() []byte {
	if x != nil {
		return x.RawIp
	}
	return nil
}

func (x *ResolveReply) GetError() *v2.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_dnsresolver_v2_dnsresolver_proto protoreflect.FileDescriptor

var file_dnsresolver_v2_dnsresolver_proto_rawDesc = []byte{
	0x0a, 0x20, 0x64, 0x6e, 0x73, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x32,
	0x2f, 0x64, 0x6e, 0x73, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1c, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x64, 0x6e, 0x73, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x32,
	0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0e,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x22, 0x8f, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x65, 0x78, 0x74, 0x75,
	0x61, 0x6c, 0x5f, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x78,
	0x74, 0x75, 0x61, 0x6c, 0x49, 0x70, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x61, 0x77, 0x5f, 0x69, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x72, 0x61, 0x77, 0x49, 0x70, 0x12, 0x35, 0x0a,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x76,
	0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x32, 0x74, 0x0a, 0x0b, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x72, 0x12, 0x65, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x12, 0x2c,
	0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64,
	0x6e, 0x73, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x76,
	0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x6e, 0x73,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x6c, 0x76, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x7c, 0x0a, 0x26, 0x6e, 0x6f,
	0x2e, 0x6e, 0x62, 0x2e, 0x6e, 0x6e, 0x61, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x6e, 0x73, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x32, 0x42, 0x12, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x01, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6c, 0x6e, 0x77, 0x61, 0x2f, 0x76, 0x65, 0x69,
	0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x6e,
	0x73, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x3b, 0x64, 0x6e, 0x73,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dnsresolver_v2_dnsresolver_proto_rawDescOnce sync.Once
	file_dnsresolver_v2_dnsresolver_proto_rawDescData = file_dnsresolver_v2_dnsresolver_proto_rawDesc
)

func file_dnsresolver_v2_dnsresolver_proto_rawDescGZIP() []byte {
	file_dnsresolver_v2_dnsresolver_proto_rawDescOnce.Do(func() {
		file_dnsresolver_v2_dnsresolver_proto_rawDescData = protoimpl.X.CompressGZIP(file_dnsresolver_v2_dnsresolver_proto_rawDescData)
	})
	return file_dnsresolver_v2_dnsresolver_proto_rawDescData
}

var file_dnsresolver_v2_dnsresolver_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_dnsresolver_v2_dnsresolver_proto_goTypes = []interface{}{
	(*ResolveRequest)(nil), // 0: veidemann.api.dnsresolver.v2.ResolveRequest
	(*ResolveReply)(nil),   // 1: veidemann.api.dnsresolver.v2.ResolveReply
	(*v2.Error)(nil),       // 2: veidemann.api.commons.v2.Error
}
var file_dnsresolver_v2_dnsresolver_proto_depIdxs = []int32{
	2, // 0: veidemann.api.dnsresolver.v2.ResolveReply.error:type_name -> veidemann.api.commons.v2.Error
	0, // 1: veidemann.api.dnsresolver.v2.DnsResolver.resolve:input_type -> veidemann.api.dnsresolver.v2.ResolveRequest
	1, // 2: veidemann.api.dnsresolver.v2.DnsResolver.resolve:output_type -> veidemann.api.dnsresolver.v2.ResolveReply
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_dnsresolver_v2_dnsresolver_proto_init() }
func file_dnsresolver_v2_dnsresolver_proto_init() {
	if File_dnsresolver_v2_dnsresolver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dnsresolver_v2_dnsresolver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveRequest); i {
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
		file_dnsresolver_v2_dnsresolver_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveReply); i {
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
			RawDescriptor: file_dnsresolver_v2_dnsresolver_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dnsresolver_v2_dnsresolver_proto_goTypes,
		DependencyIndexes: file_dnsresolver_v2_dnsresolver_proto_depIdxs,
		MessageInfos:      file_dnsresolver_v2_dnsresolver_proto_msgTypes,
	}.Build()
	File_dnsresolver_v2_dnsresolver_proto = out.File
	file_dnsresolver_v2_dnsresolver_proto_rawDesc = nil
	file_dnsresolver_v2_dnsresolver_proto_goTypes = nil
	file_dnsresolver_v2_dnsresolver_proto_depIdxs = nil
}
