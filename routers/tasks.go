package routers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gulywwx/todo-go-api/models"
	"github.com/gulywwx/todo-go-api/utils"
)

func QueryTaskList(c *gin.Context) {
	strPageSize, _ := c.GetQuery("pageSize")
	strPageNo, _ := c.GetQuery("pageNo")
	strStatus, _ := c.GetQuery("status")

	if strPageSize == "" || strPageNo == "" {
		utils.ResponseWithJson(utils.INVALID_PARAMS, "请求参数不正确", c)
		return
	}

	pageSize, _ := strconv.Atoi(strPageSize)
	pageNo, _ := strconv.Atoi(strPageNo)
	var status int = -1
	if strStatus != "" {
		status, _ = strconv.Atoi(strStatus)
	}

	fmt.Println(status)

	taskList, _ := models.GetTaskList(status)
	total := len(taskList)

	n := (pageNo - 1) * pageSize

	if total == 0 {
		utils.ResponseWithJson(utils.ERROR, gin.H{}, c)
		return
	}

	if n > total {
		n = total
	}

	end := n + pageSize
	if end > total {
		end = total
	}

	utils.ResponseWithJson(utils.SUCCESS, gin.H{
		"rows":     taskList[n:end],
		"total":    total,
		"pageNo":   pageNo,
		"pageSize": pageSize,
	}, c)

}

type millisecondsTime time.Time

func (j *millisecondsTime) UnmarshalJSON(data []byte) error {
	millis, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*j = millisecondsTime(time.Unix(0, millis*int64(time.Millisecond)))
	return nil
}

func AddTask(c *gin.Context) {
	type RequestBody struct {
		Title   string           `json:"title" binding:"required"`
		Content string           `json:"content" binding:"required"`
		Expired millisecondsTime `json:"expired" binding:"required"`
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		utils.ResponseWithJson(utils.INVALID_PARAMS, "invalid params", c)
		return
	}

	models.AddTask(body.Title, body.Content, time.Time(body.Expired))
	utils.ResponseWithJson(utils.SUCCESS, "task added successfully", c)
}

func EditTask(c *gin.Context) {
	type RequestBody struct {
		ID      string           `json:"id" binding:"required"`
		Title   string           `json:"title" binding:"required"`
		Content string           `json:"content" binding:"required"`
		Expired millisecondsTime `json:"expired" binding:"required"`
	}
	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		utils.ResponseWithJson(utils.INVALID_PARAMS, "invalid params", c)
		return
	}

	models.EditTask(body.ID, body.Title, body.Content, time.Time(body.Expired))
	utils.ResponseWithJson(utils.SUCCESS, "task edited successfully", c)
}

func UpdateTaskStatus(c *gin.Context) {
	type RequestBody struct {
		ID     string `json:"id" binding:"required"`
		Status *int   `json:"status" binding:"required"`
	}
	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		utils.ResponseWithJson(utils.INVALID_PARAMS, "invalid params", c)
		return
	}

	models.UpdateTaskStatus(body.ID, *body.Status)
	utils.ResponseWithJson(utils.SUCCESS, "task status updated successfully", c)
}

func UpdateMark(c *gin.Context) {
	type RequestBody struct {
		ID      string `json:"id" binding:"required"`
		IsMajor *bool  `json:"is_major" binding:"required"`
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
		utils.ResponseWithJson(utils.INVALID_PARAMS, "invalid params", c)
		return
	}

	if _, err := models.UpdateMark(body.ID, *body.IsMajor); err != nil {
		utils.ResponseWithJson(utils.ERROR, err.Error(), c)
	}

	utils.ResponseWithJson(utils.SUCCESS, "mark set successfully", c)

}

func DeleteTask(c *gin.Context) {
	type RequestBody struct {
		ID string `json:"id" binding:"required"`
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		utils.ResponseWithJson(utils.INVALID_PARAMS, "invalid params", c)
		return
	}

	models.DeleteTask(body.ID)
	utils.ResponseWithJson(utils.SUCCESS, "task deleted successfully", c)

}
