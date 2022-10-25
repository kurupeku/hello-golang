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
	tax_included_price := uint(float64(price) * (1 + TaxRate))
	count500 = uint(tax_included_price / Coin500)
	amari_of_count500 := uint(tax_included_price % Coin500)

	count100 = uint(amari_of_count500 / Coin100)
	amari_of_count100 := uint(amari_of_count500 % Coin100)

	count050 = uint(amari_of_count100 / Coin050)
	amari_of_count050 := uint(amari_of_count100 % Coin050)

	count010 = uint(amari_of_count050 / Coin010)
	amari_of_count010 := uint(amari_of_count050 % Coin010)

	count005 = uint(amari_of_count010 / Coin005)
	amari_of_count005 := uint(amari_of_count010 % Coin005)

	count001 = uint(amari_of_count005 / Coin001)

	return count500, count100, count050, count010, count005, count001
}
