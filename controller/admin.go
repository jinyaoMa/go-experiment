package controller

import (
	"jinyaoma/go-experiment/config"
	"jinyaoma/go-experiment/model"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	CONTROLLER_ADMIN_ERROR_BIND         = "binding error"
	CONTROLLER_ADMIN_ERROR_PASSWORD     = "password unmatched"
	CONTROLLER_ADMIN_ERROR_UNAUTHORIZED = "unauthorized user"
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

		c.Set("IsAdmin", isAdmin)

		c.Next()
	}
}

func GetIsAdmin(c *gin.Context) bool {
	var value, exist = c.Get("IsAdmin")
	if exist {
		isAdmin, ok := value.(bool)
		if ok {
			return isAdmin
		}
	}
	return false
}

func settings(c *gin.Context) {
	isAdmin := GetIsAdmin(c)
	if isAdmin {
		var roles []model.Role
		resultRoles := model.GetDB().Order("id desc").Find(&roles)
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
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"userLimit": 0,
			"roles":     make([]model.Role, 0),
		},
	})
}

type BasicUserForm struct {
	Username string `form:"username" binding:"required"`
}

type BasicAdminForm struct {
	BasicUserForm
	UserLimit int64 `form:"userLimit" binding:"required"`
}

func basic(c *gin.Context) {
	user := GetAuthUser(c)
	isAdmin := GetIsAdmin(c)
	if isAdmin {
		var form BasicAdminForm
		err := c.ShouldBindWith(&form, binding.FormPost)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error":  CONTROLLER_ADMIN_ERROR_BIND,
				"isFail": true,
			})
			return
		}

		config.User_Limit = form.UserLimit
		log.Printf("[CONFIG] User_Limit SET TO %d\n", config.User_Limit)

		resultUpdateUsername := model.GetDB().Model(&user).Updates(model.User{
			Name: form.Username,
		})
		if resultUpdateUsername.RowsAffected != 1 {
			c.JSON(http.StatusOK, gin.H{
				"error": http.StatusInternalServerError,
			})
			return
		}
	} else {
		var form BasicUserForm
		err := c.ShouldBindWith(&form, binding.FormPost)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error":  CONTROLLER_ADMIN_ERROR_BIND,
				"isFail": true,
			})
			return
		}

		resultUpdateUsername := model.GetDB().Model(&user).Updates(model.User{
			Name: form.Username,
		})
		if resultUpdateUsername.RowsAffected != 1 {
			c.JSON(http.StatusOK, gin.H{
				"error": http.StatusInternalServerError,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

type ResetPasswordForm struct {
	OldPassword string `form:"oldPassword" binding:"required"`
	NewPassword string `form:"newPassword" binding:"required"`
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

type AddRoleForm struct {
	Name       string `form:"name" binding:"required"`
	Permission string `form:"permission" binding:"required"`
	Space      uint64 `form:"space"`
}

func addRole(c *gin.Context) {
	isAdmin := GetIsAdmin(c)
	if !isAdmin {
		c.JSON(http.StatusOK, gin.H{
			"fail": CONTROLLER_ADMIN_ERROR_UNAUTHORIZED,
		})
		return
	}

	var form AddRoleForm
	err := c.ShouldBindWith(&form, binding.FormPost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":  CONTROLLER_ADMIN_ERROR_BIND,
			"isFail": true,
		})
		return
	}

	newRole := model.Role{
		Name:       form.Name,
		Permission: form.Permission,
		Space:      form.Space,
	}
	resultCreateNewRole := model.GetDB().Create(&newRole)
	if resultCreateNewRole.RowsAffected != 1 {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	var roles []model.Role
	resultRoles := model.GetDB().Order("id desc").Find(&roles)
	if resultRoles.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"roles": roles,
		},
	})
}

type SaveRoleForm struct {
	RoleID     uint   `form:"roleId" binding:"required"`
	Name       string `form:"name" binding:"required"`
	Permission string `form:"permission" binding:"required"`
	Space      uint64 `form:"space"`
}

func saveRole(c *gin.Context) {
	isAdmin := GetIsAdmin(c)
	if !isAdmin {
		c.JSON(http.StatusOK, gin.H{
			"fail": CONTROLLER_ADMIN_ERROR_UNAUTHORIZED,
		})
		return
	}

	var form SaveRoleForm
	err := c.ShouldBindWith(&form, binding.FormPost)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error":  CONTROLLER_ADMIN_ERROR_BIND,
			"isFail": true,
		})
		return
	}

	resultUpdateRole := model.GetDB().Model(&model.Role{}).Where("id = ?", form.RoleID).Updates(model.Role{
		Name:       form.Name,
		Permission: form.Permission,
		Space:      form.Space,
	})
	if resultUpdateRole.RowsAffected != 1 {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	var roles []model.Role
	resultRoles := model.GetDB().Order("id desc").Find(&roles)
	if resultRoles.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"roles": roles,
		},
	})
}

func InitAdmin(rg *gin.RouterGroup) {
	api := rg.Group("/admin").Use(Auth(), Admin())
	{
		api.POST("/settings", settings)
		api.POST("/basic", basic)
		api.POST("/resetPassword", resetPassword)
		api.POST("/role/add", addRole)
		api.POST("/role/save", saveRole)
	}
}
