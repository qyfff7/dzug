package service

import (
	"context"
	pb "dzug/relation/idl"
)

type RelationSrv struct {
	pb.UnimplementedDouyinRelationActionServiceServer
}

func (r *RelationSrv) DouyinRelationAction(context.Context, *pb.DouyinRelationActionRequest) (*pb.DouyinRelationActionResponse, error) {
	return &pb.DouyinRelationActionResponse{
		StatusCode: 200,
		StatusMsg:  "调用成功",
	}, nil
}
