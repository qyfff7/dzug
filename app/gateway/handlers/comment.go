package handlers

import (
	"dzug/app/gateway/rpc"
	"dzug/app/user/pkg/snowflake"

	pb "dzug/protos/comment"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CommentReq struct {
	VideoId    string `json:"video_id"`
	ActionType string `json:"action_type"`
}

// 评论相关操作
func CommentAction(ctx *gin.Context) {
	var CReq CommentReq

	//userId, _ := jwt.GetUserID(ctx)

	token := ctx.Query("token")

	err := ctx.ShouldBind(&CReq)
	if err != nil {
		zap.L().Fatal("绑定参数出错" + err.Error())
	}
	zap.L().Info(fmt.Sprintf("token:", token, " VideoId:", CReq.VideoId, " ActionType:", CReq.ActionType))
	videoid, _ := strconv.Atoi(CReq.VideoId)
	//videoId := ctx.Query("video_id")

	ac, _ := strconv.Atoi(CReq.ActionType)
	actionType := int32(ac)
	commentText := ctx.Query("comment_text")
	commId := snowflake.GenID()
	comid := ctx.Query("comment_id")
	commentId, _ := strconv.ParseInt(comid, 10, 64)
	ctx.JSON(http.StatusOK, pb.DouyinCommentActionResponse{
		StatusCode: 200,
		StatusMsg:  "操作成功",
	})

	if CReq.ActionType == "1" { // 进行评论
		CAction := pb.DouyinCommentActionRequest{ // 测试数据，为转换
			Token:       token,
			VideoId:     int64(videoid),
			ActionType:  int32(actionType),
			CommentText: string(commentText),
			CommentId:   int64(commId),
		}

		CResp, err := rpc.CommentAction(ctx, &CAction)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, pb.DouyinCommentActionResponse{
				StatusCode: 500,
				StatusMsg:  "评论失败",
			})
			return
		}
		ctx.JSON(http.StatusOK, CResp)
	} else if CReq.ActionType == "2" { // 删除评论操作
		CAction := pb.DouyinCommentActionRequest{ // 测试数据
			Token:      token,
			ActionType: int32(actionType),
			VideoId:    int64(videoid),
			CommentId:  int64(commentId),
		}
		CResp, err := rpc.CommentAction(ctx, &CAction)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, pb.DouyinCommentActionResponse{
				StatusCode: 500,
				StatusMsg:  "删除评论失败",
			})
			return
		}
		ctx.JSON(http.StatusOK, CResp)
	} else { // 非法操作
		ctx.JSON(http.StatusBadRequest, pb.DouyinCommentActionResponse{
			StatusCode: 400,
			StatusMsg:  "非法操作",
		})
	}
}

// 读取评论列表
func CommentList(ctx *gin.Context) {
	var CReq CommentReq
	CReq.VideoId = ctx.Query("VideoId")
	videoID, _ := strconv.ParseInt(CReq.VideoId, 10, 64)
	var commentList pb.DouyinCommentListRequest
	commentList.VideoId = videoID
	CResp, err := rpc.CommentList(ctx, &commentList)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pb.DouyinCommentListResponse{
			StatusCode: 500,
			StatusMsg:  "获取评论列表失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, CResp)

}
