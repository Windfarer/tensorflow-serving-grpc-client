// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: tensorflow_serving/config/platform_config.proto

package config

import (
	any "github.com/golang/protobuf/ptypes/any"
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

// Configuration for a servable platform e.g. tensorflow or other ML systems.
type PlatformConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The config proto for a SourceAdapter in the StoragePathSourceAdapter
	// registry.
	SourceAdapterConfig *any.Any `protobuf:"bytes,1,opt,name=source_adapter_config,json=sourceAdapterConfig,proto3" json:"source_adapter_config,omitempty"`
}

func (x *PlatformConfig) Reset() {
	*x = PlatformConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_config_platform_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformConfig) ProtoMessage() {}

func (x *PlatformConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_config_platform_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformConfig.ProtoReflect.Descriptor instead.
func (*PlatformConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_config_platform_config_proto_rawDescGZIP(), []int{0}
}

func (x *PlatformConfig) GetSourceAdapterConfig() *any.Any {
	if x != nil {
		return x.SourceAdapterConfig
	}
	return nil
}

type PlatformConfigMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A map from a platform name to a platform config. The platform name is used
	// in ModelConfig.model_platform.
	PlatformConfigs map[string]*PlatformConfig `protobuf:"bytes,1,rep,name=platform_configs,json=platformConfigs,proto3" json:"platform_configs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PlatformConfigMap) Reset() {
	*x = PlatformConfigMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_config_platform_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatformConfigMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatformConfigMap) ProtoMessage() {}

func (x *PlatformConfigMap) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_config_platform_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatformConfigMap.ProtoReflect.Descriptor instead.
func (*PlatformConfigMap) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_config_platform_config_proto_rawDescGZIP(), []int{1}
}

func (x *PlatformConfigMap) GetPlatformConfigs() map[string]*PlatformConfig {
	if x != nil {
		return x.PlatformConfigs
	}
	return nil
}

var File_tensorflow_serving_config_platform_config_proto protoreflect.FileDescriptor

var file_tensorflow_serving_config_platform_config_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5a, 0x0a, 0x0e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x48, 0x0a, 0x15, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x13, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41,
	0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0xe2, 0x01, 0x0a,
	0x11, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d,
	0x61, 0x70, 0x12, 0x65, 0x0a, 0x10, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x74,
	0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x4d, 0x61, 0x70, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x1a, 0x66, 0x0a, 0x14, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x38, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x1e, 0x5a, 0x19, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0xf8, 0x01,
	0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_serving_config_platform_config_proto_rawDescOnce sync.Once
	file_tensorflow_serving_config_platform_config_proto_rawDescData = file_tensorflow_serving_config_platform_config_proto_rawDesc
)

func file_tensorflow_serving_config_platform_config_proto_rawDescGZIP() []byte {
	file_tensorflow_serving_config_platform_config_proto_rawDescOnce.Do(func() {
		file_tensorflow_serving_config_platform_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_serving_config_platform_config_proto_rawDescData)
	})
	return file_tensorflow_serving_config_platform_config_proto_rawDescData
}

var file_tensorflow_serving_config_platform_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tensorflow_serving_config_platform_config_proto_goTypes = []interface{}{
	(*PlatformConfig)(nil),    // 0: tensorflow.serving.PlatformConfig
	(*PlatformConfigMap)(nil), // 1: tensorflow.serving.PlatformConfigMap
	nil,                       // 2: tensorflow.serving.PlatformConfigMap.PlatformConfigsEntry
	(*any.Any)(nil),           // 3: google.protobuf.Any
}
var file_tensorflow_serving_config_platform_config_proto_depIdxs = []int32{
	3, // 0: tensorflow.serving.PlatformConfig.source_adapter_config:type_name -> google.protobuf.Any
	2, // 1: tensorflow.serving.PlatformConfigMap.platform_configs:type_name -> tensorflow.serving.PlatformConfigMap.PlatformConfigsEntry
	0, // 2: tensorflow.serving.PlatformConfigMap.PlatformConfigsEntry.value:type_name -> tensorflow.serving.PlatformConfig
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_config_platform_config_proto_init() }
func file_tensorflow_serving_config_platform_config_proto_init() {
	if File_tensorflow_serving_config_platform_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_serving_config_platform_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformConfig); i {
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
		file_tensorflow_serving_config_platform_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatformConfigMap); i {
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
			RawDescriptor: file_tensorflow_serving_config_platform_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_serving_config_platform_config_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_config_platform_config_proto_depIdxs,
		MessageInfos:      file_tensorflow_serving_config_platform_config_proto_msgTypes,
	}.Build()
	File_tensorflow_serving_config_platform_config_proto = out.File
	file_tensorflow_serving_config_platform_config_proto_rawDesc = nil
	file_tensorflow_serving_config_platform_config_proto_goTypes = nil
	file_tensorflow_serving_config_platform_config_proto_depIdxs = nil
}
