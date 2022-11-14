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

	//変数"tax"を浮動小数点として扱い、price(指定金額)を小数として扱った上でtaxをかける
	//※priceをfloatへ変換しないと、taxとの計算ができない（乗算は小数同士でしかできない）
	var tax float64 = float64(price) * TaxRate

	//変数"price_with_tax(税込み金額)"を符号なし整数として扱い、符号なし整数とした税とprice(指定金額)を加算する
	//※priceとtaxを共に整数として扱わないと計算ができない（加算も小数同士でしかできない）
	var price_with_tax uint = uint(tax) + price

	//税込み金額から500円で割り、500円の枚数を出す
	count500 = price_with_tax / Coin500
	//500円の枚数を出した上で、余った金額を出す
	var reminder_500 uint = price_with_tax % Coin500
	// 下が更新前の式
	//var reminder_500 uint = price_with_tax - Coin500*count500

	//上記500円の必要枚数を引いて残った税込み金額から100円で割り、100円の枚数を出す
	count100 = reminder_500 / Coin100
	//100円の枚数を出した上で、余った金額を出す
	var reminder_100 uint = price_with_tax % Coin100
	// 下が更新前の式
	//var reminder_100 uint = reminder_500 - Coin100*count100

	//上記100円の必要枚数を引いて残った税込み金額から50円で割り、50円の枚数を出す
	count050 = reminder_100 / Coin050
	//50円の枚数を出した上で、余った金額を出す
	var reminder_50 uint = price_with_tax % Coin050
	// 下が更新前の式
	//var reminder_50 uint = reminder_100 - Coin050*count050

	//上記50円の必要枚数を引いて残った税込み金額から10円で割り、10円の枚数を出す
	count010 = reminder_50 / Coin010
	//10円の枚数を出した上で、余った金額を出す
	var reminder_10 uint = price_with_tax % Coin010
	// 下が更新前の式
	//var reminder_10 uint = reminder_50 - Coin010*count010

	//上記10円の必要枚数を引いて残った税込み金額から5円で割り、5円の枚数を出す
	count005 = reminder_10 / Coin005
	//5円の枚数を出した上で、余った金額を出す
	var reminder_5 uint = price_with_tax % Coin005
	// 下が更新前の式
	//var reminder_5 uint = reminder_10 - Coin005*count005

	//上記5円の枚数を出して余った金額が1円の最終枚数となる
	count001 = reminder_5

	//上段は各コインの必要枚数を表示、下段は計算後のあまり金額を表示
	fmt.Println(count500, count100, count050, count010, count005, count001)
	fmt.Println(reminder_500, reminder_100, reminder_50, reminder_10, reminder_5)

	//各コインの枚数（最終結果）を表示する
	return

}
