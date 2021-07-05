package controller

import (
	"jinyaoma/go-experiment/config"
	"jinyaoma/go-experiment/model"
	"jinyaoma/go-experiment/workspace"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	CONTROLLER_SIGNUP_ERROR_BIND          = "binding error"
	CONTROLLER_SIGNUP_ERROR_ACCOUNT_EXIST = "account exist"
)

type SignupForm struct {
	Username string `form:"username" binding:"required"`
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func signup(c *gin.Context) {
	var form SignupForm
	err := c.ShouldBindWith(&form, binding.FormPost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":        CONTROLLER_LOGOUT_ERROR_BIND,
			"isSignupFail": true,
		})
		return
	}

	var role model.Role
	resultFindRole := model.GetDB().First(&role, "name = ?", config.DEFAULT_SIGNUP_ROLE_NAME)
	if resultFindRole.RowsAffected != 1 {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	var userFiles []model.File
	errNewWorkspace := workspace.NewUserWorkspace(form.Account, func(path string, fileInfo os.FileInfo) {
		var typ string
		if fileInfo.IsDir() {
			typ = model.FILE_TYPE_DIRECTORY
		} else {
			typ = model.FILE_TYPE_FILE
		}
		userFiles = append(userFiles, model.File{
			Path:      path,
			Type:      typ,
			Extension: filepath.Ext(path),
			Size:      uint64(fileInfo.Size()),
		})
	})
	if errNewWorkspace != nil {
		c.JSON(http.StatusOK, gin.H{
			"fail": CONTROLLER_SIGNUP_ERROR_ACCOUNT_EXIST,
		})
		return
	}

	token, expiredAt := config.GenerateToken(128)
	newUser := model.User{
		Name:           form.Username,
		Account:        form.Account,
		Password:       form.Password,
		Token:          &token,
		TokenExpiredAt: &expiredAt,
		RoleID:         role.ID,
		Files:          userFiles,
	}
	resultCreateNewUser := model.GetDB().Create(&newUser)
	if resultCreateNewUser.RowsAffected != 1 {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	var allSpace uint64
	disk := workspace.GetDisk()
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
			"userid":    newUser.ID,
			"username":  newUser.Name,
			"role":      role.Name,
			"token":     newUser.Token,
			"usedSpace": 0,
			"allSpace":  allSpace,
			"files":     newUser.Files,
		},
	})
}

func InitSignup(rg *gin.RouterGroup) {
	rg.POST("/signup", signup)
}
