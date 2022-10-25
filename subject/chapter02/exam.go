package chapter02

//	"fmt"

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
	// 税込価格 (1円未満切り捨て)
	fullPrice := uint(float64(price) + float64(price)*TaxRate)

	count500 = fullPrice / Coin500
	count100 = (fullPrice - Coin500*count500) / Coin100
	count050 = (fullPrice - Coin500*count500 - Coin100*count100) / Coin050
	count010 = (fullPrice - Coin500*count500 - Coin100*count100 - Coin050*count050) / Coin010
	count005 = (fullPrice - Coin500*count500 - Coin100*count100 - Coin050*count050 - Coin010*count010) / Coin005
	count001 = (fullPrice - Coin500*count500 - Coin100*count100 - Coin050*count050 - Coin010*count010 - Coin005*count005) / Coin001

	//fmt.Printf("fullPrice %v \n", fullPrice)
	//fmt.Printf("count500 %v \n", count500)
	//fmt.Printf("count100 %v \n", count100)
	//fmt.Printf("count050 %v \n", count050)
	//fmt.Printf("count010 %v \n", count010)
	//fmt.Printf("count005 %v \n", count005)
	//fmt.Printf("count001 %v \n", count001)

	return
}
