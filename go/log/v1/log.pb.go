// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.2
// source: log/v1/log.proto

package log

import (
	v1 "github.com/nlnwa/veidemann-api/go/commons/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Specification of which entities to get.
type CrawlLogListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Select objects by one or more id's
	WarcId             []string      `protobuf:"bytes,2,rep,name=warc_id,json=warcId,proto3" json:"warc_id,omitempty"`
	QueryTemplate      *CrawlLog     `protobuf:"bytes,5,opt,name=query_template,json=queryTemplate,proto3" json:"query_template,omitempty"`
	QueryMask          *v1.FieldMask `protobuf:"bytes,6,opt,name=query_mask,json=queryMask,proto3" json:"query_mask,omitempty"`
	ReturnedFieldsMask *v1.FieldMask `protobuf:"bytes,7,opt,name=returned_fields_mask,json=returnedFieldsMask,proto3" json:"returned_fields_mask,omitempty"`
	OrderByPath        string        `protobuf:"bytes,8,opt,name=order_by_path,json=orderByPath,proto3" json:"order_by_path,omitempty"`
	OrderDescending    bool          `protobuf:"varint,9,opt,name=order_descending,json=orderDescending,proto3" json:"order_descending,omitempty"`
	Watch              bool          `protobuf:"varint,13,opt,name=watch,proto3" json:"watch,omitempty"`
	PageSize           int32         `protobuf:"varint,14,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Offset             int32         `protobuf:"varint,15,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *CrawlLogListRequest) Reset() {
	*x = CrawlLogListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_v1_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CrawlLogListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CrawlLogListRequest) ProtoMessage() {}

