package utils

import (
	"encoding/json"
	"io/ioutil"
)

/*
	读取全局配置
*/

// GlobalObj .
type GlobalObj struct {
	Host           string
	Port           int
	MaxConn        int
	MaxPackageSize int
}

// GlobalOblect .
var GlobalOblect *GlobalObj

func init() {
	GlobalOblect = &GlobalObj{}
	GlobalOblect.Reload()
}

// Reload .
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/conf.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &GlobalOblect); err != nil {
		panic(err)
	}
}
