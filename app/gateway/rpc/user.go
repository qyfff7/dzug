package rpc

import (
	"context"
	"dzug/discovery"
	pb "dzug/protos/user"
)

// 这里应该做的事情是与用户服务建立连接，调用远程的方法

/*func UserLogin(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
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
}*/

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
