package cw

import (
	"strings"
)

func NumJewelsInStones(jewels string, stones string) (num int) {
	for _, stone := range stones {
		if strings.ContainsRune(jewels, stone) {
			num++
		}
	}
	return num
}
