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
