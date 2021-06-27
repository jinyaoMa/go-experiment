package main

import (
	"jinyaoma/go-experiment/config"
	"jinyaoma/go-experiment/model"
	// "github.com/gin-gonic/gin"
)

func main() {
	model.InitDB(config.WORKSPACE)

	/*
		r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	*/

}
