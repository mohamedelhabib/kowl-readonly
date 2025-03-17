// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        (unknown)
// source: redpanda/api/console/v1alpha1/publish_messages.proto

package consolev1alpha1

import (
	reflect "reflect"
	sync "sync"

	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PublishMessageRequest is the request for PublishMessage call.
type PublishMessageRequest struct {
	state           protoimpl.MessageState        `protogen:"open.v1"`
	Topic           string                        `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`                                                                 // The topics to publish to.
	PartitionId     int32                         `protobuf:"varint,2,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`                                 // -1 for automatic partition assignment.
	Compression     CompressionType               `protobuf:"varint,3,opt,name=compression,proto3,enum=redpanda.api.console.v1alpha1.CompressionType" json:"compression,omitempty"` // The compression to be used.
	UseTransactions bool                          `protobuf:"varint,4,opt,name=use_transactions,json=useTransactions,proto3" json:"use_transactions,omitempty"`                     // Use transactions.
	Headers         []*KafkaRecordHeader          `protobuf:"bytes,5,rep,name=headers,proto3" json:"headers,omitempty"`                                                             // Kafka record headers.
	Key             *PublishMessagePayloadOptions `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
	Value           *PublishMessagePayloadOptions `protobuf:"bytes,7,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *PublishMessageRequest) Reset() {
	*x = PublishMessageRequest{}
	mi := &file_redpanda_api_console_v1alpha1_publish_messages_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishMessageRequest) ProtoMessage() {}

func (x *PublishMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_console_v1alpha1_publish_messages_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishMessageRequest.ProtoReflect.Descriptor instead.
func (*PublishMessageRequest) Descriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_publish_messages_proto_rawDescGZIP(), []int{0}
}

func (x *PublishMessageRequest) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *PublishMessageRequest) GetPartitionId() int32 {
	if x != nil {
		return x.PartitionId
	}
	return 0
}

func (x *PublishMessageRequest) GetCompression() CompressionType {
	if x != nil {
		return x.Compression
	}
	return CompressionType_COMPRESSION_TYPE_UNSPECIFIED
}

func (x *PublishMessageRequest) GetUseTransactions() bool {
	if x != nil {
		return x.UseTransactions
	}
	return false
}

func (x *PublishMessageRequest) GetHeaders() []*KafkaRecordHeader {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *PublishMessageRequest) GetKey() *PublishMessagePayloadOptions {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *PublishMessageRequest) GetValue() *PublishMessagePayloadOptions {
	if x != nil {
		return x.Value
	}
	return nil
}

type PublishMessagePayloadOptions struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Encoding      PayloadEncoding        `protobuf:"varint,1,opt,name=encoding,proto3,enum=redpanda.api.console.v1alpha1.PayloadEncoding" json:"encoding,omitempty"` // Payload encoding to use.
	Data          []byte                 `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`                                                             // Data.
	SchemaId      *int32                 `protobuf:"varint,9,opt,name=schema_id,json=schemaId,proto3,oneof" json:"schema_id,omitempty"`                              // Optional schema ID.
	Index         *int32                 `protobuf:"varint,10,opt,name=index,proto3,oneof" json:"index,omitempty"`                                                   // Optional index. Useful for Protobuf messages.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublishMessagePayloadOptions) Reset() {
	*x = PublishMessagePayloadOptions{}
	mi := &file_redpanda_api_console_v1alpha1_publish_messages_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishMessagePayloadOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishMessagePayloadOptions) ProtoMessage() {}

func (x *PublishMessagePayloadOptions) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_console_v1alpha1_publish_messages_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishMessagePayloadOptions.ProtoReflect.Descriptor instead.
func (*PublishMessagePayloadOptions) Descriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_publish_messages_proto_rawDescGZIP(), []int{1}
}

func (x *PublishMessagePayloadOptions) GetEncoding() PayloadEncoding {
	if x != nil {
		return x.Encoding
	}
	return PayloadEncoding_PAYLOAD_ENCODING_UNSPECIFIED
}

func (x *PublishMessagePayloadOptions) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PublishMessagePayloadOptions) GetSchemaId() int32 {
	if x != nil && x.SchemaId != nil {
		return *x.SchemaId
	}
	return 0
}

func (x *PublishMessagePayloadOptions) GetIndex() int32 {
	if x != nil && x.Index != nil {
		return *x.Index
	}
	return 0
}

// PublishMessageResponse is the response for PublishMessage call.
type PublishMessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Topic         string                 `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	PartitionId   int32                  `protobuf:"varint,2,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`
	Offset        int64                  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublishMessageResponse) Reset() {
	*x = PublishMessageResponse{}
	mi := &file_redpanda_api_console_v1alpha1_publish_messages_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublishMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishMessageResponse) ProtoMessage() {}

func (x *PublishMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_redpanda_api_console_v1alpha1_publish_messages_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishMessageResponse.ProtoReflect.Descriptor instead.
func (*PublishMessageResponse) Descriptor() ([]byte, []int) {
	return file_redpanda_api_console_v1alpha1_publish_messages_proto_rawDescGZIP(), []int{2}
}

func (x *PublishMessageResponse) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *PublishMessageResponse) GetPartitionId() int32 {
	if x != nil {
		return x.PartitionId
	}
	return 0
}

func (x *PublishMessageResponse) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_redpanda_api_console_v1alpha1_publish_messages_proto protoreflect.FileDescriptor

