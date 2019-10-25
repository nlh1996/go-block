package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// fmt.Println(r.Header["Origin"])
		return true
	},
}

// Handler .
func Handler(c *gin.Context) {
	// 升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	conn, err := InitConnection(ws); if err != nil {
		ws.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		return
	}
	p := GetInstance()
	p.Pool[conn.cid] = conn
	conn.Start()
}
