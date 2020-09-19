package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseWithJson(code int, data interface{}, c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  GetMsg(code),
		"data": data,
	})
}
