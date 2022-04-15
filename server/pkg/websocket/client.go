package websocket

import (
	"log"

	"github.com/gofiber/websocket/v2"
)

type Client struct {
	Ws *websocket.Conn
	ID string
}

func (c *Client) StartListening() {
	var packet Packet
	for {
		if err := c.Ws.ReadJSON(&packet); err != nil {
			log.Println("read:", err)
			HandleDisconnect(c)
			break
		}
		HandlePacket(c, packet)
	}
}
