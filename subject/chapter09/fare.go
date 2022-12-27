package chapter09

import (
	"github.com/kurupeku/hello-golang/helper"
)

type Fare struct {
	From string
	To   string
}

func (f *Fare) innerDistance() int {
	if f.From == f.To {
		return 0
	}

	distance := 0
	ns := f.From

	for {
		ns = helper.InnerNextStation(ns)
		distance += helper.InnerLoopDistance(ns)
		if ns == f.To {
			break
		}
	}

	return distance
}

func (f *Fare) outerDistance() int {
	if f.From == f.To {
		return 0
	}

	distance := 0
	ns := f.From

	for {
		ns = helper.OuterNextStation(ns)
		distance += helper.OuterLoopDistance(ns)
		if ns == f.To {
			break
		}
	}

	return distance
}

func (f *Fare) distance() int {
	distance := 0
	i := f.innerDistance()
	o := f.outerDistance()

	if i < o {
		distance = i
	} else {
		distance = o
	}

	return distance
}

func (f *Fare) TicketCharge() int {
	fare := 0
	distance := f.distance()

	if distance < 4000 {
		fare = 140
	} else if distance < 7000 {
		fare = 160
	} else if distance < 11000 {
		fare = 170
	} else if distance < 16000 {
		fare = 200
	} else if distance < 21000 {
		fare = 270
	} else if distance < 26000 {
		fare = 350
	} else if distance < 31000 {
		fare = 420
	} else {
		fare = 490
	}

	return fare
}

func (f *Fare) CardCharge() int {
	fare := 0
	distance := f.distance()

	if distance < 4000 {
		fare = 136
	} else if distance < 7000 {
		fare = 157
	} else if distance < 11000 {
		fare = 168
	} else if distance < 16000 {
		fare = 198
	} else if distance < 21000 {
		fare = 264
	} else if distance < 26000 {
		fare = 341
	} else if distance < 31000 {
		fare = 418
	} else {
		fare = 484
	}

	return fare
}
