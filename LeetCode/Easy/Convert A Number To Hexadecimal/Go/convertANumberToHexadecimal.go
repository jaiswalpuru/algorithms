package main

import (
	"fmt"
	"strconv"
)

var numToLetter = map[int]string{
	10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f",
}

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	isNegative := num < 0
	positiveNum := uint32(abs(num))

	//two complement
	if isNegative {
		positiveNum = (^positiveNum) + 1
	}

	base16 := ""
	for positiveNum > 0 {
		remainder := int(positiveNum % 16)
		if remainder >= 10 {
			base16 = numToLetter[remainder] + base16
		} else {
			base16 = strconv.Itoa(remainder) + base16
		}
		positiveNum /= 16
	}

	return base16
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(toHex(26))
}
