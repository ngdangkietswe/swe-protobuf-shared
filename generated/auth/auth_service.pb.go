// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: auth/auth_service.proto

package auth

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

type ChangePasswordReq struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	OldPassword     string                 `protobuf:"bytes,1,opt,name=old_password,json=oldPassword,proto3" json:"old_password,omitempty"`
	NewPassword     string                 `protobuf:"bytes,2,opt,name=new_password,json=newPassword,proto3" json:"new_password,omitempty"`
	ConfirmPassword string                 `protobuf:"bytes,3,opt,name=confirm_password,json=confirmPassword,proto3" json:"confirm_password,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ChangePasswordReq) Reset() {
	*x = ChangePasswordReq{}
	mi := &file_auth_auth_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChangePasswordReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordReq) ProtoMessage() {}

func (x *ChangePasswordReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_auth_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordReq.ProtoReflect.Descriptor instead.
func (*ChangePasswordReq) Descriptor() ([]byte, []int) {
	return file_auth_auth_service_proto_rawDescGZIP(), []int{0}
}

func (x *ChangePasswordReq) GetOldPassword() string {
	if x != nil {
		return x.OldPassword
	}
	return ""
}

func (x *ChangePasswordReq) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

func (x *ChangePasswordReq) GetConfirmPassword() string {
	if x != nil {
		return x.ConfirmPassword
	}
	return ""
}

var File_auth_auth_service_proto protoreflect.FileDescriptor

var file_auth_auth_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x0c, 0x6f,
	0x6c, 0x64, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x5f, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x72, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x32, 0xfd, 0x02, 0x0a,
	0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0c,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0a, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x20, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x47,
	0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17,
	0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x6c, 0x0a, 0x12, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x4f, 0x72, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x32, 0x46, 0x41, 0x12, 0x1b, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x72, 0x44, 0x69, 0x73,
	0x61, 0x62, 0x6c, 0x65, 0x32, 0x46, 0x41, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x72, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x32, 0x46, 0x41, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15,
	0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2f, 0x32, 0x66, 0x61, 0x12, 0x65, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x2d, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x89, 0x01, 0x0a,
	0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x42, 0x10, 0x41, 0x75, 0x74, 0x68, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x67, 0x64, 0x61, 0x6e, 0x67,
	0x6b, 0x69, 0x65, 0x74, 0x73, 0x77, 0x65, 0x2f, 0x73, 0x77, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2d, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x68, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58,
	0xaa, 0x02, 0x04, 0x41, 0x75, 0x74, 0x68, 0xca, 0x02, 0x04, 0x41, 0x75, 0x74, 0x68, 0xe2, 0x02,
	0x10, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x04, 0x41, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_auth_service_proto_rawDescOnce sync.Once
	file_auth_auth_service_proto_rawDescData = file_auth_auth_service_proto_rawDesc
)

func file_auth_auth_service_proto_rawDescGZIP() []byte {
	file_auth_auth_service_proto_rawDescOnce.Do(func() {
		file_auth_auth_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_auth_service_proto_rawDescData)
	})
	return file_auth_auth_service_proto_rawDescData
}

var file_auth_auth_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_auth_auth_service_proto_goTypes = []any{
	(*ChangePasswordReq)(nil),      // 0: auth.ChangePasswordReq
	(*User)(nil),                   // 1: auth.User
	(*LoginReq)(nil),               // 2: auth.LoginReq
	(*EnableOrDisable2FAReq)(nil),  // 3: auth.EnableOrDisable2FAReq
	(*common.UpsertResp)(nil),      // 4: common.UpsertResp
	(*LoginResp)(nil),              // 5: auth.LoginResp
	(*EnableOrDisable2FAResp)(nil), // 6: auth.EnableOrDisable2FAResp
	(*common.EmptyResp)(nil),       // 7: common.EmptyResp
}
var file_auth_auth_service_proto_depIdxs = []int32{
	1, // 0: auth.AuthService.RegisterUser:input_type -> auth.User
	2, // 1: auth.AuthService.Login:input_type -> auth.LoginReq
	3, // 2: auth.AuthService.EnableOrDisable2FA:input_type -> auth.EnableOrDisable2FAReq
	0, // 3: auth.AuthService.ChangePassword:input_type -> auth.ChangePasswordReq
	4, // 4: auth.AuthService.RegisterUser:output_type -> common.UpsertResp
	5, // 5: auth.AuthService.Login:output_type -> auth.LoginResp
	6, // 6: auth.AuthService.EnableOrDisable2FA:output_type -> auth.EnableOrDisable2FAResp
	7, // 7: auth.AuthService.ChangePassword:output_type -> common.EmptyResp
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_auth_auth_service_proto_init() }
func file_auth_auth_service_proto_init() {
	if File_auth_auth_service_proto != nil {
		return
	}
	file_auth_shared_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_auth_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_auth_service_proto_goTypes,
		DependencyIndexes: file_auth_auth_service_proto_depIdxs,
		MessageInfos:      file_auth_auth_service_proto_msgTypes,
	}.Build()
	File_auth_auth_service_proto = out.File
	file_auth_auth_service_proto_rawDesc = nil
	file_auth_auth_service_proto_goTypes = nil
	file_auth_auth_service_proto_depIdxs = nil
}
