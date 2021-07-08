// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: tensorflow/core/protobuf/saver.proto

package for_core_protos_go_proto

import (
	proto "github.com/golang/protobuf/proto"
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

// A version number that identifies a different on-disk checkpoint format.
// Usually, each subclass of BaseSaverBuilder works with a particular
// version/format.  However, it is possible that the same builder may be
// upgraded to support a newer checkpoint format in the future.
type SaverDef_CheckpointFormatVersion int32

const (
	// Internal legacy format.
	SaverDef_LEGACY SaverDef_CheckpointFormatVersion = 0
	// Deprecated format: tf.Saver() which works with tensorflow::table::Table.
	SaverDef_V1 SaverDef_CheckpointFormatVersion = 1
	// Current format: more efficient.
	SaverDef_V2 SaverDef_CheckpointFormatVersion = 2
)

// Enum value maps for SaverDef_CheckpointFormatVersion.
var (
	SaverDef_CheckpointFormatVersion_name = map[int32]string{
		0: "LEGACY",
		1: "V1",
		2: "V2",
	}
	SaverDef_CheckpointFormatVersion_value = map[string]int32{
		"LEGACY": 0,
		"V1":     1,
		"V2":     2,
	}
)

func (x SaverDef_CheckpointFormatVersion) Enum() *SaverDef_CheckpointFormatVersion {
	p := new(SaverDef_CheckpointFormatVersion)
	*p = x
	return p
}

func (x SaverDef_CheckpointFormatVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SaverDef_CheckpointFormatVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_tensorflow_core_protobuf_saver_proto_enumTypes[0].Descriptor()
}

func (SaverDef_CheckpointFormatVersion) Type() protoreflect.EnumType {
	return &file_tensorflow_core_protobuf_saver_proto_enumTypes[0]
}

func (x SaverDef_CheckpointFormatVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SaverDef_CheckpointFormatVersion.Descriptor instead.
func (SaverDef_CheckpointFormatVersion) EnumDescriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_saver_proto_rawDescGZIP(), []int{0, 0}
}

// Protocol buffer representing the configuration of a Saver.
type SaverDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the tensor in which to specify the filename when saving or
	// restoring a model checkpoint.
	FilenameTensorName string `protobuf:"bytes,1,opt,name=filename_tensor_name,json=filenameTensorName,proto3" json:"filename_tensor_name,omitempty"`
	// The operation to run when saving a model checkpoint.
	SaveTensorName string `protobuf:"bytes,2,opt,name=save_tensor_name,json=saveTensorName,proto3" json:"save_tensor_name,omitempty"`
	// The operation to run when restoring a model checkpoint.
	RestoreOpName string `protobuf:"bytes,3,opt,name=restore_op_name,json=restoreOpName,proto3" json:"restore_op_name,omitempty"`
	// Maximum number of checkpoints to keep.  If 0, no checkpoints are deleted.
	MaxToKeep int32 `protobuf:"varint,4,opt,name=max_to_keep,json=maxToKeep,proto3" json:"max_to_keep,omitempty"`
	// Shard the save files, one per device that has Variable nodes.
	Sharded bool `protobuf:"varint,5,opt,name=sharded,proto3" json:"sharded,omitempty"`
	// How often to keep an additional checkpoint. If not specified, only the last
	// "max_to_keep" checkpoints are kept; if specified, in addition to keeping
	// the last "max_to_keep" checkpoints, an additional checkpoint will be kept
	// for every n hours of training.
	KeepCheckpointEveryNHours float32                          `protobuf:"fixed32,6,opt,name=keep_checkpoint_every_n_hours,json=keepCheckpointEveryNHours,proto3" json:"keep_checkpoint_every_n_hours,omitempty"`
	Version                   SaverDef_CheckpointFormatVersion `protobuf:"varint,7,opt,name=version,proto3,enum=tensorflow.SaverDef_CheckpointFormatVersion" json:"version,omitempty"`
}

func (x *SaverDef) Reset() {
	*x = SaverDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_core_protobuf_saver_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaverDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaverDef) ProtoMessage() {}

func (x *SaverDef) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_core_protobuf_saver_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaverDef.ProtoReflect.Descriptor instead.
func (*SaverDef) Descriptor() ([]byte, []int) {
	return file_tensorflow_core_protobuf_saver_proto_rawDescGZIP(), []int{0}
}

func (x *SaverDef) GetFilenameTensorName() string {
	if x != nil {
		return x.FilenameTensorName
	}
	return ""
}

func (x *SaverDef) GetSaveTensorName() string {
	if x != nil {
		return x.SaveTensorName
	}
	return ""
}

