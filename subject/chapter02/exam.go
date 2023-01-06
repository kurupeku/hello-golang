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
	var balance uint

	taxIncludedPrice := uint(float64(price) * (TaxRate + 1)) // 税込金額
	count500, balance = NumberOfCoins(taxIncludedPrice, Coin500)
	count100, balance = NumberOfCoins(balance, Coin100)
	count050, balance = NumberOfCoins(balance, Coin050)
	count010, balance = NumberOfCoins(balance, Coin010)
	count005, balance = NumberOfCoins(balance, Coin005)
	count001, balance = NumberOfCoins(balance, Coin001)

	return
}

func NumberOfCoins(price uint, coin uint) (num uint, balance uint) {
	num = price / coin
	balance = price - coin*num

	return
}
