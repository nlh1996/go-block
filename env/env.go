package env

import "github.com/nlh1996/utils"

// GlobalObj .
type GlobalObj struct {
	Host       string
	Port       int
	MgoAddress string
	MgoPort    int
	DBName     string
}

// GlobalOblect .
var GlobalOblect *GlobalObj

func init() {
	// 默认配置
	GlobalOblect = &GlobalObj{
		Host:       "0.0.0.0",
		Port:       3000,
		MgoAddress: "localhost",
		MgoPort:    27017,
		DBName:     "transaction",
	}
	// 读取配置文件
	utils.ReadFile(GlobalOblect)
}
