package main

import (
	"chat/handler"
	"chat/websocket"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go/log"
)

func main() {
	fmt.Println("Welcome to MeChat")
	startServer()
}

func startServer() {
	pool := websocket.NewPool()
	go pool.Start()
	router := gin.Default()
	router.GET("/", handler.CheckServerStatus())
	router.GET("/chat", handler.ServeWs(pool))
	s := &http.Server{
		// Get from config.toml
		Addr:    ":3600",
		Handler: router,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Error(err)
		return
	}
}
