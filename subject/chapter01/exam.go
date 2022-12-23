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

// 引数に距離を表す文字列、戻り値が通常料金と深夜料金になるように実装してください
func Taxi(distance string) (int, int) {
	parseDistance := helper.ParseDistance(distance)
	add := (parseDistance - firstRideDistance) / perDistance

	normal := firstPrice + (add * perPrice)
	midnight := float64(normal) * 1.2 // 深夜料金は2割増

	return normal, int(midnight)
}
