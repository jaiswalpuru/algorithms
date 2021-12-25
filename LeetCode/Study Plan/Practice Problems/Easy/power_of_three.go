/*

 */
package main

import (
	"fmt"
)

func power_of_four(n int) bool {
	if n == 1 {
		return true
	}
	if n <= 0 {
		return false
	}
	for n > 1 {
		if n%3 != 0 {
			return false
		}
		n = n / 3
		if n == 1 {
			return true
		}
	}
	return true
}

func main() {
	fmt.Println(power_of_four(9))
	fmt.Println(power_of_four(1))
	fmt.Println(power_of_four(0))
}
