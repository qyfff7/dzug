package handlers

import (
	"dzug/app/gateway/rpc"
	"dzug/models"
	"dzug/protos/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strings"
)

func UserRegister(ctx *gin.Context) {

	//1.获取参数 和 参数校验
	userReq := new(user.AccountReq)
	if err := ctx.ShouldBind(userReq); err != nil {
		zap.L().Error("Register with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			models.ResponseError(ctx, models.CodeInvalidParam, models.AccountResp{
				UserID: 0,
				Token:  "",
			})
			//既然已经出错，是不是可以不返回userid和token
			//models.ResponseError(ctx, models.CodeInvalidParam,nil)
			return
		}
		models.ResponseError(ctx, models.CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		/*ctx.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), //翻译错误
		})*/
		return
	}

	//2.注册业务处理
	userResp, err := rpc.Register(ctx, userReq)
	if err != nil {
		models.ResponseErrorWithMsg(ctx, models.CodeServerBusy, err.Error(), nil)

		/*ctx.JSON(http.StatusInternalServerError, user.AccountResp{
			StatusCode: 500,
			StatusMsg:  err.Error(),
			UserId:     0,
			Token:      "",
		})*/
		return
	}
	//3.返回相应
	models.ResponseSuccess(ctx, userResp)
	//ctx.JSON(http.StatusOK, userResp)
}

func UserLogin(ctx *gin.Context) {
	//1.获取参数及参数校验
	userReq := new(user.AccountReq)
	if err := ctx.ShouldBindJSON(userReq); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			models.ResponseError(ctx, models.CodeInvalidParam, nil)
			/*ctx.JSON(http.StatusBadRequest, user.AccountResp{
				StatusCode: 400,
				StatusMsg:  "参数错误",
				UserId:     0,
				Token:      "",
			})*/
			return
		}
		models.ResponseError(ctx, models.CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		/*ctx.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), //翻译错误
		})*/
		return
	}
	userResp, err := rpc.Login(ctx, userReq)
	if err != nil {
		models.ResponseErrorWithMsg(ctx, models.CodeInvalidPassword, err.Error(), nil)
		/*ctx.JSON(http.StatusInternalServerError, user.AccountResp{
			StatusCode: 500,
			StatusMsg:  err.Error(),
			UserId:     0,
			Token:      "",
		})*/
		return
	}
	//3.返回相应
	models.ResponseSuccess(ctx, userResp)
	//ctx.JSON(http.StatusOK, userResp)
}

// UserInfo 返回用户所有信息
func UserInfo(ctx *gin.Context) {

	//1.获取参数及参数校验
	to_user_id := new(models.GetUserInfoReq)
	if err := ctx.ShouldBind(to_user_id); err != nil {
		zap.L().Error("GetUserInfo with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			models.ResponseError(ctx, models.CodeInvalidParam, nil)

			/*ctx.JSON(http.StatusBadRequest, user.GetUserInfoResp{
				StatusCode: 400,
				StatusMsg:  "获取用户信息时参数错误",
				UserInfo:   nil,
			})*/
			return
		}
		models.ResponseError(ctx, models.CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		/*ctx.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), //翻译错误
		})*/
		return
	}
	getuserInfoReq := new(user.GetUserInfoReq)
	getuserInfoReq.UserId = to_user_id.UserId

	//从请求头中获取Token
	authHeader := ctx.Request.Header.Get("Authorization") //ctx 是 Context
	parts := strings.SplitN(authHeader, " ", 2)
	getuserInfoReq.Token = parts[1]

	//4.查询用户信息
	userInfo, err := rpc.UserInfo(ctx, getuserInfoReq)
	if err != nil {
		models.ResponseErrorWithMsg(ctx, models.CodeServerBusy, err.Error(), nil)
		/*ctx.JSON(http.StatusInternalServerError, user.GetUserInfoResp{
			StatusCode: 500,
			StatusMsg:  err.Error(),
			UserInfo:   nil,
		})*/
		return
	}
	//3.返回相应
	models.ResponseSuccess(ctx, userInfo)
	//ctx.JSON(http.StatusOK, userInfoResp)

}
