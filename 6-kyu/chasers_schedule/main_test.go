package cw

import (
	"testing"
)

func TestChaser(t *testing.T) {
	var tests = []struct {
		speed int
		time  int
		want  int
	}{
		{2, 4, 10},
		{2, 3, 8},
		{1, 1, 2},
	}
	for _, s := range tests {
		got := Chaser(s.speed, s.time)
		if got != s.want {
			t.Errorf("Expected \"%d\", but got \"%d\"", s.want, got)
		}
	}
}
