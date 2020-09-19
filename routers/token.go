package routers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gulywwx/todo-go-api/utils"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		token := authHeader[len(BEARER_SCHEMA):]

		if token == "" {
			utils.ResponseWithJson(utils.ERROR_AUTH_TOKEN, "", c)
			c.Abort()
			return
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				utils.ResponseWithJson(utils.ERROR_AUTH_CHECK_TOKEN_FAIL, "", c)
				c.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt {
				utils.ResponseWithJson(utils.ERROR_AUTH_CHECK_TOKEN_TIMEOUT, "", c)
				c.Abort()
				return
			} else {
				c.Set("ID", claims.ID)
				c.Next()
			}
		}
	}

}
