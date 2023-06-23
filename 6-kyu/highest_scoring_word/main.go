package cw

import (
	"strings"
)

func High(s string) string {
	words := strings.Fields(s)
	var maxScore = 0
	var winner = ""
	for _, word := range words {
		score := 0
		for _, letter := range word {
			score += int(letter - 'a' + 1)
		}
		if score > maxScore {
			maxScore = score
			winner = word
		}
	}
	return winner
}
