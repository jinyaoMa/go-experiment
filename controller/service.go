package controller

import (
	"errors"
	"jinyaoma/go-experiment/config"
	"jinyaoma/go-experiment/model"
	"jinyaoma/go-experiment/workspace"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
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

type UploadForm struct {
	ID    uint                  `form:"id" binding:"required"`
	Token string                `form:"token" binding:"required"`
	File  *multipart.FileHeader `form:"file" binding:"required"`
	DesId uint                  `form:"desId" binding:"required"`
}

func upload(c *gin.Context) {
	var form UploadForm
	err := c.ShouldBindWith(&form, binding.FormMultipart)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":  err.Error(),
			"isFail": true,
		})
		return
	}

	var user model.User
	resultFindUser := model.GetDB().First(&user, "id = ? AND token = ? AND token_expired_at > ?", form.ID, form.Token, time.Now())
	if errors.Is(resultFindUser.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"fail": CONTROLLER_AUTH_ERROR_USER_NOT_FOUND,
		})
		return
	} else if resultFindUser.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	var role model.Role
	resultRole := model.GetDB().First(&role, user.RoleID)
	if resultRole.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	disk := workspace.GetDisk()
	var roleSpace uint64
	if role.Name == model.ROLE_ADMIN {
		roleSpace = disk.Available
	} else if role.Space < disk.Available {
		roleSpace = role.Space
	} else {
		roleSpace = disk.Available
	}
	if uint64(form.File.Size) > roleSpace {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	var des model.File
	resultDes := model.GetDB().Where("id = ? AND user_id = ?", form.DesId, user.ID).First(&des)
	if resultDes.Error != nil || des.Type != model.FILE_TYPE_DIRECTORY {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	// Upload the file to specific dst.
	userWorkspace := filepath.Join(config.WORKSPACE, user.Account)
	desPath := filepath.Join(userWorkspace, des.Path)
	filename := filepath.Base(form.File.Filename)
	savePath := filepath.Join(desPath, filename)
	errSave := c.SaveUploadedFile(form.File, savePath)
	if errSave != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	relativeSavePath, errRel := filepath.Rel(userWorkspace, savePath)
	if errRel != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	shareCode := config.GenerateShareCode(4)
	resultNewFile := model.GetDB().Create(&model.File{
		Name:      filename,
		Path:      relativeSavePath,
		Type:      model.FILE_TYPE_FILE,
		Extension: filepath.Ext(relativeSavePath),
		Size:      uint64(form.File.Size),
		ShareCode: shareCode,
		UserID:    user.ID,
	})
	if resultNewFile.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	var files []model.File
	resultFiles := model.GetDB().Find(&files, "user_id = ? AND recycled = 0", user.ID)
	if resultFiles.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	var usedSpace uint64
	resultUsedSpace := model.GetDB().Raw("select sum(size) from files where user_id = ?", user.ID).Scan(&usedSpace)
	if resultUsedSpace.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	var allSpace uint64
	disk.Update()
	if role.Name == model.ROLE_ADMIN {
		allSpace = disk.Available
	} else if role.Space < disk.Available {
		allSpace = role.Space
	} else {
		allSpace = disk.Available
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"files":     files,
			"usedSpace": usedSpace,
			"allSpace":  allSpace,
		},
	})
}

func InitService(rg *gin.RouterGroup) {
	api := rg.Group("/service")
	{
		api.GET("/download", download)
		api.POST("/upload", upload)
	}
}
