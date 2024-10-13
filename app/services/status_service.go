package services

import (
	"fmt"
	"project1/app/models"
	"project1/app/repositories"

	"gorm.io/gorm"
)

type StatusService struct {
	DB *gorm.DB
}

func (svc *StatusService) Add(status *models.Status) (any, error) {
	add := svc.DB.Create(status)
	if add.Error != nil {
		return nil, add.Error
	}

	return add, nil
}

func (svc *StatusService) Update(id *int, status *models.Status) (any, error) {
	fmt.Println("masuk update")
	update := svc.DB.Model(status).Where("id = ?", id).Updates(status)
	if update.Error != nil {
		return nil, update.Error
	}

	return update, nil
}

func (svc *StatusService) Delete(id *int, status *models.Status) error {
	return svc.DB.Where("id = ?", id).Delete(status).Error
}

func (svc *StatusService) Detail(id *int, status *models.Status) (*gorm.DB, error) {
	data := svc.DB.Where("id = ?", id).First(status)
	if data.Error != nil {
		return nil, data.Error
	}

	return data, nil
}

func NewStatusService(db *gorm.DB) repositories.StatusRepository {
	return &StatusService{
		DB: db,
	}
}
