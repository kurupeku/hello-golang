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
	p := uint(float64(price) * (1 + TaxRate))
	count500 = divmod(&p, Coin500)
	count100 = divmod(&p, Coin100)
	count050 = divmod(&p, Coin050)
	count010 = divmod(&p, Coin010)
	count005 = divmod(&p, Coin005)
	count001 = divmod(&p, Coin001)
	return
}

func divmod(price *uint, denominator uint) (count uint) {
	p := *price
	count = p / denominator
	*price = p % denominator
	return
}
