package chapter09

import "github.com/kurupeku/hello-golang/helper"

type Fare struct {
	From string
	To   string
}

func (f *Fare) innerDistance() int {
	// TODO: 内回りの料金を返すように実装

	nextStation := f.From
	distance := 0
	totalDistance := 0

	for {
		//次の駅を取得
		nextStation = helper.InnerNextStation(nextStation)
		//次の駅までの距離を取得
		distance = helper.InnerLoopDistance(nextStation)
		//距離を加算する
		totalDistance = totalDistance + distance
		//目的地であれば終了
		if nextStation == f.To {
			break
		}
	}
	return totalDistance
}

func (f *Fare) outerDistance() int {
	// TODO: 外回りの距離を返すように実装
	nextStation := f.From
	distance := 0
	totalDistance := 0

	for {
		//次の駅を取得
		nextStation = helper.OuterNextStation(nextStation)
		//次の駅までの距離を取得
		distance = helper.OuterLoopDistance(nextStation)
		//距離を加算する
		totalDistance = totalDistance + distance
		//目的地であれば終了
		if nextStation == f.To {
			break
		}
	}
	return totalDistance
}

func (f *Fare) distance() int {
	// 内回りと外回りの距離を比較し、短い方を返すように実装
	if f.innerDistance() > f.outerDistance() {
		return f.outerDistance()
	} else {
		return f.innerDistance()
	}
}

func (f *Fare) TicketCharge() int {
	// TODO: 切符の料金を返すように実装
	switch {
	case f.distance() <= 3999:
		return 140
	case f.distance() <= 6999:
		return 160
	case f.distance() <= 10999:
		return 170
	case f.distance() <= 15999:
		return 200
	case f.distance() <= 20999:
		return 270
	case f.distance() <= 25999:
		return 350
	case f.distance() <= 30999:
		return 420
	default:
		return 490
	}
}

func (f *Fare) CardCharge() int {
	// TODO: IC カードの料金を返すように実装
	switch {
	case f.distance() <= 3999:
		return 136
	case f.distance() <= 6999:
		return 157
	case f.distance() <= 10999:
		return 168
	case f.distance() <= 15999:
		return 198
	case f.distance() <= 20999:
		return 264
	case f.distance() <= 25999:
		return 341
	case f.distance() <= 30999:
		return 418
	default:
		return 484
	}
}
