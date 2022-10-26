package chapter02

import "fmt"

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
	var tmp1 float64
	var tmp2 uint

	// 結構手抜き
	tmp1 = float64(price)
	tmp1 = tmp1 + (tmp1 * TaxRate)
	tmp2 = uint(tmp1)
	fmt.Printf("price: %d\n", tmp2)

	// 順繰り順繰り割っては次に渡していく
	// 500
	count500 = tmp2 / Coin500
	tmp2 = tmp2 % Coin500
	// 100
	count100 = tmp2 / Coin100
	tmp2 = tmp2 % Coin100
	// 50
	count050 = tmp2 / Coin050
	tmp2 = tmp2 % Coin050
	// 10
	count010 = tmp2 / Coin010
	tmp2 = tmp2 % Coin010
	// 5
	count005 = tmp2 / Coin005
	tmp2 = tmp2 % Coin005
	// 1
	count001 = tmp2 / Coin001
	tmp2 = tmp2 % Coin001

	// デバッグ出力させてみる
	// fmt.Printf("500:%d\n100:%d\n50:%d\n10:%d\n5:%d\n1:%d\n", count500, count100, count050, count010, count005, count001)
	return
}
