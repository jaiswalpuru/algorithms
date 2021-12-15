/*
Given three 32-bit integers x, y, and b, return x if b is 1 and y if b is 0,
using only mathematical or bit operations.
You can assume b can only be 1 or 0.
*/

package main

import "fmt"

func return_b(x, y, b int) int {
	return x*b + y*(1-b)
}

func main() {
	fmt.Println(return_b(3, 8, 0))
	fmt.Println(return_b(3, 8, 1))
}
