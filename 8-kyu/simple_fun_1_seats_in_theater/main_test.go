package cw

import (
	"testing"
)

func TestSeatsInTheater(t *testing.T) {
	var tests = []struct {
		params []int
		wait   int
	}{
		{[]int{16, 11, 5, 3}, 96},
		{[]int{1, 1, 1, 1}, 0},
		{[]int{13, 6, 8, 3}, 18},
		{[]int{60, 100, 60, 1}, 99},
		{[]int{1000, 1000, 1000, 1000}, 0},
	}

	for _, h := range tests {
		got := SeatsInTheater(h.params[0], h.params[1], h.params[2], h.params[3])
		if got != h.wait {
			t.Errorf(
				"Expected %d, but got %d (nCols: %d, nRows: %d, col: %d, row: %d}",
				h.wait, got, h.params[0], h.params[1], h.params[2], h.params[3])
		}
	}
}
