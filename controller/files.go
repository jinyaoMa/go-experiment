package controller

import (
	"jinyaoma/go-experiment/config"
	"jinyaoma/go-experiment/model"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	"unicode/utf8"

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
	FileId      uint   `form:"fileId" binding:"required"`
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
	resultFile := model.GetDB().Where("id = ? AND user_id = ? AND recycled = 0", form.FileId, user.ID).First(&file)
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

type MoveFileForm struct {
	SrcId uint `form:"srcId" binding:"required"`
	DesId uint `form:"desId" binding:"required"`
}

func moveFile(c *gin.Context) {
	var form MoveFileForm
	err := c.ShouldBindWith(&form, binding.FormPost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":  CONTROLLER_FILES_ERROR_BIND,
			"isFail": true,
		})
		return
	}

	user := GetAuthUser(c)

	var des model.File
	resultDes := model.GetDB().Where("id = ? AND user_id = ? AND recycled = 0", form.DesId, user.ID).First(&des)
	if resultDes.Error != nil || des.Type != model.FILE_TYPE_DIRECTORY {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	var src model.File
	resultSrc := model.GetDB().Where("id = ? AND user_id = ? AND recycled = 0", form.SrcId, user.ID).First(&src)
	if resultSrc.Error != nil || src.Type != model.FILE_TYPE_FILE {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	userWorkspace := filepath.Join(config.WORKSPACE, user.Account)
	oldPath := filepath.Join(userWorkspace, src.Path)
	newPath := filepath.Join(userWorkspace, des.Path, filepath.Base(oldPath))
	errMove := os.Rename(oldPath, newPath)
	if errMove != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}
	relativeNewPath, errRel := filepath.Rel(userWorkspace, newPath)
	if errRel != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}
	resultUpdateSrc := model.GetDB().Model(&src).Updates(model.File{
		Path: relativeNewPath,
	})
	if resultUpdateSrc.Error != nil {
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

type ShareFileForm struct {
	FileId uint `form:"fileId" binding:"required"`
}

func shareFile(c *gin.Context) {
	var form ShareFileForm
	err := c.ShouldBindWith(&form, binding.FormPost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":  CONTROLLER_FILES_ERROR_BIND,
			"isFail": true,
		})
		return
	}

	user := GetAuthUser(c)

	shareCode := config.GenerateShareCode(4)
	ShareExpiredAt := time.Now().AddDate(0, 1, 0)
	resultUpdateShare := model.GetDB().Model(&model.File{}).Where("id = ? AND user_id = ? AND recycled = 0", form.FileId, user.ID).Updates(model.File{
		ShareCode:      shareCode,
		ShareExpiredAt: ShareExpiredAt,
	})
	if resultUpdateShare.Error != nil {
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
			"c":     shareCode,
			"d":     form.FileId,
			"e":     user.Account,
		},
	})
}

type NewFolderForm struct {
	FolderName string `form:"folderName" binding:"required"`
	DesId      uint   `form:"desId" binding:"required"`
}

func newFolder(c *gin.Context) {
	var form NewFolderForm
	err := c.ShouldBindWith(&form, binding.FormPost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":  CONTROLLER_FILES_ERROR_BIND,
			"isFail": true,
		})
		return
	}

	if strings.ContainsAny(form.FolderName, "\\/:*?\"<>|") {
		c.JSON(http.StatusOK, gin.H{
			"fail": CONTROLLER_FILES_ERROR_FILENAME,
		})
		return
	}

	user := GetAuthUser(c)

	var des model.File
	resultDes := model.GetDB().Where("id = ? AND user_id = ? AND recycled = 0", form.DesId, user.ID).First(&des)
	if resultDes.Error != nil || des.Type != model.FILE_TYPE_DIRECTORY {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	userWorkspace := filepath.Join(config.WORKSPACE, user.Account)
	desPath := filepath.Join(userWorkspace, des.Path)
	newFolderPath := filepath.Join(desPath, form.FolderName)

	errMkdir := os.Mkdir(newFolderPath, os.ModeDir)
	if errMkdir != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	relativeNewFolderPath, errRel := filepath.Rel(userWorkspace, newFolderPath)
	if errRel != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	shareCode := config.GenerateShareCode(4)
	resultNewFolder := model.GetDB().Create(&model.File{
		Name:      form.FolderName,
		Path:      relativeNewFolderPath,
		Type:      model.FILE_TYPE_DIRECTORY,
		Extension: filepath.Ext(relativeNewFolderPath),
		Size:      0,
		ShareCode: shareCode,
		UserID:    user.ID,
	})
	if resultNewFolder.Error != nil {
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

type RecycleForm struct {
	FileId uint `form:"fileId" binding:"required"`
}

func recycle(c *gin.Context) {
	var form RecycleForm
	err := c.ShouldBindWith(&form, binding.FormPost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":  CONTROLLER_FILES_ERROR_BIND,
			"isFail": true,
		})
		return
	}

	user := GetAuthUser(c)

	var file model.File
	resultFile := model.GetDB().Where("id = ? AND user_id = ? AND recycled = 0", form.FileId, user.ID).First(&file)
	if resultFile.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	resultUpdateRecycled := model.GetDB().Model(&file).Updates(model.File{
		Recycled: 1,
	})
	if resultUpdateRecycled.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	if file.Type == model.FILE_TYPE_DIRECTORY {
		recycledFilePath := file.Path + "\\"
		recycledFilePathLength := utf8.RuneCountInString(recycledFilePath)
		resultRecycleInner := model.GetDB().Model(&model.File{}).Where("user_id = ? AND SUBSTR(path,1,?) = ?", user.ID, recycledFilePathLength, recycledFilePath).Updates(model.File{
			Recycled: 1,
		})
		if resultRecycleInner.Error != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": http.StatusInternalServerError,
			})
			return
		}
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
	api := rg.Group("/files").Use(Auth())
	{
		api.POST(".", getFiles)
		api.POST("/renameFile", renameFile)
		api.POST("/moveFile", moveFile)
		api.POST("/shareFile", shareFile)
		api.POST("/newFolder", newFolder)
		api.POST("/recycle", recycle)
	}
}
