package chapter03

import (
	"github.com/kurupeku/hello-golang/helper"
)

// If を使用して料金の条件分岐を行ってください
func InnerChargeFromTokyo(station string) int {

	// 総距離格納用 (初期値 0)
	var totalDistance int

	// 総距離を計算
	for nextStation := "東京"; nextStation != station; {
		// 目的地にヒットしなければ、次の駅をセットしてループ
		nextStation = helper.InnerNextStation(nextStation)
		totalDistance += helper.InnerLoopDistance(nextStation)
	}

	// TODO: 実装
	return GetPrice(totalDistance)
}

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {

	// 総距離格納用 (初期値 0)
	var totalDistance int

	// 総距離を計算
	for nextStation := "東京"; nextStation != station; {
		// 目的地にヒットしなければ、次の駅をセットしてループ
		nextStation = helper.OuterNextStation(nextStation)
		totalDistance += helper.OuterLoopDistance(nextStation)
	}

	// TODO: 実装
	return GetPrice(totalDistance)
}

func GetPrice(totalDistance int) int {

	if totalDistance == 0 {
		return 0
	}
	if totalDistance < 4000 {
		return 140
	}
	if totalDistance < 7000 {
		return 160
	}
	if totalDistance < 11000 {
		return 170
	}
	if totalDistance < 16000 {
		return 200
	}
	if totalDistance < 21000 {
		return 270
	}
	if totalDistance < 26000 {
		return 350
	}
	if totalDistance < 31000 {
		return 420
	}
	return 490
}
