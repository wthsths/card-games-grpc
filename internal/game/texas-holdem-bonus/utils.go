package texasholdembonus

import (
	"fmt"
	"strings"

	"github.com/chehsunliu/poker"
)

const (
	maxStraight = 1609
)

const (
	cardValueTwo = iota
	cardValueThree
	cardValueFour
	cardValueFive
	cardValueSix
	cardValueSeven
	cardValueEight
	cardValueNine
	cardValueTen
	cardValueJack
	cardValueQueen
	cardValueKing
	cardValueAce
)

var cardValues []string = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
var cardSUits []string = []string{"d", "h", "s", "c"}

func getCardById(id int) string {
	value := int(id % 13)
	suits := int(id / 13)

	return fmt.Sprintf("%s%s", cardValues[value], cardSUits[suits])
}

func getHandCards(cards ...int) (res string) {
	for _, c := range cards {
		res += fmt.Sprintf("%s ", getCardById(c))
	}
	return strings.TrimRight(res, " ")
}

func getPokerCard(card string) poker.Card {
	return poker.NewCard(card)
}

func evaluateHand(cards ...int) int32 {
	var hand []poker.Card

	for _, c := range cards {
		hand = append(hand, getPokerCard(getCardById(c)))
	}

	rank := poker.Evaluate(hand)
	return rank
}

func getRankString(rank int32) string {
	return poker.RankString(rank)
}

func isStraightOrBetter(rank int32) bool {
	return rank <= maxStraight
}

func isAcePair(cards ...int) bool {
	if cards[0]%13 == cardValueAce && cards[1]%13 == cardValueAce {
		return true
	}
	return false
}

func isKingPair(cards ...int) bool {
	if cards[0]%13 == cardValueKing && cards[1]%13 == cardValueKing {
		return true
	}
	return false
}

func isQueenPair(cards ...int) bool {
	if cards[0]%13 == cardValueQueen && cards[1]%13 == cardValueQueen {
		return true
	}
	return false
}

func isJackPair(cards ...int) bool {
	if cards[0]%13 == cardValueJack && cards[1]%13 == cardValueJack {
		return true
	}
	return false
}

func isAceAndKing(cards ...int) bool {
	if cards[0]%13 == cardValueAce && cards[1]%13 == cardValueKing {
		return true
	}

	if cards[1]%13 == cardValueAce && cards[0]%13 == cardValueKing {
		return true
	}

	return false
}

func isSuited(cards ...int) bool {
	return cards[0]/13 == cards[1]/13
}

func isPair(cards ...int) bool {
	return cards[0]%13 == cards[1]%13
}

func isAceAndQueen(cards ...int) bool {
	if cards[0]%13 == cardValueAce && cards[1]%13 == cardValueQueen {
		return true
	}

	if cards[1]%13 == cardValueAce && cards[0]%13 == cardValueQueen {
		return true
	}

	return false
}

func isAceAndJack(cards ...int) bool {
	if cards[0]%13 == cardValueAce && cards[1]%13 == cardValueJack {
		return true
	}

	if cards[1]%13 == cardValueAce && cards[0]%13 == cardValueJack {
		return true
	}

	return false
}
