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

func calculateGeneralAmount(distanceM int) int {
	amount := firstPrice
	// 加算回数
	addCount := (distanceM - firstRideDistance) / perDistance
	amount += perPrice * addCount

	return amount
}

func calculateLateAmount(generalAmount int) int {
	// 1.2倍
	return generalAmount * 12 / 10
}

// 引数に距離を表す文字列、戻り値が通常料金と深夜料金になるように実装してください
func Taxi(distance string) (int, int) {
	// メートルに変換
	distanceM := helper.ParseDistance(distance)

	generalAmount := calculateGeneralAmount(distanceM)
	lateAmount := calculateLateAmount(generalAmount)
	
	return generalAmount, lateAmount
}
