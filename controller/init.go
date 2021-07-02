package controller

import (
	"jinyaoma/go-experiment/config"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func RunRouter() {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	router.Run(config.SERVER_PORT)
}
