package controller

import (
	"errors"
	"jinyaoma/go-experiment/config"
	"jinyaoma/go-experiment/middleware"
	"jinyaoma/go-experiment/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

const (
	CONTROLLER_AUTH_ERROR_BIND           = "binding error"
	CONTROLLER_AUTH_ERROR_USER_NOT_FOUND = "id token unmatched"
)

type AuthForm struct {
	ID    uint   `form:"id" binding:"required"`
	Token string `form:"token" binding:"required"`
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var form AuthForm
		err := c.ShouldBindWith(&form, binding.FormPost)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error":      CONTROLLER_AUTH_ERROR_BIND,
				"isAuthFail": true,
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

		c.Set("AuthForm", form)
		c.Set("User", user)

		c.Next()
	}
}

func GetAuthForm(c *gin.Context) AuthForm {
	var value, exist = c.Get("AuthForm")
	if exist {
		form, ok := value.(AuthForm)
		if ok {
			return form
		}
	}
	return AuthForm{}
}

func GetAuthUser(c *gin.Context) model.User {
	var value, exist = c.Get("User")
	if exist {
		user, ok := value.(model.User)
		if ok {
			return user
		}
	}
	return model.User{}
}

func RunRouter() {
	router := gin.Default()
	router.Use(middleware.Cors())

	api := router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		InitLogin(api)
		InitLogout(api)
		InitSignup(api)
		InitFiles(api)
		InitAdmin(api)
	}

	router.Run(config.SERVER_PORT)
}
