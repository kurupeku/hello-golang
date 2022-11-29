package chapter05

import (
	"math"
)

func DarumaDrop(daruma []int) []int {
	if len(daruma) <= 1 {
		return daruma
	}
	target := int(math.Ceil(float64(len(daruma))/2)) - 1
	return append(daruma[0:target], daruma[target+1:]...)
}

func MatrixMultiple(seed []int) [][]int {
	col := make([][]int, len(seed))
	for i := range seed {
		row := make([]int, len(seed))
		for j := range seed {
			row[j] = seed[i] * seed[j]
		}
		col[i] = row
	}
	return col
}
