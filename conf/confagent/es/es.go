package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"go.uber.org/zap"
)

// 将日志数据写入Elasticsearch

type ESClient struct {
	client      *elastic.Client  //es client
	index       string           //index
	logDataChan chan interface{} //接收数据的channel
}

var (
	esClient *ESClient //这里声明了一个空指针
)

func Init(addr, index string, goroutineNum, maxSize int) (err error) {
	client, err := elastic.NewClient(elastic.SetURL("http://" + addr))
	if err != nil {
		// Handle error
		panic(err)
	}
	//fmt.Printf("%#v\n", client)
	zap.L().Info(fmt.Sprintf("%s", client))
	esClient = &ESClient{ //  空指针不可以直接赋值，因此得初始化
		client:      client,
		index:       index,
		logDataChan: make(chan interface{}, maxSize),
	}

	zap.L().Info("connect to es success")

	// 从通道中取出数据,写入到kafka中去
	for i := 0; i < goroutineNum; i++ {
		go sendToES() //从配置文件中获取到需要的goroutine数量，动态的起goroutine
	}
	return
}

func sendToES() {
	for m1 := range esClient.logDataChan {
		_, err := esClient.client.Index().
			Index(esClient.index).
			BodyJson(m1).
			Do(context.Background())
		if err != nil {
			// Handle error
			panic(err)
		}
	}
}

// PutLogData 通过一个首字母大写的函数从包外接收msg,发送到chan中
func PutLogData(msg interface{}) {
	esClient.logDataChan <- msg
}
