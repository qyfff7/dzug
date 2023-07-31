package rpc

import (
	"context"
	"dzug/user_service/discovery"
	"dzug/user_service/idl/user"
)

func UserLogin(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	discovery.LoadClient("user", &discovery.UserClient) // 加载etcd客户端程序
	r, err := discovery.UserClient.Login(ctx, req)      // 调用登录方法
	if err != nil {
		return
	}
	return r, nil
}

func UserRegister(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	discovery.LoadClient("user", &discovery.UserClient)
	r, err := discovery.UserClient.Register(ctx, req) // 调用注册方法
	if err != nil {
		return
	}
	return r, nil
}
