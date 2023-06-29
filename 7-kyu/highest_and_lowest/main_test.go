package cw

import (
	"testing"
)

func TestTwoToOne(t *testing.T) {
	var tests = []struct {
		gotA string
		gotB string
		want string
	}{
		{"aretheyhere", "yestheyarehere", "aehrsty"},
		{"loopingisfunbutdangerous", "lessdangerousthancoding", "abcdefghilnoprstu"},
	}

	for _, s := range tests {
		got := TwoToOne(s.gotA, s.gotB)
		if got != s.want {
			t.Errorf("Expected \"%s\", but got \"%s\"", s.want, got)
		}
	}
}
