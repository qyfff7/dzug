# 项目日志与项目配置

## 目录结构：

1. `conf`:存放项目配置相关的文件
   1. `config.go`:用于项目配置的初始化
   2. `config.yaml`:项目中所有的配置文件
2. `logger`:存放项目日志相关文件
   1. `logger.go`:用于项目中日志相关的操作
   2. `douyin.log`：项目日志文件
3. `routes`:路由文件（仅做测试，实际使用时，仿照这个文件的写法即可）
   1. `routes.go`：用于声明路由

4. `main.go`:主函数，用于启动项目

---



在配置文件中，基本都加上了注释，这里着重强调一下`mode`参数的含义：

`mode`参数的作用是控制日志输出的位置，在开发阶段输出到终端更加便于调试，在项目整体运行阶段，输出到日志更好，因此`mode = develop` 表示开发阶段，日志即输出到终端，也写入日志文件，

`mode = release` 表示项目发布阶段，日志只写入日志文件。

## 使用教程

1. 查看`config.yaml`文件，根据自己的需要，修改配置文件中的字段和值
2. 如果修改了配置文件中的某些字段（修改值不需要进行后面操作），则在`config.go`文件中，同样的修改对应的结构体，（添加新的字段就在对应的结构体里面添加新的属性）**注意：后面的`tag`一定是`mapstructure:`格式的**
3. 后续在项目的任何位置，想要调用获取某个配置的值时，只需要写`conf.Config.LogConfig` `conf.Config.Mode` 等等即可调用（这其实是一个嵌套的结构体，使用.号进行一层一层调用即可）
4. 参照`main.go`文件中的写法，在自己项目启动的文件的最前面，加上==**1. 初始化配置文件**==,**==2. 初始化日志==**的代码，
5. 后续在项目的任何地方，想要写日志的时候，只需要直接写`zap.L().Debug("XXX")` 或`zap.L().Info("XXX")`或`zap.L().Error("XXX")`等即可记录日志
6. 在注册路由的时候，将`conf.Config.Mode`参数传入，并且参照`routes.go`文件的撰写格式，使用`r := gin.New()`以及`r.Use(logger.GinLogger(), logger.GinRecovery(true))`
7. 启动项目`r.Run(fmt.Sprintf(":%d", conf.Config.Port))`



项目配置文件地址可在`config,go`文件中调整，日志文件可在`config.yaml`文件中调整。

补充：

在使用`zap.L().Error()` 记录日志的过程中，如果想要记录日志的内容，可以在后面加上`zap.Error(err)`(这是记录错误类型的)，`zap.String("lalala)` 这是记录String类型的，等等有多种类型，
如果不知道想记录的类型，写`zap.Any()`
示例：`zap.L().Error("UserRegister with invalid param, ", zap.Error(err))`、`zap.L().Info("这是一个日志记录示例", zap.String("参数错误", "用户注册时，输入参数出错"))`










