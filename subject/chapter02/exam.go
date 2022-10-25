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

	//税込み価格を計算
	priceintax := price + uint(float64(price)*TaxRate)

	//500円で支払える枚数と残りの金額を計算
	count500 = priceintax / Coin500
	priceintax = priceintax - (Coin500 * count500)

	//100円で支払える枚数と残りの金額を計算　他の硬貨も同様に計算
	count100 = uint(priceintax / Coin100)
	priceintax = priceintax - (Coin100 * count100)
	count050 = uint(priceintax / Coin050)
	priceintax = priceintax - (Coin050 * count050)
	count010 = uint(priceintax / Coin010)
	priceintax = priceintax - (Coin010 * count010)
	count005 = uint(priceintax / Coin005)
	priceintax = priceintax - (Coin005 * count005)
	count001 = uint(priceintax / Coin001)
	return
}
