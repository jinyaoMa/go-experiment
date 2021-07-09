package controller

import (
	"errors"
	"jinyaoma/go-experiment/config"
	"jinyaoma/go-experiment/model"
	"jinyaoma/go-experiment/workspace"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

const (
	CONTROLLER_LOGIN_ERROR_BIND           = "binding error"
	CONTROLLER_LOGIN_ERROR_USER_NOT_FOUND = "user password unmatched"
)

type LoginForm struct {
	Account  string `form:"account" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func login(c *gin.Context) {
	var form LoginForm
	err := c.ShouldBindWith(&form, binding.FormPost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":       CONTROLLER_LOGIN_ERROR_BIND,
			"isLoginFail": true,
		})
		return
	}

	var user model.User
	resultFindUser := model.GetDB().First(&user, "account = ? AND password = ?", form.Account, form.Password)
	if errors.Is(resultFindUser.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"fail": CONTROLLER_LOGIN_ERROR_USER_NOT_FOUND,
		})
		return
	} else if resultFindUser.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	var role model.Role
	resultFindRole := model.GetDB().First(&role, user.RoleID)
	if resultFindRole.RowsAffected != 1 {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	token, expiredAt := config.GenerateToken(128)
	resultUpdateToken := model.GetDB().Model(&user).Updates(model.User{
		Token:          &token,
		TokenExpiredAt: &expiredAt,
	})
	if resultUpdateToken.RowsAffected != 1 {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	var usedSpace uint64
	resultUsedSpace := model.GetDB().Raw("select sum(size) from files where user_id = ? AND deleted_at IS NULL", user.ID).Scan(&usedSpace)
	if resultUsedSpace.Error != nil {
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
			"userid":     user.ID,
			"username":   user.Name,
			"role":       role.Name,
			"permission": role.Permission,
			"token":      user.Token,
			"usedSpace":  usedSpace,
			"allSpace":   allSpace,
		},
	})
}

func InitLogin(rg *gin.RouterGroup) {
	rg.POST("/login", login)
}
