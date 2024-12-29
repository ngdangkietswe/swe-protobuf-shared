// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: task/comment_service.proto

package task

import (
	common "github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type UpsertCommentReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *string                `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Content       string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	TaskId        string                 `protobuf:"bytes,3,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	ParentId      *string                `protobuf:"bytes,4,opt,name=parent_id,json=parentId,proto3,oneof" json:"parent_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertCommentReq) Reset() {
	*x = UpsertCommentReq{}
	mi := &file_task_comment_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertCommentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertCommentReq) ProtoMessage() {}

func (x *UpsertCommentReq) ProtoReflect() protoreflect.Message {
	mi := &file_task_comment_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertCommentReq.ProtoReflect.Descriptor instead.
func (*UpsertCommentReq) Descriptor() ([]byte, []int) {
	return file_task_comment_service_proto_rawDescGZIP(), []int{0}
}

func (x *UpsertCommentReq) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *UpsertCommentReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpsertCommentReq) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *UpsertCommentReq) GetParentId() string {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return ""
}

type ListCommentReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Pageable      *common.Pageable       `protobuf:"bytes,2,opt,name=pageable,proto3" json:"pageable,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCommentReq) Reset() {
	*x = ListCommentReq{}
	mi := &file_task_comment_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCommentReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentReq) ProtoMessage() {}

func (x *ListCommentReq) ProtoReflect() protoreflect.Message {
	mi := &file_task_comment_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentReq.ProtoReflect.Descriptor instead.
func (*ListCommentReq) Descriptor() ([]byte, []int) {
	return file_task_comment_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListCommentReq) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *ListCommentReq) GetPageable() *common.Pageable {
	if x != nil {
		return x.Pageable
	}
	return nil
}

type ListCommentResp struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Success bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// Types that are valid to be assigned to Resp:
	//
	//	*ListCommentResp_Data_
	//	*ListCommentResp_Error
	Resp          isListCommentResp_Resp `protobuf_oneof:"resp"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCommentResp) Reset() {
	*x = ListCommentResp{}
	mi := &file_task_comment_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCommentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentResp) ProtoMessage() {}

func (x *ListCommentResp) ProtoReflect() protoreflect.Message {
	mi := &file_task_comment_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentResp.ProtoReflect.Descriptor instead.
func (*ListCommentResp) Descriptor() ([]byte, []int) {
	return file_task_comment_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListCommentResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ListCommentResp) GetResp() isListCommentResp_Resp {
	if x != nil {
		return x.Resp
	}
	return nil
}

func (x *ListCommentResp) GetData() *ListCommentResp_Data {
	if x != nil {
		if x, ok := x.Resp.(*ListCommentResp_Data_); ok {
			return x.Data
		}
	}
	return nil
}

func (x *ListCommentResp) GetError() *common.Error {
	if x != nil {
		if x, ok := x.Resp.(*ListCommentResp_Error); ok {
			return x.Error
		}
	}
	return nil
}

type isListCommentResp_Resp interface {
	isListCommentResp_Resp()
}

type ListCommentResp_Data_ struct {
	Data *ListCommentResp_Data `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

