package cw

func IsValidWalk(walk []rune) bool {
	var returned struct{ x, y int }
	if len(walk) != 10 {
		return false
	}
	for _, direction := range walk {
		switch direction {
		case 'n':
			returned.x += 1
		case 's':
			returned.x -= 1
		case 'w':
			returned.y += 1
		case 'e':
			returned.y -= 1
		}
	}
	return returned.x == 0 && returned.y == 0
}
