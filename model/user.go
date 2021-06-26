package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string
	Account        string
	Password       string
	Token          *string
	TokenExpiredAt *time.Time
	RoleID         uint
	Role           Role
}

func createUserAdmin(db *gorm.DB) User {
	var admin User
	var role Role
	db.First(&admin, "account = ?", "admin")
	if admin.Account != "admin" {
		db.First(&role, "name = ?", ROLE_ADMIN)
		if role.Name == ROLE_ADMIN {
			db.Create(&User{
				Name:     "Admin",
				Account:  "admin",
				Password: "admin",
				RoleID:   role.ID,
			})
		}
	}

	fmt.Println("User admin created")
	return admin
}
