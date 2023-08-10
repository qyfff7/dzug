package oss

import (
	"bytes"
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

type Video struct {
	Title    string
	FileName string
	File     []byte
	UserID   int64
}

type VideoUrl struct {
	PlayUrl  string
	CoverUrl string
}

// UploadVideoToOss 视频上传
func UploadVideoToOss(ctx context.Context, v *Video) (*VideoUrl, error) {
	client, err := oss.New(ossVideo.EndPoint, ossVideo.AccessKeyID, ossVideo.AccessKeySecret)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	bucket, err := client.Bucket(ossVideo.Bucket)
	if err != nil {
		zap.L().Error(err.Error())
	}

	videoFileName := strconv.FormatInt(v.UserID, 10) + "/" + v.FileName
	replaceSuffixidx := strings.LastIndex(v.FileName, ".")
	coverFileName := strconv.FormatInt(v.UserID, 10) + "/" + v.FileName[0:replaceSuffixidx] + "_0.jpg"

	videoObjectKey := "play/" + videoFileName
	coverObjectKey := "cover/" + coverFileName

	err = bucket.PutObject(videoObjectKey, bytes.NewReader(v.File))
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}

	return &VideoUrl{
		PlayUrl:  ossVideo.Bucket + "." + ossVideo.EndPoint + "/" + videoObjectKey,
		CoverUrl: ossVideo.Bucket + "." + ossVideo.EndPoint + "/" + coverObjectKey,
	}, err
}
