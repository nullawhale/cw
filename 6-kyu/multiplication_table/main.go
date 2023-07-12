package cw

func MultiplicationTable(size int) [][]int {
	res := make([][]int, size)
	for i := 0; i < size; i++ {
		res[i] = make([]int, size)
		for j := 0; j < size; j++ {
			res[i][j] = (i + 1) * (j + 1)
		}
	}
	return res
}
