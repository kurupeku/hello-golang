package chapter04

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	// TODO: 実装

	if (card.Balance + card.Point) < charge {
		return false
	}

	if card.Point >= charge {
		card.Point -= charge
	} else {
		card.Balance = card.Balance + card.Point - charge
		card.Point = 0
	}

	return true
}
