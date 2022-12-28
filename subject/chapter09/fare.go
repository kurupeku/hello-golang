package chapter09

import "github.com/kurupeku/hello-golang/helper"

type Fare struct {
	From string
	To   string
}

func (f *Fare) innerDistance() int {
	current := f.From
	next := ""
	distance := 0
	for current != f.To {
		next = helper.InnerNextStation(current)
		distance += helper.InnerLoopDistance(next)
		current = next
	}
	return distance
}

func (f *Fare) outerDistance() int {
	current := f.From
	next := ""
	distance := 0
	for current != f.To {
		next = helper.OuterNextStation(current)
		distance += helper.OuterLoopDistance(next)
		current = next
	}
	return distance
}

func (f *Fare) distance() int {
	id := f.innerDistance()
	od := f.outerDistance()
	switch {
	case id < od:
		return id
	case id > od:
		return od
	case id == od:
		return id
	default:
		return 0
	}
}

func (f *Fare) TicketCharge() int {
	charge := 0
	distance := f.distance()
	switch {
	case distance <= 3999:
		charge = 140
	case 4000 <= distance && distance < 6999:
		charge = 160
	case 7000 <= distance && distance < 10999:
		charge = 170
	case 11000 <= distance && distance < 15999:
		charge = 200
	case 16000 <= distance && distance < 20999:
		charge = 270
	case 21000 <= distance && distance < 25999:
		charge = 350
	case 26000 <= distance && distance < 30999:
		charge = 420
	case 31000 <= distance:
		charge = 490
	}
	return charge
}

func (f *Fare) CardCharge() int {
	charge := 0
	distance := f.distance()
	switch {
	case distance <= 3999:
		charge = 136
	case 4000 <= distance && distance < 6999:
		charge = 157
	case 7000 <= distance && distance < 10999:
		charge = 168
	case 11000 <= distance && distance < 15999:
		charge = 198
	case 16000 <= distance && distance < 20999:
		charge = 264
	case 21000 <= distance && distance < 25999:
		charge = 341
	case 26000 <= distance && distance < 30999:
		charge = 418
	case 31000 <= distance:
		charge = 484
	}
	return charge
}
