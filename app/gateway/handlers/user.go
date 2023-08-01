package handlers

import (
	"dzug/app/gateway/rpc"
	pb "dzug/protos/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(ctx *gin.Context) {
	var userReq pb.DouyinUserRegisterRequest
	if err := ctx.Bind(&userReq); err != nil {
		ctx.JSON(http.StatusBadRequest, pb.DouyinUserRegisterResponse{
			StatusCode: 400,
			StatusMsg:  "参数错误",
			UserId:     0,
			Token:      "",
		})
		return
	}
	userResp, err := rpc.UserRegister(ctx, &userReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pb.DouyinUserRegisterResponse{
			StatusCode: 500,
			StatusMsg:  "RPC服务调用错误",
			UserId:     0,
			Token:      "",
		})
		return
	}
	ctx.JSON(http.StatusOK, userResp)
}

func UserLogin(ctx *gin.Context) {
	var userReq pb.DouyinUserLoginRequest
	if err := ctx.Bind(&userReq); err != nil {
		ctx.JSON(http.StatusBadRequest, pb.DouyinUserRegisterResponse{
			StatusCode: 400,
			StatusMsg:  "参数错误",
			UserId:     0,
			Token:      "",
		})
		return
	}

	userResp, err := rpc.UserLogin(ctx, &userReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pb.DouyinUserLoginResponse{
			StatusCode: 500,
			StatusMsg:  "RPC服务调用错误",
			UserId:     0,
			Token:      "",
		})
		return
	}

	ctx.JSON(http.StatusOK, userResp)
}
