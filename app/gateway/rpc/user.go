package rpc

import (
	"context"
	pb "dzug/protos/user"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

// Connection 与服务端建立连接
func Connection() {

}

func Login(ctx context.Context, in *pb.LoginAndRegisterRequest) (*pb.LoginAndRegisterResponse, error) {

	//1.连接到server 端，WithTransportCredentials 加入安全校验  insecure.NewCredentials() 表示不使用任何加密
	//conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// grpc.Dial 创建到给定目标的客户端连接。

	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.L().Error("与用户服务建立连接失败" + err.Error())
		return nil, err
	}
	defer conn.Close()

	//2. 建立链接
	client := pb.NewServiceClient(conn)
	defer conn.Close()
	//3.执行RPC的调用 （这个方法在服务端来实现并返回结果）
	//in.Username =
	resp, err := client.Login(context.Background(), &pb.LoginAndRegisterRequest{})
	if err != nil {

		panic(err)
	}
	fmt.Println("客户端调用服务端的登录服务，成功，得到的Userid是", resp.GetUserId(), "tocken 是", resp.GetToken())
	return &pb.LoginAndRegisterResponse{UserId: 916, Token: "这里是clinet/rpc/user,正在调用User服务的Login方法，调用成功！" + "输入的用户名是：" + in.Username + "密码是：" + in.Password}, nil

}
func Register(ctx context.Context, in *pb.LoginAndRegisterRequest) (*pb.LoginAndRegisterResponse, error) {

	conn, err := grpc.Dial("127.0.0.1:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.L().Error("与用户服务建立连接失败" + err.Error())
		return nil, err
	}
	defer conn.Close()

	//2. 建立链接
	client := pb.NewServiceClient(conn)
	defer conn.Close()
	//3.执行RPC的调用 （这个方法在服务端来实现并返回结果）
	//in.Username =
	resp, err := client.Register(context.Background(), &pb.LoginAndRegisterRequest{})
	if err != nil {

		panic(err)
	}
	fmt.Println("客户端调用服务端的登录服务，成功，得到的Userid是", resp.GetUserId(), "tocken 是", resp.GetToken())
	return &pb.LoginAndRegisterResponse{UserId: 916, Token: "这里是clinet/rpc/user,正在调用User服务的Login方法，调用成功！" + "输入的用户名是：" + in.Username + "密码是：" + in.Password}, nil
}
