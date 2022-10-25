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
	tax := float64(price) * TaxRate
	in_tax_price := price + uint(tax)
	count500 = in_tax_price / Coin500
	count100 = (in_tax_price % Coin500) / Coin100
	count050 = (in_tax_price % Coin100) / Coin050
	count010 = (in_tax_price % Coin050) / Coin010
	count005 = (in_tax_price % Coin010) / Coin005
	count001 = (in_tax_price % Coin005) / Coin001

	return
	// return count500, count100, count050, count010, count005, count001
}
