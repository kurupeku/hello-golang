package chapter04

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {

	if charge > card.Balance+card.Point {
		return false
	}

	if charge <= card.Point {
		card.Point -= charge
		return true
	}

	charge -= card.Point
	card.Point = 0
	card.Balance -= charge
	return true
}
