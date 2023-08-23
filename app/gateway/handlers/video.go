package handlers

import (
	"dzug/app/gateway/rpc"
	"dzug/app/services/user/pkg/jwt"
	"dzug/models"
	"dzug/protos/user"
	"dzug/protos/video"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

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
		authorInfo, err := rpc.UserInfo(c, u)
		if err != nil {
			zap.L().Error("获取视频作者信息失败", zap.Error(err))
		}
		author := models.UserInfoResp(authorInfo)
		v := models.GetFeedResp(v, author)
		videofeed = append(videofeed, v)
	}
	models.GetFeedSuccess(c, videofeed, videos.NextTime)
}
