package chapter09

import "github.com/kurupeku/hello-golang/helper"

type Fare struct {
	From string
	To   string
}

func (f *Fare) innerDistance() int {
	// TODO: 内回りの距離を返すように実装

	current_station := f.From
	distance := 0

	if f.From == f.To {
		return 0
	}

	for {
		next := helper.InnerNextStation(current_station)
		distance += helper.InnerLoopDistance(next)
		current_station = next

		if current_station == f.To {
			break
		}
	}

	return distance
}

func (f *Fare) outerDistance() int {
	// TODO: 外回りの距離を返すように実装

	current_station := f.From
	distance := 0

	if f.From == f.To {
		return 0
	}

	for {
		next := helper.OuterNextStation(current_station)
		distance += helper.OuterLoopDistance(next)
		current_station = next

		if current_station == f.To {
			break
		}
	}

	return distance
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

	if f.distance() == 0 {
		return 0
	} else if f.distance() <= 3999 {
		return 140
	} else if f.distance() <= 6999 {
		return 160
	} else if f.distance() <= 10999 {
		return 170
	} else if f.distance() <= 15999 {
		return 200
	} else if f.distance() <= 20999 {
		return 270
	} else if f.distance() <= 25999 {
		return 350
	} else if f.distance() <= 30999 {
		return 420
	} else {
		return 490
	}

}

func (f *Fare) CardCharge() int {
	// TODO: IC カードの料金を返すように実装

	if f.distance() == 0 {
		return 0
	} else if f.distance() <= 3999 {
		return 136
	} else if f.distance() <= 6999 {
		return 157
	} else if f.distance() <= 10999 {
		return 168
	} else if f.distance() <= 15999 {
		return 198
	} else if f.distance() <= 20999 {
		return 264
	} else if f.distance() <= 25999 {
		return 341
	} else if f.distance() <= 30999 {
		return 418
	} else {
		return 484
	}

}
