package chapter03

import "github.com/kurupeku/hello-golang/helper"
const tokyo = "東京"
// If を使用して料金の条件分岐を行ってください。
func InnerChargeFromTokyo(station string) int {
	if station == tokyo {
		return 0
	}

	var currentStation = "東京"
	var nextStation string
	var Kyori int

	for currentStation != station {
		nextStation = helper.InnerNextStation(currentStation)
		Kyori = Kyori + helper.InnerLoopDistance(nextStation)
		currentStation = nextStation
	}

	if Kyori >= 31000 {
		return 490
	}
	if Kyori >= 26000 && Kyori <= 30999 {
		return 420
	}
	if Kyori >= 21000 && Kyori <= 25999 {
		return 350
	}
	if Kyori >= 16000 && Kyori <= 20999 {
		return 270
	}
	if Kyori >= 11000 && Kyori <= 15999 {
		return 200
	}
	if Kyori >= 7000 && Kyori <= 10999 {
		return 170
	}
	if Kyori >= 4000 && Kyori <= 6999 {
		return 160
	}
	if Kyori <= 3999 {
		return 140
	}
	return 0
}
// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	if station == tokyo {
		return 0
	}

	var currentStation = "東京"
	var nextStation string
	var Kyori int

	for currentStation != station {
		nextStation = helper.OuterNextStation(currentStation)
		Kyori = Kyori + helper.OuterLoopDistance(nextStation)
		currentStation = nextStation
	}

	switch {
	case Kyori >= 31000:
		return 490
	case Kyori >= 26000 && Kyori <= 30999:
		return 420
	case Kyori >= 21000 && Kyori <= 25999:
		return 350
	case Kyori >= 16000 && Kyori <= 20999:
		return 270
	case Kyori >= 11000 && Kyori <= 15999:
		return 200
	case Kyori >= 7000 && Kyori <= 10999:
		return 170
	case Kyori >= 4000 && Kyori <= 6999:
		return 160
	case Kyori <= 3999:
		return 140
	}
	return 0
}