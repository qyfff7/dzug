package rpc

import (
	"context"
	"dzug/discovery"
	pb "dzug/protos/user"
)

// 这里应该做的事情是与用户服务建立连接，调用远程的方法

func Login(ctx context.Context, req *pb.LoginAndRegisterRequest) (*pb.LoginAndRegisterResponse, error) {

	discovery.LoadClient("user", &discovery.UserClient)
	r, err := discovery.UserClient.Login(ctx, req) // 调用注册方法
	if err != nil {
		return nil, err
	}
	return r, nil

}

func Register(ctx context.Context, req *pb.LoginAndRegisterRequest) (*pb.LoginAndRegisterResponse, error) {

	discovery.LoadClient("user", &discovery.UserClient)
	r, err := discovery.UserClient.Register(ctx, req) // 调用注册方法
	if err != nil {
		return nil, err
	}
	return r, nil

}

func UserInfo(ctx context.Context, req *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	discovery.LoadClient("user", &discovery.UserClient)
	r, err := discovery.UserClient.UserInfo(ctx, req) // 调用注册方法
	if err != nil {
		return nil, err
	}
	return r, nil

}
