package main

import "fmt"

func countOdds(low int, high int) int {
	if low%2 == 0 {
		low++
	}
	if high%2 == 0 {
		high--
	}

	if low > high {
		return 0
	}

	return (high-low)/2 + 1
}

func main() {
	fmt.Println(countOdds(3, 7))
}
