package helper

// 内回りの駅一覧
var innerStations = []string{
	"神田",
	"秋葉原",
	"御徒町",
	"上野",
	"鶯谷",
	"日暮里",
	"西日暮里",
	"田端",
	"駒込",
	"巣鴨",
	"大塚",
	"池袋",
	"目白",
	"高田馬場",
	"新大久保",
	"新宿",
	"代々木",
	"原宿",
	"渋谷",
	"恵比寿",
	"目黒",
	"五反田",
	"大崎",
	"品川",
	"高輪ゲートウェイ",
	"田町",
	"浜松町",
	"新橋",
	"有楽町",
	"東京",
}

var innerDistance = map[string]int{
	"神田":       1300,
	"秋葉原":      700,
	"御徒町":      1000,
	"上野":       600,
	"鶯谷":       1100,
	"日暮里":      1100,
	"西日暮里":     500,
	"田端":       800,
	"駒込":       1600,
	"巣鴨":       700,
	"大塚":       1100,
	"池袋":       1800,
	"目白":       1200,
	"高田馬場":     900,
	"新大久保":     1400,
	"新宿":       1300,
	"代々木":      700,
	"原宿":       1500,
	"渋谷":       1200,
	"恵比寿":      1600,
	"目黒":       1500,
	"五反田":      1200,
	"大崎":       900,
	"品川":       2000,
	"高輪ゲートウェイ": 900,
	"田町":       1300,
	"浜松町":      1500,
	"新橋":       1200,
	"有楽町":      1100,
	"東京":       800,
}

// 渡した駅の次の駅名を返す関数
// 内回りバージョン
//
// @Params station string 駅名
//
// @Return string 次の駅名
//
//	渡した駅から見て次の駅名を返す
func InnerNextStation(current string) string {
	for i, s := range innerStations {
		if s == current {
			ni := i + 1
			if ni == len(innerStations) {
				return innerStations[0]
			}

			return innerStations[ni]
		}
	}

	return current
}

// 一つ手前の駅から引数に渡した駅までの距離のメートルを int で返す関数
// 内回りバージョン
//
// @Params station string 駅名
//
//	距離を取得したい駅名
//
// @Return int 距離 (m)
//
//	一つ手前から渡した駅までの距離
func InnerLoopDistance(station string) int {
	dist, ok := innerDistance[station]
	if ok {
		return dist
	}

	return 0
}

// 外回りの駅一覧
var outerStations = []string{
	"有楽町",
	"新橋",
	"浜松町",
	"田町",
	"高輪ゲートウェイ",
	"品川",
	"大崎",
	"五反田",
	"目黒",
	"恵比寿",
	"渋谷",
	"原宿",
	"代々木",
	"新宿",
	"新大久保",
	"高田馬場",
	"目白",
	"池袋",
	"大塚",
	"巣鴨",
	"駒込",
	"田端",
	"西日暮里",
	"日暮里",
	"鶯谷",
	"上野",
	"御徒町",
	"秋葉原",
	"神田",
	"東京",
}

var outerDistance = map[string]int{
	"有楽町":      800,
	"新橋":       1100,
	"浜松町":      1200,
	"田町":       1500,
	"高輪ゲートウェイ": 1300,
	"品川":       900,
	"大崎":       2000,
	"五反田":      900,
	"目黒":       1200,
	"恵比寿":      1500,
	"渋谷":       1600,
	"原宿":       1200,
	"代々木":      1500,
	"新宿":       700,
	"新大久保":     1300,
	"高田馬場":     1400,
	"目白":       900,
	"池袋":       1200,
	"大塚":       1800,
	"巣鴨":       1100,
	"駒込":       700,
	"田端":       1600,
	"西日暮里":     500,
	"日暮里":      500,
	"鶯谷":       1100,
	"上野":       1100,
	"御徒町":      600,
	"秋葉原":      1000,
	"神田":       700,
	"東京":       1300,
}

// 渡した駅の次の駅名を返す関数
// 外回りバージョン
//
// @Params station string 駅名
//
// @Return string 次の駅名
//
//	渡した駅から見て次の駅名を返す
func OuterNextStation(current string) string {
	for i, s := range outerStations {
		if s == current {
			ni := i + 1
			if ni == len(outerStations) {
				return outerStations[0]
			}

			return outerStations[ni]
		}
	}

	return current
}

// 一つ手前の駅から引数に渡した駅までの距離のメートルを int で返す関数
// 外回りバージョン
//
// @Params station string 駅名
//
//	距離を取得したい駅名
//
// @Return int 距離 (m)
//
//	一つ手前から渡した駅までの距離
func OuterLoopDistance(station string) int {
	dist, ok := outerDistance[station]
	if ok {
		return dist
	}

	return 0
}
