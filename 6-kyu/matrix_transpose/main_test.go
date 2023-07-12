package cw

import (
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {
	var tests = []struct {
		got  [][]int
		want [][]int
	}{
		{
			[][]int{{1}},
			[][]int{{1}},
		},
		{
			[][]int{{1, 2, 3}},
			[][]int{{1}, {2}, {3}},
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}},
			[][]int{{1, 4}, {2, 5}, {3, 6}},
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
		},
		{
			[][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}, {0, 1, 0}, {1, 0, 0}},
			[][]int{{1, 0, 0, 0, 1}, {0, 1, 0, 1, 0}, {0, 0, 1, 0, 0}},
		},
	}
	for _, s := range tests {
		got := Transpose(s.got)
		if !reflect.DeepEqual(got, s.want) {
			t.Errorf("Expected \"%v\", but got \"%v\"", s.want, got)
		}
	}
}
