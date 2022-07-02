package main

import (
	"fmt"
	"os"

	"example.com/final-exam/controller"
	"example.com/final-exam/util"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	util.ConnectDB()

	router := gin.Default()
	router.GET("/airlines", controller.GetAirlines)
	router.POST("/airline/create", controller.AddAirline)

	router.GET("/routes", controller.GetRoutes)
	router.POST("/route/create", controller.AddRoute)

	router.GET("/schedule/date", controller.GetScheduleByDate)
	router.POST("/schedule/create", controller.AddSchedule)
	router.POST("/schedule/delay", controller.DelaySchedule)
	router.POST("/schedule/cancel", controller.CancelSchedule)

	router.Run("localhost:5000")
}
