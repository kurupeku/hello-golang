package chapter09

import (
	"github.com/kurupeku/hello-golang/helper"
)

type Fare struct {
	From string
	To   string
}

func (f *Fare) innerDistance() int {
	var distance int
	var nextStation string
	currentStation := f.From

	for f.To != currentStation {
		nextStation = helper.InnerNextStation(currentStation)
		distance += helper.InnerLoopDistance(nextStation)
		currentStation = nextStation
	}

	return distance
}

func (f *Fare) outerDistance() int {
	var distance int
	var nextStation string
	currentStation := f.From

	for f.To != currentStation {
		nextStation = helper.OuterNextStation(currentStation)
		distance += helper.OuterLoopDistance(nextStation)
		currentStation = nextStation
	}

	return distance
}

func (f *Fare) distance() int {
	if f.innerDistance() > f.outerDistance() {
		return f.outerDistance()
	}

	return f.innerDistance()
}

func (f *Fare) TicketCharge() int {
	var charge int

	switch {
	case f.distance() < 4000:
		charge = 140
	case 4000 <= f.distance() && f.distance() < 7000:
		charge = 160
	case 7000 <= f.distance() && f.distance() < 11000:
		charge = 170
	case 11000 <= f.distance() && f.distance() < 16000:
		charge = 200
	case 16000 <= f.distance() && f.distance() < 21000:
		charge = 270
	case 21000 <= f.distance() && f.distance() < 26000:
		charge = 350
	case 26000 <= f.distance() && f.distance() < 31000:
		charge = 420
	default:
		charge = 490
	}

	return charge
}

func (f *Fare) CardCharge() int {
	var charge int

	switch {
	case f.distance() < 4000:
		charge = 136
	case 4000 <= f.distance() && f.distance() < 7000:
		charge = 157
	case 7000 <= f.distance() && f.distance() < 11000:
		charge = 168
	case 11000 <= f.distance() && f.distance() < 16000:
		charge = 198
	case 16000 <= f.distance() && f.distance() < 21000:
		charge = 264
	case 21000 <= f.distance() && f.distance() < 26000:
		charge = 341
	case 26000 <= f.distance() && f.distance() < 31000:
		charge = 418
	default:
		charge = 484
	}

	return charge
}
