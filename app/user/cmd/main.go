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
	etcdRegister, err := discovery.NewServiceRegister(endpoints, "user", "localhost:9000", 5)
	if err != nil {
		log.Fatal(err)
	}
	defer etcdRegister.Close()

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
