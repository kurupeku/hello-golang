package chapter05

import "math"

func DarumaDrop(daruma []int) []int {
	l := len(daruma)
	if l <= 1 {
		return daruma
	}

	target := int(
		math.Ceil(float64(l)/2),
	) - 1

	newDaruma := make([]int, 0, l-1)
	for i, v := range daruma {
		if i != target {
			newDaruma = append(newDaruma, v)
		}
	}

	return newDaruma
}

func MatrixMultiple(seed []int) [][]int {
	mtx := make([][]int, 0, len(seed))
	for _, v1 := range seed {
		row := make([]int, 0, len(seed))
		for _, v2 := range seed {
			row = append(row, v1*v2)
		}
		mtx = append(mtx, row)
	}

	return mtx
}
