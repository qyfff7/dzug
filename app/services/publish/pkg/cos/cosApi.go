package cos

import (
	"bytes"
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
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

func UploadVideo(ctx context.Context, v *Video) (*VideoUrl, error) {
	u, _ := url.Parse(cosVideo.VideoBucket)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  cosVideo.SecretId,
			SecretKey: cosVideo.SecretKey,
		},
	})

	videoFileName := strconv.FormatInt(v.UserID, 10) + "/" + v.FileName
	replaceSuffixidx := strings.LastIndex(v.FileName, ".")
	coverFileName := v.FileName[0:replaceSuffixidx] + ".jpg"

	// 上传视频文件
	var file io.Reader
	file = bytes.NewReader(v.File)
	_, err := c.Object.Put(ctx, videoFileName, file, nil)
	if err != nil {
		return nil, err
	}
	videourl := &VideoUrl{
		PlayUrl:  cosVideo.VideoBucket + "/" + videoFileName,
		CoverUrl: cosVideo.CoverBucket + "/cover/" + coverFileName,
	}
	// 上传成功 返回key
	return videourl, nil
}
