package controller

import (
	"jinyaoma/go-experiment/config"
	"jinyaoma/go-experiment/model"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	CONTROLLER_FILES_ERROR_BIND     = "binding error"
	CONTROLLER_FILES_ERROR_FILENAME = "filename error"
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

type RenameFileForm struct {
	FileId      string `form:"fileId" binding:"required"`
	NewFilename string `form:"filename" binding:"required"`
}

func renameFile(c *gin.Context) {
	var form RenameFileForm
	err := c.ShouldBindWith(&form, binding.FormPost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":  CONTROLLER_FILES_ERROR_BIND,
			"isFail": true,
		})
		return
	}

	if strings.ContainsAny(form.NewFilename, "\\/:*?\"<>|") {
		c.JSON(http.StatusOK, gin.H{
			"fail": CONTROLLER_FILES_ERROR_FILENAME,
		})
		return
	}

	user := GetAuthUser(c)

	var file model.File
	resultFile := model.GetDB().Where("id = ? AND user_id = ?", form.FileId, user.ID).First(&file)
	if resultFile.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	oldPath := filepath.Join(config.WORKSPACE, user.Account, file.Path)
	currentFolder := filepath.Dir(oldPath)
	newPath := filepath.Join(currentFolder, form.NewFilename)

	errRename := os.Rename(oldPath, newPath)
	if errRename != nil {
		c.JSON(http.StatusOK, gin.H{
			"fail": CONTROLLER_FILES_ERROR_FILENAME,
		})
		return
	}

	relativeNewPath, errRel := filepath.Rel(filepath.Join(config.WORKSPACE, user.Account), newPath)
	if errRel != nil {
		c.JSON(http.StatusOK, gin.H{
			"fail": CONTROLLER_FILES_ERROR_FILENAME,
		})
		return
	}

	resultUpdateFile := model.GetDB().Model(&file).Updates(model.File{
		Name:      filepath.Base(newPath),
		Path:      relativeNewPath,
		Extension: filepath.Ext(newPath),
	})
	if resultUpdateFile.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func InitFiles(rg *gin.RouterGroup) {
	api := rg.Group("/files").Use(Auth())
	{
		api.POST(".", getFiles)
		api.POST("/renameFile", renameFile)
	}
}
