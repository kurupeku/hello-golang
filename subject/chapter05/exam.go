package chapter05

import (
	"fmt"
	"math"
)

func DarumaDrop(daruma []int) []int {
	// TODO: 実装
	fmt.Println(len(daruma))
	darumaLen := len(daruma)
	darumaPosition := int(math.Ceil(float64(darumaLen) / 2))

	if darumaLen <= 1 {
		return daruma
	} else {
		front := daruma[0 : darumaPosition-1]
		fmt.Println(front)
		back := daruma[darumaPosition:darumaLen]
		fmt.Println(back)
		newDaruma := append(front, back...)
		return newDaruma
	}
}

func MatrixMultiple(seed []int) [][]int {
	// TODO: 実装
	fmt.Println(seed)
	tree := make([][]int, len(seed))
	for i, v := range seed {
		tree[i] = make([]int, len(seed))
		for j, k := range seed {
			tree[i][j] = v * k
			fmt.Println(tree[i][j])
		}

	}
	return tree
	/*
			0*0 0*1 0*2
			1*0 1*1 1*2
		    2*0 2*1 2*2
	*/

}
