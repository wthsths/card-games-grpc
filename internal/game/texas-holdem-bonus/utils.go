package texasholdembonus

import (
	"fmt"
	"strings"

	"github.com/chehsunliu/poker"
)

const (
	maxStraight = 1609
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
