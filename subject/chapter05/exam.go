package chapter05

func DarumaDrop(daruma []int) []int {
	if dl := len(daruma); dl == 0 || dl == 1 {
		return daruma
	}

	l := len(daruma) % 2
	m := len(daruma) / 2
	if l != 0 {
		daruma = daruma[:m+copy(daruma[m:], daruma[m+1:])]
	} else {
		if daruma[m] > daruma[m-1] {
			daruma = daruma[:m-1+copy(daruma[m-1:], daruma[m:])]
		} else {
			daruma = daruma[:m+copy(daruma[m:], daruma[m+1:])]
		}
	}

	return daruma
}

func MatrixMultiple(seed []int) [][]int {
	l := len(seed)
	multiple := make([][]int, l)
	if l == 0 {
		return multiple
	}

	for i := 0; i < l; i++ {
		multiple[i] = make([]int, l)
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			multiple[i][j] = seed[i] * seed[j]
		}
	}
	return multiple
}
