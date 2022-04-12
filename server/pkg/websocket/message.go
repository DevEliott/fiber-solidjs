package websocket

import (
	"log"
)

type Message struct {
	Action   string      `json:"action"`
	Data     interface{} `json:"data,omitempty"`
	SenderID string      `json:"senderId"`
}

const (
	UpdateEasel     = "UpdateEasel"
	ExchangeLetters = "ExchangeLetters"
	All             = "All"
)

func HandleMessage(c *Client, m Message) {
	log.Println("Recv:", m)
	switch m.Action {
	case ExchangeLetters:
		message := Message{
			Action:   UpdateEasel,
			SenderID: c.ID,
			// Data:   c.easel,
		}
		message.BroadCastTo(c)
	}
}

func (m *Message) BroadCast() {
	for _, client := range Clients {
		if client.ID != m.SenderID {
			m.BroadCastTo(client)
		}
	}
}

func (m *Message) BroadCastTo(c *Client) {
	err := c.Ws.WriteJSON(m)
	if err != nil {
		log.Printf("failed to broadcast message %s to client %v", m.Data, c.Ws.RemoteAddr())
	}
}
