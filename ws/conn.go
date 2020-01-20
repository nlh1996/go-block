package ws

import (
	"encoding/json"
	"errors"
	"go-bot/env"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

// Connection .
type Connection struct {
	ID         int
	wsConnect  *websocket.Conn
	inChan     chan []byte
	outChan    chan []byte
	closeChan  chan byte
	sync.Mutex      // 对closeChan关闭上锁
	IsClosed   bool // 防止closeChan被关闭多次
	router     *Router
}

// 预先定义通道存储id
var cidCh chan int

func init() {
	// 定义可复用的id额外预留100个id
	len := env.GlobalData.Conn.MaxConnNum + 100
	cidCh = make(chan int, len)
	for i := 1; i <= len; i++ {
		cidCh <- i
	}
}

// NewConnection .
func NewConnection(wsConn *websocket.Conn) (*Connection, error) {
	conn := &Connection{
		wsConnect: wsConn,
		inChan:    make(chan []byte, 1024),
		outChan:   make(chan []byte, 1024),
		closeChan: make(chan byte, 1),
		IsClosed:  false,
		router:    &Router{},
	}
	// 连接池可用ID不足,预留的100防止高并发，其它协程抢夺
	if len(cidCh) <= 100 {
		err := errors.New("没有可用的连接，请稍后重试！")
		return nil, err
	}
	conn.ID = <-cidCh
	p := GetConnPool()
	p.Set(conn)
	return conn, nil
}

// Start .
func (conn *Connection) Start() (data []byte, err error) {
	// 启动读协程
	go conn.readLoop()
	// 启动写协程
	go conn.writeLoop()

	for {
		select {
		case data = <-conn.inChan:
			req := Request{
				conn: conn,
				data: data,
			}
			// 请求跟路由绑定
			go func(r *Request) {
				conn.router.BeforeHandle(r)
				conn.router.Handle(r)
				conn.router.AfterHandle(r)
			}(&req)
		case <-conn.closeChan:
			return
		}
	}
}

// AddRouter .
// func (conn *Connection) AddRouter(r *Router) {
// 	conn.router = r
// }

// Close .
func (conn *Connection) Close() {
	// 线程安全，可多次调用
	conn.wsConnect.Close()
	// 利用标记，让closeChan只关闭一次
	conn.Lock()

	if !conn.IsClosed {
		log.Println("连接", conn.ID, "已经关闭！！！")
		close(conn.closeChan)
		cidCh <- conn.ID
		GetConnPool().DelByID(conn.ID)
		conn.IsClosed = true
	}
	conn.Unlock()
}

// Send .
func (conn *Connection) Send(msg interface{}) (err error) {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("SEND->", string(msgBytes))

	select {
	case conn.outChan <- msgBytes:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

// 内部实现
func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = conn.wsConnect.ReadMessage(); err != nil {
			goto ERR
		}
		//阻塞在这里，等待inChan有空闲位置
		select {
		case conn.inChan <- data:
		case <-conn.closeChan: // closeChan 感知 conn断开
			goto ERR
		}
	}

ERR:
	conn.Close()
}

func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)
	for {
		select {
		case data = <-conn.outChan:
			if err = conn.wsConnect.WriteMessage(websocket.TextMessage, data); err != nil {
				goto ERR
			}
		case <-conn.closeChan:
			goto ERR
		}
	}

ERR:
	conn.Close()

}
