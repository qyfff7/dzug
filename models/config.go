package models

// ProjectConfig 项目所有的配置
type ProjectConfig struct {
	Name         string `json:"name"`
	Port         int    `json:"port"`
	Version      string `json:"version"`
	StartTime    string `json:"start_time"`
	MachineID    int64  `json:"machine_id"`
	*LogConfig   `json:"logconfig"`
	*MySQLConfig `json:"mysql"`
	*RedisConfig `json:"redis"`
	*KafkaConfig `json:"kafka"`
	*JwtConfig   `json:"jwt"`
	*VideoConfig `json:"video"`
	*Service     `json:"service"`
	*Ratelimit   `json:"ratelimit"`
	*EtcdConfig  `json:"etcd"`
	*ESConf      `json:"esconfig"`
}

// LogConfig 日志文件的配置
type LogConfig struct {
	Path       string `json:"path"`
	MaxSize    int    `json:"max_size"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
	Topic      string `json:"topic"`
	Level      string `json:"level"`
	Mode       string `json:"mode"`
}

// MySQLConfig 数据库配置
type MySQLConfig struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	User      string `json:"user"`
	Password  string `json:"password"`
	DB        string `json:"database"`
	Charset   string `json:"charset"`
	ParseTime bool   `json:"parsetime"`
	Loc       string `json:"loc"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Password     string `json:"password"`
	DB           int    `json:"db"`
	PoolSize     int    `json:"pool_size"`
	MinIdleConns int    `json:"min_idle_conns"`
	RedisExpire  int    `json:"redis_expire"`
}

// EtcdConfig etcd配置
type EtcdConfig struct {
	Addr string `json:"address"`
}

// KafkaConfig kafka配置
type KafkaConfig struct {
	Addr     string `json:"address"`
	ChanSize int64  `json:"chansize"`
}

// jwt 配置
type JwtConfig struct {
	JwtExpire int64 `json:"jwt_expire"`
}
type VideoConfig struct {
	FeedCount int64 `json:"feedcount"`
}

// Service 所有服务相关的配置（主要是服务名称和地址）
type Service struct {
	UserServiceName     string `json:"user_service_name"`
	UserServiceUrl      string `json:"user_service_url"`
	VideoServiceName    string `json:"video_service_name"`
	VideoServiceUrl     string `json:"video_service_url"`
	FavoriteServiceName string `json:"favorite_service_name"`
	FavoriteServiceUrl  string `json:"favorite_service_url"`
}

type Ratelimit struct {
	Rate int64 `json:"rate"`
	Cap  int64 `json:"cap"`
}
type ESConf struct {
	Address string `json:"address"`
	Index   string `json:"index"`
	MaxSize int    `json:"max_chan_size"`
	GoNum   int    `json:"goroutine_num"`
}
type BasicConfig struct {
	Name     string   `mapstructure:"project_name"`
	EtcdAddr []string `mapstructure:"etcd_address"`
}
