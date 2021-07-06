package controller

import (
	"jinyaoma/go-experiment/config"
	"jinyaoma/go-experiment/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	CONTROLLER_ADMIN_ERROR_BIND         = "binding error"
	CONTROLLER_ADMIN_ERROR_UNAUTHORIZED = "unauthorized user"
	CONTROLLER_ADMIN_ERROR_PASSWORD     = "password unmatched"
)

func Admin() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := GetAuthUser(c)

		var role model.Role
		resultFindRole := model.GetDB().First(&role, user.RoleID)
		if resultFindRole.RowsAffected != 1 {
			c.JSON(http.StatusOK, gin.H{
				"error": http.StatusInternalServerError,
			})
			return
		}

		isAdmin := strings.Contains(role.Permission, "ADMIN:1")
		if !isAdmin {
			c.JSON(http.StatusOK, gin.H{
				"error":      CONTROLLER_ADMIN_ERROR_UNAUTHORIZED,
				"isAuthFail": true,
			})
			return
		}

		c.Next()
	}
}

func settings(c *gin.Context) {
	var roles []model.Role
	resultRoles := model.GetDB().Find(&roles)
	if resultRoles.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"userLimit": config.User_Limit,
			"roles":     roles,
		},
	})
}

type BasicForm struct {
	UserLimit uint `form:"userLimit" binding:"required"`
}

func basic(c *gin.Context) {
	var form BasicForm
	err := c.ShouldBindWith(&form, binding.FormPost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":  CONTROLLER_ADMIN_ERROR_BIND,
			"isFail": true,
		})
		return
	}

	config.User_Limit = form.UserLimit

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

type ResetPasswordForm struct {
	OldPassword string `form:"oldPassword" binding:"required"`
	NewPassword string `form:"NewPassword" binding:"required"`
}

func resetPassword(c *gin.Context) {
	user := GetAuthUser(c)

	var form ResetPasswordForm
	err := c.ShouldBindWith(&form, binding.FormPost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":  CONTROLLER_ADMIN_ERROR_BIND,
			"isFail": true,
		})
		return
	}

	if user.Password != form.OldPassword {
		c.JSON(http.StatusOK, gin.H{
			"fail": CONTROLLER_ADMIN_ERROR_PASSWORD,
		})
		return
	}

	resultUpdatePassword := model.GetDB().Model(&user).Updates(model.User{
		Password: form.NewPassword,
	})
	if resultUpdatePassword.RowsAffected != 1 {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func InitAdmin(rg *gin.RouterGroup) {
	api := rg.Group("/admin")
	{
		api.POST("/settings", Auth(), Admin(), settings)
		api.POST("/basic", Auth(), Admin(), basic)
		api.POST("/resetPassword", Auth(), Admin(), resetPassword)
	}
}