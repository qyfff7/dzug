package service

import (
	"context"
	pb "dzug/protos/publish"
)

type VideoServer struct {
	pb.UnimplementedPublishServiceServer
}

func (p *VideoServer) PublishVideo(ctx context.Context, req *pb.PublishVideoReq) (*pb.PublishVideoResp, error) {
	resp := pb.PublishVideoResp{}
	return &resp, nil
}
func (p *VideoServer) GetVideoListByUserId(ctx context.Context, req *pb.GetVideoListByUserIdReq) (*pb.GetVideoListByUserIdResp, error) {
	resp := pb.GetVideoListByUserIdResp{}
	return &resp, nil
}
