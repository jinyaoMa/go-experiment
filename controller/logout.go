package controller

import (
	"errors"
	"jinyaoma/go-experiment/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

const (
	CONTROLLER_LOGOUT_ERROR_BIND           = "binding error"
	CONTROLLER_LOGOUT_ERROR_USER_NOT_FOUND = "id token unmatched"
)

type LogoutForm struct {
	ID    string `form:"id" binding:"required"`
	Token string `form:"token" binding:"required"`
}

func logout(c *gin.Context) {
	var form LogoutForm
	err := c.ShouldBindWith(&form, binding.FormPost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":        CONTROLLER_LOGOUT_ERROR_BIND,
			"isLogoutFail": true,
		})
		return
	}

	var user model.User
	resultFindUser := model.GetDB().First(&user, "id = ? AND token = ?", form.ID, form.Token)
	if errors.Is(resultFindUser.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"fail": CONTROLLER_LOGOUT_ERROR_USER_NOT_FOUND,
		})
		return
	} else if resultFindUser.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	time := time.Now()
	resultUpdateTokenExpiredAt := model.GetDB().Model(&user).Updates(model.User{
		TokenExpiredAt: &time,
	})
	if resultUpdateTokenExpiredAt.RowsAffected != 1 {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func InitLogout(rg *gin.RouterGroup) {
	rg.POST("/logout", logout)
}
