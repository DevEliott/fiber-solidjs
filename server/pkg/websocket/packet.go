package websocket

import (
	"log"
)

type Packet struct {
	Action   string `json:"action"`
	Data     any    `json:"data,omitempty"`
	SenderID string `json:"senderId,omitempty"`
}

const (
	UpdateEasel     = "UpdateEasel"
	ExchangeLetters = "ExchangeLetters"
	ID              = "ID"
	Message         = "Message"
)

func HandlePacket(c *Client, recv Packet) {
	log.Println("Recv:", recv)
	switch recv.Action {
	case ExchangeLetters:
		handleExchangeLetters(c, recv)
	case Message:
		handleMessage(c, recv)
	}
}

func handleMessage(c *Client, recv Packet) {
	recv.BroadCast()
}

func handleExchangeLetters(c *Client, recv Packet) {
	p := Packet{
		Action:   UpdateEasel,
		SenderID: c.ID,
		// Data:   c.easel,
	}
	p.BroadCastTo(c)
}

func (p *Packet) BroadCast() {
	for _, client := range Clients {
		log.Println(client.ID)
		if client.ID != p.SenderID {
			log.Println("Sending to client:", client.ID)
			p.BroadCastTo(client)
		}
	}
}

func (p *Packet) BroadCastTo(c *Client) {
	err := c.Ws.WriteJSON(p)
	if err != nil {
		log.Printf("failed to broadcast message %s to client %v", p.Data, c.Ws.RemoteAddr())
	}
}
