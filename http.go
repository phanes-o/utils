package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseOk(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": data,
	})
}
