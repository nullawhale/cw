package cw

import (
	"testing"
)

func TestHero(t *testing.T) {
	var tests = []struct {
		params []int
		wait   bool
	}{
		{[]int{10, 5}, true},
		{[]int{7, 4}, false},
		{[]int{4, 5}, false},
		{[]int{100, 40}, true},
		{[]int{1500, 751}, false},
		{[]int{0, 1}, false},
	}

	for _, h := range tests {
		got := Hero(h.params[0], h.params[1])
		if got != h.wait {
			t.Errorf(
				"Expected %t, but got %t (bullets: %d, dragons: %d}",
				h.wait, got, h.params[0], h.params[1])
		}
	}
}
