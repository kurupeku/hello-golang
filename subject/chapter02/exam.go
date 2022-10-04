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
	var remain uint
	withTax := price + uint(float64(price)*TaxRate)

	count500 = withTax / Coin500
	remain = withTax - Coin500*count500

	count100 = remain / Coin100
	remain = remain - Coin100*count100

	count050 = remain / Coin050
	remain = remain - Coin050*count050

	count010 = remain / Coin010
	remain = remain - Coin010*count010

	count005 = remain / Coin005
	remain = remain - Coin005*count005

	count001 = remain / Coin001
	return
}
