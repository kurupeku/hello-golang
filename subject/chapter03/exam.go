package chapter03

import (
	"fmt"

	"github.com/kurupeku/hello-golang/helper"
)

// If を使用して料金の条件分岐を行ってください
func InnerChargeFromTokyo(station string) int {
	current := station
	next := ""
	distance := 0

	switch {
	case current == "東京":
		return 0
	case current != "東京":
		for current != "東京" {
			next = helper.OuterNextStation(current)
			distance += helper.OuterLoopDistance(next)
			fmt.Println("distance=", distance, "current=", current, " next=", next)
			current = next
		}
	}
	return DistanceToCharge(distance)
}

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	current := station
	next := ""
	distance := 0

	if current == "東京" {
		return 0
	}

	for current != "東京" {
		next = helper.InnerNextStation(current)
		distance += helper.InnerLoopDistance(next)
		fmt.Println("distance=", distance, "current=", current, " next=", next)
		current = next
	}
	return DistanceToCharge(distance)
}

func DistanceToCharge(distance int) int {
	charge := 0
	switch {
	case distance <= 3999:
		charge = 140
	case 4000 <= distance && distance < 6999:
		charge = 160
	case 7000 <= distance && distance < 10999:
		charge = 170
	case 11000 <= distance && distance < 15999:
		charge = 200
	case 16000 <= distance && distance < 20999:
		charge = 270
	case 21000 <= distance && distance < 25999:
		charge = 350
	case 26000 <= distance && distance < 30999:
		charge = 420
	case 31000 <= distance:
		charge = 490
	}
	return charge
}
