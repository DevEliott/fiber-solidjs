package player

import (
	"app/server/pkg/models"
	"errors"

	"github.com/google/uuid"
)

// type Service interface {
// 	CreatePlayer(name string) *models.Player
// 	GetPlayer(ID string) (*models.Player, error)
// }

// type service struct {
// 	players []*models.Player
// }

// func NewService() Service {
// 	return &service{
// 		players: make([]*models.Player, 0),
// 	}
// }
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
	return nil, errors.New("No player matched given ID")
}
