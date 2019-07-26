package router

import (
	"go-bot/controller"
	"go-bot/middleware"

	"github.com/gin-gonic/gin"
)

// Init .
func Init() {
	router := gin.Default()
	router.Use(middleware.CrossDomain())
	v1 := router.Group("/v1")
	{
		v1.POST("/addblock", controller.AddBlock)
	}
	router.Run(":8080")
}
