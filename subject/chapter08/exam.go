package chapter08

import "math"

type Car interface {
	PricePer15Minutes() int
	MaxPrice() int
}

type Basic struct{}

type Middle struct{}

type Premium struct{}

func (c Basic) PricePer15Minutes() int { return 220 }

func (c Basic) MaxPrice() int { return 4290 }

func (c Middle) PricePer15Minutes() int { return 330 }

func (c Middle) MaxPrice() int { return 6490 }

func (c Premium) PricePer15Minutes() int { return 440 }

func (c Premium) MaxPrice() int { return 8690 }

func Calc(car Car, minutes int) int {
	payPerUsePrice := int(math.Ceil(float64(minutes%360)/15)) * car.PricePer15Minutes()

	if payPerUsePrice > car.MaxPrice() {
		payPerUsePrice = car.MaxPrice()
	}

	// 6時間毎の上限課金額を追加して返す
	return int(minutes/360)*car.MaxPrice() + payPerUsePrice
}
