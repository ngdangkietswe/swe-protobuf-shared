// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: auth/permission_internal_service.proto

package auth

import (
	common "github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PermissionOfUserResp struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Success bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// Types that are valid to be assigned to Resp:
	//
	//	*PermissionOfUserResp_Data_
	//	*PermissionOfUserResp_Error
	Resp          isPermissionOfUserResp_Resp `protobuf_oneof:"resp"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PermissionOfUserResp) Reset() {
	*x = PermissionOfUserResp{}
	mi := &file_auth_permission_internal_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PermissionOfUserResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionOfUserResp) ProtoMessage() {}

func (x *PermissionOfUserResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_permission_internal_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionOfUserResp.ProtoReflect.Descriptor instead.
func (*PermissionOfUserResp) Descriptor() ([]byte, []int) {
	return file_auth_permission_internal_service_proto_rawDescGZIP(), []int{0}
}

func (x *PermissionOfUserResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *PermissionOfUserResp) GetResp() isPermissionOfUserResp_Resp {
	if x != nil {
		return x.Resp
	}
	return nil
}

func (x *PermissionOfUserResp) GetData() *PermissionOfUserResp_Data {
	if x != nil {
		if x, ok := x.Resp.(*PermissionOfUserResp_Data_); ok {
			return x.Data
		}
	}
	return nil
}

func (x *PermissionOfUserResp) GetError() *common.Error {
	if x != nil {
		if x, ok := x.Resp.(*PermissionOfUserResp_Error); ok {
			return x.Error
		}
	}
	return nil
}

type isPermissionOfUserResp_Resp interface {
	isPermissionOfUserResp_Resp()
}

type PermissionOfUserResp_Data_ struct {
	Data *PermissionOfUserResp_Data `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

type PermissionOfUserResp_Error struct {
	Error *common.Error `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*PermissionOfUserResp_Data_) isPermissionOfUserResp_Resp() {}

func (*PermissionOfUserResp_Error) isPermissionOfUserResp_Resp() {}

type PermissionOfUserResp_Data struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Permissions   []*Permission          `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PermissionOfUserResp_Data) Reset() {
	*x = PermissionOfUserResp_Data{}
	mi := &file_auth_permission_internal_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PermissionOfUserResp_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionOfUserResp_Data) ProtoMessage() {}

func (x *PermissionOfUserResp_Data) ProtoReflect() protoreflect.Message {
	mi := &file_auth_permission_internal_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionOfUserResp_Data.ProtoReflect.Descriptor instead.
func (*PermissionOfUserResp_Data) Descriptor() ([]byte, []int) {
	return file_auth_permission_internal_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PermissionOfUserResp_Data) GetPermissions() []*Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

var File_auth_permission_internal_service_proto protoreflect.FileDescriptor

var file_auth_permission_internal_service_proto_rawDesc = string([]byte{
	0x0a, 0x26, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x13,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x01, 0x0a, 0x14, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x66, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x35, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x66, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x25, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x3a, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x32, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x70, 0x32, 0x5a, 0x0a, 0x19, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x10, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x66, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4f, 0x66, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x42, 0x97, 0x01, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x42, 0x1e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6e, 0x67, 0x64, 0x61, 0x6e, 0x67, 0x6b, 0x69, 0x65, 0x74, 0x73, 0x77, 0x65,
	0x2f, 0x73, 0x77, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x41, 0x75, 0x74, 0x68,
	0xca, 0x02, 0x04, 0x41, 0x75, 0x74, 0x68, 0xe2, 0x02, 0x10, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x41, 0x75, 0x74,
	0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_auth_permission_internal_service_proto_rawDescOnce sync.Once
	file_auth_permission_internal_service_proto_rawDescData []byte
)

func file_auth_permission_internal_service_proto_rawDescGZIP() []byte {
	file_auth_permission_internal_service_proto_rawDescOnce.Do(func() {
		file_auth_permission_internal_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_auth_permission_internal_service_proto_rawDesc), len(file_auth_permission_internal_service_proto_rawDesc)))
	})
	return file_auth_permission_internal_service_proto_rawDescData
}

var file_auth_permission_internal_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_auth_permission_internal_service_proto_goTypes = []any{
	(*PermissionOfUserResp)(nil),      // 0: auth.PermissionOfUserResp
	(*PermissionOfUserResp_Data)(nil), // 1: auth.PermissionOfUserResp.Data
	(*common.Error)(nil),              // 2: common.Error
	(*Permission)(nil),                // 3: auth.Permission
	(*common.IdReq)(nil),              // 4: common.IdReq
}
var file_auth_permission_internal_service_proto_depIdxs = []int32{
	1, // 0: auth.PermissionOfUserResp.data:type_name -> auth.PermissionOfUserResp.Data
	2, // 1: auth.PermissionOfUserResp.error:type_name -> common.Error
	3, // 2: auth.PermissionOfUserResp.Data.permissions:type_name -> auth.Permission
	4, // 3: auth.PermissionInternalService.PermissionOfUser:input_type -> common.IdReq
	0, // 4: auth.PermissionInternalService.PermissionOfUser:output_type -> auth.PermissionOfUserResp
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_auth_permission_internal_service_proto_init() }
func file_auth_permission_internal_service_proto_init() {
	if File_auth_permission_internal_service_proto != nil {
		return
	}
	file_auth_shared_proto_init()
	file_auth_permission_internal_service_proto_msgTypes[0].OneofWrappers = []any{
		(*PermissionOfUserResp_Data_)(nil),
		(*PermissionOfUserResp_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_auth_permission_internal_service_proto_rawDesc), len(file_auth_permission_internal_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_permission_internal_service_proto_goTypes,
		DependencyIndexes: file_auth_permission_internal_service_proto_depIdxs,
		MessageInfos:      file_auth_permission_internal_service_proto_msgTypes,
	}.Build()
	File_auth_permission_internal_service_proto = out.File
	file_auth_permission_internal_service_proto_goTypes = nil
	file_auth_permission_internal_service_proto_depIdxs = nil
}
