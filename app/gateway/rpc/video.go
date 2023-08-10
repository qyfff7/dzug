package rpc

import (
	"context"
	"dzug/discovery"
	pb "dzug/protos/video"
)

func Feed(ctx context.Context, req *pb.GetVideoFeedReq) (*pb.GetVideoFeedResp, error) {

	discovery.LoadClient("video", &discovery.UserClient)
	r, err := discovery.VideoClient.GetVideoFeed(ctx, req) // 调用注册方法
	if err != nil {
		return nil, err
	}
	return r, nil
}
