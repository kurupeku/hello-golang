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
	// 税込み価格を計算(小数点以下は切り捨て)
	taxPrice := uint(float64(price) + float64(TaxRate)*float64(price))

	// 500円硬貨が何枚使用できるか計算。使用枚数×500をtaxPriceから引く
	count500 = uint(taxPrice / Coin500)
	taxPrice = taxPrice - (count500 * Coin500)

	// 100円硬貨が何枚使用できるか計算。使用枚数×100をtaxPriceから引く
	count100 = uint(taxPrice / Coin100)
	taxPrice = taxPrice - (count100 * Coin100)

	// 50円硬貨が何枚使用できるか計算。使用枚数×50をtaxPriceから引く
	count050 = uint(taxPrice / Coin050)
	taxPrice = taxPrice - (count050 * Coin050)

	// 10円硬貨が何枚使用できるか計算。使用枚数×10をtaxPriceから引く
	count010 = uint(taxPrice / Coin010)
	taxPrice = taxPrice - (count010 * Coin010)

	// 5円硬貨が何枚使用できるか計算。使用枚数×5をtaxPriceから引く
	count005 = uint(taxPrice / Coin005)
	taxPrice = taxPrice - (count005 * Coin005)

	// 1円硬貨が何枚使用できるか計算。taxPriceに小数点以下は含まれないため1円硬貨で全て精算
	count001 = uint(taxPrice / Coin001)

	return
}
