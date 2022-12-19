package chapter08

import (
	//"fmt"
	"math"
)

type Car interface {
	PricePer15Minutes() int
	MaxPrice() int
}

// --- Basic 料金表---
// 空構造体
type Basic struct{}

// メソッドで固定returnを定義することで Basic.PricePer15Minutesが定義される
func (b Basic) PricePer15Minutes() int {
	return 220
}

// メソッドで固定returnを定義することで Basic.MaxPriceが定義される
func (b Basic) MaxPrice() int {
	return 4290
}

// --- Middle 料金表---
type Middle struct{}

func (m Middle) PricePer15Minutes() int {
	return 330
}

func (m Middle) MaxPrice() int {
	return 6490
}

// --- Premium 料金表---
type Premium struct{}

func (p Premium) PricePer15Minutes() int {
	return 440
}

func (p Premium) MaxPrice() int {
	return 8690
}

// --- 料金計算
func Calc(car Car, minutes int) int {
	// 6時間単位料金 計算
	maxCount := minutes / 360
	maxCharge := car.MaxPrice() * maxCount

	// 15分単位料金 計算 (Ceilは切り上げ関数)
	minCharge := 0
	minCount := int(math.Ceil(float64(minutes-maxCount*360) / 15))
	minChargeTmp := car.PricePer15Minutes() * minCount
	switch {
	case car.MaxPrice() < minChargeTmp:
		minCharge = car.MaxPrice()
	default:
		minCharge = minChargeTmp
	}

	// 料金合算
	charge := maxCharge + minCharge

	/*
		fmt.Printf("car.MaxPrice() %d minChargeTmp %d \n", car.MaxPrice(), minChargeTmp)
		fmt.Printf("maxCount %d maxCharge %d \n", maxCount, maxCharge)
		fmt.Printf("minCount %d minCharge %d \n", minCount, minCharge)
	*/

	return charge
}
