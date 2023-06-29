package cw

import (
	"sort"
	"strings"
)

func TwoToOne(s1 string, s2 string) (res string) {
	s3 := strings.Split(s1+s2, "")
	sort.Strings(s3)

	for _, l := range s3 {
		if !strings.Contains(res, l) {
			res += l
		}
	}
	return res
}
