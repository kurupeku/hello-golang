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

func (b *Basic) PricePer15Minutes() int {
	return 220
}

func (b *Basic) MaxPrice() int {
	return 4290
}

func (m *Middle) PricePer15Minutes() int {
	return 330
}

func (m *Middle) MaxPrice() int {
	return 6490
}

func (p *Premium) PricePer15Minutes() int {
	return 440
}

func (p *Premium) MaxPrice() int {
	return 8690
}

func Calc(car Car, minutes int) int {
	MaxMin := 360
	MaxPrice := car.MaxPrice()
	fee := 0

	if minutes < MaxMin {
		fee = car.PricePer15Minutes() * int(math.Ceil(float64(minutes)/15.0))
		if fee > MaxPrice {
			fee = MaxPrice
		}
	} else {
		m := minutes - MaxMin
		if m == 0 {
			fee = MaxPrice
		} else {
			fee = MaxPrice
			fee += Calc(car, m)
		}
	}

	return fee
}
