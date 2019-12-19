package ws

import (
	"encoding/json"
	"errors"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

// Connection .
type Connection struct {
	cid       int
	wsConnect *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan byte
	mutex     sync.Mutex // 对closeChan关闭上锁
	IsClosed  bool       // 防止closeChan被关闭多次
	router    *Router
}

// 预先定义通道存储id
var cidCh chan int

func init() {
	// 定义10100个有效可复用的id
	cidCh = make(chan int, 10100)
	for i := 1; i <= 10100; i++ {
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
	// 连接池可用IP不足100
	if len(cidCh) < 100 {
		err := errors.New("没有可用的连接，请稍后重试！")
		return nil, err
	}
	conn.cid = <-cidCh
	p := GetInstance()
	p.Pool[conn.cid] = conn
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
	conn.mutex.Lock()

	if !conn.IsClosed {
		log.Println("连接", conn.cid, "已经关闭！！！")
		close(conn.closeChan)
		cidCh <- conn.cid
		delete(GetInstance().Pool, conn.cid)
		conn.IsClosed = true
	}
	conn.mutex.Unlock()
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
		err = errors.New("connection is closeed")
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
