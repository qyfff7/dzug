package handlers

import (
	"dzug/app/gateway/rpc"
	"dzug/discovery"
	"dzug/models"
	"dzug/protos/user"
	pb "dzug/protos/video"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
)

// Feed 视频流
func Feed(c *gin.Context) {

	//1.新建视频流请求参数
	feedparam := new(pb.GetVideoListByTimeReq)
	if err := c.ShouldBind(feedparam); err != nil { //获取参数
		zap.L().Error("获取视频流参数出错", zap.Error(err))
		return
	}
	//2.判断当前用户是否登录
	authHeader := c.Request.Header.Get("Authorization") //ctx 是 Context
	parts := strings.SplitN(authHeader, " ", 2)

	if parts[1] == "" { //如果未登录，Token为空
		feedparam.Token = ""
	} else {
		feedparam.Token = parts[1]
	}
	//视频最新投稿时间，不填则表示当前时间
	if feedparam.LatestTime <= 0 {
		feedparam.LatestTime = time.Now().Unix()
	}
	if feedparam.LatestTime > time.Now().Unix() {
		feedparam.LatestTime = time.Now().Unix()
	}
	//3.获取视频流
	videos, err := rpc.Feed(c, feedparam)
	if err != nil {
		zap.L().Error("获取视频流出错", zap.Error(err))
		return
	}

	feed := make([]*models.Video, 0, len(videos.VideoList))

	//4. 为每个视频查询作者信息
	for _, video := range videos.VideoList {
		// 根据作者id查询作者信息
		auther := &user.GetUserInfoReq{UserId: video.AutherId, Token: feedparam.Token}
		autherInfo, err := discovery.UserClient.GetUserInfo(c, auther) // 调用查询用户信息的方法
		if err != nil {
			zap.L().Error("获取视频作者信息失败")
			continue
		}
		//查询当前用户是否给当前视频点赞

		v := &models.Video{
			Id:            video.VideoId,
			Auther:        models.UserInfoResp(autherInfo),
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    video.IsFavorite,
			Title:         video.Title,
		}
		feed = append(feed, v)
	}
	c.JSON(http.StatusOK, feed)

}
