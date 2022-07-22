package dal

import (
	"fmt"

	"github.com/jaypey/GoPlant/pkg/config"
	"github.com/jaypey/GoPlant/pkg/models"
)

func AddSensorValue(sensorValue *models.SensorValue) (uint, error) {
	result := config.GetDB().Create(sensorValue)
	if result.Error != nil {
		fmt.Println(result.Error)
		return 0, result.Error
	}
	return sensorValue.ID, nil
}
