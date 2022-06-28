package models

import (
	"gorm.io/gorm"
)

type Sensor struct {
	gorm.Model
	Name   string
	IP     string
	Port   string
	Values []SensorValue
}
