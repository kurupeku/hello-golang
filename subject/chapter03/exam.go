package chapter03

import (
	"github.com/kurupeku/hello-golang/helper"
)

// If を使用して料金の条件分岐を行ってください
func InnerChargeFromTokyo(station string) int {
	nextStation := "東京"
	totalDistance := 0
	if station == nextStation {
		return 0
	}
	for {
		nextStation = helper.InnerNextStation(nextStation)
		totalDistance += helper.InnerLoopDistance(nextStation)
		if nextStation == station {
			break
		}
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

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	nextStation := "東京"
	totalDistance := 0

	switch station {
	case nextStation:
		return 0
	}

	for {
		nextStation = helper.OuterNextStation(nextStation)
		totalDistance += helper.OuterLoopDistance(nextStation)
		if nextStation == station {
			break
		}
	}

	switch {
	case totalDistance < 4000:
		return 140
	case totalDistance < 7000:
		return 160
	case totalDistance < 11000:
		return 170
	case totalDistance < 16000:
		return 200
	case totalDistance < 21000:
		return 270
	case totalDistance < 26000:
		return 350
	case totalDistance < 31000:
		return 420
	default:
		return 490
	}
}
