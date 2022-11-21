package chapter04

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	if s := card.Balance + card.Point; charge > s {
		return false
	}

	card.Point -= charge
	if card.Point < 0 {
		card.Balance += card.Point
		card.Point = 0
	}
	return true
}
