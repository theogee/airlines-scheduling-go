package controller

import (
	"net/http"

	"example.com/final-exam/dao"
	"example.com/final-exam/model"
	v "example.com/final-exam/validator"
	"github.com/gin-gonic/gin"
)

func GetRoutes(c *gin.Context) {
	data, err := dao.GetRoutes()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"success": true, "data": data})
}

func AddRoute(c *gin.Context) {
	var newRoute model.Route
	if err := c.BindJSON(&newRoute); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	errVal, errMsg := v.ValidateRoute(newRoute)
	if errVal == true {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"success": false, "msg": errMsg})
		return
	}

	id, err := dao.AddRoute(newRoute)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"success": true, "id": id})
}
