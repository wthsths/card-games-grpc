package texasholdembonus

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/wthsths/minigames/pb"
)

type TexasHoldemBonus struct {
	entity *pb.Game
	data   *EntityData
}

const (
	GameStatusDealt = iota
	GameStatusPreflop
	GameStatusFlop
	GameStatusTurn
	GameStatusRiver
	GameStatusShowdown
)

const (
	GameResultUncompleted = iota
	GameResultDraw
	GameResultPlayerWon
	GameResultDealerWon
)

type EntityData struct {
	Deck           []int                       `json:"deck"`
	DealerCards    []int                       `json:"dealer_cards"`
	PlayerCards    []int                       `json:"player_cards"`
	CommunityCards []int                       `json:"community_cards"`
	Ante           uint64                      `json:"ante"`
	Bonus          uint64                      `json:"bonus"`
	BetFlop        uint64                      `json:"bet_flop"`
	BetTurn        uint64                      `json:"bet_turn"`
	BetRiver       uint64                      `json:"bet_river"`
	FlopWin        uint64                      `json:"flop_win"`
	TurnWin        uint64                      `json:"turn_win"`
	RiverWin       uint64                      `json:"river_win"`
	AnteWin        uint64                      `json:"ante_win"`
	BonusWin       uint64                      `json:"bonus_win"`
	TotalBet       uint64                      `json:"total_bet"`
	GameStatus     int                         `json:"game_status"`
	AllowedActions []pb.PlayGameRequest_Action `json:"allowed_actions"`
	Fold           bool                        `json:"fold"`
	GameResult     int                         `json:"game_result`
}

var (
	ErrInvalidGameType           = errors.New("invalid game type")
	ErrGameHasStartedOrCompleted = errors.New("game has already started or completed")
	ErrGameHasCompleted          = errors.New("the game has already completed")
	ErrActionNotAllowed          = errors.New("action is not allowed")
)

func NewGame() (*TexasHoldemBonus, error) {
	uuid, _ := uuid.NewUUID()

	game := &TexasHoldemBonus{
		entity: &pb.Game{
			Uuid: uuid.String(),
			Type: pb.GameType_GAME_TYPE_TEXAS_HOLDEM_BONUS,
		},
		data: &EntityData{},
	}

	for i := 0; i < 52; i++ {
		game.data.Deck = append(game.data.Deck, i)
	}

	b, err := json.Marshal(game.data)
	if err != nil {
		return nil, err
	}
	game.entity.Data = b

	return game, nil
}

func NewGameWithEntity(entity *pb.Game) (*TexasHoldemBonus, error) {
	if entity.Type != pb.GameType_GAME_TYPE_TEXAS_HOLDEM_BONUS {
		return nil, ErrInvalidGameType
	}

	game := &TexasHoldemBonus{
		entity: entity,
	}

	err := json.Unmarshal(entity.Data, &game.data)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func (game *TexasHoldemBonus) SetPlayer(id uint64) error {
	game.entity.Player = id
	return nil
}

func (game *TexasHoldemBonus) Deal() error {
	if game.entity.Status != pb.GameStatus_GAME_STATUS_CREATED {
		return ErrGameHasStartedOrCompleted
	}

	game.entity.Status = pb.GameStatus_GAME_STATUS_PLAYING

	game.shuffleDeck()

	game.data.DealerCards = append(game.data.DealerCards, game.drawCard())
	game.data.DealerCards = append(game.data.DealerCards, game.drawCard())

	game.data.PlayerCards = append(game.data.PlayerCards, game.drawCard())
	game.data.PlayerCards = append(game.data.PlayerCards, game.drawCard())

	game.data.Bonus = 0
	game.data.Ante = 10

	game.update()

	return nil
}

func (game *TexasHoldemBonus) Play() error {
	if game.data.Fold || game.data.GameStatus > GameStatusTurn {
		return ErrGameHasCompleted
	}

	if !game.isAllowedAction(pb.PlayGameRequest_ACTION_PLAY) {
		return ErrActionNotAllowed
	}

	game.data.BetFlop = game.data.Ante * 2

	for i := 0; i < 3; i++ {
		game.data.CommunityCards = append(game.data.CommunityCards, game.drawCard())
	}

	game.update()

	return nil
}

func (game *TexasHoldemBonus) Fold() error {
	if game.data.Fold || game.data.GameStatus > GameStatusTurn {
		return ErrGameHasCompleted
	}

	if !game.isAllowedAction(pb.PlayGameRequest_ACTION_FOLD) {
		return ErrActionNotAllowed
	}

	game.data.Fold = true

	game.update()

	return nil
}

func (game *TexasHoldemBonus) Check() error {
	if game.data.Fold || game.data.GameStatus > GameStatusTurn {
		return ErrGameHasCompleted
	}

	if !game.isAllowedAction(pb.PlayGameRequest_ACTION_CHECK) {
		return ErrActionNotAllowed
	}

	game.data.CommunityCards = append(game.data.CommunityCards, game.drawCard())
	game.update()

	return nil
}

func (game *TexasHoldemBonus) Bet() error {
	if game.data.Fold || game.data.GameStatus > GameStatusTurn {
		return ErrGameHasCompleted
	}

	if !game.isAllowedAction(pb.PlayGameRequest_ACTION_CHECK) {
		return ErrActionNotAllowed
	}

	if game.data.GameStatus == GameStatusFlop {
		game.data.BetTurn = game.data.Ante
	} else if game.data.GameStatus == GameStatusTurn {
		game.data.BetRiver = game.data.Ante
	}

	game.data.CommunityCards = append(game.data.CommunityCards, game.drawCard())
	game.update()

	return nil
}

func (game *TexasHoldemBonus) GetEntity() (*pb.Game, error) {
	return game.entity, nil
}

func (game *TexasHoldemBonus) shuffleDeck() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(game.data.Deck), func(i, j int) { game.data.Deck[i], game.data.Deck[j] = game.data.Deck[j], game.data.Deck[i] })
}

