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
	var taxIncludedPrice float64
	var yen uint
	taxIncludedPrice = float64(price) + float64(price)*TaxRate

	var Coin500count = uint(taxIncludedPrice) / Coin500
	var yen500 = uint(taxIncludedPrice) % Coin500
	var Coin100count = yen500 / Coin100
	var yen100 = yen % Coin100
	var Coin050count = yen100 / Coin050
	var yen050 = yen % Coin050
	var Coin010count = yen050 / Coin010
	var yen010 = yen % Coin010
	var Coin005count = yen010 / Coin005
	var yen005 = yen % Coin005
	var Coin001count = yen005 / Coin001
	return uint(Coin500count), uint(Coin100count), uint(Coin050count), uint(Coin010count), uint(Coin005count), uint(Coin001count)
}
