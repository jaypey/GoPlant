package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jaypey/GoPlant/pkg/config"
	"github.com/jaypey/GoPlant/pkg/controllers"
	"github.com/jaypey/GoPlant/pkg/udpreceiver"
)

func main() {
	config.Connect()

	router := gin.Default()
	router.GET("/sensors", controllers.GetSensors)
	router.POST("/sensor", controllers.AddSensor)
	go udpreceiver.ListenPacket()
	router.Run("localhost:8080")
}
