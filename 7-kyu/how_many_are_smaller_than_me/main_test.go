package cw

import (
	"reflect"
	"testing"
)

func TestSmaller(t *testing.T) {
	var tests = []struct {
		got  []int
		want []int
	}{
		{[]int{5, 4, 3, 2, 1}, []int{4, 3, 2, 1, 0}},
		{[]int{1, 2, 3}, []int{0, 0, 0}},
		{[]int{1, 2, 0}, []int{1, 1, 0}},
		{[]int{1, 2, 1}, []int{0, 1, 0}},
		{[]int{1, 1, -1, 0, 0}, []int{3, 3, 0, 0, 0}},
		{[]int{5, 4, 7, 9, 2, 4, 4, 5, 6}, []int{4, 1, 5, 5, 0, 0, 0, 0, 0}},
	}

	for _, s := range tests {
		got := Smaller(s.got)
		if !reflect.DeepEqual(got, s.want) {
			t.Errorf("Expected \"%v\", but got \"%v\"", s.want, got)
		}
	}
}
