package cw

import (
	"testing"
)

func TestSnakesAndLadders(t *testing.T) {
	var tests = []struct {
		dice  []int
		board []int
		want  int
	}{
		{[]int{2, 1, 5, 1, 5, 4}, []int{0, 0, 3, 0, 0, 0, 0, -2, 0, 0, 0}, 10},
	}
	for _, s := range tests {
		got := SnakesAndLadders(s.board, s.dice)
		if got != s.want {
			t.Errorf("Expected \"%d\", but got \"%d\"", s.want, got)
		}
	}
}
