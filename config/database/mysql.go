package database

import (
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// to log in file
func NewLogger() *logrus.Logger {
	log := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)

	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.JSONFormatter{})

	return log
}

func init() {
	var err error
	dsn := "root:admin@tcp(localhost:3302)/go1?charset=utf8mb4&parseTime=True"

	// newLogger := logger.Default
	// newLogger.LogMode(logger.Info)

	loggerConfig := logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: false,
		ParameterizedQueries:      false,
		Colorful:                  false,
	}

	newLogger := logger.New(
		NewLogger(),
		loggerConfig,
	)

	config := gorm.Config{
		Logger: newLogger,
	}

	DB, err = gorm.Open(mysql.Open(dsn), &config)
	if err != nil {
		log.Fatalf("Could't connect to db: %v", err)
	}

}
