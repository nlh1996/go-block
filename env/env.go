package env

import (
	"log"

	"github.com/nlh1996/utils"
)

// GlobalObj .
type GlobalObj struct {
	Host       string
	Port       int
	MgoAddress string
	MgoPort    int
	DBName     string
}

// GlobalData .
var GlobalData *GlobalObj

func init() {
	// 默认配置
	GlobalData = &GlobalObj{
		Host:       "0.0.0.0",
		Port:       3000,
		MgoAddress: "localhost",
		MgoPort:    27017,
		DBName:     "transaction",
	}
	// 读取配置文件
	if err := utils.ReadFile(GlobalData, "conf/conf.json"); err != nil {
		log.Panicln(err)
	}
}
