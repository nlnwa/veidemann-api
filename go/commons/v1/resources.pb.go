// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.2
// source: commons/v1/resources.proto

package commons

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

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Detail string `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commons_v1_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_commons_v1_resources_proto_msgTypes[0]
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
	return file_commons_v1_resources_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetCode() int32 {
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

func (x *Error) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

type FieldMask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Paths []string `protobuf:"bytes,1,rep,name=paths,proto3" json:"paths,omitempty"`
}

func (x *FieldMask) Reset() {
	*x = FieldMask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commons_v1_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldMask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldMask) ProtoMessage() {}

func (x *FieldMask) ProtoReflect() protoreflect.Message {
	mi := &file_commons_v1_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldMask.ProtoReflect.Descriptor instead.
func (*FieldMask) Descriptor() ([]byte, []int) {
	return file_commons_v1_resources_proto_rawDescGZIP(), []int{1}
}

func (x *FieldMask) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

type ParsedUri struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The entire uri
	Href string `protobuf:"bytes,1,opt,name=href,proto3" json:"href,omitempty"`
	// The scheme (protocol) part of the uri
	Scheme string `protobuf:"bytes,2,opt,name=scheme,proto3" json:"scheme,omitempty"`
	// The hostname of the uri
	Host string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	// The port number of the uri
	Port int32 `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	// The username part of the uri
	Username string `protobuf:"bytes,5,opt,name=username,proto3" json:"username,omitempty"`
	// The password part of the uri
	Password string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	// The path part of the uri
	Path string `protobuf:"bytes,7,opt,name=path,proto3" json:"path,omitempty"`
	// The query (search) part of the uri
	Query string `protobuf:"bytes,8,opt,name=query,proto3" json:"query,omitempty"`
	// The fragment (hash) part of the uri
	Fragment string `protobuf:"bytes,9,opt,name=fragment,proto3" json:"fragment,omitempty"`
}

func (x *ParsedUri) Reset() {
	*x = ParsedUri{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commons_v1_resources_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParsedUri) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParsedUri) ProtoMessage() {}

func (x *ParsedUri) ProtoReflect() protoreflect.Message {
	mi := &file_commons_v1_resources_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParsedUri.ProtoReflect.Descriptor instead.
func (*ParsedUri) Descriptor() ([]byte, []int) {
	return file_commons_v1_resources_proto_rawDescGZIP(), []int{2}
}

func (x *ParsedUri) GetHref() string {
	if x != nil {
		return x.Href
	}
	return ""
}

func (x *ParsedUri) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *ParsedUri) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ParsedUri) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ParsedUri) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ParsedUri) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *ParsedUri) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ParsedUri) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ParsedUri) GetFragment() string {
	if x != nil {
		return x.Fragment
	}
	return ""
}

