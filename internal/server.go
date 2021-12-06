package minigames

import (
	"context"
	"errors"

	"github.com/wthsths/minigames/internal/game"
	texasholdembonus "github.com/wthsths/minigames/internal/game/texas-holdem-bonus"
	"github.com/wthsths/minigames/internal/repository"
	"github.com/wthsths/minigames/pb"
)

type Server struct {
	repo repository.Repository
}

func NewServer(repo repository.Repository) (*Server, error) {
	return &Server{
		repo: repo,
	}, nil
}

func (s *Server) CreateGame(ctx context.Context, req *pb.CreateGameRequest) (*pb.CreateGameResponse, error) {
	var game game.Game
	var err error

	switch req.Type {
	case pb.GameType_GAME_TYPE_TEXAS_HOLDEM_BONUS:
		game, err = texasholdembonus.NewGame()
	default:
		return nil, errors.New("not implemented")
	}

	// If any error occured while generating a new game
	// we do response to the client
	if err != nil {
		return nil, err
	}

	game.SetPlayer(req.Player)

	entity, _ := game.GetEntity()

	// Save game
	if err := s.repo.CreateGame(ctx, entity); err != nil {
		return nil, err
	}

	return &pb.CreateGameResponse{
		Game: entity,
	}, nil
}

func (s *Server) GetGameByUuid(ctx context.Context, req *pb.GetGameByUuidRequest) (*pb.GetGameByUuidResponse, error) {
	game, err := s.repo.GetGameByUuid(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}

	return &pb.GetGameByUuidResponse{
		Game: game,
	}, nil
}

func (s *Server) PlayGame(ctx context.Context, req *pb.PlayGameRequest) (*pb.PlayGameResponse, error) {
	game, err := s.getGame(ctx, req.Uuid)
	if err != nil {
		return nil, err
	}

	switch req.Action {
	case pb.PlayGameRequest_ACTION_DEAL:
		err = game.Deal()
	case pb.PlayGameRequest_ACTION_PLAY:
		err = game.Play()
	case pb.PlayGameRequest_ACTION_FOLD:
		err = game.Fold()
	case pb.PlayGameRequest_ACTION_CHECK:
		err = game.Check()
	case pb.PlayGameRequest_ACTION_BET:
		err = game.Bet()
	default:
		return nil, errors.New("invalid action type")
	}

	if err != nil {
		return nil, err
	}

	entity, err := game.GetEntity()
	if err != nil {
		return nil, err
	}

	if err := s.repo.UpdateByUuid(ctx, req.Uuid, entity); err != nil {
		return nil, err
	}

	return &pb.PlayGameResponse{
		Game: entity,
	}, nil
}

func (s *Server) getGame(ctx context.Context, uuid string) (game.Game, error) {
	var game game.Game
	var err error
	var entity *pb.Game

	entity, err = s.repo.GetGameByUuid(ctx, uuid)
	if err != nil {
		return nil, err
	}

	switch entity.Type {
	case pb.GameType_GAME_TYPE_TEXAS_HOLDEM_BONUS:
		game, err = texasholdembonus.NewGameWithEntity(entity)
	default:
		return nil, errors.New("not implemented game")
	}

	if err != nil {
		return nil, err
	}

	return game, nil
}
