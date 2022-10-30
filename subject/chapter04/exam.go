package chapter04

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	if charge > (card.Balance + card.Point) {
		return false
	} else {
		if card.Point >= charge {
			card.Point -= charge
		} else {
			remain := charge - card.Point
			card.Point = 0
			card.Balance -= remain
		}
		return true
	}
}
