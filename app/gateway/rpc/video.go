package rpc

import (
	"context"
	"dzug/discovery"
	pb "dzug/protos/publish"
)

func PublishVideo(ctx context.Context, req *pb.PublishVideoReq) (resp *pb.PublishVideoResp, err error) {
	discovery.LoadClient("publish", &discovery.PublishClient)
	r, err := discovery.PublishClient.PublishVideo(ctx, req)
	if err != nil {
		return
	}
	return r, nil
}
