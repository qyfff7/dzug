package handlers

import (
	"dzug/app/gateway/rpc"
	"dzug/app/user/pkg/jwt"
	"dzug/models"
	"dzug/protos/user"
	"dzug/protos/video"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

/*// Feed same demo video list for every request
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, models.FeedResponse{
		Response: models.Response{StatusCode: models.CodeSuccess,
			StatusMsg: models.CodeSuccess.Msg()},
		VideoList: models.DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}*/

// Feed 视频流
func Feed(c *gin.Context) {

	//1.新建视频流请求参数
	vparams := new(video.GetVideoListByTimeReq)
	if err := c.ShouldBind(vparams); err != nil { //获取参数与参数校验
		zap.L().Error("Feed with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			models.ResponseError(c, models.CodeInvalidParam)
			return
		}
		err, _ := json.Marshal(removeTopStruct(errs.Translate(trans)))
		models.ResponseErrorWithMsg(c, models.CodeInvalidParam, string(err))
		return
	}
	//2.判断当前是否登录
	if vparams.Token != "" {
		//当前登录，进行token校验
		u, err := jwt.ParseToken(vparams.Token)
		if err != nil {
			zap.L().Error(fmt.Sprintln(models.CodeInvalidToken))
		}
		// 将当前请求的userID信息保存到请求的上下文ctx上
		c.Set(jwt.CtxUserIDKey, u.UserID)
	}

	//3.调用获取视频流服务
	videos, err := rpc.Feed(c, vparams)
	if err != nil {
		zap.L().Error("rpc调用视频流服务出错", zap.Error(err))
		return
	}

	videofeed := make([]*models.Video, 0, len(videos.VideoList))
	//4.对于每个视频,查询作者的信息
	for _, v := range videos.VideoList {
		u := &user.GetUserInfoReq{
			UserId: v.AutherId,
			Token:  vparams.Token,
		}
		zap.L().Info("查询视频作者信息" + fmt.Sprintln(u.UserId))
		zap.L().Info("查询视频作者信息" + fmt.Sprintln(u.Token))

		authorInfo, err := rpc.UserInfo(c, u)
		if err != nil {
			zap.L().Error("获取视频作者信息失败", zap.Error(err))
		}
		author := models.UserInfoResp(authorInfo)
		v := models.GetFeedResp(v, author)
		videofeed = append(videofeed, v)
	}

	models.GetFeedSuccess(c, videofeed, videos.NextTime)

	/*feed := make([]*models.Video, 0, len(videos.VideoList))
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
	c.JSON(http.StatusOK, feed)*/

}
