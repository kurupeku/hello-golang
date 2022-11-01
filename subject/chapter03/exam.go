package chapter03

import "github.com/kurupeku/hello-golang/helper"

func judgePriceTableWithIf(distance int) int {
	if distance <= 3999  { return 140 }
	if distance <= 6999  { return 160 }
	if distance <= 10999 { return 170 }
	if distance <= 15999 { return 200 }
	if distance <= 20999 { return 270 }
	if distance <= 25999 { return 350 }
	if distance <= 30999 { return 420 }
	return 490
}

func judgePriceTableWithSwitch(distance int) int {
	switch {
		case distance <= 3999:  return 140
		case distance <= 6999:  return 160
		case distance <= 10999: return 170
		case distance <= 15999: return 200
		case distance <= 20999: return 270
		case distance <= 25999: return 350
		case distance <= 30999: return 420
		default: return 490
	}
}

// If を使用して料金の条件分岐を行ってください
func InnerChargeFromTokyo(station string) int {
	Tokyo := "東京"
	if station == Tokyo {
		return 0
	}

	nextStation := helper.InnerNextStation(Tokyo)
	distance := helper.InnerLoopDistance(nextStation)

	for nextStation != station {
		nextStation = helper.InnerNextStation(nextStation)
		distance += helper.InnerLoopDistance(nextStation)
	}

	return judgePriceTableWithIf(distance)
}

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	Tokyo := "東京"
	if station == Tokyo {
		return 0
	}

	nextStation := helper.OuterNextStation(Tokyo)
	distance := helper.OuterLoopDistance(nextStation)

	for nextStation != station {
		nextStation = helper.OuterNextStation(nextStation)
		distance += helper.OuterLoopDistance(nextStation)
	}

	return judgePriceTableWithSwitch(distance)
}
