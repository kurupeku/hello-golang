package chapter04

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	// TODO: 実装

	// 支払い残高
	paymentBalance := charge

	// 残高とポイントの合計が乗車料金以下の場合は乗車できない
	if card.Balance+card.Point < charge {
		return false
	}

	// ポイントから優先的に支払う
	if card.Point >= charge {
		card.Point -= charge
	} else {
		paymentBalance -= card.Point
		card.Point = 0
		card.Balance -= paymentBalance
	}

	return true
}
