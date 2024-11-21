// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: common/shared.proto

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

type Pageable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page      int32   `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size      int32   `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Sort      string  `protobuf:"bytes,3,opt,name=sort,proto3" json:"sort,omitempty"`
	Direction string  `protobuf:"bytes,4,opt,name=direction,proto3" json:"direction,omitempty"`
	UnPaged   bool    `protobuf:"varint,5,opt,name=un_paged,json=unPaged,proto3" json:"un_paged,omitempty"`
	Search    *string `protobuf:"bytes,6,opt,name=search,proto3,oneof" json:"search,omitempty"`
}

func (x *Pageable) Reset() {
	*x = Pageable{}
	mi := &file_common_shared_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pageable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pageable) ProtoMessage() {}

func (x *Pageable) ProtoReflect() protoreflect.Message {
	mi := &file_common_shared_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pageable.ProtoReflect.Descriptor instead.
func (*Pageable) Descriptor() ([]byte, []int) {
	return file_common_shared_proto_rawDescGZIP(), []int{0}
}

func (x *Pageable) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Pageable) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Pageable) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *Pageable) GetDirection() string {
	if x != nil {
		return x.Direction
	}
	return ""
}

func (x *Pageable) GetUnPaged() bool {
	if x != nil {
		return x.UnPaged
	}
	return false
}

func (x *Pageable) GetSearch() string {
	if x != nil && x.Search != nil {
		return *x.Search
	}
	return ""
}

type PageMetaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page          int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size          int32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	TotalPages    int64 `protobuf:"varint,3,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
	TotalElements int64 `protobuf:"varint,4,opt,name=total_elements,json=totalElements,proto3" json:"total_elements,omitempty"`
	Next          bool  `protobuf:"varint,5,opt,name=next,proto3" json:"next,omitempty"`
	Previous      bool  `protobuf:"varint,6,opt,name=previous,proto3" json:"previous,omitempty"`
}

func (x *PageMetaData) Reset() {
	*x = PageMetaData{}
	mi := &file_common_shared_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PageMetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageMetaData) ProtoMessage() {}

func (x *PageMetaData) ProtoReflect() protoreflect.Message {
	mi := &file_common_shared_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageMetaData.ProtoReflect.Descriptor instead.
func (*PageMetaData) Descriptor() ([]byte, []int) {
	return file_common_shared_proto_rawDescGZIP(), []int{1}
}

func (x *PageMetaData) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageMetaData) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *PageMetaData) GetTotalPages() int64 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

func (x *PageMetaData) GetTotalElements() int64 {
	if x != nil {
		return x.TotalElements
	}
	return 0
}

func (x *PageMetaData) GetNext() bool {
	if x != nil {
		return x.Next
	}
	return false
}

func (x *PageMetaData) GetPrevious() bool {
	if x != nil {
		return x.Previous
	}
	return false
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string            `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Code    int32             `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string            `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Details map[string]string `protobuf:"bytes,4,rep,name=details,proto3" json:"details,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Error) Reset() {
	*x = Error{}
	mi := &file_common_shared_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_common_shared_proto_msgTypes[2]
	if x != nil {
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
	return file_common_shared_proto_rawDescGZIP(), []int{2}
}

func (x *Error) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetDetails() map[string]string {
	if x != nil {
		return x.Details
	}
	return nil
}

// COMMON REQUEST MESSAGES
type IdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdReq) Reset() {
	*x = IdReq{}
	mi := &file_common_shared_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdReq) ProtoMessage() {}

func (x *IdReq) ProtoReflect() protoreflect.Message {
	mi := &file_common_shared_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdReq.ProtoReflect.Descriptor instead.
func (*IdReq) Descriptor() ([]byte, []int) {
	return file_common_shared_proto_rawDescGZIP(), []int{3}
}

func (x *IdReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// COMMON RESPONSE MESSAGES
type UpsertResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// Types that are assignable to Resp:
	//
	//	*UpsertResp_Data_
	//	*UpsertResp_Error
	Resp isUpsertResp_Resp `protobuf_oneof:"resp"`
}

func (x *UpsertResp) Reset() {
	*x = UpsertResp{}
	mi := &file_common_shared_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertResp) ProtoMessage() {}

func (x *UpsertResp) ProtoReflect() protoreflect.Message {
	mi := &file_common_shared_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertResp.ProtoReflect.Descriptor instead.
func (*UpsertResp) Descriptor() ([]byte, []int) {
	return file_common_shared_proto_rawDescGZIP(), []int{4}
}

func (x *UpsertResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (m *UpsertResp) GetResp() isUpsertResp_Resp {
	if m != nil {
		return m.Resp
	}
	return nil
}

func (x *UpsertResp) GetData() *UpsertResp_Data {
	if x, ok := x.GetResp().(*UpsertResp_Data_); ok {
		return x.Data
	}
	return nil
}

func (x *UpsertResp) GetError() *Error {
	if x, ok := x.GetResp().(*UpsertResp_Error); ok {
		return x.Error
	}
	return nil
}

type isUpsertResp_Resp interface {
	isUpsertResp_Resp()
}

type UpsertResp_Data_ struct {
	Data *UpsertResp_Data `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

