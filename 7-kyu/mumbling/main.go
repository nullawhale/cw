package cw

import (
	"strings"
)

func Accum(s string) string {
	letters := make([]string, len(s))
	for i, letter := range s {
		letters[i] = strings.ToUpper(string(letter)) + strings.Repeat(strings.ToLower(string(letter)), i)
	}
	return strings.Join(letters, "-")
}
