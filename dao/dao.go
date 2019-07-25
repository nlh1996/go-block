package dao

import (
	"fmt"
	"go-bot/conn"
	"go-bot/utils"
)

var col = conn.GetCollection()

// InsertOne .
func InsertOne(data interface{}) {
	_, err := col.InsertOne(utils.GetCtx(), data)
	if err != nil {
		fmt.Println(err)
	}
}
