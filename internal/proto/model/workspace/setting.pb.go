// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: model/workspace/setting.proto

package workspace

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type SettingKey int32

const (
	SettingKey_WORKSPACE_SETTING_KEY_UNSPECIFIED SettingKey = 0
	// BASIC is the key for basic settings.
	SettingKey_BASIC SettingKey = 1
	// GENERAL is the key for general settings.
	SettingKey_GENERAL SettingKey = 2
	// STORAGE is the key for storage settings.
	SettingKey_STORAGE SettingKey = 3
	// MEMO_RELATED is the key for memo related settings.
	SettingKey_MEMO_RELATED SettingKey = 4
)

// Enum value maps for SettingKey.
var (
	SettingKey_name = map[int32]string{
		0: "WORKSPACE_SETTING_KEY_UNSPECIFIED",
		1: "BASIC",
		2: "GENERAL",
		3: "STORAGE",
		4: "MEMO_RELATED",
	}
	SettingKey_value = map[string]int32{
		"WORKSPACE_SETTING_KEY_UNSPECIFIED": 0,
		"BASIC":                             1,
		"GENERAL":                           2,
		"STORAGE":                           3,
		"MEMO_RELATED":                      4,
	}
)

func (x SettingKey) Enum() *SettingKey {
	p := new(SettingKey)
	*p = x
	return p
}

func (x SettingKey) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SettingKey) Descriptor() protoreflect.EnumDescriptor {
	return file_model_workspace_setting_proto_enumTypes[0].Descriptor()
}

func (SettingKey) Type() protoreflect.EnumType {
	return &file_model_workspace_setting_proto_enumTypes[0]
}

func (x SettingKey) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SettingKey.Descriptor instead.
func (SettingKey) EnumDescriptor() ([]byte, []int) {
	return file_model_workspace_setting_proto_rawDescGZIP(), []int{0}
}

type StorageSetting_StorageType int32

const (
	StorageSetting_STORAGE_TYPE_UNSPECIFIED StorageSetting_StorageType = 0
	// DATABASE is the database storage type.
	StorageSetting_DATABASE StorageSetting_StorageType = 1
	// LOCAL is the local storage type.
	StorageSetting_LOCAL StorageSetting_StorageType = 2
	// S3 is the S3 storage type.
	StorageSetting_S3 StorageSetting_StorageType = 3
)

// Enum value maps for StorageSetting_StorageType.
var (
	StorageSetting_StorageType_name = map[int32]string{
		0: "STORAGE_TYPE_UNSPECIFIED",
		1: "DATABASE",
		2: "LOCAL",
		3: "S3",
	}
	StorageSetting_StorageType_value = map[string]int32{
		"STORAGE_TYPE_UNSPECIFIED": 0,
		"DATABASE":                 1,
		"LOCAL":                    2,
		"S3":                       3,
	}
)

func (x StorageSetting_StorageType) Enum() *StorageSetting_StorageType {
	p := new(StorageSetting_StorageType)
	*p = x
	return p
}

func (x StorageSetting_StorageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StorageSetting_StorageType) Descriptor() protoreflect.EnumDescriptor {
	return file_model_workspace_setting_proto_enumTypes[1].Descriptor()
}

func (StorageSetting_StorageType) Type() protoreflect.EnumType {
	return &file_model_workspace_setting_proto_enumTypes[1]
}

func (x StorageSetting_StorageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StorageSetting_StorageType.Descriptor instead.
func (StorageSetting_StorageType) EnumDescriptor() ([]byte, []int) {
	return file_model_workspace_setting_proto_rawDescGZIP(), []int{3, 0}
}

type Setting struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// name is the name of the setting.
	// Format: settings/{setting}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are valid to be assigned to Value:
	//
	//	*Setting_GeneralSetting
	//	*Setting_StorageSetting
	//	*Setting_MemoRelatedSetting
	Value         isSetting_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Setting) Reset() {
	*x = Setting{}
	mi := &file_model_workspace_setting_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Setting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Setting) ProtoMessage() {}

