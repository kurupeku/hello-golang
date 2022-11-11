package chapter03

import (
	// "fmt"
	"github.com/kurupeku/hello-golang/helper"
)

// If を使用して料金の条件分岐を行ってください
func InnerChargeFromTokyo(station string) int {
	// TODO: 実装
	var start string = "東京"
	var end string = station
	var total_distance int

	if start == end {
		return 0
	}

	for start != end {
		total_distance = helper.InnerLoopDistance(start) + total_distance
		start = helper.InnerNextStation(start)
	}

	return GetChargeIf(total_distance)
}

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	// TODO: 実装
	var start string = "東京"
	var end string = station
	var total_distance int

	if start == end {
		return 0
	}

	for start != end {
		total_distance = helper.InnerLoopDistance(start) + total_distance
		start = helper.OuterNextStation(start)
	}

	return getChargeSwitch(total_distance)
}

func GetChargeIf(distance int) int {
	var charge int
	if distance <= 3999 {
		charge = 140
	} else if distance <= 6999 {
		charge = 160
	} else if distance <= 10999 {
		charge = 170
	} else if distance <= 15999 {
		charge = 200
	} else if distance <= 20999 {
		charge = 270
	} else if distance <= 25999 {
		charge = 350
	} else if distance <= 30999 {
		charge = 420
	} else {
		charge = 490
	}

	return charge
}

func getChargeSwitch(distance int) int {
	var charge int
	switch {
	case distance <= 3999:
		charge = 140
	case distance <= 6999:
		charge = 160
	case distance < 10999:
		charge = 170
	case distance < 15999:
		charge = 200
	case distance <= 20999:
		charge = 270
	case distance <= 25999:
		charge = 350
	case distance <= 30999:
		charge = 420
	case distance >= 31000:
		charge = 490
	}

	return charge
}
