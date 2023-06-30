package cw

import (
	"math"
	"strconv"
)

func DigPow(n, p int) int {
	qwe := IntToSlice(n, p+len(strconv.Itoa(n)), 0)

	if qwe%n == 0 {
		return qwe / n
	} else {
		return -1
	}
}

func IntToSlice(n, exp int, res float64) int {
	if n != 0 {
		i := n % 10
		return IntToSlice(n/10, exp-1, res+math.Pow(float64(i), float64(exp-1)))
	}
	return int(res)
}
