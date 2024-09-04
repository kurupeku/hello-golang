package chapter08

import (
	"math"
)

type Car interface {
	PricePer15Minutes() int
	MaxPrice() int
}

// Car interface を実装するようにメソッドを定義し、その実装してください
type Basic struct{}

func (b *Basic) PricePer15Minutes() int {
	return 220
}

func (b *Basic) MaxPrice() int {
	return 4290
}

// Car interface を実装するようにメソッドを定義し、その実装してください
type Middle struct{}

func (m *Middle) PricePer15Minutes() int {
	return 330
}

func (m *Middle) MaxPrice() int {
	return 6490
}

// Car interface を実装するようにメソッドを定義し、その実装してください
type Premium struct{}

func (p *Premium) PricePer15Minutes() int {
	return 440
}

func (p *Premium) MaxPrice() int {
	return 8690
}

// 渡されたクルマのグレードと駐車時間から利用料金を計算する関数を実装してください
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
