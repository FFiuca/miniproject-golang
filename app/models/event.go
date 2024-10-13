package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model

	UserID          uint           `gorm:"column:user_id"`
	StatusID        uint           `gorm:"column:status_id"`
	Title           string         `gorm:"column:title;"`
	Description     sql.NullString `gorm:"column:description;"`
	Location        sql.NullString `gorm:"column:location;"`
	Status          Status         // using default convention like laravel the rule is
	User            User
	EventReminder   []EventReminder
	EventAttachment []EventAttachment
	EventUser       []User `gorm:"many2many:event_user_pivot;foreignKey:id;joinForeignKey:event_id;references:id;joinReferences:user_id"`
}
