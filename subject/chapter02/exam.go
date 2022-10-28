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
	total_price := price + uint(float64(price)*TaxRate)
	count500 = total_price / Coin500
	total_price %= Coin500
	count100 = total_price / Coin100
	total_price %= Coin100
	count050 = total_price / Coin050
	total_price %= Coin050
	count010 = total_price / Coin010
	total_price %= Coin010
	count005 = total_price / Coin005
	count001 = total_price % Coin005
	return
}
