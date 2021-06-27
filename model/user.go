package model

import (
	"errors"
	"fmt"
	"jinyaoma/go-experiment/workspace"
	"os"

	"path/filepath"
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
	Files          []File
}

func initUserAdmin(db *gorm.DB) User {
	var admin User
	var role Role
	db.First(&admin, "account = ?", "admin")
	if admin.Account != "admin" {
		db.First(&role, "name = ?", ROLE_ADMIN)
		if role.Name == ROLE_ADMIN {
			var userFiles []File
			err := workspace.InitUserWorkspace("/admin", func(path string, fileInfo os.FileInfo) {
				var typ string
				if fileInfo.IsDir() {
					typ = FILE_TYPE_DIRECTORY
				} else {
					typ = FILE_TYPE_FILE
				}
				userFiles = append(userFiles, File{
					Path:      path,
					Type:      typ,
					Extension: filepath.Ext(path),
					Size:      fileInfo.Size(),
				})
			})
			if err != nil {
				fmt.Println("User admin workspace error")
			} else {
				if db.Create(&User{
					Name:     "Admin",
					Account:  "admin",
					Password: "admin",
					RoleID:   role.ID,
					Files:    userFiles,
				}).Error == nil {
					fmt.Println("User admin initialized")
				}
			}
		}
	}
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
