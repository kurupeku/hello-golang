package chapter08

import (
	"math"
)

type Car interface {
	PricePer15Minutes() int
	MaxPrice() int
}

type Basic struct{}

type Middle struct{}

type Premium struct{}

func (g Basic) PricePer15Minutes() int {
	return 220
}

func (g Basic) MaxPrice() int {
	return 4290
}

func (g Middle) PricePer15Minutes() int {
	return 330
}

func (g Middle) MaxPrice() int {
	return 6490
}

func (g Premium) PricePer15Minutes() int {
	return 440
}

func (g Premium) MaxPrice() int {
	return 8690
}

func Calc(car Car, minutes int) int {
	// TODO: 実装

	overtime := minutes / 360
	overtime_charge := overtime * car.MaxPrice()
	remain_charge := int(math.Ceil(float64(minutes-(overtime*360))/15) * float64(car.PricePer15Minutes()))

	if car.MaxPrice() < remain_charge {
		remain_charge = car.MaxPrice()
	}

	return overtime_charge + remain_charge

	//hour := minutes / 360
	//if hour > 1 {
	//	return hour*car.MaxPrice() + math.Ceil(float64((minutes-hour*360)/15)) *car.PricePer15Minutes()
	//}
	//return 0

}
