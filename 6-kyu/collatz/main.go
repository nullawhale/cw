package cw

import (
	"strconv"
)

func Collatz(n int) (res string) {
	res += strconv.Itoa(n)
	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		res += "->" + strconv.Itoa(n)
	}
	return
}
