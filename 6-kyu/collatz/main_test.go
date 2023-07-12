package cw

import (
	"testing"
)

func TestCollatz(t *testing.T) {
	var tests = []struct {
		got  int
		want string
	}{
		{1, "1"},
		{2, "2->1"},
		{3, "3->10->5->16->8->4->2->1"},
		{4, "4->2->1"},
	}
	for _, s := range tests {
		got := Collatz(s.got)
		if got != s.want {
			t.Errorf("Expected \"%s\", but got \"%s\"", s.want, got)
		}
	}
}
