package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gulywwx/todo-go-api/models"
	"github.com/gulywwx/todo-go-api/utils"
)

func Login(c *gin.Context) {

	type RequestBody struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		utils.ResponseWithJson(utils.INVALID_PARAMS, "invalid params", c)
		return
	}

	user, err := models.FindUser(body.Username)
	if err != nil {
		utils.ResponseWithJson(utils.ERROR, err.Error(), c)
		return
	}

	if body.Password != user.Password {
		utils.ResponseWithJson(utils.ERROR_AUTH_PASSWORD, "invalid password", c)
		return
	}

	token, err := utils.GeneratorToken(user.ID, user.Name)
	if err != nil {
		utils.ResponseWithJson(utils.ERROR, "token generation failed", c)
		return
	}

	utils.ResponseWithJson(utils.SUCCESS, gin.H{
		"user":  user,
		"token": token,
	}, c)

}

func Register(c *gin.Context) {
	type RequestBody struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		utils.ResponseWithJson(utils.INVALID_PARAMS, "invalid params", c)
		return
	}

	user, err := models.FindUser(body.Username)
	if err == nil && user != nil {
		utils.ResponseWithJson(utils.ERROR, "user has existed", c)
		return
	}

	user, err = models.CreateUser(body.Username, body.Password)
	if err != nil {
		utils.ResponseWithJson(utils.ERROR, err.Error(), c)
		return
	}

	token, err := utils.GeneratorToken(user.ID, user.Name)
	if err != nil {
		utils.ResponseWithJson(utils.ERROR, "token generation failed", c)
		return
	}

	utils.ResponseWithJson(utils.SUCCESS, gin.H{
		"user":  user,
		"token": token,
	}, c)

}

func ResetPwd(c *gin.Context) {
	type RequestBody struct {
		Username    string `json:"username" binding:"required"`
		OldPassword string `json:"oldPassword" binding:"required"`
		NewPassword string `json:"newPassword" binding:"required"`
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		utils.ResponseWithJson(utils.INVALID_PARAMS, "invalid params", c)
		return
	}
	if body.OldPassword == body.NewPassword {
		utils.ResponseWithJson(utils.ERROR, "new pass and new pass could not be same", c)
		return
	}

	err := models.ResetPass(body.Username, body.OldPassword, body.NewPassword)
	if err != nil {
		utils.ResponseWithJson(utils.ERROR, err.Error(), c)
		return
	}

	utils.ResponseWithJson(utils.SUCCESS, "password reset successfully", c)

}
