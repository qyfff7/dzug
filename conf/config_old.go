package conf

/*
import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

// Config 全局变量，用来保存项目所有的配置信息
var Config = new(ProjectConfig)

// ProjectConfig 项目所有的配置
// type logcollect []*LogConfig
type ProjectConfig struct {
	Name          string `mapstructure:"name" json:"name"`
	Port          int    `mapstructure:"port" json:"port"`
	Version       string `mapstructure:"version" json:"version"`
	StartTime     string `mapstructure:"start_time" json:"start_time"`
	Mode          string `mapstructure:"mode" json:"mode"`
	MachineID     int64  `mapstructure:"machine_id" json:"machine_id"`
	LogCollectKey string `mapstructure:"logcollectkey" json:"logcollectkey"`
	*MySQLConfig  `mapstructure:"mysql" json:"mysql"`
	*RedisConfig  `mapstructure:"redis" json:"redis"`
	//*EtcdConfig   `mapstructure:"etcd"`
	*KafkaConfig `mapstructure:"kafka" json:"kafka"`
	*JwtConfig   `mapstructure:"jwt" json:"jwt"`
	*Video       `mapstructure:"video" json:"video"`
	*Service     `mapstructure:"service" json:"service"`
	*Ratelimit   `mapstructure:"ratelimit" json:"ratelimit"`
	//*CollectEntry `mapstructure:"collectentry"`
}

// LogConfig 日志文件的配置
type LogConfig struct {
	Path       string `mapstructure:"path" json:"path"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size"`
	MaxAge     int    `mapstructure:"max_age" json:"max_age"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups"`
	Topic      string `mapstructure:"topic" json:"topic"`
	Level      string `mapstructure:"level" json:"level"`
}

// MySQLConfig 数据库配置
type MySQLConfig struct {
	Host      string `mapstructure:"host" json:"host"`
	Port      int    `mapstructure:"port" json:"port"`
	User      string `mapstructure:"user" json:"user"`
	Password  string `mapstructure:"password" json:"password"`
	DB        string `mapstructure:"database" json:"database"`
	Charset   string `mapstructure:"charset" json:"charset"`
	ParseTime bool   `mapstructure:"parsetime" json:"parsetime"`
	Loc       string `mapstructure:"loc" json:"loc"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Host         string `mapstructure:"host" json:"host"`
	Port         int    `mapstructure:"port" json:"port"`
	Password     string `mapstructure:"password" json:"password"`
	DB           int    `mapstructure:"db" json:"db"`
	PoolSize     int    `mapstructure:"pool_size" json:"pool_size"`
	MinIdleConns int    `mapstructure:"min_idle_conns" json:"min_idle_conns"`
	RedisExpire  int    `mapstructure:"redis_expire" json:"redis_expire"`
}

// EtcdConfig etcd配置
type EtcdConfig struct {
	Addr          []string `mapstructure:"address"`
	LogCollectKey string   `mapstructure:"logcollectkey"`
}

// KafkaConfig kafka配置
type KafkaConfig struct {
	Addr     string `mapstructure:"address" json:"address"`
	ChanSize int64  `mapstructure:"chansize" json:"chansize"`
}

// jwt 配置
type JwtConfig struct {
	JwtExpire int64 `mapstructure:"jwt_expire" json:"jwt_expire"`
}
type Video struct {
	FeedCount int64 `mapstructure:"feedcount" json:"feedcount"`
}

// Service 所有服务相关的配置（主要是服务名称和地址）
type Service struct {
	UserServiceName     string `mapstructure:"user_service_name" json:"user_service_name"`
	UserServiceUrl      string `mapstructure:"user_service_url" json:"user_service_url"`
	VideoServiceName    string `mapstructure:"video_service_name" json:"video_service_name"`
	VideoServiceUrl     string `mapstructure:"video_service_url" json:"video_service_url"`
	FavoriteServiceName string `mapstructure:"favorite_service_name" json:"favorite_service_name"`
	FavoriteServiceUrl  string `mapstructure:"favorite_service_url" json:"favorite_service_url"`
}

type Ratelimit struct {
	Rate int64 `mapstructure:"rate" json:"rate"`
	Cap  int64 `mapstructure:"cap" json:"cap"`
}

//// CollectEntry 要收集的日志的配置项结构体
//type CollectEntry struct {
//	Path       string `json:"path"`  //去哪个路径读取日志文件
//	Topic      string `json:"topic"` //日志文件发往kafka的哪个topic
//	MaxSize    int    `json:"max_size"`
//	MaxBackups int    `json:"max_backups"`
//	MaxAage    int    `json:"max_age"`
//	Level      string `json:"level"`
//}

// Init 从配置文件中获取项目所有的配置信息
func Init() (err error) {

	workDir, _ := os.Getwd()               // 获取当前文件夹路径
	viper.SetConfigName("config")          // 配置文件名
	viper.SetConfigType("yml")             // 配置文件格式
	viper.AddConfigPath(workDir + "/conf") // 添加配置路径

	if err = viper.ReadInConfig(); err != nil { // 查找并读取配置文件
		panic(fmt.Errorf("viper.ReadInConfig error config file: %s \n", err)) // 处理读取配置文件的错误
		return
	}
	//把读取到的配置信息，反序列化到Conf变量中
	if err := viper.Unmarshal(Config); err != nil {
		fmt.Printf("viper.Unmarshal failed ,err %v", err)
	}
	//监控配置文件的变化
	viper.WatchConfig()

	// 配置文件发生变更之后会调用的回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("项目配置文件修改了:", e.Name)
		if err = viper.Unmarshal(Config); err != nil {
			fmt.Printf("viper.Unmarshal failed ,err %v", err)
		}
	})
	return
}
*/
