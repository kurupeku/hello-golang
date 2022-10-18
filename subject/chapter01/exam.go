package chapter01

import (
	"math"
	"regexp"
	"strconv"
	"strings"
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
	var totalDistance = ConvToMeters(distance)

	// 初乗り
	var normalPrice int = firstPrice

	// 加算分 ※ 初乗りを超えた途端にメーターは上がるので、距離の端数は切り上げ
	if totalDistance > firstRideDistance {
		normalPrice += int(math.Ceil(float64(totalDistance-firstRideDistance)/perDistance)) * perPrice
	}

	// 計算結果を返す
	return normalPrice, int(1.2 * float64(normalPrice))
}

// 距離をメートルに変換する
func ConvToMeters(distance string) int {

	// 単位当たりの比率
	var unitRatio = map[string]float64{
		"mm": 0.001,
		"cm": 0.01,
		"m":  1,
		"km": 1000}

	// 数値と単位を分割
	rep := regexp.MustCompile(`([\d\.]+)([A-Za-z]*)`)

	// 単位部分を取得
	var numVal, _ = strconv.ParseFloat(rep.ReplaceAllString(distance, "$1"), 64)
	var unitStr string = strings.ToLower(rep.ReplaceAllString(distance, "$2"))

	// メートル変換した整数値を返す
	if len(unitStr) > 0 {
		return int(numVal * unitRatio[unitStr])
	} else {
		return int(numVal)
	}
}
