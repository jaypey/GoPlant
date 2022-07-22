package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jaypey/GoPlant/pkg/dal"
	"github.com/jaypey/GoPlant/pkg/models"
)

func GetSensors(c *gin.Context) {
	sensors, err := dal.GetAllSensors()
	if err != nil {
		c.AbortWithStatus(500)
	}
	c.IndentedJSON(http.StatusOK, sensors)
}

func AddSensor(c *gin.Context) {
	var newSensor models.Sensor
	if err := c.BindJSON(&newSensor); err != nil {
		return
	}
	id, err := dal.AddSensor(&newSensor)
	if err != nil {
		c.AbortWithStatus(500)
	}
	c.IndentedJSON(http.StatusCreated, id)

}
