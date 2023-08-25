package userservice

import (
	"context"
	"dzug/app/services/user/dal/redis"
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

//var UserConf = new(models.UserConfig)

func Start() (err error) {

	//1.初始化viper
	ymlPath := "/app/services/user/conf/config.yml"
	if err = conf.ViperInit(UserBaseConf, ymlPath); err != nil {
		fmt.Printf("viper 初始化失败..., err:%v\n", err)
	}
	//2.连接etcd
	UserClient, err = clientv3.New(clientv3.Config{
		Endpoints:   UserBaseConf.EtcdAddr,
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v", err)
		return
	}
	//3. 判断user配置是否已经存到etcd
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	resp, err := UserClient.Get(ctx, UserBaseConf.Name)
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	//如果user配置没有存到etcd
	if len(resp.Kvs) == 0 {
		//从yml文件中读取配置，存到etcd中
		if err := conf.ViperInit(conf.UserConf, ymlPath); err != nil {
			fmt.Printf("viper 初始化失败..., err:%v\n", err)
		}
		err = conf.PutConfigToEtcd(UserBaseConf.Name, conf.UserConf)
		if err != nil {
			fmt.Println("user配置存到etcd过程中出错：" + err.Error())
			return
		}
	} else { //如果已经存到etcd上
		err = json.Unmarshal(resp.Kvs[0].Value, &conf.UserConf)
	}
	//4.启动配置监控
	go WatchUserConf(UserBaseConf.Name)

	//5.初始化数据库
	if err = repo.Init(conf.UserConf.MySQLConfig.User,
		conf.UserConf.MySQLConfig.Password,
		conf.UserConf.MySQLConfig.Host,
		conf.UserConf.MySQLConfig.DB,
		conf.UserConf.MySQLConfig.Charset,
		conf.UserConf.MySQLConfig.Loc,
		conf.UserConf.MySQLConfig.Port,
		conf.UserConf.MySQLConfig.ParseTime); err != nil {
		return
	}
	//6.初始化user 的 redis
	if err = redis.Init(conf.UserConf.RedisConfig.Host,
		conf.UserConf.RedisConfig.Password,
		conf.UserConf.RedisConfig.DB,
		conf.UserConf.RedisConfig.Port,
		conf.UserConf.RedisConfig.PoolSize,
		conf.UserConf.RedisConfig.MinIdleConns); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	// 程序退出关闭数据库连接
	defer redis.Close()

	// 传入注册的服务名和注册的服务地址进行注册
	serviceRegister, grpcServer := discovery.InitRegister(conf.UserConf.Name, conf.UserConf.Url)
	defer serviceRegister.Close()
	defer grpcServer.Stop()
	pb.RegisterServiceServer(grpcServer, &service.Userservice{}) // 绑定grpc
	discovery.GrpcListen(grpcServer, conf.UserConf.Url)          // 开启监听
	return err
}

// WatchUserConf 监控etcd中user服务配置变化
func WatchUserConf(key string) {
	for {
		watchCh := UserClient.Watch(context.Background(), key)
		for wresp := range watchCh {
			fmt.Println("get new conf from etcd!!!")
			for _, evt := range wresp.Events {
				fmt.Printf("type:%s key:%s value:%s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
				err := json.Unmarshal(evt.Kv.Value, &conf.UserConf)
				if err != nil {
					fmt.Println("json unmarshal new conf failed, err: ", err)
					continue
				}
				if err = repo.Init(conf.UserConf.MySQLConfig.User,
					conf.UserConf.MySQLConfig.Password,
					conf.UserConf.MySQLConfig.Host,
					conf.UserConf.MySQLConfig.DB,
					conf.UserConf.MySQLConfig.Charset,
					conf.UserConf.MySQLConfig.Loc,
					conf.UserConf.MySQLConfig.Port,
					conf.UserConf.MySQLConfig.ParseTime); err != nil {
					return
				}
				//6.初始化user 的 redis
				if err = redis.Init(conf.UserConf.RedisConfig.Host,
					conf.UserConf.RedisConfig.Password,
					conf.UserConf.RedisConfig.DB,
					conf.UserConf.RedisConfig.Port,
					conf.UserConf.RedisConfig.PoolSize,
					conf.UserConf.RedisConfig.MinIdleConns); err != nil {
					fmt.Printf("init redis failed, err:%v\n", err)
					return
				}

			}
		}
	}
}
