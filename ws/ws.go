package ws

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		fmt.Println(r.Header["Origin"])
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

	conn := InitConnection(ws)
	conn.Start()

	// // 区块链单例
	// bc := block.GetInstance()
	// // 区块迭代器
	// iter := bc.NewIterator()

	// for {
	// 	// 读取ws中的数据
	// 	_, message, err := ws.ReadMessage()
	// 	if err != nil {
	// 		// 客户端关闭连接时也会进入
	// 		log.Println(err)
	// 		break
	// 	}
	// 	// JSON 反序列化struct
	// 	res := &model.Response{}
	// 	json.Unmarshal(message, res)

	// 	if err := bc.AddBlock(res.Msg); err != nil {
	// 		log.Println(err)
	// 		v := gin.H{"message": "很遗憾，什么都没有挖到。。。"}
	// 		ws.WriteJSON(v)
	// 		return
	// 	}

	// 	bk := iter.Next()
	// 	fmt.Printf("%d\n", bk.Timestamp)
	// 	res.Msg = fmt.Sprintf("Hash: %x", bk.Hash)

	// 	// JSON序列化，借助gin的gin.H实现
	// 	v := gin.H{"data": res}
	// 	ws.WriteJSON(v)

	// 	// 写入ws数据 二进制返回
	// 	// err = ws.WriteMessage(mt, message)
	//}
}
