package game

import (
	"app/server/pkg/models"
	"errors"

	"github.com/google/uuid"
)

var (
	Games []*models.Game
)

func CreateGame(name string) *models.Game {
	g := &models.Game{
		ID:      uuid.NewString(),
		Players: make([]*models.Player, 0),
	}
	Games = append(Games, g)
	return g
}

func GetGame(ID string) (*models.Game, error) {
	for _, g := range Games {
		if g.ID == ID {
			return g, nil
		}
	}
	return nil, errors.New("No Game found with the given ID")
}
