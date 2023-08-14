package handlers

import (
	"dzug/app/gateway/rpc"
	"dzug/models"
	"dzug/protos/user"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func UserRegister(ctx *gin.Context) {

	//1.获取参数 和 参数校验
	userReq := new(user.AccountReq)
	if err := ctx.ShouldBind(userReq); err != nil {
		zap.L().Error("Register with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			models.ResponseError(ctx, models.CodeInvalidParam)

			return
		}
		err, _ := json.Marshal(removeTopStruct(errs.Translate(trans)))
		models.ResponseErrorWithMsg(ctx, models.CodeInvalidParam, string(err))

		return
	}

	//2.注册业务处理
	userResp, err := rpc.Register(ctx, userReq)
	if err != nil {
		models.ResponseErrorWithMsg(ctx, models.CodeServerBusy, err.Error())
		return
	}
	//3.返回相应
	models.AccountRespSuccess(ctx, userResp)

}

func UserLogin(ctx *gin.Context) {
	//1.获取参数及参数校验
	userReq := new(user.AccountReq)
	if err := ctx.ShouldBind(userReq); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			models.ResponseError(ctx, models.CodeInvalidParam)

			return
		}
		err, _ := json.Marshal(removeTopStruct(errs.Translate(trans)))
		models.ResponseErrorWithMsg(ctx, models.CodeInvalidParam, string(err))

		return
	}
	//2.调用登录服务
	userResp, err := rpc.Login(ctx, userReq)
	if err != nil {
		models.ResponseErrorWithMsg(ctx, models.CodeInvalidPassword, err.Error())
		return
	}
	//3.返回相应
	models.AccountRespSuccess(ctx, userResp)

}

// UserInfo 返回用户所有信息
func UserInfo(ctx *gin.Context) {

	//1.获取参数及参数校验
	userInfoReq := new(user.GetUserInfoReq)
	if err := ctx.ShouldBind(userInfoReq); err != nil {
		zap.L().Error("GetUserInfo with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			models.ResponseError(ctx, models.CodeInvalidParam)
			return
		}
		err, _ := json.Marshal(removeTopStruct(errs.Translate(trans)))
		models.ResponseErrorWithMsg(ctx, models.CodeInvalidParam, string(err))
		return
	}

	//4.查询用户信息
	userInfo, err := rpc.UserInfo(ctx, userInfoReq)
	if err != nil {
		models.ResponseErrorWithMsg(ctx, models.CodeServerBusy, err.Error())
		return
	}
	//3.返回相应
	u := models.UserInfoResp(userInfo)
	models.GetUserInfoSuccess(ctx, u)

}
