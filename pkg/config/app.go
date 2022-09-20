package config

import (
	"fmt"

	"github.com/jaypey/GoPlant/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "host=172.18.0.2 user=postgres password=postgres dbname=goplantdb port=5432 sslmode=disable"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("- [Database error]")
		panic(err)
	}
	fmt.Println("+ [Database Connected]")
	d.AutoMigrate(&models.Sensor{}, &models.SensorValue{})
	fmt.Println("+ [Database Migrated]")

	db = d
}

func GetDB() *gorm.DB {
	return db
}
