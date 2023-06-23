package cw

func ChooseBestSum(t, k int, ls []int) int {

	res := -1
	for _, c := range allCombinations(k, ls) {
		if res < c && c <= t {
			res = c
		}
	}
	return res
}

func allCombinations(k int, ls []int) []int {
	if k == 1 {
		return ls
	}

	var res []int
	for i := 0; i < len(ls)-k+1; i++ {
		arr := allCombinations(k-1, ls[i+1:])
		for _, v := range arr {
			res = append(res, ls[i]+v)
		}
	}
	return res
}
