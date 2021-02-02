# chat-app
A basic gorilla websocket based chat app

## Server Guide

Run the command `go mod download` in the server directory to  install the required dependencies for the Go server.

To run the server run `go run main.go` command.

For the development of server, I have used [gorilla websocket](https://github.com/gorilla/websocket) implementation.
A single route i.e. `http://localhost:3600/chat` is exposed which can be used to integrate a connection pool.

Accessing this API will enter a user in the connection pool and they will be able to broadcast or recieve messages directly.

## Client Guide

To install the required dependencies for the client, download the `node_modules` using either of the following command:
  ```
  npm install
      or
  yarn install
  ```
  
  To start the client, run the following command :
  ```
  yarn serve
      or
  npm run serve
  ```
  
  To connect chat server to the javascript application, use the `WebSocket` class of the javascript. Following is the syntax to connect to the server.
  ```
let connection = new WebSocket("ws://localhost:3600/chat");  
```
  Please refer to the [WebSocket](https://developer.mozilla.org/en-US/docs/Web/API/WebSocket) documentation for its methods.
