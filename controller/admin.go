package controller

import (
	"github.com/gin-gonic/gin"
)

func settings(c *gin.Context) {

}

func InitAdmin(rg *gin.RouterGroup) {
	api := rg.Group("/admin")
	{
		api.POST("/settings", settings)
	}
}
