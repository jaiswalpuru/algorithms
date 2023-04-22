/*
Given an integer n, return true if it is a power of two. Otherwise, return false.

An integer n is a power of two, if there exists an integer x such that n == 2x.
*/

package main

import (
	"fmt"
	"math"
)

func power_of_two(n int) bool {
	val := math.Log2(float64(n))
	if val-float64(int(val)) != 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(power_of_two(1))
	fmt.Println(power_of_two(2))
	fmt.Println(power_of_two(16))
	fmt.Println(power_of_two(3))
}
