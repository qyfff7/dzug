package service

import (
	"context"
	"dzug/idl"
)

type UserSrv struct {
	__.UnimplementedDouyinUserServiceServer
}

func (u *UserSrv) Login(context.Context, *__.DouyinUserLoginRequest) (*__.DouyinUserLoginResponse, error) {
	return &__.DouyinUserLoginResponse{
		StatusCode: 200,
		StatusMsg:  "登录成功",
		UserId:     1,
		Token:      "test",
	}, nil
}

func (u *UserSrv) Register(context.Context, *__.DouyinUserRegisterRequest) (*__.DouyinUserRegisterResponse, error) {
	return &__.DouyinUserRegisterResponse{
		StatusCode: 200,
		StatusMsg:  "注册成功",
		UserId:     1,
		Token:      "test",
	}, nil
}
