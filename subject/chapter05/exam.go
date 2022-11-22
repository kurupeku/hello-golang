package chapter05

func DarumaDrop(daruma []int) []int {
	// 達磨の保護
	if len(daruma) < 2 {
		return daruma

	}

	// 真ん中の添え字
	sprit := (len(daruma) - 1) / 2

	// 新たな配列を作成
	after_daruma := make([]int, 0, len(daruma)-1)
	for i, val := range daruma {
		if i == sprit {
			continue
		}
		after_daruma = append(after_daruma, val)
	}
	return after_daruma
}

func MatrixMultiple(seed []int) [][]int {
	// TODO: 実装
	matrix := make([][]int, 0, len(seed))
	for i := range seed {
		table := make([]int, 0, len(seed))
		for _, val := range seed {
			table = append(table, seed[i]*val)
		}
		matrix = append(matrix, table)
	}

	return matrix
}
