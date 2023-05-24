package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Connection struct {
	wsConn *websocket.Conn
	//读取websocket的channel
	inChan chan []byte
	//给websocket写消息的channel
	outChan   chan []byte
	closeChan chan byte
	mutex     sync.Mutex
	//closeChan 状态
	isClosed bool
}

// 初始化长连接
func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConn:    wsConn,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte, 1),
	}
	//启动读协程
	go conn.readLoop()
	//启动写协程
	go conn.writeLoop()
	return
}

// 读取websocket消息
func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

// 发送消息到websocket
func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

// 关闭连接
func (conn *Connection) Close() {
	//线程安全的Close,可重入
	conn.wsConn.Close()

	//只执行一次
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			goto ERR
		}

		//如果数据量过大阻塞在这里,等待inChan有空闲的位置！
		select {
		case conn.inChan <- data:
		case <-conn.closeChan:
			//closeChan关闭的时候
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
		if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}

var (
	upgrade = websocket.Upgrader{
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func WebSocketHandler(c *gin.Context) {
	var (
		//websocket 长连接
		ws   *websocket.Conn
		err  error
		conn *Connection
		data []byte
	)
	//header中添加Upgrade:websocket
	if ws, err = upgrade.Upgrade(c.Writer, c.Request, nil); err != nil {
		return
	}

	if conn, err = InitConnection(ws); err != nil {
		goto ERR
	}

	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}

		gameMessage := GameMessage{}

		if err = json.Unmarshal(data, &gameMessage); err != nil {
			goto ERR
		}

		if err = GameManager(conn, &gameMessage); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()
}
