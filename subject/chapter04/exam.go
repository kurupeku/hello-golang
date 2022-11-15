package chapter04

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	// TODO: 実装
	if charge > (card.Point + card.Balance) {
		return false
	} else {
		if charge > card.Point {
			charge -= card.Point
			card.Balance -= charge
			card.Point = 0
		} else {
			card.Point -= charge
		}
		return true
	}
}
