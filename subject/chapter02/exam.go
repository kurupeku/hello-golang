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
	var c500 int
	var c100 int
	var c50 int
	var c10 int
	var c5 int
	var c1 int
	var tmp1 float64
	var tmp2 int

	// 結構手抜き
	tmp1 = float64(price)
	tmp1 = tmp1 + (tmp1 * TaxRate)
	tmp2 = int(tmp1)
	fmt.Printf("price: %d\n", tmp2)

	// 順繰り順繰り割っては次に渡していく
	// 500
	c500 = tmp2 / Coin500
	tmp2 = tmp2 % Coin500
	// 100
	c100 = tmp2 / Coin100
	tmp2 = tmp2 % Coin100
	// 50
	c50 = tmp2 / Coin050
	tmp2 = tmp2 % Coin050
	// 10
	c10 = tmp2 / Coin010
	tmp2 = tmp2 % Coin010
	// 5
	c5 = tmp2 / Coin005
	tmp2 = tmp2 % Coin005
	// 1
	c1 = tmp2 / Coin001
	tmp2 = tmp2 % Coin001

	// デバッグ出力させてみる
	fmt.Printf("500:%d\n100:%d\n50:%d\n10:%d\n5:%d\n1:%d\n", c500, c100, c50, c10, c5, c1)
	// どうやって戻せばいいんだっけ。とりあえずダミーを戻して後で考える。
	return 0, 0, 0, 0, 0, 0
}
