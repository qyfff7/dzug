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
	//err := ctx.ShouldBind(&fReq)
	//if err != nil {
	//	zap.L().Fatal("绑定参数出错" + err.Error())
	//}
	zap.L().Info(fmt.Sprintf("userId:", userId, " VideoId:", fReq.VideoId, " ActionType:", fReq.ActionType))
	fReq.VideoId = ctx.Query("video_id")
	fReq.ActionType = ctx.Query("action_type")
	videoid, _ := strconv.Atoi(fReq.VideoId)

	if fReq.ActionType == "1" { // 进行点赞
		fAction := pb.FavoriteRequest{
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
<<<<<<< HEAD
=======
		fResp.StatusCode = 0
>>>>>>> feat(-): message module
		ctx.JSON(0, fResp)
	} else if fReq.ActionType == "2" { // 取消点赞操作
		fAction := pb.InfavoriteRequest{
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
<<<<<<< HEAD
=======
		fResp.StatusCode = 0
>>>>>>> feat(-): message module
		ctx.JSON(0, fResp)
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
	resp := convert(fResp)
	ctx.JSON(0, resp)
	//ctx.JSON(http.StatusOK, fResp)
}

func convert(fResp *pb.FavoriteListResponse) *response {
	var resp response
	// 将fRsp转换为resp
	resp.StatusCode = fResp.StatusCode
	resp.StatusMsg = fResp.StatusMsg
	resp.VideoList = make([]*Video, len(fResp.VideoList))
	for k, v := range fResp.VideoList {
		var video Video
		video.Id = v.Id
		video.Author = &User{
			Id:              v.Author.Id,
			Name:            v.Author.Name,
			FollowCount:     v.Author.FollowCount,
			FollowerCount:   v.Author.FollowerCount,
			IsFollow:        v.Author.IsFollow,
			Avatar:          v.Author.Avatar,
			BackgroundImage: v.Author.BackgroundImage,
			Signature:       v.Author.Signature,
			TotalFavorited:  strconv.FormatInt(v.Author.TotalFavorited, 10),
			WorkCount:       v.Author.WorkCount,
			FavoriteCount:   v.Author.FavoriteCount,
		}
		video.PlayUrl = v.PlayUrl
		video.CoverUrl = v.CoverUrl
		video.FavoriteCount = v.FavoriteCount
		video.CommentCount = v.CommentCount
		video.IsFavorite = v.IsFavorite
		video.Title = v.Title
		resp.VideoList[k] = &video
	}
	return &resp
}

type response struct {
	StatusCode int32    `json:"status_code"`
	StatusMsg  string   `json:"status_msg"`
	VideoList  []*Video `json:"video_list"`
}

type Video struct {
	Id            int64  `json:"id"`             // 视频唯一标识
	Author        *User  `json:"author"`         // 视频作者信息
	PlayUrl       string `json:"play_url"`       // 视频播放地址
	CoverUrl      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	Title         string `json:"title"`
}

type User struct {
	Id              int64  `json:"id"`               // 用户id
	Name            string `json:"name"`             // 用户名称
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	Avatar          string `json:"avatar"`           //用户头像
	BackgroundImage string `json:"background_image"` //用户个人页顶部大图
	Signature       string `json:"signature"`        //个人简介
	TotalFavorited  string `json:"total_favorited"`  //获赞数量
	WorkCount       int64  `json:"work_count"`       //作品数量
	FavoriteCount   int64  `json:"favorite_count"`   //点赞数量
}
