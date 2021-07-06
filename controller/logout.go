package controller

import (
	"jinyaoma/go-experiment/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func logout(c *gin.Context) {
	user := GetAuthUser(c)

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
	rg.POST("/logout", Auth(), logout)
}
