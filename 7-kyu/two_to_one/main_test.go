package cw

import (
	"testing"
)

func TestHighAndLow(t *testing.T) {
	var tests = []struct {
		got  string
		want string
	}{
		{"1 2 3 4 5", "5 1"},
		{"1 2 -3 4 5", "5 -3"},
		{"1 9 3 4 -5", "9 -5"},
		{"8 3 -5 42 -1 0 0 -9 4 7 4 -4", "42 -9"},
		{"1 2 3", "3 1"},
		{"42", "42 42"},
	}

	for _, s := range tests {
		got := HighAndLow(s.got)
		if got != s.want {
			t.Errorf("Expected \"%s\", but got \"%s\"", s.want, got)
		}
	}
}
