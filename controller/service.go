package controller

import (
	"jinyaoma/go-experiment/config"
	"jinyaoma/go-experiment/model"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

type DownloadQuery struct {
	IndexPointToToken       int    `form:"a"`
	RelativeTokenAfterIndex string `form:"b"`
	ShareCode               string `form:"c" binding:"required"`
	FileId                  uint   `form:"d" binding:"required"`
	UserAccount             string `form:"e"`
}

func download(c *gin.Context) {
	var query DownloadQuery
	err := c.ShouldBind(&query)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	if len(query.RelativeTokenAfterIndex) < 64 {
		c.Status(http.StatusNotFound)
		return
	}

	var file model.File
	resultFile := model.GetDB().Where("id = ? AND files.share_code = ?", query.FileId, query.ShareCode).First(&file)
	if resultFile.Error != nil {
		c.Status(http.StatusNotFound)
		return
	}

	if file.ShareExpiredAt.After(time.Now()) {
		c.JSON(200, gin.H{
			"path": filepath.Join(config.WORKSPACE, query.UserAccount, file.Path),
		})
		return
	}

	var user model.User
	resultUser := model.GetDB().Where("id = ? AND SUBSTR(users.token, ?) = ?", file.UserID, query.IndexPointToToken, query.RelativeTokenAfterIndex).First(&user)
	if resultUser.Error != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(200, gin.H{
		"path": filepath.Join(config.WORKSPACE, user.Account, file.Path),
	})

	/*
		//打开文件
		_, err := os.Open(fileDir + "/" + fileName)
		//非空处理
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
			    "success": false,
			    "message": "失败",
			    "error":   "资源不存在",
			})
			c.Redirect(http.StatusFound, "/404")
			return
		}
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename="+fileName)
		c.Header("Content-Transfer-Encoding", "binary")
		c.File(fileDir + "/" + fileName)
	*/
}

func InitService(rg *gin.RouterGroup) {
	api := rg.Group("/service")
	{
		api.GET("/download", download)
	}
}
