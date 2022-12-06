package chapter05

func DarumaDrop(daruma []int) []int {
	// TODO: 実装（無理やり）

	darumaLen := len(daruma)

	if darumaLen < 2 {
		return daruma
	}

	keepLength := (darumaLen - 1) / 2

	return append(daruma[:keepLength], daruma[keepLength+1:]...)
}

func MatrixMultiple(seed []int) [][]int {
	// TODO: 実装(王道)
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
