package cw

import (
	"reflect"
	"strconv"
)

func SumMix(arr []any) (res int) {
	for _, elem := range arr {
		switch reflect.TypeOf(elem).Kind() {
		case reflect.Int:
			res += elem.(int)
		case reflect.String:
			num, _ := strconv.Atoi(elem.(string))
			res += num
		}
	}
	return res
}
