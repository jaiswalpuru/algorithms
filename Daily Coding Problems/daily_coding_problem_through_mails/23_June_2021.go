/*
	A number is considered perfect if its digits sum up to exactly 10.

	Given a positive integer n, return the n-th perfect number.

	For example, given 1, you should return 19. Given 2, you should return 28.
*/

package main

import "fmt"

func perfect_number(n int) int {
	temp := 19
	if n == 1 {
		return temp
	}

	for i := 2; i <= n; i++ {
		temp += 9
	}
	return temp
}

func main() {
	fmt.Println(perfect_number(1))
	fmt.Println(perfect_number(2))
	fmt.Println(perfect_number(3))
	fmt.Println(perfect_number(4))
	fmt.Println(perfect_number(5))
}
