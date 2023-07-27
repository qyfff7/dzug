package main

import (
	"dzug/app/user/service"
	"dzug/discovery"
	pb "dzug/idl"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	endpoints := []string{"localhost:2379"}
	// lease 应该是租约时间，这里是5秒
	etcdRegister, err := discovery.NewServiceRegister(endpoints, "user", "localhost:9000", 5)
	if err != nil {
		log.Fatal(err)
	}
	defer etcdRegister.Close()
	go etcdRegister.ListenLeaseRespChan() // 启用协程，监听续租响应通道

	// 创建grpc服务器并监听9000端口
	server := grpc.NewServer()
	defer server.Stop()
	pb.RegisterDouyinUserServiceServer(server, &service.UserSrv{})

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("listening ")
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
