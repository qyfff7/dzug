# dzug
### 抖声 demo

此demo使用框架为 grpc+etcd+gin，使用etcd来进行服务注册发现，主要使用两个服务，来测试服务注册与发现功能，其他功能基本没有实现



### 用来演示的接口为

1. [/douyin/user/login](https://apifox.com/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707522)
2. [/douyin/user/register](https://apifox.com/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707521)
3. [/douyin/relation/action](https://apifox.com/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707530)

前两者用于登录注册，最后者用于关注和取关
拆分两个微服务，12为用户服务，3为关系服务



### demo完成进度：

1. 可使用etcd完成服务注册与发现
2. 无数据库，服务端直接返回测试数据
3. 两项服务均可进行测试
4. 基本实现微服务的注册发现功能
5. jwt验证等中间件等也完全不在demo考虑范围内
6. 服务于服务之间的调用尚未实现



### 文件架构说明：

1. app：网关以及各个服务
    1. gateway：网关
        1. cmd包：启动模块
        2. handlers包：handler方法
        3. routes：路由
        4. rpc包：远程调用方法
    2. relation包：关系服务
        1. cmd包：启动模块
        2. service包：业务处理，处理远程的调用，这里直接返回数据，之后可以接上数据库等进行完善
    3. user包：用户服务
        1. 同relation
2. discovery包：服务注册发现模块
    1. discovery.go：服务发现
    2. initDiscovery.go：初始化一个服务发现程序
    3. register.go：服务注册
3. idl包：grpc远程调用生成文件



### 配置要求：

需要安装etcd并启动服务

1. etcd端口：使用默认2379端口
2. user服务：使用9000端口
3. relation服务：使用9001端口



### 启动：

在三个cmd文件下的main.go文件直接进行启动即可，postman输入对应的连接可以进行调用。

**注意**，protoc生成的pb.go文件给json添加了omitempty参数，不按照正确json格式在postman里进行调用也能正常运行，==测试只要链接正确就可以返回数据==。



![aciton测试](doc/image-20230727174553448.png)

---

以上部分已更新，以下为demo中的服务调用流程：

## demo中微服务调用流程

### gateway

gateway用来处理网关相关操作

1. main.go：启动路由程序
2. router.go：路由注册，绑定handlers中的handler
3. discovery.Init：启动服务发现程序
4. handler.go：被调用后会调用rpc.go进行远程调用
5. rpc.go：每当rpc程序被调用一次，都会从etcd中查找新的grpc链接，然后进行rpc远程调用

### user

1. main.go：启动etcd服务与grpc服务，然后进行监听
2. service.go：等待被远程调用
3. …：可以写更多的方法来进行调用，比如写dao包来被service操作，处理数据库相关的内容
4. …：可以调用discovery包下的initDiscovery.go 的Init()方法，调用其他服务，如relation的rpc方法（未测试）

### Relation

基本同user

## 开发新服务的建议流程（以comment服务为例）

1. 在idl包新建comment.proto，编写proto文件
2. 通过protoc命令生成comment.pb.go与comment_grpc.pb.go文件，并放置在idl下的commetn文件夹下（目前没有就新建）
3. 在app包建comment文件夹
4. cmd包放置main.go（仅作参考）
5. service放置rpc函数，需要继承comment_grpc.pb.go文件夹下的DouyinCommentServiceServer 接口（这个名字写法，只是参照我的写法，可能叫别的名字）
   ![image-20230730202104030](doc/image-20230730202104030.png)
   ![image-20230730202133306](doc/image-20230730202133306.png)如，这里的的service包下的UserSrv结构体就继承了对应的包
6. main函数写法，参照任意一个即可，etcd端口号与你的etcd端口相同，一般本地的不改动，都是localhost:2379，ServiceRegister这里的Value用一个新的端口，不然无法注册服务，net.Listen里的端口号与value保持一致，其他几乎不用改动，当然，除了把service.UserSrv{}注册的那个位置，要用正确的

---

待完善：

1. 日志系统
2. 错误处理模式
3. 配置管理viper
4. etcd的负载均衡

to be done...

测试本地仓库能否提交到远端仓库的lxxx分支上

再次测试本地仓库push到lxxx分支上