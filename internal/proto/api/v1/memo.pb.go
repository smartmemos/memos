// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: api/v1/memo.proto

package v1

import (
	memo "github.com/smartmemos/memos/internal/proto/model/memo"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
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

type CreateMemoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateMemoRequest) Reset() {
	*x = CreateMemoRequest{}
	mi := &file_api_v1_memo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateMemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMemoRequest) ProtoMessage() {}

func (x *CreateMemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_memo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMemoRequest.ProtoReflect.Descriptor instead.
func (*CreateMemoRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_memo_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMemoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListMemosRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListMemosRequest) Reset() {
	*x = ListMemosRequest{}
	mi := &file_api_v1_memo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListMemosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMemosRequest) ProtoMessage() {}

func (x *ListMemosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_memo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMemosRequest.ProtoReflect.Descriptor instead.
func (*ListMemosRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_memo_proto_rawDescGZIP(), []int{1}
}

type ListMemosResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Memos []*memo.Memo           `protobuf:"bytes,1,rep,name=memos,proto3" json:"memos,omitempty"`
	// A token, which can be sent as `page_token` to retrieve the next page.
	// If this field is omitted, there are no subsequent pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListMemosResponse) Reset() {
	*x = ListMemosResponse{}
	mi := &file_api_v1_memo_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListMemosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMemosResponse) ProtoMessage() {}

func (x *ListMemosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_memo_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMemosResponse.ProtoReflect.Descriptor instead.
func (*ListMemosResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_memo_proto_rawDescGZIP(), []int{2}
}

func (x *ListMemosResponse) GetMemos() []*memo.Memo {
	if x != nil {
		return x.Memos
	}
	return nil
}

func (x *ListMemosResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type GetMemoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMemoRequest) Reset() {
	*x = GetMemoRequest{}
	mi := &file_api_v1_memo_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMemoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMemoRequest) ProtoMessage() {}

func (x *GetMemoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_memo_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMemoRequest.ProtoReflect.Descriptor instead.
func (*GetMemoRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_memo_proto_rawDescGZIP(), []int{3}
}

var File_api_v1_memo_proto protoreflect.FileDescriptor

var file_api_v1_memo_proto_rawDesc = string([]byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x2f, 0x6d, 0x65,
	0x6d, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5d, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d,
	0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x6d, 0x65,
	0x6d, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6d, 0x65, 0x6d, 0x6f,
	0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x05, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x12, 0x26, 0x0a, 0x0f,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xfa, 0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x6d, 0x6f, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x6d, 0x6f, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0a, 0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x22, 0x15, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0f, 0x22, 0x0d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d,
	0x6f, 0x73, 0x12, 0x57, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x12,
	0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d,
	0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x12, 0x46, 0x0a, 0x07, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a,
	0x2e, 0x6d, 0x65, 0x6d, 0x6f, 0x2e, 0x4d, 0x65, 0x6d, 0x6f, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x6d, 0x6f,
	0x73, 0x2f, 0x2a, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x6d,
	0x6f, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_v1_memo_proto_rawDescOnce sync.Once
	file_api_v1_memo_proto_rawDescData []byte
)

func file_api_v1_memo_proto_rawDescGZIP() []byte {
	file_api_v1_memo_proto_rawDescOnce.Do(func() {
		file_api_v1_memo_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1_memo_proto_rawDesc), len(file_api_v1_memo_proto_rawDesc)))
	})
	return file_api_v1_memo_proto_rawDescData
}

var file_api_v1_memo_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_v1_memo_proto_goTypes = []any{
	(*CreateMemoRequest)(nil), // 0: api.v1.CreateMemoRequest
	(*ListMemosRequest)(nil),  // 1: api.v1.ListMemosRequest
	(*ListMemosResponse)(nil), // 2: api.v1.ListMemosResponse
	(*GetMemoRequest)(nil),    // 3: api.v1.GetMemoRequest
	(*memo.Memo)(nil),         // 4: memo.Memo
}
var file_api_v1_memo_proto_depIdxs = []int32{
	4, // 0: api.v1.ListMemosResponse.memos:type_name -> memo.Memo
	0, // 1: api.v1.MemoService.CreateMemo:input_type -> api.v1.CreateMemoRequest
	1, // 2: api.v1.MemoService.ListMemos:input_type -> api.v1.ListMemosRequest
	3, // 3: api.v1.MemoService.GetMemo:input_type -> api.v1.GetMemoRequest
	4, // 4: api.v1.MemoService.CreateMemo:output_type -> memo.Memo
	2, // 5: api.v1.MemoService.ListMemos:output_type -> api.v1.ListMemosResponse
	4, // 6: api.v1.MemoService.GetMemo:output_type -> memo.Memo
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_v1_memo_proto_init() }
func file_api_v1_memo_proto_init() {
	if File_api_v1_memo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1_memo_proto_rawDesc), len(file_api_v1_memo_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_memo_proto_goTypes,
		DependencyIndexes: file_api_v1_memo_proto_depIdxs,
		MessageInfos:      file_api_v1_memo_proto_msgTypes,
	}.Build()
	File_api_v1_memo_proto = out.File
	file_api_v1_memo_proto_goTypes = nil
	file_api_v1_memo_proto_depIdxs = nil
}
