package http

import (
	"dzug/app/gateway/rpc"
	pb "dzug/idl/relation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RelationAction(ctx *gin.Context) {
	var relationReq pb.DouyinRelationActionRequest
	if err := ctx.Bind(&relationReq); err != nil {
		ctx.JSON(http.StatusBadRequest, pb.DouyinRelationActionResponse{
			StatusCode: 400,
			StatusMsg:  "参数错误",
		})
		return
	}
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