func (x *SaverDef) GetRestoreOpName() string {
	if x != nil {
		return x.RestoreOpName
	}
	return ""
}

func (x *SaverDef) GetMaxToKeep() int32 {
	if x != nil {
		return x.MaxToKeep
	}
	return 0
}

func (x *SaverDef) GetSharded() bool {
	if x != nil {
		return x.Sharded
	}
	return false
}

func (x *SaverDef) GetKeepCheckpointEveryNHours() float32 {
	if x != nil {
		return x.KeepCheckpointEveryNHours
	}
	return 0
}

func (x *SaverDef) GetVersion() SaverDef_CheckpointFormatVersion {
	if x != nil {
		return x.Version
	}
	return SaverDef_LEGACY
}

var File_tensorflow_core_protobuf_saver_proto protoreflect.FileDescriptor

var file_tensorflow_core_protobuf_saver_proto_rawDesc = []byte{
	0x0a, 0x24, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x22, 0x89, 0x03, 0x0a, 0x08, 0x53, 0x61, 0x76, 0x65, 0x72, 0x44, 0x65, 0x66, 0x12,
	0x30, 0x0a, 0x14, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x61, 0x76, 0x65, 0x5f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x61, 0x76,
	0x65, 0x54, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x72,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x6f, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4f, 0x70, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x6f, 0x5f, 0x6b, 0x65,
	0x65, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x54, 0x6f, 0x4b,
	0x65, 0x65, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x68, 0x61, 0x72, 0x64, 0x65, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x68, 0x61, 0x72, 0x64, 0x65, 0x64, 0x12, 0x40, 0x0a,
	0x1d, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x5f, 0x65, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x6e, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x19, 0x6b, 0x65, 0x65, 0x70, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x45, 0x76, 0x65, 0x72, 0x79, 0x4e, 0x48, 0x6f, 0x75, 0x72, 0x73, 0x12,
	0x46, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x2c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x53, 0x61,
	0x76, 0x65, 0x72, 0x44, 0x65, 0x66, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x35, 0x0a, 0x17, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x45, 0x47, 0x41, 0x43, 0x59, 0x10, 0x00, 0x12, 0x06,
	0x0a, 0x02, 0x56, 0x31, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x56, 0x32, 0x10, 0x02, 0x42, 0x7e,
	0x0a, 0x13, 0x6f, 0x72, 0x67, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x75, 0x74, 0x69, 0x6c, 0x42, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_core_protobuf_saver_proto_rawDescOnce sync.Once
	file_tensorflow_core_protobuf_saver_proto_rawDescData = file_tensorflow_core_protobuf_saver_proto_rawDesc
)

func file_tensorflow_core_protobuf_saver_proto_rawDescGZIP() []byte {
	file_tensorflow_core_protobuf_saver_proto_rawDescOnce.Do(func() {
		file_tensorflow_core_protobuf_saver_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_core_protobuf_saver_proto_rawDescData)
	})
	return file_tensorflow_core_protobuf_saver_proto_rawDescData
}

var file_tensorflow_core_protobuf_saver_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tensorflow_core_protobuf_saver_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_core_protobuf_saver_proto_goTypes = []interface{}{
	(SaverDef_CheckpointFormatVersion)(0), // 0: tensorflow.SaverDef.CheckpointFormatVersion
	(*SaverDef)(nil),                      // 1: tensorflow.SaverDef
}
var file_tensorflow_core_protobuf_saver_proto_depIdxs = []int32{
	0, // 0: tensorflow.SaverDef.version:type_name -> tensorflow.SaverDef.CheckpointFormatVersion
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tensorflow_core_protobuf_saver_proto_init() }
func file_tensorflow_core_protobuf_saver_proto_init() {
	if File_tensorflow_core_protobuf_saver_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_core_protobuf_saver_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaverDef); i {
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
			RawDescriptor: file_tensorflow_core_protobuf_saver_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_core_protobuf_saver_proto_goTypes,
		DependencyIndexes: file_tensorflow_core_protobuf_saver_proto_depIdxs,
		EnumInfos:         file_tensorflow_core_protobuf_saver_proto_enumTypes,
		MessageInfos:      file_tensorflow_core_protobuf_saver_proto_msgTypes,
	}.Build()
	File_tensorflow_core_protobuf_saver_proto = out.File
	file_tensorflow_core_protobuf_saver_proto_rawDesc = nil
	file_tensorflow_core_protobuf_saver_proto_goTypes = nil
	file_tensorflow_core_protobuf_saver_proto_depIdxs = nil
}
