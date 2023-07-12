package cw

func Transpose(matrix [][]int) [][]int {

	n, m := len(matrix), len(matrix[0])

	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			res[i][j] = matrix[j][i]
		}
	}
	return res
}
