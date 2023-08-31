package cw

import "unicode"

func Solve(s string) []int {
	res := []int{0, 0, 0, 0}
	for _, letter := range s {
		switch {
		case unicode.IsUpper(letter):
			res[0] += 1
		case unicode.IsLower(letter):
			res[1] += 1
		case unicode.IsDigit(letter):
			res[2] += 1
		default:
			res[3] += 1
		}
	}
	return res
}
