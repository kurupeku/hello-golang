package chapter02

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
	var fill uint = uint(float64(price) * (TaxRate + 1))

	count500 = fill / Coin500
	fill %= Coin500

	count100 = fill / Coin100
	fill %= Coin100

	count050 = fill / Coin050
	fill %= Coin050

	count010 = fill / Coin010
	fill %= Coin010

	count005 = fill / Coin005
	count001 = fill % Coin005

	return
}
