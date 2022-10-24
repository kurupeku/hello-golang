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
	amount := price + uint((float64(price) * TaxRate))
	count500 = amount / Coin500
	amount %= Coin500
	count100 = amount / Coin100
	amount %= Coin100
	count050 = amount / Coin050
	amount %= Coin050
	count010 = amount / Coin010
	amount %= Coin010
	count005 = amount / Coin005
	amount %= Coin005
	count001 = amount
	return
}
