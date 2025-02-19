// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: integration/strava_service.proto

package integration

import (
	common "github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type IntegrateStravaAccountReq struct {
	state         protoimpl.MessageState            `protogen:"open.v1"`
	UserId        string                            `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Strava        *IntegrateStravaAccountReq_Strava `protobuf:"bytes,2,opt,name=strava,proto3" json:"strava,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IntegrateStravaAccountReq) Reset() {
	*x = IntegrateStravaAccountReq{}
	mi := &file_integration_strava_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IntegrateStravaAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegrateStravaAccountReq) ProtoMessage() {}

func (x *IntegrateStravaAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_integration_strava_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegrateStravaAccountReq.ProtoReflect.Descriptor instead.
func (*IntegrateStravaAccountReq) Descriptor() ([]byte, []int) {
	return file_integration_strava_service_proto_rawDescGZIP(), []int{0}
}

func (x *IntegrateStravaAccountReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *IntegrateStravaAccountReq) GetStrava() *IntegrateStravaAccountReq_Strava {
	if x != nil {
		return x.Strava
	}
	return nil
}

type GetStravaAccountReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStravaAccountReq) Reset() {
	*x = GetStravaAccountReq{}
	mi := &file_integration_strava_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStravaAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStravaAccountReq) ProtoMessage() {}

func (x *GetStravaAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_integration_strava_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStravaAccountReq.ProtoReflect.Descriptor instead.
func (*GetStravaAccountReq) Descriptor() ([]byte, []int) {
	return file_integration_strava_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetStravaAccountReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetStravaAccountResp struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Success bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// Types that are valid to be assigned to Resp:
	//
	//	*GetStravaAccountResp_StravaAccount
	//	*GetStravaAccountResp_Error
	Resp          isGetStravaAccountResp_Resp `protobuf_oneof:"resp"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStravaAccountResp) Reset() {
	*x = GetStravaAccountResp{}
	mi := &file_integration_strava_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStravaAccountResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStravaAccountResp) ProtoMessage() {}

func (x *GetStravaAccountResp) ProtoReflect() protoreflect.Message {
	mi := &file_integration_strava_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStravaAccountResp.ProtoReflect.Descriptor instead.
func (*GetStravaAccountResp) Descriptor() ([]byte, []int) {
	return file_integration_strava_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetStravaAccountResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetStravaAccountResp) GetResp() isGetStravaAccountResp_Resp {
	if x != nil {
		return x.Resp
	}
	return nil
}

func (x *GetStravaAccountResp) GetStravaAccount() *StravaAccount {
	if x != nil {
		if x, ok := x.Resp.(*GetStravaAccountResp_StravaAccount); ok {
			return x.StravaAccount
		}
	}
	return nil
}

func (x *GetStravaAccountResp) GetError() *common.Error {
	if x != nil {
		if x, ok := x.Resp.(*GetStravaAccountResp_Error); ok {
			return x.Error
		}
	}
	return nil
}

type isGetStravaAccountResp_Resp interface {
	isGetStravaAccountResp_Resp()
}

type GetStravaAccountResp_StravaAccount struct {
	StravaAccount *StravaAccount `protobuf:"bytes,2,opt,name=strava_account,json=stravaAccount,proto3,oneof"`
}

type GetStravaAccountResp_Error struct {
	Error *common.Error `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*GetStravaAccountResp_StravaAccount) isGetStravaAccountResp_Resp() {}

func (*GetStravaAccountResp_Error) isGetStravaAccountResp_Resp() {}

type GetStravaActivitiesReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          *StravaActivityType    `protobuf:"varint,1,opt,name=type,proto3,enum=integration.StravaActivityType,oneof" json:"type,omitempty"`
	Pageable      *common.Pageable       `protobuf:"bytes,2,opt,name=pageable,proto3" json:"pageable,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStravaActivitiesReq) Reset() {
	*x = GetStravaActivitiesReq{}
	mi := &file_integration_strava_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStravaActivitiesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStravaActivitiesReq) ProtoMessage() {}

func (x *GetStravaActivitiesReq) ProtoReflect() protoreflect.Message {
	mi := &file_integration_strava_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStravaActivitiesReq.ProtoReflect.Descriptor instead.
func (*GetStravaActivitiesReq) Descriptor() ([]byte, []int) {
	return file_integration_strava_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetStravaActivitiesReq) GetType() StravaActivityType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return StravaActivityType_RUN
}

func (x *GetStravaActivitiesReq) GetPageable() *common.Pageable {
	if x != nil {
		return x.Pageable
	}
	return nil
}

type GetStravaActivitiesResp struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Success bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	// Types that are valid to be assigned to Resp:
	//
	//	*GetStravaActivitiesResp_Data_
	//	*GetStravaActivitiesResp_Error
	Resp          isGetStravaActivitiesResp_Resp `protobuf_oneof:"resp"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStravaActivitiesResp) Reset() {
	*x = GetStravaActivitiesResp{}
	mi := &file_integration_strava_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStravaActivitiesResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStravaActivitiesResp) ProtoMessage() {}

