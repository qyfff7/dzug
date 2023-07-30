package rpc

import (
	"dzug/discovery"
	"dzug/idl/relation"
	"dzug/idl/user"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var (
	Ser            discovery.ServiceDiscovery
	UserClient     user.DouyinUserServiceClient
	RelationClient relation.DouyinRelationActionServiceClient
)

func Init() {
	endpoints := []string{"localhost:2379"}
	Ser = discovery.ServiceDiscovery{EtcdAddrs: endpoints}
	err := Ser.NewServiceDiscovery()
	if err != nil {
		log.Fatal("启动服务发现失败")
		return
	}
	//loadClient("user", &UserClient)
}

func loadClient(serviceName string, client any) {
	conn, err := connectService(serviceName)
	if err != nil {
		fmt.Println("grpc连接服务: ", serviceName, "失败, error: ", err)
		return
	}

	switch c := client.(type) {
	case *user.DouyinUserServiceClient:
		*c = user.NewDouyinUserServiceClient(conn)
	case *relation.DouyinRelationActionServiceClient:
		*c = relation.NewDouyinRelationActionServiceClient(conn)
	default:
		panic("没有这种类型的服务")
	}
}

func connectService(serviceName string) (conn *grpc.ClientConn, err error) {
	err = Ser.WatchService(serviceName)
	if err != nil {
		return nil, err
	}
	conn, err = grpc.Dial(Ser.GetServiceByKey(serviceName), grpc.WithTransportCredentials(insecure.NewCredentials()))
	return
}
