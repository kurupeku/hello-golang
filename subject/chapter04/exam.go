package chapter04

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	// TODO: 実装
	// 合計値を算出
	total := card.Balance + card.Point

	// 合計値が乗車料金より高い場合実行最終的にtrueを返す
	if total >= charge {
		// ポイントから優先的に消費
		card.Point -= charge
		// ポイントがマイナス(金額不足)の場合実行
		if card.Point < 0 {
			// 現金から足りない数を減算(ポイントは負数)
			card.Balance += card.Point
			// ポイントを実際の下限である0のリセット
			card.Point = 0
		}
		return true
	}

	return false
}
