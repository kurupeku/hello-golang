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
	price = price + uint(float64(price)*TaxRate)
	count500 = price / Coin500
	price = price - (Coin500 * count500)
	count100 = price / Coin100
	price = price - (Coin100 * count100)
	count050 = price / Coin050
	price = price - (Coin050 * count050)
	count010 = price / Coin010
	price = price - (Coin010 * count010)
	count005 = price / Coin005
	price = price - (Coin005 * count005)
	count001 = price / Coin001

	return
}
