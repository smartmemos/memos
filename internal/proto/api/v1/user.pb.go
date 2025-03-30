// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: api/v1/user.proto

package v1

import (
	user "github.com/smartmemos/memos/internal/proto/model/user"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type ListAccessTokensRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the user.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAccessTokensRequest) Reset() {
	*x = ListAccessTokensRequest{}
	mi := &file_api_v1_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAccessTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccessTokensRequest) ProtoMessage() {}

func (x *ListAccessTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccessTokensRequest.ProtoReflect.Descriptor instead.
func (*ListAccessTokensRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *ListAccessTokensRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListAccessTokensResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessTokens  []*user.AccessToken    `protobuf:"bytes,1,rep,name=access_tokens,json=accessTokens,proto3" json:"access_tokens,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAccessTokensResponse) Reset() {
	*x = ListAccessTokensResponse{}
	mi := &file_api_v1_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAccessTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccessTokensResponse) ProtoMessage() {}

func (x *ListAccessTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccessTokensResponse.ProtoReflect.Descriptor instead.
func (*ListAccessTokensResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *ListAccessTokensResponse) GetAccessTokens() []*user.AccessToken {
	if x != nil {
		return x.AccessTokens
	}
	return nil
}

type CreateAccessTokenRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the user.
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ExpiresAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=expires_at,json=expiresAt,proto3,oneof" json:"expires_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateAccessTokenRequest) Reset() {
	*x = CreateAccessTokenRequest{}
	mi := &file_api_v1_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAccessTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccessTokenRequest) ProtoMessage() {}

func (x *CreateAccessTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccessTokenRequest.ProtoReflect.Descriptor instead.
func (*CreateAccessTokenRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAccessTokenRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateAccessTokenRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateAccessTokenRequest) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

type ListAllUserStatsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAllUserStatsRequest) Reset() {
	*x = ListAllUserStatsRequest{}
	mi := &file_api_v1_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAllUserStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllUserStatsRequest) ProtoMessage() {}

func (x *ListAllUserStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllUserStatsRequest.ProtoReflect.Descriptor instead.
func (*ListAllUserStatsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_user_proto_rawDescGZIP(), []int{3}
}

type ListAllUserStatsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserStats     []*user.Stats          `protobuf:"bytes,1,rep,name=user_stats,json=userStats,proto3" json:"user_stats,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAllUserStatsResponse) Reset() {
	*x = ListAllUserStatsResponse{}
	mi := &file_api_v1_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAllUserStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllUserStatsResponse) ProtoMessage() {}

func (x *ListAllUserStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllUserStatsResponse.ProtoReflect.Descriptor instead.
func (*ListAllUserStatsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_user_proto_rawDescGZIP(), []int{4}
}

func (x *ListAllUserStatsResponse) GetUserStats() []*user.Stats {
	if x != nil {
		return x.UserStats
	}
	return nil
}

type GetUserStatsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the user.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserStatsRequest) Reset() {
	*x = GetUserStatsRequest{}
	mi := &file_api_v1_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserStatsRequest) ProtoMessage() {}

func (x *GetUserStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserStatsRequest.ProtoReflect.Descriptor instead.
func (*GetUserStatsRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_user_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserStatsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Nickname      string                 `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	AvatarUrl     string                 `protobuf:"bytes,5,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
	Description   string                 `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Password      string                 `protobuf:"bytes,7,opt,name=password,proto3" json:"password,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_api_v1_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_user_proto_rawDescGZIP(), []int{6}
}

func (x *CreateUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *CreateUserRequest) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

func (x *CreateUserRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetUserSettingRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the user.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserSettingRequest) Reset() {
	*x = GetUserSettingRequest{}
	mi := &file_api_v1_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserSettingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserSettingRequest) ProtoMessage() {}

func (x *GetUserSettingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserSettingRequest.ProtoReflect.Descriptor instead.
func (*GetUserSettingRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserSettingRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateUserSettingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Setting       *user.Setting          `protobuf:"bytes,1,opt,name=setting,proto3" json:"setting,omitempty"`
	UpdateMask    *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserSettingRequest) Reset() {
	*x = UpdateUserSettingRequest{}
	mi := &file_api_v1_user_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserSettingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserSettingRequest) ProtoMessage() {}

func (x *UpdateUserSettingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_user_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserSettingRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserSettingRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_user_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateUserSettingRequest) GetSetting() *user.Setting {
	if x != nil {
		return x.Setting
	}
	return nil
}

func (x *UpdateUserSettingRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

var File_api_v1_user_proto protoreflect.FileDescriptor

const file_api_v1_user_proto_rawDesc = "" +
	"\n" +
	"\x11api/v1/user.proto\x12\x06api.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x17google/api/client.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a google/protobuf/field_mask.proto\x1a\x15model/user/user.proto\x1a\x18model/user/setting.proto\x1a\x16model/user/stats.proto\x1a\x1dmodel/user/access_token.proto\"-\n" +
	"\x17ListAccessTokensRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"R\n" +
	"\x18ListAccessTokensResponse\x126\n" +
	"\raccess_tokens\x18\x01 \x03(\v2\x11.user.AccessTokenR\faccessTokens\"\x9f\x01\n" +
	"\x18CreateAccessTokenRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12>\n" +
	"\n" +
	"expires_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampH\x00R\texpiresAt\x88\x01\x01B\r\n" +
	"\v_expires_at\"\x19\n" +
	"\x17ListAllUserStatsRequest\"F\n" +
	"\x18ListAllUserStatsResponse\x12*\n" +
	"\n" +
	"user_stats\x18\x01 \x03(\v2\v.user.StatsR\tuserStats\")\n" +
	"\x13GetUserStatsRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"\xd7\x01\n" +
	"\x11CreateUserRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1a\n" +
	"\busername\x18\x02 \x01(\tR\busername\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x1a\n" +
	"\bnickname\x18\x04 \x01(\tR\bnickname\x12\x1d\n" +
	"\n" +
	"avatar_url\x18\x05 \x01(\tR\tavatarUrl\x12 \n" +
	"\vdescription\x18\x06 \x01(\tR\vdescription\x12\x1f\n" +
	"\bpassword\x18\a \x01(\tB\x03\xe0A\x04R\bpassword\"+\n" +
	"\x15GetUserSettingRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"\x85\x01\n" +
	"\x18UpdateUserSettingRequest\x12,\n" +
	"\asetting\x18\x01 \x01(\v2\r.user.SettingB\x03\xe0A\x02R\asetting\x12;\n" +
	"\vupdate_mask\x18\x02 \x01(\v2\x1a.google.protobuf.FieldMaskR\n" +
	"updateMask2\xe9\x06\n" +
	"\vUserService\x12h\n" +
	"\n" +
	"CreateUser\x12\x19.api.v1.CreateUserRequest\x1a\n" +
	".user.User\"3\xdaA\x16name,username,password\x82\xd3\xe4\x93\x02\x14\"\x12/api/v1/user/users\x12m\n" +
	"\x0eGetUserSetting\x12\x1d.api.v1.GetUserSettingRequest\x1a\r.user.Setting\"-\xdaA\x04name\x82\xd3\xe4\x93\x02 \x12\x1e/api/v1/{name=users/*}/setting\x12\x93\x01\n" +
	"\x11UpdateUserSetting\x12 .api.v1.UpdateUserSettingRequest\x1a\r.user.Setting\"M\xdaA\x13setting,update_mask\x82\xd3\xe4\x93\x021:\asetting2&/api/v1/{setting.name=users/*/setting}\x12t\n" +
	"\x10ListAllUserStats\x12\x1f.api.v1.ListAllUserStatsRequest\x1a .api.v1.ListAllUserStatsResponse\"\x1d\x82\xd3\xe4\x93\x02\x17\"\x15/api/v1/users/-/stats\x12e\n" +
	"\fGetUserStats\x12\x1b.api.v1.GetUserStatsRequest\x1a\v.user.Stats\"+\xdaA\x04name\x82\xd3\xe4\x93\x02\x1e\x12\x1c/api/v1/{name=users/*}/stats\x12\x80\x01\n" +
	"\x11CreateAccessToken\x12 .api.v1.CreateAccessTokenRequest\x1a\x11.user.AccessToken\"6\xdaA\x04name\x82\xd3\xe4\x93\x02):\x01*\"$/api/v1/{name=users/*}/access_tokens\x12\x8a\x01\n" +
	"\x10ListAccessTokens\x12\x1f.api.v1.ListAccessTokensRequest\x1a .api.v1.ListAccessTokensResponse\"3\xdaA\x04name\x82\xd3\xe4\x93\x02&\x12$/api/v1/{name=users/*}/access_tokensB3Z1github.com/smartmemos/memos/internal/proto/api/v1b\x06proto3"

var (
	file_api_v1_user_proto_rawDescOnce sync.Once
	file_api_v1_user_proto_rawDescData []byte
)

func file_api_v1_user_proto_rawDescGZIP() []byte {
	file_api_v1_user_proto_rawDescOnce.Do(func() {
		file_api_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1_user_proto_rawDesc), len(file_api_v1_user_proto_rawDesc)))
	})
	return file_api_v1_user_proto_rawDescData
}

var file_api_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_v1_user_proto_goTypes = []any{
	(*ListAccessTokensRequest)(nil),  // 0: api.v1.ListAccessTokensRequest
	(*ListAccessTokensResponse)(nil), // 1: api.v1.ListAccessTokensResponse
	(*CreateAccessTokenRequest)(nil), // 2: api.v1.CreateAccessTokenRequest
	(*ListAllUserStatsRequest)(nil),  // 3: api.v1.ListAllUserStatsRequest
	(*ListAllUserStatsResponse)(nil), // 4: api.v1.ListAllUserStatsResponse
	(*GetUserStatsRequest)(nil),      // 5: api.v1.GetUserStatsRequest
	(*CreateUserRequest)(nil),        // 6: api.v1.CreateUserRequest
	(*GetUserSettingRequest)(nil),    // 7: api.v1.GetUserSettingRequest
	(*UpdateUserSettingRequest)(nil), // 8: api.v1.UpdateUserSettingRequest
	(*user.AccessToken)(nil),         // 9: user.AccessToken
	(*timestamppb.Timestamp)(nil),    // 10: google.protobuf.Timestamp
	(*user.Stats)(nil),               // 11: user.Stats
	(*user.Setting)(nil),             // 12: user.Setting
	(*fieldmaskpb.FieldMask)(nil),    // 13: google.protobuf.FieldMask
	(*user.User)(nil),                // 14: user.User
}
var file_api_v1_user_proto_depIdxs = []int32{
	9,  // 0: api.v1.ListAccessTokensResponse.access_tokens:type_name -> user.AccessToken
	10, // 1: api.v1.CreateAccessTokenRequest.expires_at:type_name -> google.protobuf.Timestamp
	11, // 2: api.v1.ListAllUserStatsResponse.user_stats:type_name -> user.Stats
	12, // 3: api.v1.UpdateUserSettingRequest.setting:type_name -> user.Setting
	13, // 4: api.v1.UpdateUserSettingRequest.update_mask:type_name -> google.protobuf.FieldMask
	6,  // 5: api.v1.UserService.CreateUser:input_type -> api.v1.CreateUserRequest
	7,  // 6: api.v1.UserService.GetUserSetting:input_type -> api.v1.GetUserSettingRequest
	8,  // 7: api.v1.UserService.UpdateUserSetting:input_type -> api.v1.UpdateUserSettingRequest
	3,  // 8: api.v1.UserService.ListAllUserStats:input_type -> api.v1.ListAllUserStatsRequest
	5,  // 9: api.v1.UserService.GetUserStats:input_type -> api.v1.GetUserStatsRequest
	2,  // 10: api.v1.UserService.CreateAccessToken:input_type -> api.v1.CreateAccessTokenRequest
	0,  // 11: api.v1.UserService.ListAccessTokens:input_type -> api.v1.ListAccessTokensRequest
	14, // 12: api.v1.UserService.CreateUser:output_type -> user.User
	12, // 13: api.v1.UserService.GetUserSetting:output_type -> user.Setting
	12, // 14: api.v1.UserService.UpdateUserSetting:output_type -> user.Setting
	4,  // 15: api.v1.UserService.ListAllUserStats:output_type -> api.v1.ListAllUserStatsResponse
	11, // 16: api.v1.UserService.GetUserStats:output_type -> user.Stats
	9,  // 17: api.v1.UserService.CreateAccessToken:output_type -> user.AccessToken
	1,  // 18: api.v1.UserService.ListAccessTokens:output_type -> api.v1.ListAccessTokensResponse
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_v1_user_proto_init() }
func file_api_v1_user_proto_init() {
	if File_api_v1_user_proto != nil {
		return
	}
	file_api_v1_user_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1_user_proto_rawDesc), len(file_api_v1_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_user_proto_goTypes,
		DependencyIndexes: file_api_v1_user_proto_depIdxs,
		MessageInfos:      file_api_v1_user_proto_msgTypes,
	}.Build()
	File_api_v1_user_proto = out.File
	file_api_v1_user_proto_goTypes = nil
	file_api_v1_user_proto_depIdxs = nil
}
