package ws

import (
	"encoding/json"
	"fmt"
	"go-bot/block"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type data struct {
	Msg string `json:"msg"`
}

// Ping .
func Ping(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close()
	for {
		// 读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			// 客户端关闭连接时也会进入
			fmt.Println(err)
			break
		}
		// JSON 反序列化struct
		msg := &data{}
		json.Unmarshal(message, msg)
		fmt.Println(msg, mt)
		bc := block.GetInstance()
		if err := bc.AddBlock(msg.Msg); err != nil {
			fmt.Println(err)
			v := gin.H{"message": "很遗憾，什么都没有挖到。。。"}
			ws.WriteJSON(v)
			return
		}

		iter := bc.NewIterator()
		bk := iter.Next()
		fmt.Printf("%d\n", bk.Timestamp)
		msg.Msg = fmt.Sprintf("Hash: %x", bk.Hash)

		// JSON序列化，借助gin的gin.H实现
		v := gin.H{"message": msg.Msg}
		ws.WriteJSON(v)
			
		// 写入ws数据 二进制返回
		// err = ws.WriteMessage(mt, message)
	}
}

