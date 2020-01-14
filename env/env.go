package env

import (
	"log"

	"github.com/nlh1996/utils"
)

// GlobalObj .
type GlobalObj struct {
	Server ServerCnf
	Conn   ConnCnf
}

// ServerCnf 服务配置
type ServerCnf struct {
	Host       string
	Port       int
	MgoAddress string
	MgoPort    int
	DBName     string
}

// ConnCnf 连接配置
type ConnCnf struct {
	PoolConnNum int
	MaxConnNum  int
}

// GlobalData .
var GlobalData *GlobalObj

func init() {
	// 默认配置
	GlobalData = &GlobalObj{
		Server: ServerCnf{
			Host:       "0.0.0.0",
			Port:       3000,
			MgoAddress: "localhost",
			MgoPort:    27017,
			DBName:     "transaction",
		},
	}
	// 读取配置文件
	if err := utils.ReadFile(GlobalData, "conf/conf.json"); err != nil {
		log.Panicln(err)
	}
}
