package chapter05

import "math"

// 渡された `Slice` の真ん中の要素をふっ飛ばした `Slice` を返す関数を実装してください
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

// 渡された `Slice` 同士をかけ合わせた `Slice in Slice` の二次元配列を返す関数を実装してください
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
