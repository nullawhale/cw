package cw

import (
	"fmt"
)

func Chaser(speed, time int) (res int) {
	max := 0
	for i := time; i >= 0; i-- {
		res = 0
		//tempSpeed := speed
		var str string
		for j := 0; j < time; j++ {
			if i == j {
				res += speed * 2
				str += "s*2 + "
			} else if i < j {
				res += speed - 1
				str += "(s-1) + "
			} else {
				res += speed
				str += "s + "
			}
		}
		//fmt.Printf("speed: %d: res: %d;\n", tempSpeed, res)
		fmt.Printf("%s = %d\n", str, res)
		if res > max {
			max = res
		}
	}
	fmt.Println()
	return max
}
