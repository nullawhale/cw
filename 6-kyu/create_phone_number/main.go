package cw

import (
	"fmt"
	"strings"
)

func CreatePhoneNumber(numbers [10]uint) string {
	return fmt.Sprintf("(%s) %s-%s", ArrayToString(numbers[0:3]), ArrayToString(numbers[3:6]), ArrayToString(numbers[6:10]))
}

func ArrayToString(arr []uint) (res string) {
	return strings.Trim(strings.ReplaceAll(fmt.Sprint(arr), " ", ""), "[]")
}
