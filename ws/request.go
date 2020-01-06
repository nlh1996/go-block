package ws

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Request .
type Request struct {
	conn *Connection
	data []byte
}

// Message .
type Message struct {
	ID      uint32
	DataLen uint32
	Data    []byte
}

// GetConn .
func (r *Request) GetConn() *Connection {
	return r.conn
}

// GetData .
func (r *Request) GetData() []byte {
	return r.data
}

// Send .
func (r *Request) Send(msg interface{}) (err error) {
	return r.conn.Send(msg)
}

// Router .
type Router struct {
}

// BeforeHandle .
func (r *Router) BeforeHandle(req *Request) {
	fmt.Println("BeforeHandle call...")
	req.Send(gin.H{"msg": "BeforeHandle call..."})
}

// Handle .
func (r *Router) Handle(req *Request) {
	fmt.Println("Handle call...")
	req.Send(gin.H{"msg": "Handle call..."})
}

// AfterHandle .
func (r *Router) AfterHandle(req *Request) {
	fmt.Println("AfterHandle call...")
	req.Send(gin.H{"msg": "AfterHandle call..."})
}
