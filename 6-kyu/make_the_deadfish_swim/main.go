package cw

func Parse(data string) []int {
	res := []int{}
	temp := 0
	for _, v := range data {
		switch v {
		case 'i':
			temp += 1
		case 'd':
			temp -= 1
		case 's':
			temp *= temp
		case 'o':
			res = append(res, temp)
		}
	}

	return res
}
