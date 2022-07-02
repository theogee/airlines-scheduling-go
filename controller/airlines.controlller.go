package controller

import (
	"net/http"

	"example.com/final-exam/dao"
	"example.com/final-exam/model"
	v "example.com/final-exam/validator"
	"github.com/gin-gonic/gin"
)

func GetAirlines(c *gin.Context) {
	data, err := dao.GetAirlines()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "data": data})
}

func AddAirline(c *gin.Context) {
	var newAirline model.Airline
	if err := c.BindJSON(&newAirline); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	errVal, errMsg := v.ValidateAirline(newAirline)
	if errVal == true {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"success": false, "msg": errMsg})
		return
	}

	id, err := dao.AddAirline(newAirline)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"success": true, "id": id})
}
