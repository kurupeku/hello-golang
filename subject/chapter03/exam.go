package chapter03

import "github.com/kurupeku/hello-golang/helper"

// If を使用して料金の条件分岐を行ってください
func InnerChargeFromTokyo(station string) int {
	// TODO: 実装
	//東京であれば0
	if station == "東京" {
		return 0
	}

	totalDistance := 0
	nextStation := "東京"
	distance := 0
	for {
		//次の駅を取得
		nextStation = helper.InnerNextStation(nextStation)
		//次の駅までの距離を取得
		distance = helper.InnerLoopDistance(nextStation)
		//距離を加算する
		totalDistance = totalDistance + distance
		println(totalDistance)
		//目的地であれば終了
		if nextStation == station {
			break
		}
	}

	if totalDistance <= 3999 {
		return 140
	} else if totalDistance <= 6999 {
		return 160
	} else if totalDistance <= 10999 {
		return 170
	} else if totalDistance <= 15999 {
		return 200
	} else if totalDistance <= 20999 {
		return 270
	} else if totalDistance <= 25999 {
		return 350
	} else if totalDistance <= 30999 {
		return 420
	} else {
		return 490
	}
}

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	// TODO: 実装
	//東京であれば0
	if station == "東京" {
		return 0
	}

	totalDistance := 0
	nextStation := "東京"
	distance := 0
	for {
		//次の駅を取得
		nextStation = helper.OuterNextStation(nextStation)
		//次の駅までの距離を取得
		distance = helper.OuterLoopDistance(nextStation)
		//距離を加算する
		totalDistance = totalDistance + distance
		//目的地であれば終了
		if nextStation == station {
			break
		}
	}

	switch {
	case totalDistance <= 3999:
		return 140
	case totalDistance <= 6999:
		return 160
	case totalDistance <= 10999:
		return 170
	case totalDistance <= 15999:
		return 200
	case totalDistance <= 20999:
		return 270
	case totalDistance <= 25999:
		return 350
	case totalDistance <= 30999:
		return 420
	default:
		return 490
	}

	/*
		if totalDistance <= 3999 {
			return 140
		} else if totalDistance <= 6999 {
			return 160
		} else if totalDistance <= 10999 {
			return 170
		} else if totalDistance <= 15999 {
			return 200
		} else if totalDistance <= 20999 {
			return 270
		} else if totalDistance <= 25999 {
			return 350
		} else if totalDistance <= 30999 {
			return 420
		} else {
			return 490
		}
	*/
}
