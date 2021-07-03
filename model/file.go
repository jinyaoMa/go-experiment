package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

const (
	FILE_TYPE_DIRECTORY = "directory"
	FILE_TYPE_FILE      = "file"
)

const (
	FILE_ERROR_IGNORE_REPEAT = "ignore repeated file"
)

type File struct {
	gorm.Model
	Path           string `gorm:"not null;index"`
	Type           string `gorm:"check:type IN ('directory', 'file');not null"`
	Extension      string
	Size           uint64
	Recycled       uint `gorm:"check:recycled IN (0, 1);default:0"`
	ShareCode      *string
	ShareExpiredAt *time.Time
	UserID         uint
}

func (file *File) BeforeCreate(tx *gorm.DB) error {
	var repeatedFile File
	result := tx.First(&repeatedFile, "path = ?", file.Path)
	if result.RowsAffected == 1 {
		return errors.New(FILE_ERROR_IGNORE_REPEAT)
	}
	return nil
}