func (x *GetStravaActivitiesResp) ProtoReflect() protoreflect.Message {
	mi := &file_integration_strava_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStravaActivitiesResp.ProtoReflect.Descriptor instead.
func (*GetStravaActivitiesResp) Descriptor() ([]byte, []int) {
	return file_integration_strava_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetStravaActivitiesResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetStravaActivitiesResp) GetResp() isGetStravaActivitiesResp_Resp {
	if x != nil {
		return x.Resp
	}
	return nil
}

func (x *GetStravaActivitiesResp) GetData() *GetStravaActivitiesResp_Data {
	if x != nil {
		if x, ok := x.Resp.(*GetStravaActivitiesResp_Data_); ok {
			return x.Data
		}
	}
	return nil
}

func (x *GetStravaActivitiesResp) GetError() *common.Error {
	if x != nil {
		if x, ok := x.Resp.(*GetStravaActivitiesResp_Error); ok {
			return x.Error
		}
	}
	return nil
}

type isGetStravaActivitiesResp_Resp interface {
	isGetStravaActivitiesResp_Resp()
}

type GetStravaActivitiesResp_Data_ struct {
	Data *GetStravaActivitiesResp_Data `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

type GetStravaActivitiesResp_Error struct {
	Error *common.Error `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*GetStravaActivitiesResp_Data_) isGetStravaActivitiesResp_Resp() {}

func (*GetStravaActivitiesResp_Error) isGetStravaActivitiesResp_Resp() {}

type IntegrateStravaAccountReq_Strava struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AthleteId     int64                  `protobuf:"varint,1,opt,name=athlete_id,json=athleteId,proto3" json:"athlete_id,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	FirstName     string                 `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Profile       string                 `protobuf:"bytes,5,opt,name=profile,proto3" json:"profile,omitempty"`
	AccessToken   string                 `protobuf:"bytes,6,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,7,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	ExpiresAt     int64                  `protobuf:"varint,8,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IntegrateStravaAccountReq_Strava) Reset() {
	*x = IntegrateStravaAccountReq_Strava{}
	mi := &file_integration_strava_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IntegrateStravaAccountReq_Strava) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegrateStravaAccountReq_Strava) ProtoMessage() {}

func (x *IntegrateStravaAccountReq_Strava) ProtoReflect() protoreflect.Message {
	mi := &file_integration_strava_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegrateStravaAccountReq_Strava.ProtoReflect.Descriptor instead.
func (*IntegrateStravaAccountReq_Strava) Descriptor() ([]byte, []int) {
	return file_integration_strava_service_proto_rawDescGZIP(), []int{0, 0}
}

func (x *IntegrateStravaAccountReq_Strava) GetAthleteId() int64 {
	if x != nil {
		return x.AthleteId
	}
	return 0
}

func (x *IntegrateStravaAccountReq_Strava) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *IntegrateStravaAccountReq_Strava) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *IntegrateStravaAccountReq_Strava) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *IntegrateStravaAccountReq_Strava) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *IntegrateStravaAccountReq_Strava) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *IntegrateStravaAccountReq_Strava) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *IntegrateStravaAccountReq_Strava) GetExpiresAt() int64 {
	if x != nil {
		return x.ExpiresAt
	}
	return 0
}

