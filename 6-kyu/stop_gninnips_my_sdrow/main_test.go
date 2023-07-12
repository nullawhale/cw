package cw

import (
	"testing"
)

func TestSpinWords(t *testing.T) {
	var tests = []struct {
		got  string
		want string
	}{
		{"Welcome", "emocleW"},
		{"to", "to"},
		{"CodeWars", "sraWedoC"},
		{"Hey fellow warriors", "Hey wollef sroirraw"},
		{"Burgers are my favorite fruit", "sregruB are my etirovaf tiurf"},
		{"Pizza is the best vegetable", "azziP is the best elbategev"},
	}

	for _, s := range tests {
		got := SpinWords(s.got)
		if got != s.want {
			t.Errorf("Expected \"%s\", but got \"%s\"", s.want, got)
		}
	}
}
