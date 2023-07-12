package cw

import (
	"testing"
)

func TestIsLeapYear(t *testing.T) {
	var tests = []struct {
		got  int
		want bool
	}{
		{2020, true},
		{2000, true},
		{2015, false},
		{2100, false},
	}

	for _, s := range tests {
		got := IsLeapYear(s.got)
		if got != s.want {
			t.Errorf("Expected \"%t\", but got \"%t\"", s.want, got)
		}
	}
}
