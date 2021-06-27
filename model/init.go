package model

import (
	"path"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	DB_FILENAME = "database.db"
)

var (
	db *gorm.DB
)

func InitDB(workSpace string) *gorm.DB {
	var err error
	db, err = gorm.Open(sqlite.Open(path.Join(workSpace, DB_FILENAME)), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{}, &Role{}, &File{})
	createDefaultRoles(db)
	createUserAdmin(db)
	return db
}

func GetDB() *gorm.DB {
	return db
}
