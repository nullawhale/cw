package cw

import (
	"sort"
)

func FindDupsMiss(arr []int) (missing int, dups []int) {
	counts := map[int]int{}
	max, min, sum := arr[0], arr[0], 0
	for _, num := range arr {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
		counts[num] += 1
		if counts[num] > 1 {
			dups = append(dups, num)
		}
	}

	for key := range counts {
		sum += key
	}
	sort.Ints(dups)
	missing = (max * (max + 1) / 2) - (min * (min - 1) / 2) - sum

	return missing, dups
}
