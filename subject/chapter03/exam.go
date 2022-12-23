package chapter03

import (
	"github.com/kurupeku/hello-golang/helper"
)

const StartStation = "東京"

// If を使用して料金の条件分岐を行ってください
func InnerChargeFromTokyo(station string) int {
	// 到着駅が東京の際は0を返す
	if checkTokyo(station) {
		return 0
	}

	distance := innerDistance(station)

	var amount int
	if distance < 4000 {
		amount = 140
	} else if 4000 <= distance && distance < 7000 {
		amount = 160
	} else if 7000 <= distance && distance < 11000 {
		amount = 170
	} else if 11000 <= distance && distance < 16000 {
		amount = 200
	} else if 16000 <= distance && distance < 21000 {
		amount = 270
	} else if 21000 <= distance && distance < 26000 {
		amount = 350
	} else if 26000 <= distance && distance < 31000 {
		amount = 420
	} else {
		amount = 490
	}

	return amount
}

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	// 到着駅が東京の際は0を返す
	if checkTokyo(station) {
		return 0
	}

	distance := outerDistance(station)

	var amount int
	switch {
	case distance < 4000:
		amount = 140
	case 4000 <= distance && distance < 7000:
		amount = 160
	case 7000 <= distance && distance < 11000:
		amount = 170
	case 11000 <= distance && distance < 16000:
		amount = 200
	case 16000 <= distance && distance < 21000:
		amount = 270
	case 21000 <= distance && distance < 26000:
		amount = 350
	case 26000 <= distance && distance < 31000:
		amount = 420
	default:
		amount = 490
	}

	return amount
}

func innerDistance(station string) int {
	var distance int
	var nextStation string
	currentStation := StartStation

	for station != currentStation {
		nextStation = helper.InnerNextStation(currentStation)
		distance += helper.InnerLoopDistance(nextStation)
		currentStation = nextStation
	}

	return distance
}

func outerDistance(station string) int {
	var distance int
	var nextStation string
	currentStation := StartStation

	for station != currentStation {
		nextStation = helper.OuterNextStation(currentStation)
		distance += helper.OuterLoopDistance(nextStation)
		currentStation = nextStation
	}

	return distance
}

func checkTokyo(station string) bool {
	return station == StartStation
}
