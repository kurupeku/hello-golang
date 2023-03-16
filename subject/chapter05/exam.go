package chapter05

import (
	"fmt"
	"strconv"
)

func DarumaDrop(daruma []int) []int {
	// TODO: 実装
	fmt.Printf("daruma : %v\n", daruma)

	if len(daruma) == 0 || len(daruma) == 1 {
		return daruma
	}

	centerIndex := len(daruma) / 2
	println("centerIndex : " + strconv.Itoa(centerIndex))

	var newDaruma []int

	if len(daruma)%2 == 0 {
		newDaruma = append(daruma[:centerIndex-1], daruma[centerIndex:]...)
	} else {
		newDaruma = append(daruma[:centerIndex], daruma[centerIndex+1:]...)
	}

	fmt.Printf("newDaruma : %v\n", newDaruma)

	return newDaruma
}

func MatrixMultiple(seed []int) [][]int {
	// TODO: 実装

	outerSlice := make([][]int, 0, len(seed))

	fmt.Printf("matrix : %v\n", outerSlice)
	for i, v := range seed {
		fmt.Printf("i : %d, v : %d\n", i, v)
		innerSlice := make([]int, 0, len(seed))
		for j, v2 := range seed {
			fmt.Printf("j : %d, v2 : %d\n", j, v2)
			innerSlice = append(innerSlice, v*v2)
			fmt.Printf("innerMatrix : %v\n", innerSlice)
		}
		outerSlice = append(outerSlice, innerSlice)
	}
	return outerSlice
}
