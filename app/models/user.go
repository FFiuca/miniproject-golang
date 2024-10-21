package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	StatusID  uint `gorm:"column:status_id"`
	Email     string
	Password  string
	Status    Status  `gorm:"references:id;foreignKey:status_id"` // if different package can be cyclic import // belong to
	EventUser []Event `gorm:"many2many:event_user_pivot;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:event_id"`
	// Event     Event // cant't recursive
	// embedded.MetaTimeFull `gorm:"embedded"`
}
