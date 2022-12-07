package chapter05

func getDropDaruma(dropNum int, daruma []int) []int {
	var ans []int;
	for i, v := range daruma {
		if (i != dropNum) {
			ans = append(ans, v)
		}
	}
	return ans;
}

func DarumaDrop(daruma []int) []int {
	darumaLen := len(daruma);

	if (darumaLen == 0) {
		return []int{};
	}

	if (darumaLen == 1) {
		return daruma;
	}

	if (darumaLen % 2 == 1) {
		dropNum := int(darumaLen/2)
		return getDropDaruma(dropNum, daruma);
	} else {
		dropNum := darumaLen/2 - 1
		return getDropDaruma(dropNum, daruma);
	}
}

func MatrixMultiple(seed []int) [][]int {
	seedLen := len(seed)

	if (seedLen == 0) {
		return [][]int{}
	}

	var ans [][]int
	for _, v := range seed {
		var column []int
		for j := 0; j < seedLen; j++ {
			column = append(column, v * seed[j])
		}
		ans = append(ans, column)
	}

	return ans
}
