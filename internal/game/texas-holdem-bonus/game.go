package texasholdembonus

import (
	"encoding/json"
	"errors"
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

	DealerHandRank       int32  `json:"dealer_hand_rank"`
	DealerHandRankString string `json:"dealer_hand_rank_string"`
	PlayerHandRank       int32  `json:"player_hand_rank"`
	PlayerHandRankString string `json:"player_hand_rank_string"`
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

	b, _ := json.Marshal(game.data)
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

	game.data.PlayerCards = append(game.data.PlayerCards, game.drawCard())
	game.data.PlayerCards = append(game.data.PlayerCards, game.drawCard())

	game.data.Bonus = 10
	game.data.Ante = 10

	game.update()

	return nil
}

func (game *TexasHoldemBonus) Play() error {
	if err := game.checkIfActionAllowed(pb.PlayGameRequest_ACTION_PLAY); err != nil {
		return err
	}

	game.data.BetFlop = game.data.Ante * 2

	for i := 0; i < 3; i++ {
		game.data.CommunityCards = append(game.data.CommunityCards, game.drawCard())
	}

	game.update()

	return nil
}

func (game *TexasHoldemBonus) Fold() error {
	if err := game.checkIfActionAllowed(pb.PlayGameRequest_ACTION_FOLD); err != nil {
		return err
	}

	game.data.Fold = true

	game.update()

	return nil
}

func (game *TexasHoldemBonus) Check() error {
	if err := game.checkIfActionAllowed(pb.PlayGameRequest_ACTION_CHECK); err != nil {
		return err
	}

	game.data.CommunityCards = append(game.data.CommunityCards, game.drawCard())
	game.update()

	return nil
}

func (game *TexasHoldemBonus) Bet() error {
	if err := game.checkIfActionAllowed(pb.PlayGameRequest_ACTION_BET); err != nil {
		return err
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

	if game.data.Fold || game.data.GameStatus == GameStatusRiver {
		game.data.GameStatus = GameStatusShowdown
	}

	if game.data.GameStatus == GameStatusShowdown {
		// Show all community cards
		for i := 0; i < 5-len(game.data.CommunityCards); i++ {
			game.data.CommunityCards = append(game.data.CommunityCards, game.drawCard())
		}

		if len(game.data.DealerCards) == 0 {
			game.drawDealerCards()
		}

		var dealerHand, playerHand []int

		dealerHand = append(dealerHand, game.data.DealerCards...)
		dealerHand = append(dealerHand, game.data.CommunityCards...)

		playerHand = append(playerHand, game.data.PlayerCards...)
		playerHand = append(playerHand, game.data.CommunityCards...)

		game.data.DealerHandRank = evaluateHand(dealerHand...)
		game.data.PlayerHandRank = evaluateHand(playerHand...)

		game.data.DealerHandRankString = getRankString(game.data.DealerHandRank)
		game.data.PlayerHandRankString = getRankString(game.data.PlayerHandRank)

		if !game.data.Fold &&
			game.data.PlayerHandRank < game.data.DealerHandRank {
			game.data.FlopWin = game.data.BetFlop
			game.data.TurnWin = game.data.BetTurn
			game.data.RiverWin = game.data.BetRiver
		}

		if isStraightOrBetter(game.data.PlayerHandRank) {
			game.data.AnteWin = game.data.Ante
		}

		if isAcePair(game.data.DealerCards...) && isAcePair(game.data.PlayerCards...) {
			game.data.BonusWin = game.data.Bonus * 1000
		} else if isAcePair(game.data.PlayerCards...) {
			game.data.BonusWin = game.data.Bonus * 30
		} else if isAceAndKing(game.data.PlayerCards...) && isSuited(game.data.PlayerCards...) {
			game.data.BonusWin = game.data.Bonus * 25
		} else if isAceAndQueen(game.data.PlayerCards...) && isSuited(game.data.PlayerCards...) {
			game.data.BonusWin = game.data.Bonus * 20
		} else if isAceAndJack(game.data.PlayerCards...) && isSuited(game.data.PlayerCards...) {
			game.data.BonusWin = game.data.Bonus * 20
		} else if isAceAndKing(game.data.PlayerCards...) {
			game.data.BonusWin = game.data.Bonus * 15
		} else if isKingPair(game.data.PlayerCards...) || isQueenPair(game.data.PlayerCards...) || isJackPair(game.data.PlayerCards...) {
			game.data.BonusWin = game.data.Bonus * 10
		} else if isPair(game.data.PlayerCards...) {
			game.data.BonusWin = game.data.Bonus * 3
		}
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

func (game *TexasHoldemBonus) drawDealerCards() {
	game.data.DealerCards = append(game.data.DealerCards, game.drawCard())
	game.data.DealerCards = append(game.data.DealerCards, game.drawCard())
}

func (game *TexasHoldemBonus) checkIfActionAllowed(action pb.PlayGameRequest_Action) error {
	if game.data.Fold || game.data.GameStatus > GameStatusTurn {
		return ErrGameHasCompleted
	}

	if !game.isAllowedAction(action) {
		return ErrActionNotAllowed
	}

	return nil
}
