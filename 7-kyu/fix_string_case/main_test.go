package cw

import (
	"testing"
)

func TestSolve(t *testing.T) {
	var tests = []struct {
		got  string
		want string
	}{
		{"a", "a"},
		{"Z", "Z"},
		{"coDe", "code"},
		{"CODe", "CODE"},
		{"coDE", "code"},
		{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"},
	}

	for _, s := range tests {
		got := Solve(s.got)
		if got != s.want {
			t.Errorf("Expected \"%s\", but got \"%s\"", s.want, got)
		}
	}
}
