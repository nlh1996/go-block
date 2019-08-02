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

// Ping webSocket请求Ping 返回pong
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
		msg := &data{}
		json.Unmarshal(message, msg)
		fmt.Println(msg, mt)
		message = addBlock(msg.Msg)
		msg.Msg = fmt.Sprintf("Hash: %x", message)
		// 写入ws数据 二进制返回
		// err = ws.WriteMessage(mt, message)
		// 返回JSON字符串，借助gin的gin.H实现
		v := gin.H{"message": msg.Msg}
		err = ws.WriteJSON(v)
		if err != nil {
			break
		}
	}
}

func addBlock(msg string) []byte {
	bc := block.GetInstance()
	index := bc.AddBlock(msg)
	bk := bc.GetBlockByIndex(index)
	// for _, bk := range bc.GetBlockchain() {
	// 	fmt.Printf("Prev.hash: %x\n", bk.PrevBlockHash)
	// 	fmt.Printf("Data: \"%s\"\n", bk.Data)
	// 	fmt.Printf("Hash: %x\n", bk.Hash)
	// 	pow := block.NewProofOfWork(bk)
	// 	fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
	// }
	return bk.Hash
}
