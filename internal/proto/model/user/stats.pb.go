// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: model/user/stats.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type Stats struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The timestamps when the memos were displayed.
	// We should return raw data to the client, and let the client format the data with the user's timezone.
	MemoDisplayTimestamps []*timestamppb.Timestamp `protobuf:"bytes,2,rep,name=memo_display_timestamps,json=memoDisplayTimestamps,proto3" json:"memo_display_timestamps,omitempty"`
	// The stats of memo types.
	MemoTypeStats *Stats_MemoTypeStats `protobuf:"bytes,3,opt,name=memo_type_stats,json=memoTypeStats,proto3" json:"memo_type_stats,omitempty"`
	// The count of tags.
	// Format: "tag1": 1, "tag2": 2
	TagCount map[string]int32 `protobuf:"bytes,4,rep,name=tag_count,json=tagCount,proto3" json:"tag_count,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	// The pinned memos of the user.
	PinnedMemos    []string `protobuf:"bytes,5,rep,name=pinned_memos,json=pinnedMemos,proto3" json:"pinned_memos,omitempty"`
	TotalMemoCount int32    `protobuf:"varint,6,opt,name=total_memo_count,json=totalMemoCount,proto3" json:"total_memo_count,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Stats) Reset() {
	*x = Stats{}
	mi := &file_model_user_stats_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stats) ProtoMessage() {}

func (x *Stats) ProtoReflect() protoreflect.Message {
	mi := &file_model_user_stats_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stats.ProtoReflect.Descriptor instead.
func (*Stats) Descriptor() ([]byte, []int) {
	return file_model_user_stats_proto_rawDescGZIP(), []int{0}
}

func (x *Stats) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Stats) GetMemoDisplayTimestamps() []*timestamppb.Timestamp {
	if x != nil {
		return x.MemoDisplayTimestamps
	}
	return nil
}

func (x *Stats) GetMemoTypeStats() *Stats_MemoTypeStats {
	if x != nil {
		return x.MemoTypeStats
	}
	return nil
}

func (x *Stats) GetTagCount() map[string]int32 {
	if x != nil {
		return x.TagCount
	}
	return nil
}

func (x *Stats) GetPinnedMemos() []string {
	if x != nil {
		return x.PinnedMemos
	}
	return nil
}

func (x *Stats) GetTotalMemoCount() int32 {
	if x != nil {
		return x.TotalMemoCount
	}
	return 0
}

type Stats_MemoTypeStats struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	LinkCount     int32                  `protobuf:"varint,1,opt,name=link_count,json=linkCount,proto3" json:"link_count,omitempty"`
	CodeCount     int32                  `protobuf:"varint,2,opt,name=code_count,json=codeCount,proto3" json:"code_count,omitempty"`
	TodoCount     int32                  `protobuf:"varint,3,opt,name=todo_count,json=todoCount,proto3" json:"todo_count,omitempty"`
	UndoCount     int32                  `protobuf:"varint,4,opt,name=undo_count,json=undoCount,proto3" json:"undo_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Stats_MemoTypeStats) Reset() {
	*x = Stats_MemoTypeStats{}
	mi := &file_model_user_stats_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Stats_MemoTypeStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stats_MemoTypeStats) ProtoMessage() {}

func (x *Stats_MemoTypeStats) ProtoReflect() protoreflect.Message {
	mi := &file_model_user_stats_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stats_MemoTypeStats.ProtoReflect.Descriptor instead.
func (*Stats_MemoTypeStats) Descriptor() ([]byte, []int) {
	return file_model_user_stats_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Stats_MemoTypeStats) GetLinkCount() int32 {
	if x != nil {
		return x.LinkCount
	}
	return 0
}

func (x *Stats_MemoTypeStats) GetCodeCount() int32 {
	if x != nil {
		return x.CodeCount
	}
	return 0
}

func (x *Stats_MemoTypeStats) GetTodoCount() int32 {
	if x != nil {
		return x.TodoCount
	}
	return 0
}

func (x *Stats_MemoTypeStats) GetUndoCount() int32 {
	if x != nil {
		return x.UndoCount
	}
	return 0
}

var File_model_user_stats_proto protoreflect.FileDescriptor

const file_model_user_stats_proto_rawDesc = "" +
	"\n" +
	"\x16model/user/stats.proto\x12\x04user\x1a\x1fgoogle/protobuf/timestamp.proto\"\x82\x04\n" +
	"\x05Stats\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12R\n" +
	"\x17memo_display_timestamps\x18\x02 \x03(\v2\x1a.google.protobuf.TimestampR\x15memoDisplayTimestamps\x12A\n" +
	"\x0fmemo_type_stats\x18\x03 \x01(\v2\x19.user.Stats.MemoTypeStatsR\rmemoTypeStats\x126\n" +
	"\ttag_count\x18\x04 \x03(\v2\x19.user.Stats.TagCountEntryR\btagCount\x12!\n" +
	"\fpinned_memos\x18\x05 \x03(\tR\vpinnedMemos\x12(\n" +
	"\x10total_memo_count\x18\x06 \x01(\x05R\x0etotalMemoCount\x1a;\n" +
	"\rTagCountEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\x05R\x05value:\x028\x01\x1a\x8b\x01\n" +
	"\rMemoTypeStats\x12\x1d\n" +
	"\n" +
	"link_count\x18\x01 \x01(\x05R\tlinkCount\x12\x1d\n" +
	"\n" +
	"code_count\x18\x02 \x01(\x05R\tcodeCount\x12\x1d\n" +
	"\n" +
	"todo_count\x18\x03 \x01(\x05R\ttodoCount\x12\x1d\n" +
	"\n" +
	"undo_count\x18\x04 \x01(\x05R\tundoCountB7Z5github.com/smartmemos/memos/internal/proto/model/userb\x06proto3"

var (
	file_model_user_stats_proto_rawDescOnce sync.Once
	file_model_user_stats_proto_rawDescData []byte
)

func file_model_user_stats_proto_rawDescGZIP() []byte {
	file_model_user_stats_proto_rawDescOnce.Do(func() {
		file_model_user_stats_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_model_user_stats_proto_rawDesc), len(file_model_user_stats_proto_rawDesc)))
	})
	return file_model_user_stats_proto_rawDescData
}

var file_model_user_stats_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_model_user_stats_proto_goTypes = []any{
	(*Stats)(nil),                 // 0: user.Stats
	nil,                           // 1: user.Stats.TagCountEntry
	(*Stats_MemoTypeStats)(nil),   // 2: user.Stats.MemoTypeStats
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_model_user_stats_proto_depIdxs = []int32{
	3, // 0: user.Stats.memo_display_timestamps:type_name -> google.protobuf.Timestamp
	2, // 1: user.Stats.memo_type_stats:type_name -> user.Stats.MemoTypeStats
	1, // 2: user.Stats.tag_count:type_name -> user.Stats.TagCountEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_model_user_stats_proto_init() }
func file_model_user_stats_proto_init() {
	if File_model_user_stats_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_model_user_stats_proto_rawDesc), len(file_model_user_stats_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_user_stats_proto_goTypes,
		DependencyIndexes: file_model_user_stats_proto_depIdxs,
		MessageInfos:      file_model_user_stats_proto_msgTypes,
	}.Build()
	File_model_user_stats_proto = out.File
	file_model_user_stats_proto_goTypes = nil
	file_model_user_stats_proto_depIdxs = nil
}
