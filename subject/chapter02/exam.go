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
	price_with_tax := uint(float64(price) * (TaxRate + 1))
	count500 = price_with_tax / Coin500
	price_with_tax -= Coin500 * count500
	count100 = price_with_tax / Coin100
	price_with_tax -= Coin100 * count100
	count050 = price_with_tax / Coin050
	price_with_tax -= Coin050 * count050
	count010 = price_with_tax / Coin010
	price_with_tax -= Coin010 * count010
	count005 = price_with_tax / Coin005
	price_with_tax -= Coin005 * count005
	count001 = price_with_tax / Coin001
	return
}
