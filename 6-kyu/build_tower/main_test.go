package cw

import (
	"reflect"
	"testing"
)

func TestTowerBuilder(t *testing.T) {
	var tests = []struct {
		got  int
		want []string
	}{
		{1, []string{"*"}},
		{2, []string{" * ", "***"}},
		{3, []string{"  *  ", " *** ", "*****"}},
		{6, []string{"     *     ", "    ***    ", "   *****   ", "  *******  ", " ********* ", "***********"}},
	}
	for _, s := range tests {
		got := TowerBuilder(s.got)
		if !reflect.DeepEqual(got, s.want) {
			t.Errorf("Expected \"%s\", but got \"%s\"", s.want, got)
		}
	}
}
