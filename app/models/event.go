package models

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model

	UserID          uint           `gorm:"column:user_id" json:"user_id"`
	StatusID        uint           `gorm:"column:status_id" json:"status_id"`
	Title           string         `gorm:"column:title; json:"title"`
	Description     sql.NullString `gorm:"column:description;" json:"description"`
	Location        sql.NullString `gorm:"column:location;" json:"location"`
	Status          Status         `json:"status"` // using default convention like laravel the rule is
	EventDate       time.Time      `gorm:"column:event_date;"`
	User            User
	EventReminder   []EventReminder
	EventAttachment []EventAttachment
	EventUser       []User `gorm:"many2many:event_user_pivot;foreignKey:id;joinForeignKey:event_id;references:id;joinReferences:user_id"`
}
