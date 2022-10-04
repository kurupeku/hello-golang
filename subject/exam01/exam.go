package exam01

import "github.com/kurupeku/hello-golang/helper"

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
	dist := helper.ParseDistance(distance)
	calc := dist - firstRideDistance
	times := calc / perDistance
	price := firstPrice + perPrice*times

	return price, int(float64(price) * 1.2)
}
