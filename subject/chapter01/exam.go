package chapter01

import "github.com/kurupeku/hello-golang/helper"

// 初乗り料金
const firstPrice = 500

// 距離ごとに加算される料金
const perPrice = 100

// 初乗りの距離
const firstRideDistance = 1500

// 加算される距離
const perDistance = 250

// 深夜料金割り増し割合
const midnightPremium = 1.2

// 引数に距離を表す文字列、戻り値が通常料金と深夜料金になるように実装してください
func Taxi(distance string) (int, int) {
	// TODO: 実装

	parsedDistance := helper.ParseDistance(distance)

	// 初乗り回数と加算回数の計算
	var firstCount int
	var perCount int

	if parsedDistance > 0 {
		firstCount += 1
	} else {
		firstCount = 0
	}

	diffDistance := parsedDistance - firstRideDistance
	for i := diffDistance; i >= perDistance; i -= perDistance {
		perCount += 1
	}

	// 通常料金の計算
	generalPrice := firstPrice*firstCount + perPrice*perCount

	// 深夜料金の計算
	midnightPrice := int(float64(generalPrice) * midnightPremium)

	return generalPrice, midnightPrice
}
