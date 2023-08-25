package handlers

import (
	"dzug/app/gateway/rpc"
	"dzug/app/services/user/pkg/jwt"
	model "dzug/models"
	pb "dzug/protos/publish"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ActionResp struct {
	status_code int
	status_msg  string
}

type ListResponse struct {
	status_code int
	status_msg  string
	video_list  []model.Video
}

// UploadHandler 视频投稿
func UploadHandler(ctx *gin.Context) {
	const MaxFileSize = 30 * 1024 * 1024 // 30MB

	file, err := ctx.FormFile("data")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get file"})
		return
	}

	title := ctx.PostForm("title")
	token := ctx.PostForm("token")
	userIdBefore, err := jwt.ParseToken(token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse token"})
		return
	}
	userId := userIdBefore.UserID

	fileReader, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer fileReader.Close()

	// 限制文件大小
	if file.Size > MaxFileSize {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds the limit"})
		return
	}

	// 将视频数据读取到字节切片中
	videoData, err := ioutil.ReadAll(fileReader)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	publishReq := &pb.PublishVideoReq{
		UserId:   userId,
		Data:     videoData,
		Title:    title,
		FileName: file.Filename,
	}

	_, err = rpc.PublishVideo(ctx, publishReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ActionResp{
			status_code: http.StatusOK,
			status_msg:  "Wrong Request",
		})
	}
	rtnResp := ActionResp{
		status_code: http.StatusOK,
		status_msg:  "Success",
	}
	ctx.JSON(http.StatusOK, rtnResp)
}

// GetVideoListByUser 获取用户投稿信息
func GetVideoListByUser(ctx *gin.Context) {
	user_id := ctx.Query("user_id")
	parsedUserId, err := strconv.ParseInt(user_id, 10, 64)
	if err != nil {
		zap.L().Error(err.Error())
	}
	getVideoListReq := &pb.GetVideoListByUserIdReq{
		UserId: parsedUserId,
	}
	getVideoListResp, err := rpc.GetPublishListByUser(ctx, getVideoListReq)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pb.PublishVideoResp{
			StatusCode: 500,
			StatusMsg:  "远程RPC调用异常",
		})
	}
	//rtn := ListResponse{
	//	status_code: http.StatusOK,
	//	status_msg: "调用成功",
	//	video_list: getVideoListResp.VideoList
	//}
	ctx.JSON(http.StatusOK, getVideoListResp)
}
