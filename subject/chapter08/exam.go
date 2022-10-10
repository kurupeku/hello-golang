package chapter08

import (
	"math"
)

type Car interface {
	PricePer15Minutes() int
	MaxPrice() int
}

type Basic struct{}

func (b *Basic) PricePer15Minutes() int {
	return 220
}

func (b *Basic) MaxPrice() int {
	return 4290
}

type Middle struct{}

func (m *Middle) PricePer15Minutes() int {
	return 330
}

func (m *Middle) MaxPrice() int {
	return 6490
}

type Premium struct{}

func (p *Premium) PricePer15Minutes() int {
	return 440
}

func (p *Premium) MaxPrice() int {
	return 8690
}

func Calc(car Car, minutes int) int {
	p360 := (minutes / 360) * car.MaxPrice()

	sur := float64(minutes % 360)
	t := math.Ceil(sur / 15)
	p15 := int(t) * car.PricePer15Minutes()

	if p15 > car.MaxPrice() {
		p15 = car.MaxPrice()
	}

	return p360 + p15
}
