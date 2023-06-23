package cw

import (
	"testing"
)

func TestFakeBin(t *testing.T) {
	var tests = [][]string{
		{"45385593107843568", "01011110001100111"},
		{"509321967506747", "101000111101101"},
		{"366058562030849490134388085", "011011110000101010000011011"},
		{"15889923", "01111100"},
		{"800857237867", "100111001111"},
	}

	for i := 0; i < len(tests); i++ {
		got := FakeBin(tests[i][0])
		if got != tests[i][1] {
			t.Errorf("Expected %s, but got %s", tests[i][0], tests[i][1])
		}
	}
}
