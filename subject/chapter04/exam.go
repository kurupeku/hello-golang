package chapter04

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	// TODO: 実装
	sum := card.Balance + card.Point // 今の合計
	zandaka := sum - charge // 運賃引いたカード残高
	switch {
	case zandaka < 0:
		return false
	default:
		x := card.Point - charge // ポイントから運賃引いた額
		switch {
		case x < 0:
			y := card.Balance + x
			card.Balance = y
			card.Point = 0
			return true
		default:
			card.Point = x
			return true

		}
	}
}