package service

import (
	"context"
	"dzug/app/relation/idl"
)

type RelationSrv struct {
	__.UnimplementedDouyinRelationActionServiceServer
}

func (r *RelationSrv) DouyinRelationAction(context.Context, *__.DouyinRelationActionRequest) (*__.DouyinRelationActionResponse, error) {
	return &__.DouyinRelationActionResponse{
		StatusCode: 200,
		StatusMsg:  "调用成功",
	}, nil
}
