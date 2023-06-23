package cw

import (
	"testing"
)

func TestCreatePhoneNumber(t *testing.T) {
	var tests = []struct {
		got  [10]uint
		want string
	}{
		{[10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, "(123) 456-7890"},
	}

	for _, s := range tests {
		got := CreatePhoneNumber(s.got)
		if got != s.want {
			t.Errorf("Expected %s, but got %s", s.want, got)
		}
	}
}
