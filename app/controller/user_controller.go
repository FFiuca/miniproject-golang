package controller

import (
	"project1/app/dto"
	"project1/app/usecases"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log     *logrus.Logger
	UseCase usecases.UserUseCaseInterface
}

func NewUserController(log *logrus.Logger, useCase usecases.UserUseCaseInterface) *UserController {
	return &UserController{
		Log:     log,
		UseCase: useCase,
	}
}

func (c *UserController) Create(ctx *fiber.Ctx) error {
	request := new(dto.UserRequestCreate)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse body %v", err)
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create user : %v", err)
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

	return ctx.JSON(dto.WebResponse[*dto.UserResponse]{
		Status: 200,
		Data:   response,
	})
}

func (c *UserController) Update(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	request := new(dto.UserRequestUpdate)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.Warnf("Failed to parse body %v", err)
	}

	response, err := c.UseCase.Update(ctx.UserContext(), &id, request)
	if err != nil {
		c.Log.Warnf("Failed to update user : %v", err)
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

	return ctx.JSON(dto.WebResponse[*dto.UserResponse]{
		Status: 200,
		Data:   response,
	})
}

func (c *UserController) Detail(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	response, err := c.UseCase.Detail(ctx.UserContext(), uint(id))
	if err != nil {
		c.Log.Warnf("Failed to detail user : %v", err)
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

	return ctx.JSON(dto.WebResponse[*dto.UserResponse]{
		Status: 200,
		Data:   response,
	})
}

func (c *UserController) Delete(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	err := c.UseCase.Delete(ctx.UserContext(), uint(id))
	if err != nil {
		c.Log.Warnf("Failed to detail user : %v", err)
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

func (c *UserController) Search(ctx *fiber.Ctx) error {
	request := new(dto.UserRequestSearch)
	request.Params = map[string]any{
		"email":      ctx.Query("email", ""),
		"created_at": ctx.Query("created_at", ""),
	}
	// fmt.Println(ctx.Queries())
	data, err := c.UseCase.Search(ctx.UserContext(), request.Params)
	if err != nil {
		c.Log.Warnf("Failed to detail user : %v", err)
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

	return ctx.JSON(dto.WebResponse[[]dto.UserResponse]{
		Status: 200,
		Data:   data,
	})
}
