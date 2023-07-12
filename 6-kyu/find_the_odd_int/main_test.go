package cw

import (
	"testing"
)

func TestFindOdd(t *testing.T) {
	var tests = []struct {
		got  []int
		want int
	}{
		{[]int{7}, 7},
		{[]int{0}, 0},
		{[]int{1, 1, 2}, 2},
		{[]int{0, 1, 0, 1, 0}, 0},
		{[]int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1}, 4},
		{[]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}, 5},
		{[]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}, 5},
		{[]int{1, 1, 2, -2, 5, 2, 4, 4, -1, -2, 5}, -1},
		{[]int{20, 1, 1, 2, 2, 3, 3, 5, 5, 4, 20, 4, 5}, 5},
		{[]int{10}, 10},
		{[]int{1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1}, 10},
		{[]int{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10}, 1},
	}
	for _, s := range tests {
		got := FindOdd(s.got)
		if got != s.want {
			t.Errorf("Expected \"%d\", but got \"%d\"", s.want, got)
		}
	}
}
