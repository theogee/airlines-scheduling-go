package controller

import (
	"fmt"
	"net/http"
	"time"

	"example.com/final-exam/dao"
	"example.com/final-exam/model"
	v "example.com/final-exam/validator"
	"github.com/gin-gonic/gin"
)

func GetSchedules(c *gin.Context) {
	data, err := dao.GetSchedules()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "data": data})
}

func GetScheduleByDate(c *gin.Context) {
	date := c.Query("d")
	fmt.Println(date)
	data, err := dao.GetScheduleByDate(date)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "data": data})
}

func AddSchedule(c *gin.Context) {
	var newSchedule model.Schedule
	if err := c.BindJSON(&newSchedule); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	errVal, errMsg := v.ValidateSchedule(newSchedule)
	if errVal == true {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"success": false, "msg": errMsg})
		return
	}

	id, err := dao.AddSchedule(newSchedule)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"success": true, "id": id})
}

func DelaySchedule(c *gin.Context) {
	var delay model.Delay
	if err := c.BindJSON(&delay); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"success": false, "msg": err})
		return
	}

	errVal, errMsg := v.ValidateDelay(delay)
	if errVal == true {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"success": false, "msg": errMsg})
		return
	}

	rowsAffected, err := dao.DelaySchedule(delay)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	if rowsAffected == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"success": false, "msg": "The given scheduleID was not found"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"success": true, "msg": fmt.Sprintf("Schedule id: %v has been delayed", delay.ScheduleID)})
}

func CancelSchedule(c *gin.Context) {
	var cancel model.Cancel
	if err := c.BindJSON(&cancel); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"success": false, "msg": err})
		return
	}

	errVal, errMsg := v.ValidateCancel(cancel)
	if errVal == true {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"success": false, "msg": errMsg})
		return
	}

	rowsAffected, err := dao.CancelSchedule(cancel)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	if rowsAffected == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"success": false, "msg": "The given scheduleID was not found"})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"success": true, "msg": fmt.Sprintf("Schedule id: %v has been canceled", cancel.ScheduleID)})
}

func ScheduleAuto() {
	for {
		dao.ScheduleAuto()
		time.Sleep(time.Second * 60)
	}
}