func (game *TexasHoldemBonus) drawCard() int {
	f := len(game.data.Deck)
	rv := (game.data.Deck)[f-1]
	game.data.Deck = (game.data.Deck)[:f-1]
	return rv
}

func (game *TexasHoldemBonus) update() {
	game.data.TotalBet = game.data.BetFlop + game.data.BetTurn + game.data.BetRiver

	game.data.AllowedActions = nil

	if game.data.GameStatus == GameStatusDealt {
		game.data.AllowedActions = append(game.data.AllowedActions, pb.PlayGameRequest_ACTION_PLAY)
		game.data.AllowedActions = append(game.data.AllowedActions, pb.PlayGameRequest_ACTION_FOLD)
	} else if game.data.GameStatus < GameStatusRiver {
		game.data.AllowedActions = append(game.data.AllowedActions, pb.PlayGameRequest_ACTION_BET)
		game.data.AllowedActions = append(game.data.AllowedActions, pb.PlayGameRequest_ACTION_CHECK)
	}

	game.data.GameStatus++

	if game.data.GameStatus == GameStatusRiver {
		game.data.GameStatus = GameStatusShowdown
		game.data.AllowedActions = nil

		var dealerHand []int
		var playerHand []int

		dealerHand = append(dealerHand, game.data.DealerCards...)
		dealerHand = append(dealerHand, game.data.CommunityCards...)

		playerHand = append(playerHand, game.data.PlayerCards...)
		playerHand = append(playerHand, game.data.CommunityCards...)

		dealerRank := evaluateHand(dealerHand...)
		playerRank := evaluateHand(playerHand...)

		game.data.GameResult = GameResultDraw

		if playerRank < dealerRank {
			game.data.FlopWin = game.data.BetFlop
			game.data.TurnWin = game.data.BetTurn
			game.data.RiverWin = game.data.BetRiver

			game.data.GameResult = GameResultPlayerWon
		} else if playerRank > dealerRank {
			game.data.GameResult = GameResultDealerWon
		}

		if isStraightOrBetter(playerRank) {
			game.data.AnteWin = game.data.Ante
		}

		log.Printf("%v ===== %v", getHandCards(dealerHand...), getHandCards(playerHand...))
		log.Printf("%v (%s) === %v (%s) \n", dealerRank, getRankString(dealerRank), playerRank, getRankString(playerRank))
		log.Printf("comparing hands %+v\n", game.data)
	}

	b, _ := json.Marshal(game.data)
	game.entity.Data = b
}

func (game *TexasHoldemBonus) isAllowedAction(action pb.PlayGameRequest_Action) bool {
	for _, a := range game.data.AllowedActions {
		if a == action {
			return true
		}
	}

	return false
}
