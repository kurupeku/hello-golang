package chapter05

func DarumaDrop(daruma []int) []int {
	darumaLength := len(daruma)
	if darumaLength == 0 || darumaLength == 1 {
		return daruma
	}

	sliceNum := darumaLength / 2
	if darumaLength%2 == 0 {
		return append(daruma[0:sliceNum-1], daruma[sliceNum:]...)
	} else {
		return append(daruma[0:sliceNum], daruma[sliceNum+1:]...)
	}
}

func MatrixMultiple(seed []int) [][]int {
	seedLength := len(seed)
	multiSlice := make([][]int, seedLength)

	for i, value := range seed {
		slice := make([]int, seedLength)
		for j := 0; j < seedLength; j++ {
			slice[j] = value * seed[j]
		}
		multiSlice[i] = slice
	}

	return multiSlice
}
