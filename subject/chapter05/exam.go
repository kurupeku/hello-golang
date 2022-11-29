package chapter05

func DarumaDrop(daruma []int) []int {
	// TODO: 実装

	// 2022/11/29 Yasuhiro Fujii
	// 時間がないから超やっつけ

	var tmp int
	var tmp2 int
	var length int
	length = len(daruma)
	tmp = length % 2
	tmp2 = length / 2

	if length < 2 {
		return daruma
	} else if length == 2 {
		array_tmp := make([]int, length-1)
		array_tmp[0] = daruma[1]
		return array_tmp
	}

	var i int
	var array_tmp3 []int
	// 偶数
	if tmp == 0 {
		for i = 0; i < length; i++ {
			if i != (tmp2 - 1) {
				array_tmp3 = append(array_tmp3, daruma[i])
			}
		}
		return array_tmp3
	} else {
		var array_tmp3 []int
		for i = 0; i < length; i++ {
			if i != (tmp2) {
				array_tmp3 = append(array_tmp3, daruma[i])
			}
		}
		return array_tmp3
	}
}

func MatrixMultiple(seed []int) [][]int {
	// TODO: 実装
	var tmp int
	array := make([][]int, len(seed))
	if len(seed) == 0 {
		return [][]int{}
	}
	for i := range seed {
		array[i] = make([]int, len(seed))
		for j := range seed {
			tmp = seed[i] * seed[j]
			array[i][j] = tmp
		}
	}
	return array
}
