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
	zeikomi := uint(float64(price)*TaxRate) + price
	count500 = zeikomi / Coin500
	Amari500 := zeikomi % Coin500
	count100 = Amari500 / Coin100
	Amari100 := Amari500 % Coin100
	count050 = Amari100 / Coin050
	Amari050 := Amari100 % Coin050
	count010 = Amari050 / Coin010
	Amari010 := Amari050 % Coin010
	count005 = Amari010 / Coin005
	Amari005 := Amari010 % Coin005
	count001 = Amari005 / Coin001
	return
}
