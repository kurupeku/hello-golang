package chapter09

import "github.com/kurupeku/hello-golang/helper"

type Fare struct {
	From string
	To   string
}

func (f *Fare) innerDistance() int {
	// TODO: 内回りの距離を返すように実装
	distance := 0
	current_station := f.From
	for {
		if current_station == f.To {
			break
		}
		current_station = helper.InnerNextStation(current_station)
		distance += helper.InnerLoopDistance(current_station)
	}
	return distance
}

func (f *Fare) outerDistance() int {
	// TODO: 外回りの距離を返すように実装
	distance := 0
	current_station := f.From
	for {
		if current_station == f.To {
			break
		}
		current_station = helper.OuterNextStation(current_station)
		distance += helper.OuterLoopDistance(current_station)
	}
	return distance
}

func (f *Fare) distance() int {
	// 内回りと外回りの距離を比較し、短い方を返すように実装
	if f.innerDistance() > f.outerDistance() {
		return f.outerDistance()
	}
	return f.innerDistance()
}

func (f *Fare) TicketCharge() int {
	// TODO: 切符の料金を返すように実装
	distance := f.distance()
	if distance < 4000 {
		return 140
	}
	if distance < 7000 {
		return 160
	}
	if distance < 11000 {
		return 170
	}
	if distance < 16000 {
		return 200
	}
	if distance < 21000 {
		return 270
	}
	if distance < 26000 {
		return 350
	}
	if distance < 31000 {
		return 420
	}
	return 490
}

func (f *Fare) CardCharge() int {
	// TODO: IC カードの料金を返すように実装
	distance := f.distance()
	if distance < 4000 {
		return 136
	}
	if distance < 7000 {
		return 157
	}
	if distance < 11000 {
		return 168
	}
	if distance < 16000 {
		return 198
	}
	if distance < 21000 {
		return 264
	}
	if distance < 26000 {
		return 341
	}
	if distance < 31000 {
		return 418
	}
	return 484
}
