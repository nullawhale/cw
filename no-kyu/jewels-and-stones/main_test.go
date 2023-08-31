package cw

import (
	"testing"
)

func TestNumJewelsInStones(t *testing.T) {
	var tests = []struct {
		jewels string
		stones string
		want   int
	}{
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}

	for _, s := range tests {
		got := NumJewelsInStones(s.jewels, s.stones)
		if got != s.want {
			t.Errorf("Expected \"%d\", but got \"%d\"", s.want, got)
		}
	}
}
