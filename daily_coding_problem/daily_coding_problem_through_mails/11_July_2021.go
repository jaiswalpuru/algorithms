/*
Implement division of two positive integers without using the division, multiplication,
or modulus operators. Return the quotient as an integer, ignoring the remainder.
*/
package main

import (
	"fmt"
	"math"
)

func div(dividend, divisor int) int {

	//first check the sign of the dividend and divisor
	sign, dividend_sign, divisor_sign := 1, 0, 0
	if dividend < 0 {
		dividend_sign = 1
	}
	if divisor < 0 {
		divisor_sign = 1
	}

	if dividend_sign^divisor_sign == 1 {
		sign = -1
	}
	dividend, divisor = int(math.Abs(float64(dividend))), int(math.Abs(float64(divisor)))
	ans := 0
	for {
		dividend = dividend - divisor
		if dividend < 0 {
			return sign * ans
		}
		ans++
	}
}

func main() {
	fmt.Println(div(4, 2))
	fmt.Println(div(5, 2))
	fmt.Println(div(25, 2))
	fmt.Println(div(43, -8))
}
