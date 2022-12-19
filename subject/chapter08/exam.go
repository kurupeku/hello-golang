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

func Calc(car Car, minutes int) int {
	// TODO: 実装
	// 6時間毎の合計金額 = 6時間以内の最大料金 * 最大料金の発生回数
	TotalFor6Hour := car.MaxPrice() * (minutes / 360)

	// 15分毎の合計金額 = 15分毎の利用料金 * (6時間以内の分数 / 15分)
	TotalFor15Minutes := car.PricePer15Minutes() * int(math.Ceil(float64(minutes%360)/15))

	// 15分毎の利用料金が最大料金を超えた場合
	if TotalFor15Minutes >= car.MaxPrice() {
		// 6時間以上の合計金額 + 最大料金
		// TotalFor6Hourは切り捨てのため、15分毎の計算で最大料金になる可能性を考慮
		// ※TotalFor6Hourが0の時は純粋に最大料金が返る
		return TotalFor6Hour + car.MaxPrice()
	}

	// 上記以外は、6時間毎の合計金額 + 15分毎の合計金額
	// ※TotalFor6Hourが0の時は純粋に15分毎の合計金額が返る
	return TotalFor6Hour + TotalFor15Minutes
}

// ==========各グレードの15分毎の料金==========
// Basic
func (Basic) PricePer15Minutes() int {
	return 220
}

// Middle
func (Middle) PricePer15Minutes() int {
	return 330
}

// Premium
func (Premium) PricePer15Minutes() int {
	return 440
}

// ==========各グレードの6時間以内の最大料金==========
// Basic
func (Basic) MaxPrice() int {
	return 4290
}

// Middle
func (Middle) MaxPrice() int {
	return 6490
}

// Premium
func (Premium) MaxPrice() int {
	return 8690
}
