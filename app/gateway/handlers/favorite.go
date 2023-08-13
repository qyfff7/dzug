package handlers

import (
	"dzug/app/gateway/rpc"
	"dzug/app/user/pkg/jwt"
	pb "dzug/protos/favorite"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type favoriteReq struct {
	Token      string `json:"token"` // 暂时把token当作userid ！！！！！！！！！！！！！
	VideoId    string `json:"video_id"`
	ActionType string `json:"action_type"`
}

// FavoriteAction 点赞操作
func FavoriteAction(ctx *gin.Context) {
	var fReq favoriteReq
	if err := ctx.BindJSON(&fReq); err != nil {
		ctx.JSON(http.StatusBadRequest, pb.FavoriteResponse{
			StatusCode: 400,
			StatusMsg:  "传入参数错误",
		})
		return
	}
	userId, _ := jwt.GetUserID(ctx)
	videoId := fReq.VideoId
	fmt.Println(userId, videoId)

	ctx.JSON(http.StatusOK, pb.FavoriteResponse{
		StatusCode: 200,
		StatusMsg:  "操作成功",
	})

	//if fReq.ActionType == "1" { // 进行点赞
	//	//strconv.Atoi(fReq.Token) 仅作测试
	//	fAction := pb.FavoriteRequest{ // 测试数据，为转换
	//		UserId:  1,
	//		VideoId: 1,
	//	}
	//	fResp, err := rpc.FavoriteAction(ctx, &fAction)
	//	if err != nil {
	//		ctx.JSON(http.StatusInternalServerError, pb.FavoriteResponse{
	//			StatusCode: 500,
	//			StatusMsg:  "点赞失败",
	//		})
	//		return
	//	}
	//	ctx.JSON(http.StatusOK, fResp)
	//} else if fReq.ActionType == "2" { // 取消点赞操作
	//	fAction := pb.InfavoriteRequest{ // 测试数据
	//		UserId:  1,
	//		VideoId: 1,
	//	}
	//	fResp, err := rpc.InFavorite(ctx, &fAction)
	//	if err != nil {
	//		ctx.JSON(http.StatusInternalServerError, pb.InfavoriteResponse{
	//			StatusCode: 500,
	//			StatusMsg:  "点赞失败",
	//		})
	//		return
	//	}
	//	ctx.JSON(http.StatusOK, fResp)
	//} else { // 非法操作
	//	ctx.JSON(http.StatusBadRequest, pb.FavoriteResponse{
	//		StatusCode: 400,
	//		StatusMsg:  "非法操作",
	//	})
	//}
}

// FavoriteList 获取点赞列表
func FavoriteList(ctx *gin.Context) {
	userId, _ := jwt.GetUserID(ctx)
	var favoriteList pb.FavoriteListRequest
	favoriteList.UserId = int64(userId)
	if err := ctx.BindJSON(&favoriteList); err != nil {
		ctx.JSON(http.StatusBadRequest, pb.FavoriteListResponse{
			StatusCode: 400,
			StatusMsg:  "参数错误",
		})
		return
	}
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
