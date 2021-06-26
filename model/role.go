package model

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

const (
	ROLE_GUEST  = "guest"
	ROLE_MEMBER = "member"
	ROLE_VIP    = "vip"
	ROLE_ADMIN  = "admin"
)

type Role struct {
	gorm.Model
	Name       string
	Permission string
}

func createDefaultRoles(db *gorm.DB) {
	roles := []Role{
		{
			Name:       ROLE_ADMIN,
			Permission: "*",
		}, {
			Name: ROLE_VIP,
		}, {
			Name: ROLE_MEMBER,
		}, {
			Name: ROLE_MEMBER,
		},
	}

	db.Create(&roles)

	fmt.Println("Default roles created")
}

func (role *Role) BeforeCreate(tx *gorm.DB) error {
	var repeat Role
	result := tx.First(&repeat, "name = ?", role.Name)
	if result.RowsAffected > 0 {
		return errors.New("Ignore repeated role")
	}
	return nil
}
