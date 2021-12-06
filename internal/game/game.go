package game

import "github.com/wthsths/minigames/pb"

type Game interface {
	SetPlayer(uint64) error
	Deal() error
	Play() error
	Fold() error
	Check() error
	Bet() error
	GetEntity() (*pb.Game, error)
}
