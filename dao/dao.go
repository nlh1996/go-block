package dao

import (
	"fmt"
	"go-bot/database"
	"go-bot/utils"
)

// InsertOne .
func InsertOne(data interface{}) {
	_, err := database.GetCollection().InsertOne(utils.GetCtx(), data)
	if err != nil {
		fmt.Println(err)
	}
}
