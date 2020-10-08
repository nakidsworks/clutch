// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: sourcecontrol/v1/sourcecontrol.proto

package sourcecontrolv1

import (
	context "context"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	_ "github.com/lyft/clutch/backend/api/api/v1"
	v1 "github.com/lyft/clutch/backend/api/sourcecontrol/github/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateRepositoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner       string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Types that are assignable to Options:
	//	*CreateRepositoryRequest_CustomOptions
	//	*CreateRepositoryRequest_GithubOptions
	Options isCreateRepositoryRequest_Options `protobuf_oneof:"options"`
}

func (x *CreateRepositoryRequest) Reset() {
	*x = CreateRepositoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sourcecontrol_v1_sourcecontrol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRepositoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRepositoryRequest) ProtoMessage() {}

func (x *CreateRepositoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sourcecontrol_v1_sourcecontrol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRepositoryRequest.ProtoReflect.Descriptor instead.
func (*CreateRepositoryRequest) Descriptor() ([]byte, []int) {
	return file_sourcecontrol_v1_sourcecontrol_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRepositoryRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *CreateRepositoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRepositoryRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (m *CreateRepositoryRequest) GetOptions() isCreateRepositoryRequest_Options {
	if m != nil {
		return m.Options
	}
	return nil
}

func (x *CreateRepositoryRequest) GetCustomOptions() *any.Any {
	if x, ok := x.GetOptions().(*CreateRepositoryRequest_CustomOptions); ok {
		return x.CustomOptions
	}
	return nil
}

func (x *CreateRepositoryRequest) GetGithubOptions() *v1.CreateRepositoryOptions {
	if x, ok := x.GetOptions().(*CreateRepositoryRequest_GithubOptions); ok {
		return x.GithubOptions
	}
	return nil
}

type isCreateRepositoryRequest_Options interface {
	isCreateRepositoryRequest_Options()
}

type CreateRepositoryRequest_CustomOptions struct {
	CustomOptions *any.Any `protobuf:"bytes,4,opt,name=custom_options,json=customOptions,proto3,oneof"`
}

type CreateRepositoryRequest_GithubOptions struct {
	GithubOptions *v1.CreateRepositoryOptions `protobuf:"bytes,5,opt,name=github_options,json=githubOptions,proto3,oneof"`
}

func (*CreateRepositoryRequest_CustomOptions) isCreateRepositoryRequest_Options() {}

func (*CreateRepositoryRequest_GithubOptions) isCreateRepositoryRequest_Options() {}

type CreateRepositoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *CreateRepositoryResponse) Reset() {
	*x = CreateRepositoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sourcecontrol_v1_sourcecontrol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRepositoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRepositoryResponse) ProtoMessage() {}

func (x *CreateRepositoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sourcecontrol_v1_sourcecontrol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRepositoryResponse.ProtoReflect.Descriptor instead.
func (*CreateRepositoryResponse) Descriptor() ([]byte, []int) {
	return file_sourcecontrol_v1_sourcecontrol_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRepositoryResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_sourcecontrol_v1_sourcecontrol_proto protoreflect.FileDescriptor

var file_sourcecontrol_v1_sourcecontrol_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa8, 0x02, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x20, 0x01, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x20, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x0e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x60, 0x0a, 0x0e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x37, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x48, 0x00, 0x52, 0x0d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x0e, 0x0a, 0x07,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x22, 0x2c, 0x0a, 0x18,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x32, 0xc1, 0x01, 0x0a, 0x10, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x41, 0x50, 0x49, 0x12,
	0xac, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x30, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x63, 0x6c, 0x75, 0x74, 0x63, 0x68, 0x2e,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x27, 0x22, 0x22, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x3a, 0x01, 0x2a, 0xaa, 0xe1, 0x1c, 0x02, 0x08, 0x01, 0x42, 0x11,
	0x5a, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sourcecontrol_v1_sourcecontrol_proto_rawDescOnce sync.Once
	file_sourcecontrol_v1_sourcecontrol_proto_rawDescData = file_sourcecontrol_v1_sourcecontrol_proto_rawDesc
)

func file_sourcecontrol_v1_sourcecontrol_proto_rawDescGZIP() []byte {
	file_sourcecontrol_v1_sourcecontrol_proto_rawDescOnce.Do(func() {
		file_sourcecontrol_v1_sourcecontrol_proto_rawDescData = protoimpl.X.CompressGZIP(file_sourcecontrol_v1_sourcecontrol_proto_rawDescData)
	})
	return file_sourcecontrol_v1_sourcecontrol_proto_rawDescData
}

var file_sourcecontrol_v1_sourcecontrol_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sourcecontrol_v1_sourcecontrol_proto_goTypes = []interface{}{
	(*CreateRepositoryRequest)(nil),    // 0: clutch.sourcecontrol.v1.CreateRepositoryRequest
	(*CreateRepositoryResponse)(nil),   // 1: clutch.sourcecontrol.v1.CreateRepositoryResponse
	(*any.Any)(nil),                    // 2: google.protobuf.Any
	(*v1.CreateRepositoryOptions)(nil), // 3: clutch.sourcecontrol.github.v1.CreateRepositoryOptions
}
var file_sourcecontrol_v1_sourcecontrol_proto_depIdxs = []int32{
	2, // 0: clutch.sourcecontrol.v1.CreateRepositoryRequest.custom_options:type_name -> google.protobuf.Any
	3, // 1: clutch.sourcecontrol.v1.CreateRepositoryRequest.github_options:type_name -> clutch.sourcecontrol.github.v1.CreateRepositoryOptions
	0, // 2: clutch.sourcecontrol.v1.SourceControlAPI.CreateRepository:input_type -> clutch.sourcecontrol.v1.CreateRepositoryRequest
	1, // 3: clutch.sourcecontrol.v1.SourceControlAPI.CreateRepository:output_type -> clutch.sourcecontrol.v1.CreateRepositoryResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sourcecontrol_v1_sourcecontrol_proto_init() }
func file_sourcecontrol_v1_sourcecontrol_proto_init() {
	if File_sourcecontrol_v1_sourcecontrol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sourcecontrol_v1_sourcecontrol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRepositoryRequest); i {
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
		file_sourcecontrol_v1_sourcecontrol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRepositoryResponse); i {
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
	file_sourcecontrol_v1_sourcecontrol_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CreateRepositoryRequest_CustomOptions)(nil),
		(*CreateRepositoryRequest_GithubOptions)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sourcecontrol_v1_sourcecontrol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sourcecontrol_v1_sourcecontrol_proto_goTypes,
		DependencyIndexes: file_sourcecontrol_v1_sourcecontrol_proto_depIdxs,
		MessageInfos:      file_sourcecontrol_v1_sourcecontrol_proto_msgTypes,
	}.Build()
	File_sourcecontrol_v1_sourcecontrol_proto = out.File
	file_sourcecontrol_v1_sourcecontrol_proto_rawDesc = nil
	file_sourcecontrol_v1_sourcecontrol_proto_goTypes = nil
	file_sourcecontrol_v1_sourcecontrol_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SourceControlAPIClient is the client API for SourceControlAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SourceControlAPIClient interface {
	CreateRepository(ctx context.Context, in *CreateRepositoryRequest, opts ...grpc.CallOption) (*CreateRepositoryResponse, error)
}

type sourceControlAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewSourceControlAPIClient(cc grpc.ClientConnInterface) SourceControlAPIClient {
	return &sourceControlAPIClient{cc}
}

func (c *sourceControlAPIClient) CreateRepository(ctx context.Context, in *CreateRepositoryRequest, opts ...grpc.CallOption) (*CreateRepositoryResponse, error) {
	out := new(CreateRepositoryResponse)
	err := c.cc.Invoke(ctx, "/clutch.sourcecontrol.v1.SourceControlAPI/CreateRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SourceControlAPIServer is the server API for SourceControlAPI service.
type SourceControlAPIServer interface {
	CreateRepository(context.Context, *CreateRepositoryRequest) (*CreateRepositoryResponse, error)
}

// UnimplementedSourceControlAPIServer can be embedded to have forward compatible implementations.
type UnimplementedSourceControlAPIServer struct {
}

func (*UnimplementedSourceControlAPIServer) CreateRepository(context.Context, *CreateRepositoryRequest) (*CreateRepositoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRepository not implemented")
}

func RegisterSourceControlAPIServer(s *grpc.Server, srv SourceControlAPIServer) {
	s.RegisterService(&_SourceControlAPI_serviceDesc, srv)
}

func _SourceControlAPI_CreateRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourceControlAPIServer).CreateRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.sourcecontrol.v1.SourceControlAPI/CreateRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourceControlAPIServer).CreateRepository(ctx, req.(*CreateRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SourceControlAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.sourcecontrol.v1.SourceControlAPI",
	HandlerType: (*SourceControlAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRepository",
			Handler:    _SourceControlAPI_CreateRepository_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sourcecontrol/v1/sourcecontrol.proto",
}
