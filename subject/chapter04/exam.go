package chapter04

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	// TODO: 実装

	if charge > card.Balance+card.Point {
		return false
	}

	if charge <= card.Point {
		card.Point -= charge
	} else {
		charge -= card.Point
		card.Point = 0
		card.Balance -= charge
	}

	return true
}
