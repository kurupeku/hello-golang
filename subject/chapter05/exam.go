package chapter05

//import "math"

func DarumaDrop(daruma []int) []int {
	// TODO: 実装

	daruma2 := []int{}
	l := len(daruma)
	m := l / 2

	if l <= 1 {
		return daruma
	}

	if l%2 == 0 {
		if daruma[m-1] < daruma[m] {
			for _, v := range daruma {
				if v != daruma[m-1] {
					daruma2 = append(daruma2, v)
				}
			}
		} else {
			for _, v := range daruma {
				if v != daruma[m] {
					daruma2 = append(daruma2, v)
				}
			}
		}

	} else {
		for _, v := range daruma {
			if v != daruma[m] {
				daruma2 = append(daruma2, v)
			}
		}
	}

	return daruma2
}

func MatrixMultiple(seed []int) [][]int {
	// TODO: 実装

	matrix := [][]int{}

	for i := range seed {
		matrix2 := []int{}
		for j := range seed {
			matrix2 = append(matrix2, seed[i]*seed[j])
		}
		matrix = append(matrix, matrix2)
	}

	return matrix
}
