package cw

import (
	"testing"
)

func TestDigitalRoot(t *testing.T) {
	var tests = []struct {
		got  int
		want int
	}{
		{16, 7},
		{195, 6},
		{992, 2},
		{167346, 9},
		{0, 0},
	}
	for _, s := range tests {
		got := DigitalRoot(s.got)
		if got != s.want {
			t.Errorf("Expected \"%d\", but got \"%d\"", s.want, got)
		}
	}
}
