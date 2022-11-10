package chapter03

import (
	"github.com/kurupeku/hello-golang/helper"
)

// If を使用して料金の条件分岐を行ってください
func InnerChargeFromTokyo(station string) int {
	// TODO: 実装
	// 現在の駅を定義したり、距離を初期化したり
	var Nowstation string = "東京"
	var DestStation string = station
	var Hoshizoranodistance int = 0
	// nowStation が destStationになるまで
	// Hoshizoranodistanceに距離加算を繰り返す
	for Nowstation != DestStation {
		Nowstation = helper.InnerNextStation(Nowstation)
		Hoshizoranodistance += helper.InnerLoopDistance(Nowstation)
		// fmt.Println(Hoshizoranodistance)
	}

	// 東京からの距離に応じて運賃を定義する
	if 31000 < Hoshizoranodistance {
		return 490
	}
	if 26000 < Hoshizoranodistance {
		return 420
	}
	if 21000 < Hoshizoranodistance {
		return 350
	}
	if 16000 < Hoshizoranodistance {
		return 270
	}
	if 11000 < Hoshizoranodistance {
		return 200
	}
	if 7000 < Hoshizoranodistance {
		return 170
	}
	if 4000 < Hoshizoranodistance {
		return 160
	}
	if 0 == Hoshizoranodistance {
		return 0
	} else {
		return 140
	}
}

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	// TODO: 実装
	// 現在の駅を定義したり、距離を初期化したり
	var Nowstation string = "東京"
	var DestStation string = station
	var Hoshizoranodistance int = 0
	// nowStation が destStationになるまで
	// Hoshizoranodistanceに距離加算を繰り返す
	for Nowstation != DestStation {
		Nowstation = helper.OuterNextStation(Nowstation)
		Hoshizoranodistance += helper.OuterLoopDistance(Nowstation)
		// fmt.Println(Hoshizoranodistance)
	}

	// 東京からの距離に応じて運賃を定義する
	switch {
	case 31000 < Hoshizoranodistance:
		return 490
	case 26000 < Hoshizoranodistance:
		return 420
	case 21000 < Hoshizoranodistance:
		return 350
	case 16000 < Hoshizoranodistance:
		return 270
	case 11000 < Hoshizoranodistance:
		return 200
	case 7000 < Hoshizoranodistance:
		return 170
	case 4000 < Hoshizoranodistance:
		return 160
	case 1 < Hoshizoranodistance:
		return 140
	case 0 < Hoshizoranodistance:
		return 0
	}
	return 0
}