type UpsertResp_Error struct {
	Error *Error `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*UpsertResp_Data_) isUpsertResp_Resp() {}

func (*UpsertResp_Error) isUpsertResp_Resp() {}

type EmptyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error   *Error `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *EmptyResp) Reset() {
	*x = EmptyResp{}
	mi := &file_common_shared_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EmptyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResp) ProtoMessage() {}

func (x *EmptyResp) ProtoReflect() protoreflect.Message {
	mi := &file_common_shared_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResp.ProtoReflect.Descriptor instead.
func (*EmptyResp) Descriptor() ([]byte, []int) {
	return file_common_shared_proto_rawDescGZIP(), []int{5}
}

func (x *EmptyResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *EmptyResp) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type UpsertResp_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpsertResp_Data) Reset() {
	*x = UpsertResp_Data{}
	mi := &file_common_shared_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertResp_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertResp_Data) ProtoMessage() {}

func (x *UpsertResp_Data) ProtoReflect() protoreflect.Message {
	mi := &file_common_shared_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertResp_Data.ProtoReflect.Descriptor instead.
func (*UpsertResp_Data) Descriptor() ([]byte, []int) {
	return file_common_shared_proto_rawDescGZIP(), []int{4, 0}
}

func (x *UpsertResp_Data) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_common_shared_proto protoreflect.FileDescriptor

var file_common_shared_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0xa7, 0x01,
	0x0a, 0x08, 0x50, 0x61, 0x67, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x6e, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x75, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x64, 0x12,
	0x1b, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x00, 0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07,
	0x5f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0xae, 0x01, 0x0a, 0x0c, 0x50, 0x61, 0x67, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x22, 0xbf, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x3a,
	0x0a, 0x0c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x17, 0x0a, 0x05, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x9c, 0x01, 0x0a, 0x0a, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2d, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x44,
	0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x1a, 0x16, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x72, 0x65,
	0x73, 0x70, 0x22, 0x59, 0x0a, 0x09, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x28, 0x0a, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x90, 0x01,
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x0b, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x67, 0x64, 0x61, 0x6e, 0x67, 0x6b, 0x69,
	0x65, 0x74, 0x73, 0x77, 0x65, 0x2f, 0x73, 0x77, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2d, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58,
	0xaa, 0x02, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca, 0x02, 0x06, 0x43, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0xe2, 0x02, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_shared_proto_rawDescOnce sync.Once
	file_common_shared_proto_rawDescData = file_common_shared_proto_rawDesc
)

func file_common_shared_proto_rawDescGZIP() []byte {
	file_common_shared_proto_rawDescOnce.Do(func() {
		file_common_shared_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_shared_proto_rawDescData)
	})
	return file_common_shared_proto_rawDescData
}

var file_common_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_common_shared_proto_goTypes = []any{
	(*Pageable)(nil),        // 0: common.Pageable
	(*PageMetaData)(nil),    // 1: common.PageMetaData
	(*Error)(nil),           // 2: common.Error
	(*IdReq)(nil),           // 3: common.IdReq
	(*UpsertResp)(nil),      // 4: common.UpsertResp
	(*EmptyResp)(nil),       // 5: common.EmptyResp
	nil,                     // 6: common.Error.DetailsEntry
	(*UpsertResp_Data)(nil), // 7: common.UpsertResp.Data
}
var file_common_shared_proto_depIdxs = []int32{
	6, // 0: common.Error.details:type_name -> common.Error.DetailsEntry
	7, // 1: common.UpsertResp.data:type_name -> common.UpsertResp.Data
	2, // 2: common.UpsertResp.error:type_name -> common.Error
	2, // 3: common.EmptyResp.error:type_name -> common.Error
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_common_shared_proto_init() }
func file_common_shared_proto_init() {
	if File_common_shared_proto != nil {
		return
	}
	file_common_shared_proto_msgTypes[0].OneofWrappers = []any{}
	file_common_shared_proto_msgTypes[4].OneofWrappers = []any{
		(*UpsertResp_Data_)(nil),
		(*UpsertResp_Error)(nil),
	}
	file_common_shared_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_shared_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_shared_proto_goTypes,
		DependencyIndexes: file_common_shared_proto_depIdxs,
		MessageInfos:      file_common_shared_proto_msgTypes,
	}.Build()
	File_common_shared_proto = out.File
	file_common_shared_proto_rawDesc = nil
	file_common_shared_proto_goTypes = nil
	file_common_shared_proto_depIdxs = nil
}
