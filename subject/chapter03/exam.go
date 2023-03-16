package chapter03

import "github.com/kurupeku/hello-golang/helper"

// If を使用して料金の条件分岐を行ってください
func InnerChargeFromTokyo(station string) int {
	// TODO: 実装

	// 走行距離
	var mileage int

	// 料金
	var price int

	// 駅
	currentStation := "東京"

	if station == currentStation {
		return price
	}

	for {
		nextStation := helper.InnerNextStation(currentStation)
		mileage += helper.InnerLoopDistance(nextStation)
		currentStation = nextStation

		if currentStation == station {
			break
		}
	}

	if 0 < mileage && mileage <= 3999 {
		price = 140
	} else if mileage <= 6999 {
		price = 160
	} else if mileage <= 10999 {
		price = 170
	} else if mileage <= 15999 {
		price = 200
	} else if mileage <= 20999 {
		price = 270
	} else if mileage <= 25999 {
		price = 350
	} else if mileage <= 30999 {
		price = 420
	} else {
		price = 490
	}

	return price
}

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	// TODO: 実装

	// 走行距離
	var mileage int

	// 料金
	var price int

	// 駅
	currentStation := "東京"

	if station == currentStation {
		return price
	}

	for {
		nextStation := helper.OuterNextStation(currentStation)
		mileage += helper.OuterLoopDistance(nextStation)
		currentStation = nextStation

		if currentStation == station {
			break
		}
	}

	switch {
	case 0 < mileage && mileage <= 3999:
		price = 140
	case 4000 <= mileage && mileage <= 6999:
		price = 160
	case 7000 <= mileage && mileage <= 10999:
		price = 170
	case 11000 <= mileage && mileage <= 15999:
		price = 200
	case 16000 <= mileage && mileage <= 20999:
		price = 270
	case 21000 <= mileage && mileage <= 25999:
		price = 350
	case 26000 <= mileage && mileage <= 30999:
		price = 420
	case 31000 <= mileage:
		price = 490
	}

	return price
}
