package chapter03

import (
	"github.com/kurupeku/hello-golang/helper"
)

// If を使用して料金の条件分岐を行ってください
func InnerChargeFromTokyo(station string) int {
	var fare, distance = 0, 0
	var ns = "東京"
	if station == "東京" {
		return fare
	}
	for {
		ns = helper.InnerNextStation(ns)
		distance = distance + helper.InnerLoopDistance(ns)
		if ns == station {
			break
		}
	}

	if distance < 4000 {
		fare = 140
	} else if distance < 7000 {
		fare = 160
	} else if distance < 11000 {
		fare = 170
	} else if distance < 16000 {
		fare = 200
	} else if distance < 21000 {
		fare = 270
	} else if distance < 26000 {
		fare = 350
	} else if distance < 31000 {
		fare = 420
	} else {
		fare = 490
	}

	return fare
}

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	var fare, distance = 0, 0
	var ns = "東京"
	switch station {
	case "東京":
		return fare
	default:
	}
	for ns != station {
		ns = helper.OuterNextStation(ns)
		distance = distance + helper.OuterLoopDistance(ns)
	}

	switch {
	case distance < 4000:
		fare = 140
	case distance < 7000:
		fare = 160
	case distance < 11000:
		fare = 170
	case distance < 16000:
		fare = 200
	case distance < 21000:
		fare = 270
	case distance < 26000:
		fare = 350
	case distance < 31000:
		fare = 420
	default:
		fare = 490
	}
	return fare
}
