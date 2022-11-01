package chapter03

import (
	"github.com/kurupeku/hello-golang/helper"
)

// If を使用して料金の条件分岐を行ってください
func InnerChargeFromTokyo(station string) int {
	// TODO: 実装

	current := "東京"
	distance := 0

	if station == current {
		return 0
	}

	for {
		next := helper.InnerNextStation(current)
		distance += helper.InnerLoopDistance(next)
		current = next

		if station == current {
			break
		}
	}

	if distance > 0 && distance <= 3999 {
		return 140
	} else if distance >= 4000 && distance <= 6999 {
		return 160
	} else if distance >= 7000 && distance <= 10999 {
		return 170
	} else if distance >= 11000 && distance <= 15999 {
		return 200
	} else if distance >= 16000 && distance <= 20999 {
		return 270
	} else if distance >= 21000 && distance <= 25999 {
		return 350
	} else if distance >= 26000 && distance <= 30999 {
		return 420
	} else {
		return 490
	}
}

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	// TODO: 実装

	current := "東京"
	distance := 0

	if station == current {
		return 0
	}

	for {
		next := helper.OuterNextStation(current)
		distance += helper.OuterLoopDistance(next)
		current = next

		if station == current {
			break
		}
	}

	switch {
	case distance > 0 && distance <= 3999:
		return 140
	case distance >= 4000 && distance <= 6999:
		return 160
	case distance >= 7000 && distance <= 10999:
		return 170
	case distance >= 11000 && distance <= 15999:
		return 200
	case distance >= 16000 && distance <= 20999:
		return 270
	case distance >= 21000 && distance <= 25999:
		return 350
	case distance >= 26000 && distance <= 30999:
		return 420
	default:
		return 490
	}
}
