package websocket

import (
	"log"

	"github.com/gofiber/websocket/v2"
)

type Client struct {
	Ws   *websocket.Conn
	ID   string
	Done chan bool
}

var (
	Clients []*Client
)

func (c *Client) StartListening() {
	defer func() {
		c.Done <- true
	}()

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

func HandleDisconnect(c *Client) {
	log.Println("Disconnecting:", c.Ws.RemoteAddr())

	for i, v := range Clients {
		if v == c {
			Clients = append(Clients[:i], Clients[i+1:]...)
			break
		}
	}
	c.Ws.Close()
}
