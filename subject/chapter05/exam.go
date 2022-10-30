package chapter05

func DarumaDrop(daruma []int) []int {
	if len(daruma) <= 1 {
		return daruma
	}
	var drop int
	if len(daruma)%2 == 0 {
		drop = len(daruma)/2 - 1
	} else {
		drop = len(daruma) / 2
	}
	daruma = append(daruma[:drop], daruma[drop+1:]...)
	return daruma
}

func MatrixMultiple(seed []int) [][]int {
	var result = make([][]int, len(seed))
	for i, _ := range result {
		result[i] = make([]int, len(seed))
	}
	for i, v1 := range seed {
		for j, v2 := range seed {
			result[i][j] = v1 * v2
		}
	}
	return result
}
