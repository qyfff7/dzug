


* 用户注册、登录、查询用户信息服务完成
* JWT Token认证

存在的问题：
1. 查看用户信息时，IsFollow属性需要查relation表，
2. 视频流的处理未完成

etcd下载教程：https://www.cnblogs.com/wusenwusen/p/16572929.html



关于解析Token,获取user_id这个问题，我看了一下我的代码，在jwt.go文件中写好了生成和解析Token的函数，因此，整体操作如下：
1. 从lxxx分支上获取dzug/app/user/pkg/jwt/jwt.go 文件（这个是生成Token和解析Token的文件），以及dzug/app/gateway/middlewares/JWTmiddleware.go文件（这个是中间件，所有需要登录才能使用的路由都将使用这个中间件）
2. 当用户登录后，通过jwt的GenToken函数，利用UserID生成了Token
3. 在自己的业务代码中，当想要通过Token获取UserID时，只需执行以下代码：
```go
//  1.获取当前请求中的token
	authHeader := ctx.Request.Header.Get("Authorization")    //ctx 是 Context
	parts := strings.SplitN(authHeader, " ", 2)
	token := parts[1]
	//2.从token中解析处userID
	u, err := jwt.ParseToken(token)
	if err != nil {
		//错误处理
		return
	}
	userID := u.UserID
	//此处userID就是当前用户的ID
```
3. 然后进行相关的业务操作。
