package chapter04

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	if charge > (card.Balance + card.Point) {
		return false
	}

	if card.Point > charge {
		card.Point -= charge
	} else {
		card.Balance -= (charge - card.Point)
		card.Point = 0
	}
	return true
}
