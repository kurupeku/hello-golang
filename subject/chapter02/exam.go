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
	taxInclude := uint(float64(price) + (float64(price) * TaxRate))
	count500 = taxInclude / Coin500
	taxInclude -= Coin500 * count500

	count100 = taxInclude / Coin100
	taxInclude -= Coin100 * count100

	count050 = taxInclude / Coin050
	taxInclude -= Coin050 * count050

	count010 = taxInclude / Coin010
	taxInclude -= Coin010 * count010

	count005 = taxInclude / Coin005
	taxInclude -= Coin005 * count005

	count001 = taxInclude / Coin001
	taxInclude -= Coin001 * count001

	return
}
