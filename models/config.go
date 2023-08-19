package models

// ProjectConfig 项目所有的配置
type ProjectConfig struct {
	Name      string `mapstructure:"name" json:"name"`
	Port      int    `mapstructure:"port" json:"port"`
	Version   string `mapstructure:"version" json:"version"`
	StartTime string `mapstructure:"start_time" json:"start_time"`
	//Mode         string `mapstructure:"mode" json:"mode"`
	MachineID    int64 `mapstructure:"machine_id" json:"machine_id"`
	*LogConfig   `mapstructure:"logconfig" json:"logconfig"`
	*MySQLConfig `mapstructure:"mysql" json:"mysql"`
	*RedisConfig `mapstructure:"redis" json:"redis"`
	*KafkaConfig `mapstructure:"kafka" json:"kafka"`
	*JwtConfig   `mapstructure:"jwt" json:"jwt"`
	*VideoConfig `mapstructure:"video" json:"video"`
	*Service     `mapstructure:"service" json:"service"`
	*Ratelimit   `mapstructure:"ratelimit" json:"ratelimit"`
	*EtcdConfig  `mapstructure:"etcdconfig" json:"etcd"`
	*ESConf      `json:"esconfig"`
}

// LogConfig 日志文件的配置
type LogConfig struct {
	Path       string `mapstructure:"path" json:"path"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size"`
	MaxAge     int    `mapstructure:"max_age" json:"max_age"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups"`
	Topic      string `mapstructure:"topic" json:"topic"`
	Level      string `mapstructure:"level" json:"level"`
	Mode       string `mapstructure:"mode" json:"mode"`
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
	Addr string `mapstructure:"address" json:"address"`
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
type VideoConfig struct {
	FeedCount int64 `mapstructure:"feedcount" json:"feedcount"`
}

// Service 所有服务相关的配置（主要是服务名称和地址）
type Service struct {
	UserServiceName     string `mapstructure:"user_service_name" json:"user_service_name"`
	UserServiceUrl      string `mapstructure:"user_service_url" json:"user_service_url"`
	VideoServiceName    string `mapstructure:"video_service_name" json:"video_service_name"`
	VideoServiceUrl     string `mapstructure:"video_service_url" json:"video_service_url"`
	FavoriteServiceName string `mapstructure:"favorite_service_name" json:"favorite_service_name"`
	FavoriteServiceUrl  string `json:"favorite_service_url"`
}

type Ratelimit struct {
	Rate int64 `mapstructure:"rate" json:"rate"`
	Cap  int64 `mapstructure:"cap" json:"cap"`
}
type ESConf struct {
	Address string `json:"address"`
	Index   string `json:"index"`
	MaxSize int    `json:"max_chan_size"`
	GoNum   int    `json:"goroutine_num"`
}
