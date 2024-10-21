package services

import (
	"project1/app/models"
	"project1/app/repositories"

	"gorm.io/gorm"
)

type EventService struct {
	DB *gorm.DB
}

func (c *EventService) Add(data *models.Event) (any, error) {
	add := c.DB.Create(data)
	if add.Error != nil {
		return nil, add.Error
	}

	return add, nil
}

func (c *EventService) Update(id *int, data *models.Event) (any, error) {
	update := c.DB.Model(&models.Event{}).Where("id = ?", *id).Updates(*data)
	if update.Error != nil {
		return nil, update.Error
	}

	return update, nil

}

// harusnya return models.Event
func (c *EventService) Detail(id *int, event *models.Event) (*gorm.DB, error) {
	data := c.DB.Model(&event).Preload("Status").Preload("EventReminder").Preload("EventUser").Preload("User").Where("id = ?", *id).First(&event)
	if data.Error != nil {
		return nil, data.Error
	}

	return data, nil
}

func (c *EventService) Delete(id *int, event *models.Event) error {
	delete := c.DB.Where("id = ?", *id).Delete(event)
	if delete.Error != nil {
		return delete.Error
	}

	return nil
}

func (c *EventService) Search(param *map[string]any, event *models.Event) (*[]models.Event, error) {
	// start_date, end_date, user_id, title
	db := c.DB.Model(event)
	param2 := *param

	if param2["start_date"] != nil && param2["start_date"] != "" {
		db.Where("DATE(event_date)>=?", param2["start_date"])
	}
	if param2["end_date"] != nil && param2["end_date"] != "" {
		db.Where("DATE(event_date)<=?", param2["end_date"])
	}
	if param2["user_id"] != 0 {
		db.Where("user_id = ?", param2["user_id"])
	}
	if param2["title"] != "" {
		db.Where("title like %?%", param2["title"])
	}

	result := []models.Event{}
	get := db.Find(&result)
	if get.Error != nil {
		return nil, get.Error
	}

	return &result, nil
}

func NewEventService(db *gorm.DB) repositories.EventRepository {
	return &EventService{DB: db}
}
