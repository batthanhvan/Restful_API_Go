package database

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string `gorm:"primarykey;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type DefaultSearchModel struct {
	Skip      int
	Limit     int
	OrderBy   string
	OrderType string
	Fields    []string
}
