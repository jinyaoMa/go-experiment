package controller

import (
	"jinyaoma/go-experiment/config"

	"github.com/gin-gonic/gin"
)

func RunRouter() {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		initLogin(api)
	}

	router.Run(config.SERVER_PORT)
}
