package cw

import (
	"testing"
)

func TestHigh(t *testing.T) {
	var tests = []struct {
		got  string
		want string
	}{
		{"man i need a taxi up to ubud", "taxi"},
		{"what time are we climbing up the volcano", "volcano"},
		{"take me to semynak", "semynak"},
		{"aa b", "aa"},
		{"b aa", "b"},
		{"bb d", "bb"},
		{"d bb", "d"},
		{"aaa b", "aaa"},
	}

	for _, s := range tests {
		got := High(s.got)
		if got != s.want {
			t.Errorf("Expected \"%s\", but got \"%s\"", s.want, got)
		}
	}
}
