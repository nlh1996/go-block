package router

import (
	"go-bot/controller/clog"
	"go-bot/env"
	"go-bot/ws"
	"log"
	"go-bot/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/nlh1996/utils"
)

// Init .
func Init() {
	router := gin.Default()
	router.Use(middleware.CrossDomain())
	router.GET("/", Handler)
	router.POST("/clog", clog.LogFromClient)
	router.Run(":" + utils.IntToString(env.GlobalData.Port))
}

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// fmt.Println(r.Header["Origin"])
		return true
	},
}

// Handler .
func Handler(c *gin.Context) {
	// 升级get请求为webSocket协议
	connection, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	conn, err := ws.NewConnection(connection)
	if err != nil {
		connection.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		return
	}
	conn.Start()
}
