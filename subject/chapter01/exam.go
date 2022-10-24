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
	// 距離をメートルに変換
	var totalDistance int = helper.ParseDistance(distance)

	// 初乗り
	var normalPrice int = firstPrice

	// 加算分 ※ 初乗りを超えた途端にメーターは上がるので、距離の端数は切り上げ
	if totalDistance > firstRideDistance {
		normalPrice += (totalDistance - firstRideDistance) / perDistance * perPrice
	}

	// 計算結果を返す
	return normalPrice, int(1.2 * float64(normalPrice))
}
