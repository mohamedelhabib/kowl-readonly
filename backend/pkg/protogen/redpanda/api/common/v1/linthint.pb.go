// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: redpanda/api/common/v1/linthint.proto

package commonv1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// LintHint is a generic linting hint.
type LintHint struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Line number of the lint.
	Line int32 `protobuf:"varint,1,opt,name=line,proto3" json:"line,omitempty"`
	// Column number of the lint.
	Column int32 `protobuf:"varint,2,opt,name=column,proto3" json:"column,omitempty"`
	// The hint message.
	Hint string `protobuf:"bytes,3,opt,name=hint,proto3" json:"hint,omitempty"`
	// Optional lint type or enum.
	LintType      string `protobuf:"bytes,4,opt,name=lint_type,json=lintType,proto3" json:"lint_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LintHint) Reset() {
	*x = LintHint{}
	mi := &file_redpanda_api_common_v1_linthint_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LintHint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LintHint) ProtoMessage() {}

func (x *LintHint) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_common_v1_linthint_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LintHint.ProtoReflect.Descriptor instead.
func (*LintHint) Descriptor() ([]byte, []int) {
	return file_redpanda_api_common_v1_linthint_proto_rawDescGZIP(), []int{0}
}

func (x *LintHint) GetLine() int32 {
	if x != nil {
		return x.Line
	}
	return 0
}

func (x *LintHint) GetColumn() int32 {
	if x != nil {
		return x.Column
	}
	return 0
}

func (x *LintHint) GetHint() string {
	if x != nil {
		return x.Hint
	}
	return ""
}

func (x *LintHint) GetLintType() string {
	if x != nil {
		return x.LintType
	}
	return ""
}

var File_redpanda_api_common_v1_linthint_proto protoreflect.FileDescriptor

var file_redpanda_api_common_v1_linthint_proto_rawDesc = []byte{
	0x0a, 0x25, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x6e, 0x74, 0x68, 0x69, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64,
	0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22,
	0x67, 0x0a, 0x08, 0x4c, 0x69, 0x6e, 0x74, 0x48, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x69, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6c,
	0x69, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x69, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x42, 0xff, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d,
	0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x4c, 0x69, 0x6e, 0x74, 0x68, 0x69, 0x6e,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x57, 0x62, 0x75, 0x66, 0x2e, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x64, 0x70, 0x61,
	0x6e, 0x64, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x67,
	0x6f, 0x2f, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x52, 0x41, 0x43, 0xaa, 0x02, 0x16, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e,
	0x64, 0x61, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x16, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x5c, 0x41, 0x70, 0x69, 0x5c,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x22, 0x52, 0x65, 0x64, 0x70,
	0x61, 0x6e, 0x64, 0x61, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x19, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_redpanda_api_common_v1_linthint_proto_rawDescOnce sync.Once
	file_redpanda_api_common_v1_linthint_proto_rawDescData = file_redpanda_api_common_v1_linthint_proto_rawDesc
)

func file_redpanda_api_common_v1_linthint_proto_rawDescGZIP() []byte {
	file_redpanda_api_common_v1_linthint_proto_rawDescOnce.Do(func() {
		file_redpanda_api_common_v1_linthint_proto_rawDescData = protoimpl.X.CompressGZIP(file_redpanda_api_common_v1_linthint_proto_rawDescData)
	})
	return file_redpanda_api_common_v1_linthint_proto_rawDescData
}

var file_redpanda_api_common_v1_linthint_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_redpanda_api_common_v1_linthint_proto_goTypes = []any{
	(*LintHint)(nil), // 0: redpanda.api.common.v1.LintHint
}
var file_redpanda_api_common_v1_linthint_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_redpanda_api_common_v1_linthint_proto_init() }
func file_redpanda_api_common_v1_linthint_proto_init() {
	if File_redpanda_api_common_v1_linthint_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_redpanda_api_common_v1_linthint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_redpanda_api_common_v1_linthint_proto_goTypes,
		DependencyIndexes: file_redpanda_api_common_v1_linthint_proto_depIdxs,
		MessageInfos:      file_redpanda_api_common_v1_linthint_proto_msgTypes,
	}.Build()
	File_redpanda_api_common_v1_linthint_proto = out.File
	file_redpanda_api_common_v1_linthint_proto_rawDesc = nil
	file_redpanda_api_common_v1_linthint_proto_goTypes = nil
	file_redpanda_api_common_v1_linthint_proto_depIdxs = nil
}