type ExtractedText struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WarcId         string `protobuf:"bytes,1,opt,name=warc_id,json=warcId,proto3" json:"warc_id,omitempty"`
	Text           string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	SentenceCount  int64  `protobuf:"varint,3,opt,name=sentence_count,json=sentenceCount,proto3" json:"sentence_count,omitempty"`
	WordCount      int64  `protobuf:"varint,4,opt,name=word_count,json=wordCount,proto3" json:"word_count,omitempty"`
	LongWordCount  int64  `protobuf:"varint,5,opt,name=long_word_count,json=longWordCount,proto3" json:"long_word_count,omitempty"`
	CharacterCount int64  `protobuf:"varint,6,opt,name=character_count,json=characterCount,proto3" json:"character_count,omitempty"`
	Lix            int64  `protobuf:"varint,7,opt,name=lix,proto3" json:"lix,omitempty"`
	Language       string `protobuf:"bytes,8,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *ExtractedText) Reset() {
	*x = ExtractedText{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commons_v1_resources_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExtractedText) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtractedText) ProtoMessage() {}

func (x *ExtractedText) ProtoReflect() protoreflect.Message {
	mi := &file_commons_v1_resources_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtractedText.ProtoReflect.Descriptor instead.
func (*ExtractedText) Descriptor() ([]byte, []int) {
	return file_commons_v1_resources_proto_rawDescGZIP(), []int{3}
}

func (x *ExtractedText) GetWarcId() string {
	if x != nil {
		return x.WarcId
	}
	return ""
}

func (x *ExtractedText) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ExtractedText) GetSentenceCount() int64 {
	if x != nil {
		return x.SentenceCount
	}
	return 0
}

func (x *ExtractedText) GetWordCount() int64 {
	if x != nil {
		return x.WordCount
	}
	return 0
}

func (x *ExtractedText) GetLongWordCount() int64 {
	if x != nil {
		return x.LongWordCount
	}
	return 0
}

func (x *ExtractedText) GetCharacterCount() int64 {
	if x != nil {
		return x.CharacterCount
	}
	return 0
}

func (x *ExtractedText) GetLix() int64 {
	if x != nil {
		return x.Lix
	}
	return 0
}

func (x *ExtractedText) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

var File_commons_v1_resources_proto protoreflect.FileDescriptor

var file_commons_v1_resources_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x76, 0x65,
	0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x45, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x21, 0x0a,
	0x09, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x74, 0x68, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73,
	0x22, 0xdd, 0x01, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x73, 0x65, 0x64, 0x55, 0x72, 0x69, 0x12, 0x12,
	0x0a, 0x04, 0x68, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x72,
	0x65, 0x66, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f,
	0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x14,
	0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x22, 0x81, 0x02, 0x0a, 0x0d, 0x45, 0x78, 0x74, 0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x54, 0x65,
	0x78, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x77, 0x61, 0x72, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x61, 0x72, 0x63, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x6e, 0x63,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x64, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x64,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x77, 0x6f,
	0x72, 0x64, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x6c, 0x6f, 0x6e, 0x67, 0x57, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x27, 0x0a,
	0x0f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x69, 0x78, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x6c, 0x69, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x42, 0x6d, 0x0a, 0x22, 0x6e, 0x6f, 0x2e, 0x6e, 0x62, 0x2e, 0x6e, 0x6e,
	0x61, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x34, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6c, 0x6e, 0x77, 0x61, 0x2f,
	0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commons_v1_resources_proto_rawDescOnce sync.Once
	file_commons_v1_resources_proto_rawDescData = file_commons_v1_resources_proto_rawDesc
)

func file_commons_v1_resources_proto_rawDescGZIP() []byte {
	file_commons_v1_resources_proto_rawDescOnce.Do(func() {
		file_commons_v1_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_commons_v1_resources_proto_rawDescData)
	})
	return file_commons_v1_resources_proto_rawDescData
}

var file_commons_v1_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_commons_v1_resources_proto_goTypes = []interface{}{
	(*Error)(nil),         // 0: veidemann.api.commons.v1.Error
	(*FieldMask)(nil),     // 1: veidemann.api.commons.v1.FieldMask
	(*ParsedUri)(nil),     // 2: veidemann.api.commons.v1.ParsedUri
	(*ExtractedText)(nil), // 3: veidemann.api.commons.v1.ExtractedText
}
var file_commons_v1_resources_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_commons_v1_resources_proto_init() }
func file_commons_v1_resources_proto_init() {
	if File_commons_v1_resources_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commons_v1_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_commons_v1_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FieldMask); i {
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
		file_commons_v1_resources_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParsedUri); i {
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
		file_commons_v1_resources_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExtractedText); i {
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
			RawDescriptor: file_commons_v1_resources_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commons_v1_resources_proto_goTypes,
		DependencyIndexes: file_commons_v1_resources_proto_depIdxs,
		MessageInfos:      file_commons_v1_resources_proto_msgTypes,
	}.Build()
	File_commons_v1_resources_proto = out.File
	file_commons_v1_resources_proto_rawDesc = nil
	file_commons_v1_resources_proto_goTypes = nil
	file_commons_v1_resources_proto_depIdxs = nil
}
