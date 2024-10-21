package router

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"project1/app/controller"
	"project1/app/controller/middleware"
	"project1/app/models"
	"project1/app/services"
	"project1/app/usecases"
)

type RouteConfig struct {
	App       *fiber.App
	Log       *logrus.Logger
	DB        *gorm.DB
	Validator *validator.Validate
}

func (c *RouteConfig) Setup() {
	c.GuestRouter()
	c.AuthRouter()
	c.Dummy()
}

func (c *RouteConfig) GuestRouter() {
	userService := services.NewUserService(c.DB)

	authService := services.NewAuthService(c.DB, userService)
	authUseCase := usecases.NewAuthUseCase(c.DB, c.Log, c.Validator, authService)
	authController := controller.NewAuthController(c.Log, authUseCase)

	c.App.Post("/api/login", authController.Login)
}

func (c *RouteConfig) AuthRouter() {
	statusService := services.NewStatusService(c.DB)
	statusUseCase := usecases.NewStatusUseCase(c.DB, c.Log, c.Validator, statusService)
	statusController := controller.NewStatusController(statusUseCase, c.Log)

	c.App.Post("/api/status", statusController.Create)
	c.App.Put("/api/status/:id", statusController.Update)
	c.App.Get("/api/status/:id", statusController.Detail)
	c.App.Delete("/api/status/:id", statusController.Delete)

	userService := services.NewUserService(c.DB)
	userUserCase := usecases.NewUserUseCase(c.DB, c.Log, c.Validator, userService)
	userController := controller.NewUserController(c.Log, userUserCase)

	c.App.Post("/api/user", userController.Create)
	c.App.Put("/api/user/:id", userController.Update)
	c.App.Get("/api/user/:id", userController.Detail)
	c.App.Delete("/api/user/:id", userController.Delete)
	c.App.Get("/api/user", userController.Search)

	c.App.Use(middleware.AuthMiddleware())
	c.App.Get("/api/auth", func(ctx *fiber.Ctx) error {
		fmt.Println(ctx.Locals("metaUser"))
		return ctx.SendStatus(fiber.StatusOK)
	})

	c.App.Get("/api/coba_db", func(ctx *fiber.Ctx) error {
		a := []models.Event{}

		c.DB.Model(&a).Preload("Status").Preload("User").Preload("EventUser.Status").Find(&a)

		fmt.Println(a[0].EventUser)

		b := []models.Event{}
		c.DB.Preload("User").Find(&b)

		fmt.Println(b)

		return ctx.JSON(a)
	})

}

func (c *RouteConfig) Dummy() {
	c.App.Get("/coba", func(ctx *fiber.Ctx) error {
		a := models.User{}
		c.DB.First(&a, 1)

		fmt.Println(a.Email, a.Status.Name, reflect.TypeOf(a), reflect.TypeOf(a.Status))
		return nil
	})
}
