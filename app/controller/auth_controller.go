package controller

import (
	"project1/app/dto"
	"project1/app/usecases"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AuthController struct {
	Log     *logrus.Logger
	UseCase usecases.AuthUseCaseInterface
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	request := new(dto.AuthRequestLogin)

	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("error parse body %v", err)

		return ctx.JSON(dto.WebResponse[string]{
			Status: 500,
			Data:   "error parse body",
		})
	}

	login, err := c.UseCase.Login(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to login : %v", err)
		switch err.(type) {
		case *fiber.Error:
			return ctx.Status(fiber.StatusBadRequest).JSON(dto.WebResponse[string]{
				Status: 400,
				Data:   err.Error(),
			})
		case error:
			return ctx.Status(fiber.StatusInternalServerError).JSON(dto.WebResponse[string]{
				Status: 500,
				Data:   err.Error(),
			})
		default:
			return ctx.Status(fiber.StatusInternalServerError).JSON(dto.WebResponse[string]{
				Status: 500,
				Data:   "Unknown error",
			})
		}
	}

	return ctx.JSON(dto.WebResponse[dto.AuthResponse]{
		Status: 200,
		Data:   *login,
	})
}

func NewAuthController(log *logrus.Logger, uc usecases.AuthUseCaseInterface) *AuthController {
	return &AuthController{
		Log:     log,
		UseCase: uc,
	}
}
