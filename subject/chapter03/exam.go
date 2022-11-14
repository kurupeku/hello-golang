package chapter03

import (
	"github.com/kurupeku/hello-golang/helper"
)

// If を使用して料金の条件分岐を行ってください
func InnerChargeFromTokyo(station string) int {
	// TODO: 実装

	//問題文から目的地が東京の場合には0円を返す
	if station == "東京" {
		return 0
	}

	//出発駅に東京駅を設定する
	//var station string = "東京" としてしまうと次に求める駅を定義することができない
	//以下の通りにすると、最初に求めるのが東京なので、ループができる
	var next_station string = "東京"
	//次の駅までの距離をint(整数)で定義する
	var distance_to_next_station int
	//目的地までの総距離をint(整数)で定義する
	var total_distance int

	//駅数をiで記載し、上限30個の駅に達するまで同じ処理を繰り返す
	for i := 0; i < 30; i++ {
		//helperを用いて次の駅の名前を求め、next_stationに格納
		next_station = helper.InnerNextStation(next_station)
		//helperを用いて次の駅までの距離を求め、distance_to_next_stationに格納
		distance_to_next_station = helper.InnerLoopDistance(next_station)
		//目的地までの総距離を1回求めるごとに次の駅までの距離を加算し、total_distanceに格納
		total_distance = total_distance + distance_to_next_station
		//求める駅まで来たら、処理を終了する
		if next_station == station {
			break
		}
	}

	//目的地までの総距離が3999m以下ならば、140円を返す
	if total_distance <= 3999 {
		return 140
		//目的地までの総距離が4000m以上6999m以下ならば、160円を返す
	} else if total_distance >= 4000 && total_distance <= 6999 {
		return 160
		//目的地までの総距離が7000m以上10999m以下ならば、170円を返す
	} else if total_distance >= 7000 && total_distance <= 10999 {
		return 170
		//目的地までの総距離が11000m以上15999m以下ならば、200円を返す
	} else if total_distance >= 11000 && total_distance <= 15999 {
		return 200
		//目的地までの総距離が16000m以上20999m以下ならば、270円を返す
	} else if total_distance >= 16000 && total_distance <= 20999 {
		return 270
		//目的地までの総距離が21000m以上25999m以下ならば、350円を返す
	} else if total_distance >= 21000 && total_distance <= 25999 {
		return 350
		//目的地までの総距離が26000m以上30999m以下ならば、420円を返す
	} else if total_distance >= 26000 && total_distance <= 30999 {
		return 420
		//それ以外は490円を返す
	} else {
		return 490
	}
}

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	// TODO: 実装

	//問題文から目的地が東京の場合には0円を返す
	if station == "東京" {
		return 0
	}

	//出発駅に東京駅を設定する
	//var station string = "東京" としてしまうと次に求める駅を定義することができない
	//以下の通りにすると、最初に求めるのが東京なので、ループができる
	var next_station string = "東京"
	var distance_to_next_station int
	var total_distance int

	//駅数をiで記載し、上限30個の駅に達するまで同じ処理を繰り返す
	for i := 0; i < 30; i++ {
		//helperを用いて次の駅の名前を求め、next_stationに格納
		next_station = helper.OuterNextStation(next_station)
		//helperを用いて次の駅までの距離を求め、distance_to_next_stationに格納
		distance_to_next_station = helper.OuterLoopDistance(next_station)
		//目的地までの総距離を1回求めるごとに次の駅までの距離を加算し、total_distanceに格納
		total_distance = total_distance + distance_to_next_station
		//求める駅まで来たら、処理を終了する
		if next_station == station {
			break
		}
	}

	switch {
	//目的地までの総距離が3999m以下ならば、140円を返す
	case total_distance <= 3999:
		return 140
	//目的地までの総距離が4000m以上6999m以下ならば、160円を返す
	case total_distance >= 4000 && total_distance <= 6999:
		return 160
	//目的地までの総距離が7000m以上10999m以下ならば、170円を返す
	case total_distance >= 7000 && total_distance <= 10999:
		return 170
	//目的地までの総距離が11000m以上15999m以下ならば、200円を返す
	case total_distance >= 11000 && total_distance <= 15999:
		return 200
	//目的地までの総距離が21000m以上25999m以下ならば、270円を返す
	case total_distance >= 16000 && total_distance <= 20999:
		return 270
	//目的地までの総距離が21000m以上25999m以下ならば、350円を返す
	case total_distance >= 21000 && total_distance <= 25999:
		return 350
	//目的地までの総距離が26000m以上30999m以下ならば、420円を返す
	case total_distance >= 26000 && total_distance <= 30999:
		return 420
	//それ以外は490円を返す
	default:
		return 490
	}
}
