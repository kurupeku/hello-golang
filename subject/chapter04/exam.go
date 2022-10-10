package chapter04

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	if card.Balance+card.Point < charge {
		return false
	}

	if card.Point > charge {
		card.Point -= charge
		return true
	}

	remain := charge - card.Point
	card.Balance -= remain
	card.Point = 0

	return true
}
