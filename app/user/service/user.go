package service

import (
	"context"
	pb "dzug/idl"
)

type UserSrv struct {
	pb.UnimplementedDouyinUserServiceServer
}

func (u *UserSrv) Login(context.Context, *pb.DouyinUserLoginRequest) (*pb.DouyinUserLoginResponse, error) {
	return &pb.DouyinUserLoginResponse{
		StatusCode: 200,
		StatusMsg:  "登录成功",
		UserId:     1,
		Token:      "test",
	}, nil
}

func (u *UserSrv) Register(context.Context, *pb.DouyinUserRegisterRequest) (*pb.DouyinUserRegisterResponse, error) {
	return &pb.DouyinUserRegisterResponse{
		StatusCode: 200,
		StatusMsg:  "注册成功",
		UserId:     1,
		Token:      "test",
	}, nil
}
