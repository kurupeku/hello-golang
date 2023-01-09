package chapter05

import (
	"math"
)

func DarumaDrop(daruma []int) []int {
	if len(daruma) <= 1 {
		return daruma
	}
	i := int(math.Round(float64(len(daruma))/2)) - 1
	return append(daruma[0:i], daruma[i+1:]...)
}

func MatrixMultiple(seed []int) [][]int {
	tdArray := make([][]int, len(seed))
	for i := range seed {
		array := make([]int, len(seed))
		for j := range seed {
			array[j] = seed[i] * seed[j]
		}
		tdArray[i] = array
	}
	return tdArray
}
