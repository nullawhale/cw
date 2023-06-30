package cw

import (
	"testing"
)

func TestLongestConsec(t *testing.T) {
	var tests = []struct {
		arr  []string
		k    int
		want string
	}{
		{[]string{"tree", "foling", "trashy", "blue", "abcdef", "uvwxyz"}, 2, "folingtrashy"},
		{[]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2, "abigailtheta"},
		{[]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1,
			"oocccffuucccjjjkkkjyyyeehh"},
		{[]string{}, 3, ""},
		{[]string{"itvayloxrp", "wkppqsztdkmvcuwvereiupccauycnjutlv", "vweqilsfytihvrzlaodfixoyxvyuyvgpck"}, 2,
			"wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck"},
	}

	for _, s := range tests {
		got := LongestConsec(s.arr, s.k)
		if got != s.want {
			t.Errorf("Expected \"%s\", but got \"%s\"", s.want, got)
		}
	}
}
