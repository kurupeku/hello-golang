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
	p := float64(price) * (1 + TaxRate)
	count500, remain1 := divmod(uint(p), Coin500)
	count100, remain2 := divmod(remain1, Coin100)
	count050, remain3 := divmod(remain2, Coin050)
	count010, remain4 := divmod(remain3, Coin010)
	count005, remain5 := divmod(remain4, Coin005)
	count001, _ = divmod(remain5, Coin001)
	return
}

func divmod(numerator, denominator uint) (quotient, remainder uint) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}
