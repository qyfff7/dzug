package handlers

import (
	"dzug/app/gateway/rpc"
	"dzug/app/user/pkg/jwt"
	pb "dzug/protos/video"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func List(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"msg": "cishi",
	})
}

// Feed 视频流
func Feed(c *gin.Context) {

	//1.新建视频流请求参数
	feedparam := new(pb.GetVideoFeedReq)
	if err := c.ShouldBind(feedparam); err != nil { //获取参数
		zap.L().Error("获取视频流参数出错", zap.Error(err))
		return
	}
	//2.判断当前用户是否登录
	_, err := jwt.GetUserID(c)
	if err != nil { //如果未登录，则uid用-1表示,Token为空
		//uid = -1
		feedparam.Token = ""
	}
	//视频最新投稿时间，不填则表示当前时间
	if feedparam.LatestTime <= 0 {
		feedparam.LatestTime = time.Now().Unix()
	}
	if feedparam.LatestTime > time.Now().Unix() {
		feedparam.LatestTime = time.Now().Unix()
	}
	//3.获取视频流
	videoList, err := rpc.Feed(c, feedparam)
	if err != nil {
		zap.L().Error("获取视频流出错", zap.Error(err))

		return
	}
	//4. 返回相应
	resp := &pb.GetVideoFeedResp{
		StatusCode: 0,
		StatusMsg:  "获取视频流成功",
		NextTime:   videoList.NextTime,
		VideoInfos: videoList.VideoInfos,
	}
	c.JSON(http.StatusOK, resp)

}
