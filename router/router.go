package router

import (
	"go-bot/ws"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Init .
func Init() {
	router := gin.Default()
	router.GET("/", Handler)
	router.Run(":3000")
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
		return
	}
	conn, err := ws.InitConnection(connection)
	if err != nil {
		connection.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		return
	}
	conn.Start()
}
