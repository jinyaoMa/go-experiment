package controller

import (
	"jinyaoma/go-experiment/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getFiles(c *gin.Context) {
	user := GetAuthUser(c)

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
	rg.POST("/files", Auth(), getFiles)
}
