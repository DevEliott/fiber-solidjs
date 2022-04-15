package player

import (
	"app/server/pkg/models"
	"errors"

	"github.com/google/uuid"
)

var (
	players []*models.Player
)

func CreatePlayer(name string) *models.Player {
	p := &models.Player{
		ID:   uuid.NewString(),
		Name: name,
	}
	players = append(players, p)
	return p
}

func GetPlayer(ID string) (*models.Player, error) {
	for _, p := range players {
		if p.ID == ID {
			return p, nil
		}
	}
	return nil, errors.New("No player found with the given ID")
}
