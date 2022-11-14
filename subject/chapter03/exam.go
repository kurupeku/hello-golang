package chapter03

import (
	"github.com/kurupeku/hello-golang/helper"
)

// If を使用して料金の条件分岐を行ってください
func InnerChargeFromTokyo(station string) int {
	// calc distance
	distance := 0
	currentStation := "東京"
	for {
		if station == currentStation {
			break
		}
		currentStation = helper.InnerNextStation(currentStation)
		distance = distance + helper.InnerLoopDistance(currentStation)

	}

	// calc charge
	charge := 0
	if distance == 0 {
		charge = 0
	} else {
		if distance < 4000 {
			charge = 140
		} else {
			if distance < 7000 {
				charge = 160
			} else {
				if distance < 11000 {
					charge = 170
				} else {
					if distance < 16000 {
						charge = 200
					} else {
						if distance < 21000 {
							charge = 270
						} else {
							if distance < 26000 {
								charge = 350
							} else {
								if distance < 31000 {
									charge = 420
								} else {
									charge = 490
								}
							}
						}
					}
				}

			}
		}
	}
	return charge
}

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	// calc distance
	distance := 0
	currentStation := "東京"
	for {
		if station == currentStation {
			break
		}
		currentStation = helper.OuterNextStation(currentStation)
		distance = distance + helper.OuterLoopDistance(currentStation)
	}

	// calc charge
	charge := 0
	switch {
	case distance == 0:
		charge = 0
	case distance < 4000:
		charge = 140
	case distance < 7000:
		charge = 160
	case distance < 11000:
		charge = 170
	case distance < 16000:
		charge = 200
	case distance < 21000:
		charge = 270
	case distance < 26000:
		charge = 350
	case distance < 31000:
		charge = 420
	default:
		charge = 490
	}

	return charge
}