func (x *Setting) ProtoReflect() protoreflect.Message {
	mi := &file_model_workspace_setting_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Setting.ProtoReflect.Descriptor instead.
func (*Setting) Descriptor() ([]byte, []int) {
	return file_model_workspace_setting_proto_rawDescGZIP(), []int{0}
}

func (x *Setting) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Setting) GetValue() isSetting_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Setting) GetGeneralSetting() *GeneralSetting {
	if x != nil {
		if x, ok := x.Value.(*Setting_GeneralSetting); ok {
			return x.GeneralSetting
		}
	}
	return nil
}

func (x *Setting) GetStorageSetting() *StorageSetting {
	if x != nil {
		if x, ok := x.Value.(*Setting_StorageSetting); ok {
			return x.StorageSetting
		}
	}
	return nil
}

func (x *Setting) GetMemoRelatedSetting() *MemoRelatedSetting {
	if x != nil {
		if x, ok := x.Value.(*Setting_MemoRelatedSetting); ok {
			return x.MemoRelatedSetting
		}
	}
	return nil
}

type isSetting_Value interface {
	isSetting_Value()
}

type Setting_GeneralSetting struct {
	GeneralSetting *GeneralSetting `protobuf:"bytes,2,opt,name=general_setting,json=generalSetting,proto3,oneof"`
}

type Setting_StorageSetting struct {
	StorageSetting *StorageSetting `protobuf:"bytes,3,opt,name=storage_setting,json=storageSetting,proto3,oneof"`
}

type Setting_MemoRelatedSetting struct {
	MemoRelatedSetting *MemoRelatedSetting `protobuf:"bytes,4,opt,name=memo_related_setting,json=memoRelatedSetting,proto3,oneof"`
}

func (*Setting_GeneralSetting) isSetting_Value() {}

func (*Setting_StorageSetting) isSetting_Value() {}

func (*Setting_MemoRelatedSetting) isSetting_Value() {}

type GeneralSetting struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// disallow_user_registration disallows user registration.
	DisallowUserRegistration bool `protobuf:"varint,1,opt,name=disallow_user_registration,json=disallowUserRegistration,proto3" json:"disallow_user_registration,omitempty"`
	// disallow_password_auth disallows password authentication.
	DisallowPasswordAuth bool `protobuf:"varint,2,opt,name=disallow_password_auth,json=disallowPasswordAuth,proto3" json:"disallow_password_auth,omitempty"`
	// additional_script is the additional script.
	AdditionalScript string `protobuf:"bytes,3,opt,name=additional_script,json=additionalScript,proto3" json:"additional_script,omitempty"`
	// additional_style is the additional style.
	AdditionalStyle string `protobuf:"bytes,4,opt,name=additional_style,json=additionalStyle,proto3" json:"additional_style,omitempty"`
	// custom_profile is the custom profile.
	CustomProfile *CustomProfile `protobuf:"bytes,5,opt,name=custom_profile,json=customProfile,proto3" json:"custom_profile,omitempty"`
	// week_start_day_offset is the week start day offset from Sunday.
	// 0: Sunday, 1: Monday, 2: Tuesday, 3: Wednesday, 4: Thursday, 5: Friday, 6: Saturday
	// Default is Sunday.
	WeekStartDayOffset int32 `protobuf:"varint,6,opt,name=week_start_day_offset,json=weekStartDayOffset,proto3" json:"week_start_day_offset,omitempty"`
	// disallow_change_username disallows changing username.
	DisallowChangeUsername bool `protobuf:"varint,7,opt,name=disallow_change_username,json=disallowChangeUsername,proto3" json:"disallow_change_username,omitempty"`
	// disallow_change_nickname disallows changing nickname.
	DisallowChangeNickname bool `protobuf:"varint,8,opt,name=disallow_change_nickname,json=disallowChangeNickname,proto3" json:"disallow_change_nickname,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *GeneralSetting) Reset() {
	*x = GeneralSetting{}
	mi := &file_model_workspace_setting_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeneralSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralSetting) ProtoMessage() {}

func (x *GeneralSetting) ProtoReflect() protoreflect.Message {
	mi := &file_model_workspace_setting_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralSetting.ProtoReflect.Descriptor instead.
func (*GeneralSetting) Descriptor() ([]byte, []int) {
	return file_model_workspace_setting_proto_rawDescGZIP(), []int{1}
}

func (x *GeneralSetting) GetDisallowUserRegistration() bool {
	if x != nil {
		return x.DisallowUserRegistration
	}
	return false
}

func (x *GeneralSetting) GetDisallowPasswordAuth() bool {
	if x != nil {
		return x.DisallowPasswordAuth
	}
	return false
}

func (x *GeneralSetting) GetAdditionalScript() string {
	if x != nil {
		return x.AdditionalScript
	}
	return ""
}

func (x *GeneralSetting) GetAdditionalStyle() string {
	if x != nil {
		return x.AdditionalStyle
	}
	return ""
}

func (x *GeneralSetting) GetCustomProfile() *CustomProfile {
	if x != nil {
		return x.CustomProfile
	}
	return nil
}

func (x *GeneralSetting) GetWeekStartDayOffset() int32 {
	if x != nil {
		return x.WeekStartDayOffset
	}
	return 0
}

func (x *GeneralSetting) GetDisallowChangeUsername() bool {
	if x != nil {
		return x.DisallowChangeUsername
	}
	return false
}

func (x *GeneralSetting) GetDisallowChangeNickname() bool {
	if x != nil {
		return x.DisallowChangeNickname
	}
	return false
}

type CustomProfile struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	LogoUrl       string                 `protobuf:"bytes,3,opt,name=logo_url,json=logoUrl,proto3" json:"logo_url,omitempty"`
	Locale        string                 `protobuf:"bytes,4,opt,name=locale,proto3" json:"locale,omitempty"`
	Appearance    string                 `protobuf:"bytes,5,opt,name=appearance,proto3" json:"appearance,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomProfile) Reset() {
	*x = CustomProfile{}
	mi := &file_model_workspace_setting_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomProfile) ProtoMessage() {}

