/*
Given a positive integer num, write a function which returns True if num is a perfect square else False.

Follow up: Do not use any built-in library function such as sqrt.
*/

package main

import "fmt"

func valid_square(n int) bool {
	if n == 1 {
		return true
	}

	l, h := 2, n/2
	for l <= h {
		mid := (l + h) / 2
		if mid*mid == n {
			return true
		}
		if mid*mid > n {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}

func main() {
	fmt.Println(valid_square(7))
	// fmt.Println(valid_square(81))
}
