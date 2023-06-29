package cw

func Disemvowel(comment string) (res string) {

	for _, letter := range comment {
		switch letter {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			break
		default:
			res += string(letter)
		}
	}

	return res
}
