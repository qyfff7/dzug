package models

// ProjectConfig 项目所有的配置
type ProjectConfig struct {
	Name      string `mapstructure:"name" json:"name"`
	Port      int    `mapstructure:"port" json:"port"`
	Version   string `mapstructure:"version" json:"version"`
	StartTime string `mapstructure:"start_time" json:"start_time"`
	MachineID int64  `mapstructure:"machine_id" json:"machine_id"`
}

// LogConfig 日志文件的配置
type LogConfig struct {
	Path         string `mapstructure:"path" json:"path"`
	MaxSize      int    `mapstructure:"max_size" json:"max_size"`
	MaxAge       int    `mapstructure:"max_age" json:"max_age"`
	MaxBackups   int    `mapstructure:"max_backups" json:"max_backups"`
	Topic        string `mapstructure:"topic" json:"topic"`
	Level        string `mapstructure:"level" json:"level"`
	Mode         string `mapstructure:"mode" json:"mode"`
	*ESConfig    `mapstructure:"es" json:"es"`
	*KafkaConfig `mapstructure:"kafka" json:"kafka"`
}

// ProjectConfig 项目所有的配置
type UserConfig struct {
	Url          string `mapstructure:"url" json:"url"`
	Name         string `mapstructure:"service_name" json:"service_name"`
	*JwtConfig   `mapstructure:"jwt" json:"jwt"`
	*MySQLConfig `mapstructure:"mysql" json:"mysql"`
	*RedisConfig `mapstructure:"redis" json:"redis"`
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
type MongoDbConfig struct {
	Addr string `json:"address"`
}

// Service 所有服务相关的配置（主要是服务名称和地址）
type Service struct {
	UserServiceName     string `json:"user_service_name"`
	UserServiceUrl      string `json:"user_service_url"`
	VideoServiceName    string `json:"video_service_name"`
	VideoServiceUrl     string `json:"video_service_url"`
	FavoriteServiceName string `json:"favorite_service_name"`
	FavoriteServiceUrl  string `json:"favorite_service_url"`
	MessageServiceName  string `json:"message_service_name"`
	MessageServiceUrl   string `json:"message_service_url"`
	CommentServiceUrl   string `json:"comment_service_url"`
	CommentServiceName  string `json:"comment_service_name"`
	RelationServiceName string `json:"relation_service_name"`
	RelationServiceUrl  string `json:"relation_service_url"`
	PublishServiceName  string `json:"publish_service_name"`
	PublishServiceUrl   string `json:"publish_service_url"`
}

type Ratelimit struct {
	Rate int64 `json:"rate"`
	Cap  int64 `json:"cap"`
}

type BasicConfig struct {
	Name     string   `mapstructure:"name"`
	EtcdAddr []string `mapstructure:"etcd_address"`
}
type ESConfig struct {
	Address string `json:"address"`
	Index   string `json:"index"`
	MaxSize int    `json:"max_chan_size"`
	GoNum   int    `json:"goroutine_num"`
}
