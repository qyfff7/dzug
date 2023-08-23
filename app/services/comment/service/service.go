package service

import (
	"context"
	"dzug/app/services/comment/dal/dao"
	"dzug/app/services/comment/dal/redis"
	"dzug/app/services/user/pkg/jwt"
	"dzug/protos/comment"
	"errors"

	"go.uber.org/zap"
)

type CommentSrv struct {
	comment.UnimplementedDouyinCommentServiceServer
}

func (c *CommentSrv) Action(ctx context.Context, in *comment.DouyinCommentActionRequest) (*comment.DouyinCommentActionResponse, error) {
	videoId := in.VideoId

	token := in.Token
	//2.从token中解析处userID
	u, err := jwt.ParseToken(token)
	if err != nil {
		//错误处理
		return &comment.DouyinCommentActionResponse{
			StatusCode: 500,
			StatusMsg:  "解析错误",
		}, errors.New("Token解析错误")
	}

	userId := u.UserID
	commentid := in.CommentId
	actiontype := in.ActionType
	context := in.CommentText
	if actiontype == 1 { //新增评论
		ans := redis.AddComm(ctx, videoId, commentid)
		if ans == 0 {
			return &comment.DouyinCommentActionResponse{
				StatusCode: 500,
				StatusMsg:  "服务器错误",
			}, errors.New("redis数据库错误")
		}
		ans = 0
		ans = dao.Comm(ctx, commentid, userId, videoId, context)
		if ans == 0 {
			return &comment.DouyinCommentActionResponse{
				StatusCode: 500,
				StatusMsg:  "服务器错误",
			}, errors.New("mysql数据库错误")
		}
		return &comment.DouyinCommentActionResponse{
			StatusCode: 200,
			StatusMsg:  "评论成功",
		}, nil
	}
	if actiontype == 2 {
		ans := redis.DelComm(ctx, videoId, commentid)
		if ans == 0 {
			return &comment.DouyinCommentActionResponse{
				StatusCode: 500,
				StatusMsg:  "服务器错误",
			}, errors.New("redis数据库错误")
		}
		ans = 0
		ans = dao.Incomm(ctx, videoId, commentid)
		if ans == 0 {
			return &comment.DouyinCommentActionResponse{
				StatusCode: 500,
				StatusMsg:  "服务器错误",
			}, errors.New("mysql数据库错误")
		}
		return &comment.DouyinCommentActionResponse{
			StatusCode: 200,
			StatusMsg:  "删除评论成功",
		}, nil
	}
	return &comment.DouyinCommentActionResponse{
		StatusCode: 200,
		StatusMsg:  "评论操作成功",
	}, nil
}

func (c *CommentSrv) List(ctx context.Context, in *comment.DouyinCommentListRequest) (*comment.DouyinCommentListResponse, error) {
	videoId := in.VideoId
	commentIds, err := redis.GetComm(ctx, videoId)
	if err != nil {
		zap.L().Error("redis获取视频评论列表失败")
		return nil, err
	}
	comments, err := dao.GetcommByCommentIDs(commentIds)
	if err != nil {
		zap.L().Error("获取评论列表失败")
		return &comment.DouyinCommentListResponse{
			StatusCode: 500,
			StatusMsg:  "查看评论列表失败",
		}, nil
	}
	return &comment.DouyinCommentListResponse{
		StatusCode:  200,
		StatusMsg:   "查看评论成功",
		CommentList: comments,
	}, nil
}
