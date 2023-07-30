package main

import (
	"dzug/app/relation/service"
	"dzug/discovery"
	pb "dzug/idl/relation"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	endpoints := []string{"localhost:2379"}
	serviceRegister := &discovery.ServiceRegister{
		EtcdAddrs: endpoints,
		Lease:     5,
		Key:       "relation",
		Value:     "127.0.0.1:9001",
	}
	err := serviceRegister.NewServiceRegister()
	if err != nil {
		log.Println(err)
	}
	defer serviceRegister.Close()

	// 创建grpc服务器并监听9001端口
	server := grpc.NewServer()
	defer server.Stop()
	pb.RegisterDouyinRelationActionServiceServer(server, &service.RelationSrv{})

	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Println(err)
	}
	log.Println("listening ")
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
