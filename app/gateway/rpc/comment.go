package rpc

import (
	"context"
	"dzug/discovery"
	"dzug/protos/comment"
)

// 评论操作
func CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	discovery.LoadClient("comment", &discovery.CommentClient)
	r, err := discovery.CommentClient.Action(ctx, req)
	if err != nil {
		return
	}
	return r, nil
}

func CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	discovery.LoadClient("comment", &discovery.CommentClient)
	r, err := discovery.CommentClient.List(ctx, req)
	if err != nil {
		return
	}
	return r, nil
}
