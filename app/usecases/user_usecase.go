package usecases

import (
	"context"
	"fmt"
	"project1/app/dto"
	"project1/app/helpers"
	"project1/app/models"
	"project1/app/repositories"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserUseCaseInterface interface {
	Create(ctx context.Context, data *dto.UserRequestCreate) (*dto.UserResponse, error)
	Update(ctx context.Context, id *int, data *dto.UserRequestUpdate) (*dto.UserResponse, error)
	Detail(ctx context.Context, id uint) (*dto.UserResponse, error)
	Delete(ctx context.Context, id uint) error
	Search(ctx context.Context, data map[string]any) ([]dto.UserResponse, error)
}

type UserUseCase struct {
	DB         *gorm.DB
	Log        *logrus.Logger
	Validate   *validator.Validate
	Repository repositories.UserRepository
}

func NewUserUseCase(db *gorm.DB, log *logrus.Logger, val *validator.Validate, repo repositories.UserRepository) UserUseCaseInterface {
	return &UserUseCase{
		DB:         db,
		Log:        log,
		Validate:   val,
		Repository: repo,
	}
}

func (c *UserUseCase) Create(ctx context.Context, request *dto.UserRequestCreate) (*dto.UserResponse, error) {
	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("form validation err %+v", err)

		return nil, &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("form validation error %v", err),
		}
	}

	hashed, _ := helpers.HashPassword(request.Password)
	var user *models.User = &models.User{
		Email:    request.Email,
		Password: hashed,
		// StatusID: 1,
		Status: models.Status{
			ID: 1,
		},
	}

	_, err = c.Repository.Add(user)
	if err != nil {
		c.Log.Warnf("Error create user %v", err)

		return nil, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: fmt.Sprintf("Error create user %v", err),
		}
	}

	return dto.UserToResponse(user), nil
}

func (c *UserUseCase) Update(ctx context.Context, id *int, request *dto.UserRequestUpdate) (*dto.UserResponse, error) {
	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("form validation err %+v", err)

		return nil, &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: fmt.Sprintf("form validation error %v", err),
		}
	}

	hashed, _ := helpers.HashPassword(request.Password)
	var user *models.User = &models.User{
		Email:    request.Email,
		Password: hashed,
	}

	_, err = c.Repository.Update(id, user)
	if err != nil {
		c.Log.Warnf("Error update user %v", err)

		return nil, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: fmt.Sprintf("Error update user %v", err),
		}
	}

	return dto.UserToResponse(user), nil
}

func (c *UserUseCase) Detail(ctx context.Context, id uint) (*dto.UserResponse, error) {
	user := &models.User{}
	var id_ int = int(id)
	_, err := c.Repository.Detail(&id_, user)
	if err != nil {
		c.Log.Warnf("Error get detail %+v", err)

		return nil, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: fmt.Sprintf("Error get detail %+v", err),
		}
	}
	fmt.Println("user", user)
	return dto.UserToResponse(user), nil
}

func (c *UserUseCase) Delete(ctx context.Context, id uint) error {
	var id_ int = int(id)
	var user *models.User = &models.User{}

	err := c.Repository.Delete(&id_, user)
	if err != nil {
		c.Log.Warnf("Error delete user %+v", err)

		return &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: fmt.Sprintf("Error delete user %+v", err),
		}
	}

	return nil
}

func (c *UserUseCase) Search(ctx context.Context, data map[string]any) ([]dto.UserResponse, error) {
	var user *models.User = &models.User{}

	result, err := c.Repository.Search(&data, user)
	if err != nil {
		c.Log.Warnf("Error search user %v", err)

		return nil, &fiber.Error{
			Code:    fiber.StatusInternalServerError,
			Message: fmt.Sprintf("Error search user %v", err),
		}
	}

	var response []dto.UserResponse = make([]dto.UserResponse, 0)
	for _, v := range *result {
		response = append(response, *dto.UserToResponse(&v))
	}

	return response, nil
}
