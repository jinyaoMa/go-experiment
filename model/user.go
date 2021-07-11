package model

import (
	"errors"
	"jinyaoma/go-experiment/config"
	"jinyaoma/go-experiment/workspace"
	"log"
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
	resultAdmin := db.First(&admin, "account = ?", "admin")
	if errors.Is(resultAdmin.Error, gorm.ErrRecordNotFound) {
		resultRole := db.First(&role, "name = ?", ROLE_ADMIN)
		if resultRole.RowsAffected == 1 {
			user := User{
				Name:     "Admin",
				Account:  "admin",
				Password: "admin",
				RoleID:   role.ID,
			}
			resultUser := db.Create(&user)
			if resultUser.Error == nil {
				log.Println("User admin initialized")
			}
			var userFiles []File
			err := workspace.InitUserWorkspace("admin", func(path string, fileInfo os.FileInfo) {
				var typ string
				if fileInfo.IsDir() {
					typ = FILE_TYPE_DIRECTORY
				} else {
					typ = FILE_TYPE_FILE
				}
				shareCode := config.GenerateShareCode(4)
				userFiles = append(userFiles, File{
					Name:      filepath.Base(path),
					Path:      path,
					Type:      typ,
					Extension: filepath.Ext(path),
					Size:      uint64(fileInfo.Size()),
					ShareCode: shareCode,
					UserID:    user.ID,
				})
			})
			if err != nil {
				log.Println("User admin workspace error")
			}
			resultUserFiles := db.CreateInBatches(userFiles, 1000)
			if resultUserFiles.Error != nil {
				log.Println("User files init batches error")
			}
		}
	}
	return admin
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	var repeatedUser User
	result := tx.First(&repeatedUser, "account = ?", user.Account)
	if result.RowsAffected == 1 {
		return errors.New(USER_ERROR_IGNORE_REPEAT)
	}
	return nil
}

func GetUsers() []User {
	var users []User
	db.Find(&users)
	return users
}
