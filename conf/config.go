package conf

import (
	"dzug/conf/confagent/etcd"
	"dzug/conf/confagent/kafka"
	"dzug/conf/confagent/tailfile"
	"dzug/logger"
	"dzug/models"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"strings"
	"time"
)

// Config 全局变量，用来保存项目所有的配置信息
var Config = new(models.ProjectConfig)
var BasicConf = new(models.BasicConfig)

// Init 项目配置初始化
func Init() (err error) {

	err = viperInit()
	if err != nil {
		fmt.Println("init basic config failed, err:" + err.Error())
		return err
	}
	//1. 初始化etcd连接
	err = etcd.Init(BasicConf.EtcdAddr, BasicConf.Name)
	if err != nil {
		//zap.L().Error("init etcd failed, err:", zap.Error(err))
		fmt.Println("init etcd failed, err:" + err.Error())
		return
	}
	// 2.从etcd中拉取项目所有的配置项
	Config, err = etcd.GetProjectConf(BasicConf.Name)
	if err != nil {
		//zap.L().Error("get conf from etcd failed, err:", zap.Error(err))
		fmt.Printf("get conf from etcd failed, err:%s", err)
		return
	}

	//初始化日志
	if err = logger.Init(Config.LogConfig); err != nil {
		fmt.Printf("log file initialization error,%#v", err)
		return
	}
	defer zap.L().Sync() //把缓冲区的日志，追加到文件中
	zap.L().Info("服务启动，开始记录日志")

	//3. 初始化连接kafka生产者(做好准备工作)     (初始化kafka,初始化msg chan，起后台gorountine 去往kafka发msg)
	err = kafka.Init([]string{Config.KafkaConfig.Addr}, Config.KafkaConfig.ChanSize)
	if err != nil {
		zap.L().Error("init kafka failed, err:%v", zap.Error(err))
		return
	}
	zap.L().Info("init kafka success!")

	// 5. 根据配置中的日志路径初始化tail   （根据配置文件中指定的路径创建了一个对应的tailObj）
	err = tailfile.Init(Config.LogConfig.Path)
	if err != nil {
		zap.L().Error("init tailfile failed, err:%v", zap.Error(err))
		return
	}
	zap.L().Info("init tailfile success!")
	return nil
}

func Collectlog() (err error) {
	err = confrun(Config.LogConfig.Topic)
	if err != nil {
		zap.L().Error("Error sending log data to kafka : ", zap.Error(err))
		return
	}
	return
}

// run 真正的业务逻辑
func confrun(topic string) (err error) {

	for {
		// 循环读数据
		line, ok := <-tailfile.TailObj.Lines // chan tail.Line
		if !ok {
			zap.L().Warn("tail file close reopen, filename: " + fmt.Sprintf("%s", tailfile.TailObj.Filename))
			time.Sleep(time.Second) // 读取出错等一秒
			continue
		}
		// 如果是空行就略过

		if len(strings.Trim(line.Text, "\r")) == 0 { //strings.Trim  用来去除  "\r"
			zap.L().Info("出现空行拉,直接跳过...")
			continue
		}

		msg := &sarama.ProducerMessage{}
		msg.Topic = topic
		msg.Value = sarama.StringEncoder(line.Text)
		// 丢到通道中
		kafka.ToMsgChan(msg)
	}
}

func viperInit() (err error) {

	workDir, _ := os.Getwd()                    // 获取当前文件夹路径
	viper.SetConfigName("config")               // 配置文件名
	viper.SetConfigType("yml")                  // 配置文件格式
	viper.AddConfigPath(workDir + "/conf")      // 添加配置路径
	if err = viper.ReadInConfig(); err != nil { // 查找并读取配置文件
		panic(fmt.Errorf("viper.ReadInConfig error config file: %s \n", err)) // 处理读取配置文件的错误
		return
	}
	//把读取到的配置信息，反序列化到Conf变量中
	if err = viper.Unmarshal(BasicConf); err != nil {
		fmt.Printf("viper.Unmarshal failed ,err %v", err)
	}

	return

}
