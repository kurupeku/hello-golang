package chapter09

import (
	"github.com/kurupeku/hello-golang/helper"
)

type Fare struct {
	From string
	To   string
}

func (f *Fare) innerDistance() int {
	// TODO: 内回りの料金を返すように実装
	dist := 0
	nowstation := f.From
	for {
		nextstation := helper.InnerNextStation(nowstation)
		dist += helper.InnerLoopDistance(nextstation)
		nowstation = nextstation
		if nextstation == f.To {
			break
		}
	}
	return dist
}

func (f *Fare) outerDistance() int {
	// TODO: 外回りの距離を返すように実装
	dist := 0
	nowstation := f.From
	for {
		nextstation := helper.OuterNextStation(nowstation)
		dist += helper.OuterLoopDistance(nextstation)
		nowstation = nextstation
		if nextstation == f.To {
			break
		}
	}
	return dist
}

func (f *Fare) distance() int {
	// 内回りと外回りの距離を比較し、短い方を返すように実装
	if f.innerDistance() <= f.outerDistance() {
		return f.innerDistance()
	} else {
		return f.outerDistance()
	}
}

func (f *Fare) TicketCharge() int {
	// TODO: 切符の料金を返すように実装
	var charge int
	switch {
	case f.distance() < 4000:
		charge = 140
	case f.distance() < 7000:
		charge = 160
	case f.distance() < 11000:
		charge = 170
	case f.distance() < 16000:
		charge = 200
	case f.distance() < 21000:
		charge = 270
	case f.distance() < 26000:
		charge = 350
	case f.distance() < 31000:
		charge = 420
	case f.distance() >= 31000:
		charge = 490
	}
	return charge
}

func (f *Fare) CardCharge() int {
	// TODO: IC カードの料金を返すように実装
	var charge int
	switch {
	case f.distance() < 4000:
		charge = 136
	case f.distance() < 7000:
		charge = 157
	case f.distance() < 11000:
		charge = 168
	case f.distance() < 16000:
		charge = 198
	case f.distance() < 21000:
		charge = 264
	case f.distance() < 26000:
		charge = 341
	case f.distance() < 31000:
		charge = 418
	case f.distance() >= 31000:
		charge = 484
	}
	return charge
}
