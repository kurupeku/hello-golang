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

// 通常料金の比率
const normalPriceRate = 100

// 深夜料金の比率
const midnightPriceRate = 120

// 引数に距離を表す文字列、戻り値が通常料金と深夜料金になるように実装してください
func Taxi(distance string) (int, int) {
	var d1 int = helper.ParseDistance(distance)
	var d2 int = d1 - firstRideDistance
	var p1 int = firstPrice + (d2 / perDistance * perPrice)
	var p2 int = p1 * midnightPriceRate / normalPriceRate
	return p1, p2
}
