package chapter05

import "math"

func DarumaDrop(daruma []int) []int {
	// TODO: 実装
	if len(daruma) <= 1 {
		return daruma
	} else {
		tmp := math.Ceil(float64(len(daruma)) / 2)
		tmp2 := tmp - 1
		return append(daruma[0:int(tmp2)], daruma[int(tmp):len(daruma)]...)
	}
}

func MatrixMultiple(seed []int) [][]int {
	// TODO: 実装
	if len(seed) == 0 {
		return [][]int{}
	}
	/* 別回試す
	var result [][]int
	array := []int{}
	for i := 0; i < len(seed); i++ {
		for j := 0; j < len(seed); j++ {
			tmp := seed[i] * seed[j]
			array = append(array, int(tmp))
		}
		result = append(result, [][]int{array}...)
		array = []int{}
	}
	*/
	result := make([][]int, len(seed))
	for i, _ := range seed {
		result[i] = make([]int, len(seed))
	}
	for index, value := range seed {
		for index2, value2 := range seed {
			result[index][index2] = value * value2
		}
	}
	return result
}
