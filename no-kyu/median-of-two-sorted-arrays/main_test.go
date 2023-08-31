package cw

import (
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	var tests = []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{1, 3, 5}, []int{2, 4, 6}, 3.5},
	}

	for _, s := range tests {
		got := FindMedianSortedArrays(s.nums1, s.nums2)
		if got != s.want {
			t.Errorf("Expected \"%f\", but got \"%f\"", s.want, got)
		}
	}
}
