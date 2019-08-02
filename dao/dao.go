package dao

import (
	"fmt"
	"go-bot/conn"
	"go-bot/utils"
)

// InsertOne .
func InsertOne(data interface{}) {
	_, err := conn.GetCollection().InsertOne(utils.GetCtx(), data)
	if err != nil {
		fmt.Println(err)
	}
}
