package models

import (
	"database/sql"
	"time"
)

type EventUserPivot struct {
	ID          uint
	UserID      uint         `gorm:"column:user_id"`
	EventID     uint         `gorm:"column:event_id"`
	CreatedAt   time.Time    `gorm:"autoCreateTime"`
	UpdatedAtAt sql.NullTime `gorm:"autoUpdateTime"`
	Event       Event
	User        User
}

// custom name
func (EventUserPivot) TableName() string {
	return "event_user_pivot"
}
