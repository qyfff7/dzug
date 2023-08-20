package rpc

import (
	"context"
	"dzug/discovery"
	"dzug/protos/relation"
)

// RelationAction 关注、取关操作
func RelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (*relation.DouyinRelationActionResponse, error) {
	discovery.LoadClient("relation", &discovery.RelationClient)

	resp := &relation.DouyinRelationActionResponse{}
	resp, err := discovery.RelationClient.DouyinRelationAction(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RelationFollowList 获取关注列表操作
func RelationFollowList(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (*relation.DouyinRelationFollowListResponse, error) {
	resp := &relation.DouyinRelationFollowListResponse{}
	discovery.LoadClient("relation", &discovery.RelationClient)

	resp, err := discovery.RelationClient.DouyinRelationFollowList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RelationFollowerList 获取粉丝列表操作
func RelationFollowerList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (*relation.DouyinRelationFollowerListResponse, error) {
	resp := &relation.DouyinRelationFollowerListResponse{}
	discovery.LoadClient("relation", &discovery.RelationClient)

	resp, err := discovery.RelationClient.DouyinRelationFollowerList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RelationFriendList 获取好友列表操作
func RelationFriendList(ctx context.Context, req *relation.DouyinRelationFriendListRequest) (*relation.DouyinRelationFriendListResponse, error) {
	resp := &relation.DouyinRelationFriendListResponse{}
	discovery.LoadClient("relation", &discovery.RelationClient)

	resp, err := discovery.RelationClient.DouyinRelationFriendList(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
