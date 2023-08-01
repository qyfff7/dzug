package service

import (
	"context"
	"dzug/protos/relation"
)

type RelationSrv struct {
	relation.UnimplementedDouyinRelationActionServiceServer
}

func (r *RelationSrv) DouyinRelationAction(context.Context, *relation.DouyinRelationActionRequest) (*relation.DouyinRelationActionResponse, error) {
	return &relation.DouyinRelationActionResponse{
		StatusCode: 200,
		StatusMsg:  "调用成功，你成功进行了一次关系操作",
	}, nil
}
