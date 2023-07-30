package rpc

import (
	"context"
	"dzug/idl/user"
)

func UserLogin(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	loadClient("user", &UserClient)
	r, err := UserClient.Login(ctx, req)
	if err != nil {
		return
	}
	return r, nil
}

func UserRegister(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	loadClient("user", &UserClient)
	r, err := UserClient.Register(ctx, req)
	if err != nil {
		return
	}
	return r, nil
}
