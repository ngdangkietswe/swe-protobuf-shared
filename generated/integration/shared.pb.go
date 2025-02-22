// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: integration/shared.proto

package integration

import (
	_ "github.com/ngdangkietswe/swe-protobuf-shared/generated/common"
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

type StravaActivityType int32

const (
	StravaActivityType_RUN               StravaActivityType = 0
	StravaActivityType_RIDE              StravaActivityType = 1
	StravaActivityType_SWIM              StravaActivityType = 2
	StravaActivityType_WORKOUT           StravaActivityType = 3
	StravaActivityType_HIKE              StravaActivityType = 4
	StravaActivityType_WALK              StravaActivityType = 5
	StravaActivityType_NORDIC_SKI        StravaActivityType = 6
	StravaActivityType_ALPINE_SKI        StravaActivityType = 7
	StravaActivityType_BACKCOUNTRY_SKI   StravaActivityType = 8
	StravaActivityType_CANOEING          StravaActivityType = 9
	StravaActivityType_CROSSFIT          StravaActivityType = 10
	StravaActivityType_E_BIKE_RIDE       StravaActivityType = 11
	StravaActivityType_ELLIPTICAL        StravaActivityType = 12
	StravaActivityType_ICE_SKATE         StravaActivityType = 13
	StravaActivityType_INLINE_SKATE      StravaActivityType = 14
	StravaActivityType_KAYAKING          StravaActivityType = 15
	StravaActivityType_KITESURF          StravaActivityType = 16
	StravaActivityType_ROCK_CLIMBING     StravaActivityType = 17
	StravaActivityType_ROLLER_SKI        StravaActivityType = 18
	StravaActivityType_ROWING            StravaActivityType = 19
	StravaActivityType_SNOWBOARD         StravaActivityType = 20
	StravaActivityType_SNOWSHOE          StravaActivityType = 21
	StravaActivityType_STAIR_STEPPER     StravaActivityType = 22
	StravaActivityType_STAND_UP_PADDLING StravaActivityType = 23
	StravaActivityType_SURFING           StravaActivityType = 24
	StravaActivityType_VIRTUAL_RIDE      StravaActivityType = 25
	StravaActivityType_VIRTUAL_RUN       StravaActivityType = 26
	StravaActivityType_WEIGHT_TRAINING   StravaActivityType = 27
	StravaActivityType_WINDSURF          StravaActivityType = 28
	StravaActivityType_YOGA              StravaActivityType = 29
	StravaActivityType_OTHER             StravaActivityType = 30
)

// Enum value maps for StravaActivityType.
var (
	StravaActivityType_name = map[int32]string{
		0:  "RUN",
		1:  "RIDE",
		2:  "SWIM",
		3:  "WORKOUT",
		4:  "HIKE",
		5:  "WALK",
		6:  "NORDIC_SKI",
		7:  "ALPINE_SKI",
		8:  "BACKCOUNTRY_SKI",
		9:  "CANOEING",
		10: "CROSSFIT",
		11: "E_BIKE_RIDE",
		12: "ELLIPTICAL",
		13: "ICE_SKATE",
		14: "INLINE_SKATE",
		15: "KAYAKING",
		16: "KITESURF",
		17: "ROCK_CLIMBING",
		18: "ROLLER_SKI",
		19: "ROWING",
		20: "SNOWBOARD",
		21: "SNOWSHOE",
		22: "STAIR_STEPPER",
		23: "STAND_UP_PADDLING",
		24: "SURFING",
		25: "VIRTUAL_RIDE",
		26: "VIRTUAL_RUN",
		27: "WEIGHT_TRAINING",
		28: "WINDSURF",
		29: "YOGA",
		30: "OTHER",
	}
	StravaActivityType_value = map[string]int32{
		"RUN":               0,
		"RIDE":              1,
		"SWIM":              2,
		"WORKOUT":           3,
		"HIKE":              4,
		"WALK":              5,
		"NORDIC_SKI":        6,
		"ALPINE_SKI":        7,
		"BACKCOUNTRY_SKI":   8,
		"CANOEING":          9,
		"CROSSFIT":          10,
		"E_BIKE_RIDE":       11,
		"ELLIPTICAL":        12,
		"ICE_SKATE":         13,
		"INLINE_SKATE":      14,
		"KAYAKING":          15,
		"KITESURF":          16,
		"ROCK_CLIMBING":     17,
		"ROLLER_SKI":        18,
		"ROWING":            19,
		"SNOWBOARD":         20,
		"SNOWSHOE":          21,
		"STAIR_STEPPER":     22,
		"STAND_UP_PADDLING": 23,
		"SURFING":           24,
		"VIRTUAL_RIDE":      25,
		"VIRTUAL_RUN":       26,
		"WEIGHT_TRAINING":   27,
		"WINDSURF":          28,
		"YOGA":              29,
		"OTHER":             30,
	}
)

func (x StravaActivityType) Enum() *StravaActivityType {
	p := new(StravaActivityType)
	*p = x
	return p
}

func (x StravaActivityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StravaActivityType) Descriptor() protoreflect.EnumDescriptor {
	return file_integration_shared_proto_enumTypes[0].Descriptor()
}

func (StravaActivityType) Type() protoreflect.EnumType {
	return &file_integration_shared_proto_enumTypes[0]
}

func (x StravaActivityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StravaActivityType.Descriptor instead.
func (StravaActivityType) EnumDescriptor() ([]byte, []int) {
	return file_integration_shared_proto_rawDescGZIP(), []int{0}
}

type StravaAccount struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AccessToken   string                 `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken  string                 `protobuf:"bytes,4,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	ExpiresAt     string                 `protobuf:"bytes,5,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	TokenType     string                 `protobuf:"bytes,6,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	AthleteId     int64                  `protobuf:"varint,7,opt,name=athlete_id,json=athleteId,proto3" json:"athlete_id,omitempty"`
	Username      string                 `protobuf:"bytes,8,opt,name=username,proto3" json:"username,omitempty"`
	FirstName     string                 `protobuf:"bytes,9,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName      string                 `protobuf:"bytes,10,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StravaAccount) Reset() {
	*x = StravaAccount{}
	mi := &file_integration_shared_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StravaAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StravaAccount) ProtoMessage() {}

func (x *StravaAccount) ProtoReflect() protoreflect.Message {
	mi := &file_integration_shared_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StravaAccount.ProtoReflect.Descriptor instead.
func (*StravaAccount) Descriptor() ([]byte, []int) {
	return file_integration_shared_proto_rawDescGZIP(), []int{0}
}

func (x *StravaAccount) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StravaAccount) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *StravaAccount) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *StravaAccount) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *StravaAccount) GetExpiresAt() string {
	if x != nil {
		return x.ExpiresAt
	}
	return ""
}

