package handlers

import (
	"app/server/pkg/websocket"
	"log"

	"github.com/gofiber/fiber/v2"
	ws "github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
)

func WS() fiber.Handler {
	return ws.New(HandleNewConnection)
}

func HandleNewConnection(conn *ws.Conn) {
	c := &websocket.Client{
		Ws:   conn,
		ID:   uuid.NewString(),
		Done: make(chan bool),
	}
	websocket.Clients = append(websocket.Clients, c)
	log.Println(len(websocket.Clients))
	go c.StartListening()
	<-c.Done
}
