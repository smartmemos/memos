// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: model/user/setting.proto

package user

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

type Setting struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The preferred locale of the user.
	Locale string `protobuf:"bytes,2,opt,name=locale,proto3" json:"locale,omitempty"`
	// The preferred appearance of the user.
	Appearance string `protobuf:"bytes,3,opt,name=appearance,proto3" json:"appearance,omitempty"`
	// The default visibility of the memo.
	MemoVisibility string `protobuf:"bytes,4,opt,name=memo_visibility,json=memoVisibility,proto3" json:"memo_visibility,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Setting) Reset() {
	*x = Setting{}
	mi := &file_model_user_setting_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Setting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Setting) ProtoMessage() {}

func (x *Setting) ProtoReflect() protoreflect.Message {
	mi := &file_model_user_setting_proto_msgTypes[0]
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
	return file_model_user_setting_proto_rawDescGZIP(), []int{0}
}

func (x *Setting) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Setting) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Setting) GetAppearance() string {
	if x != nil {
		return x.Appearance
	}
	return ""
}

func (x *Setting) GetMemoVisibility() string {
	if x != nil {
		return x.MemoVisibility
	}
	return ""
}

var File_model_user_setting_proto protoreflect.FileDescriptor

const file_model_user_setting_proto_rawDesc = "" +
	"\n" +
	"\x18model/user/setting.proto\x12\x04user\x1a\x1fgoogle/protobuf/timestamp.proto\"~\n" +
	"\aSetting\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x16\n" +
	"\x06locale\x18\x02 \x01(\tR\x06locale\x12\x1e\n" +
	"\n" +
	"appearance\x18\x03 \x01(\tR\n" +
	"appearance\x12'\n" +
	"\x0fmemo_visibility\x18\x04 \x01(\tR\x0ememoVisibilityB7Z5github.com/smartmemos/memos/internal/proto/model/userb\x06proto3"

var (
	file_model_user_setting_proto_rawDescOnce sync.Once
	file_model_user_setting_proto_rawDescData []byte
)

func file_model_user_setting_proto_rawDescGZIP() []byte {
	file_model_user_setting_proto_rawDescOnce.Do(func() {
		file_model_user_setting_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_model_user_setting_proto_rawDesc), len(file_model_user_setting_proto_rawDesc)))
	})
	return file_model_user_setting_proto_rawDescData
}

var file_model_user_setting_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_model_user_setting_proto_goTypes = []any{
	(*Setting)(nil), // 0: user.Setting
}
var file_model_user_setting_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_model_user_setting_proto_init() }
func file_model_user_setting_proto_init() {
	if File_model_user_setting_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_model_user_setting_proto_rawDesc), len(file_model_user_setting_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_user_setting_proto_goTypes,
		DependencyIndexes: file_model_user_setting_proto_depIdxs,
		MessageInfos:      file_model_user_setting_proto_msgTypes,
	}.Build()
	File_model_user_setting_proto = out.File
	file_model_user_setting_proto_goTypes = nil
	file_model_user_setting_proto_depIdxs = nil
}
