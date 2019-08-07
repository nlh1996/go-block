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
