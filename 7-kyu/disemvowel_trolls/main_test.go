package cw

import (
	"testing"
)

func TestDisemvowel(t *testing.T) {
	var tests = []struct {
		got  string
		want string
	}{
		{"This website is for losers LOL!", "Ths wbst s fr lsrs LL!"},
	}

	for _, s := range tests {
		got := Disemvowel(s.got)
		if got != s.want {
			t.Errorf("Expected \"%s\", but got \"%s\"", s.want, got)
		}
	}
}
