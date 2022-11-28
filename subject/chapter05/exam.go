package chapter05

//import "fmt"

func DarumaDrop(daruma []int) []int {
	target := 0

	switch {
	case len(daruma) < 2:
		return daruma
	case len(daruma) == 2:
		return daruma[1:]
	case len(daruma)%2 == 0: //even
		target = len(daruma)/2 - 1
	case len(daruma)%2 == 1: //odd
		target = len(daruma) / 2
	}
	//fmt.Printf("daruma-len is %v\n", len(daruma))

	// 第二引数のsliceを ...で要素展開
	return append(daruma[:target], daruma[target+1:]...)

}

func MatrixMultiple(seed []int) [][]int {
	matrix := make([][]int, len(seed))

	for i := 0; i < len(seed); i++ {
		matrix[i] = make([]int, len(seed))

		for j := 0; j < len(seed); j++ {
			matrix[i][j] = seed[i] * seed[j]
		}
		//fmt.Printf("%d\n", matrix[i])
	}

	return matrix
}
