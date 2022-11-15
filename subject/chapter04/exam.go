package chapter04

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {

	if charge > (card.Point + card.Balance) {
		return false
	}

	if charge > card.Point {
		card.Balance = card.Balance - charge + card.Point
		card.Point = 0
	} else {
		card.Point -= charge
	}

	// TODO: 実装
	return true
}
