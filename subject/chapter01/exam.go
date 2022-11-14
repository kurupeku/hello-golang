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

	// 変数名testを定義（距離のメートル単位を格納する）
	var test int
	// 通常料金を変数名へ格納
	var totalamount int
	//
	var latenightcharge float64
	test = helper.ParseDistance(distance)
	totalamount = firstPrice + (test-firstRideDistance)/perDistance*perPrice
	latenightcharge = float64(totalamount) * 1.2

	fmt.Println(distance)
	fmt.Println(test)
	fmt.Println(totalamount)
	fmt.Println(latenightcharge)
	return totalamount, int(latenightcharge)
	//return 0, 0
}
