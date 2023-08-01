package conf

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Config 全局变量，用来保存项目所有的配置信息
var Config = new(ProjectConfig)

// ProjectConfig 项目所有的配置
type ProjectConfig struct {
	Name      string `mapstructure:"name"`
	Port      int    `mapstructure:"port"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	Mode      string `mapstructure:"mode"`
	MachineID int64  `mapstructure:"machine_id"`

	*LogConfig   `mapstructure:"log"`
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

// LogConfig 日志文件的配置
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

// MySQLConfig 数据库配置
type MySQLConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DB        string `mapstructure:"database"`
	Charset   string `mapstructure:"charset"`
	ParseTime bool   `mapstructure:"parsetime"`
	Loc       string `mapstructure:"loc"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// Init 从配置文件中获取项目所有的配置信息
func Init() (err error) {

	config_filepath := "./user_service/Log_Conf/conf/config.yaml" // 指定配置文件路径
	viper.SetConfigFile(config_filepath)

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
