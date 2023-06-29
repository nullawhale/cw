package cw

import (
	"reflect"
	"testing"
)

func TestChooseBestSum(t *testing.T) {
	var tests = []struct {
		items []interface{}
		k     int
		want  []interface{}
	}{
		{[]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 1, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2, []interface{}{2, 4, 6, 8, 10, 3, 7, 1, 9, 5}},
		{[]interface{}{1, 2, 3, 4, 5, 6, 7}, 3, []interface{}{3, 6, 2, 7, 5, 1, 4}},
	}
	for _, s := range tests {
		got := Josephus(s.items, s.k)
		if !reflect.DeepEqual(got, s.want) {
			t.Errorf("Expected \"%d\", but got \"%d\"", s.want, got)
		}
	}
}