type GetStravaActivitiesResp_Data struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Activities    []*StravaActivity      `protobuf:"bytes,1,rep,name=activities,proto3" json:"activities,omitempty"`
	PageMetaData  *common.PageMetaData   `protobuf:"bytes,2,opt,name=page_meta_data,json=pageMetaData,proto3" json:"page_meta_data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStravaActivitiesResp_Data) Reset() {
	*x = GetStravaActivitiesResp_Data{}
	mi := &file_integration_strava_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStravaActivitiesResp_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStravaActivitiesResp_Data) ProtoMessage() {}

func (x *GetStravaActivitiesResp_Data) ProtoReflect() protoreflect.Message {
	mi := &file_integration_strava_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStravaActivitiesResp_Data.ProtoReflect.Descriptor instead.
func (*GetStravaActivitiesResp_Data) Descriptor() ([]byte, []int) {
	return file_integration_strava_service_proto_rawDescGZIP(), []int{4, 0}
}

func (x *GetStravaActivitiesResp_Data) GetActivities() []*StravaActivity {
	if x != nil {
		return x.Activities
	}
	return nil
}

func (x *GetStravaActivitiesResp_Data) GetPageMetaData() *common.PageMetaData {
	if x != nil {
		return x.PageMetaData
	}
	return nil
}

var File_integration_strava_service_proto protoreflect.FileDescriptor

var file_integration_strava_service_proto_rawDesc = string([]byte{
	0x0a, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x74,
	0x72, 0x61, 0x76, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x18, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x02, 0x0a,
	0x19, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x76, 0x61,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x61, 0x76, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x76,
	0x61, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x2e, 0x53, 0x74, 0x72, 0x61,
	0x76, 0x61, 0x52, 0x06, 0x73, 0x74, 0x72, 0x61, 0x76, 0x61, 0x1a, 0x80, 0x02, 0x0a, 0x06, 0x53,
	0x74, 0x72, 0x61, 0x76, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x74, 0x68, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d,
	0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x22, 0x2e, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x61, 0x76, 0x61, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xa4, 0x01,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x61, 0x76, 0x61, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x43, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x61, 0x76, 0x61, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x76, 0x61, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x61, 0x76, 0x61, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x06, 0x0a, 0x04,
	0x72, 0x65, 0x73, 0x70, 0x22, 0x89, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x61,
	0x76, 0x61, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x12,
	0x38, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x72, 0x61,
	0x76, 0x61, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x22, 0xa4, 0x02, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x61, 0x76, 0x61, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x3f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x61, 0x76, 0x61, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x48,
	0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x7f,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x72, 0x61, 0x76, 0x61, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x0e, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x0c, 0x70, 0x61, 0x67, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x42,
	0x06, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x70, 0x32, 0xc0, 0x04, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x61,
	0x76, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x78, 0x0a, 0x16, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x76, 0x61, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x26, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x61, 0x76,
	0x61, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x22, 0x23,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x76, 0x61, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x65, 0x12, 0x59, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x53, 0x74, 0x72,
	0x61, 0x76, 0x61, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x2a, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x74, 0x72, 0x61, 0x76, 0x61, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x77,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x61, 0x76, 0x61, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x20, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x61, 0x76, 0x61, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x61, 0x76, 0x61, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x12,
	0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x76, 0x61, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x5b, 0x0a, 0x14, 0x53, 0x79, 0x6e, 0x63, 0x53,
	0x74, 0x72, 0x61, 0x76, 0x61, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12,
	0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22,
	0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x76, 0x61, 0x2f,
	0x73, 0x79, 0x6e, 0x63, 0x12, 0x83, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x61,
	0x76, 0x61, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x23, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x72, 0x61, 0x76, 0x61, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x1a, 0x24, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x72, 0x61, 0x76, 0x61, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12,
	0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x61, 0x76, 0x61, 0x2f,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x42, 0xb5, 0x01, 0x0a, 0x0f, 0x63,
	0x6f, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x12,
	0x53, 0x74, 0x72, 0x61, 0x76, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6e, 0x67, 0x64, 0x61, 0x6e, 0x67, 0x6b, 0x69, 0x65, 0x74, 0x73, 0x77, 0x65, 0x2f, 0x73,
	0x77, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x49, 0x58, 0x58, 0xaa, 0x02,
	0x0b, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xca, 0x02, 0x0b, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xe2, 0x02, 0x17, 0x49, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_integration_strava_service_proto_rawDescOnce sync.Once
	file_integration_strava_service_proto_rawDescData []byte
)

func file_integration_strava_service_proto_rawDescGZIP() []byte {
	file_integration_strava_service_proto_rawDescOnce.Do(func() {
		file_integration_strava_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_integration_strava_service_proto_rawDesc), len(file_integration_strava_service_proto_rawDesc)))
	})
	return file_integration_strava_service_proto_rawDescData
}

var file_integration_strava_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_integration_strava_service_proto_goTypes = []any{
	(*IntegrateStravaAccountReq)(nil),        // 0: integration.IntegrateStravaAccountReq
	(*GetStravaAccountReq)(nil),              // 1: integration.GetStravaAccountReq
	(*GetStravaAccountResp)(nil),             // 2: integration.GetStravaAccountResp
	(*GetStravaActivitiesReq)(nil),           // 3: integration.GetStravaActivitiesReq
	(*GetStravaActivitiesResp)(nil),          // 4: integration.GetStravaActivitiesResp
	(*IntegrateStravaAccountReq_Strava)(nil), // 5: integration.IntegrateStravaAccountReq.Strava
	(*GetStravaActivitiesResp_Data)(nil),     // 6: integration.GetStravaActivitiesResp.Data
	(*StravaAccount)(nil),                    // 7: integration.StravaAccount
	(*common.Error)(nil),                     // 8: common.Error
	(StravaActivityType)(0),                  // 9: integration.StravaActivityType
	(*common.Pageable)(nil),                  // 10: common.Pageable
	(*StravaActivity)(nil),                   // 11: integration.StravaActivity
	(*common.PageMetaData)(nil),              // 12: common.PageMetaData
	(*common.EmptyReq)(nil),                  // 13: common.EmptyReq
	(*common.EmptyResp)(nil),                 // 14: common.EmptyResp
}
var file_integration_strava_service_proto_depIdxs = []int32{
	5,  // 0: integration.IntegrateStravaAccountReq.strava:type_name -> integration.IntegrateStravaAccountReq.Strava
	7,  // 1: integration.GetStravaAccountResp.strava_account:type_name -> integration.StravaAccount
	8,  // 2: integration.GetStravaAccountResp.error:type_name -> common.Error
	9,  // 3: integration.GetStravaActivitiesReq.type:type_name -> integration.StravaActivityType
	10, // 4: integration.GetStravaActivitiesReq.pageable:type_name -> common.Pageable
	6,  // 5: integration.GetStravaActivitiesResp.data:type_name -> integration.GetStravaActivitiesResp.Data
	8,  // 6: integration.GetStravaActivitiesResp.error:type_name -> common.Error
	11, // 7: integration.GetStravaActivitiesResp.Data.activities:type_name -> integration.StravaActivity
	12, // 8: integration.GetStravaActivitiesResp.Data.page_meta_data:type_name -> common.PageMetaData
	0,  // 9: integration.StravaService.IntegrateStravaAccount:input_type -> integration.IntegrateStravaAccountReq
	13, // 10: integration.StravaService.RemoveStravaAccount:input_type -> common.EmptyReq
	1,  // 11: integration.StravaService.GetStravaAccount:input_type -> integration.GetStravaAccountReq
	13, // 12: integration.StravaService.SyncStravaActivities:input_type -> common.EmptyReq
	3,  // 13: integration.StravaService.GetStravaActivities:input_type -> integration.GetStravaActivitiesReq
	14, // 14: integration.StravaService.IntegrateStravaAccount:output_type -> common.EmptyResp
	14, // 15: integration.StravaService.RemoveStravaAccount:output_type -> common.EmptyResp
	2,  // 16: integration.StravaService.GetStravaAccount:output_type -> integration.GetStravaAccountResp
	14, // 17: integration.StravaService.SyncStravaActivities:output_type -> common.EmptyResp
	4,  // 18: integration.StravaService.GetStravaActivities:output_type -> integration.GetStravaActivitiesResp
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_integration_strava_service_proto_init() }
func file_integration_strava_service_proto_init() {
	if File_integration_strava_service_proto != nil {
		return
	}
	file_integration_shared_proto_init()
	file_integration_strava_service_proto_msgTypes[2].OneofWrappers = []any{
		(*GetStravaAccountResp_StravaAccount)(nil),
		(*GetStravaAccountResp_Error)(nil),
	}
	file_integration_strava_service_proto_msgTypes[3].OneofWrappers = []any{}
	file_integration_strava_service_proto_msgTypes[4].OneofWrappers = []any{
		(*GetStravaActivitiesResp_Data_)(nil),
		(*GetStravaActivitiesResp_Error)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_integration_strava_service_proto_rawDesc), len(file_integration_strava_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_integration_strava_service_proto_goTypes,
		DependencyIndexes: file_integration_strava_service_proto_depIdxs,
		MessageInfos:      file_integration_strava_service_proto_msgTypes,
	}.Build()
	File_integration_strava_service_proto = out.File
	file_integration_strava_service_proto_goTypes = nil
	file_integration_strava_service_proto_depIdxs = nil
}
