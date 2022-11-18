package chapter04

// "fmt"

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	// TODO: 実装
	var total int
	var tmp int

	// ICカード内の総残高算出
	total = card.Balance + card.Point

	// 1 そもそもICカード残高より乗車料金が高いかどうか
	// 1a乗車料金が残高以内の場合
	if total-charge >= 0 {
		// 2a ICカード内のポイントに余りがある場合
		if card.Point > 0 {
			// 3 ポイントを乗車料金分差し引く
			tmp = card.Point - charge
			// 3a ポイントから差し引いた後、ポイントがマイナスになる場合
			if tmp < 0 {
				// カードポイントを0にする
				card.Point = 0
				// カード残高から不足分を減算処理する
				card.Balance = card.Balance + tmp
				return true
				// 3b ポイントから差し引いた後、ポイントが0以上になる場合
			} else {
				card.Point = tmp
				return true
			}
			// 2b ICカード内のポイント残高が0の場合
		} else {
			card.Balance = card.Balance - charge
			return true
		}
	} else {
		// 1b 乗車料金がICカード残が会場の場合（残高不足の場合）
		return false
	}
}
