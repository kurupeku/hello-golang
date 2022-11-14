package chapter03

import (
	"fmt"

	"github.com/kurupeku/hello-golang/helper"
)

// If を使用して料金の条件分岐を行ってください
func InnerChargeFromTokyo(station string) int {
	// TODO: 実装

	// next_station（隣の駅）を文字列として扱う
	var next_station string
	// distance_to_next_station（隣の駅までの距離）を整数として扱う
	var distance_to_next_station int
	// start_station（出発駅：東京）を文字列として扱う
	var start_station string
	// total_distance（総距離）を整数として扱う
	var total_distance int
	//next_station にてhelperを用いて次の駅を出す
	next_station = helper.InnerNextStation(station)
	//distance_to_next_station にてhelperを用いて次の駅までの距離を出す
	distance_to_next_station = helper.InnerLoopDistance(station)

	// 問題分のテーブルからpriceに料金を格納する
	//var price int = (140, 160, 170, 200, 270, 350, 420, 490)
	// 問題分のテーブルからdistanceに距離を格納する
	//var distance_station int = (3999, 6999, 10999, 15999, 20999, 25999, 30999)

	//東京駅を出発駅に設定して、次の駅を求める
	if start_station == "東京" {
		fmt.Println(next_station)
	}
	//東京駅から次の駅までの距離（m）を求める
	if start_station == "東京" {
		fmt.Println(distance_to_next_station)
	}

	//駅とその次の駅、その2つの駅間の距離を表示する
	fmt.Println(station)
	fmt.Println(next_station)
	fmt.Println(distance_to_next_station)
	//fmt.Println(price)

	//目的地までの総距離を合計する。40000m未満となるまで合計する
	for total_distance := 0; total_distance < 40000; total_distance++ {

		//総距離には隣の駅までの距離を合計する
		total_distance += distance_to_next_station
		//総距離が40000mに至った場合、処理を終了する
		if total_distance == 40000 {
			break
		}
	}

	//総距離を表示する
	fmt.Println(total_distance)
	//fmt.Println(price)
	return 0
}

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	// TODO: 実装

	// next_stationを文字列として扱う
	var next_station string
	// distance_to_next_stationを整数として扱う
	var distance_to_next_station int
	// start_stationを文字列として扱う
	var start_station string
	//next_station にてhelperを用いて次の駅を呼び出す
	next_station = helper.OuterNextStation(station)
	//distance_to_next_station にてhelperを用いて次の駅までの距離を出す
	distance_to_next_station = helper.OuterLoopDistance(station)

	//東京駅を出発駅に設定して、次の駅を求める
	start_station = "東京"
	switch start_station {
	case "東京":
		//駅とその次の駅、当駅間の距離を表示する
		fmt.Println(station)
		fmt.Println(next_station)
		fmt.Println(distance_to_next_station)
	}
	return 0
}
