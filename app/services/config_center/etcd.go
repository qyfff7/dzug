package config_center

import (
	"context"
	"dzug/conf"
	"dzug/models"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

// 配置中心
var (
	ProiClient *clientv3.Client
)
var ProjConf = new(models.ProjectConfig)
var ProjBaseConf = new(models.BasicConfig)

// Init 初始化  etcd
func Init() (err error) {

	ymlPath := "/app/services/config_center/conf"
	//1.初始化viper
	if err := conf.ViperInit(ProjBaseConf, ymlPath); err != nil {
		fmt.Printf("viper 初始化失败...,baseconf   err:%v\n", err)
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
		if err := conf.ViperInit(ProjConf, "/conf"); err != nil {
			fmt.Printf("viper 初始化失败..., projconf   err:%v\n", err)
		}
		err = PutConfigToEtcd(ProjBaseConf.Name, ProjConf)
		if err != nil {
			fmt.Println("项目配置存到etcd过程中出错：" + err.Error())
			return err
		}
	} else { //如果已经存到etcd上
		var confs []*models.ProjectConfig
		ret := resp.Kvs[0]
		err = json.Unmarshal(ret.Value, &confs)
		ProjConf = confs[0]
	}
	//启动配置监控
	WatchProjConf(ProjBaseConf.Name)
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
			fmt.Println("get new conf from etcd!!!")
			for _, evt := range wresp.Events {
				fmt.Printf("type:%s key:%s value:%s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
				var newConf []models.ProjectConfig
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					fmt.Println("json unmarshal new conf failed, err: ", err)
					continue
				}
				ProjConf = &newConf[0]
			}
		}
	}
}

// WatchUserConf 监控etcd中user服务配置变化
func WatchUserConf(key string) {
	for {
		watchCh := ProiClient.Watch(context.Background(), key)
		for wresp := range watchCh {
			fmt.Println("get new conf from etcd!!!")
			for _, evt := range wresp.Events {
				fmt.Printf("type:%s key:%s value:%s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
				var newConf []models.ProjectConfig
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					fmt.Println("json unmarshal new conf failed, err: ", err)
					continue
				}
				ProjConf = &newConf[0]
			}
		}
	}
}
