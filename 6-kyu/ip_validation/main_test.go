package cw

import (
	"testing"
)

func TestFindMissingLetter(t *testing.T) {
	var tests = []struct {
		got  string
		want bool
	}{
		{"12.255.56.1", true},
		{"", false},
		{"abc.def.ghi.jkl", false},
		{"123.456.789.0", false},
		{"12.34.56", false},
		{"12.34.56 .1", false},
		{"12.34.56.-1", false},
		{"123.045.067.089", false},
		{"127.1.1.0", true},
		{"0.0.0.0", true},
		{"0.34.82.53", true},
		{"192.168.1.300", false},
	}

	for _, s := range tests {
		got := Is_valid_ip(s.got)
		if got != s.want {
			t.Errorf("Expected (\"%s\") \"%t\", but got \"%t\"", s.got, s.want, got)
		}
	}
}
