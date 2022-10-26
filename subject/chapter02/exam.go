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
	taxPriceTotalFloat64 := float64(price) * (1 + TaxRate)
	taxPriceTotal := int(taxPriceTotalFloat64)

	count500 = uint(taxPriceTotal / Coin500)
	taxPriceRest := uint(taxPriceTotal % Coin500)
	count100 = uint(taxPriceRest / Coin100)
	taxPriceRest = uint(taxPriceRest % Coin100)
	count050 = uint(taxPriceRest / Coin050)
	taxPriceRest = uint(taxPriceRest % Coin050)
	count010 = uint(taxPriceRest / Coin010)
	taxPriceRest = uint(taxPriceRest % Coin010)
	count005 = uint(taxPriceRest / Coin005)
	count001 = uint(taxPriceRest % Coin005)
	return
}
