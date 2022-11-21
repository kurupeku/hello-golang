package chapter04

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	// TODO: 実装
	now_sum_balance_and_point := card.Balance + card.Point
	num := now_sum_balance_and_point - charge
	switch {
	case num < 0:
		return false
	default:
		calc_point := card.Point - charge
		switch {
		case calc_point < 0:
			calc_balance := card.Balance + calc_point
			card.Balance = calc_balance
			card.Point = 0
			return true
		default:
			card.Point = calc_point
			return true

		}
	}
}
