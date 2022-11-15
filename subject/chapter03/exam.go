package chapter03

import (
	"github.com/kurupeku/hello-golang/helper"
)

// If を使用して料金の条件分岐を行ってください
func InnerChargeFromTokyo(station string) int {

	if station == "東京" {
		return 0
	}

	var nextstation string = "東京"
	var farnextstation int
	var totaldistance int

	for i := 0; i < 29; i++ {
		nextstation = helper.InnerNextStation(nextstation)
		farnextstation = helper.InnerLoopDistance(nextstation)
		totaldistance = totaldistance + farnextstation
		if nextstation == station {
			break
		}
	}

	if totaldistance <= 3999 {
		return 140
	} else if totaldistance >= 4000 && totaldistance <= 6999 {
		return 160
	} else if totaldistance >= 7000 && totaldistance <= 10999 {
		return 170
	} else if totaldistance >= 11000 && totaldistance <= 15999 {
		return 200
	} else if totaldistance >= 16000 && totaldistance <= 20999 {
		return 270
	} else if totaldistance >= 21000 && totaldistance <= 25999 {
		return 350
	} else if totaldistance >= 26000 && totaldistance <= 30999 {
		return 420
	} else {
		return 490
	}
}

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	if station == "東京" {
		return 0
	}

	var nextstation string = "東京"
	var farnextstation int
	var totaldistance int

	for i := 0; i < 29; i++ {
		nextstation = helper.InnerNextStation(nextstation)
		farnextstation = helper.InnerLoopDistance(nextstation)
		totaldistance = totaldistance + farnextstation
		if nextstation == station {
			break
		}
	}

	switch {
	case totaldistance <= 3999:
		return 140
	case totaldistance >= 4000 && totaldistance <= 6999:
		return 160
	case totaldistance >= 7000 && totaldistance <= 10999:
		return 170
	case totaldistance >= 11000 && totaldistance <= 15999:
		return 200
	case totaldistance >= 16000 && totaldistance <= 20999:
		return 270
	case totaldistance >= 21000 && totaldistance <= 25999:
		return 350
	case totaldistance >= 26000 && totaldistance <= 30999:
		return 420
	default:
		return 490
	}
}
