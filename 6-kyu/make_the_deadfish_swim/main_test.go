package cw

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	var tests = []struct {
		got  string
		want []int
	}{
		{"ooo", []int{0, 0, 0}},
		{"ioioio", []int{1, 2, 3}},
		{"idoiido", []int{0, 1}},
		{"isoisoiso", []int{1, 4, 25}},
		{"codewars", []int{0}},
		{"djlkiswtdidiejddscdsd", []int{}},
	}

	for _, s := range tests {
		got := Parse(s.got)
		if !reflect.DeepEqual(got, s.want) {
			t.Errorf("Expected \"%v\", but got \"%v\"", s.want, got)
		}
	}
}
