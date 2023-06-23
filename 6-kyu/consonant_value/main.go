package cw

func Solve(str string) int {
	var max int
	var temp int
	for _, letter := range str {
		switch letter {
		case 'a', 'e', 'i', 'o', 'u':
			temp = 0
			continue
		}

		temp += int(letter - 96)

		if temp > max {
			max = temp
		}
	}
	return max
}
