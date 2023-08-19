package rpc

import (
	"context"
	"dzug/discovery"
	"dzug/protos/comment"

	"go.uber.org/zap"
)

// 评论操作
func CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (resp *comment.DouyinCommentActionResponse, err error) {
	resp = &comment.DouyinCommentActionResponse{}

	discovery.LoadClient("comment", &discovery.CommentClient)
	if err != nil {
		zap.L().Error("调用点赞服务失败")
		resp.StatusCode = 500
		resp.StatusMsg = "点赞失败"
		return
	}
	r, err := discovery.CommentClient.Action(ctx, req)
	return r, nil
}

func CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (resp *comment.DouyinCommentListResponse, err error) {
	resp = &comment.DouyinCommentListResponse{}
	discovery.LoadClient("comment", &discovery.CommentClient)
	r, err := discovery.CommentClient.List(ctx, req)
	if err != nil {
		return
	}
	return r, nil
}
