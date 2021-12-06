package inmem

import (
	"context"
	"errors"
	"sync"

	"github.com/wthsths/minigames/pb"
)

type inmemRepository struct {
	sync.RWMutex
	games map[int]*pb.Game
	seq   int
}

func NewInmemRepository() (*inmemRepository, error) {
	return &inmemRepository{
		games: make(map[int]*pb.Game),
		seq:   0,
	}, nil
}

func (repo *inmemRepository) CreateGame(ctx context.Context, game *pb.Game) error {
	repo.Lock()
	defer repo.Unlock()

	repo.seq++
	repo.games[repo.seq] = game

	return nil
}

func (repo *inmemRepository) GetGameByUuid(ctx context.Context, uuid string) (*pb.Game, error) {
	repo.Lock()
	defer repo.Unlock()

	for _, game := range repo.games {
		if game.Uuid == uuid {
			return game, nil
		}
	}

	return nil, errors.New("game not found")
}

func (repo *inmemRepository) UpdateByUuid(ctx context.Context, uuid string, game *pb.Game) error {
	repo.Lock()
	defer repo.Unlock()

	for i := 0; i <= len(repo.games); i++ {
		// skip first seq
		if repo.games[i] == nil {
			continue
		}

		if repo.games[i].Uuid == uuid {
			repo.games[i] = game
			return nil
		}
	}

	return errors.New("game could not found")
}
