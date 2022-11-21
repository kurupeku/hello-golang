package chapter04

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	// TODO: 実装
	// 残高+ポイントが乗車料金より多いかの判定
	if charge <= card.Balance+card.Point {
		// 乗車料金が0より大きい かつ ポイントが0より大きい場合にループ
		for charge > 0 && card.Point > 0 {
			// 乗車料金とポイントをループする度に1マイナスする
			charge--
			card.Point--
		}
		// ポイントから引けなかった分の乗車料金を残高からマイナスする（ポイントで足りた場合は0を引く）
		card.Balance -= charge
		return true
	}
	return false
}
