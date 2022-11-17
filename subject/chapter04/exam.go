package chapter04

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	// TODO: 実装
	var c = card
	// まずは通れるか通れないかの2パターン
	// 通れないパターン
	// falseで抜けるだけ
	if charge > (c.Point + c.Balance) {
		return false
	}
	// 通れるパターンは2通り
	if (charge - c.Point) > 0 {
		/// ポイントだけで払えないパターン処理
		//// chargeからpoint分を引いた値をBalanceから減算
		c.Balance = c.Balance + c.Point - charge
		//// Pointだけで払えないからPointは全部使う
		c.Point = 0
		return true
	} else {
		/// ポイントだけで払えるパターン処理
		//// Pointだけで払えるからPointからchargeを引く
		c.Point = c.Point - charge
		return true
	}
}
