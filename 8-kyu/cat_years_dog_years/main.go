package cw

func CalculateYears(years int) (result [3]int) {
	catYears := 15
	dogYears := 15
	if years >= 2 {
		catYears += 9 + (years-2)*4
		dogYears += 9 + (years-2)*5
	}
	return [3]int{years, catYears, dogYears}
}
