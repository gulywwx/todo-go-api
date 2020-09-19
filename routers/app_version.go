package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gulywwx/todo-go-api/utils"
)

//app更新接口
func GetAppVersionTest(c *gin.Context) {

	c.JSON(utils.SUCCESS, gin.H{
		"code": utils.SUCCESS,
		"msg":  utils.GetMsg(utils.SUCCESS),
		"data": "return data successful",
	})
}
