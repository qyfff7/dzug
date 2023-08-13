package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func List(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"msg": "cishi",
	})
}

// Feed 视频流
func Feed(c *gin.Context) {

}
