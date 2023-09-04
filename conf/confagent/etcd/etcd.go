package etcd

import (
	"context"
	"dzug/models"
	"encoding/json"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// etcd 相关操作

var (
	client *clientv3.Client
)

// Init 初始化  etcd
func Init(address []string, porjname string) (err error) {
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   address,
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v", err)
		return
	}
	//InitOnce 只需要在项目最开始时执行一次，目的是将项目配置存到etcd中，之后不再需要执行
	InitOnce(porjname)

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
func InitOnce(porjname string) {
	// put     在etcd里面设置key - value
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	//项目的所有配置的json格式的数据
	str := `[{"name":"Douyin","port":8001,"version":"v0.0.1","start_time":"2023-07-30","machine_id":1,"logconfig":{"path":"./logger/douyin.log","topic":"douyinlog","max_size":200,"max_backups":7,"max_age":30,"level":"debug","mode":"develop"},"jwt":{"jwt_expire":8760},"mysql":{"host":"127.0.0.1","port":3306,"user":"root","password":"root","database":"douyin","charset":"utf8mb4","parsetime":true,"loc":"Local"},"redis":{"host":"127.0.0.1","port":6379,"password":"","db":0,"pool_size":100,"min_idle_conns":5,"redis_expire":168},"kafka":{"address":"127.0.0.1:9092","chansize":10000},"etcd":{"address":"127.0.0.1:2379"},"service":{"user_service_name":"user","user_service_url":"127.0.0.1:9001","video_service_name":"video","video_service_url":"127.0.0.1:9002","favorite_service_name":"favorite","favorite_service_url":"127.0.0.1:9003","message_service_name":"message","message_service_url":"127.0.0.1:9007","comment_service_name":"comment","comment_service_url":"127.0.0.1:9005","relation_service_name":"relation","relation_service_url":"127.0.0.1:9004","publish_service_name":"publish","publish_service_url":"127.0.0.1:9008"},"video":{"feedcount":30},"ratelimit":{"rate":2,"cap":10},"esconfig":{"address":"127.0.0.1:9200","max_chan_size":100000,"goroutine_num":100},"mongodb":{"address":"127.0.0.1:27017"}}]`
	_, err := client.Put(ctx, porjname, str)
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v", err)
		return
	}
}
