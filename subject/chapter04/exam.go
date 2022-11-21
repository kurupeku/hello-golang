package chapter04

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	// TODO: 実装
	/*
		    料金 残高 ポイント
		    if Point >= 料金
			Point = Point - 料金
			else
			Point = 0
			Balance = Balance - (料金 - Point)
		    return true
	*/

	if card.Balance+card.Point >= charge {
		if card.Point >= charge {
			card.Point = card.Point - charge
			return true
		} else {
			card.Balance = card.Balance - (charge - card.Point)
			card.Point = 0
			return true
		}

	} else {
		return false
	}

}
