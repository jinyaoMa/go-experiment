package controller

import (
	"errors"
	"jinyaoma/go-experiment/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

const (
	CONTROLLER_FILES_ERROR_BIND           = "binding error"
	CONTROLLER_FILES_ERROR_USER_NOT_FOUND = "id token unmatched"
)

type FilesForm struct {
	ID    string `form:"id" binding:"required"`
	Token string `form:"token" binding:"required"`
}

func getFiles(c *gin.Context) {
	var form FilesForm
	err := c.ShouldBindWith(&form, binding.FormPost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":        CONTROLLER_FILES_ERROR_BIND,
			"isLogoutFail": true,
		})
		return
	}

	var user model.User
	resultFindUser := model.GetDB().First(&user, "id = ? AND token = ?", form.ID, form.Token)
	if errors.Is(resultFindUser.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"fail": CONTROLLER_FILES_ERROR_USER_NOT_FOUND,
		})
		return
	} else if resultFindUser.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	var files []model.File
	resultFiles := model.GetDB().Find(&files, "user_id = ?", user.ID)
	if resultFiles.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"files": files,
		},
	})
}

func InitFiles(rg *gin.RouterGroup) {
	rg.POST("/files", getFiles)
}
