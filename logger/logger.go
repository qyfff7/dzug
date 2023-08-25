package logger

import (
	"context"
	"dzug/conf"
	"dzug/logger/logagent"
	"dzug/logger/logagent/kafka"
	"dzug/logger/logagent/tailfile"
	"dzug/models"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

var LogClient *clientv3.Client
var LogBaseConf = new(models.BasicConfig)

//var LogConf = new(models.LogConfig)

// Init 初始化Logger
func Init() (err error) {

	//1.初始化viper
	ymlPath := "/logger/conf/config.yml"
	if err = conf.ViperInit(LogBaseConf, ymlPath); err != nil {
		fmt.Printf("viper 初始化失败..., err:%v\n", err)
	}
	//2.连接etcd
	LogClient, err = clientv3.New(clientv3.Config{
		Endpoints:   LogBaseConf.EtcdAddr,
		DialTimeout: time.Second * 5,
	})
	if err != nil {
		fmt.Printf("connect to etcd failed, err:%v", err)
		return
	}
	//3. 判断配置是否已经存到etcd
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := LogClient.Get(ctx, LogBaseConf.Name)
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	//如果配置没有存到etcd
	if len(resp.Kvs) == 0 {
		//从yml文件中读取配置，存到etcd中
		if err = conf.ViperInit(conf.LogConf, ymlPath); err != nil {
			fmt.Printf("viper 失败..., err:%v\n", err)
		}
		err = conf.PutConfigToEtcd(LogBaseConf.Name, conf.LogConf)
		if err != nil {
			fmt.Println("log配置存到etcd过程中出错：" + err.Error())
			return err
		}
	} else {
		err = json.Unmarshal(resp.Kvs[0].Value, &conf.LogConf)
	}
	//4.启动配置监控
	go WatchLogConf(LogBaseConf.Name)

	//5.初始化日志
	if err = zapInit(); err != nil {
		fmt.Println("zap init failed ... ")
		return
	}
	//6.初始化kafka和es，用于日志收集
	if err = logagent.LogAgentInit(); err != nil {
		zap.L().Error("初始化kafka和es 失败,err:", zap.Error(err))
		return
	}

	//time.Sleep(time.Second * 3)
	go CollectLog()

	return

}

func zapInit() (err error) {
	logconf := conf.LogConf
	writeSyncer := getLogWriter(logconf.Path, logconf.MaxSize, logconf.MaxBackups, logconf.MaxAge)
	encoder := getEncoder()
	l := new(zapcore.Level)
	err = l.UnmarshalText([]byte(logconf.Level))
	if err != nil {
		return
	}
	//fmt.Println("在 zap 的init 中")
	var core zapcore.Core
	//fmt.Println(logconf.Mode)
	if fmt.Sprintf("%s", logconf.Mode) == "develop" {
		//开发模式，日志输出到终端
		consoleEnbcoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			//同时输出到日志文件 和 终端
			zapcore.NewCore(encoder, writeSyncer, l),
			//输出到终端
			zapcore.NewCore(consoleEnbcoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		)
	} else {
		//只将日志写入到日志文件中
		core = zapcore.NewCore(encoder, writeSyncer, l)
	}

	lg := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(lg) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可

	return
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()
		cost := time.Since(start)
		zap.L().Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					zap.L().Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					zap.L().Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

// CollectLog  收集日志
func CollectLog() (err error) {
	err = collectrun(conf.LogConf.Topic)
	if err != nil {
		zap.L().Error("Error sending log data to kafka : ", zap.Error(err))
		return
	}
	return
}

// collectrun 真正的业务逻辑
func collectrun(topic string) (err error) {
	// logfile --> TailObj --> log --> Client --> kafka
	//利用ini文件，创建kafka配置项，日志文件配置项  --> 读取出ini文件里面的信息，用来初始化 kafka 和  tail
	//--> tail得到日志文件的地址  --> TailObj对象读取出一行 log  --> 包装成一个发送到kafka所需要的 msg 对象,发送到一个channel 中
	//-->  在kafka初始化的时候，就创建一个goroutine，来从channel中读取信息， 真正发送到kafka中
	for {
		// 循环读数据
		line, ok := <-tailfile.TailObj.Lines // chan tail.Line
		if !ok {
			zap.L().Warn("tail file close reopen, filename: " + fmt.Sprintf("%s", tailfile.TailObj.Filename))
			time.Sleep(time.Second) // 读取出错等一秒
			continue
		}
		// 如果是空行就略过
		//fmt.Printf("%#v\n", line.Text)
		if len(strings.Trim(line.Text, "\r")) == 0 { //strings.Trim  用来去除  "\r"
			zap.L().Info("出现空行拉,直接跳过...")
			continue
		}

		//如果不适用channel的话，就是同步的操作，也就是读取一行日志，发送一行，这样当日志比较多的时候，是比较耗时的；
		//使用channel,可以改成异步的操作，也就是一个goroutine一直在从日志文件里面读取日志，然后发送到一个channel里面，
		//另一个 goroutine一直从该 channel 里面取日志信息，并且发送到kafka中。

		// 利用通道将同步的代码改为异步的
		// 把读出来的一行日志包装成kafka里面的msg类型
		msg := &sarama.ProducerMessage{}
		msg.Topic = topic
		msg.Value = sarama.StringEncoder(line.Text)
		// 丢到通道中
		kafka.ToMsgChan(msg)
	}
}

// WatchLogConf 监控etcd中log服务配置变化
func WatchLogConf(key string) {
	for {
		watchCh := LogClient.Watch(context.Background(), key)
		for wresp := range watchCh {
			fmt.Println("get new conf from etcd!!!")
			for _, evt := range wresp.Events {
				fmt.Printf("type:%s key:%s value:%s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
				err := json.Unmarshal(evt.Kv.Value, &conf.LogConf)
				if err != nil {
					fmt.Println("json unmarshal new conf failed, err: ", err)
					continue
				}
				if err = zapInit(); err != nil {
					fmt.Println("zap init failed ... ")
					return
				}

				if err = logagent.LogAgentInit(); err != nil {
					zap.L().Error("初始化kafka和es 失败,err:", zap.Error(err))
					return
				}

			}

		}
	}

}
