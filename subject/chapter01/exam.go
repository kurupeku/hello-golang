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
	// TODO: 実装
	// floatに変換してみたやつ
	all_distance := helper.ParseDistance(distance)
	normal_rate_price := (((all_distance - firstRideDistance) / perDistance) * perPrice) + firstPrice
	late_night_rate_price := float64(normal_rate_price) * 1.2

	return normal_rate_price, int(late_night_rate_price)

	// int で頑張ったやつ
	// all_distance := helper.ParseDistance(distance)
	// normal_rate_price := (((all_distance - firstRideDistance) / perDistance) * perPrice) + firstPrice
	// late_night_rate_price := (normal_rate_price * 12) / 10
	// return normal_rate_price, late_night_rate_price

	// ここは自分の整理用
	// billing_distance := all_distance - firstRideDistance
	// billing_count := billing_distance / perDistance
	// billing_price := billing_count * perPrice
	// normal_rate_price := billing_price + firstPrice
	// late_night_rate_price := (normal_rate_price * 12) / 10

}
