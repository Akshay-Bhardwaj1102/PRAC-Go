package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"taskapi.com/model"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("SERVER")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database...")
	}
	db.AutoMigrate(&model.Employee{})
	DB = db
}
