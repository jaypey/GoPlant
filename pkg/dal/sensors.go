package dal

import (
	"fmt"

	"github.com/jaypey/GoPlant/pkg/config"
	"github.com/jaypey/GoPlant/pkg/models"
)

func AddSensor(sensor *models.Sensor) (uint, error) {
	result := config.GetDB().Create(sensor)
	if result.Error != nil {
		fmt.Println(result.Error)
		return 0, result.Error
	}
	return sensor.ID, nil
}

func GetAllSensors() ([]models.Sensor, error) {
	var sensors []models.Sensor
	result := config.GetDB().Find(&sensors)
	if result.Error != nil {
		return sensors, result.Error
	}
	return sensors, nil
}

func GetSensor(id uint) (models.Sensor, error) {
	var sensor models.Sensor
	result := config.GetDB().Preload("Values").First(&sensor, id)
	if result.Error != nil {
		return sensor, result.Error
	}
	return sensor, nil
}

func GetSensorByNameAndIP(name string, ip string) (models.Sensor, error) {
	var sensor models.Sensor
	result := config.GetDB().Preload("Values").Where("Name = ? AND IP = ?", name, ip).First(&sensor)
	if result.Error != nil {
		return sensor, result.Error
	}
	return sensor, nil
}
