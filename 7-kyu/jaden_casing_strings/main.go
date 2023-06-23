package cw

import (
	"strings"
)

func ToJadenCase(str string) string {
	var jadenCase []string
	words := strings.Fields(str)
	for _, word := range words {
		jadenCase = append(jadenCase, strings.ToUpper(string(word[0]))+word[1:])
	}
	return strings.Join(jadenCase, " ")
}
