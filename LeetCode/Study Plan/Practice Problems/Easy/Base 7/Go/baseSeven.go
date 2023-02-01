package main

import (
	"fmt"
	"strconv"
)

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	absValue := abs(num)

	base7 := ""
	for absValue > 0 {
		base7 = strconv.Itoa(absValue%7) + base7
		absValue /= 7
	}

	if isNegative(num) {
		base7 = "-" + base7
	}
	return base7
}

// return whether the number is negative or not.
func isNegative(n int) bool { return n < 0 }

// return the absolute value
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	fmt.Println(convertToBase7(100))
}
