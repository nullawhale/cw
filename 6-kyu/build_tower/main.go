package cw

import (
	"strings"
)

func TowerBuilder(nFloors int) []string {
	var res []string
	lineLength := nFloors*2 - 1

	for i := 1; i <= nFloors; i++ {
		blockCount := i + (i - 1)
		res = append(res,
			strings.Repeat(" ", (lineLength-blockCount)/2)+
				strings.Repeat("*", blockCount)+
				strings.Repeat(" ", (lineLength-blockCount)/2))
	}
	return res
}
