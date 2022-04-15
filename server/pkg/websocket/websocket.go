package websocket

import (
	"log"
)

var (
	Clients []*Client
)

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
