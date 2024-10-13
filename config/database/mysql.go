package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	var err error
	dsn := "root:admin@tcp(localhost:3302)/go1?charset=utf8mb4&parseTime=True"

	newLogger := logger.Default
	newLogger.LogMode(logger.Info)

	config := gorm.Config{
		Logger: newLogger,
	}

	DB, err = gorm.Open(mysql.Open(dsn), &config)
	if err != nil {
		log.Fatalf("Could't connect to db: %v", err)
	}

}
