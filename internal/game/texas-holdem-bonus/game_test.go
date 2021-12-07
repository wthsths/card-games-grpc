package texasholdembonus

import (
	"encoding/json"
	"testing"

	"github.com/wthsths/minigames/pb"
)

func TestNewGame(t *testing.T) {
	_, err := NewGame()
	if err != nil {
		t.Fail()
	}
}

func TestNewGameWithEntity(t *testing.T) {
	entity := new(pb.Game)
	entity.Data, _ = json.Marshal(&EntityData{})

	_, err := NewGameWithEntity(entity)
	if err != nil {
		t.Fail()
	}
}

func TestNewGameWithEntityWithEmptyEntity(t *testing.T) {
	entity := new(pb.Game)
	_, err := NewGameWithEntity(entity)
	if err == nil {
		t.Fail()
	}
}

func TestNewGameWithEntityWithDifferentGameType(t *testing.T) {
	entity := new(pb.Game)
	entity.Type = -1
	_, err := NewGameWithEntity(entity)
	if err == nil {
		t.Fail()
	}
}

func TestSetPlayerId(t *testing.T) {
	game, _ := NewGame()
	game.SetPlayer(1)

	if game.entity.Player != 1 {
		t.Fail()
	}
}

func TestDeal(t *testing.T) {
	game, _ := NewGame()
	err := game.Deal(0, []*pb.PlayGameRequest_Bonus{})
	if err != nil {
		t.Fail()
	}
}

func TestDealIfDealtBefore(t *testing.T) {
	game, _ := NewGame()
	game.Deal(0, []*pb.PlayGameRequest_Bonus{})
	err := game.Deal(0, []*pb.PlayGameRequest_Bonus{})
	if err == nil {
		t.Fail()
	}
}

func TestPlayIfNotDealt(t *testing.T) {
	game, _ := NewGame()
	err := game.Play()
	if err == nil {
		t.Fail()
	}
}

func TestFold(t *testing.T) {
	game, _ := NewGame()
	game.Deal(0, []*pb.PlayGameRequest_Bonus{})
	err := game.Fold()
	if err != nil {
		t.Fail()
	}
}

func TestPlay(t *testing.T) {
	game, _ := NewGame()
	game.Deal(0, []*pb.PlayGameRequest_Bonus{})
	if err := game.Play(); err != nil {
		t.Fail()
	}
}

func TestFoldNotAllowed(t *testing.T) {
	game, _ := NewGame()
	game.Deal(0, []*pb.PlayGameRequest_Bonus{})
	game.Play()
	if err := game.Fold(); err == nil {
		t.Fail()
	}
}

func TestCheckNotAllowed(t *testing.T) {
	game, _ := NewGame()
	game.Deal(0, []*pb.PlayGameRequest_Bonus{})
	if err := game.Check(); err == nil {
		t.Fail()
	}
}

func TestCheck(t *testing.T) {
	game, _ := NewGame()
	game.Deal(0, []*pb.PlayGameRequest_Bonus{})
	game.Play()
	if err := game.Check(); err != nil {
		t.Fail()
	}
}

func TestAnyActionIfGameCompleted(t *testing.T) {
	game, _ := NewGame()
	game.Deal(0, []*pb.PlayGameRequest_Bonus{})
	game.Fold()
	if err := game.Fold(); err == nil {
		t.Fail()
	}
}

func TestBet(t *testing.T) {
	game, _ := NewGame()
	game.Deal(0, []*pb.PlayGameRequest_Bonus{})
	game.Play()
	if err := game.Bet(); err != nil {
		t.Fail()
	}
}

func TestBetNotAllowed(t *testing.T) {
	game, _ := NewGame()
	game.Deal(0, []*pb.PlayGameRequest_Bonus{})
	if err := game.Bet(); err == nil {
		t.Fail()
	}
}

func TestBetRiver(t *testing.T) {
	game, _ := NewGame()
	game.Deal(0, []*pb.PlayGameRequest_Bonus{})
	game.Play()
	if err := game.Bet(); err != nil {
		t.Fail()
	}
	if err := game.Bet(); err != nil {
		t.Fail()
	}
}

func TestGetEntity(t *testing.T) {
	game, _ := NewGame()
	entity, err := game.GetEntity()
	if entity == nil {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
}
