package chapter04

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	// 足りない場合
	if charge > card.Point + card.Balance {
		return false;
	}

	// ポイントで払える場合
	if charge <= card.Point {
		card.Point -= charge
		return true;
	}

	// ポイントを使い切る
	charge -= card.Point
	card.Point = 0

	// 残りを現金で
	card.Balance -= charge
	return true
}
