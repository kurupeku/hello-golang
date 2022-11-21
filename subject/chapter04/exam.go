package chapter04

type Card struct {
	Balance int
	Point   int
}

func Kaisatsu(charge int, card *Card) bool {
	// TODO: 実装

	//Balance(お金の残高)とPoint(Point残高)の合計がcharge(乗車料金)に満たない場合
	if charge > card.Balance+card.Point {
		//false(乗車不可)を返す
		return false
	}

	//Point(Point残高)がcharge(乗車料金)よりも多い場合
	if charge <= card.Point {
		//Point(Point残高)は、Point(Point残高)からcharge(乗車料金)を引いたものとなる
		card.Point = card.Point - charge
		//false(乗車可)を返す
		return true
	}

	//Point(Point残高)がcharge(乗車料金)よりも少ない場合
	if charge > card.Point {
		//最終的なBalance(お金の残高)は、()カッコ内の金額を引いたものとなる
		//()カッコ内の金額は、charge(乗車料金)からPoint(Point残高)を引いた残額となる
		card.Balance = card.Balance - (charge - card.Point)
		//Point(Point残高)は全額使い切るので、ゼロ円になる
		card.Point = 0
		//false(乗車可)を返す
	}
	return true
}
