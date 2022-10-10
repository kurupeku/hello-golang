package chapter03

import (
	"github.com/kurupeku/hello-golang/helper"
)

// If を使用して料金の条件分岐を行ってください
func InnerChargeFromTokyo(station string) int {
	var dist int
	if station == "東京" {
		return 0
	}

	next := "東京"
	for i := 0; i < 29; i++ {
		next = helper.InnerNextStation(next)
		dist += helper.InnerLoopDistance(next)
		if next == station {
			break
		}
	}

	if dist < 4000 {
		return 140
	} else if dist < 7000 {
		return 160
	} else if dist < 11000 {
		return 170
	} else if dist < 16000 {
		return 200
	} else if dist < 21000 {
		return 270
	} else if dist < 26000 {
		return 350
	} else if dist < 31000 {
		return 420
	} else {
		return 490
	}
}

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	var dist int
	if station == "東京" {
		return 0
	}

	next := "東京"
	for i := 0; i < 29; i++ {
		next = helper.OuterNextStation(next)
		dist += helper.OuterLoopDistance(next)
		if next == station {
			break
		}
	}

	switch {
	case dist < 4000:
		return 140
	case dist < 7000:
		return 160
	case dist < 11000:
		return 170
	case dist < 16000:
		return 200
	case dist < 21000:
		return 270
	case dist < 26000:
		return 350
	case dist < 31000:
		return 420
	default:
		return 490
	}
}
