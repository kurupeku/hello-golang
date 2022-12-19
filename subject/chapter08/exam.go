package chapter08

import "math"

type Car interface {
	PricePer15Minutes() int
	MaxPrice() int
}

type Basic struct{}

type Middle struct{}

type Premium struct{}

func (b Basic) PricePer15Minutes() int {
	return 220
}

func (m Middle) PricePer15Minutes() int {
	return 330
}

func (p Premium) PricePer15Minutes() int {
	return 440
}

func (b Basic) MaxPrice() int {
	return 4290
}

func (m Middle) MaxPrice() int {
	return 6490
}

func (p Premium) MaxPrice() int {
	return 8690
}

func Calc(car Car, minutes int) int {
	TotalPricePer6hours := minutes / 360 * car.MaxPrice()
	TotalPricePer15Minutes := int(math.Ceil(float64(minutes%360)/15)) * car.PricePer15Minutes()
	if TotalPricePer15Minutes >= car.MaxPrice() {
		return car.MaxPrice() + TotalPricePer6hours
	}
	return TotalPricePer6hours + TotalPricePer15Minutes
}
