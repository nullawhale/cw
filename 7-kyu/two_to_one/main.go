package cw

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) (res string) {
	letters := strings.Fields(in)
	max, _ := strconv.Atoi(letters[0])
	min := max
	for _, c := range letters {
		num, _ := strconv.Atoi(c)
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	return fmt.Sprintf("%d %d", max, min)
}
