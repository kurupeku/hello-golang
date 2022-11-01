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
	taxPrice := price + uint(float64(price)*TaxRate)
	count500 = taxPrice / Coin500
	count100 = (taxPrice % Coin500) / Coin100
	count050 = ((taxPrice % Coin500) % Coin100) / Coin050
	count010 = (((taxPrice % Coin500) % Coin100) % Coin050) / Coin010
	count005 = ((((taxPrice % Coin500) % Coin100) % Coin050) % Coin010) / Coin005
	count001 = (((((taxPrice % Coin500) % Coin100) % Coin050) % Coin010) % Coin005) / Coin001
	return
}
