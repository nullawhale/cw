package cw

func Smaller(arr []int) []int {
	result := make([]int, len(arr))
	for i, elem := range arr {
		count := 0
		for _, comp := range arr[i+1:] {
			if comp < elem {
				count++
			}
		}
		result[i] = count
	}
	return result
}
