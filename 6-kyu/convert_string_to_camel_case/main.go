package cw

import (
	"strings"
)

func ToCamelCase(s string) (res string) {
	delimiter := ""
	if strings.Contains(s, "_") {
		delimiter = "_"
	} else {
		delimiter = "-"
	}
	if s == "" {
		return ""
	}
	words := strings.Split(s, delimiter)
	res += words[0]
	for _, word := range words[1:] {
		res += strings.ToUpper(string(word[0])) + word[1:]
	}
	return res
}
