package middlewares

import (
	"dzug/app/user/pkg/jwt"
	"dzug/models"
	"github.com/gin-gonic/gin"
	"strings"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// Authorization: Bearer xxxxxxx.xxx.xxx  / X-TOKEN: xxx.xxx.xx
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			models.ResponseError(ctx, models.CodeNeedLogin, nil)
			/*	c.JSON(http.StatusOK, pb.AccountResp{
				StatusCode: 500,
				StatusMsg:  "需要登录后才能进行操作",
				UserId:     0,
				Token:      "",
			})*/
			ctx.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {

			models.ResponseError(ctx, models.CodeInvalidToken, nil)

			/*c.JSON(http.StatusOK, pb.AccountResp{
				StatusCode: 500,
				StatusMsg:  "当前Token无效",
				UserId:     0,
				Token:      "",
			})*/

			ctx.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			models.ResponseError(ctx, models.CodeInvalidToken, nil)
			/*c.JSON(http.StatusOK, pb.AccountResp{
				StatusCode: 500,
				StatusMsg:  "当前Token无效",
				UserId:     0,
				Token:      "",
			})*/
			ctx.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文ctx上
		ctx.Set(jwt.CtxUserIDKey, mc.UserID)

		ctx.Next()
		// 后续的处理请求的函数中 可以用过
		//uuid ,err:= jwt.GetUserID(c)
		//来获取当前请求的用户的userID,进而查询数据库，获取所有的信息
	}
}
