package chapter03

import "github.com/kurupeku/hello-golang/helper"

// 出発駅
const departureStation = "東京"

// 料金計算時の距離下限
var lowerDistanceLimit = []int{1, 4000, 7000, 11000, 16000, 21000, 26000, 31000}

// 料金計算時の距離上限
var upperDistanceLimit = []int{3999, 6999, 10999, 15999, 20999, 25999, 30999, 99999}

// 料金のリスト
var trainFareList = []int{140, 160, 170, 200, 270, 350, 420, 490}

// If を使用して料金の条件分岐を行ってください
func InnerChargeFromTokyo(station string) int {
	// TODO: 実装
	currentStation := departureStation
	distance, trainFare := 0, 0

	// 引数で渡された到着駅までループし、InnerLoopDistancedistanceで計算された距離をdistanceへ格納
	for {
		if currentStation != station {
			currentStation = helper.InnerNextStation(currentStation)
			distance += helper.InnerLoopDistance(currentStation)
		} else {
			break
		}
	}

	// 料金のリスト分だけループ
	for i := 0; i < len(trainFareList); i++ {
		// 計算した距離が、料金計算時の距離下限以上かつ料金計算時の距離上限以下の場合に料金リストに格納された料金をtrainFareに格納
		if distance >= lowerDistanceLimit[i] && distance <= upperDistanceLimit[i] {
			trainFare = trainFareList[i]
			break
		}
		// 上記のifに当てはまらない場合は0を格納(目的地が東京の場合等)
		trainFare = 0
	}

	// 計算された料金を返す
	return trainFare
}

// Switch を使用して料金の条件分岐を行ってください
func OuterChargeFromTokyo(station string) int {
	// TODO: 実装
	currentStation := departureStation
	distance, trainFare := 0, 0

	// 駅数だけループ（駅数を取得したい...）
	for i := 0; i < 35; i++ {
		// 引数の到着駅が東京駅以外かつ、引数の到着駅と現在の駅が異なる場合に次の駅をcurrentStationに格納しそこまでの距離をdistanceに格納
		// それ以外の場合は何もしない
		switch {
		case station != departureStation && station != currentStation:
			currentStation = helper.OuterNextStation(currentStation)
			distance += helper.OuterLoopDistance(currentStation)
		default:
		}
	}

	// 料金のリスト分だけループ
	for i := 0; i < len(trainFareList); i++ {
		switch {
		// 計算した距離が、料金計算時の距離下限以上かつ料金計算時の距離上限以下の場合に料金リストに格納された料金をtrainFareに格納
		case distance >= lowerDistanceLimit[i] && distance <= upperDistanceLimit[i]:
			trainFare = trainFareList[i]
		default:
			// 上記のに当てはまらない場合は何もしない = 0 (目的地が東京の場合等)
		}
	}

	// 計算された料金を返す
	return trainFare
}