func (x *CustomProfile) ProtoReflect() protoreflect.Message {
	mi := &file_model_workspace_setting_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomProfile.ProtoReflect.Descriptor instead.
func (*CustomProfile) Descriptor() ([]byte, []int) {
	return file_model_workspace_setting_proto_rawDescGZIP(), []int{2}
}

func (x *CustomProfile) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CustomProfile) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CustomProfile) GetLogoUrl() string {
	if x != nil {
		return x.LogoUrl
	}
	return ""
}

func (x *CustomProfile) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *CustomProfile) GetAppearance() string {
	if x != nil {
		return x.Appearance
	}
	return ""
}

type StorageSetting struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// storage_type is the storage type.
	StorageType StorageSetting_StorageType `protobuf:"varint,1,opt,name=storage_type,json=storageType,proto3,enum=workspace.StorageSetting_StorageType" json:"storage_type,omitempty"`
	// The template of file path.
	// e.g. assets/{timestamp}_{filename}
	FilepathTemplate string `protobuf:"bytes,2,opt,name=filepath_template,json=filepathTemplate,proto3" json:"filepath_template,omitempty"`
	// The max upload size in megabytes.
	UploadSizeLimitMb int64 `protobuf:"varint,3,opt,name=upload_size_limit_mb,json=uploadSizeLimitMb,proto3" json:"upload_size_limit_mb,omitempty"`
	// The S3 config.
	S3Config      *StorageSetting_S3Config `protobuf:"bytes,4,opt,name=s3_config,json=s3Config,proto3" json:"s3_config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StorageSetting) Reset() {
	*x = StorageSetting{}
	mi := &file_model_workspace_setting_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StorageSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageSetting) ProtoMessage() {}

func (x *StorageSetting) ProtoReflect() protoreflect.Message {
	mi := &file_model_workspace_setting_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageSetting.ProtoReflect.Descriptor instead.
func (*StorageSetting) Descriptor() ([]byte, []int) {
	return file_model_workspace_setting_proto_rawDescGZIP(), []int{3}
}

func (x *StorageSetting) GetStorageType() StorageSetting_StorageType {
	if x != nil {
		return x.StorageType
	}
	return StorageSetting_STORAGE_TYPE_UNSPECIFIED
}

func (x *StorageSetting) GetFilepathTemplate() string {
	if x != nil {
		return x.FilepathTemplate
	}
	return ""
}

func (x *StorageSetting) GetUploadSizeLimitMb() int64 {
	if x != nil {
		return x.UploadSizeLimitMb
	}
	return 0
}

func (x *StorageSetting) GetS3Config() *StorageSetting_S3Config {
	if x != nil {
		return x.S3Config
	}
	return nil
}

type MemoRelatedSetting struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// disallow_public_visibility disallows set memo as public visibility.
	DisallowPublicVisibility bool `protobuf:"varint,1,opt,name=disallow_public_visibility,json=disallowPublicVisibility,proto3" json:"disallow_public_visibility,omitempty"`
	// display_with_update_time orders and displays memo with update time.
	DisplayWithUpdateTime bool `protobuf:"varint,2,opt,name=display_with_update_time,json=displayWithUpdateTime,proto3" json:"display_with_update_time,omitempty"`
	// content_length_limit is the limit of content length. Unit is byte.
	ContentLengthLimit int32 `protobuf:"varint,3,opt,name=content_length_limit,json=contentLengthLimit,proto3" json:"content_length_limit,omitempty"`
	// enable_double_click_edit enables editing on double click.
	EnableDoubleClickEdit bool `protobuf:"varint,5,opt,name=enable_double_click_edit,json=enableDoubleClickEdit,proto3" json:"enable_double_click_edit,omitempty"`
	// enable_link_preview enables links preview.
	EnableLinkPreview bool `protobuf:"varint,6,opt,name=enable_link_preview,json=enableLinkPreview,proto3" json:"enable_link_preview,omitempty"`
	// enable_comment enables comment.
	EnableComment bool `protobuf:"varint,7,opt,name=enable_comment,json=enableComment,proto3" json:"enable_comment,omitempty"`
	// enable_location enables setting location for memo.
	EnableLocation bool `protobuf:"varint,8,opt,name=enable_location,json=enableLocation,proto3" json:"enable_location,omitempty"`
	// reactions is the list of reactions.
	Reactions []string `protobuf:"bytes,10,rep,name=reactions,proto3" json:"reactions,omitempty"`
	// disable_markdown_shortcuts disallow the registration of markdown shortcuts.
	DisableMarkdownShortcuts bool `protobuf:"varint,11,opt,name=disable_markdown_shortcuts,json=disableMarkdownShortcuts,proto3" json:"disable_markdown_shortcuts,omitempty"`
	// enable_blur_nsfw_content enables blurring of content marked as not safe for work (NSFW).
	EnableBlurNsfwContent bool `protobuf:"varint,12,opt,name=enable_blur_nsfw_content,json=enableBlurNsfwContent,proto3" json:"enable_blur_nsfw_content,omitempty"`
	// nsfw_tags is the list of tags that mark content as NSFW for blurring.
	NsfwTags      []string `protobuf:"bytes,13,rep,name=nsfw_tags,json=nsfwTags,proto3" json:"nsfw_tags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MemoRelatedSetting) Reset() {
	*x = MemoRelatedSetting{}
	mi := &file_model_workspace_setting_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MemoRelatedSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemoRelatedSetting) ProtoMessage() {}

