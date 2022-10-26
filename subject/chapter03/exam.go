package chapter03

import (
	"github.com/kurupeku/hello-golang/helper"
)

// If を使用して料金の条件分岐を行ってください
func InnerChargeFromTokyo(station string) int {
	if helper.InnerNextStation(station) == station {
		return 0
	}
	var dist int
	for cs := helper.InnerNextStation("東京"); cs != "東京"; cs = helper.InnerNextStation(cs) {
		dist += helper.InnerLoopDistance(cs)
		if cs == station {
			if dist < 4000 {
				return 140
			}
			if dist < 7000 {
				return 160
			}
			if dist < 11000 {
				return 170
			}
			if dist < 16000 {
				return 200
			}
			if dist < 21000 {
				return 270
			}
			if dist < 26000 {
				return 350
			}
			if dist < 31000 {
				return 420
			}
			return 490
		}
	}
	return 0
}

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	if helper.OuterNextStation(station) == station {
		return 0
	}
	var dist int
	for cs := helper.OuterNextStation("東京"); cs != "東京"; cs = helper.OuterNextStation(cs) {
		dist += helper.OuterLoopDistance(cs)
		if cs == station {
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
	}
	return 0
}
