package cw

import (
	"strconv"
)

func DeleteDigit(n int) int {
	str := strconv.Itoa(n)
	max := 0
	for i := 1; i <= len(str); i++ {
		num, _ := strconv.Atoi(str[:i-1] + str[i:])
		if num > max {
			max = num
		}
	}
	return max
}
