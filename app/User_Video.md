etcd下载教程：https://www.cnblogs.com/wusenwusen/p/16572929.html

更新关于Token和userID的问题：
1. 由于在项目中，官方要求是从url中直接获取token的值，（而不是http请求头中） 因此，获取token的方式变更为`token := ctx.Query("token")`，
2. 其他操作不变，获取userID既可以解析token，也可jwt.GetUserID(ctx)

待办事项：
1. user的信息存到redis
2. 视频流的分页查询
3. 对于日志部分，消息队列



