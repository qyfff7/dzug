package handlers

import (
	"dzug/app/gateway/rpc"
	pb "dzug/protos/relation"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RelationAction(ctx *gin.Context) {
	var relationReq pb.DouyinRelationActionRequest
	// 参数错误也能正常运行，因为json内部有个omitempty关键字
	if err := ctx.BindJSON(&relationReq); err != nil {
		ctx.JSON(http.StatusBadRequest, pb.DouyinRelationActionResponse{
			StatusCode: 400,
			StatusMsg:  "参数错误",
		})
		return
	}
	fmt.Println("我是token：", relationReq.Token)
	relationResp, err := rpc.RelationAction(ctx, &relationReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pb.DouyinRelationActionResponse{
			StatusCode: 500,
			StatusMsg:  "RPC服务调用错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, relationResp)
}