var file_redpanda_api_console_v1alpha1_publish_messages_proto_rawDesc = []byte{
	0x0a, 0x34, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2a, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed,
	0x03, 0x0a, 0x15, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e, 0xba, 0x48, 0x1b, 0x72, 0x19, 0x10, 0x01,
	0x18, 0xf9, 0x01, 0x32, 0x12, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39,
	0x2e, 0x5f, 0x5c, 0x2d, 0x5d, 0x2a, 0x24, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x33,
	0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x10, 0xba, 0x48, 0x0d, 0x1a, 0x0b, 0x28, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x50, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61,
	0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x75, 0x73, 0x65, 0x5f, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0f, 0x75, 0x73, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x4a, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x30, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x4d, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x72, 0x65, 0x64, 0x70,
	0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73,
	0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x51, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x72, 0x65, 0x64,
	0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xd3,
	0x01, 0x0a, 0x1c, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x4a, 0x0a, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2e, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x20, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x05, 0x48, 0x00, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0x88, 0x01,
	0x01, 0x12, 0x19, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x01, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a,
	0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x69, 0x64, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x22, 0x69, 0x0a, 0x16, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x42,
	0xb5, 0x02, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x14, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x63, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e,
	0x64, 0x61, 0x2d, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x67, 0x65, 0x6e, 0x2f, 0x72, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x41, 0x43, 0xaa, 0x02, 0x1d, 0x52, 0x65, 0x64, 0x70, 0x61,
	0x6e, 0x64, 0x61, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x2e,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x1d, 0x52, 0x65, 0x64, 0x70, 0x61,
	0x6e, 0x64, 0x61, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x5c,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x29, 0x52, 0x65, 0x64, 0x70, 0x61,
	0x6e, 0x64, 0x61, 0x5c, 0x41, 0x70, 0x69, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x5c,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x20, 0x52, 0x65, 0x64, 0x70, 0x61, 0x6e, 0x64, 0x61, 0x3a,
	0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x3a, 0x3a, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_redpanda_api_console_v1alpha1_publish_messages_proto_rawDescOnce sync.Once
	file_redpanda_api_console_v1alpha1_publish_messages_proto_rawDescData = file_redpanda_api_console_v1alpha1_publish_messages_proto_rawDesc
)

func file_redpanda_api_console_v1alpha1_publish_messages_proto_rawDescGZIP() []byte {
	file_redpanda_api_console_v1alpha1_publish_messages_proto_rawDescOnce.Do(func() {
		file_redpanda_api_console_v1alpha1_publish_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_redpanda_api_console_v1alpha1_publish_messages_proto_rawDescData)
	})
	return file_redpanda_api_console_v1alpha1_publish_messages_proto_rawDescData
}

var file_redpanda_api_console_v1alpha1_publish_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_redpanda_api_console_v1alpha1_publish_messages_proto_goTypes = []any{
	(*PublishMessageRequest)(nil),        // 0: redpanda.api.console.v1alpha1.PublishMessageRequest
	(*PublishMessagePayloadOptions)(nil), // 1: redpanda.api.console.v1alpha1.PublishMessagePayloadOptions
	(*PublishMessageResponse)(nil),       // 2: redpanda.api.console.v1alpha1.PublishMessageResponse
	(CompressionType)(0),                 // 3: redpanda.api.console.v1alpha1.CompressionType
	(*KafkaRecordHeader)(nil),            // 4: redpanda.api.console.v1alpha1.KafkaRecordHeader
	(PayloadEncoding)(0),                 // 5: redpanda.api.console.v1alpha1.PayloadEncoding
}
var file_redpanda_api_console_v1alpha1_publish_messages_proto_depIdxs = []int32{
	3, // 0: redpanda.api.console.v1alpha1.PublishMessageRequest.compression:type_name -> redpanda.api.console.v1alpha1.CompressionType
	4, // 1: redpanda.api.console.v1alpha1.PublishMessageRequest.headers:type_name -> redpanda.api.console.v1alpha1.KafkaRecordHeader
	1, // 2: redpanda.api.console.v1alpha1.PublishMessageRequest.key:type_name -> redpanda.api.console.v1alpha1.PublishMessagePayloadOptions
	1, // 3: redpanda.api.console.v1alpha1.PublishMessageRequest.value:type_name -> redpanda.api.console.v1alpha1.PublishMessagePayloadOptions
	5, // 4: redpanda.api.console.v1alpha1.PublishMessagePayloadOptions.encoding:type_name -> redpanda.api.console.v1alpha1.PayloadEncoding
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_redpanda_api_console_v1alpha1_publish_messages_proto_init() }
func file_redpanda_api_console_v1alpha1_publish_messages_proto_init() {
	if File_redpanda_api_console_v1alpha1_publish_messages_proto != nil {
		return
	}
	file_redpanda_api_console_v1alpha1_common_proto_init()
	file_redpanda_api_console_v1alpha1_publish_messages_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_redpanda_api_console_v1alpha1_publish_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_redpanda_api_console_v1alpha1_publish_messages_proto_goTypes,
		DependencyIndexes: file_redpanda_api_console_v1alpha1_publish_messages_proto_depIdxs,
		MessageInfos:      file_redpanda_api_console_v1alpha1_publish_messages_proto_msgTypes,
	}.Build()
	File_redpanda_api_console_v1alpha1_publish_messages_proto = out.File
	file_redpanda_api_console_v1alpha1_publish_messages_proto_rawDesc = nil
	file_redpanda_api_console_v1alpha1_publish_messages_proto_goTypes = nil
	file_redpanda_api_console_v1alpha1_publish_messages_proto_depIdxs = nil
}
