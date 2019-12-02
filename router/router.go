package router

import (
	"go-bot/env"
	"go-bot/ws"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/nlh1996/utils"
)

// Init .
func Init() {
	router := gin.Default()
	router.GET("/", Handler)
	router.Run(":" + utils.IntToString(env.GlobalOblect.Port))
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
	conn, err := ws.InitConnection(connection)
	if err != nil {
		connection.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		return
	}
	conn.Start()
}
