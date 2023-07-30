package rpc

import (
	"context"
	"dzug/discovery"
	"dzug/idl/relation"
)

func RelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	discovery.LoadClient("relation", &discovery.RelationClient) // 加载etcd客户端

	r, err := discovery.RelationClient.DouyinRelationAction(ctx, req)
	if err != nil {
		return
	}
	return r, nil
}
