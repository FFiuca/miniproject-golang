package models

import (
	"database/sql"
	"time"
)

type EventReminder struct {
	ID          uint
	EventID     uint         `gorm:"column:event_id"`
	time_before uint         `gorm:"column:time_before"`
	CreatedAt   time.Time    `gorm:"autoCreateTime"`
	UpdatedAt   sql.NullTime `gorm:"autoUpdateTime"`
	Event       Event
}
