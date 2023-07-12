package cw

import (
	"testing"
)

func TestCountBits(t *testing.T) {
	var tests = []struct {
		got  uint
		want int
	}{
		{0, 0},
		{4, 1},
		{7, 3},
		{9, 2},
		{10, 2},
	}
	for _, s := range tests {
		got := CountBits(s.got)
		if got != s.want {
			t.Errorf("Expected \"%d\", but got \"%d\"", s.want, got)
		}
	}
}
