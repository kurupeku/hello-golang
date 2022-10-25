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
	var remainderPrice uint = uint(float64(price) * (1 + TaxRate))
	count500 = uint(remainderPrice / Coin500)
	remainderPrice = uint(remainderPrice % Coin500)
	count100 = uint(remainderPrice / Coin100)
	remainderPrice = uint(remainderPrice % Coin100)
	count050 = uint(remainderPrice / Coin050)
	remainderPrice = uint(remainderPrice % Coin050)
	count010 = uint(remainderPrice / Coin010)
	remainderPrice = uint(remainderPrice % Coin010)
	count005 = uint(remainderPrice / Coin005)
	count001 = uint(remainderPrice % Coin005)

	return count500, count100, count050, count010, count005, count001
}
