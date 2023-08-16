package middlewares

import (
	"dzug/app/user/pkg/jwt"
	"dzug/models"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		//token 是存放在url请求参数中的，因此从url中获取token
		token := ctx.Query("token")
		mc, err := jwt.ParseToken(token)
		if err != nil {
			models.ResponseError(ctx, models.CodeInvalidToken)
			ctx.Abort()
			return
		}
		// 将当前请求的userID信息保存到请求的上下文ctx上
		ctx.Set(jwt.CtxUserIDKey, mc.UserID)
		ctx.Next()
	}
}
