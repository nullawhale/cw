package cw

import (
	"testing"
)

func TestGetMiddle(t *testing.T) {
	var tests = []struct {
		got  string
		want string
	}{
		{"test", "es"},
		{"testing", "t"},
		{"middle", "dd"},
		{"A", "A"},
	}

	for _, s := range tests {
		got := GetMiddle(s.got)
		if got != s.want {
			t.Errorf("Expected %s, but got %s", s.want, got)
		}
	}
}
