package models

import (
	"database/sql"
	"time"
)

type EventAttachment struct {
	ID          uint
	EventID     uint           `gorm:"column:event_id"`
	path        sql.NullString `gorm:"column:path"`
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAtAt sql.NullTime   `gorm:"autoUpdateTime"`
	Event       Event
}