func (x *CrawlLogListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_log_v1_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CrawlLogListRequest.ProtoReflect.Descriptor instead.
func (*CrawlLogListRequest) Descriptor() ([]byte, []int) {
	return file_log_v1_log_proto_rawDescGZIP(), []int{0}
}

func (x *CrawlLogListRequest) GetWarcId() []string {
	if x != nil {
		return x.WarcId
	}
	return nil
}

func (x *CrawlLogListRequest) GetQueryTemplate() *CrawlLog {
	if x != nil {
		return x.QueryTemplate
	}
	return nil
}

func (x *CrawlLogListRequest) GetQueryMask() *v1.FieldMask {
	if x != nil {
		return x.QueryMask
	}
	return nil
}

func (x *CrawlLogListRequest) GetReturnedFieldsMask() *v1.FieldMask {
	if x != nil {
		return x.ReturnedFieldsMask
	}
	return nil
}

func (x *CrawlLogListRequest) GetOrderByPath() string {
	if x != nil {
		return x.OrderByPath
	}
	return ""
}

func (x *CrawlLogListRequest) GetOrderDescending() bool {
	if x != nil {
		return x.OrderDescending
	}
	return false
}

func (x *CrawlLogListRequest) GetWatch() bool {
	if x != nil {
		return x.Watch
	}
	return false
}

func (x *CrawlLogListRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *CrawlLogListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type PageLogListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Select objects by one or more id's
	WarcId             []string      `protobuf:"bytes,2,rep,name=warc_id,json=warcId,proto3" json:"warc_id,omitempty"`
	QueryTemplate      *PageLog      `protobuf:"bytes,5,opt,name=query_template,json=queryTemplate,proto3" json:"query_template,omitempty"`
	QueryMask          *v1.FieldMask `protobuf:"bytes,6,opt,name=query_mask,json=queryMask,proto3" json:"query_mask,omitempty"`
	ReturnedFieldsMask *v1.FieldMask `protobuf:"bytes,7,opt,name=returned_fields_mask,json=returnedFieldsMask,proto3" json:"returned_fields_mask,omitempty"`
	OrderByPath        string        `protobuf:"bytes,8,opt,name=order_by_path,json=orderByPath,proto3" json:"order_by_path,omitempty"`
	OrderDescending    bool          `protobuf:"varint,9,opt,name=order_descending,json=orderDescending,proto3" json:"order_descending,omitempty"`
	Watch              bool          `protobuf:"varint,13,opt,name=watch,proto3" json:"watch,omitempty"`
	PageSize           int32         `protobuf:"varint,14,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Offset             int32         `protobuf:"varint,15,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *PageLogListRequest) Reset() {
	*x = PageLogListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_v1_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageLogListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageLogListRequest) ProtoMessage() {}

func (x *PageLogListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_log_v1_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageLogListRequest.ProtoReflect.Descriptor instead.
func (*PageLogListRequest) Descriptor() ([]byte, []int) {
	return file_log_v1_log_proto_rawDescGZIP(), []int{1}
}

func (x *PageLogListRequest) GetWarcId() []string {
	if x != nil {
		return x.WarcId
	}
	return nil
}

func (x *PageLogListRequest) GetQueryTemplate() *PageLog {
	if x != nil {
		return x.QueryTemplate
	}
	return nil
}

func (x *PageLogListRequest) GetQueryMask() *v1.FieldMask {
	if x != nil {
		return x.QueryMask
	}
	return nil
}

func (x *PageLogListRequest) GetReturnedFieldsMask() *v1.FieldMask {
	if x != nil {
		return x.ReturnedFieldsMask
	}
	return nil
}

func (x *PageLogListRequest) GetOrderByPath() string {
	if x != nil {
		return x.OrderByPath
	}
	return ""
}

func (x *PageLogListRequest) GetOrderDescending() bool {
	if x != nil {
		return x.OrderDescending
	}
	return false
}

func (x *PageLogListRequest) GetWatch() bool {
	if x != nil {
		return x.Watch
	}
	return false
}

func (x *PageLogListRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PageLogListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type WritePageLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*WritePageLogRequest_CrawlLog
	//	*WritePageLogRequest_Resource
	//	*WritePageLogRequest_Outlink
	Value isWritePageLogRequest_Value `protobuf_oneof:"value"`
}

func (x *WritePageLogRequest) Reset() {
	*x = WritePageLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_v1_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WritePageLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WritePageLogRequest) ProtoMessage() {}

func (x *WritePageLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_log_v1_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WritePageLogRequest.ProtoReflect.Descriptor instead.
func (*WritePageLogRequest) Descriptor() ([]byte, []int) {
	return file_log_v1_log_proto_rawDescGZIP(), []int{2}
}

func (m *WritePageLogRequest) GetValue() isWritePageLogRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *WritePageLogRequest) GetCrawlLog() *CrawlLog {
	if x, ok := x.GetValue().(*WritePageLogRequest_CrawlLog); ok {
		return x.CrawlLog
	}
	return nil
}

func (x *WritePageLogRequest) GetResource() *PageLog_Resource {
	if x, ok := x.GetValue().(*WritePageLogRequest_Resource); ok {
		return x.Resource
	}
	return nil
}

func (x *WritePageLogRequest) GetOutlink() string {
	if x, ok := x.GetValue().(*WritePageLogRequest_Outlink); ok {
		return x.Outlink
	}
	return ""
}

type isWritePageLogRequest_Value interface {
	isWritePageLogRequest_Value()
}

type WritePageLogRequest_CrawlLog struct {
	CrawlLog *CrawlLog `protobuf:"bytes,1,opt,name=crawlLog,proto3,oneof"`
}

type WritePageLogRequest_Resource struct {
	Resource *PageLog_Resource `protobuf:"bytes,2,opt,name=resource,proto3,oneof"`
}

type WritePageLogRequest_Outlink struct {
	Outlink string `protobuf:"bytes,3,opt,name=outlink,proto3,oneof"`
}

func (*WritePageLogRequest_CrawlLog) isWritePageLogRequest_Value() {}

func (*WritePageLogRequest_Resource) isWritePageLogRequest_Value() {}

func (*WritePageLogRequest_Outlink) isWritePageLogRequest_Value() {}

type WriteCrawlLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CrawlLog *CrawlLog `protobuf:"bytes,1,opt,name=crawlLog,proto3" json:"crawlLog,omitempty"`
}

func (x *WriteCrawlLogRequest) Reset() {
	*x = WriteCrawlLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_v1_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteCrawlLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteCrawlLogRequest) ProtoMessage() {}

func (x *WriteCrawlLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_log_v1_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteCrawlLogRequest.ProtoReflect.Descriptor instead.
func (*WriteCrawlLogRequest) Descriptor() ([]byte, []int) {
	return file_log_v1_log_proto_rawDescGZIP(), []int{3}
}

func (x *WriteCrawlLogRequest) GetCrawlLog() *CrawlLog {
	if x != nil {
		return x.CrawlLog
	}
	return nil
}

var File_log_v1_log_proto protoreflect.FileDescriptor

var file_log_v1_log_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x14, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x03, 0x0a, 0x13, 0x43, 0x72,
	0x61, 0x77, 0x6c, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x77, 0x61, 0x72, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x77, 0x61, 0x72, 0x63, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x0e, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x5f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x4c,
	0x6f, 0x67, 0x52, 0x0d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x12, 0x42, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e,
	0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x55, 0x0a, 0x14, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65,
	0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x12, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x22, 0x0a, 0x0d,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x50, 0x61, 0x74, 0x68,
	0x12, 0x29, 0x0a, 0x10, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x65, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x44, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x77,
	0x61, 0x74, 0x63, 0x68, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x77, 0x61, 0x74, 0x63,
	0x68, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0xa8, 0x03, 0x0a, 0x12, 0x50, 0x61, 0x67, 0x65, 0x4c,
	0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x77, 0x61, 0x72, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x77, 0x61, 0x72, 0x63, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x0e, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x0d, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x42, 0x0a, 0x0a,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x23, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x61, 0x73, 0x6b,
	0x12, 0x55, 0x0a, 0x14, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d,
	0x61, 0x73, 0x6b, 0x52, 0x12, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x22, 0x0a, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x62, 0x79, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x50, 0x61, 0x74, 0x68, 0x12, 0x29, 0x0a, 0x10, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x44, 0x65, 0x73, 0x63,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x61, 0x74, 0x63, 0x68, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x77, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x22, 0xbe, 0x01, 0x0a, 0x13, 0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x61, 0x67, 0x65, 0x4c,
	0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x08, 0x63, 0x72, 0x61,
	0x77, 0x6c, 0x4c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x76, 0x65,
	0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x4c, 0x6f, 0x67, 0x48, 0x00, 0x52, 0x08, 0x63,
	0x72, 0x61, 0x77, 0x6c, 0x4c, 0x6f, 0x67, 0x12, 0x44, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x76, 0x65, 0x69, 0x64,
	0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a,
	0x07, 0x6f, 0x75, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x07, 0x6f, 0x75, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x52, 0x0a, 0x14, 0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x72, 0x61, 0x77, 0x6c,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x08, 0x63, 0x72,
	0x61, 0x77, 0x6c, 0x4c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x76,
	0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x4c, 0x6f, 0x67, 0x52, 0x08, 0x63, 0x72,
	0x61, 0x77, 0x6c, 0x4c, 0x6f, 0x67, 0x32, 0xf2, 0x02, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x5e,
	0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x4c, 0x6f, 0x67, 0x73, 0x12,
	0x29, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x4c, 0x6f, 0x67, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x76, 0x65, 0x69,
	0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x4c, 0x6f, 0x67, 0x22, 0x00, 0x30, 0x01, 0x12, 0x5b,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x28,
	0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65,
	0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x22, 0x00, 0x30, 0x01, 0x12, 0x57, 0x0a, 0x0d, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x4c, 0x6f, 0x67, 0x12, 0x2a, 0x2e, 0x76,
	0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x43, 0x72, 0x61, 0x77, 0x6c, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x28, 0x01, 0x12, 0x55, 0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x61, 0x67,
	0x65, 0x4c, 0x6f, 0x67, 0x12, 0x29, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x72, 0x69, 0x74,
	0x65, 0x50, 0x61, 0x67, 0x65, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x28, 0x01, 0x42, 0x5c, 0x0a, 0x1e, 0x6e,
	0x6f, 0x2e, 0x6e, 0x62, 0x2e, 0x6e, 0x6e, 0x61, 0x2e, 0x76, 0x65, 0x69, 0x64, 0x65, 0x6d, 0x61,
	0x6e, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x4c,
	0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x6c, 0x6e, 0x77, 0x61, 0x2f, 0x76, 0x65,
	0x69, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x6c,
	0x6f, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_log_v1_log_proto_rawDescOnce sync.Once
	file_log_v1_log_proto_rawDescData = file_log_v1_log_proto_rawDesc
)

func file_log_v1_log_proto_rawDescGZIP() []byte {
	file_log_v1_log_proto_rawDescOnce.Do(func() {
		file_log_v1_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_log_v1_log_proto_rawDescData)
	})
	return file_log_v1_log_proto_rawDescData
}

