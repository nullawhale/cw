package cw

import (
	"strconv"
)

func DigitalRoot(n int) int {
	return rec(strconv.Itoa(n))
}

func rec(s string) (res int) {
	for _, l := range s {
		num, _ := strconv.Atoi(string(l))
		res += num
	}
	if len(strconv.Itoa(res)) == 1 {
		return res
	} else {
		return rec(strconv.Itoa(res))
	}
}
