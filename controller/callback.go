package controller

import (
	"encoding/json"
	"fmt"
	"go-bot/block"
	"go-bot/model"
	"go-bot/ws"
	"log"

	"github.com/gin-gonic/gin"
)

// CallBack .
// func CallBack(conn *ws.Connection, data []byte, cnt int) error {
// 	fmt.Print("callback enter")
// 	// 区块链单例
// 	bc := block.GetInstance()
// 	// 区块迭代器
// 	iter := bc.NewIterator()

// 	// JSON 反序列化struct
// 	res := &model.Response{}
// 	json.Unmarshal(data, res)
// 	if res.Code == 101 {
// 		if err := bc.AddBlock(res.Msg); err != nil {
// 			log.Println(err)
// 			msg := gin.H{"message": "很遗憾，什么都没有挖到。。。"}
// 			conn.Send(msg)
// 		}
// 		bk := iter.Next()
// 		fmt.Printf("%d\n", bk.Timestamp)
// 		res.Msg = fmt.Sprintf("Hash: %x", bk.Hash)
// 		conn.Send(res)
// 	}
// }
