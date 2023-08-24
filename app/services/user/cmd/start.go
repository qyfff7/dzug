package userservice

import (
	"context"
	"dzug/app/services/config_center"
	"dzug/app/services/user/service"
	"dzug/conf"
	"dzug/discovery"
	"dzug/models"
	pb "dzug/protos/user"
	"dzug/repo"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var UserClient *clientv3.Client
var UserBaseConf = new(models.BasicConfig)
var UserConf = new(models.UserConfig)

func Start() {

	//1.初始化viper
	ymlPath := "/app/services/user/conf"
	if err := conf.ViperInit(UserBaseConf, ymlPath); err != nil {
		fmt.Printf("viper 初始化失败...,baseconf   err:%v\n", err)
	}
	//2.连接etcd
	UserClient, err := clientv3.New(clientv3.Config{
		Endpoints:   UserBaseConf.EtcdAddr,
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v", err)
		return
	}
	//3. 判断user配置是否已经存到etcd
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := UserClient.Get(ctx, UserBaseConf.Name)
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	//如果user配置没有存到etcd
	if len(resp.Kvs) == 0 {
		//从yml文件中读取配置，存到etcd中
		if err := conf.ViperInit(UserConf, ymlPath); err != nil {
			fmt.Printf("viper 初始化失败..., projconf   err:%v\n", err)
		}
		err = config_center.PutConfigToEtcd(UserBaseConf.Name, UserConf)
		if err != nil {
			fmt.Println("项目配置存到etcd过程中出错：" + err.Error())
			return
		}
	} else { //如果已经存到etcd上
		var userconfs []*models.UserConfig
		ret := resp.Kvs[0]
		err = json.Unmarshal(ret.Value, &userconfs)
		UserConf = userconfs[0]
	}
	//4.启动配置监控
	config_center.WatchProjConf(UserBaseConf.Name)

	//5.初始化数据库
	repo.Init(UserConf.MySQLConfig)
	//6.

	// 传入注册的服务名和注册的服务地址进行注册
	serviceRegister, grpcServer := discovery.InitRegister(UserConf.Name, UserConf.Url)
	defer serviceRegister.Close()
	defer grpcServer.Stop()
	pb.RegisterServiceServer(grpcServer, &service.Userservice{}) // 绑定grpc
	discovery.GrpcListen(grpcServer, UserConf.Url)               // 开启监听
}
