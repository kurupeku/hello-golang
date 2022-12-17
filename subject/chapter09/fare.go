package chapter09

import "github.com/kurupeku/hello-golang/helper"

type Fare struct {
	From string
	To   string
}

func (f *Fare) innerDistance() int {
	var distance int
	var current = f.From
	var next string

	for {
		if current == f.To {
			break
		}
		next = helper.InnerNextStation(current)
		distance += helper.InnerLoopDistance(next)
		current = next
	}
	return distance
}

func (f *Fare) outerDistance() int {
	var distance int
	var current = f.From
	var next string

	for {
		if current == f.To {
			break
		}
		next = helper.OuterNextStation(current)
		distance += helper.OuterLoopDistance(next)
		current = next
	}
	return distance
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
	// TODO: 切符の料金を返すように実装
	distance := f.distance()

	if distance <= 3999 {
		return 140
	} else if distance >= 4000 && distance <= 6999 {
		return 160
	} else if distance >= 7000 && distance <= 10999 {
		return 170
	} else if distance >= 11000 && distance <= 15999 {
		return 200
	} else if distance >= 16000 && distance <= 20999 {
		return 270
	} else if distance >= 21000 && distance <= 25999 {
		return 350
	} else if distance >= 26000 && distance <= 30999 {
		return 420
	} else if distance >= 31000 {
		return 490
	}
	return 0
}

func (f *Fare) CardCharge() int {
	// TODO: IC カードの料金を返すように実装
	distance := f.distance()

	if distance <= 3999 {
		return 136
	} else if distance >= 4000 && distance <= 6999 {
		return 157
	} else if distance >= 7000 && distance <= 10999 {
		return 168
	} else if distance >= 11000 && distance <= 15999 {
		return 198
	} else if distance >= 16000 && distance <= 20999 {
		return 264
	} else if distance >= 21000 && distance <= 25999 {
		return 341
	} else if distance >= 26000 && distance <= 30999 {
		return 418
	} else if distance >= 31000 {
		return 484
	}
	return 0
}
