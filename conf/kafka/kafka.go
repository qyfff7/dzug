package kafka

import (
	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

// kafka相关操作
var (
	client  sarama.SyncProducer
	msgChan chan *sarama.ProducerMessage //日志消息channel
)

// Init 是初始化全局的kafka Client
func Init(address []string, chanSize int64) (err error) { //参数：一个是地址，一个是chasize ,也就是在config.ini文件里面定义的字段
	// 1. 生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // ACK
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 分区
	config.Producer.Return.Successes = true                   // 确认

	// 2. 连接kafka
	client, err = sarama.NewSyncProducer(address, config)
	if err != nil {
		zap.L().Error("kafka:producer closed, err:", zap.Error(err))
		return
	}

	// 初始化MsgChan
	msgChan = make(chan *sarama.ProducerMessage, chanSize)
	// 起一个后台的goroutine从msgchan中读数据
	go sendMsg()
	return
}

// 从MsgChan中读取msg,发送给kafka
func sendMsg() {
	for {
		select {
		case msg := <-msgChan:
			//, offset, err := client.SendMessage(msg)
			_, _, err := client.SendMessage(msg)
			//_, _, err := client.SendMessage(msg)
			if err != nil {
				zap.L().Warn("send msg failed, err:", zap.Error(err))
				return
			}
			//zap.L().Info("send msg to kafka success." + fmt.Sprintf("pid:%v  offset:%v", pid, offset))
		}
	}
}

// ToMsgChan 定义一个函数向外暴露msgChan
func ToMsgChan(msg *sarama.ProducerMessage) {
	msgChan <- msg
}
