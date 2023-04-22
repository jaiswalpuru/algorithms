/*
Given a non-negative integer c, decide whether there're two integers a and b such that a2 + b2 = c.
*/

package main

import (
	"fmt"
	"math"
)

func square_sum(n int) bool {
	if n <= 2 {
		return true
	}
	m := int(math.Sqrt(float64(n)))
	a, b := 0, m
	for a <= b {
		val := a*a + b*b
		if val == n {
			return true
		}
		if val > n {
			b -= 1
		} else {
			a += 1
		}
	}
	return false
}

func main() {
	fmt.Println(square_sum(8))
}
