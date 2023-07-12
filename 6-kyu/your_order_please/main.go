package cw

import (
	"strconv"
	"strings"
	"unicode"
)

func Order(sentence string) string {
	words := strings.Fields(sentence)
	res := make([]string, len(words))

	for _, word := range words {
		for _, letter := range word {
			if unicode.IsDigit(letter) {
				index, _ := strconv.Atoi(string(letter))
				res[index-1] = word
				break
			}
		}
	}

	return strings.Join(res, " ")
}
