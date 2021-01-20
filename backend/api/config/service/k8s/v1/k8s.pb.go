// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: config/service/k8s/v1/k8s.proto

package k8sv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of paths for `kubeconfig` files. Each config will be read and each context will become a clientset
	// in the clientset manager. If no files are provided, the service will use the K8s SDK default loading rules
	// which are detailed here: https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/,
	// i.e. look for configs:
	// - in a comma-separated `KUBECONFIG` environment variable
	// - else `$HOME/.kube/config`
	// - else in-cluster config
	// https://github.com/kubernetes/client-go/tree/master/examples/in-cluster-client-configuration
	Kubeconfigs      []string          `protobuf:"bytes,1,rep,name=kubeconfigs,proto3" json:"kubeconfigs,omitempty"`
	RestClientConfig *RestClientConfig `protobuf:"bytes,2,opt,name=rest_client_config,json=restClientConfig,proto3" json:"rest_client_config,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_k8s_v1_k8s_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_k8s_v1_k8s_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_config_service_k8s_v1_k8s_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetKubeconfigs() []string {
	if x != nil {
		return x.Kubeconfigs
	}
	return nil
}

func (x *Config) GetRestClientConfig() *RestClientConfig {
	if x != nil {
		return x.RestClientConfig
	}
	return nil
}

// These configuration values are passed directly through to the rest config object.
type RestClientConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum length of time to wait before giving up on a server request.
	// By default a value of zero means no timeout.
	// https://github.com/kubernetes/client-go/blob/00dbcca6ee44c678754d3f5fda1bd0e704b26fe2/rest/config.go#L130-L131
	Timeout *duration.Duration `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// QPS indicates the maximum QPS to the master from this client.
	// If it's zero, the created RESTClient will use DefaultQPS: 5
	// https://github.com/kubernetes/client-go/blob/00dbcca6ee44c678754d3f5fda1bd0e704b26fe2/rest/config.go#L115-L117
	Qps float32 `protobuf:"fixed32,2,opt,name=qps,proto3" json:"qps,omitempty"`
	// Maximum burst for throttle.
	// If it's zero, the created RESTClient will use DefaultBurst: 10.
	// https://github.com/kubernetes/client-go/blob/00dbcca6ee44c678754d3f5fda1bd0e704b26fe2/rest/config.go#L119-L121
	Burst uint32 `protobuf:"varint,3,opt,name=burst,proto3" json:"burst,omitempty"`
}

func (x *RestClientConfig) Reset() {
	*x = RestClientConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_service_k8s_v1_k8s_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestClientConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestClientConfig) ProtoMessage() {}

func (x *RestClientConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_service_k8s_v1_k8s_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestClientConfig.ProtoReflect.Descriptor instead.
func (*RestClientConfig) Descriptor() ([]byte, []int) {
	return file_config_service_k8s_v1_k8s_proto_rawDescGZIP(), []int{1}
}

func (x *RestClientConfig) GetTimeout() *duration.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *RestClientConfig) GetQps() float32 {
	if x != nil {
		return x.Qps
	}
	return 0
}

func (x *RestClientConfig) GetBurst() uint32 {
	if x != nil {
		return x.Burst
	}
	return 0
}

var File_config_service_k8s_v1_k8s_proto protoreflect.FileDescriptor

var file_config_service_k8s_v1_k8s_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x6b, 0x38, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x38, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1c, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x2a, 0x0a, 0x0b, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02,
	0x18, 0x01, 0x52, 0x0b, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12,
	0x5c, 0x0a, 0x12, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6c,
	0x75, 0x74, 0x63, 0x68, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x10, 0x72, 0x65, 0x73,
	0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x90, 0x01,
	0x0a, 0x10, 0x52, 0x65, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x3f, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a,
	0xfa, 0x42, 0x07, 0xaa, 0x01, 0x04, 0x32, 0x02, 0x08, 0x00, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x12, 0x1c, 0x0a, 0x03, 0x71, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x42, 0x0a, 0xfa, 0x42, 0x07, 0x0a, 0x05, 0x2d, 0x00, 0x00, 0x00, 0x00, 0x52, 0x03, 0x71, 0x70,
	0x73, 0x12, 0x1d, 0x0a, 0x05, 0x62, 0x75, 0x72, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x28, 0x00, 0x52, 0x05, 0x62, 0x75, 0x72, 0x73, 0x74,
	0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c,
	0x79, 0x66, 0x74, 0x2f, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x6b, 0x38, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6b, 0x38, 0x73,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_service_k8s_v1_k8s_proto_rawDescOnce sync.Once
	file_config_service_k8s_v1_k8s_proto_rawDescData = file_config_service_k8s_v1_k8s_proto_rawDesc
)

func file_config_service_k8s_v1_k8s_proto_rawDescGZIP() []byte {
	file_config_service_k8s_v1_k8s_proto_rawDescOnce.Do(func() {
		file_config_service_k8s_v1_k8s_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_service_k8s_v1_k8s_proto_rawDescData)
	})
	return file_config_service_k8s_v1_k8s_proto_rawDescData
}

var file_config_service_k8s_v1_k8s_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_config_service_k8s_v1_k8s_proto_goTypes = []interface{}{
	(*Config)(nil),            // 0: clutch.config.service.k8s.v1.Config
	(*RestClientConfig)(nil),  // 1: clutch.config.service.k8s.v1.RestClientConfig
	(*duration.Duration)(nil), // 2: google.protobuf.Duration
}
var file_config_service_k8s_v1_k8s_proto_depIdxs = []int32{
	1, // 0: clutch.config.service.k8s.v1.Config.rest_client_config:type_name -> clutch.config.service.k8s.v1.RestClientConfig
	2, // 1: clutch.config.service.k8s.v1.RestClientConfig.timeout:type_name -> google.protobuf.Duration
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_config_service_k8s_v1_k8s_proto_init() }
func file_config_service_k8s_v1_k8s_proto_init() {
	if File_config_service_k8s_v1_k8s_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_service_k8s_v1_k8s_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_config_service_k8s_v1_k8s_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestClientConfig); i {
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
			RawDescriptor: file_config_service_k8s_v1_k8s_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_service_k8s_v1_k8s_proto_goTypes,
		DependencyIndexes: file_config_service_k8s_v1_k8s_proto_depIdxs,
		MessageInfos:      file_config_service_k8s_v1_k8s_proto_msgTypes,
	}.Build()
	File_config_service_k8s_v1_k8s_proto = out.File
	file_config_service_k8s_v1_k8s_proto_rawDesc = nil
	file_config_service_k8s_v1_k8s_proto_goTypes = nil
	file_config_service_k8s_v1_k8s_proto_depIdxs = nil
}
