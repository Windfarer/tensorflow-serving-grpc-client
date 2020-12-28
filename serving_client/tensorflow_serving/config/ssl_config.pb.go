// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: tensorflow_serving/config/ssl_config.proto

package tensorflow_serving

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

// Configuration for a secure gRPC channel
type SSLConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// private server key for SSL
	ServerKey string `protobuf:"bytes,1,opt,name=server_key,json=serverKey,proto3" json:"server_key,omitempty"`
	// public server certificate
	ServerCert string `protobuf:"bytes,2,opt,name=server_cert,json=serverCert,proto3" json:"server_cert,omitempty"`
	//  custom certificate authority
	CustomCa string `protobuf:"bytes,3,opt,name=custom_ca,json=customCa,proto3" json:"custom_ca,omitempty"`
	// valid client certificate required ?
	ClientVerify bool `protobuf:"varint,4,opt,name=client_verify,json=clientVerify,proto3" json:"client_verify,omitempty"`
}

func (x *SSLConfig) Reset() {
	*x = SSLConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_config_ssl_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSLConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSLConfig) ProtoMessage() {}

func (x *SSLConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_config_ssl_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSLConfig.ProtoReflect.Descriptor instead.
func (*SSLConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_config_ssl_config_proto_rawDescGZIP(), []int{0}
}

func (x *SSLConfig) GetServerKey() string {
	if x != nil {
		return x.ServerKey
	}
	return ""
}

func (x *SSLConfig) GetServerCert() string {
	if x != nil {
		return x.ServerCert
	}
	return ""
}

func (x *SSLConfig) GetCustomCa() string {
	if x != nil {
		return x.CustomCa
	}
	return ""
}

func (x *SSLConfig) GetClientVerify() bool {
	if x != nil {
		return x.ClientVerify
	}
	return false
}

var File_tensorflow_serving_config_ssl_config_proto protoreflect.FileDescriptor

var file_tensorflow_serving_config_ssl_config_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x73, 0x6c, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x22, 0x8d, 0x01, 0x0a, 0x09, 0x53, 0x53, 0x4c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x65, 0x72, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x63, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x61, 0x12, 0x23, 0x0a, 0x0d, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x42, 0x03, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_serving_config_ssl_config_proto_rawDescOnce sync.Once
	file_tensorflow_serving_config_ssl_config_proto_rawDescData = file_tensorflow_serving_config_ssl_config_proto_rawDesc
)

func file_tensorflow_serving_config_ssl_config_proto_rawDescGZIP() []byte {
	file_tensorflow_serving_config_ssl_config_proto_rawDescOnce.Do(func() {
		file_tensorflow_serving_config_ssl_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_serving_config_ssl_config_proto_rawDescData)
	})
	return file_tensorflow_serving_config_ssl_config_proto_rawDescData
}

var file_tensorflow_serving_config_ssl_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_serving_config_ssl_config_proto_goTypes = []interface{}{
	(*SSLConfig)(nil), // 0: tensorflow.serving.SSLConfig
}
var file_tensorflow_serving_config_ssl_config_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_config_ssl_config_proto_init() }
func file_tensorflow_serving_config_ssl_config_proto_init() {
	if File_tensorflow_serving_config_ssl_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_serving_config_ssl_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSLConfig); i {
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
			RawDescriptor: file_tensorflow_serving_config_ssl_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_serving_config_ssl_config_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_config_ssl_config_proto_depIdxs,
		MessageInfos:      file_tensorflow_serving_config_ssl_config_proto_msgTypes,
	}.Build()
	File_tensorflow_serving_config_ssl_config_proto = out.File
	file_tensorflow_serving_config_ssl_config_proto_rawDesc = nil
	file_tensorflow_serving_config_ssl_config_proto_goTypes = nil
	file_tensorflow_serving_config_ssl_config_proto_depIdxs = nil
}
