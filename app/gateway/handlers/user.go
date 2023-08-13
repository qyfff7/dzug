package handlers

import (
	"dzug/app/gateway/rpc"
	"dzug/models"
	"dzug/protos/user"
	"encoding/json"
	"fmt"
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
			models.ResponseError(ctx, models.CodeInvalidParam)

			return
		}
		err, _ := json.Marshal(removeTopStruct(errs.Translate(trans)))
		models.ResponseErrorWithMsg(ctx, models.CodeInvalidParam, string(err))

		/*ctx.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), //翻译错误
		})*/
		return
	}

	//2.注册业务处理
	userResp, err := rpc.Register(ctx, userReq)
	if err != nil {
		models.ResponseErrorWithMsg(ctx, models.CodeServerBusy, err.Error())

		/*ctx.JSON(http.StatusInternalServerError, user.AccountResp{
			StatusCode: 500,
			StatusMsg:  err.Error(),
			UserId:     0,
			Token:      "",
		})*/
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
			/*ctx.JSON(http.StatusBadRequest, user.AccountResp{
				StatusCode: 400,
				StatusMsg:  "参数错误",
				UserId:     0,
				Token:      "",
			})*/
			return
		}
		err, _ := json.Marshal(removeTopStruct(errs.Translate(trans)))
		models.ResponseErrorWithMsg(ctx, models.CodeInvalidParam, string(err))
		/*ctx.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), //翻译错误
		})*/
		return
	}
	userResp, err := rpc.Login(ctx, userReq)
	if err != nil {
		models.ResponseErrorWithMsg(ctx, models.CodeInvalidPassword, err.Error())
		/*ctx.JSON(http.StatusInternalServerError, user.AccountResp{
			StatusCode: 500,
			StatusMsg:  err.Error(),
			UserId:     0,
			Token:      "",
		})*/
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
		return
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			models.ResponseError(ctx, models.CodeInvalidParam)

			/*ctx.JSON(http.StatusBadRequest, user.GetUserInfoResp{
				StatusCode: 400,
				StatusMsg:  "获取用户信息时参数错误",
				UserInfo:   nil,
			})*/
			return
		}
		err, _ := json.Marshal(removeTopStruct(errs.Translate(trans)))
		models.ResponseErrorWithMsg(ctx, models.CodeInvalidParam, string(err))
		/*ctx.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), //翻译错误
		})*/
		return
	}

	//getuserInfoReq := new(user.GetUserInfoReq)
	//getuserInfoReq.UserId = to_user_id.UserId

	zap.L().Info("hahhhhhhhhhhh")
	zap.L().Info(fmt.Sprintln(userInfoReq.UserId))
	zap.L().Info(fmt.Sprintln(userInfoReq.Token))
	zap.L().Info("handeler   hanlder")

	//从请求头中获取Token
	authHeader := ctx.Request.Header.Get("Authorization") //ctx 是 Context
	parts := strings.SplitN(authHeader, " ", 2)
	userInfoReq.Token = parts[1]

	//4.查询用户信息
	userInfo, err := rpc.UserInfo(ctx, userInfoReq)
	if err != nil {
		models.ResponseErrorWithMsg(ctx, models.CodeServerBusy, err.Error())
		/*ctx.JSON(http.StatusInternalServerError, user.GetUserInfoResp{
			StatusCode: 500,
			StatusMsg:  err.Error(),
			UserInfo:   nil,
		})*/
		return
	}
	//3.返回相应
	u := models.UserInfoResp(userInfo)
	models.GetUserInfoSuccess(ctx, u)

	//models.ResponseSuccess(ctx, userInfo)
	//ctx.JSON(http.StatusOK, userInfoResp)
}
