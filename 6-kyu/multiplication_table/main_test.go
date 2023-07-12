package cw

import (
	"reflect"
	"testing"
)

func TestMultiplicationTable(t *testing.T) {
	var tests = []struct {
		got  int
		want [][]int
	}{
		{1, [][]int{{1}}},
		{2, [][]int{{1, 2}, {2, 4}}},
		{3, [][]int{{1, 2, 3}, {2, 4, 6}, {3, 6, 9}}},
	}
	for _, s := range tests {
		got := MultiplicationTable(s.got)
		if !reflect.DeepEqual(got, s.want) {
			t.Errorf("Expected \"%v\", but got \"%v\"", s.want, got)
		}
	}
}
