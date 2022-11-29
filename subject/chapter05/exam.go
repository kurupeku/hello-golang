package chapter05

import "math"

func DarumaDrop(daruma []int) []int {
	// TODO: 実装
	length := float64(len(daruma))
	new_array := make([]int, 0)

	switch length {
	case 0:
		return []int{}
	case 1:
		return []int{1}
	default:
		exclude_num := math.Round(float64(length / 2))
		for i := range daruma {
			switch {
			case daruma[i] != int(exclude_num):
				new_array = append(new_array, daruma[i])
			default:
				continue
			}
		}
		return new_array
	}

}

func MatrixMultiple(seed []int) [][]int {
	// TODO: 実装
	length := len(seed)
	calc := make([][]int, length)
	for i := range seed {
		calc[i] = make([]int, length)
	}

	switch length {
	case 0:
		return [][]int{}
	case 1:
		return [][]int{{1}}
	default:
		for i := range seed {
			for j := range seed {
				value := seed[i] * seed[j]
				calc[i][j] = value
			}
		}
		return calc
	}
}
