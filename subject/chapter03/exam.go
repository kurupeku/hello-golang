package chapter03

import (
	"github.com/kurupeku/hello-golang/helper"
)

// If を使用して料金の条件分岐を行ってください
func InnerChargeFromTokyo(station string) int {
	// TODO: 実装
	// 初期値の設定
	nowStation := "東京"
	distance := 0

	// 東京駅から目的地までの距離を算出
	for nowStation != station {
		nextStation := helper.InnerNextStation(nowStation)
		distance = distance + helper.InnerLoopDistance(nextStation)
		nowStation = nextStation
	}

	// 距離に応じた賃料を返す
	if station == "東京" {
		return 0
	} else if distance <= 3999 {
		return 140
	} else if distance <= 6999 && distance >= 4000 {
		return 160
	} else if distance <= 10999 && distance >= 7000 {
		return 170
	} else if distance <= 15999 && distance >= 11000 {
		return 200
	} else if distance <= 20999 && distance >= 16000 {
		return 270
	} else if distance <= 25999 && distance >= 21000 {
		return 350
	} else if distance <= 30999 && distance >= 26000 {
		return 420
	} else if distance >= 31000 {
		return 490
	}

	return 0
}

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	// TODO: 実装

	// 初期値の設定
	nowStation := "東京"
	distance := 0

	// 東京から目的地までの距離を算出
	for nowStation != station {
		nextStation := helper.OuterNextStation(nowStation)
		distance = distance + helper.OuterLoopDistance(nextStation)
		nowStation = nextStation
	}

	// 距離に応じた賃料を返す
	switch {
	case station == "東京":
		return 0
	case distance <= 3999:
		return 140
	case distance <= 6999 && distance >= 4000:
		return 160
	case distance <= 10999 && distance >= 7000:
		return 170
	case distance <= 15999 && distance >= 11000:
		return 200
	case distance <= 20999 && distance >= 16000:
		return 270
	case distance <= 25999 && distance >= 21000:
		return 350
	case distance <= 30999 && distance >= 26000:
		return 420
	case distance >= 31000:
		return 490
	}

	return 0
}
