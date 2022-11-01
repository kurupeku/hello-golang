package chapter04

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	// TODO: 実装
	if card.Balance+card.Point >= charge {
		tmp := card.Point - charge
		if tmp >= 0 {
			card.Point = tmp
		} else {
			card.Point = 0
			card.Balance += tmp
		}
		return true
	} else {
		return false
	}

}
