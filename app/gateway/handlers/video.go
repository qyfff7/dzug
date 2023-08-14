package handlers

import (
	"dzug/app/gateway/rpc"
	pb "dzug/protos/publish"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"net/http"
	"strconv"
)

// UploadHandler 视频投稿
func UploadHandler(ctx *gin.Context) {
	user_id_ := ctx.PostForm("user_id")
	title := ctx.PostForm("title")
	file, err := ctx.FormFile("file")
	fileName := file.Filename

	user_id, _ := strconv.ParseInt(user_id_, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read file"})
		return
	}

	fileReader, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open file"})
		return
	}
	defer fileReader.Close()

	// 将视频数据读取到字节切片中
	var videoData []byte
	buf := make([]byte, 1024)
	for {
		n, err := fileReader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
			return
		}
		videoData = append(videoData, buf[:n]...)
	}

	publishReq := &pb.PublishVideoReq{
		UserId:   user_id,
		Data:     videoData,
		Title:    title,
		FileName: fileName,
	}

	publishVideoResp, err := rpc.PublishVideo(ctx, publishReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pb.PublishVideoResp{
			StatusCode: 500,
			StatusMsg:  "远程RPC调用异常",
		})
		return
	}
	ctx.JSON(http.StatusOK, publishVideoResp)
}

// GetVideoListByUser 获取用户投稿信息
func GetVideoListByUser(ctx *gin.Context) {
	user_id := ctx.PostForm("user_id")
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
	ctx.JSON(http.StatusOK, getVideoListResp)
}
