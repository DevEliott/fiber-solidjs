package handlers

import (
	"app/server/pkg/websocket"
	"log"

	"github.com/gofiber/fiber/v2"
	ws "github.com/gofiber/websocket/v2"
	"github.com/google/uuid"
)

func WS() fiber.Handler {
	return ws.New(HandleWSConnection)
}

func HandleWSConnection(conn *ws.Conn) {
	log.Println("Len before adding:", len(websocket.Clients))
	c := &websocket.Client{
		Ws: conn,
		ID: uuid.NewString(),
	}
	websocket.GeneralPubSub.AddClient(c)
	websocket.Clients = append(websocket.Clients, c)
	log.Println("Len after adding:", len(websocket.Clients))
	log.Println("Clients", websocket.Clients)
	c.StartListening()
}
