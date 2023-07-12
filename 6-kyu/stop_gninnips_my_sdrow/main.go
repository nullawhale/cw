package cw

import (
	"strings"
)

func SpinWords(str string) string {
	words := strings.Fields(str)
	for i, word := range words {
		if len(word) >= 5 {
			temp := ""
			for _, v := range word {
				temp = string(v) + temp
			}
			words[i] = temp
		} else {
			words[i] = word
		}
	}

	return strings.Join(words, " ")
}
