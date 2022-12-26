package chapter08

import "math"

type Car interface {
	PricePer15Minutes() int
	MaxPrice() int
}

type Basic struct{}

func (basic Basic) PricePer15Minutes() int {
	return 220
}

func (basic Basic) MaxPrice() int {
	return 4290
}

type Middle struct{}

func (middle Middle) PricePer15Minutes() int {
	return 330
}

func (middle Middle) MaxPrice() int {
	return 6490
}

type Premium struct{}

func (premium Premium) PricePer15Minutes() int {
	return 440
}

func (premium Premium) MaxPrice() int {
	return 8690
}

func Calc(car Car, minutes int) int {
	resetCount := minutes / 360
	pricePer6Reset := car.MaxPrice() * resetCount
	pricePer15Minutes := car.PricePer15Minutes() * int(math.Ceil(float64(minutes%360)/15))

	if pricePer15Minutes >= car.MaxPrice() {
		return pricePer6Reset + car.MaxPrice()
	}

	return pricePer6Reset + pricePer15Minutes
}
