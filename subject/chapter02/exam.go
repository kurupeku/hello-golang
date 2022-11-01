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

func MinimumCoins(price uint) (count500, count100, count050, count010, count005, count001 uint) {
	var maney int
	var IncTax float64
	IncTax = float64(price) * (1 + TaxRate)
	maney = int(IncTax)
	count500 = uint(maney / Coin500)
	maney = maney % Coin500
	count100 = uint(maney / Coin100)
	maney = maney % Coin100
	count050 = uint(maney / Coin050)
	maney = maney % Coin050
	count010 = uint(maney / Coin010)
	maney = maney % Coin010
	count005 = uint(maney / Coin005)
	maney = maney % Coin005
	count001 = uint(maney / Coin001)
	return count500, count100, count050, count010, count005, count001
}
