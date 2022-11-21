package chapter04

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	// TODO: 実装
	if charge <= card.Balance+card.Point {
		for charge > 0 && card.Point > 0 {
			card.Point--
			charge--
		}
		card.Balance -= charge
		return true
	}
	return false
}
