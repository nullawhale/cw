package cw

func FindMissingLetter(chars []rune) rune {
	missing := chars[0]
	for _, char := range chars[1:] {
		missing++
		if missing != char {
			return missing
		}
	}

	return missing
}
