package websocket

type PubSub struct {
	Clients []*Client
}

var (
	GeneralPubSub *PubSub = &PubSub{
		Clients: make([]*Client, 0),
	}
)

func (ps *PubSub) AddClient(c *Client) *PubSub {
	ps.Clients = append(ps.Clients, c)
	p := Packet{
		Action: ID,
		Data:   c.ID,
	}
	p.BroadCastTo(c)
	p = Packet{
		Action:   Message,
		Data:     "Hello from server!",
		SenderID: c.ID,
	}
	p.BroadCastTo(c)
	return ps
}