func (x *StravaAccount) GetTokenType() string {
	if x != nil {
		return x.TokenType
	}
	return ""
}

func (x *StravaAccount) GetAthleteId() int64 {
	if x != nil {
		return x.AthleteId
	}
	return 0
}

func (x *StravaAccount) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *StravaAccount) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *StravaAccount) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *StravaAccount) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *StravaAccount) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type StravaActivity struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	Id                 string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId             string                 `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AthleteId          int64                  `protobuf:"varint,3,opt,name=athlete_id,json=athleteId,proto3" json:"athlete_id,omitempty"`
	StravaActivityId   int64                  `protobuf:"varint,4,opt,name=strava_activity_id,json=stravaActivityId,proto3" json:"strava_activity_id,omitempty"`
	ActivityName       string                 `protobuf:"bytes,5,opt,name=activity_name,json=activityName,proto3" json:"activity_name,omitempty"`
	ActivityType       int32                  `protobuf:"varint,6,opt,name=activity_type,json=activityType,proto3" json:"activity_type,omitempty"`
	ActivityUrl        string                 `protobuf:"bytes,7,opt,name=activity_url,json=activityUrl,proto3" json:"activity_url,omitempty"`
	StartDate          string                 `protobuf:"bytes,8,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	Distance           string                 `protobuf:"bytes,9,opt,name=distance,proto3" json:"distance,omitempty"`                           // Ex: 3.5 (km)
	MovingTime         string                 `protobuf:"bytes,10,opt,name=moving_time,json=movingTime,proto3" json:"moving_time,omitempty"`    // Ex: 12:30 (HH:mm)
	ElapsedTime        string                 `protobuf:"bytes,11,opt,name=elapsed_time,json=elapsedTime,proto3" json:"elapsed_time,omitempty"` // Ex: 12:30 (HH:mm)
	TotalElevationGain float64                `protobuf:"fixed64,12,opt,name=total_elevation_gain,json=totalElevationGain,proto3" json:"total_elevation_gain,omitempty"`
	AverageSpeed       float64                `protobuf:"fixed64,13,opt,name=average_speed,json=averageSpeed,proto3" json:"average_speed,omitempty"` // Ex: 4.5 (km/h)
	MaxSpeed           float64                `protobuf:"fixed64,14,opt,name=max_speed,json=maxSpeed,proto3" json:"max_speed,omitempty"`             // Ex: 10.0 (km/h)
	Pace               string                 `protobuf:"bytes,15,opt,name=pace,proto3" json:"pace,omitempty"`                                       // Ex: 05:00/km
	CreatedAt          string                 `protobuf:"bytes,16,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *StravaActivity) Reset() {
	*x = StravaActivity{}
	mi := &file_integration_shared_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StravaActivity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StravaActivity) ProtoMessage() {}

func (x *StravaActivity) ProtoReflect() protoreflect.Message {
	mi := &file_integration_shared_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StravaActivity.ProtoReflect.Descriptor instead.
func (*StravaActivity) Descriptor() ([]byte, []int) {
	return file_integration_shared_proto_rawDescGZIP(), []int{1}
}

func (x *StravaActivity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StravaActivity) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *StravaActivity) GetAthleteId() int64 {
	if x != nil {
		return x.AthleteId
	}
	return 0
}

func (x *StravaActivity) GetStravaActivityId() int64 {
	if x != nil {
		return x.StravaActivityId
	}
	return 0
}

func (x *StravaActivity) GetActivityName() string {
	if x != nil {
		return x.ActivityName
	}
	return ""
}

func (x *StravaActivity) GetActivityType() int32 {
	if x != nil {
		return x.ActivityType
	}
	return 0
}

func (x *StravaActivity) GetActivityUrl() string {
	if x != nil {
		return x.ActivityUrl
	}
	return ""
}

func (x *StravaActivity) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *StravaActivity) GetDistance() string {
	if x != nil {
		return x.Distance
	}
	return ""
}

func (x *StravaActivity) GetMovingTime() string {
	if x != nil {
		return x.MovingTime
	}
	return ""
}

func (x *StravaActivity) GetElapsedTime() string {
	if x != nil {
		return x.ElapsedTime
	}
	return ""
}

func (x *StravaActivity) GetTotalElevationGain() float64 {
	if x != nil {
		return x.TotalElevationGain
	}
	return 0
}

func (x *StravaActivity) GetAverageSpeed() float64 {
	if x != nil {
		return x.AverageSpeed
	}
	return 0
}

func (x *StravaActivity) GetMaxSpeed() float64 {
	if x != nil {
		return x.MaxSpeed
	}
	return 0
}

func (x *StravaActivity) GetPace() string {
	if x != nil {
		return x.Pace
	}
	return ""
}

func (x *StravaActivity) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

var File_integration_shared_proto protoreflect.FileDescriptor

var file_integration_shared_proto_rawDesc = string([]byte{
	0x0a, 0x18, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x02, 0x0a,
	0x0d, 0x53, 0x74, 0x72, 0x61, 0x76, 0x61, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x61, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x61, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x99, 0x04, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x61, 0x76, 0x61, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x61, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a,
	0x12, 0x73, 0x74, 0x72, 0x61, 0x76, 0x61, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x73, 0x74, 0x72, 0x61, 0x76,
	0x61, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x6f, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x76, 0x69, 0x6e, 0x67,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6c, 0x61, 0x70,
	0x73, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x61, 0x69, 0x6e, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x6c, 0x65, 0x76,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x61, 0x69, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x76, 0x65,
	0x72, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0c, 0x61, 0x76, 0x65, 0x72, 0x61, 0x67, 0x65, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x2a, 0xdf,
	0x03, 0x0a, 0x12, 0x53, 0x74, 0x72, 0x61, 0x76, 0x61, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x52, 0x55, 0x4e, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x52, 0x49, 0x44, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x57, 0x49, 0x4d,
	0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x4f, 0x52, 0x4b, 0x4f, 0x55, 0x54, 0x10, 0x03, 0x12,
	0x08, 0x0a, 0x04, 0x48, 0x49, 0x4b, 0x45, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x41, 0x4c,
	0x4b, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x4f, 0x52, 0x44, 0x49, 0x43, 0x5f, 0x53, 0x4b,
	0x49, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x4c, 0x50, 0x49, 0x4e, 0x45, 0x5f, 0x53, 0x4b,
	0x49, 0x10, 0x07, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x41, 0x43, 0x4b, 0x43, 0x4f, 0x55, 0x4e, 0x54,
	0x52, 0x59, 0x5f, 0x53, 0x4b, 0x49, 0x10, 0x08, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x41, 0x4e, 0x4f,
	0x45, 0x49, 0x4e, 0x47, 0x10, 0x09, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x4f, 0x53, 0x53, 0x46,
	0x49, 0x54, 0x10, 0x0a, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x5f, 0x42, 0x49, 0x4b, 0x45, 0x5f, 0x52,
	0x49, 0x44, 0x45, 0x10, 0x0b, 0x12, 0x0e, 0x0a, 0x0a, 0x45, 0x4c, 0x4c, 0x49, 0x50, 0x54, 0x49,
	0x43, 0x41, 0x4c, 0x10, 0x0c, 0x12, 0x0d, 0x0a, 0x09, 0x49, 0x43, 0x45, 0x5f, 0x53, 0x4b, 0x41,
	0x54, 0x45, 0x10, 0x0d, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x5f, 0x53,
	0x4b, 0x41, 0x54, 0x45, 0x10, 0x0e, 0x12, 0x0c, 0x0a, 0x08, 0x4b, 0x41, 0x59, 0x41, 0x4b, 0x49,
	0x4e, 0x47, 0x10, 0x0f, 0x12, 0x0c, 0x0a, 0x08, 0x4b, 0x49, 0x54, 0x45, 0x53, 0x55, 0x52, 0x46,
	0x10, 0x10, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x4f, 0x43, 0x4b, 0x5f, 0x43, 0x4c, 0x49, 0x4d, 0x42,
	0x49, 0x4e, 0x47, 0x10, 0x11, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x4f, 0x4c, 0x4c, 0x45, 0x52, 0x5f,
	0x53, 0x4b, 0x49, 0x10, 0x12, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x4f, 0x57, 0x49, 0x4e, 0x47, 0x10,
	0x13, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x4e, 0x4f, 0x57, 0x42, 0x4f, 0x41, 0x52, 0x44, 0x10, 0x14,
	0x12, 0x0c, 0x0a, 0x08, 0x53, 0x4e, 0x4f, 0x57, 0x53, 0x48, 0x4f, 0x45, 0x10, 0x15, 0x12, 0x11,
	0x0a, 0x0d, 0x53, 0x54, 0x41, 0x49, 0x52, 0x5f, 0x53, 0x54, 0x45, 0x50, 0x50, 0x45, 0x52, 0x10,
	0x16, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x5f, 0x55, 0x50, 0x5f, 0x50, 0x41,
	0x44, 0x44, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x17, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x52, 0x46,
	0x49, 0x4e, 0x47, 0x10, 0x18, 0x12, 0x10, 0x0a, 0x0c, 0x56, 0x49, 0x52, 0x54, 0x55, 0x41, 0x4c,
	0x5f, 0x52, 0x49, 0x44, 0x45, 0x10, 0x19, 0x12, 0x0f, 0x0a, 0x0b, 0x56, 0x49, 0x52, 0x54, 0x55,
	0x41, 0x4c, 0x5f, 0x52, 0x55, 0x4e, 0x10, 0x1a, 0x12, 0x13, 0x0a, 0x0f, 0x57, 0x45, 0x49, 0x47,
	0x48, 0x54, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x1b, 0x12, 0x0c, 0x0a,
	0x08, 0x57, 0x49, 0x4e, 0x44, 0x53, 0x55, 0x52, 0x46, 0x10, 0x1c, 0x12, 0x08, 0x0a, 0x04, 0x59,
	0x4f, 0x47, 0x41, 0x10, 0x1d, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x1e,
	0x42, 0xae, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6e, 0x67, 0x64, 0x61, 0x6e, 0x67, 0x6b, 0x69, 0x65, 0x74, 0x73, 0x77, 0x65, 0x2f, 0x73, 0x77,
	0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x49, 0x58, 0x58, 0xaa, 0x02, 0x0b,
	0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xca, 0x02, 0x0b, 0x49, 0x6e,
	0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xe2, 0x02, 0x17, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_integration_shared_proto_rawDescOnce sync.Once
	file_integration_shared_proto_rawDescData []byte
)

func file_integration_shared_proto_rawDescGZIP() []byte {
	file_integration_shared_proto_rawDescOnce.Do(func() {
		file_integration_shared_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_integration_shared_proto_rawDesc), len(file_integration_shared_proto_rawDesc)))
	})
	return file_integration_shared_proto_rawDescData
}

var file_integration_shared_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_integration_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_integration_shared_proto_goTypes = []any{
	(StravaActivityType)(0), // 0: integration.StravaActivityType
	(*StravaAccount)(nil),   // 1: integration.StravaAccount
	(*StravaActivity)(nil),  // 2: integration.StravaActivity
}
var file_integration_shared_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_integration_shared_proto_init() }
func file_integration_shared_proto_init() {
	if File_integration_shared_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_integration_shared_proto_rawDesc), len(file_integration_shared_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_integration_shared_proto_goTypes,
		DependencyIndexes: file_integration_shared_proto_depIdxs,
		EnumInfos:         file_integration_shared_proto_enumTypes,
		MessageInfos:      file_integration_shared_proto_msgTypes,
	}.Build()
	File_integration_shared_proto = out.File
	file_integration_shared_proto_goTypes = nil
	file_integration_shared_proto_depIdxs = nil
}
