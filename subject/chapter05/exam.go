package chapter05

func DarumaDrop(daruma []int) []int {
	// TODO: 実装
	// 中央値(切り捨て)
	var median int = len(daruma) / 2

	// dropした後のdaruma(slice)
	var afterDaruma []int

	// darumaのlengthが1以下の場合はdarumaをそのまま返す
	if len(daruma) <= 1 {
		return daruma

		// darumaのlengthを二分したあまりが0の場合 かつ darumaのlengthを二分した値の配列番号に格納されている値の方が大きい場合 にmedianをマイナス1する
		// 詳細：DarumaDrop関数のメモを参照
	} else if len(daruma)%2 == 0 && daruma[median] > daruma[median-1] {
		median--
	}

	// darumaの要素分ループ(indexは捨てて、値はiへ格納)
	for _, value := range daruma {
		// rangeで取得したdarumaの値iが、中央値と異なる場合はappendでafterDarumaに要素を追加する
		if value != daruma[median] {
			afterDaruma = append(afterDaruma, value)
		}
	}

	// afterDarumaを返す
	return afterDaruma
}

func MatrixMultiple(seed []int) [][]int {
	// TODO: 実装
	// 最終的な計算結果
	var result [][]int

	// seedが0の場合は[][]int{}を返す
	if len(seed) == 0 {
		return [][]int{}
	}

	// seedのindex分ループ
	for i, _ := range seed {
		// 途中の計算結果（iでループする度にリセットされる）
		var calcResult []int
		// seedのvalue分ループ
		for _, value := range seed {
			// calcResultにseed[i]とvalueを掛けた結果を格納
			calcResult = append(calcResult, seed[i]*value)
		}
		// resultにcalcResultを格納（1行ずつ）
		result = append(result, calcResult)
	}
	// resultを返す
	return result
}

// DarumaDrop関数のメモ
// 偶数
// 1 2 "3" 4 5 6  ←length
// 0 1 "2 3" 4 5  ←配列番号
// この場合lengthを二分すると3だけど比較は 配列番号3 と 配列番号2(3-1) になる

// 奇数
// 1 "2" 3 4 5  ←length
// 0 1 "2" 3 4  ←配列番号
// この場合lengthを二分すれば配列番号てきにはちょうど2（5/2は切り捨てで2になるから）
