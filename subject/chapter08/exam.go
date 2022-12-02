package chapter08

import "math"

type Car interface {
	PricePer15Minutes() int
	MaxPrice() int
}

type Basic struct{}

type Middle struct{}

type Premium struct{}

func Calc(car Car, minutes int) int {
	// TODO: 実装
	maxprice := minutes / 360
	perprice := minutes % 360
	tmp := int(math.Ceil((float64(perprice) / 15.0))) * car.PricePer15Minutes()
	if tmp >= car.MaxPrice() {
		return maxprice*car.MaxPrice() + car.MaxPrice()
	}
	return maxprice*car.MaxPrice() + tmp
}

func (Basic) PricePer15Minutes() int {
	return 220
}
func (Basic) MaxPrice() int {
	return 4290
}
func (Middle) PricePer15Minutes() int {
	return 330
}
func (Middle) MaxPrice() int {
	return 6490
}
func (Premium) PricePer15Minutes() int {
	return 440
}
func (Premium) MaxPrice() int {
	return 8690
}
