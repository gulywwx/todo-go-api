package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("/api")
	{
		api.POST("login", Login)
		api.POST("register", Register)
		api.POST("resetPwd", JWTAuth(), ResetPwd)

		api.GET("queryTaskList", JWTAuth(), QueryTaskList)
		api.POST("addTask", JWTAuth(), AddTask)
		api.PUT("editTask", JWTAuth(), EditTask)
		api.PUT("updateTaskStatus", JWTAuth(), UpdateTaskStatus)
		api.PUT("updateMark", JWTAuth(), UpdateMark)
		api.DELETE("deleteTask", JWTAuth(), DeleteTask)

		api.GET("version", JWTAuth(), GetAppVersionTest)
	}

	return r
}
