package chapter09

import "github.com/kurupeku/hello-golang/helper"

type Fare struct {
	From string
	To   string
}

func (f *Fare) innerDistance() int {
	currentStation := f.From
	distance := 0
	for {
		nextStation := helper.InnerNextStation(currentStation)
		distance += helper.InnerLoopDistance(nextStation)
		currentStation = nextStation
		if currentStation == f.To {
			break
		}
	}

	return distance
}

func (f *Fare) outerDistance() int {
	currentStation := f.From
	distance := 0
	for {
		nextStation := helper.OuterNextStation(currentStation)
		distance += helper.OuterLoopDistance(nextStation)
		currentStation = nextStation
		if currentStation == f.To {
			break
		}
	}

	return distance
}

func (f *Fare) distance() int {
	inDis := f.innerDistance()
	outDis := f.outerDistance()
	if inDis > outDis {
		return outDis
	}
	return inDis
}

func (f *Fare) TicketCharge() int {
	var charge int
	dis := f.distance()
	switch {
	case dis < 4000:
		charge = 140
	case dis < 7000:
		charge = 160
	case dis < 11000:
		charge = 170
	case dis < 16000:
		charge = 200
	case dis < 21000:
		charge = 270
	case dis < 26000:
		charge = 350
	case dis < 31000:
		charge = 420
	case dis >= 31000:
		charge = 490
	}
	return charge
}

func (f *Fare) CardCharge() int {
	var charge int
	dis := f.distance()
	switch {
	case dis < 4000:
		charge = 136
	case dis < 7000:
		charge = 157
	case dis < 11000:
		charge = 168
	case dis < 16000:
		charge = 198
	case dis < 21000:
		charge = 264
	case dis < 26000:
		charge = 341
	case dis < 31000:
		charge = 418
	case dis >= 31000:
		charge = 484
	}
	return charge
}
