// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: video.proto

package video

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
	VideoService_GetVideoFeed_FullMethodName = "/video.VideoService/GetVideoFeed"
)

// VideoServiceClient is the client API for VideoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoServiceClient interface {
	GetVideoFeed(ctx context.Context, in *GetVideoFeedReq, opts ...grpc.CallOption) (*GetVideoFeedResp, error)
}

type videoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoServiceClient(cc grpc.ClientConnInterface) VideoServiceClient {
	return &videoServiceClient{cc}
}

func (c *videoServiceClient) GetVideoFeed(ctx context.Context, in *GetVideoFeedReq, opts ...grpc.CallOption) (*GetVideoFeedResp, error) {
	out := new(GetVideoFeedResp)
	err := c.cc.Invoke(ctx, VideoService_GetVideoFeed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServiceServer is the server API for VideoService service.
// All implementations must embed UnimplementedVideoServiceServer
// for forward compatibility
type VideoServiceServer interface {
	GetVideoFeed(context.Context, *GetVideoFeedReq) (*GetVideoFeedResp, error)
	mustEmbedUnimplementedVideoServiceServer()
}

// UnimplementedVideoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVideoServiceServer struct {
}

func (UnimplementedVideoServiceServer) GetVideoFeed(context.Context, *GetVideoFeedReq) (*GetVideoFeedResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoFeed not implemented")
}
func (UnimplementedVideoServiceServer) mustEmbedUnimplementedVideoServiceServer() {}

// UnsafeVideoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoServiceServer will
// result in compilation errors.
type UnsafeVideoServiceServer interface {
	mustEmbedUnimplementedVideoServiceServer()
}

func RegisterVideoServiceServer(s grpc.ServiceRegistrar, srv VideoServiceServer) {
	s.RegisterService(&VideoService_ServiceDesc, srv)
}

func _VideoService_GetVideoFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoFeedReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).GetVideoFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoService_GetVideoFeed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).GetVideoFeed(ctx, req.(*GetVideoFeedReq))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoService_ServiceDesc is the grpc.ServiceDesc for VideoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "video.VideoService",
	HandlerType: (*VideoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVideoFeed",
			Handler:    _VideoService_GetVideoFeed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "video.proto",
}
