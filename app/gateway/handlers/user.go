package handlers

import (
	"dzug/app/gateway/rpc"
	pb "dzug/protos/user"
	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func UserRegister(ctx *gin.Context) {

	//1.获取参数 和 参数校验
	userReq := new(pb.LoginAndRegisterRequest)
	if err := ctx.ShouldBindJSON(userReq); err != nil {
		zap.L().Error("Register with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ctx.JSON(http.StatusBadRequest, pb.LoginAndRegisterResponse{
				StatusCode: 400,
				StatusMsg:  "参数错误",
				UserId:     0,
				Token:      "",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), //翻译错误
		})
		return
	}

	//2.注册业务处理
	userResp, err := rpc.Register(ctx, userReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pb.LoginAndRegisterResponse{
			StatusCode: 500,
			StatusMsg:  err.Error(),
			UserId:     0,
			Token:      "",
		})
		return
	}
	//3.返回相应
	ctx.JSON(http.StatusOK, userResp)
}

func UserLogin(ctx *gin.Context) {
	//1.获取参数及参数校验
	userReq := new(pb.LoginAndRegisterRequest)
	if err := ctx.ShouldBindJSON(userReq); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ctx.JSON(http.StatusBadRequest, pb.LoginAndRegisterResponse{
				StatusCode: 400,
				StatusMsg:  "参数错误",
				UserId:     0,
				Token:      "",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)), //翻译错误
		})
		return
	}
	userResp, err := rpc.Login(ctx, userReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pb.LoginAndRegisterResponse{
			StatusCode: 500,
			StatusMsg:  err.Error(),
			UserId:     0,
			Token:      "",
		})
		return
	}
	//3.返回相应
	ctx.JSON(http.StatusOK, userResp)
}

// UserInfo 返回用户所有信息
func UserInfo(ctx *gin.Context) {

	//1.获取参数 和 参数校验

	//2.业务处理

	//3.返回相应

}
