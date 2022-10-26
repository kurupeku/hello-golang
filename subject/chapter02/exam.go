package chapter02

const (
	Coin500         = 500
	Coin100         = 100
	Coin050         = 50
	Coin010         = 10
	Coin005         = 5
	Coin001         = 1
	TaxRate      float64 = 0.1
)

// 税込みに変換
func changeTaxIncluded(price uint) uint {
	return uint(float64(price) * (TaxRate + 1))
}

// コインの使用枚数
func countCoin(price uint, coin uint) uint {
	return price / coin
}

func MinimumCoins(price uint) (count500, count100, count050, count010, count005, count001 uint) {
	currentPrice := changeTaxIncluded(price)
	count500 = countCoin(currentPrice, Coin500)

	currentPrice -= count500 * Coin500
	count100 = countCoin(currentPrice, Coin100)

	currentPrice -= count100 * Coin100
	count050 = countCoin(currentPrice, Coin050)

	currentPrice -= count050 * Coin050
	count010 = countCoin(currentPrice, Coin010)

	currentPrice -= count010 * Coin010
	count005 = countCoin(currentPrice, Coin005)

	currentPrice -= count005 * Coin005
	count001 = countCoin(currentPrice, Coin001)
	return
}
