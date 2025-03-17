// This file is a trick to force protoc-gen-openapiv2 into including the types used here into the openapi spec. They are not normally included, because they are not explicitly referenced in any proto (as protobuf ANY is used in errordetails).

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: redpanda/api/dataplane/v1alpha1/dummy.proto

package dataplanev1alpha1

import (
	reflect "reflect"
	sync "sync"

	errdetails "google.golang.org/genproto/googleapis/rpc/errdetails"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DummyMethodResponse struct {
	state                  protoimpl.MessageState   `protogen:"open.v1"`
	BadRequest             *errdetails.BadRequest   `protobuf:"bytes,1,opt,name=bad_request,json=badRequest,proto3" json:"bad_request,omitempty"`
	ErrorInfo              *errdetails.ErrorInfo    `protobuf:"bytes,2,opt,name=error_info,json=errorInfo,proto3" json:"error_info,omitempty"`
	QuotaFailure           *errdetails.QuotaFailure `protobuf:"bytes,3,opt,name=quota_failure,json=quotaFailure,proto3" json:"quota_failure,omitempty"`
	Help                   *errdetails.Help         `protobuf:"bytes,4,opt,name=help,proto3" json:"help,omitempty"`
	DeployTransformRequest *DeployTransformRequest  `protobuf:"bytes,5,opt,name=deploy_transform_request,json=deployTransformRequest,proto3" json:"deploy_transform_request,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *DummyMethodResponse) Reset() {
	*x = DummyMethodResponse{}
	mi := &file_redpanda_api_dataplane_v1alpha1_dummy_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DummyMethodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DummyMethodResponse) ProtoMessage() {}

func (x *DummyMethodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_dataplane_v1alpha1_dummy_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DummyMethodResponse.ProtoReflect.Descriptor instead.
func (*DummyMethodResponse) Descriptor() ([]byte, []int) {
	return file_redpanda_api_dataplane_v1alpha1_dummy_proto_rawDescGZIP(), []int{0}
}

func (x *DummyMethodResponse) GetBadRequest() *errdetails.BadRequest {
	if x != nil {
		return x.BadRequest
	}
	return nil
}

func (x *DummyMethodResponse) GetErrorInfo() *errdetails.ErrorInfo {
	if x != nil {
		return x.ErrorInfo
	}
	return nil
}

func (x *DummyMethodResponse) GetQuotaFailure() *errdetails.QuotaFailure {
	if x != nil {
		return x.QuotaFailure
	}
	return nil
}

func (x *DummyMethodResponse) GetHelp() *errdetails.Help {
	if x != nil {
		return x.Help
	}
	return nil
}

func (x *DummyMethodResponse) GetDeployTransformRequest() *DeployTransformRequest {
	if x != nil {
		return x.DeployTransformRequest
	}
	return nil
}

var File_redpanda_api_dataplane_v1alpha1_dummy_proto protoreflect.FileDescriptor

var file_redpanda_api_dataplane_v1alpha1_dummy_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x72,
	0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x72, 0x65, 0x64,
	0x70, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x02, 0x0a,
	0x13, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x62, 0x61, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x0a, 0x62, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a,
	0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x3d, 0x0a, 0x0d, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x5f, 0x66, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x46, 0x61, 0x69,
	0x6c, 0x75, 0x72, 0x65, 0x52, 0x0c, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x46, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x68, 0x65, 0x6c, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x65,
	0x6c, 0x70, 0x52, 0x04, 0x68, 0x65, 0x6c, 0x70, 0x12, 0x71, 0x0a, 0x18, 0x64, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x72, 0x65, 0x64,
	0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x16, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x6b, 0x0a, 0x0c, 0x44,
	0x75, 0x6d, 0x6d, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x0b, 0x44,
	0x75, 0x6d, 0x6d, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x34, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xb9, 0x02, 0x0a, 0x23, 0x63, 0x6f, 0x6d,
	0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x42, 0x0a, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x67,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x64, 0x70, 0x61,
	0x6e, 0x64, 0x61, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65,
	0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x41, 0x44, 0xaa, 0x02, 0x1f,
	0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca,
	0x02, 0x1f, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x44,
	0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0xe2, 0x02, 0x2b, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x5c, 0x41, 0x70, 0x69,
	0x5c, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x22, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a,
	0x3a, 0x44, 0x61, 0x74, 0x61, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_redpanda_api_dataplane_v1alpha1_dummy_proto_rawDescOnce sync.Once
	file_redpanda_api_dataplane_v1alpha1_dummy_proto_rawDescData = file_redpanda_api_dataplane_v1alpha1_dummy_proto_rawDesc
)

func file_redpanda_api_dataplane_v1alpha1_dummy_proto_rawDescGZIP() []byte {
	file_redpanda_api_dataplane_v1alpha1_dummy_proto_rawDescOnce.Do(func() {
		file_redpanda_api_dataplane_v1alpha1_dummy_proto_rawDescData = protoimpl.X.CompressGZIP(file_redpanda_api_dataplane_v1alpha1_dummy_proto_rawDescData)
	})
	return file_redpanda_api_dataplane_v1alpha1_dummy_proto_rawDescData
}

var file_redpanda_api_dataplane_v1alpha1_dummy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_redpanda_api_dataplane_v1alpha1_dummy_proto_goTypes = []any{
	(*DummyMethodResponse)(nil),     // 0: redpanda.api.dataplane.v1alpha1.DummyMethodResponse
	(*errdetails.BadRequest)(nil),   // 1: google.rpc.BadRequest
	(*errdetails.ErrorInfo)(nil),    // 2: google.rpc.ErrorInfo
	(*errdetails.QuotaFailure)(nil), // 3: google.rpc.QuotaFailure
	(*errdetails.Help)(nil),         // 4: google.rpc.Help
	(*DeployTransformRequest)(nil),  // 5: redpanda.api.dataplane.v1alpha1.DeployTransformRequest
	(*emptypb.Empty)(nil),           // 6: google.protobuf.Empty
}
var file_redpanda_api_dataplane_v1alpha1_dummy_proto_depIdxs = []int32{
	1, // 0: redpanda.api.dataplane.v1alpha1.DummyMethodResponse.bad_request:type_name -> google.rpc.BadRequest
	2, // 1: redpanda.api.dataplane.v1alpha1.DummyMethodResponse.error_info:type_name -> google.rpc.ErrorInfo
	3, // 2: redpanda.api.dataplane.v1alpha1.DummyMethodResponse.quota_failure:type_name -> google.rpc.QuotaFailure
	4, // 3: redpanda.api.dataplane.v1alpha1.DummyMethodResponse.help:type_name -> google.rpc.Help
	5, // 4: redpanda.api.dataplane.v1alpha1.DummyMethodResponse.deploy_transform_request:type_name -> redpanda.api.dataplane.v1alpha1.DeployTransformRequest
	6, // 5: redpanda.api.dataplane.v1alpha1.DummyService.DummyMethod:input_type -> google.protobuf.Empty
	0, // 6: redpanda.api.dataplane.v1alpha1.DummyService.DummyMethod:output_type -> redpanda.api.dataplane.v1alpha1.DummyMethodResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_redpanda_api_dataplane_v1alpha1_dummy_proto_init() }
func file_redpanda_api_dataplane_v1alpha1_dummy_proto_init() {
	if File_redpanda_api_dataplane_v1alpha1_dummy_proto != nil {
		return
	}
	file_redpanda_api_dataplane_v1alpha1_transform_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_redpanda_api_dataplane_v1alpha1_dummy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_redpanda_api_dataplane_v1alpha1_dummy_proto_goTypes,
		DependencyIndexes: file_redpanda_api_dataplane_v1alpha1_dummy_proto_depIdxs,
		MessageInfos:      file_redpanda_api_dataplane_v1alpha1_dummy_proto_msgTypes,
	}.Build()
	File_redpanda_api_dataplane_v1alpha1_dummy_proto = out.File
	file_redpanda_api_dataplane_v1alpha1_dummy_proto_rawDesc = nil
	file_redpanda_api_dataplane_v1alpha1_dummy_proto_goTypes = nil
	file_redpanda_api_dataplane_v1alpha1_dummy_proto_depIdxs = nil
}
