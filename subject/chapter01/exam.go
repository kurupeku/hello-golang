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

var tujo, shinya int

// 引数に距離を表す文字列、戻り値が通常料金と深夜料金になるように実装してください
func Taxi(distance string) (int, int) {
	// TODO: 実装

	// 距離を数値に変換
	dist := helper.ParseDistance(distance)

	//通常料金を計算
	tujo = firstPrice + ((dist - firstRideDistance) / perDistance * perPrice)

	//深夜料金を計算　通常料金*1.2
	shinya = tujo * 12 / 10

	return tujo, shinya
}
