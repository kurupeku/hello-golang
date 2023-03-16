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

	// 税込金額
	taxIncludedPrice := uint(price) + uint(float64(price)*TaxRate)

	// 差額金額
	diffPrice := taxIncludedPrice

	for i := diffPrice; i >= Coin500; i -= Coin500 {
		count500 += 1
		diffPrice -= Coin500
	}
	for i := diffPrice; i >= Coin100; i -= Coin100 {
		count100 += 1
		diffPrice -= Coin100
	}
	for i := diffPrice; i >= Coin050; i -= Coin050 {
		count050 += 1
		diffPrice -= Coin050
	}
	for i := diffPrice; i >= Coin010; i -= Coin010 {
		count010 += 1
		diffPrice -= Coin010
	}
	for i := diffPrice; i >= Coin005; i -= Coin005 {
		count005 += 1
		diffPrice -= Coin005
	}
	for i := diffPrice; i >= Coin001; i -= Coin001 {
		count001 += 1
		diffPrice -= Coin001
	}

	return
}
