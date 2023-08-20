package middlewares

import (
	"dzug/app/redis"
	"dzug/app/user/pkg/jwt"
	"dzug/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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

		oldtoken, err := redis.Rdb.HGet(ctx, redis.GetRedisKey(redis.KeyUserInfo, strconv.Itoa(int(mc.UserID))), "token").Result()

		if err != nil {
			zap.L().Error("获取redis中的用户 token 出错", zap.Error(err))
			models.ResponseError(ctx, models.CodeServerBusy)
			ctx.Abort()
			return
		}
		if token != oldtoken {
			zap.L().Error("当前账号已在其他设备登录，请重新登录后使用")
			models.ResponseError(ctx, models.CodeInvalidToken)
			ctx.Abort()
			return
		}

		// 将当前请求的userID信息保存到请求的上下文ctx上
		ctx.Set(jwt.CtxUserIDKey, mc.UserID)
		ctx.Next()
	}
}
