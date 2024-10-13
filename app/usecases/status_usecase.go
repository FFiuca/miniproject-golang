package usecases

import (
	"context"
	"fmt"
	"project1/app/dto"
	"project1/app/models"
	"project1/app/repositories"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type StatusUseCaseInterface interface {
	Create(ctx context.Context, data *dto.StatusRequestCreate) (*dto.StatusResponse, error)
	Update(ctx context.Context, id *int, data *dto.StatusRequestUpdate) (*dto.StatusResponse, error)
	Detail(ctx context.Context, id uint) (*dto.StatusResponse, error)
	Delete(ctx context.Context, id uint) error
}

type StatusUseCase struct {
	DB         *gorm.DB
	Log        *logrus.Logger
	Validate   *validator.Validate
	Repository repositories.StatusRepository
}

func NewStatusUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, repo repositories.StatusRepository) StatusUseCaseInterface {
	return &StatusUseCase{
		DB:         db,
		Log:        log,
		Validate:   validate,
		Repository: repo,
	}
}

func (c *StatusUseCase) Create(ctx context.Context, data *dto.StatusRequestCreate) (*dto.StatusResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(data)
	fmt.Println("sini", err)
	if err != nil {
		c.Log.Warnf("Form validation error : %+v", err)
		return nil, &fiber.Error{
			Code:    400,
			Message: fmt.Sprintf("%+v", err),
		}
	}
	fmt.Println("sini2")

	status := &models.Status{
		Name: data.Name,
	}

	_, err = c.Repository.Add(status)
	if err != nil {
		c.Log.Warnf("Add error %s", err)
		return nil, err
	}

	if err = tx.Commit().Error; err != nil {
		c.Log.Warnf("Commit error %s", err)
		return nil, err
	}

	return dto.StatusToResponse(status), nil
}

func (c *StatusUseCase) Update(ctx context.Context, id *int, data *dto.StatusRequestUpdate) (*dto.StatusResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	data.ID = uint(*id)

	err := c.Validate.Struct(data)
	if err != nil {
		c.Log.Warnf("Form validation error: %+v", err)

		return nil, &fiber.Error{
			Code:    400,
			Message: fmt.Sprintf("%+v", err),
		}
	}

	status := &models.Status{
		Name: data.Name,
	}

	update, err := c.Repository.Update(id, status)
	if err != nil {
		c.Log.Warnf("UPdate status error: %+v", err)

		return nil, &fiber.Error{
			Code:    500,
			Message: fmt.Sprintf("%+v", err),
		}
	}
	fmt.Println("update", update)
	return dto.StatusToResponse(status), nil
}

func (c *StatusUseCase) Detail(ctx context.Context, id uint) (*dto.StatusResponse, error) {
	id_ := int(id)
	status := &models.Status{}
	_, err := c.Repository.Detail(&id_, status)
	if err != nil {
		c.Log.Warnf("Get status error %+v", err)

		return nil, &fiber.Error{
			Code:    500,
			Message: fmt.Sprintf("%+v", err),
		}
	}

	return dto.StatusToResponse(status), nil
}

func (c *StatusUseCase) Delete(ctx context.Context, id uint) error {
	id_ := int(id)
	status := &models.Status{}

	err := c.Repository.Delete(&id_, status)
	if err != nil {
		c.Log.Warnf("Delete status error %+v", err)

		return err
	}

	return nil
}
