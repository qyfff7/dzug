package conf

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func ViperInit(config interface{}, YmlFilePath string) (err error) {
	workDir, _ := os.Getwd()
	viper.SetConfigFile(workDir + YmlFilePath)
	if err = viper.ReadInConfig(); err != nil { // 查找并读取配置文件
		panic(fmt.Errorf("viper.ReadInConfig error config file: %s \n", err)) // 处理读取配置文件的错误
		return
	}
	//把读取到的配置信息，反序列化到Conf变量中
	if err = viper.Unmarshal(config); err != nil {
		fmt.Printf("viper.Unmarshal failed ,err %v", err)
	}
	return
}
