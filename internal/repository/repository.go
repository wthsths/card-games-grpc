package repository

import (
	"context"

	"github.com/wthsths/minigames/pb"
)

type Repository interface {
	CreateGame(context.Context, *pb.Game) error
	GetGameByUuid(context.Context, string) (*pb.Game, error)
	UpdateByUuid(context.Context, string, *pb.Game) error
	GetGamesByPlayer(context.Context, uint64) ([]*pb.Game, error)
}
