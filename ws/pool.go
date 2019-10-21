package ws

import (
	"sync"

	"github.com/gorilla/websocket"
)

// Pool .
type Pool struct {
	mu    sync.Mutex
	conns []*Conn
}

var instance *Pool

// GetPool .
func GetPool() *Pool {
	if instance == nil {
		instance = &Pool{}
	}
	return instance
}

// Conn .
type Conn struct {
	ws        *websocket.Conn
	cid       string
	timeStamp uint32
}

// AddConn .
func (p *Pool) AddConn(ws *websocket.Conn) {

}

// ReadMessage .
func (conn *Conn) ReadMessage() {

}

// WriteMessage .
func (conn *Conn) WriteMessage() {

}
