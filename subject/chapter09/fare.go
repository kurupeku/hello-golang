package chapter09

import "github.com/kurupeku/hello-golang/helper"

type Fare struct {
	From string
	To   string
}

func (f *Fare) innerDistance() int {
	var totalDistance int
	var station = f.From
	var nextStation string

	for {
		if station == f.To {
			break
		}
		nextStation = helper.InnerNextStation(station)
		totalDistance += helper.InnerLoopDistance(nextStation)
		station = nextStation
	}
	return totalDistance
}

func (f *Fare) outerDistance() int {
	var totalDistance int
	var station = f.From
	var nextStation string

	for {
		if station == f.To {
			break
		}
		nextStation = helper.InnerNextStation(station)
		totalDistance += helper.InnerLoopDistance(nextStation)
		station = nextStation
	}
	return totalDistance
}

func (f *Fare) distance() int {
	inner := f.innerDistance()
	outer := f.outerDistance()
	if inner <= outer {
		return inner
	} else {
		return outer
	}
}

func (f *Fare) TicketCharge() int {
	totalDistance := f.distance()
	switch {
	case totalDistance == 0:
		return 0
	case totalDistance < 4000:
		return 140
	case totalDistance < 7000:
		return 160
	case totalDistance < 11000:
		return 170
	case totalDistance < 16000:
		return 200
	case totalDistance < 21000:
		return 270
	case totalDistance < 26000:
		return 350
	case totalDistance < 31000:
		return 420
	default:
		return 490
	}
}

func (f *Fare) CardCharge() int {
	totalDistance := f.distance()
	switch {
	case totalDistance == 0:
		return 0
	case totalDistance < 4000:
		return 136
	case totalDistance < 7000:
		return 157
	case totalDistance < 11000:
		return 168
	case totalDistance < 16000:
		return 198
	case totalDistance < 21000:
		return 264
	case totalDistance < 26000:
		return 341
	case totalDistance < 31000:
		return 418
	default:
		return 484
	}
}
