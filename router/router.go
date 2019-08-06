package router

import (
	"go-bot/ws"

	"github.com/gin-gonic/gin"
)

// Init .
func Init() {
	router := gin.Default()
	router.GET("/", ws.Ping)
	router.Run(":3000")
}
