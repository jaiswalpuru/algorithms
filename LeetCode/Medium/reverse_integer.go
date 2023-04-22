/*
Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the
signed 32-bit integer range [-231, 231 - 1], then return 0.

Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
*/

package main

import (
	"fmt"
	"math"
)

func reverse_integer(n int) int {
	flag := false
	if n < 0 {
		flag = true
		n *= -1
	}
	reverse_num := 0
	for n > 0 {
		reverse_num = reverse_num*10 + (n % 10)
		n = n / 10
		if reverse_num > int(math.MaxInt32) {
			return 0
		}
	}

	if flag {
		reverse_num *= -1
	}
	return reverse_num
}

func main() {
	fmt.Println(reverse_integer(1534236469))
	fmt.Println(reverse_integer(-321))
}
