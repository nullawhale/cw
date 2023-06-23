package cw

import (
	"strconv"
	"strings"
)

func FakeBin(x string) string {
	a := strings.Split(x, "")
	for i, s := range a {
		_s, _ := strconv.Atoi(s)
		if _s < 5 {
			a[i] = "0"
		} else {
			a[i] = "1"
		}
	}
	return strings.Join(a, "")
}
