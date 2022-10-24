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
	fn := func(count *uint, coin uint) {
		*count = amount / coin
		amount %= coin
	}
	fn(&count500, Coin500)
	fn(&count100, Coin100)
	fn(&count050, Coin050)
	fn(&count010, Coin010)
	fn(&count005, Coin005)
	fn(&count001, Coin001)
	return
}
