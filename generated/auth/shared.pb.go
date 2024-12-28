// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        (unknown)
// source: auth/shared.proto

package auth

import (
	common "github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
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

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            *string                `protobuf:"bytes,1,opt,name=id,proto3,oneof" json:"id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Email         string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_auth_shared_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_auth_shared_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_auth_shared_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type LoginReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password      string                 `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Otp           *string                `protobuf:"bytes,3,opt,name=otp,proto3,oneof" json:"otp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	mi := &file_auth_shared_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_shared_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_auth_shared_proto_rawDescGZIP(), []int{1}
}

func (x *LoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginReq) GetOtp() string {
	if x != nil && x.Otp != nil {
		return *x.Otp
	}
	return ""
}

type LoginResp struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Success bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// Types that are valid to be assigned to Resp:
	//
	//	*LoginResp_Data_
	//	*LoginResp_Error
	Resp          isLoginResp_Resp `protobuf_oneof:"resp"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoginResp) Reset() {
	*x = LoginResp{}
	mi := &file_auth_shared_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp) ProtoMessage() {}

func (x *LoginResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_shared_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp.ProtoReflect.Descriptor instead.
func (*LoginResp) Descriptor() ([]byte, []int) {
	return file_auth_shared_proto_rawDescGZIP(), []int{2}
}

func (x *LoginResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *LoginResp) GetResp() isLoginResp_Resp {
	if x != nil {
		return x.Resp
	}
	return nil
}

func (x *LoginResp) GetData() *LoginResp_Data {
	if x != nil {
		if x, ok := x.Resp.(*LoginResp_Data_); ok {
			return x.Data
		}
	}
	return nil
}

func (x *LoginResp) GetError() *common.Error {
	if x != nil {
		if x, ok := x.Resp.(*LoginResp_Error); ok {
			return x.Error
		}
	}
	return nil
}

type isLoginResp_Resp interface {
	isLoginResp_Resp()
}

type LoginResp_Data_ struct {
	Data *LoginResp_Data `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

type LoginResp_Error struct {
	Error *common.Error `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*LoginResp_Data_) isLoginResp_Resp() {}

func (*LoginResp_Error) isLoginResp_Resp() {}

// TWO FACTOR AUTHENTICATION
type EnableOrDisable2FAReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Enable        bool                   `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EnableOrDisable2FAReq) Reset() {
	*x = EnableOrDisable2FAReq{}
	mi := &file_auth_shared_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnableOrDisable2FAReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableOrDisable2FAReq) ProtoMessage() {}

func (x *EnableOrDisable2FAReq) ProtoReflect() protoreflect.Message {
	mi := &file_auth_shared_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableOrDisable2FAReq.ProtoReflect.Descriptor instead.
func (*EnableOrDisable2FAReq) Descriptor() ([]byte, []int) {
	return file_auth_shared_proto_rawDescGZIP(), []int{3}
}

func (x *EnableOrDisable2FAReq) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

type EnableOrDisable2FAResp struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Success        bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error          *common.Error          `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"`
	QrCodeImageUrl *string                `protobuf:"bytes,3,opt,name=qr_code_image_url,json=qrCodeImageUrl,proto3,oneof" json:"qr_code_image_url,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *EnableOrDisable2FAResp) Reset() {
	*x = EnableOrDisable2FAResp{}
	mi := &file_auth_shared_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnableOrDisable2FAResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnableOrDisable2FAResp) ProtoMessage() {}

func (x *EnableOrDisable2FAResp) ProtoReflect() protoreflect.Message {
	mi := &file_auth_shared_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnableOrDisable2FAResp.ProtoReflect.Descriptor instead.
func (*EnableOrDisable2FAResp) Descriptor() ([]byte, []int) {
	return file_auth_shared_proto_rawDescGZIP(), []int{4}
}

func (x *EnableOrDisable2FAResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *EnableOrDisable2FAResp) GetError() *common.Error {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *EnableOrDisable2FAResp) GetQrCodeImageUrl() string {
	if x != nil && x.QrCodeImageUrl != nil {
		return *x.QrCodeImageUrl
	}
	return ""
}

// PERMISSION
type Action struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Action) Reset() {
	*x = Action{}
	mi := &file_auth_shared_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_auth_shared_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_auth_shared_proto_rawDescGZIP(), []int{5}
}

func (x *Action) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Action) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Action) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Resource struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Resource) Reset() {
	*x = Resource{}
	mi := &file_auth_shared_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_auth_shared_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_auth_shared_proto_rawDescGZIP(), []int{6}
}

func (x *Resource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Resource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Resource) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type Permission struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Action        *Action                `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	Resource      *Resource              `protobuf:"bytes,4,opt,name=resource,proto3" json:"resource,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Permission) Reset() {
	*x = Permission{}
	mi := &file_auth_shared_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_auth_shared_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_auth_shared_proto_rawDescGZIP(), []int{7}
}

func (x *Permission) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Permission) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Permission) GetAction() *Action {
	if x != nil {
		return x.Action
	}
	return nil
}

func (x *Permission) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

type LoginResp_Data struct {
	state                 protoimpl.MessageState `protogen:"open.v1"`
	TokenType             string                 `protobuf:"bytes,1,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	AccessToken           string                 `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	AccessTokenExpiresIn  string                 `protobuf:"bytes,3,opt,name=access_token_expires_in,json=accessTokenExpiresIn,proto3" json:"access_token_expires_in,omitempty"`
	RefreshToken          string                 `protobuf:"bytes,4,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	RefreshTokenExpiresIn string                 `protobuf:"bytes,5,opt,name=refresh_token_expires_in,json=refreshTokenExpiresIn,proto3" json:"refresh_token_expires_in,omitempty"`
	TwoFactorAuth         bool                   `protobuf:"varint,6,opt,name=two_factor_auth,json=twoFactorAuth,proto3" json:"two_factor_auth,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *LoginResp_Data) Reset() {
	*x = LoginResp_Data{}
	mi := &file_auth_shared_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResp_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResp_Data) ProtoMessage() {}

