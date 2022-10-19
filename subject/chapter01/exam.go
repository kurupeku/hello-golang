// Yasuhiro Fujii
// 2022/10/17 002

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
	var tmp1 int
	var normal1 int
	var night1 float32
	tmp1 = helper.ParseDistance(distance)
	// 0除算エラーを防ぐため。
	if tmp1 <= firstRideDistance {
		normal1 = firstPrice
	} else {
		normal1 = (((tmp1 - firstRideDistance) / perDistance) * perPrice) + firstPrice
	}
	night1 = float32(normal1) * 1.2

	return normal1, int(night1)
}
