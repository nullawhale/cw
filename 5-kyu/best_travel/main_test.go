package cw

import (
	"reflect"
	"testing"
)

func TestChooseBestSum(t *testing.T) {
	var tests = []struct {
		t    int
		k    int
		ls   []int
		want int
	}{
		{163, 3, []int{50, 55, 56, 57, 58}, 163},
		{163, 3, []int{50}, -1},
		{230, 3, []int{91, 74, 73, 85, 73, 81, 87}, 228},
		{331, 2, []int{91, 74, 73, 85, 73, 81, 87}, 178},
		{331, 4, []int{91, 74, 73, 85, 73, 81, 87}, 331},
		{331, 5, []int{91, 74, 73, 85, 73, 81, 87}, -1},
	}
	for _, s := range tests {
		got := ChooseBestSum(s.t, s.k, s.ls)
		if !reflect.DeepEqual(got, s.want) {
			t.Errorf("Expected \"%d\", but got \"%d\"", s.want, got)
		}
	}
}
