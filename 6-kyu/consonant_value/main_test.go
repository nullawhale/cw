package cw

import (
	"testing"
)

func TestTowerBuilder(t *testing.T) {
	var tests = []struct {
		got  string
		want int
	}{
		{"a", 0},
		{"bcd", 9},
		{"zea", 26},
		{"az", 26},
		{"baz", 26},
		{"aeiou", 0},
		{"uaoczei", 29},
		{"abababababfapeifapefijaefaepfjavnefjnfbhwyfnjsifjapnes", 143},
		{"codewars", 37},
		{"bup", 16},
	}
	for _, s := range tests {
		got := Solve(s.got)
		if got != s.want {
			t.Errorf("Expected \"%d\", but got \"%d\"", s.want, got)
		}
	}
}
