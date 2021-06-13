/*
	Implement integer exponentiation. That is, implement the pow(x, y) function,
	where x and y are integers and returns x^y.

	Do this faster than the naive method of repeated multiplication.

	For example, pow(2, 10) should return 1024.
*/

package main

import "fmt"

func fast_exp(a, b int) int {
	if b == 0 {
		return 1
	}
	if b%2 == 0 {
		half_ans := fast_exp(a, b/2)
		return half_ans * half_ans
	} else {
		half_ans := fast_exp(a, b/2)
		return half_ans * half_ans * a
	}
}

func main() {
	fmt.Println(fast_exp(2, 10))
}
