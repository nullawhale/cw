package cw

import (
	"testing"
)

func TestDNAtoRNA(t *testing.T) {
	var tests = []struct {
		got  string
		want string
	}{
		{"GCAT", "GCAU"},
		{"GCTAADT", "GCUAADU"},
	}

	for _, s := range tests {
		got := DNAtoRNA(s.got)
		if got != s.want {
			t.Errorf("Expected \"%s\", but got \"%s\"", s.want, got)
		}
	}
}
