package models

import (
	"gorm.io/gorm"
)

type SensorValue struct {
	gorm.Model
	Value    float64
	SensorID uint
}