func (x *MemoRelatedSetting) ProtoReflect() protoreflect.Message {
	mi := &file_model_workspace_setting_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemoRelatedSetting.ProtoReflect.Descriptor instead.
func (*MemoRelatedSetting) Descriptor() ([]byte, []int) {
	return file_model_workspace_setting_proto_rawDescGZIP(), []int{4}
}

func (x *MemoRelatedSetting) GetDisallowPublicVisibility() bool {
	if x != nil {
		return x.DisallowPublicVisibility
	}
	return false
}

func (x *MemoRelatedSetting) GetDisplayWithUpdateTime() bool {
	if x != nil {
		return x.DisplayWithUpdateTime
	}
	return false
}

func (x *MemoRelatedSetting) GetContentLengthLimit() int32 {
	if x != nil {
		return x.ContentLengthLimit
	}
	return 0
}

func (x *MemoRelatedSetting) GetEnableDoubleClickEdit() bool {
	if x != nil {
		return x.EnableDoubleClickEdit
	}
	return false
}

func (x *MemoRelatedSetting) GetEnableLinkPreview() bool {
	if x != nil {
		return x.EnableLinkPreview
	}
	return false
}

func (x *MemoRelatedSetting) GetEnableComment() bool {
	if x != nil {
		return x.EnableComment
	}
	return false
}

func (x *MemoRelatedSetting) GetEnableLocation() bool {
	if x != nil {
		return x.EnableLocation
	}
	return false
}

func (x *MemoRelatedSetting) GetReactions() []string {
	if x != nil {
		return x.Reactions
	}
	return nil
}

func (x *MemoRelatedSetting) GetDisableMarkdownShortcuts() bool {
	if x != nil {
		return x.DisableMarkdownShortcuts
	}
	return false
}

func (x *MemoRelatedSetting) GetEnableBlurNsfwContent() bool {
	if x != nil {
		return x.EnableBlurNsfwContent
	}
	return false
}

func (x *MemoRelatedSetting) GetNsfwTags() []string {
	if x != nil {
		return x.NsfwTags
	}
	return nil
}

// Reference: https://developers.cloudflare.com/r2/examples/aws/aws-sdk-go/
type StorageSetting_S3Config struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	AccessKeyId     string                 `protobuf:"bytes,1,opt,name=access_key_id,json=accessKeyId,proto3" json:"access_key_id,omitempty"`
	AccessKeySecret string                 `protobuf:"bytes,2,opt,name=access_key_secret,json=accessKeySecret,proto3" json:"access_key_secret,omitempty"`
	Endpoint        string                 `protobuf:"bytes,3,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Region          string                 `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	Bucket          string                 `protobuf:"bytes,5,opt,name=bucket,proto3" json:"bucket,omitempty"`
	UsePathStyle    bool                   `protobuf:"varint,6,opt,name=use_path_style,json=usePathStyle,proto3" json:"use_path_style,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *StorageSetting_S3Config) Reset() {
	*x = StorageSetting_S3Config{}
	mi := &file_model_workspace_setting_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StorageSetting_S3Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageSetting_S3Config) ProtoMessage() {}

func (x *StorageSetting_S3Config) ProtoReflect() protoreflect.Message {
	mi := &file_model_workspace_setting_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageSetting_S3Config.ProtoReflect.Descriptor instead.
func (*StorageSetting_S3Config) Descriptor() ([]byte, []int) {
	return file_model_workspace_setting_proto_rawDescGZIP(), []int{3, 0}
}

func (x *StorageSetting_S3Config) GetAccessKeyId() string {
	if x != nil {
		return x.AccessKeyId
	}
	return ""
}

func (x *StorageSetting_S3Config) GetAccessKeySecret() string {
	if x != nil {
		return x.AccessKeySecret
	}
	return ""
}

func (x *StorageSetting_S3Config) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *StorageSetting_S3Config) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *StorageSetting_S3Config) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *StorageSetting_S3Config) GetUsePathStyle() bool {
	if x != nil {
		return x.UsePathStyle
	}
	return false
}

var File_model_workspace_setting_proto protoreflect.FileDescriptor

const file_model_workspace_setting_proto_rawDesc = "" +
	"\n" +
	"\x1dmodel/workspace/setting.proto\x12\tworkspace\x1a\x1fgoogle/protobuf/timestamp.proto\"\x85\x02\n" +
	"\aSetting\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12D\n" +
	"\x0fgeneral_setting\x18\x02 \x01(\v2\x19.workspace.GeneralSettingH\x00R\x0egeneralSetting\x12D\n" +
	"\x0fstorage_setting\x18\x03 \x01(\v2\x19.workspace.StorageSettingH\x00R\x0estorageSetting\x12Q\n" +
	"\x14memo_related_setting\x18\x04 \x01(\v2\x1d.workspace.MemoRelatedSettingH\x00R\x12memoRelatedSettingB\a\n" +
	"\x05value\"\xc4\x03\n" +
	"\x0eGeneralSetting\x12<\n" +
	"\x1adisallow_user_registration\x18\x01 \x01(\bR\x18disallowUserRegistration\x124\n" +
	"\x16disallow_password_auth\x18\x02 \x01(\bR\x14disallowPasswordAuth\x12+\n" +
	"\x11additional_script\x18\x03 \x01(\tR\x10additionalScript\x12)\n" +
	"\x10additional_style\x18\x04 \x01(\tR\x0fadditionalStyle\x12?\n" +
	"\x0ecustom_profile\x18\x05 \x01(\v2\x18.workspace.CustomProfileR\rcustomProfile\x121\n" +
	"\x15week_start_day_offset\x18\x06 \x01(\x05R\x12weekStartDayOffset\x128\n" +
	"\x18disallow_change_username\x18\a \x01(\bR\x16disallowChangeUsername\x128\n" +
	"\x18disallow_change_nickname\x18\b \x01(\bR\x16disallowChangeNickname\"\x9a\x01\n" +
	"\rCustomProfile\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12\x19\n" +
	"\blogo_url\x18\x03 \x01(\tR\alogoUrl\x12\x16\n" +
	"\x06locale\x18\x04 \x01(\tR\x06locale\x12\x1e\n" +
	"\n" +
	"appearance\x18\x05 \x01(\tR\n" +
	"appearance\"\x96\x04\n" +
	"\x0eStorageSetting\x12H\n" +
	"\fstorage_type\x18\x01 \x01(\x0e2%.workspace.StorageSetting.StorageTypeR\vstorageType\x12+\n" +
	"\x11filepath_template\x18\x02 \x01(\tR\x10filepathTemplate\x12/\n" +
	"\x14upload_size_limit_mb\x18\x03 \x01(\x03R\x11uploadSizeLimitMb\x12?\n" +
	"\ts3_config\x18\x04 \x01(\v2\".workspace.StorageSetting.S3ConfigR\bs3Config\x1a\xcc\x01\n" +
	"\bS3Config\x12\"\n" +
	"\raccess_key_id\x18\x01 \x01(\tR\vaccessKeyId\x12*\n" +
	"\x11access_key_secret\x18\x02 \x01(\tR\x0faccessKeySecret\x12\x1a\n" +
	"\bendpoint\x18\x03 \x01(\tR\bendpoint\x12\x16\n" +
	"\x06region\x18\x04 \x01(\tR\x06region\x12\x16\n" +
	"\x06bucket\x18\x05 \x01(\tR\x06bucket\x12$\n" +
	"\x0euse_path_style\x18\x06 \x01(\bR\fusePathStyle\"L\n" +
	"\vStorageType\x12\x1c\n" +
	"\x18STORAGE_TYPE_UNSPECIFIED\x10\x00\x12\f\n" +
	"\bDATABASE\x10\x01\x12\t\n" +
	"\x05LOCAL\x10\x02\x12\x06\n" +
	"\x02S3\x10\x03\"\xae\x04\n" +
	"\x12MemoRelatedSetting\x12<\n" +
	"\x1adisallow_public_visibility\x18\x01 \x01(\bR\x18disallowPublicVisibility\x127\n" +
	"\x18display_with_update_time\x18\x02 \x01(\bR\x15displayWithUpdateTime\x120\n" +
	"\x14content_length_limit\x18\x03 \x01(\x05R\x12contentLengthLimit\x127\n" +
	"\x18enable_double_click_edit\x18\x05 \x01(\bR\x15enableDoubleClickEdit\x12.\n" +
	"\x13enable_link_preview\x18\x06 \x01(\bR\x11enableLinkPreview\x12%\n" +
	"\x0eenable_comment\x18\a \x01(\bR\renableComment\x12'\n" +
	"\x0fenable_location\x18\b \x01(\bR\x0eenableLocation\x12\x1c\n" +
	"\treactions\x18\n" +
	" \x03(\tR\treactions\x12<\n" +
	"\x1adisable_markdown_shortcuts\x18\v \x01(\bR\x18disableMarkdownShortcuts\x127\n" +
	"\x18enable_blur_nsfw_content\x18\f \x01(\bR\x15enableBlurNsfwContent\x12\x1b\n" +
	"\tnsfw_tags\x18\r \x03(\tR\bnsfwTagsJ\x04\b\x04\x10\x05*j\n" +
	"\n" +
	"SettingKey\x12%\n" +
	"!WORKSPACE_SETTING_KEY_UNSPECIFIED\x10\x00\x12\t\n" +
	"\x05BASIC\x10\x01\x12\v\n" +
	"\aGENERAL\x10\x02\x12\v\n" +
	"\aSTORAGE\x10\x03\x12\x10\n" +
	"\fMEMO_RELATED\x10\x04B<Z:github.com/smartmemos/memos/internal/proto/model/workspaceb\x06proto3"

var (
	file_model_workspace_setting_proto_rawDescOnce sync.Once
	file_model_workspace_setting_proto_rawDescData []byte
)

func file_model_workspace_setting_proto_rawDescGZIP() []byte {
	file_model_workspace_setting_proto_rawDescOnce.Do(func() {
		file_model_workspace_setting_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_model_workspace_setting_proto_rawDesc), len(file_model_workspace_setting_proto_rawDesc)))
	})
	return file_model_workspace_setting_proto_rawDescData
}

var file_model_workspace_setting_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_model_workspace_setting_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_model_workspace_setting_proto_goTypes = []any{
	(SettingKey)(0),                 // 0: workspace.SettingKey
	(StorageSetting_StorageType)(0), // 1: workspace.StorageSetting.StorageType
	(*Setting)(nil),                 // 2: workspace.Setting
	(*GeneralSetting)(nil),          // 3: workspace.GeneralSetting
	(*CustomProfile)(nil),           // 4: workspace.CustomProfile
	(*StorageSetting)(nil),          // 5: workspace.StorageSetting
	(*MemoRelatedSetting)(nil),      // 6: workspace.MemoRelatedSetting
	(*StorageSetting_S3Config)(nil), // 7: workspace.StorageSetting.S3Config
}
var file_model_workspace_setting_proto_depIdxs = []int32{
	3, // 0: workspace.Setting.general_setting:type_name -> workspace.GeneralSetting
	5, // 1: workspace.Setting.storage_setting:type_name -> workspace.StorageSetting
	6, // 2: workspace.Setting.memo_related_setting:type_name -> workspace.MemoRelatedSetting
	4, // 3: workspace.GeneralSetting.custom_profile:type_name -> workspace.CustomProfile
	1, // 4: workspace.StorageSetting.storage_type:type_name -> workspace.StorageSetting.StorageType
	7, // 5: workspace.StorageSetting.s3_config:type_name -> workspace.StorageSetting.S3Config
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_model_workspace_setting_proto_init() }
func file_model_workspace_setting_proto_init() {
	if File_model_workspace_setting_proto != nil {
		return
	}
	file_model_workspace_setting_proto_msgTypes[0].OneofWrappers = []any{
		(*Setting_GeneralSetting)(nil),
		(*Setting_StorageSetting)(nil),
		(*Setting_MemoRelatedSetting)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_model_workspace_setting_proto_rawDesc), len(file_model_workspace_setting_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_workspace_setting_proto_goTypes,
		DependencyIndexes: file_model_workspace_setting_proto_depIdxs,
		EnumInfos:         file_model_workspace_setting_proto_enumTypes,
		MessageInfos:      file_model_workspace_setting_proto_msgTypes,
	}.Build()
	File_model_workspace_setting_proto = out.File
	file_model_workspace_setting_proto_goTypes = nil
	file_model_workspace_setting_proto_depIdxs = nil
}
