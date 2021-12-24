/*
Given an integer n, return true if it is a power of four. Otherwise, return false.

An integer n is a power of four, if there exists an integer x such that n == 4x.
*/

package main

import (
	"fmt"
	"math"
)

func power_of_four(n int) bool {
	num := math.Log2(float64(n)) / 2

	if num-float64(int(num)) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(power_of_four(16))
	fmt.Println(power_of_four(1))
	fmt.Println(power_of_four(5))
}
