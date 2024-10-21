package usecases

import (
	"context"
	"project1/app/dto"
	"project1/app/services"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthUseCaseInterface interface {
	Login(ctx context.Context, request *dto.AuthRequestLogin) (*dto.AuthResponse, error)
}

type AuthUseCase struct {
	DB       *gorm.DB
	Log      *logrus.Logger
	Validate *validator.Validate
	Service  services.AuthService
}

func (c *AuthUseCase) Login(ctx context.Context, request *dto.AuthRequestLogin) (*dto.AuthResponse, error) {
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Validate error %v", err)

		return nil, &fiber.Error{
			Code:    200,
			Message: err.Error(),
		}
	}

	login, err := c.Service.Login(request.Email, request.Password)
	if err != nil {
		c.Log.Warnf("error login %v", err)

		return nil, err
	}

	token, err := c.Service.CreateToken(login)
	if err != nil {
		c.Log.Warnf("error create token %v", err)

		return nil, err
	}

	return dto.AuthToResponse(token), nil
}

func NewAuthUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, svc services.AuthService) AuthUseCaseInterface {
	return &AuthUseCase{
		DB:       db,
		Log:      log,
		Validate: validate,
		Service:  svc,
	}
}
