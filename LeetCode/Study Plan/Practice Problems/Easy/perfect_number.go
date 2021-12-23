/*
A perfect number is a positive integer that is equal to the sum of its positive divisors, excluding the number itself.
A divisor of an integer x is an integer that can divide x evenly.

Given an integer n, return true if n is a perfect number, otherwise return false.
*/
package main

import "fmt"

func perfect_number(n int) bool {
	sum := 0
	for i := n / 2; i > 0; i-- {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}

func main() {
	fmt.Println(perfect_number(1))
	fmt.Println(perfect_number(7))
}
