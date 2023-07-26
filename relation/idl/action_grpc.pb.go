// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.5
// source: action.proto

package __

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
	DouyinRelationActionService_DouyinRelationAction_FullMethodName = "/idl.DouyinRelationActionService/DouyinRelationAction"
)

// DouyinRelationActionServiceClient is the client API for DouyinRelationActionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DouyinRelationActionServiceClient interface {
	DouyinRelationAction(ctx context.Context, in *DouyinRelationActionRequest, opts ...grpc.CallOption) (*DouyinRelationActionResponse, error)
}

type douyinRelationActionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDouyinRelationActionServiceClient(cc grpc.ClientConnInterface) DouyinRelationActionServiceClient {
	return &douyinRelationActionServiceClient{cc}
}

func (c *douyinRelationActionServiceClient) DouyinRelationAction(ctx context.Context, in *DouyinRelationActionRequest, opts ...grpc.CallOption) (*DouyinRelationActionResponse, error) {
	out := new(DouyinRelationActionResponse)
	err := c.cc.Invoke(ctx, DouyinRelationActionService_DouyinRelationAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DouyinRelationActionServiceServer is the server API for DouyinRelationActionService service.
// All implementations must embed UnimplementedDouyinRelationActionServiceServer
// for forward compatibility
type DouyinRelationActionServiceServer interface {
	DouyinRelationAction(context.Context, *DouyinRelationActionRequest) (*DouyinRelationActionResponse, error)
	mustEmbedUnimplementedDouyinRelationActionServiceServer()
}

// UnimplementedDouyinRelationActionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDouyinRelationActionServiceServer struct {
}

func (UnimplementedDouyinRelationActionServiceServer) DouyinRelationAction(context.Context, *DouyinRelationActionRequest) (*DouyinRelationActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DouyinRelationAction not implemented")
}
func (UnimplementedDouyinRelationActionServiceServer) mustEmbedUnimplementedDouyinRelationActionServiceServer() {
}

// UnsafeDouyinRelationActionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DouyinRelationActionServiceServer will
// result in compilation errors.
type UnsafeDouyinRelationActionServiceServer interface {
	mustEmbedUnimplementedDouyinRelationActionServiceServer()
}

func RegisterDouyinRelationActionServiceServer(s grpc.ServiceRegistrar, srv DouyinRelationActionServiceServer) {
	s.RegisterService(&DouyinRelationActionService_ServiceDesc, srv)
}

func _DouyinRelationActionService_DouyinRelationAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinRelationActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DouyinRelationActionServiceServer).DouyinRelationAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DouyinRelationActionService_DouyinRelationAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DouyinRelationActionServiceServer).DouyinRelationAction(ctx, req.(*DouyinRelationActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DouyinRelationActionService_ServiceDesc is the grpc.ServiceDesc for DouyinRelationActionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DouyinRelationActionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "idl.DouyinRelationActionService",
	HandlerType: (*DouyinRelationActionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DouyinRelationAction",
			Handler:    _DouyinRelationActionService_DouyinRelationAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "action.proto",
}
