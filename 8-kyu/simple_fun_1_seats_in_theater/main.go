package cw

func SeatsInTheater(nCols int, nRows int, col int, row int) int {
	return (nCols - (col - 1)) * (nRows - row)
}
