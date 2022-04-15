package models

type Game struct {
	ID      string    `json:"id"`
	Players []*Player `json:"players"`
}
