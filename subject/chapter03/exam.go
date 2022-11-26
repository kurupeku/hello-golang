package chapter03

import "github.com/kurupeku/hello-golang/helper"

// If を使用して料金の条件分岐を行ってください
func InnerChargeFromTokyo(station string) int {

	// 初期化
	var Nowstation string = "東京"
	var DestStation string = station
	var Hoshizoranodistance int = 0
	var price int = 0
	for Nowstation != DestStation {
		Nowstation = helper.InnerNextStation(Nowstation)
		Hoshizoranodistance += helper.InnerLoopDistance(Nowstation)
	}

	// 運賃
	if 31000 < Hoshizoranodistance {
		price = 490
	} else if 26000 < Hoshizoranodistance {
		price = 420
	} else if 21000 < Hoshizoranodistance {
		price = 350
	} else if 16000 < Hoshizoranodistance {
		price = 270
	} else if 11000 < Hoshizoranodistance {
		price = 200
	} else if 7000 < Hoshizoranodistance {
		price = 170
	} else if 4000 < Hoshizoranodistance {
		price = 160
	} else if 1 < Hoshizoranodistance {
		price = 140
	}
	return price
}

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	// TODO: 実装
	// 初期化
	var Nowstation string = "東京"
	var DestStation string = station
	var price int = 0
	var Hoshizoranodistance int = 0
	for Nowstation != DestStation {
		Nowstation = helper.OuterNextStation(Nowstation)
		Hoshizoranodistance += helper.OuterLoopDistance(Nowstation)
	}

	// 東京からの距離に応じて運賃を定義する
	switch {
	case 31000 < Hoshizoranodistance:
		price = 490
	case 26000 < Hoshizoranodistance:
		price = 420
	case 21000 < Hoshizoranodistance:
		price = 350
	case 16000 < Hoshizoranodistance:
		price = 270
	case 11000 < Hoshizoranodistance:
		price = 200
	case 7000 < Hoshizoranodistance:
		price = 170
	case 4000 < Hoshizoranodistance:
		price = 160
	case 1 < Hoshizoranodistance:
		price = 140
	}
	return price
}
