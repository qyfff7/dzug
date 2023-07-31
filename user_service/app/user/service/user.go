package service

import (
	"context"
	user2 "dzug/user_service/idl/user"
)

type UserSrv struct {
	user2.UnimplementedDouyinUserServiceServer
}

func (u *UserSrv) Login(context.Context, *user2.DouyinUserLoginRequest) (*user2.DouyinUserLoginResponse, error) {
	return &user2.DouyinUserLoginResponse{
		StatusCode: 200,
		StatusMsg:  "登录成功",
		UserId:     1,
		Token:      "test",
	}, nil
}

func (u *UserSrv) Register(context.Context, *user2.DouyinUserRegisterRequest) (*user2.DouyinUserRegisterResponse, error) {
	return &user2.DouyinUserRegisterResponse{
		StatusCode: 200,
		StatusMsg:  "注册成功",
		UserId:     1,
		Token:      "test",
	}, nil
}
