package chapter04

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	if (card.Balance + card.Point) < charge {
		return false
	}

	if charge <= card.Point {
		card.Point = card.Point - charge
		return true
	}

	card.Balance = card.Balance - (charge - card.Point)
	card.Point = 0
	return true
}
