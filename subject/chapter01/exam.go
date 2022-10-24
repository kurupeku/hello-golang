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
	var dayPrice, nightPrice, addedDistance, addedPrice int

	// 追加距離の単位数
	addedDistance = (helper.ParseDistance(distance) - firstRideDistance) / perDistance
	addedPrice = addedDistance * perPrice
	dayPrice = firstPrice + addedPrice
	nightPrice = dayPrice * 12 / 10

	//fmt.Println(addedDistance)
	//fmt.Println(addedPrice)

	return dayPrice, nightPrice
}
