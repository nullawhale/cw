package cw

func FindOdd(seq []int) int {
	counts := map[int]int{}
	for _, num := range seq {
		counts[num] += 1
	}

	for key, val := range counts {
		if val%2 != 0 {
			return key
		}
	}
	return 0
}
