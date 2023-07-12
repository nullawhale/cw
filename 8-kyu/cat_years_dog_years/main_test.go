package cw

import (
	"testing"
)

func TestCalculateYears(t *testing.T) {
	var tests = []struct {
		got  int
		want [3]int
	}{
		{1, [3]int{1, 15, 15}},
		{2, [3]int{2, 24, 24}},
		{10, [3]int{10, 56, 64}},
	}

	for _, s := range tests {
		got := CalculateYears(s.got)
		if got != s.want {
			t.Errorf("Expected \"%v\", but got \"%v\"", s.want, got)
		}
	}
}
