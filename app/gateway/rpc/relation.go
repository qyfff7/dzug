package rpc

import (
	"context"
	"dzug/idl/relation"
)

func RelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	loadClient("relation", &RelationClient)

	r, err := RelationClient.DouyinRelationAction(ctx, req)
	if err != nil {
		return
	}
	return r, nil
}
