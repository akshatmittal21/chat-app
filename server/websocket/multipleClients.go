package websocket

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// Client a websocket client
type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

// Message To write message to client
type Message struct {
	SenderID    string `json:"senderId"`
	MessageType string `json:"messageType"`
	Time        int64  `json:"time"`
	Type        int    `json:"type"`
	Body        string `json:"body"`
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		msgTime := time.Now().Unix()
		message := Message{SenderID: c.ID, MessageType: "message", Time: msgTime, Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}
