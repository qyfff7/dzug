package rpc

import (
	"context"
	"dzug/discovery"
	pb "dzug/idl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func UserLogin(ctx context.Context, req *pb.DouyinUserLoginRequest) (resp *pb.DouyinUserLoginResponse, err error) {
	var endpoints = []string{"localhost:2379"}
	ser := discovery.NewServiceDiscovery(endpoints)
	defer ser.Close()
	ser.WatchService("user")
	target := ser.GetServices()[0]
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials())) // grpc.WithInsecure() // 不使用TLS认证
	if err != nil {
		log.Fatalf("net.Connect err : %v", err)
	}
	defer conn.Close()

	UserClient := pb.NewDouyinUserServiceClient(conn)
	r, err := UserClient.Login(ctx, req)
	if err != nil {
		return
	}
	return r, nil
}

func UserRegister(ctx context.Context, req *pb.DouyinUserRegisterRequest) (resp *pb.DouyinUserRegisterResponse, err error) {
	var endpoints = []string{"localhost:2379"}
	ser := discovery.NewServiceDiscovery(endpoints)
	defer ser.Close()
	ser.WatchService("user")
	target := ser.GetServices()[0]
	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials())) // grpc.WithInsecure() // 不使用TLS认证
	if err != nil {
		log.Fatalf("net.Connect err : %v", err)
	}
	defer conn.Close()

	UserClient := pb.NewDouyinUserServiceClient(conn)
	r, err := UserClient.Register(ctx, req)
	if err != nil {
		return
	}
	return r, nil
}
