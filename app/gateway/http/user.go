package http

import (
	pb "dzug/idl"
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
}
