package cw

func LongestConsec(strarr []string, k int) (res string) {
	for i := 0; i <= len(strarr)-k; i++ {
		maxStr := ""
		for j := 0; j < k; j++ {
			maxStr += strarr[j+i]
		}
		if len(maxStr) > len(res) {
			res = maxStr
		}
	}

	if len(strarr) == 0 || k > len(strarr) || k <= 0 {
		return ""
	}
	return res
}
