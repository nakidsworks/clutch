// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package authnv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// AuthnAPIClient is the client API for AuthnAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthnAPIClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Callback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*CallbackResponse, error)
}

type authnAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthnAPIClient(cc grpc.ClientConnInterface) AuthnAPIClient {
	return &authnAPIClient{cc}
}

func (c *authnAPIClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/clutch.authn.v1.AuthnAPI/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authnAPIClient) Callback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*CallbackResponse, error) {
	out := new(CallbackResponse)
	err := c.cc.Invoke(ctx, "/clutch.authn.v1.AuthnAPI/Callback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthnAPIServer is the server API for AuthnAPI service.
// All implementations should embed UnimplementedAuthnAPIServer
// for forward compatibility
type AuthnAPIServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Callback(context.Context, *CallbackRequest) (*CallbackResponse, error)
}

// UnimplementedAuthnAPIServer should be embedded to have forward compatible implementations.
type UnimplementedAuthnAPIServer struct {
}

func (UnimplementedAuthnAPIServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthnAPIServer) Callback(context.Context, *CallbackRequest) (*CallbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Callback not implemented")
}

// UnsafeAuthnAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthnAPIServer will
// result in compilation errors.
type UnsafeAuthnAPIServer interface {
	mustEmbedUnimplementedAuthnAPIServer()
}

func RegisterAuthnAPIServer(s grpc.ServiceRegistrar, srv AuthnAPIServer) {
	s.RegisterService(&_AuthnAPI_serviceDesc, srv)
}

func _AuthnAPI_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthnAPIServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.authn.v1.AuthnAPI/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthnAPIServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthnAPI_Callback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthnAPIServer).Callback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/clutch.authn.v1.AuthnAPI/Callback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthnAPIServer).Callback(ctx, req.(*CallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthnAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "clutch.authn.v1.AuthnAPI",
	HandlerType: (*AuthnAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthnAPI_Login_Handler,
		},
		{
			MethodName: "Callback",
			Handler:    _AuthnAPI_Callback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authn/v1/authn.proto",
}
