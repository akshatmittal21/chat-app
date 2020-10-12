package websocket

import (
	"fmt"
	"time"
)

// Pool : a pool of connection
type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

// NewPool is to initliase new pool
func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

// Start : to start the pool
func (pool *Pool) Start() {
	for {
		select {
		case newClient := <-pool.Register:
			pool.Clients[newClient] = true
			msgTime := time.Now().Unix()

			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client := range pool.Clients {
				if client.ID == newClient.ID {
					client.Conn.WriteJSON(struct {
						UserID string `json:"userId"`
					}{UserID: newClient.ID})
				}
				client.Conn.WriteJSON(Message{SenderID: newClient.ID, Time: msgTime, MessageType: "info", Type: 1, Body: "New User Joined..."})
			}
			break
		case rmClient := <-pool.Unregister:
			delete(pool.Clients, rmClient)
			msgTime := time.Now().Unix()

			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{SenderID: rmClient.ID, Time: msgTime, MessageType: "info", Type: 1, Body: "User Disconnected..."})
			}
			break
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
