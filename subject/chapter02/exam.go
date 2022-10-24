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
	fn := func(coin uint) (count uint) {
		count = amount / coin
		amount %= coin
		return
	}
	count500 = fn(Coin500)
	count100 = fn(Coin100)
	count050 = fn(Coin050)
	count010 = fn(Coin010)
	count005 = fn(Coin005)
	count001 = fn(Coin001)
	return
}
