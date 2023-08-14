package handlers

import (
	"dzug/app/gateway/rpc"
	"dzug/app/user/pkg/jwt"
	pb "dzug/protos/favorite"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type favoriteReq struct {
	VideoId    string `json:"video_id"`
	ActionType string `json:"action_type"`
}

// FavoriteAction 点赞操作
func FavoriteAction(ctx *gin.Context) {
	var fReq favoriteReq
	userId, _ := jwt.GetUserID(ctx)
	err := ctx.ShouldBind(&fReq)
	if err != nil {
		zap.L().Fatal("绑定参数出错" + err.Error())
	}
	zap.L().Info(fmt.Sprintf("userId:", userId, " VideoId:", fReq.VideoId, " ActionType:", fReq.ActionType))
	videoid, _ := strconv.Atoi(fReq.VideoId)
	//videoId := ctx.Query("video_id")
	//actionType := ctx.Query("action_type")

	if fReq.ActionType == "1" { // 进行点赞
		fAction := pb.FavoriteRequest{ // 测试数据，为转换
			UserId:  userId,
			VideoId: int64(videoid),
		}
		fResp, err := rpc.FavoriteAction(ctx, &fAction)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, pb.FavoriteResponse{
				StatusCode: 500,
				StatusMsg:  "点赞失败",
			})
			return
		}
		ctx.JSON(http.StatusOK, fResp)
	} else if fReq.ActionType == "2" { // 取消点赞操作
		fAction := pb.InfavoriteRequest{ // 测试数据
			UserId:  userId,
			VideoId: int64(videoid),
		}
		fResp, err := rpc.InFavorite(ctx, &fAction)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, pb.InfavoriteResponse{
				StatusCode: 500,
				StatusMsg:  "取消点赞失败",
			})
			return
		}
		ctx.JSON(http.StatusOK, fResp)
	} else { // 非法操作
		ctx.JSON(http.StatusBadRequest, pb.FavoriteResponse{
			StatusCode: 400,
			StatusMsg:  "非法操作",
		})
	}
}

// FavoriteList 获取点赞列表
func FavoriteList(ctx *gin.Context) {
	userId, _ := jwt.GetUserID(ctx)
	var favoriteList pb.FavoriteListRequest
	favoriteList.UserId = userId
	fResp, err := rpc.FavoriteList(ctx, &favoriteList)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pb.FavoriteListResponse{
			StatusCode: 500,
			StatusMsg:  "获取点赞列表失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, fResp)
}
