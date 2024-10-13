package models

import (
	"project1/app/models/embedded"
)

type Status struct {
	ID                    uint   `gorm:"primaryKey;column:id"`
	Name                  string `gorm:"column:name"`
	embedded.MetaTimeFull `gorm:"embedded"`
	Users                 []User `gorm:"foreignKey:status_id;references:id"` // has many
}

func (s *Status) TableName() string {
	return "master_statuses"
}
