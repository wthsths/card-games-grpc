package game

import "github.com/wthsths/card-games-grpc/pb"

type Game interface {
	SetPlayer(uint64) error
	Deal(uint64, []*pb.PlayGameRequest_Bonus) error
	Play() error
	Fold() error
	Check() error
	Bet() error
	GetEntity() (*pb.Game, error)
}
