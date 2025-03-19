// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: model/memo/resource.proto

package memo

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Resource struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The name of the resource.
	// Format: resources/{resource}, resource is the user defined if or uuid.
	Name         string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreateTime   *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Filename     string                 `protobuf:"bytes,4,opt,name=filename,proto3" json:"filename,omitempty"`
	Content      []byte                 `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	ExternalLink string                 `protobuf:"bytes,6,opt,name=external_link,json=externalLink,proto3" json:"external_link,omitempty"`
	Type         string                 `protobuf:"bytes,7,opt,name=type,proto3" json:"type,omitempty"`
	Size         int64                  `protobuf:"varint,8,opt,name=size,proto3" json:"size,omitempty"`
	// The related memo. Refer to `Memo.name`.
	Memo          *string `protobuf:"bytes,9,opt,name=memo,proto3,oneof" json:"memo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Resource) Reset() {
	*x = Resource{}
	mi := &file_model_memo_resource_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_model_memo_resource_proto_msgTypes[0]
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
	return file_model_memo_resource_proto_rawDescGZIP(), []int{0}
}

func (x *Resource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Resource) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Resource) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *Resource) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Resource) GetExternalLink() string {
	if x != nil {
		return x.ExternalLink
	}
	return ""
}

func (x *Resource) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Resource) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Resource) GetMemo() string {
	if x != nil && x.Memo != nil {
		return *x.Memo
	}
	return ""
}

var File_model_memo_resource_proto protoreflect.FileDescriptor

var file_model_memo_resource_proto_rawDesc = string([]byte{
	0x0a, 0x19, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x65, 0x6d,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x98, 0x02, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06,
	0xe0, 0x41, 0x03, 0xe0, 0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x03, 0xe0, 0x41, 0x04,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6d, 0x65, 0x6d, 0x6f, 0x88, 0x01, 0x01, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x6d, 0x65, 0x6d, 0x6f, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x42, 0x37,
	0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x73, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2f, 0x6d, 0x65, 0x6d, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_model_memo_resource_proto_rawDescOnce sync.Once
	file_model_memo_resource_proto_rawDescData []byte
)

func file_model_memo_resource_proto_rawDescGZIP() []byte {
	file_model_memo_resource_proto_rawDescOnce.Do(func() {
		file_model_memo_resource_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_model_memo_resource_proto_rawDesc), len(file_model_memo_resource_proto_rawDesc)))
	})
	return file_model_memo_resource_proto_rawDescData
}

var file_model_memo_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_model_memo_resource_proto_goTypes = []any{
	(*Resource)(nil),              // 0: memo.Resource
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_model_memo_resource_proto_depIdxs = []int32{
	1, // 0: memo.Resource.create_time:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_model_memo_resource_proto_init() }
func file_model_memo_resource_proto_init() {
	if File_model_memo_resource_proto != nil {
		return
	}
	file_model_memo_resource_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_model_memo_resource_proto_rawDesc), len(file_model_memo_resource_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_memo_resource_proto_goTypes,
		DependencyIndexes: file_model_memo_resource_proto_depIdxs,
		MessageInfos:      file_model_memo_resource_proto_msgTypes,
	}.Build()
	File_model_memo_resource_proto = out.File
	file_model_memo_resource_proto_goTypes = nil
	file_model_memo_resource_proto_depIdxs = nil
}
