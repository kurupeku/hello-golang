package chapter04

type Card struct {
	Balance int
	Point   int
}

// 運賃とIC カードを渡すと、カード内のチャージ残高とポイントで乗車可能かを判定する関数を実装してください
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
