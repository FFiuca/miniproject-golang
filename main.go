package main

import (
	"fmt"
	"os"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"project1/app/controller/router"
	"project1/config/database"
)

var app *fiber.App
var db *gorm.DB
var log *logrus.Logger

// var validation *validator.Validate

func NewLogger() *logrus.Logger {
	log := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)

	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.JSONFormatter{})

	return log
}

func Catch() {
	fmt.Println("masuk catch global")
	if recover() != nil {
		fmt.Println("panic triggered")
	}
}

func main() {
	defer Catch()

	config := fiber.Config{
		// Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
	}

	app := fiber.New(config)

	routeConfig := router.RouteConfig{
		App:       app,
		Log:       NewLogger(),
		DB:        database.DB,
		Validator: validator.New(),
	}

	routeConfig.Setup()

	app.Static("/static", "./static")

	app.Listen(":8888")
}
