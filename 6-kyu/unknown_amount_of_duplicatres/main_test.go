package cw

import (
	"reflect"
	"testing"
)

func TestFindDupsMiss(t *testing.T) {
	var tests = []struct {
		got     []int
		missing int
		dups    []int
	}{
		{[]int{10, 9, 8, 9, 6, 1, 2, 4, 3, 2, 5, 5, 3}, 7, []int{2, 3, 5, 9}},
		{[]int{20, 19, 6, 9, 7, 17, 16, 17, 12, 5, 6, 8, 9, 10, 14, 13, 11, 14, 15, 19}, 18, []int{6, 9, 14, 17, 19}},
		{[]int{24, 25, 34, 40, 38, 26, 33, 29, 50, 31, 33, 56, 35, 36, 53, 49, 57, 27, 37, 40, 48, 44, 32, 35, 45, 52, 43, 47, 26, 51, 55, 28, 41, 42, 46, 51, 25, 30, 44, 54}, 39, []int{25, 26, 33, 35, 40, 44, 51}},
	}
	for _, s := range tests {
		missing, dups := FindDupsMiss(s.got)
		if missing != s.missing {
			t.Errorf("Expected \"%d\", but got \"%d\"", s.missing, missing)
		}
		if !reflect.DeepEqual(dups, s.dups) {
			t.Errorf("Expected \"%v\", but got \"%v\"\n\n", s.dups, dups)
		}
	}
}
