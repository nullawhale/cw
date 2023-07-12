package cw

import (
	"strings"
	"unicode"
)

func Solve(str string) string {
	l, u := 0, 0
	for _, letter := range str {
		if unicode.IsUpper(letter) {
			u++
		} else {
			l++
		}
	}

	if l >= u {
		return strings.ToLower(str)
	}
	return strings.ToUpper(str)
}