type ListCommentResp_Error struct {
	Error *common.Error `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*ListCommentResp_Data_) isListCommentResp_Resp() {}

func (*ListCommentResp_Error) isListCommentResp_Resp() {}

type ListCommentResp_Data struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Comments      []*Comment             `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	PageMetaData  *common.PageMetaData   `protobuf:"bytes,2,opt,name=page_meta_data,json=pageMetaData,proto3" json:"page_meta_data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCommentResp_Data) Reset() {
	*x = ListCommentResp_Data{}
	mi := &file_task_comment_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCommentResp_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCommentResp_Data) ProtoMessage() {}

func (x *ListCommentResp_Data) ProtoReflect() protoreflect.Message {
	mi := &file_task_comment_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCommentResp_Data.ProtoReflect.Descriptor instead.
func (*ListCommentResp_Data) Descriptor() ([]byte, []int) {
	return file_task_comment_service_proto_rawDescGZIP(), []int{2, 0}
}

func (x *ListCommentResp_Data) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

func (x *ListCommentResp_Data) GetPageMetaData() *common.PageMetaData {
	if x != nil {
		return x.PageMetaData
	}
	return nil
}

var File_task_comment_service_proto protoreflect.FileDescriptor

var file_task_comment_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x61,
	0x73, 0x6b, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x73,
	0x65, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x13, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88,
	0x01, 0x01, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x57, 0x0a, 0x0e,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x61, 0x62, 0x6c, 0x65, 0x22, 0xfb, 0x01, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x6d, 0x0a, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x29, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x3a, 0x0a, 0x0e, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x50, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x70,
	0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x42, 0x06, 0x0a, 0x04, 0x72,
	0x65, 0x73, 0x70, 0x32, 0x92, 0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0d, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x55,
	0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x54, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x14, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x18, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x50, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x17, 0x2a, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x8c, 0x01, 0x0a, 0x08, 0x63, 0x6f, 0x6d,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x42, 0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x67, 0x64, 0x61, 0x6e, 0x67, 0x6b,
	0x69, 0x65, 0x74, 0x73, 0x77, 0x65, 0x2f, 0x73, 0x77, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2d, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa,
	0x02, 0x04, 0x54, 0x61, 0x73, 0x6b, 0xca, 0x02, 0x04, 0x54, 0x61, 0x73, 0x6b, 0xe2, 0x02, 0x10,
	0x54, 0x61, 0x73, 0x6b, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x04, 0x54, 0x61, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_task_comment_service_proto_rawDescOnce sync.Once
	file_task_comment_service_proto_rawDescData = file_task_comment_service_proto_rawDesc
)

func file_task_comment_service_proto_rawDescGZIP() []byte {
	file_task_comment_service_proto_rawDescOnce.Do(func() {
		file_task_comment_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_task_comment_service_proto_rawDescData)
	})
	return file_task_comment_service_proto_rawDescData
}

var file_task_comment_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_task_comment_service_proto_goTypes = []any{
	(*UpsertCommentReq)(nil),     // 0: task.UpsertCommentReq
	(*ListCommentReq)(nil),       // 1: task.ListCommentReq
	(*ListCommentResp)(nil),      // 2: task.ListCommentResp
	(*ListCommentResp_Data)(nil), // 3: task.ListCommentResp.Data
	(*common.Pageable)(nil),      // 4: common.Pageable
	(*common.Error)(nil),         // 5: common.Error
	(*Comment)(nil),              // 6: task.Comment
	(*common.PageMetaData)(nil),  // 7: common.PageMetaData
	(*common.IdReq)(nil),         // 8: common.IdReq
	(*common.UpsertResp)(nil),    // 9: common.UpsertResp
	(*common.EmptyResp)(nil),     // 10: common.EmptyResp
}
var file_task_comment_service_proto_depIdxs = []int32{
	4,  // 0: task.ListCommentReq.pageable:type_name -> common.Pageable
	3,  // 1: task.ListCommentResp.data:type_name -> task.ListCommentResp.Data
	5,  // 2: task.ListCommentResp.error:type_name -> common.Error
	6,  // 3: task.ListCommentResp.Data.comments:type_name -> task.Comment
	7,  // 4: task.ListCommentResp.Data.page_meta_data:type_name -> common.PageMetaData
	0,  // 5: task.CommentService.UpsertComment:input_type -> task.UpsertCommentReq
	1,  // 6: task.CommentService.ListComment:input_type -> task.ListCommentReq
	8,  // 7: task.CommentService.DeleteComment:input_type -> common.IdReq
	9,  // 8: task.CommentService.UpsertComment:output_type -> common.UpsertResp
	2,  // 9: task.CommentService.ListComment:output_type -> task.ListCommentResp
	10, // 10: task.CommentService.DeleteComment:output_type -> common.EmptyResp
	8,  // [8:11] is the sub-list for method output_type
	5,  // [5:8] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_task_comment_service_proto_init() }
func file_task_comment_service_proto_init() {
	if File_task_comment_service_proto != nil {
		return
	}
	file_task_shared_proto_init()
	file_task_comment_service_proto_msgTypes[0].OneofWrappers = []any{}
	file_task_comment_service_proto_msgTypes[2].OneofWrappers = []any{
		(*ListCommentResp_Data_)(nil),
		(*ListCommentResp_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_task_comment_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_task_comment_service_proto_goTypes,
		DependencyIndexes: file_task_comment_service_proto_depIdxs,
		MessageInfos:      file_task_comment_service_proto_msgTypes,
	}.Build()
	File_task_comment_service_proto = out.File
	file_task_comment_service_proto_rawDesc = nil
	file_task_comment_service_proto_goTypes = nil
	file_task_comment_service_proto_depIdxs = nil
}