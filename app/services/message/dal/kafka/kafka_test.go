package kafka

import (
	"fmt"
	"github.com/IBM/sarama"
	"testing"
	"time"
)

func handle(message *sarama.ConsumerMessage) error {
	fmt.Printf("message: topic %s , key %s, value %s\n\n", message.Topic, message.Key, message.Value)
	return nil
}

func TestKafka(t *testing.T) {
	InitConsumer()
	go ConsumeMsg("dousheng", handle)
	time.Sleep(1 * time.Second)
	InitProducer()
	SendMsg("dousheng", "test", "hello world")
	time.Sleep(5 * time.Second)
}
