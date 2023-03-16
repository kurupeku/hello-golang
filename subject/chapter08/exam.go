package chapter08

import "math"

type Car interface {
	PricePer15Minutes() int
	MaxPrice() int
}

const (
	maxPriceHour             = 360
	basicPricePer15Minutes   = 220
	basicMaxPrice            = 4290
	middlePricePer15Minutes  = 330
	middleMaxPrice           = 6490
	premiumPricePer15Minutes = 440
	premiumMaxPrice          = 8690
)

type Basic struct{}

func (b Basic) PricePer15Minutes() int {
	return basicPricePer15Minutes
}

func (b Basic) MaxPrice() int {
	return basicMaxPrice
}

type Middle struct{}

func (m Middle) PricePer15Minutes() int {
	return middlePricePer15Minutes
}

func (m Middle) MaxPrice() int {
	return middleMaxPrice
}

type Premium struct{}

func (p Premium) PricePer15Minutes() int {
	return premiumPricePer15Minutes
}

func (p Premium) MaxPrice() int {
	return premiumMaxPrice
}

func Calc(car Car, minutes int) int {
	// TODO: 実装

	// 最大料金を何回適用するか
	maxPriceCount := math.Floor((float64(minutes) / float64(maxPriceHour)))
	totalMaxPrice := car.MaxPrice() * int(maxPriceCount)

	// 最大料金分の時間を除いた、残りの時間から15ごとの料金による合計を計算する
	diffMinutes := minutes - int(maxPriceHour*maxPriceCount)
	per15MinutesCount := math.Ceil(float64(diffMinutes) / float64(15))
	totalPricePer15Minutes := car.PricePer15Minutes() * int(per15MinutesCount)

	// 残りの時間から、15ごとの計算料金か、最大料金のどちらを適用するか
	var diffPrice int
	if totalPricePer15Minutes > car.MaxPrice() {
		diffPrice = car.MaxPrice()
	} else {
		diffPrice = totalPricePer15Minutes
	}

	total := totalMaxPrice + diffPrice

	return total
}
