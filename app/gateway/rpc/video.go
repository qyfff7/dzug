package rpc

import (
	"context"
	"dzug/discovery"
	"dzug/protos/video"
)

func Feed(ctx context.Context, req *video.GetVideoListByTimeReq) (*video.GetVideoListByTimeResp, error) {

	discovery.LoadClient("video", &discovery.VideoClient)
	r, err := discovery.VideoClient.GetVideoListByTime(ctx, req) // 调用注册方法
	if err != nil {
		return nil, err
	}
	return r, nil
}
