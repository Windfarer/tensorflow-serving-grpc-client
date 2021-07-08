// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: tensorflow/core/protobuf/snapshot.proto

package tensorflow

import (
	proto "github.com/golang/protobuf/proto"
	tensor_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_go_proto"
	tensor_shape_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto"
	types_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Each SnapshotRecord represents one batch of pre-processed input data. A batch
// consists of a list of tensors that we encode as TensorProtos. This message
// doesn't store the structure of the batch.
type SnapshotRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tensor []*tensor_go_proto.TensorProto `protobuf:"bytes,1,rep,name=tensor,proto3" json:"tensor,omitempty"`
}

func (x *SnapshotRecord) Reset() {
	*x = SnapshotRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_snapshot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotRecord) ProtoMessage() {}

func (x *SnapshotRecord) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_snapshot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotRecord.ProtoReflect.Descriptor instead.
func (*SnapshotRecord) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_snapshot_proto_rawDescGZIP(), []int{0}
}

func (x *SnapshotRecord) GetTensor() []*tensor_go_proto.TensorProto {
	if x != nil {
		return x.Tensor
	}
	return nil
}

// This stores the metadata information present in each snapshot record.
type SnapshotMetadataRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Stores the fingerprint of the graph that describes the dataset that is
	// snapshotted.
	GraphHash string `protobuf:"bytes,1,opt,name=graph_hash,json=graphHash,proto3" json:"graph_hash,omitempty"`
	// Run ID that this snapshot corresponds to.
	RunId string `protobuf:"bytes,2,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
	// Time when we started creating this snapshot.
	CreationTimestamp int64 `protobuf:"varint,3,opt,name=creation_timestamp,json=creationTimestamp,proto3" json:"creation_timestamp,omitempty"`
	// Version of the snapshot data file format.
	Version int64 `protobuf:"varint,4,opt,name=version,proto3" json:"version,omitempty"`
	// A list of tensor dtype corresponding to each element of the snapshot.
	Dtype []types_go_proto.DataType `protobuf:"varint,5,rep,packed,name=dtype,proto3,enum=tensorflow.DataType" json:"dtype,omitempty"`
	// The number of elements in the snapshot.
	NumElements int64 `protobuf:"varint,6,opt,name=num_elements,json=numElements,proto3" json:"num_elements,omitempty"`
	Finalized   bool  `protobuf:"varint,1000,opt,name=finalized,proto3" json:"finalized,omitempty"`
}

func (x *SnapshotMetadataRecord) Reset() {
	*x = SnapshotMetadataRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_snapshot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotMetadataRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotMetadataRecord) ProtoMessage() {}

func (x *SnapshotMetadataRecord) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_snapshot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotMetadataRecord.ProtoReflect.Descriptor instead.
func (*SnapshotMetadataRecord) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_snapshot_proto_rawDescGZIP(), []int{1}
}

func (x *SnapshotMetadataRecord) GetGraphHash() string {
	if x != nil {
		return x.GraphHash
	}
	return ""
}

func (x *SnapshotMetadataRecord) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

func (x *SnapshotMetadataRecord) GetCreationTimestamp() int64 {
	if x != nil {
		return x.CreationTimestamp
	}
	return 0
}

func (x *SnapshotMetadataRecord) GetVersion() int64 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *SnapshotMetadataRecord) GetDtype() []types_go_proto.DataType {
	if x != nil {
		return x.Dtype
	}
	return nil
}

func (x *SnapshotMetadataRecord) GetNumElements() int64 {
	if x != nil {
		return x.NumElements
	}
	return 0
}

func (x *SnapshotMetadataRecord) GetFinalized() bool {
	if x != nil {
		return x.Finalized
	}
	return false
}

// Metadata for a single tensor in the Snapshot Record.
type TensorMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TensorShape *tensor_shape_go_proto.TensorShapeProto `protobuf:"bytes,2,opt,name=tensor_shape,json=tensorShape,proto3" json:"tensor_shape,omitempty"`
	// Number of uncompressed bytes used to store the tensor representation.
	TensorSizeBytes int64 `protobuf:"varint,3,opt,name=tensor_size_bytes,json=tensorSizeBytes,proto3" json:"tensor_size_bytes,omitempty"`
}

func (x *TensorMetadata) Reset() {
	*x = TensorMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_snapshot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TensorMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TensorMetadata) ProtoMessage() {}

func (x *TensorMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_snapshot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TensorMetadata.ProtoReflect.Descriptor instead.
func (*TensorMetadata) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_snapshot_proto_rawDescGZIP(), []int{2}
}

func (x *TensorMetadata) GetTensorShape() *tensor_shape_go_proto.TensorShapeProto {
	if x != nil {
		return x.TensorShape
	}
	return nil
}

func (x *TensorMetadata) GetTensorSizeBytes() int64 {
	if x != nil {
		return x.TensorSizeBytes
	}
	return 0
}

// Metadata for all the tensors in a Snapshot Record.
type SnapshotTensorMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TensorMetadata []*TensorMetadata `protobuf:"bytes,1,rep,name=tensor_metadata,json=tensorMetadata,proto3" json:"tensor_metadata,omitempty"`
}

func (x *SnapshotTensorMetadata) Reset() {
	*x = SnapshotTensorMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_snapshot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotTensorMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotTensorMetadata) ProtoMessage() {}

func (x *SnapshotTensorMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_snapshot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotTensorMetadata.ProtoReflect.Descriptor instead.
func (*SnapshotTensorMetadata) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_snapshot_proto_rawDescGZIP(), []int{3}
}

func (x *SnapshotTensorMetadata) GetTensorMetadata() []*TensorMetadata {
	if x != nil {
		return x.TensorMetadata
	}
	return nil
}

var File_tensorflow_core_protobuf_snapshot_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_snapshot_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x26, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f,
	0x73, 0x68, 0x61, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x41, 0x0a, 0x0e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x52, 0x06, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x22, 0x85, 0x02, 0x0a, 0x16, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x61, 0x70, 0x68, 0x48, 0x61, 0x73, 0x68, 0x12,
	0x15, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x72, 0x75, 0x6e, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x11, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x2a, 0x0a, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x14,
	0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x64, 0x74, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6e,
	0x75, 0x6d, 0x5f, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x6e, 0x75, 0x6d, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1d,
	0x0a, 0x09, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x18, 0xe8, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x22, 0x7d, 0x0a,
	0x0e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x3f, 0x0a, 0x0c, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x68, 0x61, 0x70, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x52, 0x0b, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x53, 0x68, 0x61, 0x70, 0x65,
	0x12, 0x2a, 0x0a, 0x11, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x5d, 0x0a, 0x16,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x43, 0x0a, 0x0f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x54, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x0e, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_snapshot_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_snapshot_proto_rawDescData = file_tensorflow_core_protobuf_snapshot_proto_rawDesc
)

func file_tensorflow_core_protobuf_snapshot_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_snapshot_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_snapshot_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_snapshot_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_snapshot_proto_rawDescData
}

var file_tensorflow_core_protobuf_snapshot_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_tensorflow_core_protobuf_snapshot_proto_goTypes = []interface{}{
	(*SnapshotRecord)(nil),                         // 0: tensorflow.SnapshotRecord
	(*SnapshotMetadataRecord)(nil),                 // 1: tensorflow.SnapshotMetadataRecord
	(*TensorMetadata)(nil),                         // 2: tensorflow.TensorMetadata
	(*SnapshotTensorMetadata)(nil),                 // 3: tensorflow.SnapshotTensorMetadata
	(*tensor_go_proto.TensorProto)(nil),            // 4: tensorflow.TensorProto
	(types_go_proto.DataType)(0),                   // 5: tensorflow.DataType
	(*tensor_shape_go_proto.TensorShapeProto)(nil), // 6: tensorflow.TensorShapeProto
}
var file_tensorflow_core_protobuf_snapshot_proto_depIdxs = []int32{
	4, // 0: tensorflow.SnapshotRecord.tensor:type_name -> tensorflow.TensorProto
	5, // 1: tensorflow.SnapshotMetadataRecord.dtype:type_name -> tensorflow.DataType
	6, // 2: tensorflow.TensorMetadata.tensor_shape:type_name -> tensorflow.TensorShapeProto
	2, // 3: tensorflow.SnapshotTensorMetadata.tensor_metadata:type_name -> tensorflow.TensorMetadata
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_snapshot_proto_init() }
func file_tensorflow_core_protobuf_snapshot_proto_init() {
	if File_tensorflow_core_protobuf_snapshot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_snapshot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotRecord); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorflow_core_protobuf_snapshot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotMetadataRecord); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorflow_core_protobuf_snapshot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TensorMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_tensorflow_core_protobuf_snapshot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotTensorMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_core_protobuf_snapshot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_snapshot_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_snapshot_proto_depIdxs,
		MessageInfos:      file_tensorflow_core_protobuf_snapshot_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_snapshot_proto = out.File
	file_tensorflow_core_protobuf_snapshot_proto_rawDesc = nil
	file_tensorflow_core_protobuf_snapshot_proto_goTypes = nil
	file_tensorflow_core_protobuf_snapshot_proto_depIdxs = nil
}
