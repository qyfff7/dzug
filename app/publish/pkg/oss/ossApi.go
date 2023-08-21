package oss

import (
	"bytes"
	"context"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
	"io"
	"strconv"
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

	tmp := ossVideo

	bucket, err := client.Bucket(tmp.Bucket)
	if err != nil {
		zap.L().Error(err.Error())
	}

	videoFileName := strconv.FormatInt(v.UserID, 10) + "/" + v.FileName
	videoObjectKey := "play/" + videoFileName

	var file io.Reader
	file = bytes.NewReader(v.File)

	err = bucket.PutObject(videoObjectKey, file)
	if err != nil {
		zap.L().Error(err.Error())
		return nil, err
	}
	dataUrl := ossVideo.Bucket + "." + ossVideo.EndPoint + "/" + videoObjectKey
	coverUrl := dataUrl + "?x-oss-process=video/snapshot,t_0,f_jpg"

	return &VideoUrl{
		PlayUrl:  dataUrl,
		CoverUrl: coverUrl,
	}, nil
}
