package service

import (
	"context"
	"dzug/idl/user"
)

type UserSrv struct {
	user.UnimplementedDouyinUserServiceServer
}

func (u *UserSrv) Login(context.Context, *user.DouyinUserLoginRequest) (*user.DouyinUserLoginResponse, error) {
	return &user.DouyinUserLoginResponse{
		StatusCode: 200,
		StatusMsg:  "登录成功",
		UserId:     1,
		Token:      "test",
	}, nil
}

func (u *UserSrv) Register(context.Context, *user.DouyinUserRegisterRequest) (*user.DouyinUserRegisterResponse, error) {
	return &user.DouyinUserRegisterResponse{
		StatusCode: 200,
		StatusMsg:  "注册成功",
		UserId:     1,
		Token:      "test",
	}, nil
}
