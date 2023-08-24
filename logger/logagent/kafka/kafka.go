package kafka

import (
	"dzug/logger/logagent/es"
	"encoding/json"
	"fmt"
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
			_, _, err := client.SendMessage(msg)
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

// ConsumerInit 初始化kafka连接   从kafka里面取出日志数据
func ConsumerInit(addr []string, topic string) (err error) {
	// 创建新的消费者
	consumer, err := sarama.NewConsumer(addr, nil)
	if err != nil {
		//fmt.Printf("fail to start consumer, err:%v\n", err)
		zap.L().Error("fail to start consumer, err: ", zap.Error(err))
		return
	}
	// 拿到指定topic下面的所有分区列表
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	//fmt.Println(partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		var pc sarama.PartitionConsumer
		pc, err = consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		//defer pc.AsyncClose()
		// 异步从每个分区消费信息
		zap.L().Info("start to consume...")
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				//logDataChan<-msg // 为了将同步流程异步化,所以将取出的日志数据先放到channel中
				//fmt.Println(msg.Topic, string(msg.Value))
				var m1 map[string]interface{}
				err = json.Unmarshal(msg.Value, &m1)
				if err != nil {
					fmt.Printf("unmarshal msg failed, err:%v\n", err)
					continue
				}
				es.PutLogData(m1)
			}
		}(pc)
	}
	return
}
