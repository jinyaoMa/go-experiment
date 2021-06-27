package model

import (
	"time"

	"gorm.io/gorm"
)

const (
	FILE_TYPE_DIRECTORY = "directory"
	FILE_TYPE_FILE      = "file"
)

type File struct {
	gorm.Model
	Path           string `gorm:"not null"`
	Type           string `gorm:"check:type IN ('directory', 'file');not null"`
	Recycled       uint   `gorm:"check:recycled IN (0, 1);default:0"`
	ShareCode      *string
	ShareExpiredAt *time.Time
	UserID         uint
}
