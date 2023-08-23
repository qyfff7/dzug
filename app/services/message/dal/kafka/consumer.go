package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
	"sync"
)

type handlerMsg func(message *sarama.ConsumerMessage) error

//func MessageConsumer(){
//	fmt.Println("message start consuming")
//	partitionConsumers := make([]sarama.PartitionConsumer, 0)
//	partitionList, err := KafkaConsumer.Partitions("message")
//	if err != nil {
//		fmt.Println("fail to get list of partition,err:", err)
//	}
//	fmt.Println(len(partitionList))
//	for p := range partitionList {
//		//针对每一个分区创建一个对应分区的消费者
//		pc, err := KafkaConsumer.ConsumePartition("message", int32(p), sarama.OffsetNewest)
//		if err != nil {
//			fmt.Printf("failed to start consumer for partition %d,err:%v\n", p, err)
//		}
//		partitionConsumers = append(partitionConsumers, pc)
//		fmt.Println(len(pc.Messages()))
//		//异步从每个分区消费信息
//		go func(sarama.PartitionConsumer) {
//			for {
//				for msg := range pc.Messages() {
//					fmt.Printf("message: topic %s , key %s, value %s\n", msg.Topic, msg.Key, msg.Value)
//					if err := handler(msg); err != nil {
//						zap.L().Fatal("消息消费处理失败", zap.String("partition", string(msg.Partition)), zap.String("Key", string(msg.Key)), zap.String("Value", string(msg.Value)))
//					}
//				}
//			}
//		}(pc)
//	}
//}

var wg sync.WaitGroup

func ConsumeMsg(topic string, handler handlerMsg) {
	//根据topic获取所有的分区列表
	fmt.Println("start consuming")
	partitionConsumers := make([]sarama.PartitionConsumer, 0)
	partitionList, err := KafkaConsumer.Partitions(topic)
	if err != nil {
		fmt.Println("fail to get list of partition,err:", err)
	}
	fmt.Println(len(partitionList))
	pcnum := 0
	for p := range partitionList {
		//针对每一个分区创建一个对应分区的消费者
		pc, err := KafkaConsumer.ConsumePartition(topic, int32(p), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", p, err)
		}
		defer pc.AsyncClose()
		wg.Add(1)
		partitionConsumers = append(partitionConsumers, pc)
		pcnum++
		//异步从每个分区消费信息
		go func(sarama.PartitionConsumer, int) {
			for {
				for msg := range pc.Messages() {
					fmt.Printf("consumer num: %d, message: topic %s , key %s, value %s\n", pcnum, msg.Topic, msg.Key, msg.Value)
					if err := handler(msg); err != nil {
						zap.L().Fatal("消息消费处理失败", zap.String("partition", string(msg.Partition)), zap.String("Key", string(msg.Key)), zap.String("Value", string(msg.Value)))
					}
				}
			}
		}(pc, pcnum)
	}
	defer ClosePartitionConsumer(partitionConsumers)
	wg.Wait()
}

func ClosePartitionConsumer(pcs []sarama.PartitionConsumer) {
	for _, pc := range pcs {
		pc.Close()
	}
}
