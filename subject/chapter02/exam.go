package chapter02

import (
	"fmt"
)

const (
	Coin500         = 500
	Coin100         = 100
	Coin050         = 50
	Coin010         = 10
	Coin005         = 5
	Coin001         = 1
	TaxRate float64 = 0.1
)

func MinimumCoins(price uint) (count500, count100, count050, count010, count005, count001 uint) {
	// TODO: 実装

	var tax float64 = float64(price) * TaxRate
	var price_with_tax uint = uint(tax) + price

	count500 = price_with_tax / Coin500
	//float64(price) * TaxRate / Coin500
	reminder_500 := price_with_tax - Coin500*count500

	count100 = reminder_500 / Coin100
	reminder_100 := reminder_500 - Coin100*count100

	count050 = reminder_100 / Coin050
	reminder_50 := reminder_100 - Coin050*count050

	count010 = reminder_50 / Coin010
	reminder_10 := reminder_50 - Coin010*count010

	count005 = reminder_10 / Coin005
	reminder_5 := reminder_10 - Coin005*count005

	count001 = reminder_5
	//reminder_1 := reminder_5 - Coin001*count001
	//count001 = count005 % Coin005

	//求めたいのはここ→各コインの枚数
	fmt.Println(count500, count100, count050, count010, count005, count001)
	fmt.Println(reminder_500, reminder_100, reminder_50, reminder_10, reminder_5)

	//各コインの枚数をプリントする
	return //int(numbers_500, numbers_100, numbers_50, numbers_10, numbers_5), reminder_1)

}
