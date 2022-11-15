package chapter03

import "github.com/kurupeku/hello-golang/helper"

// If を使用して料金の条件分岐を行ってください
func InnerChargeFromTokyo(station string) int {
	// TODO: 実装
	if station == "東京" {
		return 0
	} else {
		next_station_distance := 0
		next_station := helper.InnerNextStation("東京")
		for {
			if next_station == station {
				next_station_distance = next_station_distance + helper.InnerLoopDistance(station)
				break
			} else {
				next_station_distance = next_station_distance + helper.InnerLoopDistance(next_station)
				next_station = helper.InnerNextStation(next_station)
			}
		}
		if next_station_distance < 4000 {
			return 140
		} else if next_station_distance < 7000 {
			return 160
		} else if next_station_distance < 11000 {
			return 170
		} else if next_station_distance < 16000 {
			return 200
		} else if next_station_distance < 21000 {
			return 270
		} else if next_station_distance < 26000 {
			return 350
		} else if next_station_distance < 31000 {
			return 420
		} else {
			return 490
		}
	}
}

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	// TODO: 実装
	if station == "東京" {
		return 0
	} else {
		next_station_distance := 0
		next_station := helper.OuterNextStation("東京")
	culk_dist:
		for {
			switch next_station {
			case station:
				next_station_distance = next_station_distance + helper.OuterLoopDistance(station)
				break culk_dist
			default:
				next_station_distance = next_station_distance + helper.OuterLoopDistance(next_station)
				next_station = helper.OuterNextStation(next_station)
				continue
			}
		}
		distance := next_station_distance
		switch {
		case distance <= 4000:
			return 140
		case distance < 7000:
			return 160
		case distance < 11000:
			return 170
		case distance < 16000:
			return 200
		case distance < 21000:
			return 270
		case distance < 26000:
			return 350
		case distance < 31000:
			return 420
		default:
			return 490
		}
	}
}
