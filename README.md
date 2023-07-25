# dzug
dousheng demo

准备使用go-zero进行一次测试
主要使用接口为：

1. /douyin/user/login
2. /douyin/user/register
3. /douyin/relation/action

前两者用于登录注册，最后者用于关注和取关
拆分两个微服务，12为用户服务，3为关系服务

1. 数据库建立
2. go-zero模板
3. 业务逻辑
4. 代码测试

---

demo进度

1. 主要参考go-zero 文档：https://go-zero.dev/docs/tasks
2. 工具安装均已完成 
3. 首先创建文件结构 
4. 然后编写api文件  （goctl api go --api user.api --dir .）
5. proto文件  （goctl rpc protoc ./rpc/user.proto --go_out=./rpc/pb --go-grpc_out=./rpc/pb --zrpc_out=./rpc）
5. rpc文件，写法都参考文档即可
6. 数据库未完成，逻辑部分处理暂停
---
假如使用zero，前期使用工具配置可能比较困难，同时，gin框架和gorm应该不再需要，go-zero自带有api控制和自有orm

算是一个比较成熟的框架