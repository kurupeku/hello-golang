package chapter04

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	// TODO: 実装
	var c = card
	// まずは通れるか通れないかの2パターン
	// 通れないパターン(ポイント＋残高をchargeが超えている場合)
	// falseで抜けるだけ
	if charge > (c.Point + c.Balance) {
		return false
	}
	// 通れるパターンは2通り
	// ポイントだけで通れる場合
	// ポイントと残高で通れる場合
	if charge < c.Point {
		/// ポイントだけで払えるパターン処理
		/// Pointからchargeを引く
		c.Point = c.Point - charge
		return true
	} else {
		/// ポイントだけで払えないパターン処理
		/// 残高＋ポイント－charge
		c.Balance = c.Balance + c.Point - charge
		/// ポイントは全部使うから0
		c.Point = 0
		return true
	}
}
