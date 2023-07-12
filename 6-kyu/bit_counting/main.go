package cw

func CountBits(num uint) (res int) {
	for num > 0 {
		if num%2 == 1 {
			res++
		}
		num /= 2
	}
	return int(res)
}
