package ws

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-bot/block"
	"go-bot/model"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Connection .
type Connection struct {
	wsConnect *websocket.Conn
	inChan    chan []byte
	outChan   chan []byte
	closeChan chan byte

	mutex    sync.Mutex // 对closeChan关闭上锁
	isClosed bool       // 防止closeChan被关闭多次
}

// InitConnection .
func InitConnection(wsConn *websocket.Conn) *Connection {
	conn := &Connection{
		wsConnect: wsConn,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte, 1),
	}
	return conn
}

// Start .
func (conn *Connection) Start() (data []byte, err error) {
	// 区块链单例
	bc := block.GetInstance()
	// 区块迭代器
	iter := bc.NewIterator()
	// 启动读协程
	go conn.readLoop()
	// 启动写协程
	go conn.writeLoop()

	for {
		select {
		case data = <-conn.inChan:
			// JSON 反序列化struct
			res := &model.Response{}
			json.Unmarshal(data, res)
			if err := bc.AddBlock(res.Msg); err != nil {
				log.Println(err)
				v := gin.H{"message": "很遗憾，什么都没有挖到。。。"}
				conn.wsConnect.WriteJSON(v)
				conn.closeChan <- 0
			}

			bk := iter.Next()
			fmt.Printf("%d\n", bk.Timestamp)
			res.Msg = fmt.Sprintf("Hash: %x", bk.Hash)

			// JSON序列化，借助gin的gin.H实现
			v := gin.H{"data": res}
			conn.wsConnect.WriteJSON(v)
		case <-conn.closeChan:
			err = errors.New("connection is closeed")
			return
		}
	}
}

// WriteMessage .
func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("connection is closeed")
	}
	return
}

// Close .
func (conn *Connection) Close() {
	// 线程安全，可多次调用
	conn.wsConnect.Close()
	// 利用标记，让closeChan只关闭一次
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		close(conn.inChan)
		close(conn.outChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
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
		case <-conn.closeChan:
			goto ERR
		}
		if err = conn.wsConnect.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()

}
