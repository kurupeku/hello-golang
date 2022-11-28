package chapter05

func DarumaDrop(daruma []int) []int {
	// TODO: 実装

	// 配列の長さを取得
	darumaLength := len(daruma)

	// 要素数１個以下はスルー
	if darumaLength < 2 { return daruma }

	// 奇数偶数判定用
	fractionNum := darumaLength % 2

	// 両端に必ず残る個数
	keepLength := darumaLength / 2 + fractionNum - 1

	// 最初に冒頭だけ切り取る
	excludedDaruma := daruma[:keepLength]

	// 偶数個の場合、真ん中の２個を比較して大きい方を取得
	if largerNum := &daruma[keepLength]; fractionNum == 0 {
		if *largerNum < daruma[keepLength + 1] { largerNum = &daruma[keepLength + 1] }
		excludedDaruma = append(excludedDaruma, *largerNum)
	}

	// 末尾を連結して返す
	return append(excludedDaruma, daruma[darumaLength - keepLength:]...)
}

func MatrixMultiple(seed []int) [][]int {

	// TODO: 実装

	// 要素数を取得して、戻り値用の配列を作成
	seedLength := len(seed)
	matrixMultipleArray := make([][]int, seedLength)

	for i, v := range seed {

		// 配列in配列
		matrixMultipleArray[i] = make([]int, seedLength)

		// 値を格納
		for t, w := range seed {
			matrixMultipleArray[i][t] = v * w
		}
	}

	return matrixMultipleArray
}
