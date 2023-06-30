package cw

import (
	"testing"
)

func TestDigPow(t *testing.T) {
	var tests = []struct {
		n    int
		p    int
		want int
	}{
		{89, 1, 1},
		{92, 1, -1},
		{695, 2, 2},
		{46288, 3, 51},
	}
	for _, s := range tests {
		got := DigPow(s.n, s.p)
		if got != s.want {
			t.Errorf("Expected \"%d\", but got \"%d\"", s.want, got)
		}
	}
}
