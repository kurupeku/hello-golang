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

const warimashi = 1.2

// 引数に距離を表す文字列、戻り値が通常料金と深夜料金になるように実装してください
func Taxi(distance string) (int, int) {
	// TODO: 実装
	var i int
	i = helper.ParseDistance(distance)
	var x int
	x = firstPrice + (i-firstRideDistance)/perDistance*perPrice
	var y float32
	y = float32(x) * warimashi
	println(x, y)
	return x, int(y)
}
