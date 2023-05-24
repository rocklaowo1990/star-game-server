package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type SendMessage struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}

type ReplyMessage struct {
	Port    int    `json:"code"`
	Type    int    `json:"type"`
	Content string `json:"content"`
}

type Client struct {
	Uid     string          `json:"uid"`
	SendUid string          `json:"sendUid"`
	Socket  *websocket.Conn `json:"socket"`
	Send    chan []byte     `json:"send"`
}

type Broadcast struct {
	Client  *Client `json:"client"`
	Message []byte  `json:"message"`
	Type    int     `json:"type"`
}

type ClientManager struct {
	Clients   map[string]*Client `json:"clients"`
	Broadcast chan *Broadcast    `json:"broadcast"`
	Reply     chan *Client       `json:"reply"`
	SignUp    chan *Client       `json:"signUp"`
	SignOut   chan *Client       `json:"signOut"`
}

type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,,omitempty"`
}

var Manager = ClientManager{
	Clients:   make(map[string]*Client),
	Broadcast: make(chan *Broadcast),
	SignUp:    make(chan *Client),
	SignOut:   make(chan *Client),
}

func Handler(c *gin.Context) {
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}).Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		http.NotFound(c.Writer, c.Request)
	}

	client := &Client{
		Uid:     "777",
		SendUid: "999",
		Socket:  conn,
		Send:    make(chan []byte),
	}
	fmt.Println(conn.ReadMessage())
	Manager.SignUp <- client

	go client.Read()
	go client.Write()
}

func (c *Client) Read() {
	defer func() {
		Manager.SignOut <- c
		_ = c.Socket.Close()
	}()

	for {
		c.Socket.PingHandler()
		sendMessage := new(SendMessage)

		//c.Socket.SendMessage(sendMessage)
		messageType, content, err := c.Socket.ReadMessage()
		if err != nil {
			fmt.Println(sendMessage)
			fmt.Println("=> 数据格式不正确", err)
			Manager.SignOut <- c
			_ = c.Socket.Close()
			break
		}

		fmt.Println(messageType, string(content))

		if sendMessage.Type == 1 { //发送消息
			r1, _ := RedisClient.Get(Context, c.Uid).Result()
			r2, _ := RedisClient.Get(Context, c.SendUid).Result()

			if r1 > "3" && r2 == "" {
				replyMessage := ReplyMessage{
					Port:    8080,
					Type:    200,
					Content: "达到限制",
				}

				msg, _ := json.Marshal(replyMessage)
				_ = c.Socket.WriteMessage(websocket.TextMessage, msg)
				continue
			}

			RedisClient.Incr(Context, c.Uid)
			_, _ = RedisClient.Expire(Context, c.Uid, time.Hour*24*30*3).Result()

			fmt.Println(c.Uid, "发送消息给 ", c.SendUid, ": =>", sendMessage.Content)
			Manager.Broadcast <- &Broadcast{
				Client:  c,
				Message: []byte(sendMessage.Content),
			}
		}
	}
}

func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			replyMessage := ReplyMessage{
				Port:    8080,
				Type:    200,
				Content: string(message),
			}

			msg, _ := json.Marshal(replyMessage)
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)

		case message, ok := <-c.Send:
			if !ok {
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			replyMessage := ReplyMessage{
				Port:    8080,
				Type:    200,
				Content: string(message),
			}

			msg, _ := json.Marshal(replyMessage)
			_ = c.Socket.WriteMessage(websocket.TextMessage, msg)

		}

	}
}

func (manager *ClientManager) Client() {
	for {
		// fmt.Println("=> 监听管道通信")
		select {
		case conn := <-Manager.SignUp:
			fmt.Println("=> 有新的连接", conn.Uid)
			Manager.Clients[conn.Uid] = conn
			replyMessage := ReplyMessage{
				Port:    8080,
				Type:    200,
				Content: "已经连接到服务器了",
			}
			msg, _ := json.Marshal(replyMessage)
			_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
		case conn := <-Manager.SignOut:
			fmt.Println("连接失败", conn.Uid)
			if _, ok := Manager.Clients[conn.Uid]; ok {
				replyMessage := &ReplyMessage{
					Port:    8080,
					Type:    200,
					Content: "连接中断",
				}
				msg, _ := json.Marshal(replyMessage)
				_ = conn.Socket.WriteMessage(websocket.TextMessage, msg)
				close(conn.Send)
				delete(manager.Clients, conn.Uid)
			}
		}
	}
}
