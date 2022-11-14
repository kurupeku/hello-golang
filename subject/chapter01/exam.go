package chapter01

import (
	"fmt"
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

	var i int
	var normal_ryokin int
	var night_ryokin float64
	i = helper.ParseDistance(distance)
	normal_ryokin = firstPrice + (i - firstRideDistance) / perDistance * perPrice
	night_ryokin = float64(normal_ryokin) * 1.2

	fmt.Println("走行距離は", i,"m")
	fmt.Println("通常料金は", normal_ryokin,"円")
	fmt.Println("割増料金は", night_ryokin,"円")
	return normal_ryokin, int(night_ryokin)
	//return 0, 0
}