var file_log_v1_log_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_log_v1_log_proto_goTypes = []interface{}{
	(*CrawlLogListRequest)(nil),  // 0: veidemann.api.log.v1.CrawlLogListRequest
	(*PageLogListRequest)(nil),   // 1: veidemann.api.log.v1.PageLogListRequest
	(*WritePageLogRequest)(nil),  // 2: veidemann.api.log.v1.WritePageLogRequest
	(*WriteCrawlLogRequest)(nil), // 3: veidemann.api.log.v1.WriteCrawlLogRequest
	(*CrawlLog)(nil),             // 4: veidemann.api.log.v1.CrawlLog
	(*v1.FieldMask)(nil),         // 5: veidemann.api.commons.v1.FieldMask
	(*PageLog)(nil),              // 6: veidemann.api.log.v1.PageLog
	(*PageLog_Resource)(nil),     // 7: veidemann.api.log.v1.PageLog.Resource
	(*emptypb.Empty)(nil),        // 8: google.protobuf.Empty
}
var file_log_v1_log_proto_depIdxs = []int32{
	4,  // 0: veidemann.api.log.v1.CrawlLogListRequest.query_template:type_name -> veidemann.api.log.v1.CrawlLog
	5,  // 1: veidemann.api.log.v1.CrawlLogListRequest.query_mask:type_name -> veidemann.api.commons.v1.FieldMask
	5,  // 2: veidemann.api.log.v1.CrawlLogListRequest.returned_fields_mask:type_name -> veidemann.api.commons.v1.FieldMask
	6,  // 3: veidemann.api.log.v1.PageLogListRequest.query_template:type_name -> veidemann.api.log.v1.PageLog
	5,  // 4: veidemann.api.log.v1.PageLogListRequest.query_mask:type_name -> veidemann.api.commons.v1.FieldMask
	5,  // 5: veidemann.api.log.v1.PageLogListRequest.returned_fields_mask:type_name -> veidemann.api.commons.v1.FieldMask
	4,  // 6: veidemann.api.log.v1.WritePageLogRequest.crawlLog:type_name -> veidemann.api.log.v1.CrawlLog
	7,  // 7: veidemann.api.log.v1.WritePageLogRequest.resource:type_name -> veidemann.api.log.v1.PageLog.Resource
	4,  // 8: veidemann.api.log.v1.WriteCrawlLogRequest.crawlLog:type_name -> veidemann.api.log.v1.CrawlLog
	0,  // 9: veidemann.api.log.v1.Log.ListCrawlLogs:input_type -> veidemann.api.log.v1.CrawlLogListRequest
	1,  // 10: veidemann.api.log.v1.Log.ListPageLogs:input_type -> veidemann.api.log.v1.PageLogListRequest
	3,  // 11: veidemann.api.log.v1.Log.WriteCrawlLog:input_type -> veidemann.api.log.v1.WriteCrawlLogRequest
	2,  // 12: veidemann.api.log.v1.Log.WritePageLog:input_type -> veidemann.api.log.v1.WritePageLogRequest
	4,  // 13: veidemann.api.log.v1.Log.ListCrawlLogs:output_type -> veidemann.api.log.v1.CrawlLog
	6,  // 14: veidemann.api.log.v1.Log.ListPageLogs:output_type -> veidemann.api.log.v1.PageLog
	8,  // 15: veidemann.api.log.v1.Log.WriteCrawlLog:output_type -> google.protobuf.Empty
	8,  // 16: veidemann.api.log.v1.Log.WritePageLog:output_type -> google.protobuf.Empty
	13, // [13:17] is the sub-list for method output_type
	9,  // [9:13] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_log_v1_log_proto_init() }
func file_log_v1_log_proto_init() {
	if File_log_v1_log_proto != nil {
		return
	}
	file_log_v1_resources_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_log_v1_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CrawlLogListRequest); i {
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
		file_log_v1_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageLogListRequest); i {
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
		file_log_v1_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WritePageLogRequest); i {
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
		file_log_v1_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteCrawlLogRequest); i {
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
	file_log_v1_log_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*WritePageLogRequest_CrawlLog)(nil),
		(*WritePageLogRequest_Resource)(nil),
		(*WritePageLogRequest_Outlink)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_log_v1_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_log_v1_log_proto_goTypes,
		DependencyIndexes: file_log_v1_log_proto_depIdxs,
		MessageInfos:      file_log_v1_log_proto_msgTypes,
	}.Build()
	File_log_v1_log_proto = out.File
	file_log_v1_log_proto_rawDesc = nil
	file_log_v1_log_proto_goTypes = nil
	file_log_v1_log_proto_depIdxs = nil
}
