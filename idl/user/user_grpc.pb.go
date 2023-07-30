// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.5
// source: user.proto

package user

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DouyinUserService_Register_FullMethodName = "/idl.DouyinUserService/Register"
	DouyinUserService_Login_FullMethodName    = "/idl.DouyinUserService/Login"
)

// DouyinUserServiceClient is the client API for DouyinUserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DouyinUserServiceClient interface {
	Register(ctx context.Context, in *DouyinUserRegisterRequest, opts ...grpc.CallOption) (*DouyinUserRegisterResponse, error)
	Login(ctx context.Context, in *DouyinUserLoginRequest, opts ...grpc.CallOption) (*DouyinUserLoginResponse, error)
}

type douyinUserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDouyinUserServiceClient(cc grpc.ClientConnInterface) DouyinUserServiceClient {
	return &douyinUserServiceClient{cc}
}

func (c *douyinUserServiceClient) Register(ctx context.Context, in *DouyinUserRegisterRequest, opts ...grpc.CallOption) (*DouyinUserRegisterResponse, error) {
	out := new(DouyinUserRegisterResponse)
	err := c.cc.Invoke(ctx, DouyinUserService_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *douyinUserServiceClient) Login(ctx context.Context, in *DouyinUserLoginRequest, opts ...grpc.CallOption) (*DouyinUserLoginResponse, error) {
	out := new(DouyinUserLoginResponse)
	err := c.cc.Invoke(ctx, DouyinUserService_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DouyinUserServiceServer is the server API for DouyinUserService service.
// All implementations must embed UnimplementedDouyinUserServiceServer
// for forward compatibility
type DouyinUserServiceServer interface {
	Register(context.Context, *DouyinUserRegisterRequest) (*DouyinUserRegisterResponse, error)
	Login(context.Context, *DouyinUserLoginRequest) (*DouyinUserLoginResponse, error)
	mustEmbedUnimplementedDouyinUserServiceServer()
}

// UnimplementedDouyinUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDouyinUserServiceServer struct {
}

func (UnimplementedDouyinUserServiceServer) Register(context.Context, *DouyinUserRegisterRequest) (*DouyinUserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedDouyinUserServiceServer) Login(context.Context, *DouyinUserLoginRequest) (*DouyinUserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedDouyinUserServiceServer) mustEmbedUnimplementedDouyinUserServiceServer() {}

// UnsafeDouyinUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DouyinUserServiceServer will
// result in compilation errors.
type UnsafeDouyinUserServiceServer interface {
	mustEmbedUnimplementedDouyinUserServiceServer()
}

func RegisterDouyinUserServiceServer(s grpc.ServiceRegistrar, srv DouyinUserServiceServer) {
	s.RegisterService(&DouyinUserService_ServiceDesc, srv)
}

func _DouyinUserService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinUserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DouyinUserServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DouyinUserService_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DouyinUserServiceServer).Register(ctx, req.(*DouyinUserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DouyinUserService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinUserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DouyinUserServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DouyinUserService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DouyinUserServiceServer).Login(ctx, req.(*DouyinUserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DouyinUserService_ServiceDesc is the grpc.ServiceDesc for DouyinUserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DouyinUserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "idl.DouyinUserService",
	HandlerType: (*DouyinUserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _DouyinUserService_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _DouyinUserService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
