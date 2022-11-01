package chapter03

import "github.com/kurupeku/hello-golang/helper"

// If を使用して料金の条件分岐を行ってください
func InnerChargeFromTokyo(station string) int {
	// TODO: 実装
	current_station := "東京"
	distance := 0

	//東京なら0を返す
	if station == current_station {
		return 0
	}

	//目的地まで距離計算
	for {
		current_station = helper.InnerNextStation(current_station)
		distance += helper.InnerLoopDistance(current_station)
		if current_station == station {
			break
		}
	}

	//料金
	if distance < 4000 {
		return 140
	}
	if distance < 7000 {
		return 160
	}
	if distance < 11000 {
		return 170
	}
	if distance < 16000 {
		return 200
	}
	if distance < 21000 {
		return 270
	}
	if distance < 26000 {
		return 350
	}
	if distance < 31000 {
		return 420
	}
	return 490
}

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	// TODO: 実装
	current_station := "東京"
	distance := 0

	//東京なら0を返す
	switch station {
	case current_station:
		return 0
	}

	//目的地まで距離計算
	for {
		current_station = helper.OuterNextStation(current_station)
		distance += helper.OuterLoopDistance(current_station)
		if current_station == station {
			break
		}
	}

	//料金
	switch {
	case distance < 4000:
		return 140
	case distance < 7000:
		return 160
	case distance < 11000:
		return 170
	case distance < 16000:
		return 200
	case distance < 21000:
		return 270
	case distance < 26000:
		return 350
	case distance < 31000:
		return 420
	default:
		return 490
	}
}
