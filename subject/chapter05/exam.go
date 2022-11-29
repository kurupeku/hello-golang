package chapter05

func DarumaDrop(daruma []int) []int {
	// TODO: 実装

	// 配列の長さを取得
	darumaLength := len(daruma)

	// 要素数１個以下はスルー
	if darumaLength < 2 { return daruma }

	// 冒頭の個数
	keepLength := (darumaLength - 1) / 2

	// 冒頭だけ取得
	excludedDaruma := daruma[:keepLength]

	// 末尾を連結
	return append(excludedDaruma, daruma[keepLength+1:]...)
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
