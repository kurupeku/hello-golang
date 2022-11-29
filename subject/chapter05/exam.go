package chapter05

func DarumaDrop(daruma []int) []int {
	// TODO: 実装（無理やり）
	darumaResult := daruma
	darumaLen := len(daruma)
	var length int
	if darumaLen < 2 {
		return daruma
	} else if darumaLen%2 == 0 {
		length = darumaLen / 2

		if daruma[length] > daruma[length-1] {
			length -= 1
		}

	} else if darumaLen%2 == 1 {
		length = darumaLen / 2
	}

	if length == 0 {
		darumaResult = daruma[length+1:]
	} else {
		darumaResult = append(daruma[:length], daruma[length+1:]...)
	}

	return darumaResult
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
