package chapter09

import "github.com/kurupeku/hello-golang/helper"

// 料金計算時の距離下限
var lowerDistanceLimit = []int{1, 4000, 7000, 11000, 16000, 21000, 26000, 31000}

// 料金計算時の距離上限
var upperDistanceLimit = []int{3999, 6999, 10999, 15999, 20999, 25999, 30999, 99999}

// 切符の料金リスト
var ticketFareList = []int{140, 160, 170, 200, 270, 350, 420, 490}

// ICカードの料金リスト
var cardFareList = []int{136, 157, 168, 198, 264, 341, 418, 484}

type Fare struct {
	From string
	To   string
}

func (f *Fare) innerDistance() int {
	// TODO: 内回りの距離を返すように実装

	// 現在の停車駅をcurrentStationに格納
	currentStation := f.From

	// 距離を初期化
	distance := 0

	// 引数で渡された到着駅までループし、InnerLoopDistanceで計算された距離をdistanceへ格納
	for {
		if currentStation != f.To {
			currentStation = helper.InnerNextStation(currentStation)
			distance += helper.InnerLoopDistance(currentStation)
		} else {
			break
		}
	}
	return distance
}

func (f *Fare) outerDistance() int {
	// TODO: 外回りの距離を返すように実装

	// 現在の停車駅をcurrentStationに格納
	currentStation := f.From

	// 距離を初期化
	distance := 0

	// 引数で渡された到着駅までループし、InnerLoopDistanceで計算された距離をdistanceへ格納
	for {
		if currentStation != f.To {
			currentStation = helper.OuterNextStation(currentStation)
			distance += helper.OuterLoopDistance(currentStation)
		} else {
			break
		}
	}
	return distance
}

func (f *Fare) distance() int {
	// 内回りと外回りの距離を比較し、短い方を返すように実装

	// 内回りが外回りより距離が長い場合は外回りの距離を返す
	if f.innerDistance() > f.outerDistance() {
		return f.outerDistance()
	}
	// それ以外は内回りの距離を返す
	return f.innerDistance()
}

func (f *Fare) TicketCharge() int {
	// TODO: 切符の料金を返すように実装

	// 距離を計算してdistanceに格納
	distance := f.distance()

	// 乗車料金trainFareの初期化
	trainFare := 0

	// 料金のリスト分だけループ
	for i := 0; i < len(ticketFareList); i++ {
		// 計算した距離が、料金計算時の距離下限以上かつ料金計算時の距離上限以下の場合に料金リストに格納された料金をtrainFareに格納
		if distance >= lowerDistanceLimit[i] && distance <= upperDistanceLimit[i] {
			trainFare = ticketFareList[i]
			break
		}
		// 上記のifに当てはまらない場合は0を格納(出発地と目的地が同じ場合等)
		trainFare = 0
	}

	// 計算された料金を返す
	return trainFare

	//return f.Charge("ticket")
}

func (f *Fare) CardCharge() int {
	// TODO: IC カードの料金を返すように実装

	// 距離を計算してdistanceに格納
	distance := f.distance()

	// 乗車料金trainFareの初期化
	trainFare := 0

	// 料金のリスト分だけループ
	for i := 0; i < len(cardFareList); i++ {
		// 計算した距離が、料金計算時の距離下限以上かつ料金計算時の距離上限以下の場合に料金リストに格納された料金をtrainFareに格納
		if distance >= lowerDistanceLimit[i] && distance <= upperDistanceLimit[i] {
			trainFare = cardFareList[i]
			break
		}
		// 上記のifに当てはまらない場合は0を格納(出発地と目的地が同じ場合等)
		trainFare = 0
	}

	// 計算された料金を返す
	return trainFare

	//return f.Charge("card")
}

// 関数集約しようとした残骸↓↓↓
/*func (f *Fare) Charge(item string) int {

	// 距離を取得してdistanceへ格納
	distance := f.distance()

	trainFare := 0
	fareList := []int{}

	switch item {
	case "ticket":
		fareList = ticketFareList[:]
	case "card":
		fareList = cardFareList[:]
	}

	// チケットの料金リスト分だけループ
	for i := 0; i < len(fareList); i++ {
		// 計算した距離が、料金計算時の距離下限以上かつ料金計算時の距離上限以下の場合に料金リストに格納された料金をtrainFareに格納
		if distance >= lowerDistanceLimit[i] && distance <= upperDistanceLimit[i] {
			trainFare = fareList[i]
			break
		}
		// 上記のifに当てはまらない場合は0を格納(出発地と目的地が同じ場合等)
		trainFare = 0
	}
}*/