func (x *LoginResp_Data) ProtoReflect() protoreflect.Message {
	mi := &file_auth_shared_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResp_Data.ProtoReflect.Descriptor instead.
func (*LoginResp_Data) Descriptor() ([]byte, []int) {
	return file_auth_shared_proto_rawDescGZIP(), []int{2, 0}
}

func (x *LoginResp_Data) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *LoginResp_Data) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginResp_Data) GetAccessTokenExpiresIn() string {
	if x != nil {
		return x.AccessTokenExpiresIn
	}
	return ""
}

func (x *LoginResp_Data) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *LoginResp_Data) GetRefreshTokenExpiresIn() string {
	if x != nil {
		return x.RefreshTokenExpiresIn
	}
	return ""
}

func (x *LoginResp_Data) GetTwoFactorAuth() bool {
	if x != nil {
		return x.TwoFactorAuth
	}
	return false
}

var File_auth_shared_proto protoreflect.FileDescriptor

var file_auth_shared_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64,
	0x22, 0x61, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x15, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x03, 0x6f, 0x74, 0x70, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f,
	0x6f, 0x74, 0x70, 0x22, 0x88, 0x03, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2a, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x48,
	0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x85,
	0x02, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x35, 0x0a, 0x17, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x73, 0x5f, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x37, 0x0a, 0x18, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x69,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x12, 0x26,
	0x0a, 0x0f, 0x74, 0x77, 0x6f, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x75, 0x74,
	0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x74, 0x77, 0x6f, 0x46, 0x61, 0x63, 0x74,
	0x6f, 0x72, 0x41, 0x75, 0x74, 0x68, 0x42, 0x06, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x70, 0x22, 0x2f,
	0x0a, 0x15, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x72, 0x44, 0x69, 0x73, 0x61, 0x62, 0x6c,
	0x65, 0x32, 0x46, 0x41, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x22,
	0xac, 0x01, 0x0a, 0x16, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x72, 0x44, 0x69, 0x73, 0x61,
	0x62, 0x6c, 0x65, 0x32, 0x46, 0x41, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x28, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2e,
	0x0a, 0x11, 0x71, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0e, 0x71, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x71, 0x72, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0x4e,
	0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x50,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x90, 0x01, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x24, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x42, 0x84, 0x01, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x42, 0x0b, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x67, 0x64, 0x61,
	0x6e, 0x67, 0x6b, 0x69, 0x65, 0x74, 0x73, 0x77, 0x65, 0x2f, 0x73, 0x77, 0x65, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x61, 0x75, 0x74, 0x68, 0xa2, 0x02, 0x03, 0x41,
	0x58, 0x58, 0xaa, 0x02, 0x04, 0x41, 0x75, 0x74, 0x68, 0xca, 0x02, 0x04, 0x41, 0x75, 0x74, 0x68,
	0xe2, 0x02, 0x10, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x41, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_auth_shared_proto_rawDescOnce sync.Once
	file_auth_shared_proto_rawDescData = file_auth_shared_proto_rawDesc
)

func file_auth_shared_proto_rawDescGZIP() []byte {
	file_auth_shared_proto_rawDescOnce.Do(func() {
		file_auth_shared_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_shared_proto_rawDescData)
	})
	return file_auth_shared_proto_rawDescData
}

var file_auth_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_auth_shared_proto_goTypes = []any{
	(*User)(nil),                   // 0: auth.User
	(*LoginReq)(nil),               // 1: auth.LoginReq
	(*LoginResp)(nil),              // 2: auth.LoginResp
	(*EnableOrDisable2FAReq)(nil),  // 3: auth.EnableOrDisable2FAReq
	(*EnableOrDisable2FAResp)(nil), // 4: auth.EnableOrDisable2FAResp
	(*Action)(nil),                 // 5: auth.Action
	(*Resource)(nil),               // 6: auth.Resource
	(*Permission)(nil),             // 7: auth.Permission
	(*LoginResp_Data)(nil),         // 8: auth.LoginResp.Data
	(*common.Error)(nil),           // 9: common.Error
}
var file_auth_shared_proto_depIdxs = []int32{
	8, // 0: auth.LoginResp.data:type_name -> auth.LoginResp.Data
	9, // 1: auth.LoginResp.error:type_name -> common.Error
	9, // 2: auth.EnableOrDisable2FAResp.error:type_name -> common.Error
	5, // 3: auth.Permission.action:type_name -> auth.Action
	6, // 4: auth.Permission.resource:type_name -> auth.Resource
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_auth_shared_proto_init() }
func file_auth_shared_proto_init() {
	if File_auth_shared_proto != nil {
		return
	}
	file_auth_shared_proto_msgTypes[0].OneofWrappers = []any{}
	file_auth_shared_proto_msgTypes[1].OneofWrappers = []any{}
	file_auth_shared_proto_msgTypes[2].OneofWrappers = []any{
		(*LoginResp_Data_)(nil),
		(*LoginResp_Error)(nil),
	}
	file_auth_shared_proto_msgTypes[4].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_shared_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_shared_proto_goTypes,
		DependencyIndexes: file_auth_shared_proto_depIdxs,
		MessageInfos:      file_auth_shared_proto_msgTypes,
	}.Build()
	File_auth_shared_proto = out.File
	file_auth_shared_proto_rawDesc = nil
	file_auth_shared_proto_goTypes = nil
	file_auth_shared_proto_depIdxs = nil
}
