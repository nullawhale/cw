package cw

import (
	"testing"
)

func TestIsValidWalk(t *testing.T) {
	var tests = []struct {
		got  []rune
		want bool
	}{
		{[]rune{'n', 's', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}, true},
		{[]rune{'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e', 'w', 'e'}, false},
		{[]rune{'w'}, false},
		{[]rune{'n', 'n', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}, false},
		{[]rune{'e', 'e', 'e', 'e', 'w', 'w', 's', 's', 's', 's'}, false},
		{[]rune{'s', 'w', 's', 'w', 'w', 'w', 's', 'e', 'w', 's'}, false},
	}
	for _, s := range tests {
		got := IsValidWalk(s.got)
		if got != s.want {
			t.Errorf("Expected \"%t\", but got \"%t\"", s.want, got)
		}
	}
}
