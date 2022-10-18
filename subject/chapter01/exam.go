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
	dis := helper.ParseDistance(distance)
	sabunDis := dis - firstRideDistance
	t := sabunDis / perDistance
	tp := firstPrice + (perPrice * t)

	return tp, tp + tp/5
}
