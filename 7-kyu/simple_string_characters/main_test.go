package cw

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	var tests = []struct {
		got  string
		want []int
	}{
		{"Codewars@codewars123.com", []int{1, 18, 3, 2}},
		{"bgA5<1d-tOwUZTS8yQ", []int{7, 6, 3, 2}},
		{"P*K4%>mQUDaG$h=cx2?.Czt7!Zn16p@5H", []int{9, 9, 6, 9}},
		{"RYT'>s&gO-.CM9AKeH?,5317tWGpS<*x2ukXZD", []int{15, 8, 6, 9}},
		{"$Cnl)Sr<7bBW-&qLHI!mY41ODe", []int{10, 7, 3, 6}},
	}

	for _, s := range tests {
		got := Solve(s.got)
		if !reflect.DeepEqual(got, s.want) {
			t.Errorf("Expected \"%v\", but got \"%v\"", s.want, got)
		}
	}
}
