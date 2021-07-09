package model

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

const (
	ROLE_MEMBER = "member"
	ROLE_ADMIN  = "admin"
)

const (
	ROLE_ERROR_IGNORE_REPEAT = "ignore repeated role"
)

type Role struct {
	gorm.Model
	Name       string
	Permission string `gorm:"default:'CORE:1,SHARE:0,ADMIN:0'"`
	Space      uint64 `gorm:"default:0"`
}

func initDefaultRoles(db *gorm.DB) {
	roles := []Role{
		{
			Name:       ROLE_ADMIN,
			Permission: "CORE:1,ADMIN:1",
		}, {
			Name:       ROLE_MEMBER,
			Permission: "CORE:1,ADMIN:0",
			Space:      8 * 1024 * 1024 * 1024,
		},
	}

	if db.Create(&roles).Error == nil {
		log.Println("Default roles initialized")
	}
}

func (role *Role) BeforeCreate(tx *gorm.DB) error {
	var repeatedRole Role
	result := tx.First(&repeatedRole, "name = ?", role.Name)
	if result.RowsAffected == 1 {
		return errors.New(ROLE_ERROR_IGNORE_REPEAT)
	}
	return nil
}
