package oss

type OssVideo struct {
	Bucket          string
	EndPoint        string
	AccessKeyID     string
	AccessKeySecret string
}

var ossVideo OssVideo

func Init() {
	ossVideo = OssVideo{
		Bucket:          "byte-camp-video",
		EndPoint:        "oss-cn-beijing.aliyuncs.com",
		AccessKeyID:     "LTAI5tFbJAeUWwCP5uCunYt2",
		AccessKeySecret: "zEcp5j5uvRuNYHoyFo7scab1qyeB20",
	}
}
