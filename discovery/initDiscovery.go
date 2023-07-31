package discovery

import (
	"dzug/idl/relation"
	"dzug/idl/user"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var (
	SerDiscovery ServiceDiscovery

	UserClient     user.DouyinUserServiceClient
	RelationClient relation.DouyinRelationActionServiceClient
)

// InitDiscovery 初始化一个服务发现程序
func InitDiscovery() {
	endpoints := []string{"localhost:2379"}               // etcd地址
	SerDiscovery = ServiceDiscovery{EtcdAddrs: endpoints} // 放入etcd地址
	err := SerDiscovery.NewServiceDiscovery()             // 实例化
	if err != nil {
		log.Println("启动服务发现失败: ", err)
		return
	}
}

// LoadClient 加载etcd客户端调用实例，每一次客户端调用一个方法都会调用这个方法
// 先去etcd中拿去现在的链接，再去通过grpc进行远程调用
func LoadClient(serviceName string, client any) {
	conn, err := connectService(serviceName) // 找到grpc连接链接
	if err != nil {
		log.Println("grpc连接服务: ", serviceName, "失败, error: ", err)
		return
	}

	switch c := client.(type) {
	case *user.DouyinUserServiceClient:
		*c = user.NewDouyinUserServiceClient(conn)
	case *relation.DouyinRelationActionServiceClient:
		*c = relation.NewDouyinRelationActionServiceClient(conn)
	default:
		fmt.Println("没有这种类型的服务")
	}
}

// connectService 通过服务名字找到对应的链接
// 比如，传入user，会找到etcd上存储的user的链接
func connectService(serviceName string) (conn *grpc.ClientConn, err error) {
	err = SerDiscovery.WatchService(serviceName)
	if err != nil {
		return nil, err
	}
	conn, err = grpc.Dial(SerDiscovery.GetServiceByKey(serviceName), grpc.WithTransportCredentials(insecure.NewCredentials()))
	return
}
