package cw

import (
	"testing"
)

func TestFindMissingLetter(t *testing.T) {
	var tests = []struct {
		got  []rune
		want rune
	}{
		{[]rune{'a', 'b', 'c', 'd', 'f'}, 'e'},
		{[]rune{'O', 'Q', 'R', 'S'}, 'P'},
	}

	for _, s := range tests {
		got := FindMissingLetter(s.got)
		if got != s.want {
			t.Errorf("Expected \"%s\", but got \"%s\"", string(s.want), string(got))
		}
	}
}
