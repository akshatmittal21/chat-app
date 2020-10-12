package handler

import (
	"chat/websocket"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go/log"
)

// CheckServerStatus : to check server status
func CheckServerStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Chat app started successfully")
		// Add http code instead of 200
		c.JSON(http.StatusOK, "Server Running Successfully")
	}
}

// ServeWs : to start websocket
func ServeWs(pool *websocket.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("WebSocket Endpoint Hit")
		ws, err := websocket.Upgrader(c.Writer, c.Request)
		if err != nil {
			log.Error(err)
		}
		ID, err := uuid.NewRandom()
		if err != nil {
			log.Error(err)
		}
		client := websocket.Client{
			ID:   ID.String(),
			Conn: ws,
			Pool: pool,
		}

		pool.Register <- &client
		client.Read()
	}
}
