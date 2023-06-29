package cw

import (
	"testing"
)

func TestToCamelCase(t *testing.T) {
	var tests = []struct {
		got  string
		want string
	}{
		{"", ""},
		{"The_Stealth_Warrior", "TheStealthWarrior"},
		{"the-Stealth-Warrior", "theStealthWarrior"},
	}

	for _, s := range tests {
		got := ToCamelCase(s.got)
		if got != s.want {
			t.Errorf("Expected \"%s\", but got \"%s\"", s.want, got)
		}
	}
}
