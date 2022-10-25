package chapter01

import (
	"github.com/kurupeku/hello-golang/helper"
)

// 初乗り料金
const firstPrice = 500

// 距離ごとに加算される料金
const perPrice = 100

// 初乗りの距離
const firstRideDistance = 1500

// 加算される距離
const perDistance = 250

// 引数に距離を表す文字列、戻り値が通常料金と深夜料金になるように実装してください
func Taxi(distance string) (int, int) {
	// TODO: 実装

	// :=これは変数の定義もまとめてやってくれるらしい
	rawDistance := helper.ParseDistance(distance)
	excludeFirstRideDistance := rawDistance - firstRideDistance

	// 250m走るまでは加算されないのでintの割り算を行う(余りが出ない)
	addPrice := (excludeFirstRideDistance / perDistance) * perPrice

	// 昼の料金 (初乗り+走料)
	noonPrice := firstPrice + addPrice

	// 深夜の料金 昼の料金*1.2
	// 2割増しのためfloatに型変換
	// returnがint指定なので再変換
	// 料金表の定義から小数点以下が出ることはないので考慮はしていない(1円単位の設定なし)
	midnightPrice := int(float64(noonPrice) * 1.2)

	return noonPrice, midnightPrice
}
