package model

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

const (
	ROLE_MEMBER = "member"
	ROLE_VIP    = "vip"
	ROLE_ADMIN  = "admin"
)

const (
	ROLE_ERROR_IGNORE_REPEAT = "ignore repeated role"
)

type Role struct {
	gorm.Model
	Name       string
	Permission string
	Space      int `gorm:"default:0"`
}

func initDefaultRoles(db *gorm.DB) {
	roles := []Role{
		{
			Name:       ROLE_ADMIN,
			Permission: "*",
			Space:      -1,
		}, {
			Name: ROLE_VIP,
		}, {
			Name: ROLE_MEMBER,
		},
	}

	if db.Create(&roles).Error == nil {
		fmt.Println("Default roles initialized")
	}
}

func (role *Role) BeforeCreate(tx *gorm.DB) error {
	var repeatedRole Role
	result := tx.First(&repeatedRole, "name = ?", role.Name)
	if result.RowsAffected > 0 {
		return errors.New(ROLE_ERROR_IGNORE_REPEAT)
	}
	return nil
}
