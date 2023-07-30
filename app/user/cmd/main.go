package main

import (
	"dzug/app/user/service"
	"dzug/discovery"
	pb "dzug/idl/user"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	endpoints := []string{"localhost:2379"}
	serviceRegister := &discovery.ServiceRegister{
		EtcdAddrs: endpoints,
		Lease:     5,
		Key:       "user",
		Value:     "127.0.0.1:9000",
	}
	err := serviceRegister.NewServiceRegister()
	if err != nil {
		log.Println(err)
	}
	defer serviceRegister.Close()

	// 创建grpc服务器并监听9000端口
	server := grpc.NewServer()
	defer server.Stop()
	pb.RegisterDouyinUserServiceServer(server, &service.UserSrv{})

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Println(err)
	}
	log.Println("listening ")
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
