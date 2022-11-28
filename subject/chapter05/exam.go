package chapter05

import (
	"math"
)

func DarumaDrop(daruma []int) []int {
	daruma_len := len(daruma)
	if daruma_len <= 1 {
		return daruma
	}

	if daruma_len == 2 {
		daruma = daruma[1:]
		return daruma
	}

	var first_half, latter_half []int

	if (daruma_len % 2) == 0 {
		number_of_pulled_out := int(math.Floor(float64(daruma_len / 2)))
		first_half = daruma[:number_of_pulled_out-1]
		latter_half = daruma[number_of_pulled_out:]
	}

	if (daruma_len % 2) == 1 {
		number_of_pulled_out := int(math.Ceil(float64(daruma_len / 2)))
		first_half = daruma[:number_of_pulled_out]
		latter_half = daruma[number_of_pulled_out+1:]
	}

	daruma = append(first_half, latter_half...)
	return daruma
}

func MatrixMultiple(seed []int) [][]int {
	var answer [][]int
	var pre_answer []int

	if len(seed) == 0 {
		return [][]int{}
	}

	for _, v := range seed {
		pre_answer = nil
		for _, x := range seed {
			pre_answer = append(pre_answer, v*x)
		}
		answer = append(answer, pre_answer)
	}
	return answer

}
