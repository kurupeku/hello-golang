package chapter09

import (
	// "fmt"

	"github.com/kurupeku/hello-golang/helper"
)

type Fare struct {
	From string
	To   string
}

func (f *Fare) innerDistance() int {
	// 内回りの距離を返すように実装
	distance := 0
	currentStation := f.From
	for {
		if currentStation == f.To {
			break
		}
		currentStation = helper.InnerNextStation(currentStation)
		distance = distance + helper.InnerLoopDistance(currentStation)

	}
	return distance
}

func (f *Fare) outerDistance() int {
	// 外回りの距離を返すように実装
	distance := 0
	currentStation := f.From
	for {
		if currentStation == f.To {
			break
		}
		currentStation = helper.OuterNextStation(currentStation)
		distance = distance + helper.OuterLoopDistance(currentStation)
	}

	return distance
}

func (f *Fare) distance() int {
	// 内回りと外回りの距離を比較し、短い方を返すように実装
	distance := 0

	switch {
	case f.innerDistance() < f.outerDistance():
		distance = f.innerDistance()
	default:
		distance = f.outerDistance()
	}
	/*
		fmt.Printf("innerDistance is = %v \n", f.innerDistance())
		fmt.Printf("outerDistance is = %v \n", f.outerDistance())
		fmt.Printf("distance is = %v \n", distance)
	*/

	return distance
}

func (f *Fare) TicketCharge() int {
	// 切符の料金を返すように実装
	distance := f.distance()
	charge := 0

	switch {
	case distance == 0:
		charge = 0
	case distance < 4000:
		charge = 140
	case distance < 7000:
		charge = 160
	case distance < 11000:
		charge = 170
	case distance < 16000:
		charge = 200
	case distance < 21000:
		charge = 270
	case distance < 26000:
		charge = 350
	case distance < 31000:
		charge = 420
	default:
		charge = 490
	}

	return charge
}

func (f *Fare) CardCharge() int {
	// 切符の料金を返すように実装
	distance := f.distance()
	charge := 0

	switch {
	case distance == 0:
		charge = 0
	case distance < 4000:
		charge = 136
	case distance < 7000:
		charge = 157
	case distance < 11000:
		charge = 168
	case distance < 16000:
		charge = 198
	case distance < 21000:
		charge = 264
	case distance < 26000:
		charge = 341
	case distance < 31000:
		charge = 418
	default:
		charge = 484
	}

	return charge
}
