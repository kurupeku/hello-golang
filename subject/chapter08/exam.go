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

func (basic Basic) PricePer15Minutes() int {
	return 220
}

func (basic Basic) MaxPrice() int {
	return 4290
}

func (middle Middle) PricePer15Minutes() int {
	return 330
}

func (middle Middle) MaxPrice() int {
	return 6490
}

func (premium Premium) PricePer15Minutes() int {
	return 440
}

func (premium Premium) MaxPrice() int {
	return 8690
}

func Calc(car Car, minutes int) int {
	var price int
	calcminutes := minutes

	if minutes >= 360 {
		count := minutes / 360
		price += car.MaxPrice() * count
		calcminutes -= 360 * count
		price += maxOrPayPerUsePrice(car, calcminutes)
	} else {
		price += maxOrPayPerUsePrice(car, minutes)
	}

	return price
}

func countPer15Minutes(minutes int) int {
	return int(math.Ceil((float64(minutes) / 15.0)))
}

func payPerUse(car Car, count int) int {
	return car.PricePer15Minutes() * count
}

func maxOrPayPerUsePrice(car Car, minutes int) int {
	count := countPer15Minutes(minutes)

	if payPerUse(car, count) > car.MaxPrice() {
		return car.MaxPrice()
	} else {
		return payPerUse(car, count)
	}
}
