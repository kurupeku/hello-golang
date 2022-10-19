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
	// 受け取ったstring型の数字をメートルかつint型に直す
	var hoge int
	hoge = helper.ParseDistance(distance)

	// 通常料金を計算する
	var x int
	x = (((hoge - firstRideDistance) / perDistance) * perPrice) + firstPrice

	// 深夜料金は小数点を含む計算をするので、float32型を宣言する
	var y float32
	y = float32(x) * 1.2

	// int型で返さないと怒られるので、yはint型に直してreturnを返す
	return x, int(y)
}
