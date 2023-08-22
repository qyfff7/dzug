package tailfile

import (
	"fmt"
	"github.com/hpcloud/tail"
	"go.uber.org/zap"
)

var (
	TailObj *tail.Tail
)

// tail相关

func Init(filename string) (err error) {

	config := tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	// 打开文件开始读取数据
	TailObj, err = tail.TailFile(filename, config)
	if err != nil {
		zap.L().Error("tailfile: create tailObj for path:%s "+fmt.Sprintf("%s", filename)+" failed, err: ", zap.Error(err))
		return
	}
	return
}
