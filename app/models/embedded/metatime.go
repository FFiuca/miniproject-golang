package embedded

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type MetaTimeFull struct {
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt sql.NullTime   `gorm:"column:created_at;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

type MetaTime struct {
	CreatedAt time.Time    `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt sql.NullTime `gorm:"column:updated_at;autoUpdateTime"`
}
