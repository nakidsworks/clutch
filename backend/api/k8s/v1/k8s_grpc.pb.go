// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package k8sv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// K8SAPIClient is the client API for K8SAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type K8SAPIClient interface {
	DescribePod(ctx context.Context, in *DescribePodRequest, opts ...grpc.CallOption) (*DescribePodResponse, error)
	ListPods(ctx context.Context, in *ListPodsRequest, opts ...grpc.CallOption) (*ListPodsResponse, error)
	DeletePod(ctx context.Context, in *DeletePodRequest, opts ...grpc.CallOption) (*DeletePodResponse, error)
	UpdatePod(ctx context.Context, in *UpdatePodRequest, opts ...grpc.CallOption) (*UpdatePodResponse, error)
	ResizeHPA(ctx context.Context, in *ResizeHPARequest, opts ...grpc.CallOption) (*ResizeHPAResponse, error)
	UpdateDeployment(ctx context.Context, in *UpdateDeploymentRequest, opts ...grpc.CallOption) (*UpdateDeploymentResponse, error)
	DeleteDeployment(ctx context.Context, in *DeleteDeploymentRequest, opts ...grpc.CallOption) (*DeleteDeploymentResponse, error)
}

type k8SAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewK8SAPIClient(cc grpc.ClientConnInterface) K8SAPIClient {
	return &k8SAPIClient{cc}
}

func (c *k8SAPIClient) DescribePod(ctx context.Context, in *DescribePodRequest, opts ...grpc.CallOption) (*DescribePodResponse, error) {
	out := new(DescribePodResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/DescribePod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) ListPods(ctx context.Context, in *ListPodsRequest, opts ...grpc.CallOption) (*ListPodsResponse, error) {
	out := new(ListPodsResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/ListPods", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) DeletePod(ctx context.Context, in *DeletePodRequest, opts ...grpc.CallOption) (*DeletePodResponse, error) {
	out := new(DeletePodResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/DeletePod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) UpdatePod(ctx context.Context, in *UpdatePodRequest, opts ...grpc.CallOption) (*UpdatePodResponse, error) {
	out := new(UpdatePodResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/UpdatePod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) ResizeHPA(ctx context.Context, in *ResizeHPARequest, opts ...grpc.CallOption) (*ResizeHPAResponse, error) {
	out := new(ResizeHPAResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/ResizeHPA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) UpdateDeployment(ctx context.Context, in *UpdateDeploymentRequest, opts ...grpc.CallOption) (*UpdateDeploymentResponse, error) {
	out := new(UpdateDeploymentResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/UpdateDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k8SAPIClient) DeleteDeployment(ctx context.Context, in *DeleteDeploymentRequest, opts ...grpc.CallOption) (*DeleteDeploymentResponse, error) {
	out := new(DeleteDeploymentResponse)
	err := c.cc.Invoke(ctx, "/clutch.k8s.v1.K8sAPI/DeleteDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// K8SAPIServer is the server API for K8SAPI service.
// All implementations should embed UnimplementedK8SAPIServer
// for forward compatibility
type K8SAPIServer interface {
	DescribePod(context.Context, *DescribePodRequest) (*DescribePodResponse, error)
	ListPods(context.Context, *ListPodsRequest) (*ListPodsResponse, error)
	DeletePod(context.Context, *DeletePodRequest) (*DeletePodResponse, error)
	UpdatePod(context.Context, *UpdatePodRequest) (*UpdatePodResponse, error)
	ResizeHPA(context.Context, *ResizeHPARequest) (*ResizeHPAResponse, error)
	UpdateDeployment(context.Context, *UpdateDeploymentRequest) (*UpdateDeploymentResponse, error)
	DeleteDeployment(context.Context, *DeleteDeploymentRequest) (*DeleteDeploymentResponse, error)
}

// UnimplementedK8SAPIServer should be embedded to have forward compatible implementations.
type UnimplementedK8SAPIServer struct {
}

func (UnimplementedK8SAPIServer) DescribePod(context.Context, *DescribePodRequest) (*DescribePodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribePod not implemented")
}
func (UnimplementedK8SAPIServer) ListPods(context.Context, *ListPodsRequest) (*ListPodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPods not implemented")
}
func (UnimplementedK8SAPIServer) DeletePod(context.Context, *DeletePodRequest) (*DeletePodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePod not implemented")
}
func (UnimplementedK8SAPIServer) UpdatePod(context.Context, *UpdatePodRequest) (*UpdatePodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePod not implemented")
}
func (UnimplementedK8SAPIServer) ResizeHPA(context.Context, *ResizeHPARequest) (*ResizeHPAResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResizeHPA not implemented")
}
func (UnimplementedK8SAPIServer) UpdateDeployment(context.Context, *UpdateDeploymentRequest) (*UpdateDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDeployment not implemented")
}
func (UnimplementedK8SAPIServer) DeleteDeployment(context.Context, *DeleteDeploymentRequest) (*DeleteDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDeployment not implemented")
}

// UnsafeK8SAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to K8SAPIServer will
// result in compilation errors.
type UnsafeK8SAPIServer interface {
	mustEmbedUnimplementedK8SAPIServer()
}

func RegisterK8SAPIServer(s *grpc.Server, srv K8SAPIServer) {
	s.RegisterService(&_K8SAPI_serviceDesc, srv)
}

func _K8SAPI_DescribePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribePodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).DescribePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/DescribePod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).DescribePod(ctx, req.(*DescribePodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_ListPods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).ListPods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/ListPods",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).ListPods(ctx, req.(*ListPodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_DeletePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).DeletePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/DeletePod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).DeletePod(ctx, req.(*DeletePodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_UpdatePod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).UpdatePod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/UpdatePod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).UpdatePod(ctx, req.(*UpdatePodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_ResizeHPA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResizeHPARequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).ResizeHPA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/ResizeHPA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).ResizeHPA(ctx, req.(*ResizeHPARequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_UpdateDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).UpdateDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/UpdateDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).UpdateDeployment(ctx, req.(*UpdateDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _K8SAPI_DeleteDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K8SAPIServer).DeleteDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.k8s.v1.K8sAPI/DeleteDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K8SAPIServer).DeleteDeployment(ctx, req.(*DeleteDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _K8SAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.k8s.v1.K8sAPI",
	HandlerType: (*K8SAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescribePod",
			Handler:    _K8SAPI_DescribePod_Handler,
		},
		{
			MethodName: "ListPods",
			Handler:    _K8SAPI_ListPods_Handler,
		},
		{
			MethodName: "DeletePod",
			Handler:    _K8SAPI_DeletePod_Handler,
		},
		{
			MethodName: "UpdatePod",
			Handler:    _K8SAPI_UpdatePod_Handler,
		},
		{
			MethodName: "ResizeHPA",
			Handler:    _K8SAPI_ResizeHPA_Handler,
		},
		{
			MethodName: "UpdateDeployment",
			Handler:    _K8SAPI_UpdateDeployment_Handler,
		},
		{
			MethodName: "DeleteDeployment",
			Handler:    _K8SAPI_DeleteDeployment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "k8s/v1/k8s.proto",
}