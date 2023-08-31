package cw

import (
	"testing"
)

func TestDeleteDigit(t *testing.T) {
	var tests = []struct {
		got  int
		want int
	}{
		{152, 52},
		{1001, 101},
		{10, 1},
	}
	for _, s := range tests {
		got := DeleteDigit(s.got)
		if got != s.want {
			t.Errorf("Expected \"%d\", but got \"%d\"", s.want, got)
		}
	}
}
