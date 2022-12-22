package chapter09

import "github.com/kurupeku/hello-golang/helper"

type Fare struct {
	From string
	To   string
}

func (f *Fare) innerDistance() int {
	var distance int
	for station := f.From; station != f.To; station = helper.InnerNextStation(station) {
		distance += helper.InnerLoopDistance(station)
	}
	return distance
}

func (f *Fare) outerDistance() int {
	var distance int
	for station := f.From; station != f.To; station = helper.OuterNextStation(station) {
		distance += helper.OuterLoopDistance(station)
	}
	return distance
}

func (f *Fare) distance() int {
	if f.outerDistance() > f.innerDistance() {
		return f.innerDistance()
	} else {
		return f.outerDistance()
	}
}

func (f *Fare) TicketCharge() int {
	distance := f.distance()
	switch {
		case distance <= 3999:  return 140
		case distance <= 6999:  return 160
		case distance <= 10999: return 170
		case distance <= 15999: return 200
		case distance <= 20999: return 270
		case distance <= 25999: return 350
		case distance <= 30999: return 420
		default: return 490
	}
}

func (f *Fare) CardCharge() int {
	distance := f.distance()
	switch {
		case distance <= 3999:  return 136
		case distance <= 6999:  return 157
		case distance <= 10999: return 168
		case distance <= 15999: return 198
		case distance <= 20999: return 264
		case distance <= 25999: return 341
		case distance <= 30999: return 418
		default: return 484
	}
}
