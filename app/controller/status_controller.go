package controller

import (
	"fmt"
	"project1/app/dto"
	"project1/app/usecases"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type StatusController struct {
	Log     *logrus.Logger
	UseCase usecases.StatusUseCaseInterface
}

func NewStatusController(useCase usecases.StatusUseCaseInterface, log *logrus.Logger) *StatusController {
	return &StatusController{
		Log:     log,
		UseCase: useCase,
	}
}

func (c *StatusController) Create(ctx *fiber.Ctx) error {
	request := new(dto.StatusRequestCreate)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse body : %+v", err)
	}
	fmt.Println("check", request)
	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create status : %v", err)
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
	fmt.Println("cok")

	return ctx.JSON(dto.WebResponse[dto.StatusResponse]{
		Status: 200,
		Data:   *response,
	})
}

func (c *StatusController) Update(ctx *fiber.Ctx) error {
	// like java, everything all is string
	id, _ := strconv.Atoi(ctx.Params("id"))

	request := new(dto.StatusRequestUpdate)
	fmt.Println("check request", request)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse body : %+v", err)
	}
	fmt.Println("sini1")
	response, err := c.UseCase.Update(ctx.UserContext(), &id, request)
	if err != nil {
		c.Log.Warnf("Failed to update status : %v", err)
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
	fmt.Println("sini2", response)

	return ctx.JSON(dto.WebResponse[dto.StatusResponse]{
		Status: 200,
		Data:   *response,
	})

}

func (c *StatusController) Detail(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	response, err := c.UseCase.Detail(ctx.UserContext(), uint(id))
	if err != nil {
		c.Log.Warnf("Failed to detail status : %v", err)
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

	return ctx.JSON(dto.WebResponse[dto.StatusResponse]{
		Status: 200,
		Data:   *response,
	})
}

func (c *StatusController) Delete(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	err := c.UseCase.Delete(ctx.UserContext(), uint(id))
	if err != nil {
		c.Log.Warnf("Failed to delete status : %v", err)
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

	return ctx.JSON(dto.WebResponse[bool]{
		Status: 200,
		Data:   true,
	})
}
