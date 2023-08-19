package etcd

import (
	"context"
	"dzug/models"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
)

// etcd 相关操作

var (
	client *clientv3.Client
)

// Init 初始化  etcd
func Init(address []string) (err error) {
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   address,
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v", err)
		return
	}
	return
}

// GetProjectConf 拉取日志收集配置项的函数
func GetProjectConf(key string) (config *models.ProjectConfig, err error) {
	var configlist []*models.ProjectConfig

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	resp, err := client.Get(ctx, key)
	if err != nil {
		//zap.L().Error("get conf from etcd by key:" + fmt.Sprintf("%s", key) + " failed ,err:%v" + fmt.Sprintf("%s", err))
		fmt.Printf("get conf from etcd by key:%s ,err:%v", key, err)
		return
	}
	if len(resp.Kvs) == 0 {
		//zap.L().Warn("get len:0 conf from etcd by key:%s" + fmt.Sprintf("%s", key))
		fmt.Printf("get len:0 conf from etcd by key:%s", key)
		return
	}

	ret := resp.Kvs[0] //取一个
	// ret.Value // json格式字符串
	//fmt.Printf("%s", ret.Value)

	//将从etcd中去取出来的值ret.Value利用Unmarshal方法反序列化出来，存放在collectEntryList上
	err = json.Unmarshal(ret.Value, &configlist)
	if err != nil {
		//zap.L().Error("json unmarshal failed, err:", zap.Error(err))
		fmt.Printf("json unmarshal failed, err:%v", err)
		return nil, err
	}
	config = configlist[0]
	return
}

// GetLogConf 拉取日志收集配置项的函数
func GetLogConf(key string) (logconflist []*models.LogConfig, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	resp, err := client.Get(ctx, key)
	if err != nil {
		//zap.L().Error("get conf from etcd by key:" + fmt.Sprintf("%s", key) + " failed ,err:%v" + fmt.Sprintf("%s", err))
		fmt.Printf("get conf from etcd by key:%s ,err:%v", key, err)
		return
	}
	if len(resp.Kvs) == 0 {
		//zap.L().Warn("get len:0 conf from etcd by key:%s" + fmt.Sprintf("%s", key))
		fmt.Printf("get len:0 conf from etcd by key:%s", key)
		return
	}

	ret := resp.Kvs[0] //取一个
	//// ret.Value // json格式字符串
	//fmt.Printf("%s", ret.Value)

	//将从etcd中去取出来的值ret.Value利用Unmarshal方法反序列化出来，存放在collectEntryList上
	err = json.Unmarshal(ret.Value, &logconflist)
	if err != nil {
		//zap.L().Error("json unmarshal failed, err:", zap.Error(err))
		fmt.Printf("json unmarshal failed, err:%v", err)
		return nil, err
	}
	return
}
