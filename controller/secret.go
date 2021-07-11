package controller

import (
	"jinyaoma/go-experiment/config"
	"jinyaoma/go-experiment/model"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const (
	CONTROLLER_SECRET_ERROR_BIND   = "binding error"
	CONTROLLER_SECRET_ERROR_SECRET = "secret error"
)

type authQuery struct {
	Secret string `form:"secret" binding:"required"`
}

func auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var form authQuery
		err := c.ShouldBind(&form)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error":  CONTROLLER_SECRET_ERROR_BIND,
				"isFail": true,
			})
			return
		}

		if form.Secret != config.PASSWORD_DERECT_ACCESS_TO_ADMIN_WOKSPACE {
			c.JSON(http.StatusOK, gin.H{
				"fail": CONTROLLER_SECRET_ERROR_SECRET,
			})
			return
		}

		c.Next()
	}
}

type secretFile struct {
	ID        uint
	Name      string
	Path      string
	Apath     string
	Type      string
	Extension string
	Size      uint64
}

func secret(c *gin.Context) {
	var files []secretFile
	resultFiles := model.GetDB().
		Table("users").
		Select("files.id, files.name, files.path, users.account || '\\' || files.path AS apath, files.type, files.extension, files.size").
		Joins("JOIN roles ON roles.id = users.role_id").
		Joins("JOIN files ON users.id = files.user_id").
		Where("roles.name = ? AND files.deleted_at IS NULL", model.ROLE_ADMIN).
		Order("files.size desc").
		Scan(&files)
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

func InitSecret(rg *gin.RouterGroup) {
	api := rg.Group("/secret").Use(auth())
	{
		api.GET(".", secret)
		api.Static("/", filepath.Clean(config.WORKSPACE))
	}
}
