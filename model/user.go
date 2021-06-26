package model

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

const (
	USER_ERROR_IGNORE_REPEAT = "ignore repeated user"
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

func (user *User) BeforeCreate(tx *gorm.DB) error {
	var repeatedUser User
	result := tx.First(&repeatedUser, "account = ?", user.Account)
	if result.RowsAffected > 0 {
		return errors.New(USER_ERROR_IGNORE_REPEAT)
	}
	return nil
}

func GetUsers() []User {
	var users []User
	db.Find(&users)
	return users
}
