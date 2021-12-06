package texasholdembonus

import (
	"testing"
)

func TestGetCardById(t *testing.T) {
	cardIds := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 13, 14, 15, 27, 28, 51}
	expects := []string{"2d", "3d", "4d", "5d", "6d", "7d", "8d", "9d", "Td", "2h", "3h", "4h", "3s", "4s", "Ac"}

	for i := 0; i < len(cardIds); i++ {
		r := getCardById(cardIds[i])
		if r != expects[i] {
			t.Errorf("expected %v got %v", expects[i], r)
		}
	}
}

func TestGetHandCards(t *testing.T) {
	cards := []int{0, 1, 2}
	expect := "2d 3d 4d"

	result := getHandCards(cards...)
	if expect != result {
		t.Errorf("expected %s got %s", expect, result)
	}
}
