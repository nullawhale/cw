package cw

import (
	"testing"
)

func TestOrder(t *testing.T) {
	var tests = []struct {
		got  string
		want string
	}{
		{"is2 Thi1s T4est 3a", "Thi1s is2 3a T4est"},
		{"4of Fo1r pe6ople g3ood th5e the2", "Fo1r the2 g3ood 4of th5e pe6ople"},
		{"", ""},
	}
	for _, s := range tests {
		got := Order(s.got)
		if got != s.want {
			t.Errorf("Expected \"%s\", but got \"%s\"", s.want, got)
		}
	}
}
