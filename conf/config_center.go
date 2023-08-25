package conf

import (
	"context"
	"dzug/app/services/user/pkg/snowflake"
	"dzug/models"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"os"
	"time"
)

// 配置中心
var (
	ProiClient   *clientv3.Client
	UserConf     = new(models.UserConfig)
	ProjConf     = new(models.ProjectConfig)
	LogConf      = new(models.LogConfig)
	ProjBaseConf = new(models.BasicConfig)
)

// Init 初始化  etcd
func Init() (err error) {

	ymlPath := "/conf/config.yml"
	//1.初始化viper
	if err = ViperInit(ProjBaseConf, ymlPath); err != nil {
		fmt.Printf("viper 初始化失败...,err:%v\n", err)
	}
	//2.连接etcd
	ProiClient, err = clientv3.New(clientv3.Config{
		Endpoints:   ProjBaseConf.EtcdAddr,
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v", err)
		return
	}
	//3. 判断项目配置是否已经存到etcd
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := ProiClient.Get(ctx, ProjBaseConf.Name)
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	//如果项目配置没有存到etcd
	if len(resp.Kvs) == 0 {
		//从yml文件中读取配置，存到etcd中
		if err = ViperInit(ProjConf, ymlPath); err != nil {
			fmt.Printf("viper 初始化失败..., err:%v\n", err)
		}
		err = PutConfigToEtcd(ProjBaseConf.Name, ProjConf)
		if err != nil {
			fmt.Println("项目配置存到etcd过程中出错：" + err.Error())
			return err
		}
	} else { //如果已经存到etcd上

		err = json.Unmarshal(resp.Kvs[0].Value, &ProjConf)

	}
	//启动配置监控
	go WatchProjConf(ProjBaseConf.Name)

	// snowflake初始化
	if err = snowflake.Init(ProjConf.StartTime, ProjConf.MachineID); err != nil {
		zap.L().Error("snowflake initialization error", zap.Error(err))
		return
	}
	return
}
func PutConfigToEtcd(key string, projconf interface{}) (err error) {
	// put     在etcd里面设置key - value
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	//项目的所有配置的json格式的数据
	str, err := json.Marshal(projconf)
	_, err = ProiClient.Put(ctx, key, string(str))
	if err != nil {
		fmt.Printf("put config to etcd failed, err:%v", err)
		return
	}
	return err
}

// WatchProjConf 监控etcd中项目配置变化
func WatchProjConf(key string) {
	for {
		watchCh := ProiClient.Watch(context.Background(), key)
		for wresp := range watchCh {
			fmt.Println("get new project conf from etcd!!!")
			for _, evt := range wresp.Events {
				fmt.Printf("type:%s key:%s value:%s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)

				err := json.Unmarshal(evt.Kv.Value, &ProjConf)
				if err != nil {
					fmt.Println("json unmarshal new conf failed, err: ", err)
					continue
				}
			}
		}
	}
}

func ViperInit(config interface{}, YmlFilePath string) (err error) {
	workDir, _ := os.Getwd()
	fmt.Println(YmlFilePath)
	viper.SetConfigFile(workDir + YmlFilePath)
	if err = viper.ReadInConfig(); err != nil { // 查找并读取配置文件
		panic(fmt.Errorf("viper.ReadInConfig error config file: %s \n", err)) // 处理读取配置文件的错误
		return
	}
	//把读取到的配置信息，反序列化到Conf变量中
	if err = viper.Unmarshal(config); err != nil {
		fmt.Printf("viper.Unmarshal failed ,err %v", err)
	}
	return
}
