package controller

import (
	"jinyaoma/go-experiment/config"
	"jinyaoma/go-experiment/model"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"unicode/utf8"

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

	var file model.File
	resultFile := model.GetDB().Where("id = ? AND share_code = ? AND type = ? AND recycled = 0", query.FileId, query.ShareCode, model.FILE_TYPE_FILE).First(&file)
	if resultFile.Error != nil {
		c.Status(http.StatusNotFound)
		return
	}

	var path string
	if file.ShareExpiredAt.After(time.Now()) && utf8.RuneCountInString(query.UserAccount) > 0 {
		var user model.User
		resultUser := model.GetDB().Where("id = ? AND account = ?", file.UserID, query.UserAccount).First(&user)
		if resultUser.Error != nil {
			c.Status(http.StatusNotFound)
			return
		}
		path = filepath.Join(config.WORKSPACE, query.UserAccount, file.Path)
	} else {
		if utf8.RuneCountInString(query.RelativeTokenAfterIndex) < 64 {
			c.Status(http.StatusNotFound)
			return
		}
		var user model.User
		resultUser := model.GetDB().Where("id = ? AND SUBSTR(users.token, ?) = ?", file.UserID, query.IndexPointToToken, query.RelativeTokenAfterIndex).First(&user)
		if resultUser.Error != nil {
			c.Status(http.StatusNotFound)
			return
		}
		path = filepath.Join(config.WORKSPACE, user.Account, file.Path)
	}

	_, errOpenFile := os.Open(path)
	if errOpenFile != nil {
		model.GetDB().Delete(&file)
		c.Status(http.StatusNotFound)
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+filepath.Base(path))
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(path)
}

func InitService(rg *gin.RouterGroup) {
	api := rg.Group("/service")
	{
		api.GET("/download", download)
	}
}
